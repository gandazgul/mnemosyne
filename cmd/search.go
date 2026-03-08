package cmd

import (
	"fmt"
	"strings"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/db"
	"github.com/spf13/cobra"
)

// searchCmd searches documents using full-text search (FTS5) or vector similarity search.
var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search documents by keyword or semantic similarity",
	Long: `Search documents in a collection using full-text search (FTS5 with BM25 ranking)
or vector similarity search (cosine distance via sqlite-vec).

Search modes:
  fts     Full-text keyword search using BM25 ranking
  vector  Semantic similarity search using embedding vectors

The query supports basic keyword matching in FTS mode. Use double quotes for exact phrases:
  mnemosyne search "exact phrase"
  mnemosyne search golang concurrency
  mnemosyne search --mode vector "how do goroutines work"

If --name is not provided, the current directory name is used.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		limitFlag, _ := cmd.Flags().GetInt("limit")
		modeFlag, _ := cmd.Flags().GetString("mode")

		query := strings.TrimSpace(strings.Join(args, " "))
		if query == "" {
			return fmt.Errorf("search query cannot be empty")
		}

		// Validate search mode.
		switch modeFlag {
		case "fts", "vector":
			// valid
		default:
			return fmt.Errorf("invalid search mode %q; use 'fts' or 'vector'", modeFlag)
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

		switch modeFlag {
		case "fts":
			return searchFTS(database, collection.ID, collectionName, query, limitFlag)
		case "vector":
			return searchVector(database, collection.ID, collectionName, query, limitFlag)
		}

		return nil
	},
}

// searchFTS performs a full-text search using FTS5 with BM25 ranking.
func searchFTS(database *db.DB, collectionID int64, collectionName, query string, limit int) error {
	results, err := database.SearchFTS(collectionID, query, limit)
	if err != nil {
		return fmt.Errorf("searching: %w", err)
	}

	if len(results) == 0 {
		fmt.Printf("No results for %q in collection %q (mode: fts)\n", query, collectionName)
		return nil
	}

	fmt.Printf("Search results for %q in collection %q (%d found, mode: fts)\n",
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
}

// searchVector performs a semantic similarity search using vector embeddings.
func searchVector(database *db.DB, collectionID int64, collectionName, query string, limit int) error {
	// Load config and initialize the embedder for query embedding.
	cfg := config.Load()

	if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
		return fmt.Errorf("ensuring vector table: %w", err)
	}

	embedder, err := openEmbedder(cfg)
	if err != nil {
		return fmt.Errorf("loading embedding model: %w", err)
	}
	defer embedder.Close()

	// Embed the query using the query prefix.
	queryVec, err := embedder.EmbedQuery(query)
	if err != nil {
		return fmt.Errorf("embedding query: %w", err)
	}

	results, err := database.SearchVectors(collectionID, queryVec, limit)
	if err != nil {
		return fmt.Errorf("vector search: %w", err)
	}

	if len(results) == 0 {
		fmt.Printf("No results for %q in collection %q (mode: vector)\n", query, collectionName)
		return nil
	}

	fmt.Printf("Search results for %q in collection %q (%d found, mode: vector)\n",
		query, collectionName, len(results))
	fmt.Println(strings.Repeat("-", 60))

	for i, r := range results {
		fmt.Printf("  %d. [%d]  distance: %.4f\n", i+1, r.ID, r.Distance)
		fmt.Printf("     %s\n", r.Content)
		fmt.Printf("     added %s\n", r.CreatedAt.Format("2006-01-02 15:04:05"))
		if i < len(results)-1 {
			fmt.Println()
		}
	}

	return nil
}

func init() {
	searchCmd.Flags().String("name", "", "collection name (defaults to current directory name)")
	searchCmd.Flags().Int("limit", 10, "maximum number of results to return")
	searchCmd.Flags().String("mode", "vector", "search mode: fts, vector")
	rootCmd.AddCommand(searchCmd)
}
