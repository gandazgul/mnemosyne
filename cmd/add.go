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
		var content string

		switch {
		case fileFlag != "":
			data, err := os.ReadFile(fileFlag)
			if err != nil {
				return fmt.Errorf("reading file %s: %w", fileFlag, err)
			}
			content = strings.TrimSpace(string(data))

		case stdinFlag:
			var lines []string
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("reading stdin: %w", err)
			}
			content = strings.TrimSpace(strings.Join(lines, "\n"))

		case len(args) > 0:
			content = strings.TrimSpace(strings.Join(args, " "))

		default:
			return fmt.Errorf("no content provided; pass text as argument, use --file, or use --stdin")
		}

		if content == "" {
			return fmt.Errorf("content is empty")
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
		embedder, err := openEmbedder(cfg)
		if err != nil {
			return fmt.Errorf("loading embedding model: %w", err)
		}
		defer embedder.Close()

		// Generate the document embedding using the document prefix.
		vec, err := embedder.EmbedDocument(content)
		if err != nil {
			return fmt.Errorf("embedding document: %w", err)
		}

		// Insert the document.
		doc, err := database.InsertDocument(collection.ID, content, nil)
		if err != nil {
			return fmt.Errorf("adding document: %w", err)
		}

		// Store the embedding vector alongside the document.
		if err := database.InsertVector(doc.ID, collection.ID, vec); err != nil {
			return fmt.Errorf("storing embedding: %w", err)
		}

		// Show a preview: first 80 characters.
		preview := content
		if len(preview) > 80 {
			preview = preview[:80] + "..."
		}

		cmd.Printf("Added document %d to collection %q (embedded %d dims)\n",
			doc.ID, collectionName, len(vec))
		cmd.Printf("  %s\n", preview)

		return nil
	},
}

func init() {
	addCmd.Flags().String("name", "", "collection name (defaults to current directory name)")
	addCmd.Flags().String("file", "", "read content from a file")
	addCmd.Flags().Bool("stdin", false, "read content from stdin")
	rootCmd.AddCommand(addCmd)
}
