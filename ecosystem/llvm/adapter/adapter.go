package adapter

/*
#cgo CFLAGS: -I/usr/lib/llvm-18/include
#cgo LDFLAGS: -L/usr/lib/llvm-18/lib -lLLVM-18
#include <stdlib.h>
#include <llvm-c/Core.h>
#include <llvm-c/Analysis.h>
#include <llvm-c/IRReader.h>
#include <llvm-c/BitReader.h>
#include <llvm-c/BitWriter.h>
#include <llvm-c/Error.h>
#include <llvm-c/Target.h>
#include <llvm-c/TargetMachine.h>
#include <llvm-c/Transforms/PassBuilder.h>
static int initialize_native(void) {
    return LLVMInitializeNativeTarget() || LLVMInitializeNativeAsmPrinter() || LLVMInitializeNativeAsmParser();
}
*/
import "C"

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
)

type Handle interface{ llvmHandle() }

func (*node) llvmHandle() {}

type contextState struct {
	mu       sync.Mutex
	raw      C.LLVMContextRef
	modules  map[*node]bool
	builders map[*node]bool
}

type node struct {
	context *contextState
	kind    int
	module  *node
	epoch   uint64
	raw     unsafe.Pointer
	owner   C.LLVMValueRef
	block   C.LLVMBasicBlockRef
	closed  bool
}

const (
	contextKind = iota
	moduleKind
	typeKind
	valueKind
	functionKind
	blockKind
	builderKind
)

func NewContext() Handle {
	ctx := &contextState{raw: C.LLVMContextCreate(), modules: make(map[*node]bool), builders: make(map[*node]bool)}
	return &node{context: ctx, kind: contextKind}
}

func enter(handles ...Handle) (*contextState, []*node, string) {
	if len(handles) == 0 {
		return nil, nil, "argument: missing handle"
	}
	first, ok := handles[0].(*node)
	if !ok || first == nil {
		return nil, nil, "argument: invalid handle"
	}
	ctx := first.context
	ctx.mu.Lock()
	fail := func(message string) (*contextState, []*node, string) { ctx.mu.Unlock(); return nil, nil, message }
	if ctx.raw == nil {
		return fail("closed: context")
	}
	values := make([]*node, len(handles))
	for i, handle := range handles {
		value, ok := handle.(*node)
		if !ok || value == nil {
			return fail("argument: invalid handle")
		}
		if value.context != ctx {
			return fail("argument: handles belong to different contexts")
		}
		if value.closed {
			return fail("closed: resource")
		}
		if value.module != nil {
			if value.module.closed {
				return fail("closed: module")
			}
			if value.epoch != value.module.epoch {
				return fail("stale: handle invalidated by module transformation")
			}
		}
		values[i] = value
	}
	return ctx, values, ""
}

func Close(handle Handle) string {
	value, ok := handle.(*node)
	if !ok || value == nil {
		return "argument: invalid handle"
	}
	ctx := value.context
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if value.closed || ctx.raw == nil {
		return ""
	}
	switch value.kind {
	case contextKind:
		for builder := range ctx.builders {
			C.LLVMDisposeBuilder(C.LLVMBuilderRef(builder.raw))
			builder.closed = true
		}
		for module := range ctx.modules {
			C.LLVMDisposeModule(C.LLVMModuleRef(module.raw))
			module.closed = true
		}
		C.LLVMContextDispose(ctx.raw)
		ctx.raw = nil
		clear(ctx.builders)
		clear(ctx.modules)
	case moduleKind:
		for builder := range ctx.builders {
			if builder.module == value {
				C.LLVMClearInsertionPosition(C.LLVMBuilderRef(builder.raw))
				builder.module = nil
				builder.owner = nil
				builder.block = nil
			}
		}
		C.LLVMDisposeModule(C.LLVMModuleRef(value.raw))
		delete(ctx.modules, value)
	case builderKind:
		C.LLVMDisposeBuilder(C.LLVMBuilderRef(value.raw))
		delete(ctx.builders, value)
	default:
		return "argument: resource is owned by its context or module"
	}
	value.closed = true
	return ""
}

func IsClosed(handle Handle) bool {
	value, ok := handle.(*node)
	if !ok || value == nil {
		return true
	}
	ctx := value.context
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.raw == nil || value.closed || value.module != nil && (value.module.closed || value.epoch != value.module.epoch)
}

func text(value string) (*C.char, func(), string) {
	if strings.ContainsRune(value, 0) {
		return nil, func() {}, "argument: text contains NUL"
	}
	raw := C.CString(value)
	return raw, func() { C.free(unsafe.Pointer(raw)) }, ""
}

func message(raw *C.char) string {
	if raw == nil {
		return ""
	}
	value := C.GoString(raw)
	C.LLVMDisposeMessage(raw)
	return value
}

func wrapType(ctx *contextState, raw C.LLVMTypeRef) Handle {
	return &node{context: ctx, kind: typeKind, raw: unsafe.Pointer(raw)}
}
func wrapValue(ctx *contextState, module *node, owner C.LLVMValueRef, raw C.LLVMValueRef) Handle {
	result := &node{context: ctx, kind: valueKind, module: module, owner: owner, raw: unsafe.Pointer(raw)}
	if module != nil {
		result.epoch = module.epoch
	}
	return result
}
func newModule(ctx *contextState, raw C.LLVMModuleRef) Handle {
	result := &node{context: ctx, kind: moduleKind, raw: unsafe.Pointer(raw)}
	ctx.modules[result] = true
	return result
}

