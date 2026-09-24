package native

import (
	"fmt"
	"goml.dev/cabi"
	"strings"
	"sync"
	"unsafe"
)

var gomlCLibrary = cabi.Library{Names: []string{"libsqlite3.so.0"}}
var gomlCOnce sync.Once
var gomlCError error
var gomlCAddresses [7]uint64

func gomlCLoad() error {
	gomlCOnce.Do(func() {
		gomlCAddresses[0], gomlCError = gomlCLibrary.Symbol("sqlite3_close")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[1], gomlCError = gomlCLibrary.Symbol("sqlite3_column_int64")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[2], gomlCError = gomlCLibrary.Symbol("sqlite3_errmsg")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[3], gomlCError = gomlCLibrary.Symbol("sqlite3_finalize")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[4], gomlCError = gomlCLibrary.Symbol("sqlite3_open")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[5], gomlCError = gomlCLibrary.Symbol("sqlite3_prepare_v2")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[6], gomlCError = gomlCLibrary.Symbol("sqlite3_step")
		if gomlCError != nil {
			return
		}
	})
	return gomlCError
}
func gomlDynamic_sqlite3_close(p0 unsafe.Pointer) int32 {
	f := cabi.Frame{Function: gomlCAddresses[0], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
	return int32(f.Result)
}
func gomlDynamic_sqlite3_column_int64(p0 unsafe.Pointer, p1 int32) int64 {
	f := cabi.Frame{Function: gomlCAddresses[1], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(p1)
	f.Call()
	return int64(f.Result)
}
func gomlDynamic_sqlite3_errmsg(p0 unsafe.Pointer) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[2], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_sqlite3_finalize(p0 unsafe.Pointer) int32 {
	f := cabi.Frame{Function: gomlCAddresses[3], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
	return int32(f.Result)
}
func gomlDynamic_sqlite3_open(p0 unsafe.Pointer, p1 unsafe.Pointer) int32 {
	f := cabi.Frame{Function: gomlCAddresses[4], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(uintptr(p1))
	f.Call()
	return int32(f.Result)
}
func gomlDynamic_sqlite3_prepare_v2(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 int32, p3 unsafe.Pointer, p4 unsafe.Pointer) int32 {
	f := cabi.Frame{Function: gomlCAddresses[5], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(uintptr(p1))
	f.Integer[2] = uint64(p2)
	f.Integer[3] = uint64(uintptr(p3))
	f.Integer[4] = uint64(uintptr(p4))
	f.Call()
	return int32(f.Result)
}
func gomlDynamic_sqlite3_step(p0 unsafe.Pointer) int32 {
	f := cabi.Frame{Function: gomlCAddresses[6], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
	return int32(f.Result)
}

type Database struct{ raw unsafe.Pointer }

func GomlC_NullDatabase() Database                  { return Database{} }
func GomlC_IsNullDatabase(value Database) bool      { return value.raw == nil }
func GomlC_EqualDatabase(left, right Database) bool { return left.raw == right.raw }

type Statement struct{ raw unsafe.Pointer }

func GomlC_NullStatement() Statement                  { return Statement{} }
func GomlC_IsNullStatement(value Statement) bool      { return value.raw == nil }
func GomlC_EqualStatement(left, right Statement) bool { return left.raw == right.raw }
func GomlC_open(a0 string) (int32, Database, error) {
	if err := gomlCLoad(); err != nil {
		return 0, Database{}, err
	}
	if strings.IndexByte(a0, 0) >= 0 || len(a0) > 67108864 {
		return 0, Database{}, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg0, _ := cabi.Copy([]byte(a0), true)
	if cArg0 == nil {
		return 0, Database{}, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg0)
	cArg1Slot, _ := cabi.Alloc(8)
	if cArg1Slot == nil {
		return 0, Database{}, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg1Slot)
	cResult := gomlDynamic_sqlite3_open(cArg0, cArg1Slot)
	cArg1 := *(*unsafe.Pointer)(cArg1Slot)
	return int32(cResult), Database{raw: cArg1}, nil
}
func GomlC_close(a0 Database) (int32, error) {
	if err := gomlCLoad(); err != nil {
		return 0, err
	}
	cResult := gomlDynamic_sqlite3_close(a0.raw)
	return int32(cResult), nil
}
func GomlC_prepare(a0 Database, a1 string, a2 int32) (int32, Statement, string, bool, error) {
	if err := gomlCLoad(); err != nil {
		return 0, Statement{}, "", false, err
	}
	if strings.IndexByte(a1, 0) >= 0 || len(a1) > 67108864 {
		return 0, Statement{}, "", false, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg1, _ := cabi.Copy([]byte(a1), true)
	if cArg1 == nil {
		return 0, Statement{}, "", false, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg1)
	cArg3Slot, _ := cabi.Alloc(8)
	if cArg3Slot == nil {
		return 0, Statement{}, "", false, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg3Slot)
	cArg4Slot, _ := cabi.Alloc(8)
	if cArg4Slot == nil {
		return 0, Statement{}, "", false, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg4Slot)
	cResult := gomlDynamic_sqlite3_prepare_v2(a0.raw, cArg1, int32(a2), cArg3Slot, cArg4Slot)
	cArg3 := *(*unsafe.Pointer)(cArg3Slot)
	cArg4 := *(*unsafe.Pointer)(cArg4Slot)
	cArg4Text, cArg4Error := cabi.String(cArg4, 1048576)
	if cArg4Error != nil {
		return 0, Statement{}, "", false, fmt.Errorf("C string exceeds configured copy limit")
	}
	return int32(cResult), Statement{raw: cArg3}, cArg4Text, cArg4 != nil, nil
}
func GomlC_step(a0 Statement) (int32, error) {
	if err := gomlCLoad(); err != nil {
		return 0, err
	}
	cResult := gomlDynamic_sqlite3_step(a0.raw)
	return int32(cResult), nil
}
func GomlC_column_int64(a0 Statement, a1 int32) (int64, error) {
	if err := gomlCLoad(); err != nil {
		return 0, err
	}
	cResult := gomlDynamic_sqlite3_column_int64(a0.raw, int32(a1))
	return int64(cResult), nil
}
func GomlC_finalize(a0 Statement) (int32, error) {
	if err := gomlCLoad(); err != nil {
		return 0, err
	}
	cResult := gomlDynamic_sqlite3_finalize(a0.raw)
	return int32(cResult), nil
}
func GomlC_error_message(a0 Database) (string, bool, error) {
	if err := gomlCLoad(); err != nil {
		return "", false, err
	}
	cResult := gomlDynamic_sqlite3_errmsg(a0.raw)
	cResultText, cResultError := cabi.String(cResult, 1048576)
	if cResultError != nil {
		return "", false, fmt.Errorf("C string exceeds configured copy limit")
	}
	return cResultText, cResult != nil, nil
}
func GomlC_OK() int32  { return 0 }
func GomlC_ROW() int32 { return 100 }
