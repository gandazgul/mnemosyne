package db

import (
	"fmt"
	"strings"
	"time"
)

// SearchResult represents a document matched by FTS5 search,
// including its BM25 relevance score.
type SearchResult struct {
	// Document fields.
	ID           int64
	CollectionID int64
	Content      string
	Metadata     *string
	CreatedAt    time.Time

	// Rank is the BM25 relevance score from FTS5.
	// Lower (more negative) values indicate higher relevance.
	// We negate it so higher values = more relevant.
	Rank float64
}

// SearchFTS performs a full-text search using FTS5 with BM25 ranking,
// scoped to the given collection. Returns results ordered by relevance
// (highest rank first).
//
// The query string is passed directly to FTS5's MATCH operator. FTS5
// supports operators like AND, OR, NOT, and phrase queries ("exact phrase").
// Special characters that could break the query are escaped.
//
// Use limit <= 0 for no limit.
func (db *DB) SearchFTS(collectionID int64, query string, tags []string, limit int) (results []SearchResult, err error) {
	// Sanitize the query: remove characters that are FTS5 operators/syntax
	// that the user probably doesn't intend. We keep alphanumeric, spaces,
	// and double quotes (for phrase queries).
	sanitized := sanitizeFTSQuery(query)
	if sanitized == "" {
		return nil, nil
	}

	// The query joins docs_fts with documents to:
	// 1. Filter by collection_id (FTS5 table is global, not per-collection)
	// 2. Retrieve full document fields
	// 3. Rank by BM25 (bm25() returns negative values; more negative = more relevant)
	//    We negate it so higher values = more relevant, which is more intuitive.
	sql := `
		SELECT d.id, d.collection_id, d.content, d.metadata, d.created_at,
		       -bm25(docs_fts) AS rank
		FROM docs_fts
		JOIN documents d ON d.id = docs_fts.rowid
		WHERE docs_fts MATCH ?
		  AND d.collection_id = ?`

	var args []interface{}
	args = append(args, sanitized, collectionID)

	for _, tag := range tags {
		sql += " AND EXISTS (SELECT 1 FROM json_each(d.metadata, '$.tags') WHERE value = ?)"
		args = append(args, tag)
	}

	sql += " ORDER BY rank DESC"

	if limit > 0 {
		sql += " LIMIT ?"
		args = append(args, limit)
	}

	rows, err := db.conn.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("FTS search: %w", err)
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var r SearchResult
		if err := rows.Scan(&r.ID, &r.CollectionID, &r.Content, &r.Metadata, &r.CreatedAt, &r.Rank); err != nil {
			return nil, fmt.Errorf("scanning FTS result: %w", err)
		}
		results = append(results, r)
	}

	return results, rows.Err()
}

// sanitizeFTSQuery cleans a user query for safe use with FTS5 MATCH.
// It removes special FTS5 syntax characters that could cause parse errors,
// while preserving quoted phrases and basic word tokens.
func sanitizeFTSQuery(query string) string {
	query = strings.TrimSpace(query)
	if query == "" {
		return ""
	}

	// Remove characters that are FTS5 column filters or operators.
	// Keep: letters, digits, spaces, double quotes (for phrases), hyphens, underscores.
	var b strings.Builder
	for _, r := range query {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == ' ', r == '"', r == '-', r == '_':
			b.WriteRune(r)
		default:
			// Replace other characters with a space to avoid breaking tokenization.
			b.WriteRune(' ')
		}
	}

	// Collapse multiple spaces and trim.
	result := strings.Join(strings.Fields(b.String()), " ")

	// Ensure balanced double quotes (unbalanced quotes cause FTS5 parse errors).
	if strings.Count(result, `"`)%2 != 0 {
		result = strings.ReplaceAll(result, `"`, "")
	}

	return result
}