func NewModule(handle Handle, name string) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != contextKind {
		return nil, "argument: expected context"
	}
	raw, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	return newModule(ctx, C.LLVMModuleCreateWithNameInContext(raw, ctx.raw)), ""
}

func Parse(handle Handle, source string, bitcode bool) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != contextKind {
		return nil, "argument: expected context"
	}
	raw := C.CBytes([]byte(source))
	defer C.free(raw)
	name := C.CString("goml")
	defer C.free(unsafe.Pointer(name))
	buffer := C.LLVMCreateMemoryBufferWithMemoryRangeCopy((*C.char)(raw), C.size_t(len(source)), name)
	var module C.LLVMModuleRef
	var diagnostic *C.char
	if bitcode {
		defer C.LLVMDisposeMemoryBuffer(buffer)
		if C.LLVMParseBitcodeInContext(ctx.raw, buffer, &module, &diagnostic) != 0 {
			return nil, "llvm: " + message(diagnostic)
		}
	} else if C.LLVMParseIRInContext(ctx.raw, buffer, &module, &diagnostic) != 0 {
		return nil, "llvm: " + message(diagnostic)
	}
	return newModule(ctx, module), ""
}

func Print(handle Handle) (string, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return "", err
	}
	defer ctx.mu.Unlock()
	n := nodes[0]
	switch n.kind {
	case moduleKind:
		return message(C.LLVMPrintModuleToString(C.LLVMModuleRef(n.raw))), ""
	case typeKind:
		return message(C.LLVMPrintTypeToString(C.LLVMTypeRef(n.raw))), ""
	case valueKind, functionKind:
		return message(C.LLVMPrintValueToString(C.LLVMValueRef(n.raw))), ""
	default:
		return "", "argument: resource cannot be printed"
	}
}

func verify(module *node) string {
	var diagnostic *C.char
	status := C.LLVMVerifyModule(C.LLVMModuleRef(module.raw), C.LLVMReturnStatusAction, &diagnostic)
	result := message(diagnostic)
	if status != 0 {
		return "verification: " + result
	}
	return ""
}

func Verify(handle Handle) string {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != moduleKind {
		return "argument: expected module"
	}
	return verify(nodes[0])
}

func Clone(handle Handle) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != moduleKind {
		return nil, "argument: expected module"
	}
	return newModule(ctx, C.LLVMCloneModule(C.LLVMModuleRef(nodes[0].raw))), ""
}

func Bitcode(handle Handle) ([]byte, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != moduleKind {
		return nil, "argument: expected module"
	}
	if err := verify(nodes[0]); err != "" {
		return nil, err
	}
	buffer := C.LLVMWriteBitcodeToMemoryBuffer(C.LLVMModuleRef(nodes[0].raw))
	defer C.LLVMDisposeMemoryBuffer(buffer)
	length := C.LLVMGetBufferSize(buffer)
	if uint64(length) > 2147483647 {
		return nil, "argument: bitcode exceeds 2 GiB"
	}
	return C.GoBytes(unsafe.Pointer(C.LLVMGetBufferStart(buffer)), C.int(length)), ""
}

func Primitive(handle Handle, kind int, parameter uint64) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != contextKind {
		return nil, "argument: expected context"
	}
	var result C.LLVMTypeRef
	switch kind {
	case 0:
		result = C.LLVMVoidTypeInContext(ctx.raw)
	case 1:
		if parameter == 0 || parameter > 65536 {
			return nil, "argument: integer width must be 1..65536"
		}
		result = C.LLVMIntTypeInContext(ctx.raw, C.uint(parameter))
	case 2:
		result = C.LLVMFloatTypeInContext(ctx.raw)
	case 3:
		result = C.LLVMDoubleTypeInContext(ctx.raw)
	case 4:
		if parameter > 16777215 {
			return nil, "argument: address space exceeds 24 bits"
		}
		result = C.LLVMPointerTypeInContext(ctx.raw, C.uint(parameter))
	default:
		return nil, "argument: unknown primitive type"
	}
	return wrapType(ctx, result), ""
}

func firstClass(ty C.LLVMTypeRef) bool {
	switch C.LLVMGetTypeKind(ty) {
	case C.LLVMIntegerTypeKind, C.LLVMFloatTypeKind, C.LLVMDoubleTypeKind, C.LLVMPointerTypeKind, C.LLVMStructTypeKind, C.LLVMArrayTypeKind:
		return true
	default:
		return false
	}
}

