package search_test

import (
	"math"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gandazgul/mnemoteca/internal/db"
	"github.com/gandazgul/mnemoteca/internal/search"
)

// mockEmbedder is a test embedder that returns deterministic vectors.
// It generates vectors based on simple keyword matching so we can
// control which documents are "semantically similar" to a query.
type mockEmbedder struct {
	dim int
	// keywords maps words to a dimension index. When a text contains the word,
	// that dimension is set to 1.0 (and the vector is normalized).
	keywords map[string]int
}

func newMockEmbedder(keywords []string) *mockEmbedder {
	m := &mockEmbedder{
		dim:      len(keywords),
		keywords: make(map[string]int, len(keywords)),
	}
	for i, kw := range keywords {
		m.keywords[kw] = i
	}
	return m
}

func (m *mockEmbedder) Embed(text string) ([]float32, error) {
	return m.embed(text), nil
}

func (m *mockEmbedder) EmbedBatch(texts []string) ([][]float32, error) {
	result := make([][]float32, len(texts))
	for i, t := range texts {
		result[i] = m.embed(t)
	}
	return result, nil
}

func (m *mockEmbedder) EmbedQuery(query string) ([]float32, error) {
	return m.embed(query), nil
}

func (m *mockEmbedder) EmbedDocument(doc string) ([]float32, error) {
	return m.embed(doc), nil
}

func (m *mockEmbedder) Dimensions() int {
	return m.dim
}

func (m *mockEmbedder) Close() error {
	return nil
}

// embed creates a vector where each dimension corresponds to a keyword.
// If the text contains the keyword, that dimension gets 1.0.
// All dimensions start with a small baseline (0.01) to ensure non-zero vectors
// even when no keywords match (zero vectors cause NULL distances in sqlite-vec).
// The vector is L2-normalized.
func (m *mockEmbedder) embed(text string) []float32 {
	vec := make([]float32, m.dim)
	lower := strings.ToLower(text)
	for i := range vec {
		vec[i] = 0.01 // small baseline
	}
	for kw, idx := range m.keywords {
		if strings.Contains(lower, kw) {
			vec[idx] = 1.0
		}
	}

	// L2 normalize.
	var norm float64
	for _, v := range vec {
		norm += float64(v * v)
	}
	if norm > 0 {
		norm = math.Sqrt(norm)
		for i := range vec {
			vec[i] = float32(float64(vec[i]) / norm)
		}
	}
	return vec
}

// testDB creates a temporary database for testing.
func testDB(t *testing.T) *db.DB {
	t.Helper()
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")
	database, err := db.Open(dbPath)
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}
	t.Cleanup(func() {
		database.Close() //nolint:errcheck
	})
	return database
}

// seedCollection creates a collection and inserts documents with embeddings.
func seedCollection(t *testing.T, database *db.DB, embedder *mockEmbedder, name string, docs []string) int64 {
	t.Helper()

	if err := database.EnsureVectorTable(embedder.Dimensions()); err != nil {
		t.Fatalf("EnsureVectorTable: %v", err)
	}

	c, err := database.CreateCollection(name)
	if err != nil {
		t.Fatalf("CreateCollection: %v", err)
	}

	for _, content := range docs {
		doc, err := database.InsertDocument(c.ID, content, nil)
		if err != nil {
			t.Fatalf("InsertDocument: %v", err)
		}

		vec, _ := embedder.EmbedDocument(content)
		if err := database.InsertVector(doc.ID, c.ID, vec); err != nil {
			t.Fatalf("InsertVector: %v", err)
		}
	}

	return c.ID
}

func TestEngine_Search(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "cooking", "programming", "goroutine", "channel"})
	database := testDB(t)

	docs := []string{
		"Go programming with goroutines and channels",      // strong FTS + vector match
		"Python programming language overview",             // partial match
		"Cooking recipes for beginners",                    // irrelevant
		"Go channels are used for goroutine communication", // strong on different keywords
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	engine := search.NewEngine(database, embedder, nil)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "go programming goroutine",
		Limit:            10,
		RRFK:             60,
		Fusion:           search.FusionRRF,
		ReRankCandidates: 10,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected results")
	}

	// All results should have RRF scores > 0.
	for _, r := range results {
		if r.RRFScore <= 0 {
			t.Errorf("expected positive RRFScore, got %f for doc %d", r.RRFScore, r.DocumentID)
		}
		if len(r.Sources) == 0 {
			t.Errorf("expected at least one source for doc %d", r.DocumentID)
		}
	}

	// Results should be sorted by RRFScore descending.
	for i := 1; i < len(results); i++ {
		if results[i].RRFScore > results[i-1].RRFScore {
			t.Errorf("results not sorted by RRFScore: [%d]=%f > [%d]=%f",
				i, results[i].RRFScore, i-1, results[i-1].RRFScore)
		}
	}

	// Documents found by both sources should have both scores populated.
	for _, r := range results {
		if len(r.Sources) == 2 {
			if r.FTSRank <= 0 {
				t.Errorf("doc %d found by both but FTSRank is %f", r.DocumentID, r.FTSRank)
			}
		}
	}
}

func TestEngine_Search_BothSourcesBoosted(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "cooking", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",    // matches both FTS and vector for "go programming"
		"Cooking recipes for dinner", // irrelevant to query
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	engine := search.NewEngine(database, embedder, nil)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "go programming",
		Limit:            10,
		RRFK:             60,
		Fusion:           search.FusionRRF,
		ReRankCandidates: 10,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected results")
	}

	// The "Go programming" doc should be first since it matches both sources.
	if results[0].Content != "Go programming language" {
		t.Errorf("expected 'Go programming language' first, got %q", results[0].Content)
	}
	if len(results[0].Sources) != 2 {
		t.Errorf("expected 2 sources for top result, got %v", results[0].Sources)
	}
}

