package db

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

// VectorResult represents a document matched by vector similarity search,
// including its cosine distance from the query vector.
type VectorResult struct {
	// Document fields (populated by joining back to the documents table).
	ID           int64
	CollectionID int64
	Content      string
	Metadata     *string
	CreatedAt    time.Time

	// Distance is the cosine distance from the query vector.
	// Lower values indicate higher similarity (0 = identical, 2 = opposite).
	Distance float64
}

// EnsureVectorTable creates the docs_vec virtual table if it doesn't already exist.
// The dimension parameter specifies the embedding vector size (e.g. 768, 512, 256).
//
// This is separate from migrate() because:
//   - The vector dimension comes from the embedding model config, not the DB layer.
//   - Commands that don't need vectors (list, init) don't need to know about dimensions.
//   - The dimension is baked into the DDL and only matters on first creation.
func (db *DB) EnsureVectorTable(dimension int) error {
	// The collection_id metadata column allows sqlite-vec to filter by collection
	// during KNN search, which is more efficient than post-filtering via a join.
	ddl := fmt.Sprintf(`CREATE VIRTUAL TABLE IF NOT EXISTS docs_vec USING vec0(
		document_id INTEGER PRIMARY KEY,
		embedding float[%d] distance_metric=cosine,
		collection_id INTEGER
	)`, dimension)

	if _, err := db.conn.Exec(ddl); err != nil {
		return fmt.Errorf("creating docs_vec table (dim=%d): %w", dimension, err)
	}

	return nil
}

// InsertVector stores an embedding vector for a document in the vector index.
// The documentID must correspond to an existing row in the documents table.
// The collectionID is stored as a metadata column for efficient KNN filtering.
func (db *DB) InsertVector(documentID, collectionID int64, embedding []float32) error {
	_, err := db.conn.Exec(
		"INSERT INTO docs_vec (document_id, embedding, collection_id) VALUES (?, ?, ?)",
		documentID, SerializeFloat32(embedding), collectionID,
	)
	if err != nil {
		return fmt.Errorf("inserting vector for document %d: %w", documentID, err)
	}

	return nil
}

// SearchVectors performs a KNN vector similarity search using sqlite-vec,
// scoped to the given collection. Returns results ordered by cosine distance
// (most similar first).
//
// The query embedding is matched against stored vectors using the MATCH operator.
// The collection_id metadata column filters results within the KNN scan itself,
// so the requested limit is always satisfied (unlike post-filtering via a join).
func (db *DB) SearchVectors(collectionID int64, queryEmbedding []float32, tags []string, limit int) (results []VectorResult, err error) {
	if limit <= 0 {
		limit = 10
	}

	// The KNN query:
	// - MATCH runs the vector similarity search
	// - collection_id = ? filters within the KNN scan
	// We join back to the documents table to get the full document content.
	query := `
		SELECT d.id, d.collection_id, d.content, d.metadata, d.created_at,
		       v.distance
		FROM docs_vec v
		JOIN documents d ON d.id = v.document_id
		WHERE v.embedding MATCH ?
		  AND v.collection_id = ?`

	// If filtering by tags, sqlite-vec evaluates k=? BEFORE filtering on d.metadata.
	// We fetch more candidates from the KNN scan when filtering to compensate.
	scanLimit := limit
	if len(tags) > 0 {
		scanLimit = limit * 10
	}
	query += " AND v.k = ?"

	args := []interface{}{SerializeFloat32(queryEmbedding), collectionID, scanLimit}

	for _, tag := range tags {
		query += " AND EXISTS (SELECT 1 FROM json_each(d.metadata, '$.tags') WHERE value = ?)"
		args = append(args, tag)
	}

	if len(tags) > 0 {
		query += " LIMIT ?"
		args = append(args, limit)
	}

	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("vector search: %w", err)
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var r VectorResult
		if err := rows.Scan(&r.ID, &r.CollectionID, &r.Content, &r.Metadata, &r.CreatedAt, &r.Distance); err != nil {
			return nil, fmt.Errorf("scanning vector result: %w", err)
		}
		results = append(results, r)
	}

	return results, rows.Err()
}

// DeleteVector removes the embedding vector for a given document ID.
// This should be called when a document is deleted, since sqlite-vec virtual
// tables do not support ON DELETE CASCADE triggers.
func (db *DB) DeleteVector(documentID int64) error {
	_, err := db.conn.Exec("DELETE FROM docs_vec WHERE document_id = ?", documentID)
	if err != nil {
		return fmt.Errorf("deleting vector for document %d: %w", documentID, err)
	}

	return nil
}

// DeleteVectorsByCollection removes all embedding vectors for documents in the
// given collection. This should be called before deleting a collection, since
// CASCADE on the documents table won't propagate to the sqlite-vec virtual table.
func (db *DB) DeleteVectorsByCollection(collectionID int64) error {
	// Delete vectors whose document_id belongs to the collection.
	_, err := db.conn.Exec(
		"DELETE FROM docs_vec WHERE collection_id = ?",
		collectionID,
	)
	if err != nil {
		return fmt.Errorf("deleting vectors for collection %d: %w", collectionID, err)
	}

	return nil
}

// SerializeFloat32 converts a []float32 slice into a little-endian byte slice
// suitable for insertion into sqlite-vec. Each float32 is encoded as 4 bytes
// in IEEE 754 format, matching the binary format sqlite-vec expects.
//
// This is the Go equivalent of Python's struct.pack("%sf" % len(vec), *vec).
func SerializeFloat32(v []float32) []byte {
	buf := make([]byte, len(v)*4)
	for i, val := range v {
		binary.LittleEndian.PutUint32(buf[i*4:], math.Float32bits(val))
	}
	return buf
}