func Composite(handle Handle, kind int, members []Handle, count uint64, flag bool) (Handle, string) {
	ctx, nodes, err := enter(append([]Handle{handle}, members...)...)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != contextKind {
		return nil, "argument: expected context"
	}
	types := make([]C.LLVMTypeRef, len(members))
	for i, n := range nodes[1:] {
		if n.kind != typeKind {
			return nil, "argument: expected types"
		}
		types[i] = C.LLVMTypeRef(n.raw)
	}
	var result C.LLVMTypeRef
	truth := C.LLVMBool(0)
	if flag {
		truth = 1
	}
	switch kind {
	case 0:
		if len(types) != 1 || !firstClass(types[0]) || C.LLVMTypeIsSized(types[0]) == 0 || count > 2147483647 {
			return nil, "argument: invalid array type or length"
		}
		result = C.LLVMArrayType2(types[0], C.uint64_t(count))
	case 1:
		for _, ty := range types {
			if !firstClass(ty) {
				return nil, "argument: invalid struct member"
			}
		}
		result = C.LLVMStructTypeInContext(ctx.raw, unsafe.SliceData(types), C.uint(len(types)), truth)
	case 2:
		if len(types) == 0 {
			return nil, "argument: missing return type"
		}
		if !firstClass(types[0]) && C.LLVMGetTypeKind(types[0]) != C.LLVMVoidTypeKind {
			return nil, "argument: invalid return type"
		}
		for _, ty := range types[1:] {
			if !firstClass(ty) {
				return nil, "argument: invalid function parameter"
			}
		}
		result = C.LLVMFunctionType(types[0], unsafe.SliceData(types[1:]), C.uint(len(types)-1), truth)
	default:
		return nil, "argument: unknown composite type"
	}
	return wrapType(ctx, result), ""
}

func TypeEqual(left, right Handle) (bool, string) {
	ctx, nodes, err := enter(left, right)
	if err != "" {
		return false, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != typeKind || nodes[1].kind != typeKind {
		return false, "argument: expected types"
	}
	return nodes[0].raw == nodes[1].raw, ""
}

func ValueType(handle Handle) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	n := nodes[0]
	if n.kind == functionKind {
		return wrapType(ctx, C.LLVMGlobalGetValueType(C.LLVMValueRef(n.raw))), ""
	}
	if n.kind != valueKind {
		return nil, "argument: expected value"
	}
	return wrapType(ctx, C.LLVMTypeOf(C.LLVMValueRef(n.raw))), ""
}

func Constant(handle Handle, kind int, integer uint64, real float64) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != typeKind {
		return nil, "argument: expected type"
	}
	ty := C.LLVMTypeRef(nodes[0].raw)
	var result C.LLVMValueRef
	switch kind {
	case 0, 1:
		if C.LLVMGetTypeKind(ty) != C.LLVMIntegerTypeKind {
			return nil, "argument: integer constant requires integer type"
		}
		extend := C.LLVMBool(0)
		if kind == 1 {
			extend = 1
		}
		result = C.LLVMConstInt(ty, C.ulonglong(integer), extend)
	case 2:
		if C.LLVMGetTypeKind(ty) != C.LLVMFloatTypeKind && C.LLVMGetTypeKind(ty) != C.LLVMDoubleTypeKind {
			return nil, "argument: floating constant requires float type"
		}
		result = C.LLVMConstReal(ty, C.double(real))
	case 3:
		if !firstClass(ty) {
			return nil, "argument: invalid null constant type"
		}
		result = C.LLVMConstNull(ty)
	case 4:
		if !firstClass(ty) {
			return nil, "argument: invalid undef constant type"
		}
		result = C.LLVMGetUndef(ty)
	default:
		return nil, "argument: unknown constant kind"
	}
	return wrapValue(ctx, nil, nil, result), ""
}

func Aggregate(handle Handle, elements []Handle) (Handle, string) {
	ctx, nodes, err := enter(append([]Handle{handle}, elements...)...)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != typeKind {
		return nil, "argument: expected type"
	}
	ty := C.LLVMTypeRef(nodes[0].raw)
	kind := C.LLVMGetTypeKind(ty)
	var types []C.LLVMTypeRef
	if kind == C.LLVMStructTypeKind {
		types = make([]C.LLVMTypeRef, int(C.LLVMCountStructElementTypes(ty)))
		C.LLVMGetStructElementTypes(ty, unsafe.SliceData(types))
	} else if kind == C.LLVMArrayTypeKind {
		if uint64(len(elements)) != uint64(C.LLVMGetArrayLength2(ty)) {
			return nil, "argument: array constant length mismatch"
		}
		types = make([]C.LLVMTypeRef, len(elements))
		for i := range types {
			types[i] = C.LLVMGetElementType(ty)
		}
	} else {
		return nil, "argument: expected array or struct type"
	}
	if len(types) != len(elements) {
		return nil, "argument: aggregate constant length mismatch"
	}
	values := make([]C.LLVMValueRef, len(elements))
	for i, n := range nodes[1:] {
		if n.kind != valueKind || n.module != nil {
			return nil, "argument: expected context-owned constants"
		}
		values[i] = C.LLVMValueRef(n.raw)
		if C.LLVMIsConstant(values[i]) == 0 || C.LLVMTypeOf(values[i]) != types[i] {
			return nil, "argument: constant element type mismatch"
		}
	}
	var result C.LLVMValueRef
	if kind == C.LLVMStructTypeKind {
		result = C.LLVMConstNamedStruct(ty, unsafe.SliceData(values), C.uint(len(values)))
	} else {
		result = C.LLVMConstArray2(C.LLVMGetElementType(ty), unsafe.SliceData(values), C.uint64_t(len(values)))
	}
	return wrapValue(ctx, nil, nil, result), ""
}

