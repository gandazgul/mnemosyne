package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd lists documents in a collection.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List documents in a collection",
	Long: `List all documents in the specified collection.

If --name is not provided, the current directory name is used.
Use --limit to restrict the number of results.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		nameFlag, _ := cmd.Flags().GetString("name")
		limitFlag, _ := cmd.Flags().GetInt("limit")

		collectionName, err := resolveCollectionName(nameFlag)
		if err != nil {
			return err
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer func() {
			if cerr := database.Close(); cerr != nil && err == nil {
				err = cerr
			}
		}()

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
			fmt.Printf("No documents in collection %q\n", collectionName)
			return nil
		}

		count, err := database.CountDocuments(collection.ID)
		if err != nil {
			return fmt.Errorf("counting documents: %w", err)
		}

		fmt.Printf("Collection %q (%d documents)\n", collectionName, count)
		fmt.Println(strings.Repeat("-", 60))

		for _, doc := range docs {
			preview := doc.Content
			if len(preview) > 80 {
				preview = preview[:80] + "..."
			}
			fmt.Printf("  [%d]  %s\n", doc.ID, preview)
			fmt.Printf("        added %s\n", doc.CreatedAt.Format("2006-01-02 15:04:05"))
		}

		if limitFlag > 0 && int64(limitFlag) < count {
			fmt.Printf("\nShowing %d of %d documents. Use --limit to see more.\n", limitFlag, count)
		}

		return nil
	},
}

func init() {
	listCmd.Flags().String("name", "", "collection name (defaults to current directory name)")
	listCmd.Flags().Int("limit", 20, "maximum number of documents to show")
	rootCmd.AddCommand(listCmd)
}
