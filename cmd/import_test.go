package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gandazgul/mnemoteca/internal/backup"
	"github.com/gandazgul/mnemoteca/internal/db"
)

func TestImportCmd_NoArgs(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"import"})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error when no args or flags provided")
	}
}

func TestImportCmd_DirAndFileConflict(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"import", "--dir", tmpDir, "somefile.jsonl"})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error when --dir used with a file argument")
	}
}

func TestImportCmd_DirAndNameConflict(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"import", "--dir", tmpDir, "--name", "foo"})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error when --dir used with --name")
	}
}

func TestImportCmd_SingleFile(t *testing.T) {
	// Reset flags that may have been set by previous tests (Cobra flag state
	// persists because rootCmd is a package-level variable).
	importCmd.Flags().Set("dir", "")  //nolint:errcheck
	importCmd.Flags().Set("name", "") //nolint:errcheck

	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))

	// Create a JSONL export file to import.
	header := backup.Header{
		Version:    backup.FormatVersion,
		Collection: "testcol",
		DocCount:   1,
	}
	doc := backup.DocRecord{
		Content: "hello world",
		Vector:  []float32{0.1, 0.2, 0.3},
	}

	headerJSON, _ := json.Marshal(header)
	docJSON, _ := json.Marshal(doc)
	exportFile := filepath.Join(tmpDir, "testcol.jsonl")
	if err := os.WriteFile(exportFile, []byte(string(headerJSON)+"\n"+string(docJSON)+"\n"), 0644); err != nil {
		t.Fatalf("writing export file: %v", err)
	}

	// Ensure vector table exists so import can work.
	database, err := db.Open(filepath.Join(tmpDir, "mnemoteca.db"))
	if err != nil {
		t.Fatalf("opening DB: %v", err)
	}
	if err := database.EnsureVectorTable(3); err != nil {
		t.Fatalf("ensuring vector table: %v", err)
	}
	database.Close() //nolint:errcheck

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"import", exportFile})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("import failed: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, "Imported 1 documents") {
		t.Errorf("unexpected output: %s", output)
	}
	if !strings.Contains(output, "testcol") {
		t.Errorf("expected collection name in output: %s", output)
	}
}

func TestImportCmd_NoEmbeddingsFile(t *testing.T) {
	importCmd.Flags().Set("dir", "")  //nolint:errcheck
	importCmd.Flags().Set("name", "") //nolint:errcheck

	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))

	// Create a JSONL export file with no vectors (simulates --no-embeddings export).
	header := backup.Header{
		Version:    backup.FormatVersion,
		Collection: "noembed-col",
		DocCount:   1,
	}
	doc := backup.DocRecord{
		Content:            "hello world",
		OriginalDocumentID: 42,
	}

	headerJSON, _ := json.Marshal(header)
	docJSON, _ := json.Marshal(doc)
	exportFile := filepath.Join(tmpDir, "noembed-col.jsonl")
	if err := os.WriteFile(exportFile, []byte(string(headerJSON)+"\n"+string(docJSON)+"\n"), 0644); err != nil {
		t.Fatalf("writing export file: %v", err)
	}

	// Ensure vector table exists.
	database, err := db.Open(filepath.Join(tmpDir, "mnemoteca.db"))
	if err != nil {
		t.Fatalf("opening DB: %v", err)
	}
	if err := database.EnsureVectorTable(3); err != nil {
		t.Fatalf("ensuring vector table: %v", err)
	}
	database.Close() //nolint:errcheck

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	// This will attempt to import a file without vectors, which triggers
	// the lazy embedder path. Since we can't actually load ONNX in tests,
	// we expect the import to fail with an embedding error. This test
	// verifies the code path is exercised (the error comes from the
	// embedder, not from the backup package).
	rootCmd.SetArgs([]string{"import", exportFile})
	err = rootCmd.Execute()
	// The command should fail because the embedder can't initialize in tests.
	if err == nil {
		// If by some miracle it works, check the output.
		output := outBuf.String()
		if !strings.Contains(output, "Imported") {
			t.Errorf("unexpected success without error: %s", output)
		}
	}
	// The error should mention embedding or setup, not "no embedding and no embedder".
	if err != nil && strings.Contains(err.Error(), "no embedding and no embedder") {
		t.Errorf("should have attempted auto-embedding, got: %v", err)
	}
}

func TestImportCmd_EmptyDir(t *testing.T) {
	// Reset flags that may have been set by previous tests.
	importCmd.Flags().Set("dir", "")  //nolint:errcheck
	importCmd.Flags().Set("name", "") //nolint:errcheck

	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))

	emptyDir := filepath.Join(tmpDir, "empty")
	if err := os.MkdirAll(emptyDir, 0755); err != nil {
		t.Fatalf("creating empty dir: %v", err)
	}

	// Ensure vector table exists.
	database, err := db.Open(filepath.Join(tmpDir, "mnemoteca.db"))
	if err != nil {
		t.Fatalf("opening DB: %v", err)
	}
	if err := database.EnsureVectorTable(3); err != nil {
		t.Fatalf("ensuring vector table: %v", err)
	}
	database.Close() //nolint:errcheck

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"import", "--dir", emptyDir})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("import --dir on empty dir failed: %v", err)
	}

	if !strings.Contains(outBuf.String(), "No .jsonl files found") {
		t.Errorf("unexpected output: %s", outBuf.String())
	}
}

