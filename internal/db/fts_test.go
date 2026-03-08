package db_test

import (
	"testing"
)

func TestSearchFTS_BasicMatch(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("search-test")

	database.InsertDocument(c.ID, "Go is a statically typed programming language", nil)
	database.InsertDocument(c.ID, "Python is a dynamically typed language", nil)
	database.InsertDocument(c.ID, "Rust focuses on memory safety", nil)

	results, err := database.SearchFTS(c.ID, "programming language", 0)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	// "Go is a statically typed programming language" should match.
	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}

	// The result with "programming language" (exact match of both terms) should rank first.
	found := false
	for _, r := range results {
		if r.Content == "Go is a statically typed programming language" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find 'Go is a statically typed programming language' in results")
	}
}

func TestSearchFTS_RankedByRelevance(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("rank-test")

	// Insert documents with varying relevance to "golang concurrency".
	database.InsertDocument(c.ID, "Golang concurrency patterns using goroutines and channels in golang", nil)
	database.InsertDocument(c.ID, "Introduction to concurrency", nil)
	database.InsertDocument(c.ID, "Cooking recipes for beginners", nil)

	results, err := database.SearchFTS(c.ID, "golang concurrency", 0)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	if len(results) < 1 {
		t.Fatal("expected at least one result")
	}

	// The most relevant document should be first (highest rank score).
	if results[0].Content != "Golang concurrency patterns using goroutines and channels in golang" {
		t.Errorf("most relevant result = %q, want the golang concurrency document", results[0].Content)
	}

	// All results should have positive rank (since we negate bm25()).
	for i, r := range results {
		if r.Rank <= 0 {
			t.Errorf("result[%d].Rank = %f, want positive value", i, r.Rank)
		}
	}
}

func TestSearchFTS_CollectionScoped(t *testing.T) {
	database := testDB(t)

	c1, _ := database.CreateCollection("collection-a")
	c2, _ := database.CreateCollection("collection-b")

	database.InsertDocument(c1.ID, "Go programming in collection A", nil)
	database.InsertDocument(c2.ID, "Go programming in collection B", nil)

	// Search in collection A only.
	results, err := database.SearchFTS(c1.ID, "programming", 0)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result in collection A, got %d", len(results))
	}
	if results[0].CollectionID != c1.ID {
		t.Errorf("result collection_id = %d, want %d", results[0].CollectionID, c1.ID)
	}

	// Search in collection B only.
	results, err = database.SearchFTS(c2.ID, "programming", 0)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result in collection B, got %d", len(results))
	}
	if results[0].CollectionID != c2.ID {
		t.Errorf("result collection_id = %d, want %d", results[0].CollectionID, c2.ID)
	}
}

func TestSearchFTS_NoResults(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("empty-search")

	database.InsertDocument(c.ID, "Hello world", nil)

	results, err := database.SearchFTS(c.ID, "nonexistentterm", 0)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected 0 results, got %d", len(results))
	}
}

func TestSearchFTS_EmptyQuery(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("empty-query")

	database.InsertDocument(c.ID, "Some document", nil)

	results, err := database.SearchFTS(c.ID, "", 0)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected 0 results for empty query, got %d", len(results))
	}
}

func TestSearchFTS_WithLimit(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("limit-test")

	database.InsertDocument(c.ID, "Go tutorial part one", nil)
	database.InsertDocument(c.ID, "Go tutorial part two", nil)
	database.InsertDocument(c.ID, "Go tutorial part three", nil)

	results, err := database.SearchFTS(c.ID, "Go tutorial", 2)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	if len(results) != 2 {
		t.Errorf("expected 2 results with limit, got %d", len(results))
	}
}

