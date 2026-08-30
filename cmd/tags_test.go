package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestTagsCmd_GlobalFreshDatabaseLazilyInitializes(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))
	resetTagsFlagsForTest(t)

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"tags", "--global", "--format", "plain"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("fresh global tags error: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, `No tags found in collection "global"`) {
		t.Fatalf("expected empty global tags output, got: %s", output)
	}
	if strings.Contains(output, "mnemoteca init") {
		t.Fatalf("global tags should not print init guidance, got: %s", output)
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

func resetTagsFlagsForTest(t *testing.T) {
	t.Helper()
	for name, value := range map[string]string{
		"name":   "",
		"global": "false",
		"format": "color",
	} {
		if err := tagsCmd.Flags().Set(name, value); err != nil {
			t.Fatalf("reset flag %s: %v", name, err)
		}
	}
}
