package db_test

import (
	"encoding/binary"
	"math"
	"testing"

	"github.com/gandazgul/mnemosyne/internal/db"
)

// testDBWithVectors creates a test database and ensures the vector table exists
// with the given dimension. This helper is used by all vector tests.
func testDBWithVectors(t *testing.T, dimension int) *db.DB {
	t.Helper()
	database := testDB(t)

	if err := database.EnsureVectorTable(dimension); err != nil {
		t.Fatalf("EnsureVectorTable(%d) error = %v", dimension, err)
	}

	return database
}

func TestSerializeFloat32(t *testing.T) {
	vec := []float32{1.0, 2.0, 3.0, -1.5}
	buf := db.SerializeFloat32(vec)

	// Each float32 is 4 bytes, so total should be 16 bytes.
	if len(buf) != 16 {
		t.Fatalf("SerializeFloat32 len = %d, want 16", len(buf))
	}

	// Verify each float32 is correctly encoded as little-endian IEEE 754.
	for i, expected := range vec {
		bits := binary.LittleEndian.Uint32(buf[i*4:])
		got := math.Float32frombits(bits)
		if got != expected {
			t.Errorf("SerializeFloat32[%d] = %f, want %f", i, got, expected)
		}
	}
}

func TestSerializeFloat32_Empty(t *testing.T) {
	buf := db.SerializeFloat32([]float32{})
	if len(buf) != 0 {
		t.Errorf("SerializeFloat32(empty) len = %d, want 0", len(buf))
	}
}

func TestEnsureVectorTable_Creates(t *testing.T) {
	database := testDB(t)

	// Table should not exist yet.
	var name string
	err := database.Conn().QueryRow(
		"SELECT name FROM sqlite_master WHERE type='table' AND name='docs_vec'",
	).Scan(&name)
	if err == nil {
		t.Fatal("docs_vec table should not exist before EnsureVectorTable")
	}

	// Create the vector table.
	if err := database.EnsureVectorTable(4); err != nil {
		t.Fatalf("EnsureVectorTable() error = %v", err)
	}

	// Verify the table exists.
	err = database.Conn().QueryRow(
		"SELECT name FROM sqlite_master WHERE type='table' AND name='docs_vec'",
	).Scan(&name)
	if err != nil {
		t.Fatalf("docs_vec table not found after EnsureVectorTable: %v", err)
	}
}

func TestEnsureVectorTable_Idempotent(t *testing.T) {
	database := testDB(t)

	// Calling EnsureVectorTable multiple times should not error.
	if err := database.EnsureVectorTable(4); err != nil {
		t.Fatalf("first EnsureVectorTable() error = %v", err)
	}
	if err := database.EnsureVectorTable(4); err != nil {
		t.Fatalf("second EnsureVectorTable() error = %v", err)
	}
}

func TestInsertVector(t *testing.T) {
	database := testDBWithVectors(t, 4)

	c, _ := database.CreateCollection("vec-test")
	doc, _ := database.InsertDocument(c.ID, "test document", nil)

	vec := []float32{0.1, 0.2, 0.3, 0.4}
	if err := database.InsertVector(doc.ID, c.ID, vec); err != nil {
		t.Fatalf("InsertVector() error = %v", err)
	}

	// Verify the vector was inserted by counting rows.
	var count int
	err := database.Conn().QueryRow("SELECT count(*) FROM docs_vec").Scan(&count)
	if err != nil {
		t.Fatalf("counting docs_vec rows: %v", err)
	}
	if count != 1 {
		t.Errorf("docs_vec row count = %d, want 1", count)
	}
}

