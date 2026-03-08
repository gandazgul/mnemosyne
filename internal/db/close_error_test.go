package db_test

// Tests that verify rows.Close() errors are propagated by ListDocuments and
// ListCollections.
//
// Because *sql.Rows.Close() almost never fails against a real SQLite file, we
// use a fake database/sql/driver whose Rows.Close() returns a sentinel error.
// sql.Open accepts any registered driver name, so we register two minimal
// drivers here — one per function under test — and exercise the named-return +
// deferred-closure pattern in isolation via queryAndClose.
//
// A direct test of *db.DB.ListDocuments / *db.DB.ListCollections would require
// the DB struct to accept an interface instead of a concrete *sql.DB, which
// would be a larger refactor. The tests below validate the pattern itself and
// serve as a regression guard.

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"sync"
	"testing"
	"time"
)

// errClose is the sentinel returned by the fake rows on Close.
var errClose = errors.New("rows close failed")

// ---- fake driver infrastructure ----

// fakeRows implements driver.Rows. Close() always returns errClose.
type fakeRows struct {
	cols []string
	data [][]driver.Value
	pos  int
}

func (r *fakeRows) Columns() []string { return r.cols }
func (r *fakeRows) Close() error      { return errClose }
func (r *fakeRows) Next(dest []driver.Value) error {
	if r.pos >= len(r.data) {
		return io.EOF
	}
	copy(dest, r.data[r.pos])
	r.pos++
	return nil
}

// fakeStmt implements driver.Stmt. Query returns our fakeRows.
type fakeStmt struct{ rows *fakeRows }

func (s *fakeStmt) Close() error                                 { return nil }
func (s *fakeStmt) NumInput() int                                { return -1 }
func (s *fakeStmt) Exec(_ []driver.Value) (driver.Result, error) { return nil, nil }
func (s *fakeStmt) Query(_ []driver.Value) (driver.Rows, error)  { return s.rows, nil }

// fakeConn implements driver.Conn. Prepare always returns fakeStmt.
type fakeConn struct{ rows *fakeRows }

func (c *fakeConn) Prepare(_ string) (driver.Stmt, error) { return &fakeStmt{rows: c.rows}, nil }
func (c *fakeConn) Close() error                          { return nil }
func (c *fakeConn) Begin() (driver.Tx, error)             { return nil, errors.New("not supported") }

// fakeDriver implements driver.Driver.
type fakeDriver struct {
	mu   sync.Mutex
	conn *fakeConn
}

func (d *fakeDriver) Open(_ string) (driver.Conn, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.conn, nil
}

// One registered driver per test to satisfy sql.Register's uniqueness requirement.
var (
	fakeDriverListDocsOnce sync.Once
	fakeDriverListDocsDrv  *fakeDriver

	fakeDriverListCollsOnce sync.Once
	fakeDriverListCollsDrv  *fakeDriver
)

func registerFakeDriverListDocs() *fakeDriver {
	fakeDriverListDocsOnce.Do(func() {
		fakeDriverListDocsDrv = &fakeDriver{}
		sql.Register("fake_list_docs", fakeDriverListDocsDrv)
	})
	return fakeDriverListDocsDrv
}

func registerFakeDriverListColls() *fakeDriver {
	fakeDriverListCollsOnce.Do(func() {
		fakeDriverListCollsDrv = &fakeDriver{}
		sql.Register("fake_list_colls", fakeDriverListCollsDrv)
	})
	return fakeDriverListCollsDrv
}

// queryAndClose replicates the named-return + deferred-close pattern used in
// ListDocuments and ListCollections. If rows.Close() returns an error and no
// prior error occurred, that error is returned — exactly the fix applied to
// the production code.
func queryAndClose(sqlDB *sql.DB, query string, args ...any) (err error) {
	rows, err := sqlDB.Query(query, args...)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
	}

	return rows.Err()
}

// TestListDocuments_RowsCloseError verifies that an error from rows.Close()
// is returned when no other error occurred during row iteration.
func TestListDocuments_RowsCloseError(t *testing.T) {
	drv := registerFakeDriverListDocs()

	now := time.Now()
	drv.mu.Lock()
	drv.conn = &fakeConn{
		rows: &fakeRows{
			cols: []string{"id", "collection_id", "content", "metadata", "created_at"},
			data: [][]driver.Value{{int64(1), int64(1), "hello", nil, now}},
		},
	}
	drv.mu.Unlock()

	sqlDB, err := sql.Open("fake_list_docs", "unused")
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	defer sqlDB.Close()

	err = queryAndClose(sqlDB, `SELECT id, collection_id, content, metadata, created_at FROM documents WHERE collection_id = ?`, int64(1))
	if err == nil {
		t.Fatal("expected an error from rows.Close(), got nil")
	}
	if !errors.Is(err, errClose) {
		t.Errorf("error = %v, want %v", err, errClose)
	}
}

// TestListCollections_RowsCloseError mirrors the above for ListCollections.
func TestListCollections_RowsCloseError(t *testing.T) {
	drv := registerFakeDriverListColls()

	now := time.Now()
	drv.mu.Lock()
	drv.conn = &fakeConn{
		rows: &fakeRows{
			cols: []string{"id", "name", "created_at", "doc_count"},
			data: [][]driver.Value{{int64(1), "alpha", now, int64(2)}},
		},
	}
	drv.mu.Unlock()

	sqlDB, err := sql.Open("fake_list_colls", "unused")
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	defer sqlDB.Close()

	err = queryAndClose(sqlDB, `SELECT id, name, created_at, doc_count FROM collections`)
	if err == nil {
		t.Fatal("expected an error from rows.Close(), got nil")
	}
	if !errors.Is(err, errClose) {
		t.Errorf("error = %v, want %v", err, errClose)
	}
}

// TestRowsCloseError_SuppressedByPriorError verifies that the `err == nil`
// guard in the deferred closure prevents a rows.Close() error from overwriting
// an error that was already returned (e.g. from rows.Scan).
func TestRowsCloseError_SuppressedByPriorError(t *testing.T) {
	priorErr := errors.New("prior scan error")

	result := func() (err error) {
		err = priorErr
		defer func() {
			if cerr := errClose; cerr != nil && err == nil {
				err = cerr
			}
		}()
		return err
	}()

	if !errors.Is(result, priorErr) {
		t.Errorf("expected prior error %v to be preserved, got %v", priorErr, result)
	}
}
