package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestCollectionsCmd_Empty(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"collections"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("collections command failed: %v", err)
	}

	out := outBuf.String()
	if !strings.Contains(out, "No collections found.") {
		t.Errorf("expected 'No collections found.', got %q", out)
	}
}

func TestCollectionsCmd_WithData(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	// Create test DB
	database, err := openDB()
	if err != nil {
		t.Fatalf("openDB: %v", err)
	}

	c, _ := database.CreateCollection("test_col")
	_, _ = database.InsertDocument(c.ID, "doc 1", nil)
	_, _ = database.InsertDocument(c.ID, "doc 2", nil)

	c2, _ := database.CreateCollection("another_col")
	_, _ = database.InsertDocument(c2.ID, "doc 1", nil)
	database.Close() //nolint:errcheck

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"collections"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("collections command failed: %v", err)
	}

	out := outBuf.String()
	if !strings.Contains(out, "Collections (2 total)") {
		t.Errorf("output should show 2 collections: %q", out)
	}
	if !strings.Contains(out, "test_col") || !strings.Contains(out, "2") {
		t.Errorf("output should show test_col with 2 docs: %q", out)
	}
	if !strings.Contains(out, "another_col") || !strings.Contains(out, "1") {
		t.Errorf("output should show another_col with 1 doc: %q", out)
	}
}

func TestCollectionsCmd_InvalidDBPath(t *testing.T) {
	t.Setenv("MNEMOSYNE_DB_PATH", "/this/path/should/not/exist/or/be/writable/db.sqlite")

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"collections"})
	err := rootCmd.Execute()
	if err == nil {
		t.Errorf("collections command should fail with invalid db path")
	}
}
