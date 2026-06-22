package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/embedding"
	"github.com/gandazgul/mnemosyne/internal/reranker"
	"github.com/gandazgul/mnemosyne/internal/search"
	"github.com/spf13/cobra"
)

// searchCmd searches documents using hybrid search combining full-text search
// and vector similarity search.
var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search documents using hybrid keyword + semantic search",
	Long: `Search documents in a collection using hybrid search, which combines
semantic vector search with a small lexical BM25 boost by default. Use
--fusion rrf to retrieve independent full-text and vector candidate lists and
combine them with Reciprocal Rank Fusion (RRF).

Documents found by both keyword match and semantic similarity are boosted
above those found by only one method.

Examples:
  mnemosyne search "exact phrase"
  mnemosyne search golang concurrency
  mnemosyne search --limit 5 "how do goroutines work"
  mnemosyne search -f json --limit 10 "benchmark query"
  mnemosyne search --fts-only --no-rerank "benchmark query"
  mnemosyne search --vector-only --no-rerank "benchmark query"
  mnemosyne search --fusion rrf --no-rerank "benchmark query"
  mnemosyne search --fusion vector-bm25 --bm25-weight 0.10 --rerank-candidates 300 --no-rerank "benchmark query"

If --name is not provided, the current directory name is used.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		globalFlag, _ := cmd.Flags().GetBool("global")
		limitFlag, _ := cmd.Flags().GetInt("limit")
		rrfKFlag, _ := cmd.Flags().GetInt("rrf-k")
		rerankCandidatesFlag, _ := cmd.Flags().GetInt("rerank-candidates")
		noRerankFlag, _ := cmd.Flags().GetBool("no-rerank")
		fusionFlag, _ := cmd.Flags().GetString("fusion")
		bm25WeightFlag, _ := cmd.Flags().GetFloat64("bm25-weight")
		thresholdFlag, _ := cmd.Flags().GetFloat64("threshold")
		noThresholdFlag, _ := cmd.Flags().GetBool("no-threshold")
		debugFlag, _ := cmd.Flags().GetBool("debug")
		formatFlag, _ := cmd.Flags().GetString("format")
		tagsFlag, _ := cmd.Flags().GetStringSlice("tag")
		ftsOnlyFlag, _ := cmd.Flags().GetBool("fts-only")
		vectorOnlyFlag, _ := cmd.Flags().GetBool("vector-only")

		if err := validateSearchFormat(formatFlag); err != nil {
			return err
		}
		if ftsOnlyFlag && vectorOnlyFlag {
			return fmt.Errorf("cannot use both --fts-only and --vector-only")
		}
		if plain(formatFlag) {
			color.NoColor = true
		}

		query := strings.TrimSpace(strings.Join(args, " "))
		if query == "" {
			return fmt.Errorf("search query cannot be empty")
		}

		// Resolve collection.
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

		// Load config for search settings and embedder.
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		// Determine RRF k: flag overrides config.
		rrfK := cfg.Search.RRFK
		if rrfKFlag > 0 {
			rrfK = rrfKFlag
		}

		rerankCandidates := cfg.Search.ReRankCandidates
		if rerankCandidatesFlag > 0 {
			rerankCandidates = rerankCandidatesFlag
		}

		fusion := cfg.Search.Fusion
		if fusion == "" {
			fusion = search.FusionVectorBM25
		}
		if cmd.Flags().Changed("fusion") {
			fusion = fusionFlag
		}
		if (ftsOnlyFlag || vectorOnlyFlag) && !cmd.Flags().Changed("fusion") {
			fusion = search.FusionRRF
		}
		if fusion != search.FusionRRF && fusion != search.FusionVectorBM25 {
			return fmt.Errorf("invalid fusion %q (expected %q or %q)", fusion, search.FusionRRF, search.FusionVectorBM25)
		}
		if fusion == search.FusionVectorBM25 && ftsOnlyFlag {
			return fmt.Errorf("--fusion %s cannot be used with --fts-only", search.FusionVectorBM25)
		}

		bm25Weight := cfg.Search.BM25Weight
		if bm25Weight == 0 {
			bm25Weight = 0.10
		}
		if cmd.Flags().Changed("bm25-weight") {
			bm25Weight = bm25WeightFlag
		}
		if bm25Weight < 0 || bm25Weight > 1 {
			return fmt.Errorf("bm25 weight must be between 0 and 1")
		}

		useVector := !ftsOnlyFlag
		useReranker := !noRerankFlag

		var embedder embedding.Embedder
		if useVector || useReranker {
			if useVector {
				if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
					return fmt.Errorf("ensuring vector table: %w", err)
				}
			}

			embedder, cfg, err = openEmbedder(cmd.Context())
			if err != nil {
				return fmt.Errorf("loading embedding model: %w", err)
			}
			defer embedder.Close() //nolint:errcheck
		}

		var rr reranker.Reranker
		if useReranker {
			rr, err = openReranker(cfg)
			if err != nil {
				return fmt.Errorf("loading reranker model: %w", err)
			}
			if rr != nil {
				defer rr.Close() //nolint:errcheck
			}
		}

		// Determine thresholds: use config defaults, override with --threshold flag.
		rerankerThreshold := cfg.Search.RerankerThreshold
		rrfThreshold := cfg.Search.RRFThreshold
		if cmd.Flags().Changed("threshold") {
			rerankerThreshold = thresholdFlag
			rrfThreshold = thresholdFlag
		}

		engine := search.NewEngine(database, embedder, rr)
		results, err := engine.Search(search.Options{
			CollectionID:      collection.ID,
			Query:             query,
			Limit:             limitFlag,
			RRFK:              rrfK,
			Fusion:            fusion,
			BM25Weight:        bm25Weight,
			ReRankCandidates:  rerankCandidates,
			RerankerThreshold: rerankerThreshold,
			RRFThreshold:      rrfThreshold,
			DisableThreshold:  noThresholdFlag,
			NoRerank:          noRerankFlag,
			FTSOnly:           ftsOnlyFlag,
			VectorOnly:        vectorOnlyFlag,
			Tags:              tagsFlag,
		})
		if err != nil {
			return fmt.Errorf("searching: %w", err)
		}

		return printSearchResults(results, query, collectionName, formatFlag, debugFlag)
	},
}

type searchJSONOutput struct {
	Query      string             `json:"query"`
	Collection string             `json:"collection"`
	Count      int                `json:"count"`
	Results    []searchJSONResult `json:"results"`
}

type searchJSONResult struct {
	Rank          int      `json:"rank"`
	DocumentID    int64    `json:"document_id"`
	CollectionID  int64    `json:"collection_id"`
	Content       string   `json:"content"`
	Metadata      any      `json:"metadata,omitempty"`
	MetadataRaw   string   `json:"metadata_raw,omitempty"`
	CreatedAt     string   `json:"created_at"`
	RRFScore      float64  `json:"rrf_score"`
	FTSRank       float64  `json:"fts_rank"`
	BM25Score     float64  `json:"bm25_score"`
	VecDistance   float64  `json:"vec_distance"`
	RerankerScore float32  `json:"reranker_score"`
	IsReranked    bool     `json:"is_reranked"`
	Sources       []string `json:"sources"`
}

// printSearchResults formats and prints search results to stdout.
func printSearchResults(results []search.Result, query, collectionName, formatFlag string, debugFlag bool) error {
	if formatFlag == formatJSON {
		return printSearchResultsJSON(results, query, collectionName)
	}

	if len(results) == 0 {
		fmt.Printf("No results for %q in collection %q\n", query, collectionName)
		return nil
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
		// Timestamp — always shown.
		ts := r.CreatedAt.Format("2006-01-02 15:04:05")
		if debugFlag {
			// Debug: show scores, sources, and component details.
			if plain(formatFlag) {
				fmt.Printf("%d. [%d] ts: %s score: %.6f sources: %s\n",
					i+1, r.DocumentID, ts, r.RRFScore, strings.Join(r.Sources, "+"))
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
				case "bm25":
					details = append(details, fmt.Sprintf("bm25=%.4f", r.BM25Score))
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
				fmt.Printf("%d. [%d] %s - ", i+1, r.DocumentID, ts)
			} else {
				fmt.Printf("%s %s - ",
					boldWhite(fmt.Sprintf("%d.", i+1)),
					boldYellow(fmt.Sprintf("[%d]", r.DocumentID)))
			}
		}

		// Content — always shown.
		fmt.Printf("%s\n", r.Content)

		if i < len(results)-1 {
			fmt.Println()
		}
	}

	return nil
}

func printSearchResultsJSON(results []search.Result, query, collectionName string) error {
	out := searchJSONOutput{
		Query:      query,
		Collection: collectionName,
		Count:      len(results),
		Results:    make([]searchJSONResult, 0, len(results)),
	}

	for i, r := range results {
		item := searchJSONResult{
			Rank:          i + 1,
			DocumentID:    r.DocumentID,
			CollectionID:  r.CollectionID,
			Content:       r.Content,
			CreatedAt:     r.CreatedAt.Format(time.RFC3339Nano),
			RRFScore:      r.RRFScore,
			FTSRank:       r.FTSRank,
			BM25Score:     r.BM25Score,
			VecDistance:   r.VecDistance,
			RerankerScore: r.RerankerScore,
			IsReranked:    r.IsReranked,
			Sources:       r.Sources,
		}

		if r.Metadata != nil && *r.Metadata != "" {
			var metadata any
			if err := json.Unmarshal([]byte(*r.Metadata), &metadata); err == nil {
				item.Metadata = metadata
			} else {
				item.MetadataRaw = *r.Metadata
			}
		}

		out.Results = append(out.Results, item)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func init() {
	searchCmd.Flags().StringP("name", "n", "", "collection name (defaults to current directory name)")
	searchCmd.Flags().BoolP("global", "g", false, "use the global collection")
	searchCmd.Flags().Int("limit", 3, "maximum number of results to return")
	searchCmd.Flags().Int("rrf-k", 0, "RRF fusion constant (default from config, typically 60)")
	searchCmd.Flags().Int("rerank-candidates", 0, "number of candidates to pass to the reranker")
	searchCmd.Flags().Bool("no-rerank", false, "disable the cross-encoder reranking step")
	searchCmd.Flags().String("fusion", "", "fusion strategy: vector-bm25 or rrf (default from config)")
	searchCmd.Flags().Float64("bm25-weight", -1, "BM25 lexical weight for --fusion vector-bm25, 0.0 to 1.0 (default from config)")
	searchCmd.Flags().Bool("fts-only", false, "use only full-text search candidates")
	searchCmd.Flags().Bool("vector-only", false, "use only vector search candidates")
	searchCmd.Flags().Float64("threshold", 0.0, "minimum score for a result to be included (overrides config rank/RRF limits if set)")
	searchCmd.Flags().Bool("no-threshold", false, "disable score-based filtering (return all results)")
	searchCmd.Flags().Bool("debug", false, "show scores, ranks, and sources for each result")
	searchCmd.Flags().StringP("format", "f", "color", "output format: color (default), plain, or json")
	searchCmd.Flags().StringSliceP("tag", "t", nil, "filter results by one or more tags (must match all)")
	rootCmd.AddCommand(searchCmd)
}
