package native

import (
	"fmt"
	"goml.dev/cabi"
	"strings"
	"sync"
	"unsafe"
)

var gomlCLibrary = cabi.Library{Names: []string{"libLLVM-18.so"}}
var gomlCOnce sync.Once
var gomlCError error
var gomlCAddresses [16]uint64

func gomlCLoad() error {
	gomlCOnce.Do(func() {
		gomlCAddresses[0], gomlCError = gomlCLibrary.Symbol("LLVMAddFunction")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[1], gomlCError = gomlCLibrary.Symbol("LLVMAppendBasicBlockInContext")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[2], gomlCError = gomlCLibrary.Symbol("LLVMBuildRet")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[3], gomlCError = gomlCLibrary.Symbol("LLVMConstInt")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[4], gomlCError = gomlCLibrary.Symbol("LLVMContextCreate")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[5], gomlCError = gomlCLibrary.Symbol("LLVMContextDispose")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[6], gomlCError = gomlCLibrary.Symbol("LLVMCreateBuilderInContext")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[7], gomlCError = gomlCLibrary.Symbol("LLVMDisposeBuilder")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[8], gomlCError = gomlCLibrary.Symbol("LLVMDisposeMessage")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[9], gomlCError = gomlCLibrary.Symbol("LLVMDisposeModule")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[10], gomlCError = gomlCLibrary.Symbol("LLVMFunctionType")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[11], gomlCError = gomlCLibrary.Symbol("LLVMInt32TypeInContext")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[12], gomlCError = gomlCLibrary.Symbol("LLVMModuleCreateWithNameInContext")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[13], gomlCError = gomlCLibrary.Symbol("LLVMPositionBuilderAtEnd")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[14], gomlCError = gomlCLibrary.Symbol("LLVMPrintModuleToString")
		if gomlCError != nil {
			return
		}
		gomlCAddresses[15], gomlCError = gomlCLibrary.Symbol("LLVMVerifyModule")
		if gomlCError != nil {
			return
		}
	})
	return gomlCError
}
func gomlDynamic_LLVMAddFunction(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[0], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(uintptr(p1))
	f.Integer[2] = uint64(uintptr(p2))
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMAppendBasicBlockInContext(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[1], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(uintptr(p1))
	f.Integer[2] = uint64(uintptr(p2))
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMBuildRet(p0 unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[2], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(uintptr(p1))
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMConstInt(p0 unsafe.Pointer, p1 uint64, p2 int32) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[3], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(p1)
	f.Integer[2] = uint64(p2)
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMContextCreate() unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[4], StackCount: 0}

	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMContextDispose(p0 unsafe.Pointer) {
	f := cabi.Frame{Function: gomlCAddresses[5], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
}
func gomlDynamic_LLVMCreateBuilderInContext(p0 unsafe.Pointer) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[6], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMDisposeBuilder(p0 unsafe.Pointer) {
	f := cabi.Frame{Function: gomlCAddresses[7], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
}
func gomlDynamic_LLVMDisposeMessage(p0 unsafe.Pointer) {
	f := cabi.Frame{Function: gomlCAddresses[8], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
}
func gomlDynamic_LLVMDisposeModule(p0 unsafe.Pointer) {
	f := cabi.Frame{Function: gomlCAddresses[9], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
}
func gomlDynamic_LLVMFunctionType(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 uint32, p3 int32) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[10], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(uintptr(p1))
	f.Integer[2] = uint64(p2)
	f.Integer[3] = uint64(p3)
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMInt32TypeInContext(p0 unsafe.Pointer) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[11], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMModuleCreateWithNameInContext(p0 unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[12], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(uintptr(p1))
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMPositionBuilderAtEnd(p0 unsafe.Pointer, p1 unsafe.Pointer) {
	f := cabi.Frame{Function: gomlCAddresses[13], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(uintptr(p1))
	f.Call()
}
func gomlDynamic_LLVMPrintModuleToString(p0 unsafe.Pointer) unsafe.Pointer {
	f := cabi.Frame{Function: gomlCAddresses[14], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Call()
	return cabi.Pointer(f.Result)
}
func gomlDynamic_LLVMVerifyModule(p0 unsafe.Pointer, p1 uint32, p2 unsafe.Pointer) int32 {
	f := cabi.Frame{Function: gomlCAddresses[15], StackCount: 0}
	f.Integer[0] = uint64(uintptr(p0))
	f.Integer[1] = uint64(p1)
	f.Integer[2] = uint64(uintptr(p2))
	f.Call()
	return int32(f.Result)
}

type Context struct{ raw unsafe.Pointer }

func GomlC_NullContext() Context                  { return Context{} }
func GomlC_IsNullContext(value Context) bool      { return value.raw == nil }
func GomlC_EqualContext(left, right Context) bool { return left.raw == right.raw }

type Module struct{ raw unsafe.Pointer }

func GomlC_NullModule() Module                  { return Module{} }
func GomlC_IsNullModule(value Module) bool      { return value.raw == nil }
func GomlC_EqualModule(left, right Module) bool { return left.raw == right.raw }

type Type struct{ raw unsafe.Pointer }

func GomlC_NullType() Type                  { return Type{} }
func GomlC_IsNullType(value Type) bool      { return value.raw == nil }
func GomlC_EqualType(left, right Type) bool { return left.raw == right.raw }

type TypeArray struct{ raw unsafe.Pointer }

func GomlC_NullTypeArray() TypeArray                  { return TypeArray{} }
func GomlC_IsNullTypeArray(value TypeArray) bool      { return value.raw == nil }
func GomlC_EqualTypeArray(left, right TypeArray) bool { return left.raw == right.raw }

type Value struct{ raw unsafe.Pointer }

func GomlC_NullValue() Value                  { return Value{} }
func GomlC_IsNullValue(value Value) bool      { return value.raw == nil }
func GomlC_EqualValue(left, right Value) bool { return left.raw == right.raw }

type Block struct{ raw unsafe.Pointer }

func GomlC_NullBlock() Block                  { return Block{} }
func GomlC_IsNullBlock(value Block) bool      { return value.raw == nil }
func GomlC_EqualBlock(left, right Block) bool { return left.raw == right.raw }

type Builder struct{ raw unsafe.Pointer }

func GomlC_NullBuilder() Builder                  { return Builder{} }
func GomlC_IsNullBuilder(value Builder) bool      { return value.raw == nil }
func GomlC_EqualBuilder(left, right Builder) bool { return left.raw == right.raw }
func GomlC_context_create() (Context, error) {
	if err := gomlCLoad(); err != nil {
		return Context{}, err
	}
	cResult := gomlDynamic_LLVMContextCreate()
	return Context{raw: cResult}, nil
}
func GomlC_context_dispose(a0 Context) error {
	if err := gomlCLoad(); err != nil {
		return err
	}
	gomlDynamic_LLVMContextDispose(a0.raw)
	return nil
}
func GomlC_module_create(a0 string, a1 Context) (Module, error) {
	if err := gomlCLoad(); err != nil {
		return Module{}, err
	}
	if strings.IndexByte(a0, 0) >= 0 || len(a0) > 67108864 {
		return Module{}, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg0, _ := cabi.Copy([]byte(a0), true)
	if cArg0 == nil {
		return Module{}, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg0)
	cResult := gomlDynamic_LLVMModuleCreateWithNameInContext(cArg0, a1.raw)
	return Module{raw: cResult}, nil
}
func GomlC_module_dispose(a0 Module) error {
	if err := gomlCLoad(); err != nil {
		return err
	}
	gomlDynamic_LLVMDisposeModule(a0.raw)
	return nil
}
func GomlC_int32_type(a0 Context) (Type, error) {
	if err := gomlCLoad(); err != nil {
		return Type{}, err
	}
	cResult := gomlDynamic_LLVMInt32TypeInContext(a0.raw)
	return Type{raw: cResult}, nil
}
func GomlC_function_type(a0 Type, a1 TypeArray, a2 uint32, a3 int32) (Type, error) {
	if err := gomlCLoad(); err != nil {
		return Type{}, err
	}
	cResult := gomlDynamic_LLVMFunctionType(a0.raw, a1.raw, uint32(a2), int32(a3))
	return Type{raw: cResult}, nil
}
func GomlC_add_function(a0 Module, a1 string, a2 Type) (Value, error) {
	if err := gomlCLoad(); err != nil {
		return Value{}, err
	}
	if strings.IndexByte(a1, 0) >= 0 || len(a1) > 67108864 {
		return Value{}, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg1, _ := cabi.Copy([]byte(a1), true)
	if cArg1 == nil {
		return Value{}, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg1)
	cResult := gomlDynamic_LLVMAddFunction(a0.raw, cArg1, a2.raw)
	return Value{raw: cResult}, nil
}
func GomlC_append_block(a0 Context, a1 Value, a2 string) (Block, error) {
	if err := gomlCLoad(); err != nil {
		return Block{}, err
	}
	if strings.IndexByte(a2, 0) >= 0 || len(a2) > 67108864 {
		return Block{}, fmt.Errorf("C string contains NUL or exceeds 64 MiB")
	}
	cArg2, _ := cabi.Copy([]byte(a2), true)
	if cArg2 == nil {
		return Block{}, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg2)
	cResult := gomlDynamic_LLVMAppendBasicBlockInContext(a0.raw, a1.raw, cArg2)
	return Block{raw: cResult}, nil
}
func GomlC_builder_create(a0 Context) (Builder, error) {
	if err := gomlCLoad(); err != nil {
		return Builder{}, err
	}
	cResult := gomlDynamic_LLVMCreateBuilderInContext(a0.raw)
	return Builder{raw: cResult}, nil
}
func GomlC_builder_dispose(a0 Builder) error {
	if err := gomlCLoad(); err != nil {
		return err
	}
	gomlDynamic_LLVMDisposeBuilder(a0.raw)
	return nil
}
func GomlC_position_at_end(a0 Builder, a1 Block) error {
	if err := gomlCLoad(); err != nil {
		return err
	}
	gomlDynamic_LLVMPositionBuilderAtEnd(a0.raw, a1.raw)
	return nil
}
func GomlC_constant_int(a0 Type, a1 uint64, a2 int32) (Value, error) {
	if err := gomlCLoad(); err != nil {
		return Value{}, err
	}
	cResult := gomlDynamic_LLVMConstInt(a0.raw, uint64(a1), int32(a2))
	return Value{raw: cResult}, nil
}
func GomlC_build_return(a0 Builder, a1 Value) (Value, error) {
	if err := gomlCLoad(); err != nil {
		return Value{}, err
	}
	cResult := gomlDynamic_LLVMBuildRet(a0.raw, a1.raw)
	return Value{raw: cResult}, nil
}
func GomlC_verify(a0 Module, a1 int64) (int32, string, bool, error) {
	if err := gomlCLoad(); err != nil {
		return 0, "", false, err
	}
	if int64(uint32(a1)) != a1 {
		return 0, "", false, fmt.Errorf("C enum argument is out of range")
	}
	cArg2Slot, _ := cabi.Alloc(8)
	if cArg2Slot == nil {
		return 0, "", false, fmt.Errorf("C allocation failed")
	}
	defer cabi.Free(cArg2Slot)
	cResult := gomlDynamic_LLVMVerifyModule(a0.raw, uint32(a1), cArg2Slot)
	cArg2 := *(*unsafe.Pointer)(cArg2Slot)
	defer gomlDynamic_LLVMDisposeMessage(cArg2)
	cArg2Text, cArg2Error := cabi.String(cArg2, 1048576)
	if cArg2Error != nil {
		return 0, "", false, fmt.Errorf("C string exceeds configured copy limit")
	}
	return int32(cResult), cArg2Text, cArg2 != nil, nil
}
func GomlC_print_module(a0 Module) (string, bool, error) {
	if err := gomlCLoad(); err != nil {
		return "", false, err
	}
	cResult := gomlDynamic_LLVMPrintModuleToString(a0.raw)
	defer gomlDynamic_LLVMDisposeMessage(cResult)
	cResultText, cResultError := cabi.String(cResult, 1048576)
	if cResultError != nil {
		return "", false, fmt.Errorf("C string exceeds configured copy limit")
	}
	return cResultText, cResult != nil, nil
}
func GomlC_RETURN_STATUS() int32 { return 2 }
