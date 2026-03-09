package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/spf13/cobra"
)

// addCmd stores a document in the specified collection.
var addCmd = &cobra.Command{
	Use:   "add [text]",
	Short: "Add a document to a collection",
	Long: `Store a document in the specified collection.

Text can be provided as a positional argument, read from a file with --file,
or piped via stdin with --stdin.

The document is embedded using the configured ONNX model and stored alongside
its vector representation for semantic search.

The collection must already exist (use 'mnemosyne init' first).
If --name is not provided, the current directory name is used.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		fileFlag, _ := cmd.Flags().GetString("file")
		stdinFlag, _ := cmd.Flags().GetBool("stdin")

		// Determine the content to store.
		var rawContent string

		switch {
		case fileFlag != "":
			data, err := os.ReadFile(fileFlag)
			if err != nil {
				return fmt.Errorf("reading file %s: %w", fileFlag, err)
			}
			rawContent = strings.TrimSpace(string(data))

		case stdinFlag:
			var lines []string
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("reading stdin: %w", err)
			}
			rawContent = strings.TrimSpace(strings.Join(lines, "\n"))

		case len(args) > 0:
			rawContent = strings.TrimSpace(strings.Join(args, " "))

		default:
			return fmt.Errorf("no content provided; pass text as argument, use --file, or use --stdin")
		}

		if rawContent == "" {
			return fmt.Errorf("content is empty")
		}

		// Split content into chunks if it came from file or stdin.
		var chunks []string
		if fileFlag != "" || stdinFlag {
			// Basic chunking: split by double newlines (paragraphs)
			parts := strings.Split(rawContent, "\n\n")
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if p != "" {
					// If a chunk is still very large, further split by single newlines.
					// This is a naive approach; a proper tokenizer-based chunker would be better.
					if len(p) > 2000 {
						lines := strings.Split(p, "\n")
						var currentChunk strings.Builder
						for _, line := range lines {
							line = strings.TrimSpace(line)
							if line == "" {
								continue
							}
							if currentChunk.Len()+len(line) > 2000 && currentChunk.Len() > 0 {
								chunks = append(chunks, currentChunk.String())
								currentChunk.Reset()
							}
							if currentChunk.Len() > 0 {
								currentChunk.WriteString(" ")
							}
							currentChunk.WriteString(line)
						}
						if currentChunk.Len() > 0 {
							chunks = append(chunks, currentChunk.String())
						}
					} else {
						chunks = append(chunks, p)
					}
				}
			}
		} else {
			chunks = []string{rawContent}
		}

		if len(chunks) == 0 {
			return fmt.Errorf("no valid content after chunking")
		}

		// Resolve collection.
		collectionName, err := resolveCollectionName(nameFlag)
		if err != nil {
			return err
		}

		// Load config (needed for embedder and vector table dimensions).
		cfg := config.Load()

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close()

		// Ensure the vector table exists with the configured embedding dimensions.
		if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
			return fmt.Errorf("ensuring vector table: %w", err)
		}

		collection, err := database.GetCollectionByName(collectionName)
		if err != nil {
			return fmt.Errorf("looking up collection: %w", err)
		}
		if collection == nil {
			return fmt.Errorf("collection %q does not exist; run 'mnemosyne init --name %s' first",
				collectionName, collectionName)
		}

		// Initialize the embedder to generate a vector for this document.
		embedder, err := openEmbedder(cmd.Context(), cfg)
		if err != nil {
			return fmt.Errorf("loading embedding model: %w", err)
		}
		defer embedder.Close()

		for i, chunk := range chunks {
			// Generate the document embedding using the document prefix.
			vec, err := embedder.EmbedDocument(chunk)
			if err != nil {
				return fmt.Errorf("embedding document chunk %d: %w", i+1, err)
			}

			// Insert the document and its vector atomically.
			doc, err := database.InsertDocumentWithVector(collection.ID, chunk, nil, vec)
			if err != nil {
				return fmt.Errorf("adding document chunk %d: %w", i+1, err)
			}

			// Show a preview: first 80 characters.
			preview := chunk
			if len(preview) > 80 {
				preview = strings.ReplaceAll(preview[:80], "\n", " ") + "..."
			}

			cmd.Printf("Added document %d to collection %q (embedded %d dims)\n",
				doc.ID, collectionName, len(vec))
			cmd.Printf("  %s\n", preview)
		}

		return nil
	},
}

func init() {
	addCmd.Flags().String("name", "", "collection name (defaults to current directory name)")
	addCmd.Flags().String("file", "", "read content from a file")
	addCmd.Flags().Bool("stdin", false, "read content from stdin")
	rootCmd.AddCommand(addCmd)
}
