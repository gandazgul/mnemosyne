package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gandazgul/mnemosyne/internal/backup"
	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/db"
	"github.com/spf13/cobra"
)

// exportCmd exports collections to JSONL files for backup.
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export collections to JSONL files",
	Long: `Export one or more collections to JSONL files for backup and transfer.

Each exported file contains a header line with metadata followed by one line
per document, including the raw vector embedding. This makes imports fast
and model-independent (no re-embedding required).

Use --no-embeddings to exclude vector data from the export. This produces
much smaller files but requires re-embedding on import.

Examples:
  mnemosyne export --name my-project              # → my-project.jsonl
  mnemosyne export --name my-project -o backup.jsonl
  mnemosyne export --name my-project --no-embeddings
  mnemosyne export --all                           # → one .jsonl per collection
  mnemosyne export --all -o ./backups/             # → ./backups/<name>.jsonl`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		globalFlag, _ := cmd.Flags().GetBool("global")
		allFlag, _ := cmd.Flags().GetBool("all")
		outputFlag, _ := cmd.Flags().GetString("output")
		yesFlag, _ := cmd.Flags().GetBool("yes")
		noEmbeddings, _ := cmd.Flags().GetBool("no-embeddings")

		if allFlag && (nameFlag != "" || globalFlag) {
			return fmt.Errorf("cannot use --all with --name or --global")
		}

		// Ensure vector table exists (needed for the JOIN query).
		cfg := config.Load()

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close() //nolint:errcheck

		if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
			return fmt.Errorf("ensuring vector table: %w", err)
		}

		if allFlag {
			return exportAll(cmd, database, outputFlag, yesFlag, noEmbeddings)
		}

		var collectionName string
		if allFlag {
			// When --all is used, we don't need a collection name
			collectionName = ""
		} else {
			collectionName, err = resolveCollectionName(nameFlag, globalFlag)
			if err != nil {
				return err
			}
		}

		return exportSingleCollection(cmd, database, collectionName, outputFlag, noEmbeddings)
	},
}

// exportSingleCollection exports one collection to a JSONL file.
func exportSingleCollection(cmd *cobra.Command, database *db.DB, collectionName, outputFlag string, noEmbeddings bool) error {
	outputPath := outputFlag
	if outputPath == "" {
		outputPath = collectionName + ".jsonl"
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating output file %s: %w", outputPath, err)
	}
	defer f.Close() //nolint:errcheck

	w := bufio.NewWriter(f)

	count, err := backup.ExportCollection(w, database, collectionName, noEmbeddings)
	if err != nil {
		os.Remove(outputPath) //nolint:errcheck // best-effort cleanup
		return err
	}

	if err := w.Flush(); err != nil {
		os.Remove(outputPath) //nolint:errcheck // best-effort cleanup
		return fmt.Errorf("flushing output: %w", err)
	}

	cmd.Printf("Exported %d documents from %q → %s\n", count, collectionName, outputPath)
	return nil
}

// exportAll exports every collection to individual JSONL files.
func exportAll(cmd *cobra.Command, database *db.DB, outputFlag string, skipConfirm, noEmbeddings bool) error {
	collections, err := database.ListCollections()
	if err != nil {
		return fmt.Errorf("listing collections: %w", err)
	}

	if len(collections) == 0 {
		cmd.Println("No collections to export.")
		return nil
	}

	// Calculate totals for confirmation.
	var totalDocs int64
	for _, c := range collections {
		totalDocs += c.DocumentCount
	}

	if !skipConfirm {
		cmd.Printf("Export all? %d collections, %d documents. Continue? [y/N] ",
			len(collections), totalDocs)

		input, err := readInput(cmd)
		if err != nil && err.Error() != "EOF" {
			return err
		}

		input = strings.TrimSpace(input)
		input = strings.TrimRight(input, "\r")
		input = strings.ToLower(input)
		if input != "y" && input != "yes" {
			cmd.Println("Aborted.")
			return nil
		}
	}

	// Determine output directory.
	outDir := outputFlag
	if outDir == "" {
		outDir = "."
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("creating output directory %s: %w", outDir, err)
	}

	for _, c := range collections {
		outputPath := filepath.Join(outDir, c.Name+".jsonl")
		f, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("creating output file %s: %w", outputPath, err)
		}

		w := bufio.NewWriter(f)
		count, err := backup.ExportCollection(w, database, c.Name, noEmbeddings)
		if err != nil {
			f.Close() //nolint:errcheck
			return err
		}

		if err := w.Flush(); err != nil {
			f.Close() //nolint:errcheck
			return fmt.Errorf("flushing output: %w", err)
		}
		f.Close() //nolint:errcheck

		cmd.Printf("Exported %d documents from %q → %s\n", count, c.Name, outputPath)
	}

	cmd.Printf("\nDone. Exported %d collections (%d documents total).\n", len(collections), totalDocs)
	return nil
}

func init() {
	exportCmd.Flags().StringP("name", "n", "", "collection name to export")
	exportCmd.Flags().BoolP("global", "g", false, "export the global collection")
	exportCmd.Flags().Bool("all", false, "export all collections")
	exportCmd.Flags().StringP("output", "o", "", "output file or directory path")
	exportCmd.Flags().Bool("yes", false, "skip confirmation prompt for --all")
	exportCmd.Flags().Bool("no-embeddings", false, "exclude vector embeddings from export")
	rootCmd.AddCommand(exportCmd)
}
