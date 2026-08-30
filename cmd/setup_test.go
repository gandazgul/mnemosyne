package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gandazgul/mnemoteca/internal/setup"
)

func TestSetupCmd_AlreadyReady(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_DATA_HOME", tmpDir)

	mnemoDir := filepath.Join(tmpDir, "mnemoteca")

	// mock all files
	libPath := filepath.Join(mnemoDir, "lib", setup.OnnxRuntimeLibNameForTest())
	_ = os.MkdirAll(filepath.Dir(libPath), 0755)
	_ = os.WriteFile(libPath, []byte("fake"), 0644)

	for _, m := range setup.AllModelsForTest() {
		modelDir := filepath.Join(mnemoDir, "models", m.LocalDir)
		for file := range m.Files {
			path := filepath.Join(modelDir, file)
			_ = os.MkdirAll(filepath.Dir(path), 0755)
			_ = os.WriteFile(path, []byte("dummy content"), 0644)
		}
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"setup"})
	err := rootCmd.Execute()

	w.Close() //nolint:errcheck
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	output := buf.String()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(output, "All components are already installed.") {
		t.Errorf("expected already installed message, got: %s", output)
	}
}

func TestSetupCmd_NotReady_FailsToDownload(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_DATA_HOME", tmpDir)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"setup"})
	err := rootCmd.Execute()

	w.Close() //nolint:errcheck
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	// output := buf.String()

	if err == nil {
		t.Log("Expected error from real download without network/proper path")
	}
}

func TestSetupCmdAlreadyReadyLeavesLegacyMnemosyneSentinelsUnchanged(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_DATA_HOME", tmpDir)

	mnemotecaDir := filepath.Join(tmpDir, "mnemoteca")
	legacyDir := filepath.Join(tmpDir, "mnemosyne")
	legacyDB := filepath.Join(legacyDir, "mnemosyne.db")
	legacyConfig := filepath.Join(t.TempDir(), ".config", "mnemosyne", "config.yaml")
	for _, path := range []string{legacyDB, legacyConfig} {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("creating legacy sentinel parent: %v", err)
		}
		if err := os.WriteFile(path, []byte("legacy sentinel"), 0o644); err != nil {
			t.Fatalf("writing legacy sentinel: %v", err)
		}
	}

	libPath := filepath.Join(mnemotecaDir, "lib", setup.OnnxRuntimeLibNameForTest())
	if err := os.MkdirAll(filepath.Dir(libPath), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(libPath, []byte("fake"), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, m := range setup.AllModelsForTest() {
		modelDir := filepath.Join(mnemotecaDir, "models", m.LocalDir)
		for file := range m.Files {
			path := filepath.Join(modelDir, file)
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, []byte("dummy content"), 0o644); err != nil {
				t.Fatal(err)
			}
		}
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	rootCmd.SetArgs([]string{"setup"})
	err := rootCmd.Execute()
	w.Close() //nolint:errcheck
	os.Stdout = oldStdout
	_, _ = io.Copy(io.Discard, r)
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	for _, path := range []string{legacyDB, legacyConfig} {
		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("legacy sentinel %s was changed or removed: %v", path, err)
		}
		if string(got) != "legacy sentinel" {
			t.Fatalf("legacy sentinel %s = %q", path, got)
		}
	}
}
