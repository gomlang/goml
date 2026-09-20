package adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	"modernc.org/sqlite"
)

var ErrClosed = errors.New("sqlite resource is closed")
var ErrBusy = errors.New("sqlite connection has an active cursor or nested transaction")
var ErrArgument = errors.New("invalid sqlite argument")
var ErrType = errors.New("unsupported sqlite value")

type Control interface{ control() }

func (*ControlState) control() {}

type ControlState struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewControl(milliseconds int64) (Control, error) {
	if milliseconds < 0 || milliseconds > math.MaxInt64/int64(time.Millisecond) {
		return nil, ErrArgument
	}
	if milliseconds == 0 {
		ctx, cancel := context.WithCancel(context.Background())
		return &ControlState{ctx, cancel}, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(milliseconds)*time.Millisecond)
	return &ControlState{ctx, cancel}, nil
}
func Cancel(cHandle Control) {
	c, _ := cHandle.(*ControlState)

	if c != nil {
		c.cancel()
	}
}
func operation(d *database, cHandle Control) (context.Context, func(), error) {
	c, _ := cHandle.(*ControlState)

	if c == nil {
		return nil, nil, ErrArgument
	}
	if err := c.ctx.Err(); err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithCancel(d.ctx)
	stop := context.AfterFunc(c.ctx, cancel)
	if c.ctx.Err() != nil {
		cancel()
	}
	return ctx, func() { stop(); cancel() }, nil
}
func cause(err error, cHandle Control) error {
	c, _ := cHandle.(*ControlState)

	if err != nil && c != nil && c.ctx.Err() != nil {
		return fmt.Errorf("%w: %v", c.ctx.Err(), err)
	}
	return err
}

type Session interface{ session() }

func (*SessionState) session() {}

type SessionState struct {
	db          *database
	name        string
	probe       string
	transaction bool
	done        bool
}
type database struct {
	mu         gate
	pool       *sql.DB
	conn       *sql.Conn
	ctx        context.Context
	cancel     context.CancelFunc
	closed     bool
	stack      []*SessionState
	statements map[*StatementState]bool
	cursors    map[*CursorState]bool
	sequence   uint64
}

func Open(dsn string, busyMilliseconds int64, foreignKeys bool, cHandle Control) (Session, error) {
	c, _ := cHandle.(*ControlState)

	if dsn == "" || strings.ContainsRune(dsn, 0) || !utf8.ValidString(dsn) || busyMilliseconds < 0 || busyMilliseconds > 2147483647 || c == nil {
		return nil, ErrArgument
	}
	separator := "?"
	if strings.Contains(dsn, "?") {
		separator = "&"
	}
	fk := "OFF"
	if foreignKeys {
		fk = "ON"
	}
	dsn += separator + "_pragma=" + url.QueryEscape(fmt.Sprintf("busy_timeout(%d)", busyMilliseconds)) + "&_pragma=" + url.QueryEscape("foreign_keys("+fk+")")
	pool, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	pool.SetMaxOpenConns(1)
	pool.SetMaxIdleConns(1)
	root, cancel := context.WithCancel(context.Background())
	d := &database{mu: newGate(), pool: pool, ctx: root, cancel: cancel, statements: make(map[*StatementState]bool), cursors: make(map[*CursorState]bool)}
	ctx, finish, err := operation(d, c)
	if err != nil {
		cancel()
		pool.Close()
		return nil, err
	}
	defer finish()
	conn, err := pool.Conn(ctx)
	if err == nil {
		err = conn.PingContext(ctx)
	}
	if err != nil {
		if conn != nil {
			conn.Close()
		}
		cancel()
		pool.Close()
		return nil, cause(err, c)
	}
	d.conn = conn
	return &SessionState{db: d}, nil
}
func (s *SessionState) check(exclusive bool) error {
	d := s.db
	if s.done || d.closed {
		return ErrClosed
	}
	if len(d.stack) != 0 && d.stack[len(d.stack)-1] != s {
		return ErrBusy
	}
	if s.transaction && len(d.stack) == 0 {
		return ErrClosed
	}
	if exclusive && len(d.cursors) != 0 {
		return ErrBusy
	}
	return nil
}
func IsClosed(sHandle Session) bool {
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return true
	}
	s.db.mu.Lock()
	defer s.db.mu.Unlock()
	return s.done || s.db.closed
}
func (d *database) closeResources(owner *SessionState, descendants bool) error {
	owns := func(s *SessionState) bool {
		if owner == nil || s == owner {
			return true
		}
		if !descendants {
			return false
		}
		found := false
		for _, item := range d.stack {
			if item == owner {
				found = true
			}
			if found && item == s {
				return true
			}
		}
		return false
	}
	var result error
	for c := range d.cursors {
		if owns(c.owner) {
			result = errors.Join(result, c.close())
		}
	}
	for st := range d.statements {
		if owns(st.owner) {
			result = errors.Join(result, st.close())
		}
	}
	return result
}
func Close(sHandle Session) error {
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return nil
	}
	if s.transaction {
		return Rollback(s)
	}
	d := s.db
	d.cancel()
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.closed {
		return nil
	}
	d.closed = true
	s.done = true
	result := d.closeResources(nil, true)
	if len(d.stack) != 0 {
		_, err := d.conn.ExecContext(context.Background(), "ROLLBACK")
		result = errors.Join(result, ignoreInactive(err))
	}
	for _, item := range d.stack {
		item.done = true
	}
	d.stack = nil
	return errors.Join(result, d.conn.Close(), d.pool.Close())
}
func ignoreInactive(err error) error {

	if err != nil && strings.Contains(err.Error(), "no transaction is active") {
		return nil
	}
	return err
}

