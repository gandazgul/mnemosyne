package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/reranker"
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
		rerankCandidatesFlag, _ := cmd.Flags().GetInt("rerank-candidates")
		noRerankFlag, _ := cmd.Flags().GetBool("no-rerank")
		thresholdFlag, _ := cmd.Flags().GetFloat64("threshold")
		debugFlag, _ := cmd.Flags().GetBool("debug")
		formatFlag, _ := cmd.Flags().GetString("format")

		if err := validateFormat(formatFlag); err != nil {
			return err
		}
		if plain(formatFlag) {
			color.NoColor = true
		}

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

		rerankCandidates := cfg.Search.ReRankCandidates
		if rerankCandidatesFlag > 0 {
			rerankCandidates = rerankCandidatesFlag
		}

		if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
			return fmt.Errorf("ensuring vector table: %w", err)
		}

		embedder, err := openEmbedder(cfg)
		if err != nil {
			return fmt.Errorf("loading embedding model: %w", err)
		}
		defer embedder.Close()

		var rr reranker.Reranker
		if !noRerankFlag {
			rr, err = openReranker(cfg)
			if err != nil {
				return fmt.Errorf("loading reranker model: %w", err)
			}
			if rr != nil {
				defer rr.Close()
			}
		}

		engine := search.NewEngine(database, embedder, rr)
		results, err := engine.Search(search.Options{
			CollectionID:     collection.ID,
			Query:            query,
			Limit:            limitFlag,
			RRFK:             rrfK,
			ReRankCandidates: rerankCandidates,
			Threshold:        thresholdFlag,
			ApplyThreshold:   cmd.Flags().Changed("threshold"),
			NoRerank:         noRerankFlag,
		})
		if err != nil {
			return fmt.Errorf("searching: %w", err)
		}

		printSearchResults(results, query, collectionName, formatFlag, debugFlag)
		return nil
	},
}

// printSearchResults formats and prints search results to stdout.
func printSearchResults(results []search.Result, query, collectionName, formatFlag string, debugFlag bool) {
	if len(results) == 0 {
		fmt.Printf("No results for %q in collection %q\n", query, collectionName)
		return
	}

	// Header.
	if plain(formatFlag) {
		fmt.Printf("Search results for %q in collection %q (%d found)\n",
			query, collectionName, len(results))
	} else {
		fmt.Printf("%s %q in collection %s (%d found)\n",
			boldWhite("Search results for"),
			query,
			boldCyan(collectionName),
			len(results))
		fmt.Println(dimWhite(strings.Repeat("─", 60)))
	}

	for i, r := range results {
		if debugFlag {
			// Debug: show scores, sources, and component details.
			if plain(formatFlag) {
				fmt.Printf("  %d. [%d] score: %.6f sources: %s\n",
					i+1, r.DocumentID, r.RRFScore, strings.Join(r.Sources, "+"))
			} else {
				fmt.Printf("  %s %s score: %s sources: %s\n",
					boldWhite(fmt.Sprintf("%d.", i+1)),
					boldYellow(fmt.Sprintf("[%d]", r.DocumentID)),
					green(fmt.Sprintf("%.6f", r.RRFScore)),
					yellow(strings.Join(r.Sources, "+")))
			}

			// Show component scores for transparency.
			var details []string
			if r.IsReranked {
				details = append(details, fmt.Sprintf("rerank=%.4f", r.RerankerScore))
			}
			for _, src := range r.Sources {
				switch src {
				case "fts":
					details = append(details, fmt.Sprintf("fts_rank=%.4f", r.FTSRank))
				case "vector":
					details = append(details, fmt.Sprintf("vec_dist=%.4f", r.VecDistance))
				}
			}
			if len(details) > 0 {
				if plain(formatFlag) {
					fmt.Printf("     (%s)\n", strings.Join(details, ", "))
				} else {
					fmt.Printf("     %s\n", dimWhite("("+strings.Join(details, ", ")+")"))
				}
			}
		} else {
			// Normal: show rank and document ID.
			if plain(formatFlag) {
				fmt.Printf("  %d. [%d]", i+1, r.DocumentID)
			} else {
				fmt.Printf("  %s %s",
					boldWhite(fmt.Sprintf("%d.", i+1)),
					boldYellow(fmt.Sprintf("[%d]", r.DocumentID)))
			}
		}

		// Content — always shown.
		fmt.Printf("  %s\n", r.Content)

		// Timestamp — always shown.
		ts := r.CreatedAt.Format("2006-01-02 15:04:05")
		if plain(formatFlag) {
			fmt.Printf("     %s\n", ts)
		} else {
			fmt.Printf("     %s\n", dimWhite(ts))
		}

		if i < len(results)-1 {
			fmt.Println()
		}
	}
}

func init() {
	searchCmd.Flags().String("name", "", "collection name (defaults to current directory name)")
	searchCmd.Flags().Int("limit", 3, "maximum number of results to return")
	searchCmd.Flags().Int("rrf-k", 0, "RRF fusion constant (default from config, typically 60)")
	searchCmd.Flags().Int("rerank-candidates", 0, "number of candidates to pass to the reranker")
	searchCmd.Flags().Bool("no-rerank", false, "disable the cross-encoder reranking step")
	searchCmd.Flags().Float64("threshold", 0.0, "minimum score for a result to be included (e.g., 0.0 or 5.0 for reranker, 0.016 for RRF)")
	searchCmd.Flags().Bool("debug", false, "show scores, ranks, and sources for each result")
	searchCmd.Flags().String("format", "color", "output format: color (default) or plain")
	rootCmd.AddCommand(searchCmd)
}
