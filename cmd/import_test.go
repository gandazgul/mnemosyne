package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gandazgul/mnemosyne/internal/backup"
	"github.com/gandazgul/mnemosyne/internal/db"
)

func TestImportCmd_NoArgs(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

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
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

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
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

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
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

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
	database, err := db.Open(filepath.Join(tmpDir, "mnemosyne.db"))
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
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

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
	database, err := db.Open(filepath.Join(tmpDir, "mnemosyne.db"))
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
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	emptyDir := filepath.Join(tmpDir, "empty")
	if err := os.MkdirAll(emptyDir, 0755); err != nil {
		t.Fatalf("creating empty dir: %v", err)
	}

	// Ensure vector table exists.
	database, err := db.Open(filepath.Join(tmpDir, "mnemosyne.db"))
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