type Value struct {
	kind    int
	integer int64
	real    float64
	text    string
	blob    []byte
}

func Null() Value {
	return Value{}
}
func Integer(value int64) Value {
	return Value{kind: 1, integer: value}
}
func Real(value float64) Value {
	return Value{kind: 2, real: value}
}
func Text(value string) Value {
	return Value{kind: 3, text: value}
}
func Blob(value []byte) Value {

	data := make([]byte, len(value))
	copy(data, value)
	return Value{kind: 4, blob: data}
}
func ValueKind(value Value) int {
	return value.kind
}
func ValueInteger(value Value) int64 {
	return value.integer
}
func ValueReal(value Value) float64 {
	return value.real
}
func ValueText(value Value) string {
	return value.text
}
func ValueBlob(value Value) []byte {

	result := make([]byte, len(value.blob))
	copy(result, value.blob)
	return result
}
func (v Value) native() (any, error) {
	switch v.kind {
	case 0:
		return nil, nil
	case 1:
		return v.integer, nil
	case 2:
		return v.real, nil
	case 3:
		return v.text, nil
	case 4:
		result := make([]byte, len(v.blob))
		copy(result, v.blob)
		return result, nil
	default:
		return nil, ErrType
	}
}
func fromNative(value any) (Value, error) {
	switch value := value.(type) {
	case nil:
		return Null(), nil
	case int64:
		return Integer(value), nil
	case float64:
		return Real(value), nil
	case bool:
		if value {
			return Integer(1), nil
		}
		return Integer(0), nil
	case string:
		return Text(value), nil
	case []byte:
		return Blob(value), nil
	case time.Time:
		return Text(value.Format(time.RFC3339Nano)), nil
	default:
		return Value{}, fmt.Errorf("%w: %T", ErrType, value)
	}
}
func parameters(values []Value, names []string) ([]any, error) {
	if len(names) != 0 && len(names) != len(values) {
		return nil, ErrArgument
	}
	if len(values) > 32766 {
		return nil, ErrArgument
	}
	result := make([]any, len(values))
	seen := make(map[string]bool)
	for i, value := range values {
		native, err := value.native()
		if err != nil {
			return nil, err
		}
		name := ""
		if len(names) != 0 {
			name = names[i]
		}
		if name != "" {
			if seen[name] {
				return nil, fmt.Errorf("%w: duplicate parameter %s", ErrArgument, name)
			}
			seen[name] = true
			for j, ch := range name {
				if !(ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || j > 0 && (ch >= '0' && ch <= '9' || ch == '_')) {
					return nil, fmt.Errorf("%w: parameter name", ErrArgument)
				}
			}
			result[i] = sql.Named(name, native)
		} else {
			result[i] = native
		}
	}
	return result, nil
}
func sqlText(text string) error {

	if len(text) == 0 || len(text) > 1048576 || strings.ContainsRune(text, 0) || !utf8.ValidString(text) {
		return fmt.Errorf("%w: SQL text", ErrArgument)
	}
	return singleStatement(text)
}
func Execute(sHandle Session, text string, values []Value, names []string, cHandle Control) (int64, int64, error) {
	c, _ := cHandle.(*ControlState)
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return 0, 0, ErrClosed
	}
	if err := sqlText(text); err != nil {
		return 0, 0, err
	}
	params, err := parameters(values, names)
	if err != nil {
		return 0, 0, err
	}
	d := s.db
	if err := d.mu.lock(c, d.ctx); err != nil {
		return 0, 0, err
	}
	defer d.mu.Unlock()
	if err := s.check(true); err != nil {
		return 0, 0, err
	}
	ctx, finish, err := operation(d, c)
	if err != nil {
		return 0, 0, err
	}
	defer finish()
	result, err := d.conn.ExecContext(ctx, text, params...)
	if err != nil {
		d.recoverTransaction()
		return 0, 0, cause(err, c)
	}
	changes, err := result.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	id, err := result.LastInsertId()
	return changes, id, err
}

