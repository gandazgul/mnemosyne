package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd lists documents in a collection.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List documents in a collection",
	Long: `List all documents in the specified collection.

If --name is not provided, the current directory name is used.
Use --limit to restrict the number of results.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		globalFlag, _ := cmd.Flags().GetBool("global")
		limitFlag, _ := cmd.Flags().GetInt("limit")
		formatFlag, _ := cmd.Flags().GetString("format")

		if err := validateFormat(formatFlag); err != nil {
			return err
		}
		if plain(formatFlag) {
			color.NoColor = true
		}

		collectionName, err := resolveCollectionName(nameFlag, globalFlag)
		if err != nil {
			return err
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close() //nolint:errcheck

		collection, err := database.GetCollectionByName(collectionName)
		if err != nil {
			return fmt.Errorf("looking up collection: %w", err)
		}
		if collection == nil {
			return fmt.Errorf("collection %q does not exist; run 'mnemosyne init --name %s' first",
				collectionName, collectionName)
		}

		docs, err := database.ListDocuments(collection.ID, limitFlag)
		if err != nil {
			return fmt.Errorf("listing documents: %w", err)
		}

		if len(docs) == 0 {
			cmd.Printf("No documents in collection %q\n", collectionName)
			return nil
		}

		count, err := database.CountDocuments(collection.ID)
		if err != nil {
			return fmt.Errorf("counting documents: %w", err)
		}

		if plain(formatFlag) {
			cmd.Printf("Collection %q (%d documents)\n", collectionName, count)
		} else {
			cmd.Printf("Collection %s (%d documents)\n",
				boldCyan(collectionName), count)
			cmd.Println(dimWhite(strings.Repeat("─", 60)))
		}

		for _, doc := range docs {
			preview := doc.Content
			if len(preview) > 80 {
				preview = preview[:80] + "..."
			}

			ts := doc.CreatedAt.Format("2006-01-02 15:04:05")

			if plain(formatFlag) {
				cmd.Printf("  [%d]  %s\n", doc.ID, preview)
				cmd.Printf("        %s\n", ts)
			} else {
				cmd.Printf("  %s  %s\n",
					boldYellow(fmt.Sprintf("[%d]", doc.ID)),
					preview)
				cmd.Printf("        %s\n", dimWhite(ts))
			}
		}

		if limitFlag > 0 && int64(limitFlag) < count {
			cmd.Printf("\nShowing %d of %d documents. Use --limit to see more.\n", limitFlag, count)
		}

		return nil
	},
}

func init() {
	listCmd.Flags().StringP("name", "n", "", "collection name (defaults to current directory name)")
	listCmd.Flags().BoolP("global", "g", false, "use the global collection")
	listCmd.Flags().Int("limit", 20, "maximum number of documents to show")
	listCmd.Flags().String("format", "color", "output format: color (default) or plain")
	rootCmd.AddCommand(listCmd)
}
