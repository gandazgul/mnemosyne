package backup

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gandazgul/mnemosyne/internal/db"
)

// setupTestDB creates a temporary database with a collection and some test documents.
func setupTestDB(t *testing.T) (*db.DB, string) {
	t.Helper()

	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.db")

	database, err := db.Open(dbPath)
	if err != nil {
		t.Fatalf("opening test DB: %v", err)
	}

	if err := database.EnsureVectorTable(3); err != nil {
		t.Fatalf("ensuring vector table: %v", err)
	}

	t.Cleanup(func() { database.Close() }) //nolint:errcheck

	return database, tmpDir
}

func TestExportCollection(t *testing.T) {
	database, _ := setupTestDB(t)

	// Create a collection and add documents.
	col, err := database.CreateCollection("test-export")
	if err != nil {
		t.Fatalf("creating collection: %v", err)
	}

	vec1 := []float32{0.1, 0.2, 0.3}
	vec2 := []float32{0.4, 0.5, 0.6}
	meta := `{"tags":["core"]}`

	_, err = database.InsertDocumentWithVector(col.ID, "hello world", nil, vec1)
	if err != nil {
		t.Fatalf("inserting doc 1: %v", err)
	}
	_, err = database.InsertDocumentWithVector(col.ID, "goodbye world", &meta, vec2)
	if err != nil {
		t.Fatalf("inserting doc 2: %v", err)
	}

	// Export to buffer.
	var buf bytes.Buffer
	count, err := ExportCollection(&buf, database, "test-export")
	if err != nil {
		t.Fatalf("exporting: %v", err)
	}
	if count != 2 {
		t.Errorf("expected 2 exported, got %d", count)
	}

	// Parse output.
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines (header + 2 docs), got %d", len(lines))
	}

	// Check header.
	var header Header
	if err := json.Unmarshal([]byte(lines[0]), &header); err != nil {
		t.Fatalf("parsing header: %v", err)
	}
	if header.Version != FormatVersion {
		t.Errorf("expected version %d, got %d", FormatVersion, header.Version)
	}
	if header.Collection != "test-export" {
		t.Errorf("expected collection 'test-export', got %q", header.Collection)
	}
	if header.DocCount != 2 {
		t.Errorf("expected doc_count 2, got %d", header.DocCount)
	}

	// Check first document.
	var doc1 DocRecord
	if err := json.Unmarshal([]byte(lines[1]), &doc1); err != nil {
		t.Fatalf("parsing doc 1: %v", err)
	}
	if doc1.Content != "hello world" {
		t.Errorf("doc1 content = %q, want 'hello world'", doc1.Content)
	}
	if doc1.Metadata != nil {
		t.Errorf("doc1 metadata should be nil, got %v", doc1.Metadata)
	}
	if len(doc1.Vector) != 3 {
		t.Fatalf("doc1 vector length = %d, want 3", len(doc1.Vector))
	}
	// Check vector values are close (float32 precision).
	for i, want := range vec1 {
		if got := doc1.Vector[i]; got != want {
			t.Errorf("doc1 vector[%d] = %f, want %f", i, got, want)
		}
	}

	// Check second document.
	var doc2 DocRecord
	if err := json.Unmarshal([]byte(lines[2]), &doc2); err != nil {
		t.Fatalf("parsing doc 2: %v", err)
	}
	if doc2.Content != "goodbye world" {
		t.Errorf("doc2 content = %q", doc2.Content)
	}
	if doc2.Metadata == nil || *doc2.Metadata != meta {
		t.Errorf("doc2 metadata = %v, want %q", doc2.Metadata, meta)
	}
}

func TestExportCollection_NotFound(t *testing.T) {
	database, _ := setupTestDB(t)

	var buf bytes.Buffer
	_, err := ExportCollection(&buf, database, "nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent collection")
	}
}

