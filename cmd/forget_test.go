package cmd

import (
	"bytes"
	"os"
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

	// Test interactive non-TTY input with Unix line endings (\n)
	t.Run("InteractiveUnix", func(t *testing.T) {
		db, _ := openDB()
		_, _, _ = db.GetOrCreateCollection("col_unix")
		db.Close() //nolint:errcheck

		outBuf.Reset()
		rootCmd.SetArgs([]string{"forget", "--name", "col_unix"})

		// Provide the input as if typed
		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		r, w, _ := os.Pipe()
		os.Stdin = r
		_, _ = w.Write([]byte("col_unix\n"))
		_ = w.Close()

		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("unexpected error for unix interactive input: %v", err)
		}
		output := outBuf.String()
		if !strings.Contains(output, `Deleted collection "col_unix"`) {
			t.Errorf("expected deleted message, got: %s", output)
		}
	})

	// Test interactive non-TTY input with Windows/PTY carriage-return line endings (\r\n or just \r)
	t.Run("InteractiveWindowsCR", func(t *testing.T) {
		db, _ := openDB()
		_, _, _ = db.GetOrCreateCollection("col_win")
		db.Close() //nolint:errcheck

		outBuf.Reset()
		rootCmd.SetArgs([]string{"forget", "--name", "col_win"})

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		r, w, _ := os.Pipe()
		os.Stdin = r
		// Write with just \r to mimic a raw terminal Enter key missing an ICRNL translation,
		// or \r\n to mimic a windows PTY.
		_, _ = w.Write([]byte("col_win\r"))
		_ = w.Close()

		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("unexpected error for windows/CR interactive input: %v", err)
		}

		output := outBuf.String()
		if !strings.Contains(output, `Deleted collection "col_win"`) {
			t.Errorf("expected deleted message, got: %s", output)
		}
	})

	// Test mismatch interactive input
	t.Run("InteractiveMismatch", func(t *testing.T) {
		db, _ := openDB()
		_, _, _ = db.GetOrCreateCollection("col_mismatch")
		db.Close() //nolint:errcheck

		outBuf.Reset()
		rootCmd.SetArgs([]string{"forget", "--name", "col_mismatch"})
		_ = forgetCmd.Flags().Set("yes", "false")

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		r, w, _ := os.Pipe()
		os.Stdin = r
		_, _ = w.Write([]byte("wrong_name\n"))
		_ = w.Close()

		err := rootCmd.Execute()
		if err == nil {
			t.Errorf("expected an error for mismatched confirmation name")
		} else if !strings.Contains(err.Error(), "confirmation did not match") {
			t.Errorf("expected 'confirmation did not match', got: %v", err)
		}
	})

	// Test missing collection
	t.Run("MissingCollection", func(t *testing.T) {
		outBuf.Reset()
		rootCmd.SetArgs([]string{"forget", "--name", "non_existent", "--yes"})
		err := rootCmd.Execute()
		if err == nil {
			t.Error("expected error for missing collection")
		}
	})
}
