package search

import (
	"fmt"
	"sort"

	"github.com/gandazgul/mnemosyne/internal/db"
	"github.com/gandazgul/mnemosyne/internal/embedding"
	"github.com/gandazgul/mnemosyne/internal/reranker"
)

const (
	FusionRRF        = "rrf"
	FusionVectorBM25 = "vector-bm25"

	defaultBM25Weight = 0.10
)

// Options configures a search operation.
type Options struct {
	// CollectionID is the collection to search within.
	CollectionID int64

	// Query is the user's search query text.
	Query string

	// Limit is the maximum number of final results to return.
	Limit int

	// RRFK is the RRF constant (typically 60). Higher values make the
	// ranking more "flat" (less differentiation between positions).
	RRFK int

	// Fusion selects the candidate fusion/ranking strategy. Empty means vector-bm25.
	Fusion string

	// BM25Weight controls how much vector-bm25 fusion should weight in-memory
	// BM25 over vector candidates. The vector weight is 1-BM25Weight.
	BM25Weight float64

	// ReRankCandidates is the number of candidates to retrieve from each
	// source before fusion. Each source fetches this many results, then
	// RRF fusion produces the final top-Limit results.
	// If zero, defaults to Limit.
	ReRankCandidates int

	// RerankerThreshold is the minimum reranker score (logit) for inclusion.
	// Applied to reranked results. Default from config: 0.0.
	RerankerThreshold float64

	// RRFThreshold is the minimum RRF fusion score for inclusion.
	// Applied when reranking is disabled. Default from config: 0.01.
	RRFThreshold float64

	// DisableThreshold skips all score-based filtering when true.
	DisableThreshold bool

	// NoRerank disables the cross-encoder reranking step, even if a
	// reranker is available in the engine.
	NoRerank bool

	// FTSOnly disables vector candidate retrieval.
	FTSOnly bool

	// VectorOnly disables full-text candidate retrieval.
	VectorOnly bool

	// Tags filters results to only those matching the specified tags.
	Tags []string
}

// Engine performs hybrid search combining FTS5 and vector similarity,
// fused via Reciprocal Rank Fusion.
type Engine struct {
	db       *db.DB
	embedder embedding.Embedder
	reranker reranker.Reranker
}

// NewEngine creates a search engine with the given database, embedder, and optional reranker.
func NewEngine(database *db.DB, embedder embedding.Embedder, reranker reranker.Reranker) *Engine {
	return &Engine{
		db:       database,
		embedder: embedder,
		reranker: reranker,
	}
}

// Search retrieves candidates, applies the configured fusion strategy, and then
// optionally reranks with a cross-encoder.
func (e *Engine) Search(opts Options) ([]Result, error) {
	if opts.FTSOnly && opts.VectorOnly {
		return nil, fmt.Errorf("cannot use both FTSOnly and VectorOnly")
	}

	fusion := opts.Fusion
	if fusion == "" {
		if opts.FTSOnly || opts.VectorOnly {
			fusion = FusionRRF
		} else {
			fusion = FusionVectorBM25
		}
	}
	switch fusion {
	case FusionRRF, FusionVectorBM25:
	default:
		return nil, fmt.Errorf("unknown fusion mode %q", fusion)
	}
	if fusion == FusionVectorBM25 && opts.FTSOnly {
		return nil, fmt.Errorf("fusion mode %q requires vector candidates", fusion)
	}

	useFTS := !opts.VectorOnly && fusion == FusionRRF
	useVector := !opts.FTSOnly

	if useVector && e.embedder == nil {
		return nil, fmt.Errorf("search requires an embedder")
	}

	// Determine how many candidates to retrieve from each source.
	// Fetch a richer candidate pool than the final limit so fusion has room to
	// rerank before trimming.
	candidates := opts.ReRankCandidates
	if candidates <= 0 {
		candidates = opts.Limit
	}

	var ftsResults []db.SearchResult
	var vecResults []db.VectorResult
	var rankedLists []RankedList

	if useFTS {
		var err error
		ftsResults, err = e.db.SearchFTS(opts.CollectionID, opts.Query, opts.Tags, candidates)
		if err != nil {
			return nil, fmt.Errorf("FTS search: %w", err)
		}

		ftsList := RankedList{Name: "fts", Entries: make([]RankedEntry, len(ftsResults))}
		for i, r := range ftsResults {
			ftsList.Entries[i] = RankedEntry{DocumentID: r.ID, Rank: i + 1}
		}
		rankedLists = append(rankedLists, ftsList)
	}

	if useVector {
		queryVec, err := e.embedder.EmbedQuery(opts.Query)
		if err != nil {
			return nil, fmt.Errorf("embedding query: %w", err)
		}

		vecResults, err = e.db.SearchVectors(opts.CollectionID, queryVec, opts.Tags, candidates)
		if err != nil {
			return nil, fmt.Errorf("vector search: %w", err)
		}

		vecList := RankedList{Name: "vector", Entries: make([]RankedEntry, len(vecResults))}
		for i, r := range vecResults {
			vecList.Entries[i] = RankedEntry{DocumentID: r.ID, Rank: i + 1}
		}
		rankedLists = append(rankedLists, vecList)
	}

	var results []Result
	if fusion == FusionVectorBM25 {
		results = vectorBM25Results(opts, vecResults)
	} else {
		results = rrfResults(opts, ftsResults, vecResults, rankedLists)
	}

	// Apply Reranker if configured and enabled
	if e.reranker != nil && !opts.NoRerank {
		// Take top ReRankCandidates
		rerankLimit := opts.ReRankCandidates
		if rerankLimit <= 0 {
			rerankLimit = opts.Limit
		}

		if len(results) > rerankLimit {
			results = results[:rerankLimit]
		}

		// Extract document contents for the reranker
		docs := make([]string, len(results))
		for i, r := range results {
			docs[i] = r.Content
		}

		// Score with cross-encoder
		rerankScores, err := e.reranker.Score(opts.Query, docs)
		if err != nil {
			return nil, fmt.Errorf("reranking failed: %w", err)
		}

		// Assign scores
		for i := range results {
			results[i].RerankerScore = rerankScores[i]
			results[i].IsReranked = true
		}

		// Re-sort by reranker score
		SortByRerankerScore(results)
	}

	// Filter by threshold (always applied unless explicitly disabled).
	if !opts.DisableThreshold {
		filtered := results[:0]
		for _, r := range results {
			if r.IsReranked {
				if float64(r.RerankerScore) >= opts.RerankerThreshold {
					filtered = append(filtered, r)
				}
			} else {
				if r.RRFScore >= opts.RRFThreshold {
					filtered = append(filtered, r)
				}
			}
		}
		results = filtered
	}

	// Trim to the requested limit.
	if opts.Limit > 0 && len(results) > opts.Limit {
		results = results[:opts.Limit]
	}

	return results, nil
}

