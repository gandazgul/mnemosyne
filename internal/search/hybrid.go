package search

import (
	"fmt"

	"github.com/gandazgul/mnemosyne/internal/db"
	"github.com/gandazgul/mnemosyne/internal/embedding"
	"github.com/gandazgul/mnemosyne/internal/reranker"
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

	// ReRankCandidates is the number of candidates to retrieve from each
	// source before fusion. Each source fetches this many results, then
	// RRF fusion produces the final top-Limit results.
	// If zero, defaults to Limit.
	ReRankCandidates int

	// Threshold is the minimum score a result must have to be included.
	// If reranking is enabled, this applies to the RerankerScore (logits).
	// If reranking is disabled, this applies to the RRFScore.
	// Only applied if ApplyThreshold is true.
	Threshold float64

	// ApplyThreshold indicates whether Threshold should be used to filter results.
	ApplyThreshold bool

	// NoRerank disables the cross-encoder reranking step, even if a
	// reranker is available in the engine.
	NoRerank bool

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

// Search runs both FTS and vector searches, then fuses results with RRF.
// Documents found by both keyword match and semantic similarity are boosted
// above those found by only one method.
func (e *Engine) Search(opts Options) ([]Result, error) {
	if e.embedder == nil {
		return nil, fmt.Errorf("search requires an embedder")
	}

	// Determine how many candidates to retrieve from each source.
	// We fetch more than the final limit to give RRF a richer pool to fuse.
	candidates := opts.ReRankCandidates
	if candidates <= 0 {
		candidates = opts.Limit
	}

	// Run both searches.
	ftsResults, err := e.db.SearchFTS(opts.CollectionID, opts.Query, opts.Tags, candidates)
	if err != nil {
		return nil, fmt.Errorf("FTS search: %w", err)
	}

	queryVec, err := e.embedder.EmbedQuery(opts.Query)
	if err != nil {
		return nil, fmt.Errorf("embedding query: %w", err)
	}

	vecResults, err := e.db.SearchVectors(opts.CollectionID, queryVec, opts.Tags, candidates)
	if err != nil {
		return nil, fmt.Errorf("vector search: %w", err)
	}

	// Build ranked lists for RRF.
	ftsList := RankedList{Name: "fts", Entries: make([]RankedEntry, len(ftsResults))}
	for i, r := range ftsResults {
		ftsList.Entries[i] = RankedEntry{DocumentID: r.ID, Rank: i + 1}
	}

	vecList := RankedList{Name: "vector", Entries: make([]RankedEntry, len(vecResults))}
	for i, r := range vecResults {
		vecList.Entries[i] = RankedEntry{DocumentID: r.ID, Rank: i + 1}
	}

	// Fuse with RRF.
	rrfK := opts.RRFK
	if rrfK <= 0 {
		rrfK = 60 // sensible default
	}
	rrfScores := FuseRRF(rrfK, ftsList, vecList)

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

	// Filter by threshold if requested.
	if opts.ApplyThreshold {
		filtered := results[:0]
		for _, r := range results {
			if r.IsReranked {
				// Reranker score is float32, Threshold is float64
				if float64(r.RerankerScore) >= opts.Threshold {
					filtered = append(filtered, r)
				}
			} else {
				if r.RRFScore >= opts.Threshold {
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
