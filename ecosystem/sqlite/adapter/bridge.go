package adapter

type Failure interface{ failure() }
type FailureState struct{ cause error }

func (*FailureState) failure() {}
func failure(err error) Failure {
	if err == nil {
		return nil
	}
	return &FailureState{cause: err}
}
func FailureIsNil(value Failure) bool { return value == nil }
func FailureError(value Failure) error {
	state, ok := value.(*FailureState)
	if !ok || state == nil {
		return nil
	}
	return state.cause
}
func FailureClass(value Failure) int { return ErrorClass(FailureError(value)) }
func FailureCode(value Failure) int  { return ErrorCode(FailureError(value)) }
func FailureMessage(value Failure) string {
	err := FailureError(value)
	if err == nil {
		return ""
	}
	return err.Error()
}

func BridgeNewControl(milliseconds int64) (Control, Failure) {
	value0, err := NewControl(milliseconds)
	return value0, failure(err)
}

func BridgeOpen(dsn string, busyMilliseconds int64, foreignKeys bool, cHandle Control) (Session, Failure) {
	value0, err := Open(dsn, busyMilliseconds, foreignKeys, cHandle)
	return value0, failure(err)
}

func BridgeClose(sHandle Session) Failure {
	err := Close(sHandle)
	return failure(err)
}

func BridgeExecute(sHandle Session, text string, values []Value, names []string, cHandle Control) (int64, int64, Failure) {
	value0, value1, err := Execute(sHandle, text, values, names, cHandle)
	return value0, value1, failure(err)
}

func BridgePrepare(sHandle Session, text string, cHandle Control) (Statement, Failure) {
	value0, err := Prepare(sHandle, text, cHandle)
	return value0, failure(err)
}

func BridgeCloseStatement(sHandle Statement) Failure {
	err := CloseStatement(sHandle)
	return failure(err)
}

func BridgeExecuteStatement(sHandle Statement, values []Value, names []string, cHandle Control) (int64, int64, Failure) {
	value0, value1, err := ExecuteStatement(sHandle, values, names, cHandle)
	return value0, value1, failure(err)
}

func BridgeQuery(sHandle Session, text string, values []Value, names []string, cHandle Control) (Cursor, Failure) {
	value0, err := Query(sHandle, text, values, names, cHandle)
	return value0, failure(err)
}

func BridgeQueryStatement(sHandle Statement, values []Value, names []string, cHandle Control) (Cursor, Failure) {
	value0, err := QueryStatement(sHandle, values, names, cHandle)
	return value0, failure(err)
}

func BridgeCloseCursor(cHandle Cursor) Failure {
	err := CloseCursor(cHandle)
	return failure(err)
}

func BridgeNext(cHandle Cursor) (bool, Row, Failure) {
	value0, value1, err := Next(cHandle)
	return value0, value1, failure(err)
}

func BridgeBegin(sHandle Session, mode int, cHandle Control) (Session, Failure) {
	value0, err := Begin(sHandle, mode, cHandle)
	return value0, failure(err)
}

func BridgeSavepoint(sHandle Session, cHandle Control) (Session, Failure) {
	value0, err := Savepoint(sHandle, cHandle)
	return value0, failure(err)
}

func BridgeCommit(sHandle Session, cHandle Control) Failure {
	err := Commit(sHandle, cHandle)
	return failure(err)
}

func BridgeRollback(sHandle Session) Failure {
	err := Rollback(sHandle)
	return failure(err)
}