func Function(handle Handle, name string, signature Handle, create bool) (Handle, string) {
	handles := []Handle{handle}
	if create {
		handles = append(handles, signature)
	}
	ctx, nodes, err := enter(handles...)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	module := nodes[0]
	if module.kind != moduleKind {
		return nil, "argument: expected module"
	}
	if name == "" {
		return nil, "argument: function name is empty"
	}
	rawName, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	fn := C.LLVMGetNamedFunction(C.LLVMModuleRef(module.raw), rawName)
	if create {
		if nodes[1].kind != typeKind || C.LLVMGetTypeKind(C.LLVMTypeRef(nodes[1].raw)) != C.LLVMFunctionTypeKind {
			return nil, "argument: expected function type"
		}
		if fn != nil || C.LLVMGetNamedGlobal(C.LLVMModuleRef(module.raw), rawName) != nil || C.LLVMGetNamedGlobalAlias(C.LLVMModuleRef(module.raw), rawName, C.size_t(len(name))) != nil || C.LLVMGetNamedGlobalIFunc(C.LLVMModuleRef(module.raw), rawName, C.size_t(len(name))) != nil {
			return nil, "argument: symbol already exists"
		}
		fn = C.LLVMAddFunction(C.LLVMModuleRef(module.raw), rawName, C.LLVMTypeRef(nodes[1].raw))
	}
	if fn == nil {
		return nil, "argument: function not found"
	}
	return &node{context: ctx, kind: functionKind, module: module, epoch: module.epoch, raw: unsafe.Pointer(fn)}, ""
}

func Parameter(handle Handle, index int) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	n := nodes[0]
	if n.kind != functionKind {
		return nil, "argument: expected function"
	}
	fn := C.LLVMValueRef(n.raw)
	if index < 0 || index >= int(C.LLVMCountParams(fn)) {
		return nil, "argument: parameter index out of bounds"
	}
	return wrapValue(ctx, n.module, fn, C.LLVMGetParam(fn, C.uint(index))), ""
}

func SetName(handle Handle, name string) string {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return err
	}
	defer ctx.mu.Unlock()
	n := nodes[0]
	if n.kind != valueKind {
		return "argument: expected value"
	}
	value := C.LLVMValueRef(n.raw)
	if C.LLVMIsConstant(value) != 0 || C.LLVMGetTypeKind(C.LLVMTypeOf(value)) == C.LLVMVoidTypeKind {
		return "argument: constant or void value cannot be named"
	}
	raw, free, err := text(name)
	if err != "" {
		return err
	}
	defer free()
	C.LLVMSetValueName2(value, raw, C.size_t(len(name)))
	return ""
}

func NewBlock(handle Handle, name string) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	n := nodes[0]
	if n.kind != functionKind {
		return nil, "argument: expected function"
	}
	raw, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	owner := C.LLVMValueRef(n.raw)
	return &node{context: ctx, kind: blockKind, module: n.module, epoch: n.epoch, owner: owner, raw: unsafe.Pointer(C.LLVMAppendBasicBlockInContext(ctx.raw, owner, raw))}, ""
}

func NewBuilder(handle Handle) (Handle, string) {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	if nodes[0].kind != contextKind {
		return nil, "argument: expected context"
	}
	result := &node{context: ctx, kind: builderKind, raw: unsafe.Pointer(C.LLVMCreateBuilderInContext(ctx.raw))}
	ctx.builders[result] = true
	return result, ""
}

func Position(builder, block Handle) string {
	ctx, nodes, err := enter(builder, block)
	if err != "" {
		return err
	}
	defer ctx.mu.Unlock()
	b, target := nodes[0], nodes[1]
	if b.kind != builderKind || target.kind != blockKind {
		return "argument: expected builder and block"
	}
	b.module, b.epoch, b.owner, b.block = target.module, target.epoch, target.owner, C.LLVMBasicBlockRef(target.raw)
	C.LLVMPositionBuilderAtEnd(C.LLVMBuilderRef(b.raw), b.block)
	return ""
}

func insertion(builder *node) string {
	if builder.kind != builderKind || builder.module == nil || builder.block == nil {
		return "argument: builder has no insertion block"
	}
	if C.LLVMGetBasicBlockTerminator(builder.block) != nil {
		return "argument: block already has a terminator"
	}
	return ""
}

func operand(builder, value *node) string {
	if value.kind != valueKind {
		return "argument: expected value"
	}
	if value.module != nil && value.module != builder.module {
		return "argument: value belongs to another module"
	}
	if value.owner != nil && value.owner != builder.owner {
		return "argument: value belongs to another function"
	}
	return ""
}

func instruction(builder *node, value C.LLVMValueRef) Handle {
	if C.LLVMIsConstant(value) != 0 {
		return wrapValue(builder.context, nil, nil, value)
	}
	return wrapValue(builder.context, builder.module, builder.owner, value)
}

