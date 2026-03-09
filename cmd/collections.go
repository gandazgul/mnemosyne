package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// collectionsCmd lists all collections and their document counts.
var collectionsCmd = &cobra.Command{
	Use:   "collections",
	Short: "List all collections",
	Long:  "Display a list of all collections and their document counts.",
	RunE: func(cmd *cobra.Command, args []string) error {
		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close()

		collections, err := database.ListCollections()
		if err != nil {
			return fmt.Errorf("listing collections: %w", err)
		}

		if len(collections) == 0 {
			cmd.Println("No collections found.")
			return nil
		}

		cmd.Printf("Collections (%d total)\n", len(collections))
		cmd.Println("────────────────────────────────────────────────────────────")

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 4, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tDOCUMENTS\tCREATED")

		for _, c := range collections {
			count, err := database.CountDocuments(c.ID)
			if err != nil {
				return fmt.Errorf("counting documents for %s: %w", c.Name, err)
			}
			fmt.Fprintf(w, "%d\t%s\t%d\t%s\n", c.ID, c.Name, count, c.CreatedAt.Format("2006-01-02 15:04:05"))
		}
		w.Flush()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(collectionsCmd)
}
