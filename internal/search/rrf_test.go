package search

import (
	"math"
	"testing"
)

// almostEqual compares two float64 values within a small epsilon.
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestFuseRRF_SingleList(t *testing.T) {
	list := RankedList{
		Name: "fts",
		Entries: []RankedEntry{
			{DocumentID: 1, Rank: 1},
			{DocumentID: 2, Rank: 2},
			{DocumentID: 3, Rank: 3},
		},
	}

	scores := FuseRRF(60, list)

	// With k=60:
	// doc 1: 1/(60+1) = 0.016393
	// doc 2: 1/(60+2) = 0.016129
	// doc 3: 1/(60+3) = 0.015873
	if !almostEqual(scores[1], 1.0/61.0, 1e-9) {
		t.Errorf("doc 1 score = %f, want %f", scores[1], 1.0/61.0)
	}
	if !almostEqual(scores[2], 1.0/62.0, 1e-9) {
		t.Errorf("doc 2 score = %f, want %f", scores[2], 1.0/62.0)
	}
	if !almostEqual(scores[3], 1.0/63.0, 1e-9) {
		t.Errorf("doc 3 score = %f, want %f", scores[3], 1.0/63.0)
	}
}

func TestFuseRRF_TwoLists_Overlap(t *testing.T) {
	// Two ranked lists with overlapping documents.
	ftsList := RankedList{
		Name: "fts",
		Entries: []RankedEntry{
			{DocumentID: 10, Rank: 1}, // top in FTS
			{DocumentID: 20, Rank: 2},
			{DocumentID: 30, Rank: 3},
		},
	}
	vecList := RankedList{
		Name: "vector",
		Entries: []RankedEntry{
			{DocumentID: 20, Rank: 1}, // top in vector
			{DocumentID: 10, Rank: 2},
			{DocumentID: 40, Rank: 3}, // only in vector
		},
	}

	scores := FuseRRF(60, ftsList, vecList)

	// doc 10: 1/(60+1) + 1/(60+2) = 1/61 + 1/62
	expectedDoc10 := 1.0/61.0 + 1.0/62.0
	if !almostEqual(scores[10], expectedDoc10, 1e-9) {
		t.Errorf("doc 10 score = %f, want %f", scores[10], expectedDoc10)
	}

	// doc 20: 1/(60+2) + 1/(60+1) = 1/62 + 1/61 (same as doc 10, just swapped)
	expectedDoc20 := 1.0/62.0 + 1.0/61.0
	if !almostEqual(scores[20], expectedDoc20, 1e-9) {
		t.Errorf("doc 20 score = %f, want %f", scores[20], expectedDoc20)
	}

	// doc 10 and doc 20 should have equal scores (rank 1+2 vs 2+1).
	if !almostEqual(scores[10], scores[20], 1e-9) {
		t.Errorf("doc 10 (%f) and doc 20 (%f) should be equal", scores[10], scores[20])
	}

	// doc 30: only in FTS at rank 3 -> 1/63
	if !almostEqual(scores[30], 1.0/63.0, 1e-9) {
		t.Errorf("doc 30 score = %f, want %f", scores[30], 1.0/63.0)
	}

	// doc 40: only in vector at rank 3 -> 1/63
	if !almostEqual(scores[40], 1.0/63.0, 1e-9) {
		t.Errorf("doc 40 score = %f, want %f", scores[40], 1.0/63.0)
	}
}

func TestFuseRRF_NoOverlap(t *testing.T) {
	// Two lists with no overlapping documents.
	list1 := RankedList{
		Name: "fts",
		Entries: []RankedEntry{
			{DocumentID: 1, Rank: 1},
			{DocumentID: 2, Rank: 2},
		},
	}
	list2 := RankedList{
		Name: "vector",
		Entries: []RankedEntry{
			{DocumentID: 3, Rank: 1},
			{DocumentID: 4, Rank: 2},
		},
	}

	scores := FuseRRF(60, list1, list2)

	// Each doc only appears in one list, so scores are just 1/(k+rank).
	if len(scores) != 4 {
		t.Fatalf("expected 4 scored documents, got %d", len(scores))
	}

	if !almostEqual(scores[1], 1.0/61.0, 1e-9) {
		t.Errorf("doc 1 score = %f, want %f", scores[1], 1.0/61.0)
	}
	if !almostEqual(scores[3], 1.0/61.0, 1e-9) {
		t.Errorf("doc 3 score = %f, want %f", scores[3], 1.0/61.0)
	}
}

