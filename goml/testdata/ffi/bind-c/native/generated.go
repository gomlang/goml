package native

/*
#cgo CFLAGS: -std=c11 -I${SRCDIR}/..
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <sample.h>
_Static_assert(sizeof(int64_t) == 8 && (((int64_t)-1 < (int64_t)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(int64_t) == 8 && (((int64_t)-1 < (int64_t)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int64_t (*goml_c_signature_signed)(int64_t);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_signed), goml_c_signature_signed), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(uint64_t) == 8 && (((uint64_t)-1 < (uint64_t)0) == 0), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(uint64_t) == 8 && (((uint64_t)-1 < (uint64_t)0) == 0), "C ABI changed; regenerate GoML bindings");
typedef uint64_t (*goml_c_signature_unsigned)(uint64_t);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_unsigned), goml_c_signature_unsigned), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(double) == 8, "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(double) == 8, "C ABI changed; regenerate GoML bindings");
typedef double (*goml_c_signature_floating)(double);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_float), goml_c_signature_floating), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(bool) == sizeof(_Bool), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(bool) == sizeof(_Bool), "C ABI changed; regenerate GoML bindings");
typedef bool (*goml_c_signature_boolean)(bool);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_bool), goml_c_signature_boolean), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(SampleMode) <= 4, "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(SampleMode) <= 4, "C ABI changed; regenerate GoML bindings");
typedef SampleMode (*goml_c_signature_mode)(SampleMode);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_mode), goml_c_signature_mode), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int64_t) == 8 && (((int64_t)-1 < (int64_t)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef SampleCounter (*goml_c_signature_create)(int64_t);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_create), goml_c_signature_create), "C signature changed; regenerate GoML bindings");
typedef void (*goml_c_signature_destroy)(SampleCounter);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_destroy), goml_c_signature_destroy), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int64_t) == 8 && (((int64_t)-1 < (int64_t)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int64_t (*goml_c_signature_read)(SampleCounter);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_read), goml_c_signature_read), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int (*goml_c_signature_open)(SampleCounter *);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_open), goml_c_signature_open), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int32_t) == 4 && (((int32_t)-1 < (int32_t)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef void (*goml_c_signature_divide)(int32_t, int32_t *, int32_t *);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_divide), goml_c_signature_divide), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(size_t) == 8 && (((size_t)-1 < (size_t)0) == 0), "C ABI changed; regenerate GoML bindings");
typedef void (*goml_c_signature_invert)(unsigned char *, size_t);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_invert), goml_c_signature_invert), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(uint64_t) == 8 && (((uint64_t)-1 < (uint64_t)0) == 0), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(uint8_t) == 1 && (((uint8_t)-1 < (uint8_t)0) == 0), "C ABI changed; regenerate GoML bindings");
typedef uint64_t (*goml_c_signature_sum)(const unsigned char *, uint8_t);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_sum), goml_c_signature_sum), "C signature changed; regenerate GoML bindings");
typedef char * (*goml_c_signature_greet)(const char *);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_greet), goml_c_signature_greet), "C signature changed; regenerate GoML bindings");
typedef const char * (*goml_c_signature_null_text)(void);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_null_text), goml_c_signature_null_text), "C signature changed; regenerate GoML bindings");
typedef const char * (*goml_c_signature_raw_text)(void);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_raw_text), goml_c_signature_raw_text), "C signature changed; regenerate GoML bindings");
typedef char * (*goml_c_signature_long_text)(void);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_long_text), goml_c_signature_long_text), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int (*goml_c_signature_out_text)(char **);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_out_text), goml_c_signature_out_text), "C signature changed; regenerate GoML bindings");
typedef char * (*goml_c_signature_two_texts)(char **);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_two_texts), goml_c_signature_two_texts), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(int) == 4 && (((int)-1 < (int)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef int (*goml_c_signature_allocations)(void);
_Static_assert(__builtin_types_compatible_p(__typeof__(&sample_allocations), goml_c_signature_allocations), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(const unsigned long) == 8 && (((const unsigned long)-1 < (const unsigned long)0) == 0), "C ABI changed; regenerate GoML bindings");
_Static_assert((SAMPLE_MASK) == (const unsigned long)18446744073709551615ULL, "C constant changed; regenerate GoML bindings");
static size_t goml_c_bounded_length(const char *p, size_t limit) { size_t n = 0; while (n < limit && p[n] != 0) { n++; } return n; }
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

type Counter struct{ raw C.SampleCounter }

func GomlC_NullCounter() Counter                  { return Counter{} }
func GomlC_IsNullCounter(value Counter) bool      { return value.raw == nil }
func GomlC_EqualCounter(left, right Counter) bool { return left.raw == right.raw }
func GomlC_signed(a0 int64) (int64, error) {
	cResult := C.sample_signed(C.int64_t(a0))
	return int64(cResult), nil
}
func GomlC_unsigned(a0 uint64) (uint64, error) {
	cResult := C.sample_unsigned(C.uint64_t(a0))
	return uint64(cResult), nil
}
func GomlC_floating(a0 float64) (float64, error) {
	cResult := C.sample_float(C.double(a0))
	return float64(cResult), nil
}
func GomlC_boolean(a0 bool) (bool, error) {
	cResult := C.sample_bool(C.bool(a0))
	return bool(cResult), nil
}
func GomlC_mode(a0 int64) (int64, error) {
	if int64(C.SampleMode(a0)) != a0 {
		return 0, fmt.Errorf("C enum argument is out of range")
	}
	cResult := C.sample_mode(C.SampleMode(a0))
	return int64(cResult), nil
}
func GomlC_create(a0 int64) (Counter, error) {
	cResult := C.sample_create(C.int64_t(a0))
	return Counter{raw: cResult}, nil
}
func GomlC_destroy(a0 Counter) error {
	C.sample_destroy(a0.raw)
	return nil
}
func GomlC_read(a0 Counter) (int64, error) {
	cResult := C.sample_read(a0.raw)
	return int64(cResult), nil
}
func GomlC_open() (int32, Counter, error) {
	var cArg0 C.SampleCounter
	cResult := C.sample_open(&cArg0)
	return int32(cResult), Counter{raw: cArg0}, nil
}
func GomlC_divide(a0 int32) (int32, int32, error) {
	var cArg1 C.int
	var cArg2 C.int
	C.sample_divide(C.int32_t(a0), &cArg1, &cArg2)
	return int32(cArg1), int32(cArg2), nil
}
func GomlC_invert(a0 []byte) ([]byte, error) {
	if len(a0) > 67108864 {
		return nil, fmt.Errorf("C buffer exceeds 64 MiB")
	}
	var cArg0 unsafe.Pointer
	if len(a0) > 0 {
		cArg0 = C.CBytes(a0)
		defer C.free(cArg0)
	}
	if uint64(C.size_t(len(a0))) != uint64(len(a0)) {
		return nil, fmt.Errorf("buffer length does not fit C parameter")
	}
	C.sample_invert((*C.uchar)(cArg0), C.size_t(len(a0)))
	return C.GoBytes(cArg0, C.int(len(a0))), nil
}
func GomlC_sum(a0 []byte) (uint64, error) {
	if len(a0) > 67108864 {
		return 0, fmt.Errorf("C buffer exceeds 64 MiB")
	}
	var cArg0 unsafe.Pointer
	if len(a0) > 0 {
		cArg0 = C.CBytes(a0)
		defer C.free(cArg0)
	}
	if uint64(C.uint8_t(len(a0))) != uint64(len(a0)) {
		return 0, fmt.Errorf("buffer length does not fit C parameter")
	}
	cResult := C.sample_sum((*C.uchar)(cArg0), C.uint8_t(len(a0)))
	return uint64(cResult), nil
}
func GomlC_greet(a0 string) (string, bool, error) {
	if strings.IndexByte(a0, 0) >= 0 || len(a0) > 67108864 {
		return "", false, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg0 := C.CString(a0)
	defer C.free(unsafe.Pointer(cArg0))
	cResult := C.sample_greet(cArg0)
	defer C.sample_free(unsafe.Pointer(cResult))
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
func GomlC_null_text() (string, bool, error) {
	cResult := C.sample_null_text()
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
func GomlC_raw_text() (string, bool, error) {
	cResult := C.sample_raw_text()
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
func GomlC_long_text() (string, bool, error) {
	cResult := C.sample_long_text()
	defer C.sample_free(unsafe.Pointer(cResult))
	var cResultText string
	if cResult != nil {
		n := C.goml_c_bounded_length(cResult, 3)
		if n > 2 {
			return "", false, fmt.Errorf("C string exceeds configured copy limit")
		}
		cResultText = C.GoStringN(cResult, C.int(n))
	}
	return cResultText, cResult != nil, nil
}
func GomlC_out_text() (int32, string, bool, error) {
	var cArg0 *C.char
	cResult := C.sample_out_text(&cArg0)
	defer C.sample_free(unsafe.Pointer(cArg0))
	var cArg0Text string
	if cArg0 != nil {
		n := C.goml_c_bounded_length(cArg0, 1048577)
		if n > 1048576 {
			return 0, "", false, fmt.Errorf("C string exceeds configured copy limit")
		}
		cArg0Text = C.GoStringN(cArg0, C.int(n))
	}
	return int32(cResult), cArg0Text, cArg0 != nil, nil
}
func GomlC_two_texts() (string, bool, string, bool, error) {
	var cArg0 *C.char
	cResult := C.sample_two_texts(&cArg0)
	defer C.sample_free(unsafe.Pointer(cResult))
	defer C.sample_free(unsafe.Pointer(cArg0))
	var cResultText string
	if cResult != nil {
		n := C.goml_c_bounded_length(cResult, 1048577)
		if n > 1048576 {
			return "", false, "", false, fmt.Errorf("C string exceeds configured copy limit")
		}
		cResultText = C.GoStringN(cResult, C.int(n))
	}
	var cArg0Text string
	if cArg0 != nil {
		n := C.goml_c_bounded_length(cArg0, 3)
		if n > 2 {
			return "", false, "", false, fmt.Errorf("C string exceeds configured copy limit")
		}
		cArg0Text = C.GoStringN(cArg0, C.int(n))
	}
	return cResultText, cResult != nil, cArg0Text, cArg0 != nil, nil
}
func GomlC_allocations() (int32, error) {
	cResult := C.sample_allocations()
	return int32(cResult), nil
}
func GomlC_sample_mask() uint64 { return uint64(C.SAMPLE_MASK) }
