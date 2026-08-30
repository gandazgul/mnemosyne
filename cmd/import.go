package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gandazgul/mnemoteca/internal/backup"
	"github.com/gandazgul/mnemoteca/internal/config"
	"github.com/gandazgul/mnemoteca/internal/db"
	"github.com/gandazgul/mnemoteca/internal/embedding"
	"github.com/spf13/cobra"
)

// importCmd imports collections from JSONL files.
var importCmd = &cobra.Command{
	Use:   "import <file.jsonl>",
	Short: "Import collections from JSONL files or agent memory",
	Long: `Import one or more collections from JSONL files exported by 'mnemoteca export'.

When the file includes vector embeddings, the import is fast and model-independent
(no re-embedding required).

When importing a file exported with --no-embeddings, the import command will
automatically generate embeddings using the configured embedder. This requires
the embedding model to be available (it is auto-downloaded on first use).

If the collection already exists, documents are appended to it.

Use --agent to import memories from another agent's memory system. Agent imports
append new Mnemoteca documents and do not modify source agent memory files.

Examples:
  mnemoteca import my-project.jsonl                # import into original collection name
  mnemoteca import my-project.jsonl --name other    # override collection name
  mnemoteca import --dir ./backups/                 # import all .jsonl files from directory
  mnemoteca import --agent claude --dry-run         # preview Claude Code memories
  mnemoteca import --agent claude                   # import Claude Code memories`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		dirFlag, _ := cmd.Flags().GetString("dir")
		agentFlag, _ := cmd.Flags().GetString("agent")
		globalFlag, _ := cmd.Flags().GetBool("global")
		dryRunFlag, _ := cmd.Flags().GetBool("dry-run")
		includeUserFlag, _ := cmd.Flags().GetBool("include-user")
		pathFlags, _ := cmd.Flags().GetStringSlice("path")

		if agentFlag != "" {
			if len(args) > 0 {
				return fmt.Errorf("cannot use --agent with a file argument; use --path for agent-specific source paths")
			}
			if dirFlag != "" {
				return fmt.Errorf("cannot use --agent with --dir")
			}
			return importAgent(cmd, agentFlag, importAgentOptions{
				Name:        nameFlag,
				Global:      globalFlag,
				DryRun:      dryRunFlag,
				IncludeUser: includeUserFlag,
				Paths:       pathFlags,
			})
		}

		if globalFlag || dryRunFlag || includeUserFlag || hasNonEmptyPath(pathFlags) {
			return fmt.Errorf("--global, --dry-run, --include-user, and --path require --agent")
		}

		if dirFlag != "" && len(args) > 0 {
			return fmt.Errorf("cannot use --dir with a file argument")
		}
		if dirFlag != "" && nameFlag != "" {
			return fmt.Errorf("cannot use --dir with --name (each file uses its own collection name)")
		}
		if dirFlag == "" && len(args) == 0 {
			return fmt.Errorf("specify a file to import or use --dir")
		}

		// Ensure vector table exists.
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

		// Create a lazy embedder that initializes ONNX Runtime and loads the
		// model only when a document without a vector is encountered.
		var (
			embedOnce   sync.Once
			embedder    embedding.Embedder
			embedderErr error
		)
		lazyEmbedFn := func(content string) ([]float32, error) {
			embedOnce.Do(func() {
				embedder, _, embedderErr = openEmbedder(context.Background())
			})
			if embedderErr != nil {
				return nil, embedderErr
			}
			return embedder.EmbedDocument(content)
		}
		// Schedule cleanup of embedder if it was initialized.
		defer func() {
			if embedder != nil {
				embedder.Close() //nolint:errcheck
			}
		}()

		if dirFlag != "" {
			return importDir(cmd, database, dirFlag, lazyEmbedFn)
		}

		return importFile(cmd, database, args[0], nameFlag, lazyEmbedFn)
	},
}

// importFile imports a single JSONL file into the database.
func importFile(cmd *cobra.Command, database *db.DB, filePath, overrideName string, embedFn backup.EmbedFunc) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("opening file %s: %w", filePath, err)
	}
	defer f.Close() //nolint:errcheck

	header, count, err := backup.ImportCollection(f, database, overrideName, embedFn)
	if err != nil {
		return err
	}

	collectionName := overrideName
	if collectionName == "" && header != nil {
		collectionName = header.Collection
	}

	cmd.Printf("Imported %d documents into %q from %s\n", count, collectionName, filePath)
	return nil
}

// importDir imports all .jsonl files from a directory.
func importDir(cmd *cobra.Command, database *db.DB, dirPath string, embedFn backup.EmbedFunc) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("reading directory %s: %w", dirPath, err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(strings.ToLower(e.Name()), ".jsonl") {
			files = append(files, filepath.Join(dirPath, e.Name()))
		}
	}

	if len(files) == 0 {
		cmd.Println("No .jsonl files found in directory.")
		return nil
	}

	var totalImported int64
	for _, filePath := range files {
		f, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("opening file %s: %w", filePath, err)
		}

		header, count, err := backup.ImportCollection(f, database, "", embedFn)
		f.Close() //nolint:errcheck

		if err != nil {
			return fmt.Errorf("importing %s: %w", filePath, err)
		}

		collectionName := ""
		if header != nil {
			collectionName = header.Collection
		}
		cmd.Printf("Imported %d documents into %q from %s\n", count, collectionName, filePath)
		totalImported += count
	}

	cmd.Printf("\nDone. Imported %d files (%d documents total).\n", len(files), totalImported)
	return nil
}

func init() {
	importCmd.Flags().StringP("name", "n", "", "override collection name")
	importCmd.Flags().StringP("dir", "d", "", "import all .jsonl files from directory")
	importCmd.Flags().String("agent", "", "import memories from an agent memory system (supported: claude)")
	importCmd.Flags().BoolP("global", "g", false, "use the global collection for agent imports")
	importCmd.Flags().Bool("dry-run", false, "preview an agent import without writing to Mnemoteca")
	importCmd.Flags().Bool("include-user", false, "include user-level memories for agent imports")
	importCmd.Flags().StringSlice("path", nil, "specific agent memory file or directory to import (repeatable)")
	rootCmd.AddCommand(importCmd)
}

type importAgentOptions struct {
	Name        string
	Global      bool
	DryRun      bool
	IncludeUser bool
	Paths       []string
}

func importAgent(cmd *cobra.Command, agent string, opts importAgentOptions) error {
	switch strings.ToLower(agent) {
	case "claude", "claude-code":
		return importClaudeAgent(cmd, opts)
	default:
		return fmt.Errorf("unsupported import agent %q (supported: claude)", agent)
	}
}

func hasNonEmptyPath(paths []string) bool {
	for _, path := range paths {
		if strings.TrimSpace(path) != "" {
			return true
		}
	}
	return false
}
