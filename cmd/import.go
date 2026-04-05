package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gandazgul/mnemosyne/internal/backup"
	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/db"
	"github.com/spf13/cobra"
)

// importCmd imports collections from JSONL files.
var importCmd = &cobra.Command{
	Use:   "import <file.jsonl>",
	Short: "Import collections from JSONL files",
	Long: `Import one or more collections from JSONL files exported by 'mnemosyne export'.

The import is fast because raw vector embeddings are included in the file,
so no re-embedding is required.

If the collection already exists, documents are appended to it.

Examples:
  mnemosyne import my-project.jsonl                # import into original collection name
  mnemosyne import my-project.jsonl --name other    # override collection name
  mnemosyne import --dir ./backups/                 # import all .jsonl files from directory`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		dirFlag, _ := cmd.Flags().GetString("dir")

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
		cfg := config.Load()

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close() //nolint:errcheck

		if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
			return fmt.Errorf("ensuring vector table: %w", err)
		}

		if dirFlag != "" {
			return importDir(cmd, database, dirFlag)
		}

		return importFile(cmd, database, args[0], nameFlag)
	},
}

// importFile imports a single JSONL file into the database.
func importFile(cmd *cobra.Command, database *db.DB, filePath, overrideName string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("opening file %s: %w", filePath, err)
	}
	defer f.Close() //nolint:errcheck

	header, count, err := backup.ImportCollection(f, database, overrideName)
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
func importDir(cmd *cobra.Command, database *db.DB, dirPath string) error {
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

		header, count, err := backup.ImportCollection(f, database, "")
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
	rootCmd.AddCommand(importCmd)
}