func TestSearchFTS_SpecialCharacters(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("special-chars")

	database.InsertDocument(c.ID, "Hello world from Go", nil)

	// These queries contain FTS5 special characters that should be sanitized.
	queries := []string{
		"hello*",        // asterisk (FTS5 prefix operator)
		"hello + world", // plus sign
		"hello : world", // colon (column filter)
		"(hello)",       // parentheses (grouping)
		"hello ^ world", // caret
		"~hello",        // tilde
	}

	for _, q := range queries {
		results, err := database.SearchFTS(c.ID, q, 0)
		if err != nil {
			t.Errorf("SearchFTS(%q) error = %v", q, err)
			continue
		}
		// Should either return results or return empty -- but not error.
		_ = results
	}
}

func TestSearchFTS_PhraseQuery(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("phrase-test")

	database.InsertDocument(c.ID, "The quick brown fox jumps over the lazy dog", nil)
	database.InsertDocument(c.ID, "A brown and quick animal", nil)

	// Phrase query: "quick brown" should match exact sequence.
	results, err := database.SearchFTS(c.ID, `"quick brown"`, 0)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result for phrase query, got %d", len(results))
	}
	if results[0].Content != "The quick brown fox jumps over the lazy dog" {
		t.Errorf("wrong result: %q", results[0].Content)
	}
}

func TestSearchFTS_DocumentFields(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("fields-test")

	meta := `{"source": "test"}`
	database.InsertDocument(c.ID, "searchable content here", &meta)

	results, err := database.SearchFTS(c.ID, "searchable", 0)
	if err != nil {
		t.Fatalf("SearchFTS() error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	r := results[0]
	if r.ID == 0 {
		t.Error("expected non-zero ID")
	}
	if r.CollectionID != c.ID {
		t.Errorf("CollectionID = %d, want %d", r.CollectionID, c.ID)
	}
	if r.Content != "searchable content here" {
		t.Errorf("Content = %q, want %q", r.Content, "searchable content here")
	}
	if r.Metadata == nil || *r.Metadata != meta {
		t.Errorf("Metadata = %v, want %q", r.Metadata, meta)
	}
	if r.CreatedAt.IsZero() {
		t.Error("expected non-zero CreatedAt")
	}
	if r.Rank <= 0 {
		t.Errorf("Rank = %f, want positive value", r.Rank)
	}
}

func TestSearchFTS_AfterDelete(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("delete-sync")

	doc, _ := database.InsertDocument(c.ID, "temporary document about testing", nil)

	// Verify it's searchable.
	results, _ := database.SearchFTS(c.ID, "temporary", 0)
	if len(results) != 1 {
		t.Fatalf("expected 1 result before delete, got %d", len(results))
	}

	// Delete the document.
	database.DeleteDocument(doc.ID)

	// Verify it's no longer searchable (trigger should have removed FTS entry).
	results, err := database.SearchFTS(c.ID, "temporary", 0)
	if err != nil {
		t.Fatalf("SearchFTS() after delete error = %v", err)
	}
	if len(results) != 0 {
		t.Errorf("expected 0 results after delete, got %d", len(results))
	}
}

func TestSearchFTS_AfterCollectionDelete(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("cascade-fts")

	database.InsertDocument(c.ID, "document in doomed collection", nil)

	// Verify searchable.
	results, _ := database.SearchFTS(c.ID, "doomed", 0)
	if len(results) != 1 {
		t.Fatalf("expected 1 result before cascade delete, got %d", len(results))
	}

	// Delete the entire collection (CASCADE deletes documents, triggers update FTS).
	database.DeleteCollection("cascade-fts")

	// Verify FTS is cleaned up.
	results, err := database.SearchFTS(c.ID, "doomed", 0)
	if err != nil {
		t.Fatalf("SearchFTS() after cascade delete error = %v", err)
	}
	if len(results) != 0 {
		t.Errorf("expected 0 results after cascade delete, got %d", len(results))
	}
}

func TestSanitizeFTSQuery_UnbalancedQuotes(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("unbalanced-quotes")

	database.InsertDocument(c.ID, "hello world from Go", nil)

	// Unbalanced quote should not cause an error.
	results, err := database.SearchFTS(c.ID, `"hello`, 0)
	if err != nil {
		t.Fatalf("SearchFTS() with unbalanced quote error = %v", err)
	}

	// Should still find results (quotes stripped when unbalanced).
	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}