func TestImportCmd_FinalMnemosyneV1FixturePreservesCollectionsMetadataTagsAndVectors(t *testing.T) {
	resetImportFlagsForTest()
	resetExportFlagsForTest(t)

	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "mnemoteca.db")
	t.Setenv("MNEMOTECA_DB_PATH", dbPath)

	fixtureDir := filepath.Join("..", "internal", "backup", "testdata", "final-mnemosyne-v1")

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"import", "--dir", fixtureDir})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("importing final Mnemosyne fixture: %v\nOutput:\n%s", err, outBuf.String())
	}
	if !strings.Contains(outBuf.String(), `Imported 2 files (2 documents total)`) {
		t.Fatalf("expected fixture directory import count, got: %s", outBuf.String())
	}

	database, err := db.Open(dbPath)
	if err != nil {
		t.Fatalf("opening imported database: %v", err)
	}
	defer database.Close() //nolint:errcheck

	assertImportedFixtureCollection(t, database, "global", "Final Mnemosyne global core memory export fixture", `{"tags":["core","global"],"source":"final-mnemosyne-fixture"}`, []string{"core", "global"})
	assertImportedFixtureCollection(t, database, "rename-project", "Final Mnemosyne project memory export fixture", `{"tags":["ordinary","rename"],"nested":{"preserved":true}}`, []string{"ordinary", "rename"})

	globalOut := filepath.Join(tmpDir, "global.out.jsonl")
	resetExportFlagsForTest(t)
	outBuf.Reset()
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"export", "--global", "--output", globalOut})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("exporting imported global fixture: %v\nOutput:\n%s", err, outBuf.String())
	}
	assertExportedFixture(t, globalOut, "global", "Final Mnemosyne global core memory export fixture", `{"tags":["core","global"],"source":"final-mnemosyne-fixture"}`, ascendingFixtureVector())

	projectOut := filepath.Join(tmpDir, "project.out.jsonl")
	resetExportFlagsForTest(t)
	outBuf.Reset()
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"export", "--name", "rename-project", "--output", projectOut})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("exporting imported project fixture: %v\nOutput:\n%s", err, outBuf.String())
	}
	assertExportedFixture(t, projectOut, "rename-project", "Final Mnemosyne project memory export fixture", `{"tags":["ordinary","rename"],"nested":{"preserved":true}}`, descendingFixtureVector())
}

func assertImportedFixtureCollection(t *testing.T, database *db.DB, name, content, metadata string, tags []string) {
	t.Helper()

	col, err := database.GetCollectionByName(name)
	if err != nil {
		t.Fatalf("getting collection %q: %v", name, err)
	}
	if col == nil {
		t.Fatalf("collection %q was not imported", name)
	}
	count, err := database.CountDocuments(col.ID, nil)
	if err != nil {
		t.Fatalf("counting collection %q: %v", name, err)
	}
	if count != 1 {
		t.Fatalf("collection %q document count = %d, want 1", name, count)
	}
	docs, err := database.ListDocuments(col.ID, nil, 0)
	if err != nil {
		t.Fatalf("listing collection %q: %v", name, err)
	}
	if len(docs) != 1 || docs[0].Content != content {
		t.Fatalf("collection %q docs = %#v, want content %q", name, docs, content)
	}
	if docs[0].Metadata == nil || *docs[0].Metadata != metadata {
		t.Fatalf("collection %q metadata = %v, want %q", name, docs[0].Metadata, metadata)
	}
	gotTags, err := database.GetTags(col.ID)
	if err != nil {
		t.Fatalf("getting tags for %q: %v", name, err)
	}
	if strings.Join(gotTags, ",") != strings.Join(tags, ",") {
		t.Fatalf("collection %q tags = %v, want %v", name, gotTags, tags)
	}
}

func assertExportedFixture(t *testing.T, path, collection, content, metadata string, wantVector []float32) {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading exported fixture %s: %v", path, err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) != 2 {
		t.Fatalf("export %s has %d JSONL lines, want 2", path, len(lines))
	}
	var header backup.Header
	if err := json.Unmarshal([]byte(lines[0]), &header); err != nil {
		t.Fatalf("parsing fixture header: %v", err)
	}
	if header.Version != backup.FormatVersion || header.Collection != collection || header.DocCount != 1 {
		t.Fatalf("header = %#v, want version %d collection %q count 1", header, backup.FormatVersion, collection)
	}
	var doc backup.DocRecord
	if err := json.Unmarshal([]byte(lines[1]), &doc); err != nil {
		t.Fatalf("parsing fixture doc: %v", err)
	}
	if doc.Content != content {
		t.Fatalf("content = %q, want %q", doc.Content, content)
	}
	if doc.Metadata == nil || *doc.Metadata != metadata {
		t.Fatalf("metadata = %v, want %q", doc.Metadata, metadata)
	}
	if len(doc.Vector) != len(wantVector) {
		t.Fatalf("vector length = %d, want %d", len(doc.Vector), len(wantVector))
	}
	for i := range wantVector {
		if doc.Vector[i] != wantVector[i] {
			t.Fatalf("vector[%d] = %f, want %f", i, doc.Vector[i], wantVector[i])
		}
	}
}

func ascendingFixtureVector() []float32 {
	vec := make([]float32, 256)
	for i := range vec {
		vec[i] = float32(i) / 1000
	}
	return vec
}

func descendingFixtureVector() []float32 {
	vec := make([]float32, 256)
	for i := range vec {
		vec[i] = float32(255-i) / 1000
	}
	return vec
}