func Binary(builder Handle, operation int, left, right Handle, name string) (Handle, string) {
	ctx, nodes, err := enter(builder, left, right)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	b := nodes[0]
	if err := insertion(b); err != "" {
		return nil, err
	}
	for _, n := range nodes[1:] {
		if err := operand(b, n); err != "" {
			return nil, err
		}
	}
	a, z := C.LLVMValueRef(nodes[1].raw), C.LLVMValueRef(nodes[2].raw)
	ty := C.LLVMTypeOf(a)
	if ty != C.LLVMTypeOf(z) {
		return nil, "argument: binary operand types differ"
	}
	kind := C.LLVMGetTypeKind(ty)
	if operation >= 0 && operation <= 12 {
		if kind != C.LLVMIntegerTypeKind {
			return nil, "argument: integer operands required"
		}
	} else if operation >= 13 && operation <= 17 {
		if kind != C.LLVMFloatTypeKind && kind != C.LLVMDoubleTypeKind {
			return nil, "argument: floating operands required"
		}
	} else {
		return nil, "argument: unknown binary operation"
	}
	raw, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	r := C.LLVMBuilderRef(b.raw)
	var result C.LLVMValueRef
	switch operation {
	case 0:
		result = C.LLVMBuildAdd(r, a, z, raw)
	case 1:
		result = C.LLVMBuildSub(r, a, z, raw)
	case 2:
		result = C.LLVMBuildMul(r, a, z, raw)
	case 3:
		result = C.LLVMBuildSDiv(r, a, z, raw)
	case 4:
		result = C.LLVMBuildUDiv(r, a, z, raw)
	case 5:
		result = C.LLVMBuildSRem(r, a, z, raw)
	case 6:
		result = C.LLVMBuildURem(r, a, z, raw)
	case 7:
		result = C.LLVMBuildAnd(r, a, z, raw)
	case 8:
		result = C.LLVMBuildOr(r, a, z, raw)
	case 9:
		result = C.LLVMBuildXor(r, a, z, raw)
	case 10:
		result = C.LLVMBuildShl(r, a, z, raw)
	case 11:
		result = C.LLVMBuildAShr(r, a, z, raw)
	case 12:
		result = C.LLVMBuildLShr(r, a, z, raw)
	case 13:
		result = C.LLVMBuildFAdd(r, a, z, raw)
	case 14:
		result = C.LLVMBuildFSub(r, a, z, raw)
	case 15:
		result = C.LLVMBuildFMul(r, a, z, raw)
	case 16:
		result = C.LLVMBuildFDiv(r, a, z, raw)
	case 17:
		result = C.LLVMBuildFRem(r, a, z, raw)
	}
	return instruction(b, result), ""
}

func Compare(builder Handle, predicate int, floating bool, left, right Handle, name string) (Handle, string) {
	ctx, nodes, err := enter(builder, left, right)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	b := nodes[0]
	if err := insertion(b); err != "" {
		return nil, err
	}
	for _, n := range nodes[1:] {
		if err := operand(b, n); err != "" {
			return nil, err
		}
	}
	a, z := C.LLVMValueRef(nodes[1].raw), C.LLVMValueRef(nodes[2].raw)
	if C.LLVMTypeOf(a) != C.LLVMTypeOf(z) {
		return nil, "argument: comparison operand types differ"
	}
	kind := C.LLVMGetTypeKind(C.LLVMTypeOf(a))
	if floating {
		if predicate < 0 || predicate > 15 || kind != C.LLVMFloatTypeKind && kind != C.LLVMDoubleTypeKind {
			return nil, "argument: invalid floating comparison"
		}
	} else if predicate < 32 || predicate > 41 || kind != C.LLVMIntegerTypeKind && kind != C.LLVMPointerTypeKind {
		return nil, "argument: invalid integer comparison"
	}
	raw, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	var result C.LLVMValueRef
	if floating {
		result = C.LLVMBuildFCmp(C.LLVMBuilderRef(b.raw), C.LLVMRealPredicate(predicate), a, z, raw)
	} else {
		result = C.LLVMBuildICmp(C.LLVMBuilderRef(b.raw), C.LLVMIntPredicate(predicate), a, z, raw)
	}
	return instruction(b, result), ""
}

func Cast(builder Handle, operation int, value, target Handle, name string) (Handle, string) {
	ctx, nodes, err := enter(builder, value, target)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	b, n, t := nodes[0], nodes[1], nodes[2]
	if err := insertion(b); err != "" {
		return nil, err
	}
	if err := operand(b, n); err != "" {
		return nil, err
	}
	if t.kind != typeKind {
		return nil, "argument: expected destination type"
	}
	source, dest := C.LLVMTypeOf(C.LLVMValueRef(n.raw)), C.LLVMTypeRef(t.raw)
	a, z := C.LLVMGetTypeKind(source), C.LLVMGetTypeKind(dest)
	integerA, integerZ := a == C.LLVMIntegerTypeKind, z == C.LLVMIntegerTypeKind
	floatA, floatZ := a == C.LLVMFloatTypeKind || a == C.LLVMDoubleTypeKind, z == C.LLVMFloatTypeKind || z == C.LLVMDoubleTypeKind
	valid := false
	var opcode C.LLVMOpcode
	switch operation {
	case 0, 1:
		valid = integerA && integerZ && C.LLVMGetIntTypeWidth(source) < C.LLVMGetIntTypeWidth(dest)
		if operation == 0 {
			opcode = C.LLVMZExt
		} else {
			opcode = C.LLVMSExt
		}
	case 2:
		valid = integerA && integerZ && C.LLVMGetIntTypeWidth(source) > C.LLVMGetIntTypeWidth(dest)
		opcode = C.LLVMTrunc
	case 3:
		valid = integerA && floatZ
		opcode = C.LLVMSIToFP
	case 4:
		valid = integerA && floatZ
		opcode = C.LLVMUIToFP
	case 5:
		valid = floatA && integerZ
		opcode = C.LLVMFPToSI
	case 6:
		valid = floatA && integerZ
		opcode = C.LLVMFPToUI
	case 7:
		valid = a == C.LLVMFloatTypeKind && z == C.LLVMDoubleTypeKind
		opcode = C.LLVMFPExt
	case 8:
		valid = a == C.LLVMDoubleTypeKind && z == C.LLVMFloatTypeKind
		opcode = C.LLVMFPTrunc
	case 9:
		valid = a == C.LLVMPointerTypeKind && integerZ
		opcode = C.LLVMPtrToInt
	case 10:
		valid = integerA && z == C.LLVMPointerTypeKind
		opcode = C.LLVMIntToPtr
	case 11:
		bits := func(ty C.LLVMTypeRef) int {
			switch C.LLVMGetTypeKind(ty) {
			case C.LLVMIntegerTypeKind:
				return int(C.LLVMGetIntTypeWidth(ty))
			case C.LLVMFloatTypeKind:
				return 32
			case C.LLVMDoubleTypeKind:
				return 64
			}
			return 0
		}
		valid = bits(source) > 0 && bits(source) == bits(dest) || a == C.LLVMPointerTypeKind && z == C.LLVMPointerTypeKind && C.LLVMGetPointerAddressSpace(source) == C.LLVMGetPointerAddressSpace(dest)
		opcode = C.LLVMBitCast
	}
	if !valid {
		return nil, "argument: invalid cast or widths"
	}
	raw, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	return instruction(b, C.LLVMBuildCast(C.LLVMBuilderRef(b.raw), opcode, C.LLVMValueRef(n.raw), dest, raw)), ""
}

