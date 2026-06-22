package db

import (
	"fmt"
	"strings"
	"time"
	"unicode"
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
	// 3. Rank by FTS5's hidden rank column. By default this is equivalent to
	//    bm25(docs_fts), where lower values are more relevant. Ordering by the
	//    hidden rank column lets FTS5 optimize ORDER BY/LIMIT internally instead
	//    of forcing SQLite to materialize and sort all matches.
	//    We negate it in the result so higher values = more relevant.
	sql := `
		SELECT d.id, d.collection_id, d.content, d.metadata, d.created_at,
		       docs_fts.rank AS rank
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

	sql += " ORDER BY rank"

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
		r.Rank = -r.Rank
		results = append(results, r)
	}

	return results, rows.Err()
}

// sanitizeFTSQuery converts a user query into a safe FTS5 MATCH expression.
// Ordinary natural-language text is treated as a retrieval query: low-signal
// stopwords are removed and remaining terms are OR'd so sparse questions do
// not become expensive all-terms-required boolean queries. Explicit quoted
// phrases are preserved as exact phrase clauses.
func sanitizeFTSQuery(query string) string {
	query = strings.TrimSpace(query)
	if query == "" {
		return ""
	}

	var clauses []string

	if strings.Count(query, `"`)%2 == 0 {
		for {
			start := strings.Index(query, `"`)
			if start < 0 {
				clauses = append(clauses, ftsTokenClauses(query)...)
				break
			}

			clauses = append(clauses, ftsTokenClauses(query[:start])...)
			query = query[start+1:]

			end := strings.Index(query, `"`)
			if end < 0 {
				// This should be unreachable because quote count is even, but keep
				// the fallback defensive.
				clauses = append(clauses, ftsTokenClauses(query)...)
				break
			}

			if phrase := sanitizeFTSPhrase(query[:end]); phrase != "" {
				clauses = append(clauses, phrase)
			}
			query = query[end+1:]
		}
	} else {
		// Unbalanced quotes cause FTS5 parse errors. Treat the quotes as noise.
		clauses = append(clauses, ftsTokenClauses(strings.ReplaceAll(query, `"`, " "))...)
	}

	if len(clauses) == 0 {
		return ""
	}
	return strings.Join(clauses, " OR ")
}

func ftsTokenClauses(query string) []string {
	var clauses []string
	var b strings.Builder
	for _, r := range query {
		switch {
		case unicode.IsLetter(r), unicode.IsDigit(r), r == '_':
			b.WriteRune(r)
		default:
			// Replace other characters with a space to avoid breaking tokenization.
			b.WriteRune(' ')
		}
	}

	for _, token := range strings.Fields(b.String()) {
		if isFTSStopword(token) || isLowSignalFTSToken(token) {
			continue
		}
		clauses = append(clauses, token)
	}

	return clauses
}

func sanitizeFTSPhrase(phrase string) string {
	var b strings.Builder
	for _, r := range phrase {
		switch {
		case unicode.IsLetter(r), unicode.IsDigit(r), r == '_':
			b.WriteRune(r)
		default:
			b.WriteRune(' ')
		}
	}

	tokens := strings.Fields(b.String())
	if len(tokens) == 0 {
		return ""
	}
	return `"` + strings.Join(tokens, " ") + `"`
}

func isLowSignalFTSToken(token string) bool {
	runes := []rune(token)
	return len(runes) == 1 && unicode.IsDigit(runes[0])
}

func isFTSStopword(token string) bool {
	_, ok := ftsStopwords[strings.ToLower(token)]
	return ok
}

var ftsStopwords = map[string]struct{}{
	"a": {}, "about": {}, "above": {}, "after": {}, "again": {}, "against": {},
	"all": {}, "am": {}, "an": {}, "and": {}, "any": {}, "are": {}, "as": {},
	"at": {}, "be": {}, "because": {}, "been": {}, "before": {}, "being": {},
	"below": {}, "between": {}, "both": {}, "but": {}, "by": {}, "can": {},
	"could": {}, "did": {}, "do": {}, "does": {}, "doing": {}, "down": {},
	"during": {}, "each": {}, "few": {}, "for": {}, "from": {}, "further": {},
	"had": {}, "has": {}, "have": {}, "having": {}, "he": {}, "her": {},
	"here": {}, "hers": {}, "herself": {}, "him": {}, "himself": {}, "his": {},
	"how": {}, "i": {}, "if": {}, "in": {}, "into": {}, "is": {}, "it": {},
	"its": {}, "itself": {}, "just": {}, "me": {}, "more": {}, "most": {},
	"my": {}, "myself": {}, "no": {}, "nor": {}, "not": {}, "now": {},
	"of": {}, "off": {}, "on": {}, "once": {}, "only": {}, "or": {},
	"other": {}, "our": {}, "ours": {}, "ourselves": {}, "out": {}, "over": {},
	"own": {}, "same": {}, "she": {}, "should": {}, "so": {}, "some": {},
	"such": {}, "than": {}, "that": {}, "the": {}, "their": {}, "theirs": {},
	"them": {}, "themselves": {}, "then": {}, "there": {}, "these": {},
	"they": {}, "this": {}, "those": {}, "through": {}, "to": {}, "too": {},
	"under": {}, "until": {}, "up": {}, "very": {}, "was": {}, "we": {},
	"were": {}, "what": {}, "when": {}, "where": {}, "which": {}, "while": {},
	"who": {}, "whom": {}, "why": {}, "will": {}, "with": {}, "you": {},
	"your": {}, "yours": {}, "yourself": {}, "yourselves": {},
}
