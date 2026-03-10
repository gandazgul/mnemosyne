package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestForgetCmd(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	db, err := openDB()
	if err != nil {
		t.Fatalf("setup db: %v", err)
	}
	_, _, _ = db.GetOrCreateCollection("col_a")
	db.Close() //nolint:errcheck

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	// Test --yes flag for automatic confirmation
	rootCmd.SetArgs([]string{"forget", "--name", "col_a", "--yes"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, "Deleted collection ") {
		t.Errorf("expected deleted message, got: %s", output)
	}

	// Test missing collection
	rootCmd.SetArgs([]string{"forget", "--name", "non_existent", "--yes"})
	err = rootCmd.Execute()
	if err == nil {
		t.Error("expected error for missing collection")
	}
}
