package cmd

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func executeStatsForTest(t *testing.T, args ...string) (string, error) {
	t.Helper()

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs(args)
	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
	})

	err := rootCmd.Execute()
	return outBuf.String(), err
}

func TestStatsCmd_EmptyDB(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "mnemoteca.db")
	t.Setenv("MNEMOTECA_DB_PATH", dbPath)

	out, err := executeStatsForTest(t, "stats")
	if err != nil {
		t.Fatalf("stats command failed: %v", err)
	}

	if !strings.Contains(out, "Mnemoteca Statistics") {
		t.Errorf("output missing title: %q", out)
	}
	if !strings.Contains(out, "Database Path:   "+dbPath) {
		t.Errorf("output should show resolved database path: %q", out)
	}
	if !strings.Contains(out, "Collections:     0") {
		t.Errorf("output should show 0 collections: %q", out)
	}
	if !strings.Contains(out, "Total Documents: 0") {
		t.Errorf("output should show 0 documents: %q", out)
	}
}

func TestStatsCmd_JSONEmptyDB(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "mnemoteca.db")
	t.Setenv("MNEMOTECA_DB_PATH", dbPath)

	out, err := executeStatsForTest(t, "stats", "--format", "json")
	if err != nil {
		t.Fatalf("stats command failed: %v", err)
	}

	var payload statsOutput
	if err := json.Unmarshal([]byte(out), &payload); err != nil {
		t.Fatalf("stats JSON did not parse: %v\n%s", err, out)
	}
	if payload.DatabasePath != dbPath {
		t.Fatalf("database_path = %q, want %q", payload.DatabasePath, dbPath)
	}
	if payload.CollectionCount != 0 {
		t.Fatalf("collection_count = %d, want 0", payload.CollectionCount)
	}
	if payload.DocumentCount != 0 {
		t.Fatalf("document_count = %d, want 0", payload.DocumentCount)
	}
}

func TestStatsCmd_JSONWithData(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "mnemoteca.db")
	t.Setenv("MNEMOTECA_DB_PATH", dbPath)

	database, err := openDB()
	if err != nil {
		t.Fatalf("openDB: %v", err)
	}

	first, err := database.CreateCollection("first")
	if err != nil {
		t.Fatalf("CreateCollection first: %v", err)
	}
	second, err := database.CreateCollection("second")
	if err != nil {
		t.Fatalf("CreateCollection second: %v", err)
	}
	_, _ = database.InsertDocument(first.ID, "doc 1", nil)
	_, _ = database.InsertDocument(first.ID, "doc 2", nil)
	_, _ = database.InsertDocument(second.ID, "doc 3", nil)
	database.Close() //nolint:errcheck

	out, err := executeStatsForTest(t, "stats", "--format", "json")
	if err != nil {
		t.Fatalf("stats command failed: %v", err)
	}

	var payload statsOutput
	if err := json.Unmarshal([]byte(out), &payload); err != nil {
		t.Fatalf("stats JSON did not parse: %v\n%s", err, out)
	}
	if payload.DatabasePath != dbPath {
		t.Fatalf("database_path = %q, want %q", payload.DatabasePath, dbPath)
	}
	if payload.CollectionCount != 2 {
		t.Fatalf("collection_count = %d, want 2", payload.CollectionCount)
	}
	if payload.DocumentCount != 3 {
		t.Fatalf("document_count = %d, want 3", payload.DocumentCount)
	}
}

func TestStatsCmd_WithData(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))

	database, err := openDB()
	if err != nil {
		t.Fatalf("openDB: %v", err)
	}

	c, _ := database.CreateCollection("test_col")
	_, _ = database.InsertDocument(c.ID, "doc 1", nil)
	_, _ = database.InsertDocument(c.ID, "doc 2", nil)
	database.Close() //nolint:errcheck

	out, err := executeStatsForTest(t, "stats", "--format", "plain")
	if err != nil {
		t.Fatalf("stats command failed: %v", err)
	}

	if !strings.Contains(out, "Collections:     1") {
		t.Errorf("output should show 1 collection: %q", out)
	}
	if !strings.Contains(out, "Total Documents: 2") {
		t.Errorf("output should show 2 documents: %q", out)
	}
}

func TestStatsCmd_InvalidFormat(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))

	_, err := executeStatsForTest(t, "stats", "--format", "xml")
	if err == nil {
		t.Fatal("stats command should fail with invalid format")
	}
	if !strings.Contains(err.Error(), "invalid format") {
		t.Fatalf("error = %v, want invalid format", err)
	}
}

func TestStatsCmd_InvalidDBPath(t *testing.T) {
	t.Setenv("MNEMOTECA_DB_PATH", "/this/path/should/not/exist/or/be/writable/db.sqlite")

	_, err := executeStatsForTest(t, "stats")
	if err == nil {
		t.Errorf("stats command should fail with invalid db path")
	}
}
