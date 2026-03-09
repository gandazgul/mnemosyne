package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestListCmd(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	db, err := openDB()
	if err != nil {
		t.Fatalf("setup db: %v", err)
	}
	db.GetOrCreateCollection("col_a")
	db.Close()

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	
	// Since list defaults to current dir (cmd), pass name explicitly
	rootCmd.SetArgs([]string{"list", "--name", "col_a"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	output := outBuf.String()
	if !strings.Contains(output, `No documents in collection "col_a"`) {
		t.Errorf("expected collection name in output, got: %s", output)
	}
}
