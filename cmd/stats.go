package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gandazgul/mnemoteca/internal/config"
	"github.com/spf13/cobra"
)

type statsOutput struct {
	DatabasePath    string `json:"database_path"`
	CollectionCount int    `json:"collection_count"`
	DocumentCount   int64  `json:"document_count"`
}

// statsCmd displays system statistics.
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show database size and total document counts",
	Long:  "Display statistics about the mnemoteca database and models.",
	RunE: func(cmd *cobra.Command, args []string) error {
		formatFlag, _ := cmd.Flags().GetString("format")
		if formatFlag != formatColor && formatFlag != formatPlain && formatFlag != formatJSON {
			return fmt.Errorf("invalid format %q; must be one of: color, plain, json", formatFlag)
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

		collections, err := database.ListCollections()
		if err != nil {
			return fmt.Errorf("listing collections: %w", err)
		}

		var totalDocs int64
		for _, c := range collections {
			count, err := database.CountDocuments(c.ID, nil)
			if err != nil {
				return fmt.Errorf("counting documents: %w", err)
			}
			totalDocs += count
		}

		resolvedDBPath, err := filepath.Abs(cfg.DBPath)
		if err != nil {
			return fmt.Errorf("resolving database path: %w", err)
		}

		if formatFlag == formatJSON {
			payload := statsOutput{
				DatabasePath:    resolvedDBPath,
				CollectionCount: len(collections),
				DocumentCount:   totalDocs,
			}
			enc := json.NewEncoder(cmd.OutOrStdout())
			return enc.Encode(payload)
		}

		info, err := os.Stat(cfg.DBPath)
		var sizeStr string
		if err != nil {
			sizeStr = "unknown"
		} else {
			sizeStr = fmt.Sprintf("%.2f MB", float64(info.Size())/1024/1024)
		}

		cmd.Println("Mnemoteca Statistics")
		cmd.Println("────────────────────────────────────────────────────────────")
		cmd.Printf("Database Path:   %s\n", resolvedDBPath)
		cmd.Printf("Database Size:   %s\n", sizeStr)
		cmd.Printf("Collections:     %d\n", len(collections))
		cmd.Printf("Total Documents: %d\n", totalDocs)
		cmd.Printf("Embedding Model: %s\n", cfg.Embedding.ModelPath)
		cmd.Printf("Dimensions:      %d\n", cfg.Embedding.Dimensions)
		cmd.Printf("Reranker Model:  %s (Enabled: %t)\n", cfg.Reranker.ModelPath, cfg.Reranker.Enabled)

		return nil
	},
}

func init() {
	statsCmd.Flags().StringP("format", "f", formatColor, "output format: color (default), plain, or json")
	rootCmd.AddCommand(statsCmd)
}
