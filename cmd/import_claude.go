package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gandazgul/mnemoteca/internal/chunker"
	"github.com/gandazgul/mnemoteca/internal/config"
	"github.com/gandazgul/mnemoteca/internal/db"
	"github.com/spf13/cobra"
)

type claudeMemorySource struct {
	Path  string
	Scope string
}

const claudeImportNonDestructiveMessage = `Non-destructive import: Mnemoteca will only read Claude Code memory files and append them as new Mnemoteca documents. It will not edit, delete, move, overwrite, or truncate any Claude Code files or existing Mnemoteca memories.`

func importClaudeAgent(cmd *cobra.Command, opts importAgentOptions) error {
	collectionName, err := resolveCollectionName(opts.Name, opts.Global)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting current directory: %w", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("getting home directory: %w", err)
	}

	sources, err := discoverClaudeMemorySources(cwd, home, opts.Paths, opts.IncludeUser)
	if err != nil {
		return err
	}

	cmd.Println(claudeImportNonDestructiveMessage)
	if opts.DryRun {
		cmd.Println("Dry run: no database writes or embedding calls will be made.")
	}
	cmd.Println()

	if len(sources) == 0 {
		cmd.Println("No Claude Code memory files found.")
		return nil
	}

	cmd.Printf("Found %d Claude Code memory file(s):\n", len(sources))
	for _, source := range sources {
		cmd.Printf("  %s  %s\n", source.Scope, source.Path)
	}

	if opts.DryRun {
		return nil
	}

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	database, err := openDB()
	if err != nil {
		return err
	}
	defer database.Close() //nolint:errcheck

	if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
		return fmt.Errorf("ensuring vector table: %w", err)
	}

	collection, created, err := database.GetOrCreateCollection(collectionName)
	if err != nil {
		return fmt.Errorf("getting or creating collection %q: %w", collectionName, err)
	}
	if created {
		cmd.Printf("\nCreated collection %q.\n", collectionName)
	} else {
		cmd.Printf("\nAppending to existing collection %q.\n", collectionName)
	}

	embedder, _, err := openEmbedder(cmd.Context())
	if err != nil {
		return fmt.Errorf("loading embedding model: %w", err)
	}
	defer embedder.Close() //nolint:errcheck

	var totalChunks int
	var importedFiles int
	for _, source := range sources {
		n, err := importClaudeMemorySource(database, collection.ID, source, func(content string) ([]float32, error) {
			return embedder.EmbedDocument(content)
		})
		if err != nil {
			return err
		}
		if n == 0 {
			cmd.Printf("Skipped empty file: %s\n", source.Path)
			continue
		}
		importedFiles++
		totalChunks += n
		cmd.Printf("Imported %d document(s) from %s\n", n, source.Path)
	}

	cmd.Printf("\nDone. Imported %d Claude Code file(s) as %d Mnemoteca document(s).\n", importedFiles, totalChunks)
	return nil
}

func discoverClaudeMemorySources(cwd, home string, explicitPaths []string, includeUser bool) ([]claudeMemorySource, error) {
	var sources []claudeMemorySource
	seen := map[string]bool{}

	addSource := func(path, scope string) error {
		expanded, err := expandHome(path, home)
		if err != nil {
			return err
		}
		cleaned, err := filepath.Abs(expanded)
		if err != nil {
			return fmt.Errorf("resolving %s: %w", path, err)
		}
		info, err := os.Stat(cleaned)
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
			return fmt.Errorf("checking %s: %w", cleaned, err)
		}
		if info.IsDir() {
			return collectClaudeMarkdownFiles(cleaned, scope, seen, &sources)
		}
		if strings.EqualFold(filepath.Ext(cleaned), ".md") && !seen[cleaned] {
			seen[cleaned] = true
			sources = append(sources, claudeMemorySource{Path: cleaned, Scope: scope})
		}
		return nil
	}

	if len(explicitPaths) > 0 {
		for _, path := range explicitPaths {
			if strings.TrimSpace(path) == "" {
				continue
			}
			if err := addSource(path, "claude-explicit"); err != nil {
				return nil, err
			}
		}
		sortClaudeMemorySources(sources)
		return sources, nil
	}

	projectFiles := []string{
		filepath.Join(cwd, "CLAUDE.md"),
		filepath.Join(cwd, ".claude", "CLAUDE.md"),
		filepath.Join(cwd, "CLAUDE.local.md"),
	}
	for _, path := range projectFiles {
		if err := addSource(path, "claude-project"); err != nil {
			return nil, err
		}
	}
	if err := addSource(filepath.Join(cwd, ".claude", "rules"), "claude-project-rule"); err != nil {
		return nil, err
	}

	for _, root := range candidateClaudeProjectRoots(cwd) {
		memoryDir := filepath.Join(home, ".claude", "projects", claudeProjectDirName(root), "memory")
		if err := addSource(memoryDir, "claude-auto-memory"); err != nil {
			return nil, err
		}
	}

	if includeUser {
		if err := addSource(filepath.Join(home, ".claude", "CLAUDE.md"), "claude-user"); err != nil {
			return nil, err
		}
		if err := addSource(filepath.Join(home, ".claude", "rules"), "claude-user-rule"); err != nil {
			return nil, err
		}
	}

	sortClaudeMemorySources(sources)
	return sources, nil
}

