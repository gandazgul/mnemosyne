package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// searchCmd searches documents using full-text search (FTS5 + BM25).
var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search documents by keyword",
	Long: `Search documents in a collection using full-text search (FTS5 with BM25 ranking).

The query supports basic keyword matching. Use double quotes for exact phrases:
  mnemosyne search "exact phrase"
  mnemosyne search golang concurrency

If --name is not provided, the current directory name is used.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		limitFlag, _ := cmd.Flags().GetInt("limit")

		query := strings.TrimSpace(strings.Join(args, " "))
		if query == "" {
			return fmt.Errorf("search query cannot be empty")
		}

		// Resolve collection.
		collectionName, err := resolveCollectionName(nameFlag)
		if err != nil {
			return err
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close()

		collection, err := database.GetCollectionByName(collectionName)
		if err != nil {
			return fmt.Errorf("looking up collection: %w", err)
		}
		if collection == nil {
			return fmt.Errorf("collection %q does not exist; run 'mnemosyne init --name %s' first",
				collectionName, collectionName)
		}

		// Perform FTS5 search.
		results, err := database.SearchFTS(collection.ID, query, limitFlag)
		if err != nil {
			return fmt.Errorf("searching: %w", err)
		}

		if len(results) == 0 {
			fmt.Printf("No results for %q in collection %q\n", query, collectionName)
			return nil
		}

		fmt.Printf("Search results for %q in collection %q (%d found)\n",
			query, collectionName, len(results))
		fmt.Println(strings.Repeat("-", 60))

		for i, r := range results {
			fmt.Printf("  %d. [%d]  score: %.4f\n", i+1, r.ID, r.Rank)
			fmt.Printf("     %s\n", r.Content)
			fmt.Printf("     added %s\n", r.CreatedAt.Format("2006-01-02 15:04:05"))
			if i < len(results)-1 {
				fmt.Println()
			}
		}

		return nil
	},
}

func init() {
	searchCmd.Flags().String("name", "", "collection name (defaults to current directory name)")
	searchCmd.Flags().Int("limit", 10, "maximum number of results to return")
	rootCmd.AddCommand(searchCmd)
}
