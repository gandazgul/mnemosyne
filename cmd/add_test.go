package cmd

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestAddCmd_File(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	// Since adding requires opening the embedder, which is hard to mock cleanly
	// without 500mb models or breaking `main_test.go`, we will test the error cases
	// for `addCmd` which still provides good coverage.

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	// Error: Neither text, file, nor stdin provided
	rootCmd.SetArgs([]string{"add", "--name", "col_a"})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error when no content provided")
	}

	// Error: File reading fails
	rootCmd.SetArgs([]string{"add", "--name", "col_a", "--file", filepath.Join(tmpDir, "does-not-exist.txt")})
	err = rootCmd.Execute()
	if err == nil {
		t.Error("expected error when file doesn't exist")
	}
}