func TestEngine_Search_Limit(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "python", "rust", "java", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Python programming language",
		"Rust programming language",
		"Java programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	engine := search.NewEngine(database, embedder, nil)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "programming language",
		Limit:            2,
		RRFK:             60,
		Fusion:           search.FusionRRF,
		ReRankCandidates: 10,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) > 2 {
		t.Errorf("expected at most 2 results, got %d", len(results))
	}
}

func TestEngine_Search_FTSOnly(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	engine := search.NewEngine(database, nil, nil)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "go programming",
		Limit:            10,
		RRFK:             60,
		ReRankCandidates: 10,
		FTSOnly:          true,
		DisableThreshold: true,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected FTS-only results")
	}
	for _, r := range results {
		if len(r.Sources) != 1 || r.Sources[0] != "fts" {
			t.Fatalf("expected only fts source, got %v", r.Sources)
		}
		if r.FTSRank <= 0 {
			t.Fatalf("expected FTSRank to be populated, got %f", r.FTSRank)
		}
		if r.VecDistance != 0 {
			t.Fatalf("expected VecDistance to be unset, got %f", r.VecDistance)
		}
	}
}

func TestEngine_Search_VectorOnly(t *testing.T) {
	embedder := newMockEmbedder([]string{"go", "programming"})
	database := testDB(t)

	docs := []string{
		"Go programming language",
		"Cooking recipes",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	engine := search.NewEngine(database, embedder, nil)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "go programming",
		Limit:            10,
		RRFK:             60,
		ReRankCandidates: 10,
		VectorOnly:       true,
		DisableThreshold: true,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected vector-only results")
	}
	for _, r := range results {
		if len(r.Sources) != 1 || r.Sources[0] != "vector" {
			t.Fatalf("expected only vector source, got %v", r.Sources)
		}
		if r.FTSRank != 0 {
			t.Fatalf("expected FTSRank to be unset, got %f", r.FTSRank)
		}
	}
}

func TestEngine_Search_VectorBM25Fusion(t *testing.T) {
	embedder := newMockEmbedder([]string{"semantic", "concept"})
	database := testDB(t)

	docs := []string{
		"semantic concept general discussion",
		"exactterm exactterm exactterm exactterm",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	engine := search.NewEngine(database, embedder, nil)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "semantic exactterm",
		Limit:            10,
		ReRankCandidates: 10,
		Fusion:           search.FusionVectorBM25,
		BM25Weight:       1.0,
		DisableThreshold: true,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) < 2 {
		t.Fatalf("expected at least 2 results, got %d", len(results))
	}
	if results[0].Content != "exactterm exactterm exactterm exactterm" {
		t.Fatalf("expected lexical match to be promoted, got %q", results[0].Content)
	}
	if results[0].BM25Score <= 0 {
		t.Fatalf("expected BM25Score to be populated, got %f", results[0].BM25Score)
	}
	if !contains(results[0].Sources, "bm25") || !contains(results[0].Sources, "vector") {
		t.Fatalf("expected vector and bm25 sources, got %v", results[0].Sources)
	}
}

func TestEngine_Search_DefaultFusionIsVectorBM25(t *testing.T) {
	embedder := newMockEmbedder([]string{"semantic", "concept"})
	database := testDB(t)

	docs := []string{
		"semantic concept general discussion",
		"exactterm exactterm exactterm exactterm",
	}

	collID := seedCollection(t, database, embedder, "test", docs)

	engine := search.NewEngine(database, embedder, nil)
	results, err := engine.Search(search.Options{
		CollectionID:     collID,
		Query:            "semantic exactterm",
		Limit:            10,
		ReRankCandidates: 10,
		BM25Weight:       1.0,
		DisableThreshold: true,
	})
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected default vector-BM25 results")
	}
	if !contains(results[0].Sources, "bm25") || !contains(results[0].Sources, "vector") {
		t.Fatalf("expected default fusion to include vector and bm25 sources, got %v", results[0].Sources)
	}
}

func TestEngine_Search_SourceModeConflict(t *testing.T) {
	database := testDB(t)
	engine := search.NewEngine(database, nil, nil)

	_, err := engine.Search(search.Options{
		CollectionID: 1,
		Query:        "test",
		Limit:        10,
		FTSOnly:      true,
		VectorOnly:   true,
	})
	if err == nil {
		t.Fatal("expected error when both source-only modes are enabled")
	}
}

func TestEngine_Search_VectorBM25RejectsFTSOnly(t *testing.T) {
	database := testDB(t)
	engine := search.NewEngine(database, nil, nil)

	_, err := engine.Search(search.Options{
		CollectionID: 1,
		Query:        "test",
		Limit:        10,
		Fusion:       search.FusionVectorBM25,
		FTSOnly:      true,
	})
	if err == nil {
		t.Fatal("expected error when vector-bm25 fusion is used with FTSOnly")
	}
}

func TestEngine_Search_RequiresEmbedder(t *testing.T) {
	database := testDB(t)
	engine := search.NewEngine(database, nil, nil)

	_, err := engine.Search(search.Options{
		CollectionID: 1,
		Query:        "test",
		Limit:        10,
	})
	if err == nil {
		t.Fatal("expected error when searching without embedder")
	}
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