func TestSearchVectors_BasicKNN(t *testing.T) {
	database := testDBWithVectors(t, 4)

	c, _ := database.CreateCollection("knn-test")

	// Insert three documents with known vectors.
	// doc1 is closest to the query vector, doc3 is farthest.
	doc1, _ := database.InsertDocument(c.ID, "very close match", nil)
	doc2, _ := database.InsertDocument(c.ID, "somewhat close match", nil)
	doc3, _ := database.InsertDocument(c.ID, "far away match", nil)

	_ = database.InsertVector(doc1.ID, c.ID, []float32{0.9, 0.1, 0.0, 0.0})
	_ = database.InsertVector(doc2.ID, c.ID, []float32{0.5, 0.5, 0.0, 0.0})
	_ = database.InsertVector(doc3.ID, c.ID, []float32{0.0, 0.0, 0.9, 0.1})

	// Query with a vector close to doc1.
	query := []float32{1.0, 0.0, 0.0, 0.0}
	results, err := database.SearchVectors(c.ID, query, nil, 3)
	if err != nil {
		t.Fatalf("SearchVectors() error = %v", err)
	}

	if len(results) != 3 {
		t.Fatalf("got %d results, want 3", len(results))
	}

	// doc1 should be the closest (lowest distance).
	if results[0].ID != doc1.ID {
		t.Errorf("first result ID = %d, want %d (doc1)", results[0].ID, doc1.ID)
	}

	// Distances should be in ascending order.
	for i := 1; i < len(results); i++ {
		if results[i].Distance < results[i-1].Distance {
			t.Errorf("results not sorted by distance: [%d].distance=%f < [%d].distance=%f",
				i, results[i].Distance, i-1, results[i-1].Distance)
		}
	}
}