func TestFuseRRF_EmptyLists(t *testing.T) {
	scores := FuseRRF(60)
	if len(scores) != 0 {
		t.Errorf("expected empty scores for no lists, got %d entries", len(scores))
	}

	scores = FuseRRF(60, RankedList{Name: "empty"})
	if len(scores) != 0 {
		t.Errorf("expected empty scores for empty list, got %d entries", len(scores))
	}
}

func TestFuseRRF_DifferentK(t *testing.T) {
	list := RankedList{
		Name: "fts",
		Entries: []RankedEntry{
			{DocumentID: 1, Rank: 1},
			{DocumentID: 2, Rank: 2},
		},
	}

	// With k=1, top-ranked results get much higher scores.
	scores := FuseRRF(1, list)
	// doc 1: 1/(1+1) = 0.5
	// doc 2: 1/(1+2) = 0.333
	if !almostEqual(scores[1], 0.5, 1e-9) {
		t.Errorf("doc 1 score with k=1: %f, want 0.5", scores[1])
	}
	if !almostEqual(scores[2], 1.0/3.0, 1e-9) {
		t.Errorf("doc 2 score with k=1: %f, want %f", scores[2], 1.0/3.0)
	}

	// Ratio with k=1: 0.5/0.333 = 1.5
	// Ratio with k=60: (1/61)/(1/62) = 62/61 ≈ 1.016
	// Higher k makes rankings more "flat" (less differentiation).
	ratioK1 := scores[1] / scores[2]

	scoresK60 := FuseRRF(60, list)
	ratioK60 := scoresK60[1] / scoresK60[2]

	if ratioK1 <= ratioK60 {
		t.Errorf("k=1 ratio (%f) should be larger than k=60 ratio (%f)", ratioK1, ratioK60)
	}
}

func TestFuseRRF_DocumentInBothLists_BoostedAboveSingleList(t *testing.T) {
	// A document appearing in both lists should always score higher than
	// a document appearing in only one list at the same rank.
	list1 := RankedList{
		Name: "fts",
		Entries: []RankedEntry{
			{DocumentID: 1, Rank: 1}, // appears in both
			{DocumentID: 2, Rank: 2}, // only in FTS
		},
	}
	list2 := RankedList{
		Name: "vector",
		Entries: []RankedEntry{
			{DocumentID: 1, Rank: 1}, // appears in both
			{DocumentID: 3, Rank: 2}, // only in vector
		},
	}

	scores := FuseRRF(60, list1, list2)

	// doc 1 appears rank 1 in both -> 2 * 1/(60+1)
	// doc 2 appears rank 2 in FTS only -> 1/(60+2)
	// doc 1 should be strictly higher than doc 2
	if scores[1] <= scores[2] {
		t.Errorf("doc 1 (%f) should be higher than doc 2 (%f)", scores[1], scores[2])
	}
	if scores[1] <= scores[3] {
		t.Errorf("doc 1 (%f) should be higher than doc 3 (%f)", scores[1], scores[3])
	}
}

func TestSortByRRFScore(t *testing.T) {
	results := []Result{
		{DocumentID: 1, RRFScore: 0.01},
		{DocumentID: 2, RRFScore: 0.05},
		{DocumentID: 3, RRFScore: 0.03},
	}

	SortByRRFScore(results)

	if results[0].DocumentID != 2 {
		t.Errorf("expected doc 2 first, got doc %d", results[0].DocumentID)
	}
	if results[1].DocumentID != 3 {
		t.Errorf("expected doc 3 second, got doc %d", results[1].DocumentID)
	}
	if results[2].DocumentID != 1 {
		t.Errorf("expected doc 1 third, got doc %d", results[2].DocumentID)
	}
}
