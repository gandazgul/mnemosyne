package db_test

import (
	"testing"
)

func TestInsertDocument(t *testing.T) {
	database := testDB(t)

	c, _ := database.CreateCollection("docs")

	doc, err := database.InsertDocument(c.ID, "hello world", nil)
	if err != nil {
		t.Fatalf("InsertDocument() error = %v", err)
	}

	if doc.ID == 0 {
		t.Error("expected non-zero ID")
	}
	if doc.Content != "hello world" {
		t.Errorf("Content = %q, want %q", doc.Content, "hello world")
	}
	if doc.CollectionID != c.ID {
		t.Errorf("CollectionID = %d, want %d", doc.CollectionID, c.ID)
	}
}

func TestInsertDocument_WithMetadata(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("docs")

	meta := `{"source": "test.txt"}`
	doc, err := database.InsertDocument(c.ID, "content", &meta)
	if err != nil {
		t.Fatalf("InsertDocument() error = %v", err)
	}

	if doc.Metadata == nil {
		t.Fatal("expected non-nil metadata")
	}
	if *doc.Metadata != meta {
		t.Errorf("Metadata = %q, want %q", *doc.Metadata, meta)
	}
}

func TestGetDocumentByID(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("docs")

	inserted, _ := database.InsertDocument(c.ID, "find me", nil)

	doc, err := database.GetDocumentByID(inserted.ID)
	if err != nil {
		t.Fatalf("GetDocumentByID() error = %v", err)
	}
	if doc == nil {
		t.Fatal("GetDocumentByID() returned nil")
	}
	if doc.Content != "find me" {
		t.Errorf("Content = %q, want %q", doc.Content, "find me")
	}
}

func TestGetDocumentByID_NotFound(t *testing.T) {
	database := testDB(t)

	doc, err := database.GetDocumentByID(999)
	if err != nil {
		t.Fatalf("GetDocumentByID() error = %v", err)
	}
	if doc != nil {
		t.Errorf("expected nil, got %v", doc)
	}
}

func TestListDocuments(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("docs")

	_, _ = database.InsertDocument(c.ID, "first", nil)
	_, _ = database.InsertDocument(c.ID, "second", nil)
	_, _ = database.InsertDocument(c.ID, "third", nil)

	// List all (no limit).
	docs, err := database.ListDocuments(c.ID, 0)
	if err != nil {
		t.Fatalf("ListDocuments() error = %v", err)
	}
	if len(docs) != 3 {
		t.Fatalf("expected 3 documents, got %d", len(docs))
	}

	// Newest first.
	if docs[0].Content != "third" {
		t.Errorf("first doc = %q, want %q (newest first)", docs[0].Content, "third")
	}
}

func TestListDocuments_WithLimit(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("docs")

	_, _ = database.InsertDocument(c.ID, "a", nil)
	_, _ = database.InsertDocument(c.ID, "b", nil)
	_, _ = database.InsertDocument(c.ID, "c", nil)

	docs, err := database.ListDocuments(c.ID, 2)
	if err != nil {
		t.Fatalf("ListDocuments() error = %v", err)
	}
	if len(docs) != 2 {
		t.Errorf("expected 2 documents, got %d", len(docs))
	}
}

func TestListDocuments_CollectionScoped(t *testing.T) {
	database := testDB(t)

	c1, _ := database.CreateCollection("collection-1")
	c2, _ := database.CreateCollection("collection-2")

	_, _ = database.InsertDocument(c1.ID, "doc in c1", nil)
	_, _ = database.InsertDocument(c2.ID, "doc in c2", nil)
	_, _ = database.InsertDocument(c2.ID, "another in c2", nil)

	docs1, _ := database.ListDocuments(c1.ID, 0)
	docs2, _ := database.ListDocuments(c2.ID, 0)

	if len(docs1) != 1 {
		t.Errorf("collection-1 has %d docs, want 1", len(docs1))
	}
	if len(docs2) != 2 {
		t.Errorf("collection-2 has %d docs, want 2", len(docs2))
	}
}

func TestDeleteDocument(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("docs")

	doc, _ := database.InsertDocument(c.ID, "delete me", nil)

	err := database.DeleteDocument(doc.ID)
	if err != nil {
		t.Fatalf("DeleteDocument() error = %v", err)
	}

	// Verify it's gone.
	found, _ := database.GetDocumentByID(doc.ID)
	if found != nil {
		t.Error("document still exists after delete")
	}
}

func TestDeleteDocument_NotFound(t *testing.T) {
	database := testDB(t)

	err := database.DeleteDocument(999)
	if err == nil {
		t.Error("expected error when deleting non-existent document, got nil")
	}
}

func TestCountDocuments(t *testing.T) {
	database := testDB(t)
	c, _ := database.CreateCollection("docs")

	count, _ := database.CountDocuments(c.ID)
	if count != 0 {
		t.Errorf("empty collection count = %d, want 0", count)
	}

	_, _ = database.InsertDocument(c.ID, "one", nil)
	_, _ = database.InsertDocument(c.ID, "two", nil)

	count, _ = database.CountDocuments(c.ID)
	if count != 2 {
		t.Errorf("count = %d, want 2", count)
	}
}

func TestInsertDocumentWithVector(t *testing.T) {
	database := testDBWithVectors(t, 4)
	c, _ := database.CreateCollection("docs")

	vec := []float32{1.0, 0.5, 0.25, 0.125}
	doc, err := database.InsertDocumentWithVector(c.ID, "vector doc", nil, vec)
	if err != nil {
		t.Fatalf("InsertDocumentWithVector() error = %v", err)
	}

	if doc == nil || doc.ID == 0 {
		t.Fatal("expected non-zero ID")
	}

	results, err := database.SearchVectors(c.ID, vec, 1)
	if err != nil || len(results) == 0 || results[0].ID != doc.ID {
		t.Errorf("expected to find inserted vector")
	}
}
