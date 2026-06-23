package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDiscoverClaudeMemorySources_DefaultProjectAndAutoMemory(t *testing.T) {
	tmpDir := t.TempDir()
	home := filepath.Join(tmpDir, "home")
	project := filepath.Join(tmpDir, "work", "project")
	if err := os.MkdirAll(filepath.Join(project, ".claude", "rules"), 0755); err != nil {
		t.Fatalf("creating project dirs: %v", err)
	}

	writeFile(t, filepath.Join(project, "CLAUDE.md"), "# Project memory")
	writeFile(t, filepath.Join(project, ".claude", "CLAUDE.md"), "# Alternate project memory")
	writeFile(t, filepath.Join(project, "CLAUDE.local.md"), "# Local memory")
	writeFile(t, filepath.Join(project, ".claude", "rules", "testing.md"), "# Testing rule")
	writeFile(t, filepath.Join(project, ".claude", "rules", "ignore.txt"), "ignored")

	autoMemoryDir := filepath.Join(home, ".claude", "projects", claudeProjectDirName(project), "memory")
	writeFile(t, filepath.Join(autoMemoryDir, "MEMORY.md"), "# Auto memory")

	writeFile(t, filepath.Join(home, ".claude", "CLAUDE.md"), "# User memory")

	sources, err := discoverClaudeMemorySources(project, home, nil, false)
	if err != nil {
		t.Fatalf("discovering sources: %v", err)
	}

	got := sourcePaths(sources)
	want := []string{
		filepath.Join(home, ".claude", "projects", claudeProjectDirName(project), "memory", "MEMORY.md"),
		filepath.Join(project, ".claude", "CLAUDE.md"),
		filepath.Join(project, "CLAUDE.local.md"),
		filepath.Join(project, "CLAUDE.md"),
		filepath.Join(project, ".claude", "rules", "testing.md"),
	}
	assertSameStrings(t, got, want)
}

func TestDiscoverClaudeMemorySources_IncludeUser(t *testing.T) {
	tmpDir := t.TempDir()
	home := filepath.Join(tmpDir, "home")
	project := filepath.Join(tmpDir, "project")
	if err := os.MkdirAll(project, 0755); err != nil {
		t.Fatalf("creating project: %v", err)
	}

	writeFile(t, filepath.Join(home, ".claude", "CLAUDE.md"), "# User memory")
	writeFile(t, filepath.Join(home, ".claude", "rules", "style.md"), "# Style rule")

	sources, err := discoverClaudeMemorySources(project, home, nil, true)
	if err != nil {
		t.Fatalf("discovering sources: %v", err)
	}

	got := sourcePaths(sources)
	want := []string{
		filepath.Join(home, ".claude", "CLAUDE.md"),
		filepath.Join(home, ".claude", "rules", "style.md"),
	}
	assertSameStrings(t, got, want)
}

func TestDiscoverClaudeMemorySources_ExplicitPath(t *testing.T) {
	tmpDir := t.TempDir()
	home := filepath.Join(tmpDir, "home")
	project := filepath.Join(tmpDir, "project")
	explicitDir := filepath.Join(tmpDir, "explicit")
	if err := os.MkdirAll(project, 0755); err != nil {
		t.Fatalf("creating project: %v", err)
	}

	writeFile(t, filepath.Join(project, "CLAUDE.md"), "# Project memory")
	writeFile(t, filepath.Join(explicitDir, "one.md"), "# One")
	writeFile(t, filepath.Join(explicitDir, "nested", "two.md"), "# Two")
	writeFile(t, filepath.Join(explicitDir, "ignored.txt"), "ignored")

	sources, err := discoverClaudeMemorySources(project, home, []string{explicitDir}, true)
	if err != nil {
		t.Fatalf("discovering sources: %v", err)
	}

	got := sourcePaths(sources)
	want := []string{
		filepath.Join(explicitDir, "nested", "two.md"),
		filepath.Join(explicitDir, "one.md"),
	}
	assertSameStrings(t, got, want)
}

func TestImportCmdAgentClaudeDryRunPrintsNonDestructiveMessage(t *testing.T) {
	tmpDir := t.TempDir()
	home := filepath.Join(tmpDir, "home")
	project := filepath.Join(tmpDir, "project")
	if err := os.MkdirAll(project, 0755); err != nil {
		t.Fatalf("creating project: %v", err)
	}
	writeFile(t, filepath.Join(project, "CLAUDE.md"), "# Project memory")

	t.Setenv("HOME", home)
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "should-not-be-created", "mnemosyne.db"))
	t.Chdir(project)

	resetImportFlagsForTest()
	defer resetImportFlagsForTest()

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"import", "--agent", "claude", "--dry-run"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("dry run failed: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, "Non-destructive import:") {
		t.Fatalf("expected non-destructive message, got: %s", output)
	}
	if !strings.Contains(output, "It will not edit, delete, move, overwrite, or truncate any Claude Code files or existing Mnemosyne memories.") {
		t.Fatalf("expected explicit safety details, got: %s", output)
	}
	if !strings.Contains(output, "Dry run: no database writes or embedding calls will be made.") {
		t.Fatalf("expected dry-run message, got: %s", output)
	}
	if _, err := os.Stat(filepath.Join(tmpDir, "should-not-be-created")); !os.IsNotExist(err) {
		t.Fatalf("dry run should not create database directory, stat err: %v", err)
	}
}

func TestImportCmdAgentRejectsFileArgument(t *testing.T) {
	resetImportFlagsForTest()
	defer resetImportFlagsForTest()

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"import", "memories.jsonl", "--agent", "claude"})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected --agent with file argument to fail")
	}
	if !strings.Contains(err.Error(), "cannot use --agent with a file argument") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestImportCmdUnsupportedAgent(t *testing.T) {
	resetImportFlagsForTest()
	defer resetImportFlagsForTest()

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"import", "--agent", "unknown"})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected unsupported agent to fail")
	}
	if !strings.Contains(err.Error(), `unsupported import agent "unknown"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func sourcePaths(sources []claudeMemorySource) []string {
	paths := make([]string, 0, len(sources))
	for _, source := range sources {
		paths = append(paths, source.Path)
	}
	return paths
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("creating parent directory for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("writing %s: %v", path, err)
	}
}

func assertSameStrings(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %d paths, want %d\ngot:  %v\nwant: %v", len(got), len(want), got, want)
	}
	wantSet := map[string]bool{}
	for _, value := range want {
		wantSet[value] = true
	}
	for _, value := range got {
		if !wantSet[value] {
			t.Fatalf("unexpected value %q\ngot:  %v\nwant: %v", value, got, want)
		}
	}
}

func resetImportFlagsForTest() {
	importCmd.Flags().Set("name", "")              //nolint:errcheck
	importCmd.Flags().Set("dir", "")               //nolint:errcheck
	importCmd.Flags().Set("agent", "")             //nolint:errcheck
	importCmd.Flags().Set("global", "false")       //nolint:errcheck
	importCmd.Flags().Set("dry-run", "false")      //nolint:errcheck
	importCmd.Flags().Set("include-user", "false") //nolint:errcheck
}
