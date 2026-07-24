package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	dbpkg "github.com/gandazgul/mnemosyne/internal/db"
)

func TestResolveCollectionName(t *testing.T) {
	// 1. Name provided explicitly
	name, err := resolveCollectionName("explicit_name", false)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if name != "explicit_name" {
		t.Errorf("expected 'explicit_name', got %v", name)
	}

	// 2. Global flag provided
	name, err = resolveCollectionName("", true)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if name != "global" {
		t.Errorf("expected 'global', got %v", name)
	}

	// 3. Global flag and conflicting name provided
	_, err = resolveCollectionName("other_name", true)
	if err == nil {
		t.Errorf("expected error for conflicting flags, got nil")
	}

	// 4. Global flag and matching name provided (now an error)
	_, err = resolveCollectionName("global", true)
	if err == nil {
		t.Errorf("expected error for using both flags, got nil")
	}

	// 5. Name is 'global', no global flag
	_, err = resolveCollectionName("global", false)
	if err == nil {
		t.Errorf("expected error when using 'global' as name, got nil")
	}

	// 6. No name provided (uses cwd)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %v", err)
	}
	expected := filepath.Base(cwd)

	name, err = resolveCollectionName("", false)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if name != expected {
		t.Errorf("expected '%s', got '%s'", expected, name)
	}
}

func TestOpenDB(t *testing.T) {
	// Temporarily override DBPath to a test file so we don't clobber the real DB
	tmpDB := filepath.Join(t.TempDir(), "test.db")
	t.Setenv("MNEMOSYNE_DB_PATH", tmpDB)

	db, err := openDB()
	if err != nil {
		t.Fatalf("expected no error opening test DB, got %v", err)
	}
	defer db.Close() //nolint:errcheck

	if db == nil {
		t.Error("expected valid DB instance, got nil")
	}
}

func TestGetSelectedCollection_LazilyCreatesGlobal(t *testing.T) {
	tmpDB := filepath.Join(t.TempDir(), "mnemosyne.db")
	t.Setenv("MNEMOSYNE_DB_PATH", tmpDB)

	database, err := openDB()
	if err != nil {
		t.Fatalf("openDB: %v", err)
	}
	defer database.Close() //nolint:errcheck

	first, err := getSelectedCollection(database, "global", true)
	if err != nil {
		t.Fatalf("first global lookup: %v", err)
	}
	if first.Name != "global" {
		t.Fatalf("collection name = %q, want global", first.Name)
	}

	second, err := getSelectedCollection(database, "global", true)
	if err != nil {
		t.Fatalf("second global lookup: %v", err)
	}
	if second.ID != first.ID {
		t.Fatalf("second lookup ID = %d, want %d", second.ID, first.ID)
	}

	if got := countCollectionsNamedForTest(t, database, "global"); got != 1 {
		t.Fatalf("global collection count = %d, want 1", got)
	}
}

func TestGetSelectedCollection_MissingNamedCollectionIsStrict(t *testing.T) {
	tmpDB := filepath.Join(t.TempDir(), "mnemosyne.db")
	t.Setenv("MNEMOSYNE_DB_PATH", tmpDB)

	database, err := openDB()
	if err != nil {
		t.Fatalf("openDB: %v", err)
	}
	defer database.Close() //nolint:errcheck

	_, err = getSelectedCollection(database, "missing", false)
	if err == nil {
		t.Fatal("expected missing named collection error")
	}
	if !strings.Contains(err.Error(), "mnemosyne init --name missing") {
		t.Fatalf("expected init guidance for missing named collection, got %v", err)
	}
	if got := countCollectionsNamedForTest(t, database, "missing"); got != 0 {
		t.Fatalf("missing named collection count = %d, want 0", got)
	}
}

func countCollectionsNamedForTest(t *testing.T, database *dbpkg.DB, name string) int {
	t.Helper()
	collections, err := database.ListCollections()
	if err != nil {
		t.Fatalf("ListCollections: %v", err)
	}
	count := 0
	for _, collection := range collections {
		if collection.Name == name {
			count++
		}
	}
	return count
}
