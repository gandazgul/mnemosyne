package db

import (
	"database/sql"
	"fmt"
	"time"
)

// Document represents a stored document with its content and metadata.
type Document struct {
	ID           int64
	CollectionID int64
	Content      string
	Metadata     *string // nil if no metadata
	CreatedAt    time.Time
}

// InsertDocument stores a new document in the given collection.
// The metadata parameter is optional and stored as a JSON string.
func (db *DB) InsertDocument(collectionID int64, content string, metadata *string) (*Document, error) {
	result, err := db.conn.Exec(
		"INSERT INTO documents (collection_id, content, metadata) VALUES (?, ?, ?)",
		collectionID, content, metadata,
	)
	if err != nil {
		return nil, fmt.Errorf("inserting document: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("getting last insert id: %w", err)
	}

	return &Document{
		ID:           id,
		CollectionID: collectionID,
		Content:      content,
		Metadata:     metadata,
		CreatedAt:    time.Now(),
	}, nil
}

// InsertDocumentWithVector stores a new document and its vector embedding atomically.
func (db *DB) InsertDocumentWithVector(collectionID int64, content string, metadata *string, embedding []float32) (*Document, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return nil, fmt.Errorf("beginning transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	result, err := tx.Exec(
		"INSERT INTO documents (collection_id, content, metadata) VALUES (?, ?, ?)",
		collectionID, content, metadata,
	)
	if err != nil {
		return nil, fmt.Errorf("inserting document: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("getting last insert id: %w", err)
	}

	vecBlob := SerializeFloat32(embedding)

	_, err = tx.Exec(
		"INSERT INTO docs_vec (document_id, collection_id, embedding) VALUES (?, ?, ?)",
		id, collectionID, vecBlob,
	)
	if err != nil {
		return nil, fmt.Errorf("inserting vector: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return &Document{
		ID:           id,
		CollectionID: collectionID,
		Content:      content,
		Metadata:     metadata,
		CreatedAt:    time.Now(),
	}, nil
}

// GetDocumentByID retrieves a single document by its ID.
// Returns nil, nil if the document does not exist.
func (db *DB) GetDocumentByID(id int64) (*Document, error) {
	var doc Document
	err := db.conn.QueryRow(
		"SELECT id, collection_id, content, metadata, created_at FROM documents WHERE id = ?",
		id,
	).Scan(&doc.ID, &doc.CollectionID, &doc.Content, &doc.Metadata, &doc.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("querying document %d: %w", id, err)
	}

	return &doc, nil
}

// ListDocuments returns documents in a collection, ordered by creation time (newest first).
// Use limit <= 0 for no limit.
func (db *DB) ListDocuments(collectionID int64, limit int) (docs []Document, err error) {
	query := `
		SELECT id, collection_id, content, metadata, created_at
		FROM documents
		WHERE collection_id = ?
		ORDER BY id DESC`

	var rows *sql.Rows

	if limit > 0 {
		query += " LIMIT ?"
		rows, err = db.conn.Query(query, collectionID, limit)
	} else {
		rows, err = db.conn.Query(query, collectionID)
	}

	if err != nil {
		return nil, fmt.Errorf("listing documents: %w", err)
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var doc Document
		if err := rows.Scan(&doc.ID, &doc.CollectionID, &doc.Content, &doc.Metadata, &doc.CreatedAt); err != nil {
			return nil, fmt.Errorf("scanning document row: %w", err)
		}
		docs = append(docs, doc)
	}

	return docs, rows.Err()
}

// DeleteDocument removes a document by its ID.
// Returns an error if the document does not exist.
func (db *DB) DeleteDocument(id int64) error {
	result, err := db.conn.Exec("DELETE FROM documents WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("deleting document %d: %w", id, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("document %d not found", id)
	}

	return nil
}

// CountDocuments returns the number of documents in a collection.
func (db *DB) CountDocuments(collectionID int64) (int64, error) {
	var count int64
	err := db.conn.QueryRow(
		"SELECT COUNT(*) FROM documents WHERE collection_id = ?",
		collectionID,
	).Scan(&count)

	if err != nil {
		return 0, fmt.Errorf("counting documents: %w", err)
	}

	return count, nil
}
