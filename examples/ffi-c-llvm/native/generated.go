package native

/*
#cgo CFLAGS: -std=c11 -I${SRCDIR}/.. -I/usr/lib/llvm-18/include
#cgo LDFLAGS: -L/usr/lib/llvm-18/lib -lLLVM-18
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <llvm-c/Core.h>
#include <llvm-c/Analysis.h>
typedef LLVMContextRef (*goml_c_signature_context_create)(void);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMContextCreate), goml_c_signature_context_create), "C signature changed; regenerate GoML bindings");
typedef void (*goml_c_signature_context_dispose)(LLVMContextRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMContextDispose), goml_c_signature_context_dispose), "C signature changed; regenerate GoML bindings");
typedef LLVMModuleRef (*goml_c_signature_module_create)(const char *, LLVMContextRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMModuleCreateWithNameInContext), goml_c_signature_module_create), "C signature changed; regenerate GoML bindings");
typedef void (*goml_c_signature_module_dispose)(LLVMModuleRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMDisposeModule), goml_c_signature_module_dispose), "C signature changed; regenerate GoML bindings");
typedef LLVMTypeRef (*goml_c_signature_int32_type)(LLVMContextRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMInt32TypeInContext), goml_c_signature_int32_type), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(unsigned int) == 4 && (((unsigned int)-1 < (unsigned int)0) == 0), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(LLVMBool) == 4 && (((LLVMBool)-1 < (LLVMBool)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef LLVMTypeRef (*goml_c_signature_function_type)(LLVMTypeRef, LLVMTypeRef *, unsigned int, LLVMBool);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMFunctionType), goml_c_signature_function_type), "C signature changed; regenerate GoML bindings");
typedef LLVMValueRef (*goml_c_signature_add_function)(LLVMModuleRef, const char *, LLVMTypeRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMAddFunction), goml_c_signature_add_function), "C signature changed; regenerate GoML bindings");
typedef LLVMBasicBlockRef (*goml_c_signature_append_block)(LLVMContextRef, LLVMValueRef, const char *);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMAppendBasicBlockInContext), goml_c_signature_append_block), "C signature changed; regenerate GoML bindings");
typedef LLVMBuilderRef (*goml_c_signature_builder_create)(LLVMContextRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMCreateBuilderInContext), goml_c_signature_builder_create), "C signature changed; regenerate GoML bindings");
typedef void (*goml_c_signature_builder_dispose)(LLVMBuilderRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMDisposeBuilder), goml_c_signature_builder_dispose), "C signature changed; regenerate GoML bindings");
typedef void (*goml_c_signature_position_at_end)(LLVMBuilderRef, LLVMBasicBlockRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMPositionBuilderAtEnd), goml_c_signature_position_at_end), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(unsigned long long) == 8 && (((unsigned long long)-1 < (unsigned long long)0) == 0), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(LLVMBool) == 4 && (((LLVMBool)-1 < (LLVMBool)0) == 1), "C ABI changed; regenerate GoML bindings");
typedef LLVMValueRef (*goml_c_signature_constant_int)(LLVMTypeRef, unsigned long long, LLVMBool);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMConstInt), goml_c_signature_constant_int), "C signature changed; regenerate GoML bindings");
typedef LLVMValueRef (*goml_c_signature_build_return)(LLVMBuilderRef, LLVMValueRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMBuildRet), goml_c_signature_build_return), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(LLVMBool) == 4 && (((LLVMBool)-1 < (LLVMBool)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert(sizeof(LLVMVerifierFailureAction) <= 4, "C ABI changed; regenerate GoML bindings");
typedef LLVMBool (*goml_c_signature_verify)(LLVMModuleRef, LLVMVerifierFailureAction, char **);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMVerifyModule), goml_c_signature_verify), "C signature changed; regenerate GoML bindings");
typedef char * (*goml_c_signature_print_module)(LLVMModuleRef);
_Static_assert(__builtin_types_compatible_p(__typeof__(&LLVMPrintModuleToString), goml_c_signature_print_module), "C signature changed; regenerate GoML bindings");
_Static_assert(sizeof(const int) == 4 && (((const int)-1 < (const int)0) == 1), "C ABI changed; regenerate GoML bindings");
_Static_assert((LLVMReturnStatusAction) == (const int)2ULL, "C constant changed; regenerate GoML bindings");
static size_t goml_c_bounded_length(const char *p, size_t limit) { size_t n = 0; while (n < limit && p[n] != 0) { n++; } return n; }
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

type Context struct{ raw C.LLVMContextRef }

func GomlC_NullContext() Context                  { return Context{} }
func GomlC_IsNullContext(value Context) bool      { return value.raw == nil }
func GomlC_EqualContext(left, right Context) bool { return left.raw == right.raw }

type Module struct{ raw C.LLVMModuleRef }

func GomlC_NullModule() Module                  { return Module{} }
func GomlC_IsNullModule(value Module) bool      { return value.raw == nil }
func GomlC_EqualModule(left, right Module) bool { return left.raw == right.raw }

type Type struct{ raw C.LLVMTypeRef }

func GomlC_NullType() Type                  { return Type{} }
func GomlC_IsNullType(value Type) bool      { return value.raw == nil }
func GomlC_EqualType(left, right Type) bool { return left.raw == right.raw }

type TypeArray struct{ raw *C.LLVMTypeRef }

func GomlC_NullTypeArray() TypeArray                  { return TypeArray{} }
func GomlC_IsNullTypeArray(value TypeArray) bool      { return value.raw == nil }
func GomlC_EqualTypeArray(left, right TypeArray) bool { return left.raw == right.raw }

type Value struct{ raw C.LLVMValueRef }

func GomlC_NullValue() Value                  { return Value{} }
func GomlC_IsNullValue(value Value) bool      { return value.raw == nil }
func GomlC_EqualValue(left, right Value) bool { return left.raw == right.raw }

type Block struct{ raw C.LLVMBasicBlockRef }

func GomlC_NullBlock() Block                  { return Block{} }
func GomlC_IsNullBlock(value Block) bool      { return value.raw == nil }
func GomlC_EqualBlock(left, right Block) bool { return left.raw == right.raw }

type Builder struct{ raw C.LLVMBuilderRef }

func GomlC_NullBuilder() Builder                  { return Builder{} }
func GomlC_IsNullBuilder(value Builder) bool      { return value.raw == nil }
func GomlC_EqualBuilder(left, right Builder) bool { return left.raw == right.raw }
func GomlC_context_create() (Context, error) {
	cResult := C.LLVMContextCreate()
	return Context{raw: cResult}, nil
}
func GomlC_context_dispose(a0 Context) error {
	C.LLVMContextDispose(a0.raw)
	return nil
}
func GomlC_module_create(a0 string, a1 Context) (Module, error) {
	if strings.IndexByte(a0, 0) >= 0 || len(a0) > 67108864 {
		return Module{}, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg0 := C.CString(a0)
	defer C.free(unsafe.Pointer(cArg0))
	cResult := C.LLVMModuleCreateWithNameInContext(cArg0, a1.raw)
	return Module{raw: cResult}, nil
}
func GomlC_module_dispose(a0 Module) error {
	C.LLVMDisposeModule(a0.raw)
	return nil
}
func GomlC_int32_type(a0 Context) (Type, error) {
	cResult := C.LLVMInt32TypeInContext(a0.raw)
	return Type{raw: cResult}, nil
}
func GomlC_function_type(a0 Type, a1 TypeArray, a2 uint32, a3 int32) (Type, error) {
	cResult := C.LLVMFunctionType(a0.raw, a1.raw, C.uint(a2), C.LLVMBool(a3))
	return Type{raw: cResult}, nil
}
func GomlC_add_function(a0 Module, a1 string, a2 Type) (Value, error) {
	if strings.IndexByte(a1, 0) >= 0 || len(a1) > 67108864 {
		return Value{}, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg1 := C.CString(a1)
	defer C.free(unsafe.Pointer(cArg1))
	cResult := C.LLVMAddFunction(a0.raw, cArg1, a2.raw)
	return Value{raw: cResult}, nil
}
func GomlC_append_block(a0 Context, a1 Value, a2 string) (Block, error) {
	if strings.IndexByte(a2, 0) >= 0 || len(a2) > 67108864 {
		return Block{}, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg2 := C.CString(a2)
	defer C.free(unsafe.Pointer(cArg2))
	cResult := C.LLVMAppendBasicBlockInContext(a0.raw, a1.raw, cArg2)
	return Block{raw: cResult}, nil
}
func GomlC_builder_create(a0 Context) (Builder, error) {
	cResult := C.LLVMCreateBuilderInContext(a0.raw)
	return Builder{raw: cResult}, nil
}
func GomlC_builder_dispose(a0 Builder) error {
	C.LLVMDisposeBuilder(a0.raw)
	return nil
}
func GomlC_position_at_end(a0 Builder, a1 Block) error {
	C.LLVMPositionBuilderAtEnd(a0.raw, a1.raw)
	return nil
}
func GomlC_constant_int(a0 Type, a1 uint64, a2 int32) (Value, error) {
	cResult := C.LLVMConstInt(a0.raw, C.ulonglong(a1), C.LLVMBool(a2))
	return Value{raw: cResult}, nil
}
func GomlC_build_return(a0 Builder, a1 Value) (Value, error) {
	cResult := C.LLVMBuildRet(a0.raw, a1.raw)
	return Value{raw: cResult}, nil
}
func GomlC_verify(a0 Module, a1 int64) (int32, string, bool, error) {
	if int64(C.LLVMVerifierFailureAction(a1)) != a1 {
		return 0, "", false, fmt.Errorf("C enum argument is out of range")
	}
	var cArg2 *C.char
	cResult := C.LLVMVerifyModule(a0.raw, C.LLVMVerifierFailureAction(a1), &cArg2)
	defer C.LLVMDisposeMessage(cArg2)
	var cArg2Text string
	if cArg2 != nil {
		n := C.goml_c_bounded_length(cArg2, 1048577)
		if n > 1048576 {
			return 0, "", false, fmt.Errorf("C string exceeds configured copy limit")
		}
		cArg2Text = C.GoStringN(cArg2, C.int(n))
	}
	return int32(cResult), cArg2Text, cArg2 != nil, nil
}
func GomlC_print_module(a0 Module) (string, bool, error) {
	cResult := C.LLVMPrintModuleToString(a0.raw)
	defer C.LLVMDisposeMessage(cResult)
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
func GomlC_RETURN_STATUS() int32 { return int32(C.LLVMReturnStatusAction) }
