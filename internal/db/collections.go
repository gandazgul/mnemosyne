package db

import (
	"database/sql"
	"fmt"
	"time"
)

// Collection represents a named group of documents.
type Collection struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

// CreateCollection inserts a new collection with the given name.
// Returns the created collection, or an error if the name already exists.
func (db *DB) CreateCollection(name string) (*Collection, error) {
	result, err := db.conn.Exec(
		"INSERT INTO collections (name) VALUES (?)",
		name,
	)
	if err != nil {
		return nil, fmt.Errorf("inserting collection %q: %w", name, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("getting last insert id: %w", err)
	}

	return &Collection{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}

// GetCollectionByName retrieves a collection by its name.
// Returns nil, nil if the collection does not exist.
func (db *DB) GetCollectionByName(name string) (*Collection, error) {
	var c Collection
	err := db.conn.QueryRow(
		"SELECT id, name, created_at FROM collections WHERE name = ?",
		name,
	).Scan(&c.ID, &c.Name, &c.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("querying collection %q: %w", name, err)
	}

	return &c, nil
}

// GetOrCreateCollection retrieves a collection by name, creating it if it doesn't exist.
// This is useful for the init command which should be idempotent.
func (db *DB) GetOrCreateCollection(name string) (*Collection, bool, error) {
	c, err := db.GetCollectionByName(name)
	if err != nil {
		return nil, false, err
	}
	if c != nil {
		return c, false, nil // already existed
	}

	c, err = db.CreateCollection(name)
	if err != nil {
		return nil, false, err
	}

	return c, true, nil // newly created
}

// CollectionInfo holds collection details plus document count.
type CollectionInfo struct {
	Collection
	DocumentCount int64
}

// ListCollections returns all collections with their document counts.
func (db *DB) ListCollections() ([]CollectionInfo, error) {
	rows, err := db.conn.Query(`
		SELECT c.id, c.name, c.created_at, COUNT(d.id) as doc_count
		FROM collections c
		LEFT JOIN documents d ON d.collection_id = c.id
		GROUP BY c.id
		ORDER BY c.name
	`)
	if err != nil {
		return nil, fmt.Errorf("listing collections: %w", err)
	}
	defer rows.Close()

	var collections []CollectionInfo
	for rows.Next() {
		var ci CollectionInfo
		if err := rows.Scan(&ci.ID, &ci.Name, &ci.CreatedAt, &ci.DocumentCount); err != nil {
			return nil, fmt.Errorf("scanning collection row: %w", err)
		}
		collections = append(collections, ci)
	}

	return collections, rows.Err()
}

// DeleteCollection removes a collection and all its documents (via CASCADE).
func (db *DB) DeleteCollection(name string) error {
	result, err := db.conn.Exec("DELETE FROM collections WHERE name = ?", name)
	if err != nil {
		return fmt.Errorf("deleting collection %q: %w", name, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("collection %q not found", name)
	}

	return nil
}
