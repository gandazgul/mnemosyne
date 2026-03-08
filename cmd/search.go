package cmd

import (
	"fmt"
	"strings"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/search"
	"github.com/spf13/cobra"
)

// searchCmd searches documents using hybrid search combining full-text search
// (FTS5 with BM25 ranking) and vector similarity search (cosine distance),
// fused via Reciprocal Rank Fusion (RRF).
var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search documents using hybrid keyword + semantic search",
	Long: `Search documents in a collection using hybrid search, which combines
full-text keyword search (BM25) and semantic vector search (cosine similarity),
fused via Reciprocal Rank Fusion (RRF).

Documents found by both keyword match and semantic similarity are boosted
above those found by only one method.

Examples:
  mnemosyne search "exact phrase"
  mnemosyne search golang concurrency
  mnemosyne search --limit 5 "how do goroutines work"

If --name is not provided, the current directory name is used.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		limitFlag, _ := cmd.Flags().GetInt("limit")
		rrfKFlag, _ := cmd.Flags().GetInt("rrf-k")

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

		// Load config for search settings and embedder.
		cfg := config.Load()

		// Determine RRF k: flag overrides config.
		rrfK := cfg.Search.RRFK
		if rrfKFlag > 0 {
			rrfK = rrfKFlag
		}

		if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
			return fmt.Errorf("ensuring vector table: %w", err)
		}

		embedder, err := openEmbedder(cfg)
		if err != nil {
			return fmt.Errorf("loading embedding model: %w", err)
		}
		defer embedder.Close()

		engine := search.NewEngine(database, embedder)
		results, err := engine.Search(search.Options{
			CollectionID:     collection.ID,
			Query:            query,
			Limit:            limitFlag,
			RRFK:             rrfK,
			ReRankCandidates: cfg.Search.ReRankCandidates,
		})
		if err != nil {
			return fmt.Errorf("searching: %w", err)
		}

		if len(results) == 0 {
			fmt.Printf("No results for %q in collection %q\n", query, collectionName)
			return nil
		}

		// Display results.
		fmt.Printf("Search results for %q in collection %q (%d found)\n",
			query, collectionName, len(results))
		fmt.Println(strings.Repeat("-", 60))

		for i, r := range results {
			fmt.Printf("  %d. [%d]  score: %.6f  sources: %s\n",
				i+1, r.DocumentID, r.RRFScore, strings.Join(r.Sources, "+"))

			// Show component scores for transparency.
			var details []string
			for _, src := range r.Sources {
				switch src {
				case "fts":
					details = append(details, fmt.Sprintf("fts_rank=%.4f", r.FTSRank))
				case "vector":
					details = append(details, fmt.Sprintf("vec_dist=%.4f", r.VecDistance))
				}
			}
			if len(details) > 0 {
				fmt.Printf("     (%s)\n", strings.Join(details, ", "))
			}

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
	searchCmd.Flags().Int("rrf-k", 0, "RRF fusion constant (default from config, typically 60)")
	rootCmd.AddCommand(searchCmd)
}