func importClaudeMemorySource(database *db.DB, collectionID int64, source claudeMemorySource, embed func(string) ([]float32, error)) (int, error) {
	data, err := os.ReadFile(source.Path)
	if err != nil {
		return 0, fmt.Errorf("reading %s: %w", source.Path, err)
	}

	content := strings.TrimSpace(string(data))
	if content == "" {
		return 0, nil
	}

	var chunks []string
	for _, semanticChunk := range chunker.ChunkDocument([]byte(content), 2000) {
		chunk := strings.TrimSpace(semanticChunk.Content)
		if chunk != "" {
			chunks = append(chunks, chunk)
		}
	}
	if len(chunks) == 0 {
		chunks = []string{content}
	}

	metadataMap := map[string]interface{}{
		"source":       "claude-code",
		"source_path":  source.Path,
		"source_scope": source.Scope,
		"tags":         []string{"claude-code", "imported", source.Scope},
	}
	metadataBytes, err := json.Marshal(metadataMap)
	if err != nil {
		return 0, fmt.Errorf("encoding metadata for %s: %w", source.Path, err)
	}
	metadata := string(metadataBytes)

	for i, chunk := range chunks {
		vec, err := embed(chunk)
		if err != nil {
			return i, fmt.Errorf("embedding %s chunk %d: %w", source.Path, i+1, err)
		}
		if _, err := database.InsertDocumentWithVector(collectionID, chunk, &metadata, vec); err != nil {
			return i, fmt.Errorf("inserting %s chunk %d: %w", source.Path, i+1, err)
		}
	}

	return len(chunks), nil
}

func collectClaudeMarkdownFiles(dir, scope string, seen map[string]bool, sources *[]claudeMemorySource) error {
	return filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking %s: %w", path, err)
		}
		if entry.IsDir() {
			return nil
		}
		if !strings.EqualFold(filepath.Ext(path), ".md") {
			return nil
		}
		cleaned, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("resolving %s: %w", path, err)
		}
		if seen[cleaned] {
			return nil
		}
		seen[cleaned] = true
		*sources = append(*sources, claudeMemorySource{Path: cleaned, Scope: scope})
		return nil
	})
}

func candidateClaudeProjectRoots(cwd string) []string {
	cleaned, err := filepath.Abs(cwd)
	if err != nil {
		cleaned = filepath.Clean(cwd)
	}

	roots := []string{cleaned}
	if gitRoot, ok := findGitRoot(cleaned); ok && gitRoot != cleaned {
		roots = append(roots, gitRoot)
	}
	return roots
}

func findGitRoot(start string) (string, bool) {
	dir := filepath.Clean(start)
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
}

func claudeProjectDirName(path string) string {
	cleaned := filepath.Clean(path)
	slashed := filepath.ToSlash(cleaned)
	replaced := strings.ReplaceAll(slashed, "/", "-")
	replaced = strings.ReplaceAll(replaced, ":", "-")
	return replaced
}

func expandHome(path, home string) (string, error) {
	if path == "~" {
		return home, nil
	}
	if strings.HasPrefix(path, "~/") {
		return filepath.Join(home, strings.TrimPrefix(path, "~/")), nil
	}
	return path, nil
}

func sortClaudeMemorySources(sources []claudeMemorySource) {
	sort.Slice(sources, func(i, j int) bool {
		if sources[i].Scope == sources[j].Scope {
			return sources[i].Path < sources[j].Path
		}
		return sources[i].Scope < sources[j].Scope
	})
}