func TestImportCollection(t *testing.T) {
	database, _ := setupTestDB(t)

	// Build JSONL content.
	header := Header{
		Version:    FormatVersion,
		ExportedAt: "2025-01-01T00:00:00Z",
		Collection: "imported-col",
		DocCount:   2,
	}
	doc1 := DocRecord{
		Content: "first doc",
		Vector:  []float32{1.0, 2.0, 3.0},
	}
	meta := `{"tags":["test"]}`
	doc2 := DocRecord{
		Content:  "second doc",
		Metadata: &meta,
		Vector:   []float32{4.0, 5.0, 6.0},
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(header) //nolint:errcheck
	enc.Encode(doc1)   //nolint:errcheck
	enc.Encode(doc2)   //nolint:errcheck

	// Import.
	hdr, count, err := ImportCollection(&buf, database, "")
	if err != nil {
		t.Fatalf("importing: %v", err)
	}
	if count != 2 {
		t.Errorf("expected 2 imported, got %d", count)
	}
	if hdr.Collection != "imported-col" {
		t.Errorf("header collection = %q", hdr.Collection)
	}

	// Verify collection exists.
	col, err := database.GetCollectionByName("imported-col")
	if err != nil {
		t.Fatalf("getting collection: %v", err)
	}
	if col == nil {
		t.Fatal("collection not found after import")
	}

	// Verify document count.
	docCount, err := database.CountDocuments(col.ID, nil)
	if err != nil {
		t.Fatalf("counting documents: %v", err)
	}
	if docCount != 2 {
		t.Errorf("document count = %d, want 2", docCount)
	}
}

func TestImportCollection_OverrideName(t *testing.T) {
	database, _ := setupTestDB(t)

	header := Header{
		Version:    FormatVersion,
		ExportedAt: "2025-01-01T00:00:00Z",
		Collection: "original-name",
		DocCount:   1,
	}
	doc := DocRecord{
		Content: "test",
		Vector:  []float32{1.0, 2.0, 3.0},
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(header) //nolint:errcheck
	enc.Encode(doc)    //nolint:errcheck

	_, count, err := ImportCollection(&buf, database, "override-name")
	if err != nil {
		t.Fatalf("importing with override: %v", err)
	}
	if count != 1 {
		t.Errorf("expected 1 imported, got %d", count)
	}

	// Original name should NOT exist.
	col, _ := database.GetCollectionByName("original-name")
	if col != nil {
		t.Error("original collection should not exist when override is used")
	}

	// Override name should exist.
	col, err = database.GetCollectionByName("override-name")
	if err != nil {
		t.Fatalf("getting overridden collection: %v", err)
	}
	if col == nil {
		t.Fatal("overridden collection not found")
	}
}

func TestImportCollection_VersionMismatch(t *testing.T) {
	database, _ := setupTestDB(t)

	header := Header{
		Version:    999,
		ExportedAt: "2025-01-01T00:00:00Z",
		Collection: "test",
		DocCount:   0,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(header) //nolint:errcheck

	_, _, err := ImportCollection(&buf, database, "")
	if err == nil {
		t.Error("expected error for version mismatch")
	}
	if !strings.Contains(err.Error(), "unsupported format version") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestImportCollection_EmptyFile(t *testing.T) {
	database, _ := setupTestDB(t)

	var buf bytes.Buffer
	_, _, err := ImportCollection(&buf, database, "")
	if err == nil {
		t.Error("expected error for empty file")
	}
}

func TestRoundTrip(t *testing.T) {
	database, _ := setupTestDB(t)

	// Create collection with documents.
	col, err := database.CreateCollection("roundtrip")
	if err != nil {
		t.Fatalf("creating collection: %v", err)
	}

	vecs := [][]float32{
		{0.1, 0.2, 0.3},
		{0.4, 0.5, 0.6},
		{0.7, 0.8, 0.9},
	}
	contents := []string{"alpha", "beta", "gamma"}

	for i, content := range contents {
		_, err := database.InsertDocumentWithVector(col.ID, content, nil, vecs[i])
		if err != nil {
			t.Fatalf("inserting doc %d: %v", i, err)
		}
	}

	// Export.
	var buf bytes.Buffer
	count, err := ExportCollection(&buf, database, "roundtrip")
	if err != nil {
		t.Fatalf("exporting: %v", err)
	}
	if count != 3 {
		t.Errorf("exported %d, want 3", count)
	}

	// Import into new collection name.
	importBuf := bytes.NewReader(buf.Bytes())
	_, imported, err := ImportCollection(importBuf, database, "roundtrip-copy")
	if err != nil {
		t.Fatalf("importing: %v", err)
	}
	if imported != 3 {
		t.Errorf("imported %d, want 3", imported)
	}

	// Verify the copy has the same documents.
	copiedCol, err := database.GetCollectionByName("roundtrip-copy")
	if err != nil {
		t.Fatalf("getting copied collection: %v", err)
	}
	if copiedCol == nil {
		t.Fatal("copied collection not found")
	}

	docCount, err := database.CountDocuments(copiedCol.ID, nil)
	if err != nil {
		t.Fatalf("counting: %v", err)
	}
	if docCount != 3 {
		t.Errorf("copied doc count = %d, want 3", docCount)
	}

	// Verify documents content.
	docs, err := database.ListDocuments(copiedCol.ID, nil, 0)
	if err != nil {
		t.Fatalf("listing copied docs: %v", err)
	}
	if len(docs) != 3 {
		t.Fatalf("listed %d docs, want 3", len(docs))
	}

	// Documents are returned newest first, so reverse order.
	contentSet := make(map[string]bool)
	for _, d := range docs {
		contentSet[d.Content] = true
	}
	for _, c := range contents {
		if !contentSet[c] {
			t.Errorf("missing content %q in imported docs", c)
		}
	}
}