func Emit(builder Handle, operation int, arguments []Handle, name string) (Handle, string) {
	ctx, nodes, err := enter(append([]Handle{builder}, arguments...)...)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	b := nodes[0]
	if err := insertion(b); err != "" {
		return nil, err
	}
	ns := nodes[1:]
	counts := []int{1, 0, 1, 3, 1, 2, 2, 3, 1, 0}
	if operation < 0 || operation >= len(counts) || len(ns) != counts[operation] {
		return nil, "argument: invalid instruction argument count"
	}
	raw, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	r := C.LLVMBuilderRef(b.raw)
	value := func(index int) (C.LLVMValueRef, string) {
		n := ns[index]
		if err := operand(b, n); err != "" {
			return nil, err
		}
		return C.LLVMValueRef(n.raw), ""
	}
	block := func(index int) (C.LLVMBasicBlockRef, string) {
		n := ns[index]
		if n.kind != blockKind || n.module != b.module || n.owner != b.owner {
			return nil, "argument: branch block belongs to another function"
		}
		return C.LLVMBasicBlockRef(n.raw), ""
	}
	ty := func(index int) (C.LLVMTypeRef, string) {
		n := ns[index]
		if n.kind != typeKind {
			return nil, "argument: expected type"
		}
		t := C.LLVMTypeRef(n.raw)
		if !firstClass(t) || C.LLVMTypeIsSized(t) == 0 {
			return nil, "argument: expected sized first-class type"
		}
		return t, ""
	}
	var result C.LLVMValueRef
	switch operation {
	case 0:
		v, e := value(0)
		if e != "" {
			return nil, e
		}
		if C.LLVMGetTypeKind(C.LLVMTypeOf(v)) == C.LLVMVoidTypeKind || C.LLVMTypeOf(v) != C.LLVMGetReturnType(C.LLVMGlobalGetValueType(b.owner)) {
			return nil, "argument: return type mismatch"
		}
		result = C.LLVMBuildRet(r, v)
	case 1:
		if C.LLVMGetTypeKind(C.LLVMGetReturnType(C.LLVMGlobalGetValueType(b.owner))) != C.LLVMVoidTypeKind {
			return nil, "argument: function does not return void"
		}
		result = C.LLVMBuildRetVoid(r)
	case 2:
		target, e := block(0)
		if e != "" {
			return nil, e
		}
		result = C.LLVMBuildBr(r, target)
	case 3:
		condition, e := value(0)
		if e != "" {
			return nil, e
		}
		if C.LLVMTypeOf(condition) != C.LLVMInt1TypeInContext(ctx.raw) {
			return nil, "argument: branch condition must be i1"
		}
		yes, e := block(1)
		if e != "" {
			return nil, e
		}
		no, e := block(2)
		if e != "" {
			return nil, e
		}
		result = C.LLVMBuildCondBr(r, condition, yes, no)
	case 4:
		t, e := ty(0)
		if e != "" {
			return nil, e
		}
		result = C.LLVMBuildAlloca(r, t, raw)
	case 5:
		t, e := ty(0)
		if e != "" {
			return nil, e
		}
		p, e := value(1)
		if e != "" {
			return nil, e
		}
		if C.LLVMGetTypeKind(C.LLVMTypeOf(p)) != C.LLVMPointerTypeKind {
			return nil, "argument: load requires pointer"
		}
		result = C.LLVMBuildLoad2(r, t, p, raw)
	case 6:
		v, e := value(0)
		if e != "" {
			return nil, e
		}
		p, e := value(1)
		if e != "" {
			return nil, e
		}
		if !firstClass(C.LLVMTypeOf(v)) || C.LLVMTypeIsSized(C.LLVMTypeOf(v)) == 0 || C.LLVMGetTypeKind(C.LLVMTypeOf(p)) != C.LLVMPointerTypeKind {
			return nil, "argument: invalid store types"
		}
		result = C.LLVMBuildStore(r, v, p)
	case 7:
		c, e := value(0)
		if e != "" {
			return nil, e
		}
		a, e := value(1)
		if e != "" {
			return nil, e
		}
		z, e := value(2)
		if e != "" {
			return nil, e
		}
		if C.LLVMTypeOf(c) != C.LLVMInt1TypeInContext(ctx.raw) || C.LLVMTypeOf(a) != C.LLVMTypeOf(z) || !firstClass(C.LLVMTypeOf(a)) {
			return nil, "argument: invalid select types"
		}
		result = C.LLVMBuildSelect(r, c, a, z, raw)
	case 8:
		t, e := ty(0)
		if e != "" {
			return nil, e
		}
		for i := C.LLVMGetFirstInstruction(b.block); i != nil; i = C.LLVMGetNextInstruction(i) {
			if C.LLVMIsAPHINode(i) == nil {
				return nil, "argument: phi must precede non-phi instructions"
			}
		}
		result = C.LLVMBuildPhi(r, t, raw)
	case 9:
		result = C.LLVMBuildUnreachable(r)
	}
	return instruction(b, result), ""
}

