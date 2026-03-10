package cmd

import (
	"fmt"
	"os"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/spf13/cobra"
)

// statsCmd displays system statistics.
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show database size and total document counts",
	Long:  "Display statistics about the mnemosyne database and models.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Load()

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close()

		collections, err := database.ListCollections()
		if err != nil {
			return fmt.Errorf("listing collections: %w", err)
		}

		var totalDocs int64
		for _, c := range collections {
			count, err := database.CountDocuments(c.ID)
			if err != nil {
				return fmt.Errorf("counting documents: %w", err)
			}
			totalDocs += count
		}

		info, err := os.Stat(cfg.DBPath)
		var sizeStr string
		if err != nil {
			sizeStr = "unknown"
		} else {
			sizeStr = fmt.Sprintf("%.2f MB", float64(info.Size())/1024/1024)
		}

		cmd.Println("Mnemosyne Statistics")
		cmd.Println("────────────────────────────────────────────────────────────")
		cmd.Printf("Database Path:   %s\n", cfg.DBPath)
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
	rootCmd.AddCommand(statsCmd)
}
