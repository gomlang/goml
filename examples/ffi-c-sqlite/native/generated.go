package native

/*
#cgo CFLAGS: -std=c11 -I${SRCDIR}/..
#cgo LDFLAGS: -lsqlite3
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <sqlite3.h>
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int (*goml_c_signature_open)(const char *, sqlite3 **);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sqlite3_open), goml_c_signature_open), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int (*goml_c_signature_close)(sqlite3 *);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sqlite3_close), goml_c_signature_close), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int (*goml_c_signature_prepare)(sqlite3 *, const char *, int, sqlite3_stmt **, const char **);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sqlite3_prepare_v2), goml_c_signature_prepare), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int (*goml_c_signature_step)(sqlite3_stmt *);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sqlite3_step), goml_c_signature_step), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(sqlite3_int64) == 8 && (((sqlite3_int64)-1 < (sqlite3_int64)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef sqlite3_int64 (*goml_c_signature_column_int64)(sqlite3_stmt *, int);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sqlite3_column_int64), goml_c_signature_column_int64), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int (*goml_c_signature_finalize)(sqlite3_stmt *);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sqlite3_finalize), goml_c_signature_finalize), "C signature changed; regenerate GoML bindings");
typedef const char * (*goml_c_signature_error_message)(sqlite3 *);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sqlite3_errmsg), goml_c_signature_error_message), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(const int) == 4 && (((const int)-1 < (const int)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert((SQLITE_OK) == (const int)0ULL, "C constant changed; regenerate GoML bindings");
_Static_assert(sizeof(const int) == 4 && (((const int)-1 < (const int)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert((SQLITE_ROW) == (const int)100ULL, "C constant changed; regenerate GoML bindings");
static size_t goml_c_bounded_length(const char *p, size_t limit) { size_t n = 0; while (n < limit && p[n] != 0) { n++; } return n; }
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

type Database struct{ raw *C.sqlite3 }

func GomlC_NullDatabase() Database                  { return Database{} }
func GomlC_IsNullDatabase(value Database) bool      { return value.raw == nil }
func GomlC_EqualDatabase(left, right Database) bool { return left.raw == right.raw }

type Statement struct{ raw *C.sqlite3_stmt }

func GomlC_NullStatement() Statement                  { return Statement{} }
func GomlC_IsNullStatement(value Statement) bool      { return value.raw == nil }
func GomlC_EqualStatement(left, right Statement) bool { return left.raw == right.raw }
func GomlC_open(a0 string) (int32, Database, error) {
	if strings.IndexByte(a0, 0) >= 0 || len(a0) > 67108864 {
		return 0, Database{}, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg0 := C.CString(a0)
	defer C.free(unsafe.Pointer(cArg0))
	var cArg1 *C.sqlite3
	cResult := C.sqlite3_open(cArg0, &cArg1)
	return int32(cResult), Database{raw: cArg1}, nil
}
func GomlC_close(a0 Database) (int32, error) {
	cResult := C.sqlite3_close(a0.raw)
	return int32(cResult), nil
}
func GomlC_prepare(a0 Database, a1 string, a2 int32) (int32, Statement, string, bool, error) {
	if strings.IndexByte(a1, 0) >= 0 || len(a1) > 67108864 {
		return 0, Statement{}, "", false, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg1 := C.CString(a1)
	defer C.free(unsafe.Pointer(cArg1))
	var cArg3 *C.sqlite3_stmt
	var cArg4 *C.char
	cResult := C.sqlite3_prepare_v2(a0.raw, cArg1, C.int(a2), &cArg3, &cArg4)
	var cArg4Text string
	if cArg4 != nil {
		n := C.goml_c_bounded_length(cArg4, 1048577)
		if n > 1048576 {
			return 0, Statement{}, "", false, fmt.Errorf("C string exceeds configured copy limit")
		}
		cArg4Text = C.GoStringN(cArg4, C.int(n))
	}
	return int32(cResult), Statement{raw: cArg3}, cArg4Text, cArg4 != nil, nil
}
func GomlC_step(a0 Statement) (int32, error) {
	cResult := C.sqlite3_step(a0.raw)
	return int32(cResult), nil
}
func GomlC_column_int64(a0 Statement, a1 int32) (int64, error) {
	cResult := C.sqlite3_column_int64(a0.raw, C.int(a1))
	return int64(cResult), nil
}
func GomlC_finalize(a0 Statement) (int32, error) {
	cResult := C.sqlite3_finalize(a0.raw)
	return int32(cResult), nil
}
func GomlC_error_message(a0 Database) (string, bool, error) {
	cResult := C.sqlite3_errmsg(a0.raw)
	var cResultText string
	if cResult != nil {
		n := C.goml_c_bounded_length(cResult, 1048577)
		if n > 1048576 {
			return "", false, fmt.Errorf("C string exceeds configured copy limit")
		}
		cResultText = C.GoStringN(cResult, C.int(n))
	}
	return cResultText, cResult != nil, nil
}
func GomlC_OK() int32  { return int32(C.SQLITE_OK) }
func GomlC_ROW() int32 { return int32(C.SQLITE_ROW) }
