package db_test

import (
	"testing"
)

func TestCreateCollection(t *testing.T) {
	database := testDB(t)

	c, err := database.CreateCollection("test-collection")
	if err != nil {
		t.Fatalf("CreateCollection() error = %v", err)
	}

	if c.ID == 0 {
		t.Error("expected non-zero ID")
	}
	if c.Name != "test-collection" {
		t.Errorf("Name = %q, want %q", c.Name, "test-collection")
	}
}

func TestCreateCollection_DuplicateName(t *testing.T) {
	database := testDB(t)

	_, err := database.CreateCollection("dup")
	if err != nil {
		t.Fatalf("first CreateCollection() error = %v", err)
	}

	_, err = database.CreateCollection("dup")
	if err == nil {
		t.Error("expected error on duplicate collection name, got nil")
	}
}

func TestGetCollectionByName(t *testing.T) {
	database := testDB(t)

	// Create a collection.
	created, err := database.CreateCollection("find-me")
	if err != nil {
		t.Fatalf("CreateCollection() error = %v", err)
	}

	// Find it by name.
	found, err := database.GetCollectionByName("find-me")
	if err != nil {
		t.Fatalf("GetCollectionByName() error = %v", err)
	}
	if found == nil {
		t.Fatal("GetCollectionByName() returned nil")
	}
	if found.ID != created.ID {
		t.Errorf("ID = %d, want %d", found.ID, created.ID)
	}
}

func TestGetCollectionByName_NotFound(t *testing.T) {
	database := testDB(t)

	found, err := database.GetCollectionByName("does-not-exist")
	if err != nil {
		t.Fatalf("GetCollectionByName() error = %v", err)
	}
	if found != nil {
		t.Errorf("expected nil, got %v", found)
	}
}

func TestGetOrCreateCollection(t *testing.T) {
	database := testDB(t)

	// First call should create.
	c1, created1, err := database.GetOrCreateCollection("idempotent")
	if err != nil {
		t.Fatalf("GetOrCreateCollection() error = %v", err)
	}
	if !created1 {
		t.Error("expected created=true on first call")
	}

	// Second call should return existing.
	c2, created2, err := database.GetOrCreateCollection("idempotent")
	if err != nil {
		t.Fatalf("GetOrCreateCollection() second call error = %v", err)
	}
	if created2 {
		t.Error("expected created=false on second call")
	}
	if c1.ID != c2.ID {
		t.Errorf("IDs don't match: %d vs %d", c1.ID, c2.ID)
	}
}

func TestListCollections(t *testing.T) {
	database := testDB(t)

	// Empty list.
	collections, err := database.ListCollections()
	if err != nil {
		t.Fatalf("ListCollections() error = %v", err)
	}
	if len(collections) != 0 {
		t.Errorf("expected 0 collections, got %d", len(collections))
	}

	// Create two collections and add a doc to one.
	c1, _ := database.CreateCollection("alpha")
	database.CreateCollection("beta")
	database.InsertDocument(c1.ID, "test doc", nil)

	collections, err = database.ListCollections()
	if err != nil {
		t.Fatalf("ListCollections() error = %v", err)
	}
	if len(collections) != 2 {
		t.Fatalf("expected 2 collections, got %d", len(collections))
	}

	// Should be sorted by name.
	if collections[0].Name != "alpha" {
		t.Errorf("first collection name = %q, want %q", collections[0].Name, "alpha")
	}
	if collections[0].DocumentCount != 1 {
		t.Errorf("alpha doc count = %d, want 1", collections[0].DocumentCount)
	}
	if collections[1].Name != "beta" {
		t.Errorf("second collection name = %q, want %q", collections[1].Name, "beta")
	}
	if collections[1].DocumentCount != 0 {
		t.Errorf("beta doc count = %d, want 0", collections[1].DocumentCount)
	}
}

func TestDeleteCollection(t *testing.T) {
	database := testDB(t)

	c, _ := database.CreateCollection("to-delete")
	// Add a document to verify CASCADE.
	database.InsertDocument(c.ID, "orphan doc", nil)

	err := database.DeleteCollection("to-delete")
	if err != nil {
		t.Fatalf("DeleteCollection() error = %v", err)
	}

	// Verify collection is gone.
	found, _ := database.GetCollectionByName("to-delete")
	if found != nil {
		t.Error("collection still exists after delete")
	}

	// Verify documents were cascaded.
	docs, _ := database.ListDocuments(c.ID, 0)
	if len(docs) != 0 {
		t.Errorf("expected 0 documents after cascade, got %d", len(docs))
	}
}

func TestDeleteCollection_NotFound(t *testing.T) {
	database := testDB(t)

	err := database.DeleteCollection("nope")
	if err == nil {
		t.Error("expected error when deleting non-existent collection, got nil")
	}
}
