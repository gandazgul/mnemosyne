// Package search implements hybrid search combining full-text search (FTS5)
// and vector similarity search (sqlite-vec), fused via Reciprocal Rank Fusion.
package search

import (
	"sort"
	"time"
)

// Result is a unified search result that can originate from FTS, vector search,
// or a fusion of both. It carries the original scores from each source along
// with the fused RRF score.
type Result struct {
	// Document fields.
	DocumentID   int64
	CollectionID int64
	Content      string
	Metadata     *string
	CreatedAt    time.Time

	// RRFScore is the Reciprocal Rank Fusion score (higher = more relevant).
	// Only set when results are fused from multiple sources.
	RRFScore float64

	// FTSRank is the BM25 relevance score from FTS5 (higher = more relevant).
	// Zero if this document was not found by FTS search.
	FTSRank float64

	// VecDistance is the cosine distance from vector search (lower = more similar).
	// Zero if this document was not found by vector search.
	VecDistance float64

	// Sources lists which search methods found this document ("fts", "vector").
	Sources []string
}

// RankedEntry represents a single document in a ranked list, identified by its
// position (1-indexed rank) and document ID.
type RankedEntry struct {
	DocumentID int64
	Rank       int // 1-indexed position in the ranked list
}

// RankedList is an ordered list of ranked entries from a single search source.
type RankedList struct {
	Name    string // source name, e.g. "fts" or "vector"
	Entries []RankedEntry
}

// FuseRRF combines multiple ranked lists using Reciprocal Rank Fusion.
//
// The RRF score for a document is computed as:
//
//	RRF_score(d) = SUM over all rankings R of: 1 / (k + rank_R(d))
//
// where k is a constant (typically 60) that dampens the contribution of
// high-ranked documents, and rank_R(d) is the 1-indexed position of document d
// in ranking R. Documents not present in a ranking receive no contribution from
// that ranking (they are simply skipped, not penalized).
//
// The returned scores are a map from document ID to RRF score, sorted by score
// descending is left to the caller.
func FuseRRF(k int, lists ...RankedList) map[int64]float64 {
	scores := make(map[int64]float64)

	for _, list := range lists {
		for _, entry := range list.Entries {
			scores[entry.DocumentID] += 1.0 / float64(k+entry.Rank)
		}
	}

	return scores
}

// SortByRRFScore sorts a slice of Results by RRFScore descending (highest first).
func SortByRRFScore(results []Result) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].RRFScore > results[j].RRFScore
	})
}
