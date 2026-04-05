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

func TestExportCmd_NoFlags(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"export"})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error when no flags provided")
	}
}

func TestExportCmd_CollectionNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"export", "--name", "nonexistent"})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error for nonexistent collection")
	}
}

func TestExportCmd_SingleCollection(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	// Create collection with a document.
	database, err := db.Open(filepath.Join(tmpDir, "mnemosyne.db"))
	if err != nil {
		t.Fatalf("opening DB: %v", err)
	}
	if err := database.EnsureVectorTable(3); err != nil {
		t.Fatalf("ensuring vector table: %v", err)
	}
	col, err := database.CreateCollection("myproject")
	if err != nil {
		t.Fatalf("creating collection: %v", err)
	}
	_, err = database.InsertDocumentWithVector(col.ID, "test content", nil, []float32{1, 2, 3})
	if err != nil {
		t.Fatalf("inserting: %v", err)
	}
	database.Close() //nolint:errcheck

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	outputPath := filepath.Join(tmpDir, "out.jsonl")
	rootCmd.SetArgs([]string{"export", "--name", "myproject", "-o", outputPath})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("export failed: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, "Exported 1 documents") {
		t.Errorf("unexpected output: %s", output)
	}

	// Verify file exists and has valid JSONL.
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("reading output file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines (header + 1 doc), got %d", len(lines))
	}

	var header backup.Header
	if err := json.Unmarshal([]byte(lines[0]), &header); err != nil {
		t.Fatalf("parsing header: %v", err)
	}
	if header.Collection != "myproject" {
		t.Errorf("header collection = %q", header.Collection)
	}
}

// Tests using --all come last because Cobra bool flags persist across test
// runs in the same process (rootCmd is a package-level variable).

func TestExportCmd_AllEmpty(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"export", "--all", "--yes", "--name", ""})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("export --all failed: %v", err)
	}

	if !strings.Contains(outBuf.String(), "No collections to export") {
		t.Errorf("unexpected output: %s", outBuf.String())
	}
}

func TestExportCmd_ConflictingFlags(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"export", "--all", "--name", "foo"})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error when --all and --name both provided")
	}
}
