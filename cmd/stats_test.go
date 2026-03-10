package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatsCmd_EmptyDB(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"stats"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("stats command failed: %v", err)
	}

	out := outBuf.String()
	if !strings.Contains(out, "Mnemosyne Statistics") {
		t.Errorf("output missing title: %q", out)
	}
	if !strings.Contains(out, "Collections:     0") {
		t.Errorf("output should show 0 collections: %q", out)
	}
	if !strings.Contains(out, "Total Documents: 0") {
		t.Errorf("output should show 0 documents: %q", out)
	}
}

func TestStatsCmd_WithData(t *testing.T) {
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
	database.Close()

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"stats"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("stats command failed: %v", err)
	}

	out := outBuf.String()
	if !strings.Contains(out, "Collections:     1") {
		t.Errorf("output should show 1 collection: %q", out)
	}
	if !strings.Contains(out, "Total Documents: 2") {
		t.Errorf("output should show 2 documents: %q", out)
	}
}

func TestStatsCmd_InvalidDBPath(t *testing.T) {
	t.Setenv("MNEMOSYNE_DB_PATH", "/this/path/should/not/exist/or/be/writable/db.sqlite")

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"stats"})
	err := rootCmd.Execute()
	if err == nil {
		t.Errorf("stats command should fail with invalid db path")
	}
}
