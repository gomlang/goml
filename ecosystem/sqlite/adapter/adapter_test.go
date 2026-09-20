package adapter

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"
)

func testControl(t *testing.T, milliseconds int64) Control {
	t.Helper()
	c, err := NewControl(milliseconds)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { Cancel(c) })
	return c
}
func testDatabase(t *testing.T, path string, busy int64) Session {
	t.Helper()
	db, err := Open(path, busy, true, testControl(t, 0))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Close(db); err != nil {
			t.Error(err)
		}
	})
	return db
}
func TestFileLockContentionAndRecovery(t *testing.T) {
	path := filepath.Join(t.TempDir(), "locks.db")
	a, b := testDatabase(t, path, 25), testDatabase(t, path, 25)
	c := testControl(t, 0)
	if _, _, err := Execute(a, "CREATE TABLE values_table(x)", nil, nil, c); err != nil {
		t.Fatal(err)
	}
	tx, err := Begin(a, 1, c)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := Execute(tx, "INSERT INTO values_table VALUES(7)", nil, nil, c); err != nil {
		t.Fatal(err)
	}
	if _, err := Begin(b, 1, c); ErrorCode(err)&255 != 5 {
		t.Fatalf("expected SQLITE_BUSY, got %v", err)
	}
	if _, _, depth := Resources(b); depth != 0 {
		t.Fatal("failed begin leaked transaction")
	}
	if err := Rollback(tx); err != nil {
		t.Fatal(err)
	}
	if _, _, err := Execute(b, "INSERT INTO values_table VALUES(9)", nil, nil, c); err != nil {
		t.Fatal(err)
	}
	rows, err := Query(a, "SELECT x FROM values_table", nil, nil, c)
	if err != nil {
		t.Fatal(err)
	}
	defer CloseCursor(rows)
	present, row, err := Next(rows)
	if err != nil || !present || ValueInteger(RowValues(row)[0]) != 9 {
		t.Fatalf("unexpected committed state: %v %v %v", present, row, err)
	}
}
func TestDeadlineWhileWaitingForConnectionGate(t *testing.T) {
	db := testDatabase(t, ":memory:", 25)
	state := db.(*SessionState)
	state.db.mu.Lock()
	defer state.db.mu.Unlock()
	deadline := testControl(t, 20)
	finished := make(chan error, 1)
	go func() { _, _, err := Execute(db, "SELECT 1", nil, nil, deadline); finished <- err }()
	select {
	case err := <-finished:
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("expected deadline, got %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("operation did not cancel while waiting for gate")
	}
}
func TestBindingCopiesAndOpaqueFailureRetainsErrorIdentity(t *testing.T) {
	input := []byte{0, 255, 128}
	value := Blob(input)
	input[1] = 1
	output := ValueBlob(value)
	output[2] = 2
	if ValueBlob(value)[1] != 255 || ValueBlob(value)[2] != 128 {
		t.Fatal("blob aliases input or getter output")
	}
	wrapped := failure(ErrArgument)
	if FailureIsNil(wrapped) || !errors.Is(FailureError(wrapped), ErrArgument) || FailureClass(wrapped) != 3 {
		t.Fatal("opaque boundary lost error identity")
	}
	if !FailureIsNil(failure(nil)) {
		t.Fatal("nil failure was not preserved")
	}
}

func TestCommitFailureDetectsEndedNativeTransaction(t *testing.T) {
	db := testDatabase(t, ":memory:", 25)
	c := testControl(t, 0)
	if _, _, err := Execute(db, "CREATE TABLE values_table(x)", nil, nil, c); err != nil {
		t.Fatal(err)
	}
	tx, err := Begin(db, 0, c)
	if err != nil {
		t.Fatal(err)
	}
	statement, err := Prepare(tx, "INSERT INTO values_table VALUES(1)", c)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := ExecuteStatement(statement, nil, nil, c); err != nil {
		t.Fatal(err)
	}
	state := db.(*SessionState)
	state.db.mu.Lock()
	_, err = state.db.conn.ExecContext(context.Background(), "ROLLBACK")
	state.db.mu.Unlock()
	if err != nil {
		t.Fatal(err)
	}
	if err := Commit(tx, c); err == nil {
		t.Fatal("commit should report the ended transaction")
	}
	if !IsClosed(tx) {
		t.Fatal("failed commit retained an invalid transaction handle")
	}
	if _, _, err := Execute(tx, "INSERT INTO values_table VALUES(2)", nil, nil, c); !errors.Is(err, ErrClosed) {
		t.Fatalf("invalid handle could execute: %v", err)
	}
	if statements, cursors, depth := Resources(db); statements != 0 || cursors != 0 || depth != 0 {
		t.Fatal("resources survived failed commit")
	}
	if _, _, err := Execute(db, "INSERT INTO values_table VALUES(3)", nil, nil, c); err != nil {
		t.Fatal(err)
	}
}