func Call(builder, function Handle, arguments []Handle, name string) (Handle, string) {
	ctx, nodes, err := enter(append([]Handle{builder, function}, arguments...)...)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	b, fn := nodes[0], nodes[1]
	if err := insertion(b); err != "" {
		return nil, err
	}
	if fn.kind != functionKind || fn.module != b.module {
		return nil, "argument: callee belongs to another module or is not a function"
	}
	functionType := C.LLVMGlobalGetValueType(C.LLVMValueRef(fn.raw))
	count := int(C.LLVMCountParamTypes(functionType))
	if len(arguments) < count || len(arguments) != count && C.LLVMIsFunctionVarArg(functionType) == 0 {
		return nil, "argument: call argument count mismatch"
	}
	types := make([]C.LLVMTypeRef, count)
	C.LLVMGetParamTypes(functionType, unsafe.SliceData(types))
	values := make([]C.LLVMValueRef, len(arguments))
	for i, n := range nodes[2:] {
		if err := operand(b, n); err != "" {
			return nil, err
		}
		values[i] = C.LLVMValueRef(n.raw)
		if !firstClass(C.LLVMTypeOf(values[i])) || i < count && C.LLVMTypeOf(values[i]) != types[i] {
			return nil, "argument: call argument type mismatch"
		}
	}
	if C.LLVMGetTypeKind(C.LLVMGetReturnType(functionType)) == C.LLVMVoidTypeKind && name != "" {
		return nil, "argument: void call cannot have a name"
	}
	raw, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	return instruction(b, C.LLVMBuildCall2(C.LLVMBuilderRef(b.raw), functionType, C.LLVMValueRef(fn.raw), unsafe.SliceData(values), C.uint(len(values)), raw)), ""
}

func Incoming(phi Handle, values, blocks []Handle) string {
	if len(values) != len(blocks) || len(values) == 0 {
		return "argument: phi requires matching nonempty values and blocks"
	}
	ctx, nodes, err := enter(append(append([]Handle{phi}, values...), blocks...)...)
	if err != "" {
		return err
	}
	defer ctx.mu.Unlock()
	p := nodes[0]
	if p.kind != valueKind || p.module == nil || C.LLVMIsAPHINode(C.LLVMValueRef(p.raw)) == nil {
		return "argument: expected phi instruction"
	}
	raw := C.LLVMValueRef(p.raw)
	seen := make(map[C.LLVMBasicBlockRef]bool)
	for i := C.uint(0); i < C.LLVMCountIncoming(raw); i++ {
		seen[C.LLVMGetIncomingBlock(raw, i)] = true
	}
	vs := make([]C.LLVMValueRef, len(values))
	bs := make([]C.LLVMBasicBlockRef, len(blocks))
	for i := range values {
		v, b := nodes[i+1], nodes[len(values)+i+1]
		if err := operand(p, v); err != "" {
			return err
		}
		if b.kind != blockKind || b.module != p.module || b.owner != p.owner {
			return "argument: incoming block belongs to another function"
		}
		vs[i], bs[i] = C.LLVMValueRef(v.raw), C.LLVMBasicBlockRef(b.raw)
		if C.LLVMTypeOf(vs[i]) != C.LLVMTypeOf(raw) {
			return "argument: phi incoming type mismatch"
		}
		if seen[bs[i]] {
			return "argument: duplicate phi predecessor"
		}
		seen[bs[i]] = true
	}
	C.LLVMAddIncoming(raw, unsafe.SliceData(vs), unsafe.SliceData(bs), C.uint(len(vs)))
	return ""
}

