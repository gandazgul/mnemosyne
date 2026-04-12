package search_test

import (
	"testing"

	"github.com/gandazgul/mnemosyne/internal/search"
)

// mockReranker returns deterministic scores for testing.
type mockReranker struct {
	scores map[string]float32
}

func newMockReranker(scores map[string]float32) *mockReranker {
	return &mockReranker{scores: scores}
}

func (m *mockReranker) Score(query string, documents []string) ([]float32, error) {
	results := make([]float32, len(documents))
	for i, doc := range documents {
		if score, ok := m.scores[doc]; ok {
			results[i] = score
		} else {
			results[i] = -10.0 // Default low score
		}
	}
	return results, nil
}

func (m *mockReranker) Close() error {
	return nil
}

func TestEngine_Search_Rerank(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Python programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	// Make the python doc score higher in the reranker
	reranker := newMockReranker(map[string]float32{
		"Go programming language":     1.5,
		"Python programming language": 5.0,
		"Cooking recipes":             -5.0,
	})

	engine := search.NewEngine(database, embedder, reranker)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "programming language",
		Limit:            10,
		RRFK:             60,
		ReRankCandidates: 10,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected results")
	}

	// Python should be first because of the reranker score
	if results[0].Content != "Python programming language" {
		t.Errorf("expected Python first due to reranker, got %q", results[0].Content)
	}

	for _, r := range results {
		if !r.IsReranked {
			t.Errorf("expected doc %d to be marked as reranked", r.DocumentID)
		}
	}

	// Results should be sorted by RerankerScore descending.
	for i := 1; i < len(results); i++ {
		if results[i].RerankerScore > results[i-1].RerankerScore {
			t.Errorf("results not sorted by RerankerScore: [%d]=%f > [%d]=%f",
				i, results[i].RerankerScore, i-1, results[i-1].RerankerScore)
		}
	}
}

func TestEngine_Search_Rerank_Threshold(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Python programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	// Make the python doc score higher in the reranker
	reranker := newMockReranker(map[string]float32{
		"Go programming language":     1.5,
		"Python programming language": 5.0,
		"Cooking recipes":             -5.0,
	})

	engine := search.NewEngine(database, embedder, reranker)
	results, err := engine.Search(search.Options{
		CollectionID:      collID,
		Query:             "programming language",
		Limit:             10,
		RRFK:              60,
		ReRankCandidates:  10,
		RerankerThreshold: 2.0, // Only Python should pass
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected exactly 1 result above threshold, got %d", len(results))
	}

	if results[0].Content != "Python programming language" {
		t.Errorf("expected Python to pass threshold, got %q", results[0].Content)
	}
}

func TestEngine_Search_NoRerank_Threshold(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Python programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	engine := search.NewEngine(database, embedder, nil) // No reranker
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "programming language",
		Limit:            10,
		RRFK:             60,
		ReRankCandidates: 10,
		RRFThreshold:     0.02, // A reasonable RRF threshold for this mock
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	// RRF scores will be around 1/61 + 1/62 ≈ 0.03 for the matches, and 0 for cooking.
	if len(results) != 2 {
		for i, r := range results {
			t.Logf("Result %d: %q (RRFScore: %f, VecDistance: %f, FTSRank: %f)", i, r.Content, r.RRFScore, r.VecDistance, r.FTSRank)
		}
		t.Fatalf("expected exactly 2 results above threshold, got %d", len(results))
	}

	for _, r := range results {
		if r.Content == "Cooking recipes" {
			t.Errorf("did not expect 'Cooking recipes' to pass threshold")
		}
		if r.IsReranked {
			t.Errorf("expected doc not to be reranked")
		}
	}
}

func TestEngine_Search_Rerank_DefaultThreshold(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Python programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	// Cooking scores negative = irrelevant, Go scores just above 0
	reranker := newMockReranker(map[string]float32{
		"Go programming language":     0.5,
		"Python programming language": 5.0,
		"Cooking recipes":             -5.0,
	})

	engine := search.NewEngine(database, embedder, reranker)
	// No explicit threshold set — zero-value RerankerThreshold (0.0) filters negatives
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "programming language",
		Limit:            10,
		RRFK:             60,
		ReRankCandidates: 10,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	// Cooking (-5.0) should be filtered by default threshold of 0.0
	for _, r := range results {
		if r.Content == "Cooking recipes" {
			t.Errorf("expected 'Cooking recipes' to be filtered by default threshold, got score %f", r.RerankerScore)
		}
	}

	if len(results) != 2 {
		t.Errorf("expected 2 results after default threshold, got %d", len(results))
	}
}

func TestEngine_Search_DisableThreshold(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Python programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	reranker := newMockReranker(map[string]float32{
		"Go programming language":     1.5,
		"Python programming language": 5.0,
		"Cooking recipes":             -5.0,
	})

	engine := search.NewEngine(database, embedder, reranker)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "programming language",
		Limit:            10,
		RRFK:             60,
		ReRankCandidates: 10,
		DisableThreshold: true, // All results should be returned
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	// With threshold disabled, all 3 docs should be returned including Cooking
	if len(results) != 3 {
		t.Errorf("expected 3 results with threshold disabled, got %d", len(results))
	}

	foundCooking := false
	for _, r := range results {
		if r.Content == "Cooking recipes" {
			foundCooking = true
		}
	}
	if !foundCooking {
		t.Errorf("expected 'Cooking recipes' to be included when threshold is disabled")
	}
}

func TestEngine_Search_NoRerank_Flag(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Python programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	reranker := newMockReranker(map[string]float32{
		"Go programming language":     1.5,
		"Python programming language": 5.0,
		"Cooking recipes":             -5.0,
	})

	engine := search.NewEngine(database, embedder, reranker)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "programming language",
		Limit:            10,
		RRFK:             60,
		ReRankCandidates: 10,
		NoRerank:         true, // This should bypass the reranker
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	for _, r := range results {
		if r.IsReranked {
			t.Errorf("expected doc not to be reranked when NoRerank is true")
		}
	}
}