func TestSearchVectors_ReturnsDocumentFields(t *testing.T) {
	database := testDBWithVectors(t, 4)

	c, _ := database.CreateCollection("fields-test")
	doc, _ := database.InsertDocument(c.ID, "the document content", nil)
	_ = database.InsertVector(doc.ID, c.ID, []float32{1.0, 0.0, 0.0, 0.0})

	results, err := database.SearchVectors(c.ID, []float32{1.0, 0.0, 0.0, 0.0}, nil, 1)
	if err != nil {
		t.Fatalf("SearchVectors() error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("got %d results, want 1", len(results))
	}

	r := results[0]
	if r.ID != doc.ID {
		t.Errorf("ID = %d, want %d", r.ID, doc.ID)
	}
	if r.CollectionID != c.ID {
		t.Errorf("CollectionID = %d, want %d", r.CollectionID, c.ID)
	}
	if r.Content != "the document content" {
		t.Errorf("Content = %q, want %q", r.Content, "the document content")
	}
	if r.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
}

func TestSearchVectors_CollectionScoped(t *testing.T) {
	database := testDBWithVectors(t, 4)

	// Create two collections with similar vectors.
	c1, _ := database.CreateCollection("collection-a")
	c2, _ := database.CreateCollection("collection-b")

	doc1, _ := database.InsertDocument(c1.ID, "doc in collection a", nil)
	doc2, _ := database.InsertDocument(c2.ID, "doc in collection b", nil)

	// Both documents have the same embedding.
	vec := []float32{1.0, 0.0, 0.0, 0.0}
	_ = database.InsertVector(doc1.ID, c1.ID, vec)
	_ = database.InsertVector(doc2.ID, c2.ID, vec)

	// Search in collection A should only return doc1.
	results, err := database.SearchVectors(c1.ID, vec, nil, 10)
	if err != nil {
		t.Fatalf("SearchVectors(c1) error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("got %d results for collection A, want 1", len(results))
	}
	if results[0].ID != doc1.ID {
		t.Errorf("result ID = %d, want %d (doc1 from collection A)", results[0].ID, doc1.ID)
	}

	// Search in collection B should only return doc2.
	results, err = database.SearchVectors(c2.ID, vec, nil, 10)
	if err != nil {
		t.Fatalf("SearchVectors(c2) error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("got %d results for collection B, want 1", len(results))
	}
	if results[0].ID != doc2.ID {
		t.Errorf("result ID = %d, want %d (doc2 from collection B)", results[0].ID, doc2.ID)
	}
}

func TestSearchVectors_Limit(t *testing.T) {
	database := testDBWithVectors(t, 4)
	c, _ := database.CreateCollection("limit-test")

	// Insert 5 documents.
	for i := 0; i < 5; i++ {
		doc, _ := database.InsertDocument(c.ID, "document", nil)
		_ = database.InsertVector(doc.ID, c.ID, []float32{float32(i) * 0.1, 0.1, 0.1, 0.1})
	}

	// Request only 2 results.
	results, err := database.SearchVectors(c.ID, []float32{0.0, 0.1, 0.1, 0.1}, nil, 2)
	if err != nil {
		t.Fatalf("SearchVectors() error = %v", err)
	}

	if len(results) != 2 {
		t.Errorf("got %d results, want 2", len(results))
	}
}

func TestSearchVectors_NoResults(t *testing.T) {
	database := testDBWithVectors(t, 4)
	c, _ := database.CreateCollection("empty-test")

	// Search with no documents in the collection.
	results, err := database.SearchVectors(c.ID, []float32{1.0, 0.0, 0.0, 0.0}, nil, 10)
	if err != nil {
		t.Fatalf("SearchVectors() error = %v", err)
	}

	if len(results) != 0 {
		t.Errorf("got %d results, want 0", len(results))
	}
}

func TestDeleteVector(t *testing.T) {
	database := testDBWithVectors(t, 4)

	c, _ := database.CreateCollection("delete-vec-test")
	doc, _ := database.InsertDocument(c.ID, "will be deleted", nil)
	_ = database.InsertVector(doc.ID, c.ID, []float32{1.0, 0.0, 0.0, 0.0})

	// Verify vector exists.
	results, err := database.SearchVectors(c.ID, []float32{1.0, 0.0, 0.0, 0.0}, nil, 10)
	if err != nil {
		t.Fatalf("SearchVectors() error = %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result before delete, got %d", len(results))
	}

	// Delete the vector.
	if err := database.DeleteVector(doc.ID); err != nil {
		t.Fatalf("DeleteVector() error = %v", err)
	}

	// Verify vector is gone.
	results, err = database.SearchVectors(c.ID, []float32{1.0, 0.0, 0.0, 0.0}, nil, 10)
	if err != nil {
		t.Fatalf("SearchVectors() after delete error = %v", err)
	}
	if len(results) != 0 {
		t.Errorf("expected 0 results after delete, got %d", len(results))
	}
}

func TestDeleteVector_NonExistent(t *testing.T) {
	database := testDBWithVectors(t, 4)

	// Deleting a non-existent vector should not error.
	if err := database.DeleteVector(99999); err != nil {
		t.Errorf("DeleteVector(non-existent) error = %v", err)
	}
}

func TestDeleteVectorsByCollection(t *testing.T) {
	database := testDBWithVectors(t, 4)

	c1, _ := database.CreateCollection("coll-1")
	c2, _ := database.CreateCollection("coll-2")

	// Insert docs with vectors in both collections.
	doc1, _ := database.InsertDocument(c1.ID, "doc in c1", nil)
	doc2, _ := database.InsertDocument(c1.ID, "another doc in c1", nil)
	doc3, _ := database.InsertDocument(c2.ID, "doc in c2", nil)

	_ = database.InsertVector(doc1.ID, c1.ID, []float32{1.0, 0.0, 0.0, 0.0})
	_ = database.InsertVector(doc2.ID, c1.ID, []float32{0.0, 1.0, 0.0, 0.0})
	_ = database.InsertVector(doc3.ID, c2.ID, []float32{0.0, 0.0, 1.0, 0.0})

	// Delete all vectors in collection 1.
	if err := database.DeleteVectorsByCollection(c1.ID); err != nil {
		t.Fatalf("DeleteVectorsByCollection() error = %v", err)
	}

	// Collection 1 should have no vector results.
	results, err := database.SearchVectors(c1.ID, []float32{1.0, 0.0, 0.0, 0.0}, nil, 10)
	if err != nil {
		t.Fatalf("SearchVectors(c1) error = %v", err)
	}
	if len(results) != 0 {
		t.Errorf("c1: got %d results after delete, want 0", len(results))
	}

	// Collection 2 should still have its vector.
	results, err = database.SearchVectors(c2.ID, []float32{0.0, 0.0, 1.0, 0.0}, nil, 10)
	if err != nil {
		t.Fatalf("SearchVectors(c2) error = %v", err)
	}
	if len(results) != 1 {
		t.Errorf("c2: got %d results after c1 delete, want 1", len(results))
	}
}
