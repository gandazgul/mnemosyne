package db

import (
	"fmt"
)

// GetTags returns all distinct tags used in a collection.
func (db *DB) GetTags(collectionID int64) ([]string, error) {
	query := `
		SELECT DISTINCT value
		FROM documents, json_each(metadata, '$.tags')
		WHERE collection_id = ?
		ORDER BY value`

	rows, err := db.conn.Query(query, collectionID)
	if err != nil {
		return nil, fmt.Errorf("querying tags: %w", err)
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, fmt.Errorf("scanning tag: %w", err)
		}
		tags = append(tags, tag)
	}

	return tags, rows.Err()
}
