package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestListCmd_GlobalFreshDatabaseLazilyInitializes(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))
	resetListFlagsForTest(t)

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"list", "--global", "--format", "plain"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("fresh global list error: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, `No documents in collection "global"`) {
		t.Fatalf("expected empty global list output, got: %s", output)
	}
	if strings.Contains(output, "mnemoteca init") {
		t.Fatalf("global list should not print init guidance, got: %s", output)
	}

	database, err := openDB()
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer database.Close() //nolint:errcheck
	if got := countCollectionsNamedForTest(t, database, "global"); got != 1 {
		t.Fatalf("global collection count = %d, want 1", got)
	}
}

func TestListCmd_MissingNamedCollectionRemainsStrict(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))
	resetListFlagsForTest(t)

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"list", "--name", "missing", "--format", "plain"})
	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected missing named collection error")
	}
	if !strings.Contains(err.Error(), "mnemoteca init --name missing") {
		t.Fatalf("expected valid init guidance, got %v", err)
	}

	database, openErr := openDB()
	if openErr != nil {
		t.Fatalf("open db: %v", openErr)
	}
	defer database.Close() //nolint:errcheck
	if got := countCollectionsNamedForTest(t, database, "missing"); got != 0 {
		t.Fatalf("missing collection count = %d, want 0", got)
	}
}

func TestListCmd(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))
	resetListFlagsForTest(t)

	db, err := openDB()
	if err != nil {
		t.Fatalf("setup db: %v", err)
	}
	_, _, _ = db.GetOrCreateCollection("col_a")
	db.Close() //nolint:errcheck

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

func resetListFlagsForTest(t *testing.T) {
	t.Helper()
	for name, value := range map[string]string{
		"name":   "",
		"global": "false",
		"limit":  "20",
		"format": "color",
	} {
		if err := listCmd.Flags().Set(name, value); err != nil {
			t.Fatalf("reset flag %s: %v", name, err)
		}
	}

	flag := listCmd.Flags().Lookup("tag")
	if flag == nil {
		t.Fatal("tag flag not found")
	}
	sliceValue, ok := flag.Value.(interface{ Replace([]string) error })
	if !ok {
		t.Fatal("tag flag does not support slice replacement")
	}
	if err := sliceValue.Replace(nil); err != nil {
		t.Fatalf("reset tag flag: %v", err)
	}
}