type Statement interface{ statement() }

func (*StatementState) statement() {}

type StatementState struct {
	owner  *SessionState
	stmt   *sql.Stmt
	closed bool
}

func Prepare(sHandle Session, text string, cHandle Control) (Statement, error) {
	c, _ := cHandle.(*ControlState)
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return nil, ErrClosed
	}
	if err := sqlText(text); err != nil {
		return nil, err
	}
	d := s.db
	if err := d.mu.lock(c, d.ctx); err != nil {
		return nil, err
	}
	defer d.mu.Unlock()
	if err := s.check(true); err != nil {
		return nil, err
	}
	ctx, finish, err := operation(d, c)
	if err != nil {
		return nil, err
	}
	defer finish()
	st, err := d.conn.PrepareContext(ctx, text)
	if err != nil {
		d.recoverTransaction()
		return nil, cause(err, c)
	}
	result := &StatementState{owner: s, stmt: st}
	d.statements[result] = true
	return result, nil
}
func (s *StatementState) close() error {
	if s.closed {
		return nil
	}
	s.closed = true
	delete(s.owner.db.statements, s)
	var result error
	for c := range s.owner.db.cursors {
		if c.statement == s {
			result = errors.Join(result, c.close())
		}
	}
	return errors.Join(result, s.stmt.Close())
}
func CloseStatement(sHandle Statement) error {
	s, _ := sHandle.(*StatementState)

	if s == nil {
		return nil
	}
	s.owner.db.mu.Lock()
	defer s.owner.db.mu.Unlock()
	return s.close()
}
func ExecuteStatement(sHandle Statement, values []Value, names []string, cHandle Control) (int64, int64, error) {
	c, _ := cHandle.(*ControlState)
	s, _ := sHandle.(*StatementState)

	if s == nil {
		return 0, 0, ErrClosed
	}
	params, err := parameters(values, names)
	if err != nil {
		return 0, 0, err
	}
	d := s.owner.db
	if err := d.mu.lock(c, d.ctx); err != nil {
		return 0, 0, err
	}
	defer d.mu.Unlock()
	if s.closed {
		return 0, 0, ErrClosed
	}
	if err := s.owner.check(true); err != nil {
		return 0, 0, err
	}
	ctx, finish, err := operation(d, c)
	if err != nil {
		return 0, 0, err
	}
	defer finish()
	result, err := s.stmt.ExecContext(ctx, params...)
	if err != nil {
		d.recoverTransaction()
		return 0, 0, cause(err, c)
	}
	changes, err := result.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	id, err := result.LastInsertId()
	return changes, id, err
}