func GEP(builder, source, pointer Handle, indices []Handle, inBounds bool, name string) (Handle, string) {
	if len(indices) == 0 {
		return nil, "argument: GEP requires at least one index"
	}
	ctx, nodes, err := enter(append([]Handle{builder, source, pointer}, indices...)...)
	if err != "" {
		return nil, err
	}
	defer ctx.mu.Unlock()
	b, t, p := nodes[0], nodes[1], nodes[2]
	if err := insertion(b); err != "" {
		return nil, err
	}
	if t.kind != typeKind {
		return nil, "argument: expected GEP source type"
	}
	ty := C.LLVMTypeRef(t.raw)
	if !firstClass(ty) || C.LLVMTypeIsSized(ty) == 0 {
		return nil, "argument: GEP requires sized source type"
	}
	if err := operand(b, p); err != "" {
		return nil, err
	}
	ptr := C.LLVMValueRef(p.raw)
	if C.LLVMGetTypeKind(C.LLVMTypeOf(ptr)) != C.LLVMPointerTypeKind {
		return nil, "argument: GEP requires pointer"
	}
	values := make([]C.LLVMValueRef, len(indices))
	current := ty
	for i, n := range nodes[3:] {
		if err := operand(b, n); err != "" {
			return nil, err
		}
		v := C.LLVMValueRef(n.raw)
		values[i] = v
		if C.LLVMGetTypeKind(C.LLVMTypeOf(v)) != C.LLVMIntegerTypeKind {
			return nil, "argument: GEP indices must be integers"
		}
		if i == 0 {
			continue
		}
		switch C.LLVMGetTypeKind(current) {
		case C.LLVMArrayTypeKind:
			current = C.LLVMGetElementType(current)
		case C.LLVMStructTypeKind:
			if C.LLVMIsAConstantInt(v) == nil || C.LLVMGetIntTypeWidth(C.LLVMTypeOf(v)) != 32 {
				return nil, "argument: struct GEP index must be constant i32"
			}
			field := C.LLVMConstIntGetZExtValue(v)
			if field >= C.ulonglong(C.LLVMCountStructElementTypes(current)) {
				return nil, "argument: struct GEP index out of bounds"
			}
			current = C.LLVMStructGetTypeAtIndex(current, C.uint(field))
		default:
			return nil, "argument: GEP index enters non-aggregate type"
		}
	}
	raw, free, err := text(name)
	if err != "" {
		return nil, err
	}
	defer free()
	var result C.LLVMValueRef
	if inBounds {
		result = C.LLVMBuildInBoundsGEP2(C.LLVMBuilderRef(b.raw), ty, ptr, unsafe.SliceData(values), C.uint(len(values)), raw)
	} else {
		result = C.LLVMBuildGEP2(C.LLVMBuilderRef(b.raw), ty, ptr, unsafe.SliceData(values), C.uint(len(values)), raw)
	}
	return instruction(b, result), ""
}

func Optimize(handle Handle, pipeline string) string {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return err
	}
	defer ctx.mu.Unlock()
	module := nodes[0]
	if module.kind != moduleKind {
		return "argument: expected module"
	}
	if err := verify(module); err != "" {
		return err
	}
	if pipeline == "" {
		return "argument: empty pass pipeline"
	}
	raw, free, err := text(pipeline)
	if err != "" {
		return err
	}
	defer free()
	for builder := range ctx.builders {
		if builder.module == module {
			C.LLVMClearInsertionPosition(C.LLVMBuilderRef(builder.raw))
			builder.module = nil
			builder.owner = nil
			builder.block = nil
		}
	}
	module.epoch++
	options := C.LLVMCreatePassBuilderOptions()
	defer C.LLVMDisposePassBuilderOptions(options)
	failure := C.LLVMRunPasses(C.LLVMModuleRef(module.raw), raw, nil, options)
	if failure != nil {
		diagnostic := C.LLVMGetErrorMessage(failure)
		defer C.LLVMDisposeErrorMessage(diagnostic)
		return "llvm: " + C.GoString(diagnostic)
	}
	return verify(module)
}

var targetOnce sync.Once
var targetError string

func EmitFile(handle Handle, path string, assembly bool, level int) string {
	ctx, nodes, err := enter(handle)
	if err != "" {
		return err
	}
	defer ctx.mu.Unlock()
	module := nodes[0]
	if module.kind != moduleKind {
		return "argument: expected module"
	}
	if level < 0 || level > 3 {
		return "argument: optimization level must be 0..3"
	}
	if path == "" {
		return "argument: output path is empty"
	}
	if err := verify(module); err != "" {
		return err
	}
	rawPath, free, err := text(path)
	if err != "" {
		return err
	}
	defer free()
	targetOnce.Do(func() {
		if C.initialize_native() != 0 {
			targetError = "llvm: native target initialization failed"
		}
	})
	if targetError != "" {
		return targetError
	}
	triple := C.LLVMGetDefaultTargetTriple()
	defer C.LLVMDisposeMessage(triple)
	cpu := C.LLVMGetHostCPUName()
	defer C.LLVMDisposeMessage(cpu)
	features := C.LLVMGetHostCPUFeatures()
	defer C.LLVMDisposeMessage(features)
	var target C.LLVMTargetRef
	var diagnostic *C.char
	if C.LLVMGetTargetFromTriple(triple, &target, &diagnostic) != 0 {
		return "llvm: " + message(diagnostic)
	}
	machine := C.LLVMCreateTargetMachine(target, triple, cpu, features, C.LLVMCodeGenOptLevel(level), C.LLVMRelocPIC, C.LLVMCodeModelDefault)
	if machine == nil {
		return "llvm: target machine creation failed"
	}
	defer C.LLVMDisposeTargetMachine(machine)
	clone := C.LLVMCloneModule(C.LLVMModuleRef(module.raw))
	defer C.LLVMDisposeModule(clone)
	C.LLVMSetTarget(clone, triple)
	layout := C.LLVMCreateTargetDataLayout(machine)
	defer C.LLVMDisposeTargetData(layout)
	C.LLVMSetModuleDataLayout(clone, layout)
	format := C.LLVMObjectFile
	if assembly {
		format = C.LLVMAssemblyFile
	}
	if C.LLVMTargetMachineEmitToFile(machine, clone, rawPath, C.LLVMCodeGenFileType(format), &diagnostic) != 0 {
		return "llvm: " + message(diagnostic)
	}
	return ""
}

func Version() string {
	var major, minor, patch C.unsigned
	C.LLVMGetVersion(&major, &minor, &patch)
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
