package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestUpdateCmd_TagOperations(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	database, err := openDB()
	if err != nil {
		t.Fatalf("setup db: %v", err)
	}
	col, _, _ := database.GetOrCreateCollection("docs")
	meta := `{"tags":["core","old"],"source":"import"}`
	doc, _ := database.InsertDocument(col.ID, "some doc", &meta)
	_ = database.Close()

	out, err := executeUpdateForTest(t, "update", "--name", "docs", strconvID(t, doc.ID), "--tag", "old", "--tag", "new")
	if err != nil {
		t.Fatalf("additive update error: %v", err)
	}
	if !strings.Contains(out, "Updated document") {
		t.Fatalf("expected success output, got %q", out)
	}
	assertDocumentMetadata(t, doc.ID, []string{"core", "old", "new"}, map[string]string{"source": "import"})

	_, err = executeUpdateForTest(t, "update", "--name", "docs", strconvID(t, doc.ID), "--replace-tags", "--tag", "fresh", "--tag", "fresh")
	if err != nil {
		t.Fatalf("replacement update error: %v", err)
	}
	assertDocumentMetadata(t, doc.ID, []string{"fresh"}, map[string]string{"source": "import"})

	_, err = executeUpdateForTest(t, "update", "--name", "docs", strconvID(t, doc.ID), "--replace-tags")
	if err != nil {
		t.Fatalf("clear update error: %v", err)
	}
	assertDocumentMetadata(t, doc.ID, nil, map[string]string{"source": "import"})
}

func TestUpdateCmd_RejectsMissingAndWrongCollection(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	database, err := openDB()
	if err != nil {
		t.Fatalf("setup db: %v", err)
	}
	col1, _, _ := database.GetOrCreateCollection("docs-1")
	_, _, _ = database.GetOrCreateCollection("docs-2")
	doc, _ := database.InsertDocument(col1.ID, "some doc", nil)
	_ = database.Close()

	_, err = executeUpdateForTest(t, "update", "--name", "missing", strconvID(t, doc.ID), "--tag", "new")
	if err == nil {
		t.Fatal("expected error for unknown collection")
	}

	_, err = executeUpdateForTest(t, "update", "--name", "docs-2", strconvID(t, doc.ID), "--tag", "new")
	if err == nil {
		t.Fatal("expected error for wrong collection")
	}

	_, err = executeUpdateForTest(t, "update", "--name", "docs-1", "999", "--tag", "new")
	if err == nil {
		t.Fatal("expected error for missing document")
	}

	updatedDB, _ := openDB()
	defer updatedDB.Close() //nolint:errcheck
	updated, _ := updatedDB.GetDocumentByID(doc.ID)
	if updated.Metadata != nil {
		t.Fatalf("expected wrong-collection update to leave metadata nil, got %v", *updated.Metadata)
	}
}

func TestUpdateCmd_RejectsInvalidInputs(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	if _, err := executeUpdateForTest(t, "update", "abc", "--tag", "new"); err == nil {
		t.Fatal("expected invalid ID error")
	}
	if _, err := executeUpdateForTest(t, "update", "1"); err == nil {
		t.Fatal("expected no-op error")
	}

	file := filepath.Join(tmpDir, "content.txt")
	if err := os.WriteFile(file, []byte("file content"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}
	if _, err := executeUpdateForTest(t, "update", "1", "text", "--file", file); err == nil {
		t.Fatal("expected conflicting content source error")
	}

	empty := filepath.Join(tmpDir, "empty.txt")
	if err := os.WriteFile(empty, []byte(" \n\t"), 0644); err != nil {
		t.Fatalf("write empty file: %v", err)
	}
	if _, err := executeUpdateForTest(t, "update", "1", "--file", empty); err == nil {
		t.Fatal("expected empty file content error")
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	os.Stdin = r
	_, _ = w.Write([]byte(" \n"))
	_ = w.Close()
	if _, err := executeUpdateForTest(t, "update", "1", "--stdin"); err == nil {
		t.Fatal("expected empty stdin content error")
	}
}

func TestUpdateCmd_RejectsMalformedMetadataForTagEdit(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	database, err := openDB()
	if err != nil {
		t.Fatalf("setup db: %v", err)
	}
	col, _, _ := database.GetOrCreateCollection("docs")
	bad := `not-json`
	doc, _ := database.InsertDocument(col.ID, "some doc", &bad)
	_ = database.Close()

	_, err = executeUpdateForTest(t, "update", "--name", "docs", strconvID(t, doc.ID), "--tag", "new")
	if err == nil {
		t.Fatal("expected malformed metadata error")
	}

	updatedDB, _ := openDB()
	defer updatedDB.Close() //nolint:errcheck
	updated, _ := updatedDB.GetDocumentByID(doc.ID)
	if updated.Metadata == nil || *updated.Metadata != bad {
		t.Fatalf("expected metadata to remain %q, got %v", bad, updated.Metadata)
	}
}

func executeUpdateForTest(t *testing.T, args ...string) (string, error) {
	t.Helper()
	resetUpdateFlagsForTest(t)

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs(args)

	err := rootCmd.Execute()
	rootCmd.SetArgs(nil)
	return outBuf.String(), err
}

func resetUpdateFlagsForTest(t *testing.T) {
	t.Helper()
	for name, value := range map[string]string{
		"name":         "",
		"global":       "false",
		"file":         "",
		"stdin":        "false",
		"replace-tags": "false",
	} {
		if err := updateCmd.Flags().Set(name, value); err != nil {
			t.Fatalf("reset flag %s: %v", name, err)
		}
	}

	flag := updateCmd.Flags().Lookup("tag")
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

func assertDocumentMetadata(t *testing.T, id int64, wantTags []string, wantStrings map[string]string) {
	t.Helper()
	database, err := openDB()
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer database.Close() //nolint:errcheck

	doc, err := database.GetDocumentByID(id)
	if err != nil {
		t.Fatalf("GetDocumentByID() error = %v", err)
	}
	if doc == nil || doc.Metadata == nil {
		t.Fatalf("expected metadata for document %d", id)
	}

	var metadata struct {
		Tags   []string `json:"tags"`
		Source string   `json:"source"`
	}
	if err := json.Unmarshal([]byte(*doc.Metadata), &metadata); err != nil {
		t.Fatalf("metadata is not JSON: %v", err)
	}
	if len(metadata.Tags) != len(wantTags) {
		t.Fatalf("tags = %#v, want %#v", metadata.Tags, wantTags)
	}
	for i := range wantTags {
		if metadata.Tags[i] != wantTags[i] {
			t.Fatalf("tags = %#v, want %#v", metadata.Tags, wantTags)
		}
	}
	if source, ok := wantStrings["source"]; ok && metadata.Source != source {
		t.Fatalf("source = %q, want %q", metadata.Source, source)
	}
}

func strconvID(t *testing.T, id int64) string {
	t.Helper()
	return strconv.FormatInt(id, 10)
}
