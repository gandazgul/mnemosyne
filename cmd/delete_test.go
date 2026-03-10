package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestDeleteCmd(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	db, err := openDB()
	if err != nil {
		t.Fatalf("setup db: %v", err)
	}
	col, _, _ := db.GetOrCreateCollection("col_a")
	_, _ = db.InsertDocument(col.ID, "some doc", nil)
	db.Close()

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	// Test deleting real doc
	rootCmd.SetArgs([]string{"delete", "1"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, "Deleted document") {
		t.Errorf("expected deleted message, got: %s", output)
	}

	outBuf.Reset()

	// Test invalid doc id
	rootCmd.SetArgs([]string{"delete", "not_an_id"})
	err = rootCmd.Execute()
	if err == nil {
		t.Error("expected error for invalid id")
	}

	outBuf.Reset()

	// Test non-existent doc
	rootCmd.SetArgs([]string{"delete", "999"})
	err = rootCmd.Execute()
	if err == nil {
		t.Error("expected error for non-existent doc")
	}
}