type Column struct {
	name     string
	declared string
}

func ColumnName(c Column) string {
	return c.name
}
func ColumnType(c Column) string {
	return c.declared
}

type Row struct {
	values  []Value
	columns []Column
}

func RowValues(r Row) []Value {
	return append([]Value(nil), r.values...)
}
func RowColumns(r Row) []Column {
	return append([]Column(nil), r.columns...)
}

type Cursor interface{ cursor() }

func (*CursorState) cursor() {}

type CursorState struct {
	owner     *SessionState
	statement *StatementState
	rows      *sql.Rows
	columns   []Column
	closed    bool
	exhausted bool
	finish    func()
	control   Control
	failure   error
}

func columns(rows *sql.Rows) ([]Column, error) {
	info, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	result := make([]Column, len(info))
	for i, column := range info {
		result[i] = Column{column.Name(), column.DatabaseTypeName()}
	}
	return result, nil
}
func newCursor(owner *SessionState, statement *StatementState, rows *sql.Rows, finish func(), cHandle Control) (Cursor, error) {
	c, _ := cHandle.(*ControlState)

	cols, err := columns(rows)
	if err != nil {
		rows.Close()
		finish()
		return nil, err
	}
	cursor := &CursorState{owner: owner, statement: statement, rows: rows, columns: cols, finish: finish, control: c}
	owner.db.cursors[cursor] = true
	return cursor, nil
}
func Query(sHandle Session, text string, values []Value, names []string, cHandle Control) (Cursor, error) {
	c, _ := cHandle.(*ControlState)
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return nil, ErrClosed
	}
	if err := sqlText(text); err != nil {
		return nil, err
	}
	params, err := parameters(values, names)
	if err != nil {
		return nil, err
	}
	d := s.db
	if err := d.mu.lock(c, d.ctx); err != nil {
		return nil, err
	}
	defer d.mu.Unlock()
	if err := s.check(true); err != nil {
		return nil, err
	}
	ctx, finish, err := operation(d, c)
	if err != nil {
		return nil, err
	}
	rows, err := d.conn.QueryContext(ctx, text, params...)
	if err != nil {
		finish()
		d.recoverTransaction()
		return nil, cause(err, c)
	}
	return newCursor(s, nil, rows, finish, c)
}
func QueryStatement(sHandle Statement, values []Value, names []string, cHandle Control) (Cursor, error) {
	c, _ := cHandle.(*ControlState)
	s, _ := sHandle.(*StatementState)

	if s == nil {
		return nil, ErrClosed
	}
	params, err := parameters(values, names)
	if err != nil {
		return nil, err
	}
	d := s.owner.db
	if err := d.mu.lock(c, d.ctx); err != nil {
		return nil, err
	}
	defer d.mu.Unlock()
	if s.closed {
		return nil, ErrClosed
	}
	if err := s.owner.check(true); err != nil {
		return nil, err
	}
	ctx, finish, err := operation(d, c)
	if err != nil {
		return nil, err
	}
	rows, err := s.stmt.QueryContext(ctx, params...)
	if err != nil {
		finish()
		d.recoverTransaction()
		return nil, cause(err, c)
	}
	return newCursor(s.owner, s, rows, finish, c)
}
func (c *CursorState) close() error {
	if c.closed {
		return nil
	}
	c.closed = true
	delete(c.owner.db.cursors, c)
	err := c.rows.Close()
	c.finish()
	return err
}
func CloseCursor(cHandle Cursor) error {
	c, _ := cHandle.(*CursorState)

	if c == nil {
		return nil
	}
	c.owner.db.mu.Lock()
	defer c.owner.db.mu.Unlock()
	return c.close()
}
func CursorColumns(cHandle Cursor) []Column {
	c, _ := cHandle.(*CursorState)

	if c == nil {
		return nil
	}
	return append([]Column(nil), c.columns...)
}
func Next(cHandle Cursor) (bool, Row, error) {
	c, _ := cHandle.(*CursorState)

	if c == nil {
		return false, Row{}, ErrClosed
	}
	d := c.owner.db
	d.mu.Lock()
	defer d.mu.Unlock()
	if c.failure != nil {
		return false, Row{}, c.failure
	}
	if c.exhausted {
		return false, Row{}, nil
	}
	if c.closed || d.closed || c.owner.done {
		return false, Row{}, ErrClosed
	}
	fail := func(err error) (bool, Row, error) {
		c.failure = cause(err, c.control)
		c.close()
		d.recoverTransaction()
		return false, Row{}, c.failure
	}
	if !c.rows.Next() {
		if err := c.rows.Err(); err != nil {
			return fail(err)
		}
		if c.rows.NextResultSet() {
			return fail(fmt.Errorf("%w: multiple query result sets", ErrArgument))
		}
		if err := c.rows.Err(); err != nil {
			return fail(err)
		}
		c.exhausted = true
		return false, Row{}, c.close()
	}
	values := make([]any, len(c.columns))
	destinations := make([]any, len(values))
	for i := range values {
		destinations[i] = &values[i]
	}
	if err := c.rows.Scan(destinations...); err != nil {
		return fail(err)
	}
	result := Row{columns: append([]Column(nil), c.columns...), values: make([]Value, len(values))}
	for i, value := range values {
		converted, err := fromNative(value)
		if err != nil {
			return fail(err)
		}
		result.values[i] = converted
	}
	return true, result, nil
}
func Begin(sHandle Session, mode int, cHandle Control) (Session, error) {
	c, _ := cHandle.(*ControlState)
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return nil, ErrClosed
	}
	if mode < 0 || mode > 2 || s.transaction {
		return nil, ErrArgument
	}
	d := s.db
	if err := d.mu.lock(c, d.ctx); err != nil {
		return nil, err
	}
	defer d.mu.Unlock()
	if err := s.check(true); err != nil {
		return nil, err
	}
	ctx, finish, err := operation(d, c)
	if err != nil {
		return nil, err
	}
	defer finish()
	_, err = d.conn.ExecContext(ctx, "BEGIN "+[]string{"DEFERRED", "IMMEDIATE", "EXCLUSIVE"}[mode])
	if err != nil {
		d.conn.ExecContext(context.Background(), "ROLLBACK")
		return nil, cause(err, c)
	}
	d.sequence++
	probe := fmt.Sprintf("goml_probe_%d", d.sequence)
	if _, err := d.conn.ExecContext(ctx, "SAVEPOINT "+probe); err != nil {
		d.conn.ExecContext(context.Background(), "ROLLBACK")
		return nil, cause(err, c)
	}
	result := &SessionState{db: d, transaction: true, probe: probe}
	d.stack = append(d.stack, result)
	return result, nil
}
func Savepoint(sHandle Session, cHandle Control) (Session, error) {
	c, _ := cHandle.(*ControlState)
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return nil, ErrClosed
	}
	if !s.transaction {
		return nil, ErrArgument
	}
	d := s.db
	if err := d.mu.lock(c, d.ctx); err != nil {
		return nil, err
	}
	defer d.mu.Unlock()
	if err := s.check(true); err != nil {
		return nil, err
	}
	ctx, finish, err := operation(d, c)
	if err != nil {
		return nil, err
	}
	defer finish()
	d.sequence++
	name := fmt.Sprintf("goml_savepoint_%d", d.sequence)
	_, err = d.conn.ExecContext(ctx, "SAVEPOINT "+name)
	if err != nil {
		d.recoverTransaction()
		return nil, cause(err, c)
	}
	probe := name + "_probe"
	if _, err := d.conn.ExecContext(ctx, "SAVEPOINT "+probe); err != nil {
		d.conn.ExecContext(context.Background(), "ROLLBACK TO SAVEPOINT "+name)
		d.conn.ExecContext(context.Background(), "RELEASE SAVEPOINT "+name)
		d.recoverTransaction()
		return nil, cause(err, c)
	}
	result := &SessionState{db: d, transaction: true, name: name, probe: probe}
	d.stack = append(d.stack, result)
	return result, nil
}
func Commit(sHandle Session, cHandle Control) error {
	c, _ := cHandle.(*ControlState)
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return ErrClosed
	}
	if !s.transaction {
		return ErrArgument
	}
	d := s.db
	if err := d.mu.lock(c, d.ctx); err != nil {
		return err
	}
	defer d.mu.Unlock()
	if err := s.check(false); err != nil {
		return err
	}
	ctx, finish, err := operation(d, c)
	if err != nil {
		return err
	}
	defer finish()
	if err := d.closeResources(s, false); err != nil {
		return err
	}
	text := "COMMIT"
	if s.name != "" {
		text = "RELEASE SAVEPOINT " + s.name
	}
	_, err = d.conn.ExecContext(ctx, text)
	if err != nil {
		d.recoverTransaction()
		return cause(err, c)
	}
	s.done = true
	d.stack = d.stack[:len(d.stack)-1]
	return nil
}
func Rollback(sHandle Session) error {
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return nil
	}
	if !s.transaction {
		return ErrArgument
	}
	d := s.db
	d.mu.Lock()
	defer d.mu.Unlock()
	if s.done || d.closed {
		return nil
	}
	index := -1
	for i, item := range d.stack {
		if item == s {
			index = i
			break
		}
	}
	if index < 0 {
		return ErrClosed
	}
	result := d.closeResources(s, true)
	if s.name == "" {
		_, err := d.conn.ExecContext(context.Background(), "ROLLBACK")
		result = errors.Join(result, ignoreInactive(err))
	} else {
		_, err := d.conn.ExecContext(context.Background(), "ROLLBACK TO SAVEPOINT "+s.name)
		result = errors.Join(result, err)
		_, err = d.conn.ExecContext(context.Background(), "RELEASE SAVEPOINT "+s.name)
		result = errors.Join(result, err)
	}
	for _, item := range d.stack[index:] {
		item.done = true
	}
	d.stack = d.stack[:index]
	if result != nil {
		d.recoverTransaction()
	}
	return result
}
func ErrorClass(err error) int {

	switch {
	case err == nil:
		return 0
	case errors.Is(err, ErrClosed):
		return 1
	case errors.Is(err, ErrBusy):
		return 2
	case errors.Is(err, ErrArgument):
		return 3
	case errors.Is(err, ErrType):
		return 4
	case errors.Is(err, context.DeadlineExceeded):
		return 5
	case errors.Is(err, context.Canceled):
		return 6
	default:
		return 7
	}
}
func ErrorCode(err error) int {

	var native *sqlite.Error
	if errors.As(err, &native) {
		return native.Code()
	}
	return 0
}
func Resources(sHandle Session) (int, int, int) {
	s, _ := sHandle.(*SessionState)

	if s == nil {
		return 0, 0, 0
	}
	d := s.db
	d.mu.Lock()
	defer d.mu.Unlock()
	return len(d.statements), len(d.cursors), len(d.stack)
}

func (d *database) recoverTransaction() {
	if len(d.stack) == 0 {
		return
	}
	current := d.stack[len(d.stack)-1]
	if _, err := d.conn.ExecContext(context.Background(), "RELEASE SAVEPOINT "+current.probe); err == nil {
		if _, err := d.conn.ExecContext(context.Background(), "SAVEPOINT "+current.probe); err == nil {
			return
		}
	}
	d.closeResources(d.stack[0], true)
	_, rollback := d.conn.ExecContext(context.Background(), "ROLLBACK")
	for _, session := range d.stack {
		session.done = true
	}
	d.stack = nil
	if ignoreInactive(rollback) != nil {
		d.closed = true
		d.cancel()
		d.closeResources(nil, true)
		d.conn.Close()
		d.pool.Close()
	}
}