func rrfResults(opts Options, ftsResults []db.SearchResult, vecResults []db.VectorResult, rankedLists []RankedList) []Result {
	// Fuse with RRF.
	rrfK := opts.RRFK
	if rrfK <= 0 {
		rrfK = 60 // sensible default
	}
	rrfScores := FuseRRF(rrfK, rankedLists...)

	// Build a lookup of document details from both result sets.
	ftsLookup := make(map[int64]db.SearchResult, len(ftsResults))
	for _, r := range ftsResults {
		ftsLookup[r.ID] = r
	}
	vecLookup := make(map[int64]db.VectorResult, len(vecResults))
	for _, r := range vecResults {
		vecLookup[r.ID] = r
	}

	// Build unified results.
	results := make([]Result, 0, len(rrfScores))
	for docID, score := range rrfScores {
		r := Result{
			DocumentID: docID,
			RRFScore:   score,
		}

		// Populate document fields and source scores from whichever source(s)
		// found this document.
		var sources []string
		if fts, ok := ftsLookup[docID]; ok {
			r.CollectionID = fts.CollectionID
			r.Content = fts.Content
			r.Metadata = fts.Metadata
			r.CreatedAt = fts.CreatedAt
			r.FTSRank = fts.Rank
			sources = append(sources, "fts")
		}
		if vec, ok := vecLookup[docID]; ok {
			r.CollectionID = vec.CollectionID
			r.Content = vec.Content
			r.Metadata = vec.Metadata
			r.CreatedAt = vec.CreatedAt
			r.VecDistance = vec.Distance
			sources = append(sources, "vector")
		}
		r.Sources = sources

		results = append(results, r)
	}

	// Sort by RRF score descending.
	SortByRRFScore(results)
	return results
}

func vectorBM25Results(opts Options, vecResults []db.VectorResult) []Result {
	docs := make([]string, len(vecResults))
	for i, vec := range vecResults {
		docs[i] = vec.Content
	}

	bm25Raw := bm25Scores(opts.Query, docs)
	bm25Norm := normalizeMax(bm25Raw)

	bm25Weight := opts.BM25Weight
	if bm25Weight == 0 {
		bm25Weight = defaultBM25Weight
	}
	if bm25Weight < 0 {
		bm25Weight = 0
	}
	if bm25Weight > 1 {
		bm25Weight = 1
	}
	vectorWeight := 1 - bm25Weight

	results := make([]Result, 0, len(vecResults))
	for i, vec := range vecResults {
		vectorSimilarity := 1 - vec.Distance
		if vectorSimilarity < 0 {
			vectorSimilarity = 0
		}

		sources := []string{"vector"}
		if bm25Raw[i] > 0 {
			sources = append(sources, "bm25")
		}

		results = append(results, Result{
			DocumentID:   vec.ID,
			CollectionID: vec.CollectionID,
			Content:      vec.Content,
			Metadata:     vec.Metadata,
			CreatedAt:    vec.CreatedAt,
			RRFScore:     vectorWeight*vectorSimilarity + bm25Weight*bm25Norm[i],
			BM25Score:    bm25Raw[i],
			VecDistance:  vec.Distance,
			Sources:      sources,
		})
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].RRFScore == results[j].RRFScore {
			return results[i].VecDistance < results[j].VecDistance
		}
		return results[i].RRFScore > results[j].RRFScore
	})

	return results
}
