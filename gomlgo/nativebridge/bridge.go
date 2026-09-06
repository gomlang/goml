package nativebridge

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"sync"
	"syscall"
)

func LockImageBuild(path string) (int, string) {
	file, err := syscall.Open(path, syscall.O_CREAT|syscall.O_RDWR|syscall.O_CLOEXEC, 0o600)
	if err != nil {
		return -1, err.Error()
	}
	for {
		err = syscall.Flock(file, syscall.LOCK_EX)
		if err != syscall.EINTR {
			break
		}
	}
	if err != nil {
		syscall.Close(file)
		return -1, err.Error()
	}
	return file, ""
}

func UnlockImageBuild(file int) {
	syscall.Close(file)
}

func ReplaceProcess(program string, arguments []string, environmentKey string, environmentValue string) string {
	previous, present := os.LookupEnv(environmentKey)
	if err := os.Setenv(environmentKey, environmentValue); err != nil {
		return err.Error()
	}
	err := syscall.Exec(program, arguments, os.Environ())
	if present {
		_ = os.Setenv(environmentKey, previous)
	} else {
		_ = os.Unsetenv(environmentKey)
	}
	return err.Error()
}

type CallID uint32

type ValueRef = reflect.Value

type Binding struct {
	id       CallID
	function reflect.Value
}

type Registry struct {
	functions map[CallID]reflect.Value
	names     map[CallID]string
}

type Invocation struct {
	id        CallID
	function  ValueRef
	slice     bool
	arguments []ValueRef
	result    CallResult
	done      chan struct{}
	start     sync.Once
}

type CallResult struct {
	Values     []ValueRef
	PanicValue any
	Err        error
	Panicked   bool
	Exited     bool
	ExitCode   int
}

type exitSignal struct {
	code int
}

type nativeNil struct{}

var defaultRegistry struct {
	sync.RWMutex
	value *Registry
}

var nativeGlobals struct {
	sync.RWMutex
	values map[int]reflect.Value
}

func RegisterNativeGlobal(id int, value any) error {
	if id < 0 {
		return fmt.Errorf("native global ID is negative")
	}
	ref := reflect.ValueOf(value)
	if !ref.IsValid() {
		return fmt.Errorf("native global %d is invalid", id)
	}
	nativeGlobals.Lock()
	defer nativeGlobals.Unlock()
	if nativeGlobals.values == nil {
		nativeGlobals.values = make(map[int]reflect.Value)
	}
	if current, exists := nativeGlobals.values[id]; exists && current.Type() != ref.Type() {
		return fmt.Errorf("native global %d is registered with conflicting types", id)
	}
	nativeGlobals.values[id] = ref
	return nil
}

func NativeGlobalValue(id int) any {
	nativeGlobals.RLock()
	value := nativeGlobals.values[id]
	nativeGlobals.RUnlock()
	return Clone(value)
}

func ValueOf(value any) ValueRef {
	return reflect.ValueOf(value)
}

func CloneValue(value any) any {
	ref, ok := value.(reflect.Value)
	if !ok {
		return reflect.Value{}
	}
	return Clone(ref)
}

func ValidValue(value any) bool {
	ref, ok := value.(reflect.Value)
	return ok && ref.IsValid()
}

func NilValue(value any) bool {
	ref, ok := value.(reflect.Value)
	if !ok || !ref.IsValid() {
		return true
	}
	switch ref.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return ref.IsNil()
	default:
		return false
	}
}

func ValueType(value any) string {
	ref, ok := value.(reflect.Value)
	if !ok || !ref.IsValid() {
		return ""
	}
	return ref.Type().String()
}

func ValueText(value any) string {
	ref, ok := value.(reflect.Value)
	if !ok || !ref.IsValid() {
		return "<nil>"
	}
	return fmt.Sprint(ref.Interface())
}

func Values(values ...any) []ValueRef {
	result := make([]ValueRef, len(values))
	for index, value := range values {
		result[index] = ValueOf(value)
	}
	return result
}

func Clone(value ValueRef) ValueRef {
	if !value.IsValid() {
		return value
	}
	cloned := reflect.New(value.Type()).Elem()
	cloned.Set(value)
	return cloned
}

func NewNilValue() any {
	return reflect.ValueOf(nativeNil{})
}

func ProgramArgsValue() any {
	arguments := []string{os.Args[0]}
	for index, value := range os.Args {
		if value == "--" {
			arguments = append(arguments, os.Args[index+1:]...)
			break
		}
	}
	return reflect.ValueOf(arguments)
}

func Bind(id CallID, function any) Binding {
	return Binding{id: id, function: reflect.ValueOf(function)}
}

func BindValue(id CallID, function ValueRef) Binding {
	return Binding{id: id, function: function}
}

func ExitBinding(id CallID) Binding {
	return Bind(id, exitAdapter)
}

func Install(bindings ...Binding) (*Registry, error) {
	functions := make(map[CallID]reflect.Value, len(bindings))
	for _, binding := range bindings {
		if !binding.function.IsValid() {
			return nil, fmt.Errorf("native call %d has no function", binding.id)
		}
		if binding.function.Kind() != reflect.Func {
			return nil, fmt.Errorf("native call %d is %s, not func", binding.id, binding.function.Kind())
		}
		if binding.function.IsNil() {
			return nil, fmt.Errorf("native call %d has a nil function", binding.id)
		}
		if _, exists := functions[binding.id]; exists {
			return nil, fmt.Errorf("native call %d is registered more than once", binding.id)
		}
		functions[binding.id] = binding.function
	}
	return &Registry{functions: functions}, nil
}

func InstallDefault(bindings ...Binding) (*Registry, error) {
	registry, err := Install(bindings...)
	if err != nil {
		return nil, err
	}
	defaultRegistry.Lock()
	defaultRegistry.value = registry
	defaultRegistry.Unlock()
	return registry, nil
}

func DefaultRegistry() *Registry {
	defaultRegistry.RLock()
	registry := defaultRegistry.value
	defaultRegistry.RUnlock()
	return registry
}

func DefaultCall(id CallID, arguments []ValueRef) CallResult {
	return DefaultRegistry().Call(id, arguments)
}

func DefaultCallSlice(id CallID, arguments []ValueRef) CallResult {
	return DefaultRegistry().CallSlice(id, arguments)
}

func NewInvocation(id int, slice bool) any {
	return reflect.ValueOf(&Invocation{id: CallID(id), slice: slice, done: make(chan struct{})})
}

func NewValueInvocation(function any, slice bool) any {
	ref, ok := function.(reflect.Value)
	if !ok {
		ref = reflect.Value{}
	}
	return reflect.ValueOf(&Invocation{function: ref, slice: slice, done: make(chan struct{})})
}

func MethodValue(receiver any, name string) any {
	value, ok := receiver.(reflect.Value)
	if !ok || !value.IsValid() {
		return reflect.Value{}
	}
	return value.MethodByName(name)
}

func InvocationAppend(invocation any, value any) {
	call, err := invocationValue(invocation)
	if err != nil {
		return
	}
	ref, ok := value.(reflect.Value)
	if !ok {
		call.result.Err = fmt.Errorf("native call %s argument %d is not a reflected value", nativeCallName(call), len(call.arguments))
		return
	}
	if !ref.IsValid() {
		call.result.Err = fmt.Errorf("native call %s argument %d is invalid", nativeCallName(call), len(call.arguments))
		return
	}
	call.arguments = append(call.arguments, ref)
}

func nativeCallName(call *Invocation) string {
	if call.function.IsValid() {
		return call.function.Type().String()
	}
	registry := DefaultRegistry()
	if registry != nil {
		if name := registry.names[call.id]; name != "" {
			return fmt.Sprintf("%d (%s)", call.id, name)
		}
	}
	return fmt.Sprint(call.id)
}

func InvocationRun(invocation any) {
	call, err := invocationValue(invocation)
	if err != nil {
		return
	}
	if call.result.Err != nil {
		return
	}
	function := call.function
	if !function.IsValid() {
		registry := DefaultRegistry()
		if registry == nil {
			call.result.Err = fmt.Errorf("native registry is not installed")
			return
		}
		function, err = registry.lookup(call.id)
		if err != nil {
			call.result.Err = err
			return
		}
	}
	arguments, err := convertedArguments(function.Type(), call.arguments, call.slice)
	if err != nil {
		call.result.Err = fmt.Errorf("native call %d: %w", call.id, err)
		return
	}
	call.result = invoke(function, arguments, call.slice)
}

func InvocationUseRuntimeStack(invocation any, goroutine uint64) bool {
	call, err := invocationValue(invocation)
	if err != nil {
		return false
	}
	call.function = reflect.ValueOf(func(buffer []byte, all bool) int {
		if all {
			return runtime.Stack(buffer, all)
		}
		return copy(buffer, fmt.Sprintf("goroutine %d [running]:\n", goroutine))
	})
	return true
}

func InvocationStart(invocation any) bool {
	call, err := invocationValue(invocation)
	if err != nil {
		return false
	}
	started := false
	call.start.Do(func() {
		started = true
		go func() {
			InvocationRun(reflect.ValueOf(call))
			close(call.done)
		}()
	})
	return started
}

func InvocationReady(invocation any) bool {
	call, err := invocationValue(invocation)
	if err != nil {
		return false
	}
	select {
	case <-call.done:
		return true
	default:
		return false
	}
}

func InvocationWait(invocation any) {
	call, err := invocationValue(invocation)
	if err != nil {
		return
	}
	<-call.done
}

func InvocationError(invocation any) string {
	call, err := invocationValue(invocation)
	if err != nil {
		return err.Error()
	}
	if call.result.Err == nil {
		return ""
	}
	return call.result.Err.Error()
}

func InvocationPanicked(invocation any) bool {
	call, err := invocationValue(invocation)
	return err == nil && call.result.Panicked
}

func InvocationExited(invocation any) bool {
	call, err := invocationValue(invocation)
	return err == nil && call.result.Exited
}

func InvocationExitCode(invocation any) int {
	call, err := invocationValue(invocation)
	if err != nil {
		return 0
	}
	return call.result.ExitCode
}

func InvocationResultCount(invocation any) int {
	call, err := invocationValue(invocation)
	if err != nil {
		return 0
	}
	return len(call.result.Values)
}

func InvocationResult(invocation any, index int) any {
	call, err := invocationValue(invocation)
	if err != nil || index < 0 || index >= len(call.result.Values) {
		return reflect.Value{}
	}
	return call.result.Values[index]
}

func InvocationArgument(invocation any, index int) any {
	call, err := invocationValue(invocation)
	if err != nil || index < 0 || index >= len(call.arguments) {
		return reflect.Value{}
	}
	return call.arguments[index]
}

func InvocationPanic(invocation any) any {
	call, err := invocationValue(invocation)
	if err != nil || !call.result.Panicked {
		return reflect.Value{}
	}
	return reflect.ValueOf(call.result.PanicValue)
}

func NewIntegerValue(typeKey string, bits uint64) any {
	value, err := newValue(typeKey)
	if err != nil {
		return reflect.Value{}
	}
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(int64(bits))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value.SetUint(bits)
	default:
		return reflect.Value{}
	}
	return value
}

func NewBoolValue(value bool) any {
	return reflect.ValueOf(value)
}

func NewFloatValue(typeKey string, bits uint64) any {
	value, err := newValue(typeKey)
	if err != nil {
		return reflect.Value{}
	}
	if value.Kind() == reflect.Float32 {
		value.SetFloat(float64(math.Float32frombits(uint32(bits))))
	} else if value.Kind() == reflect.Float64 {
		value.SetFloat(math.Float64frombits(bits))
	} else {
		return reflect.Value{}
	}
	return value
}

func NewComplexValue(typeKey string, realBits uint64, imaginaryBits uint64) any {
	value, err := newValue(typeKey)
	if err != nil {
		return reflect.Value{}
	}
	if value.Kind() == reflect.Complex64 {
		value.SetComplex(complex(
			float64(math.Float32frombits(uint32(realBits))),
			float64(math.Float32frombits(uint32(imaginaryBits))),
		))
	} else if value.Kind() == reflect.Complex128 {
		value.SetComplex(complex(math.Float64frombits(realBits), math.Float64frombits(imaginaryBits)))
	} else {
		return reflect.Value{}
	}
	return value
}

func NewStringValue(value string) any {
	return reflect.ValueOf(value)
}

func NewSliceValue(typeKey string, length int, nilValue bool) any {
	typeValue, err := typeForKey(typeKey)
	if err != nil || typeValue.Kind() != reflect.Slice || length < 0 {
		return reflect.Value{}
	}
	if nilValue {
		return reflect.Zero(typeValue)
	}
	return reflect.MakeSlice(typeValue, length, length)
}

func SetSliceValue(slice any, index int, element any) bool {
	sliceValue, ok := slice.(reflect.Value)
	if !ok || !sliceValue.IsValid() || sliceValue.Kind() != reflect.Slice || index < 0 || index >= sliceValue.Len() {
		return false
	}
	elementValue, ok := element.(reflect.Value)
	if !ok || !elementValue.IsValid() {
		return false
	}
	converted, err := convertedValue(elementValue, sliceValue.Type().Elem(), index)
	if err != nil {
		return false
	}
	sliceValue.Index(index).Set(converted)
	return true
}

func ValueBool(value any) bool {
	ref, ok := reflectedValue(value)
	return ok && ref.Kind() == reflect.Bool && ref.Bool()
}

func ValueIntegerBits(value any) uint64 {
	ref, ok := reflectedValue(value)
	if !ok {
		return 0
	}
	switch ref.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(ref.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return ref.Uint()
	default:
		return 0
	}
}

func ValueFloatBits(value any) uint64 {
	ref, ok := reflectedValue(value)
	if !ok {
		return 0
	}
	if ref.Kind() == reflect.Float32 {
		return uint64(math.Float32bits(float32(ref.Float())))
	}
	if ref.Kind() == reflect.Float64 {
		return math.Float64bits(ref.Float())
	}
	return 0
}

func ValueComplexRealBits(value any) uint64 {
	ref, ok := reflectedValue(value)
	if !ok {
		return 0
	}
	if ref.Kind() == reflect.Complex64 {
		return uint64(math.Float32bits(float32(real(ref.Complex()))))
	}
	if ref.Kind() == reflect.Complex128 {
		return math.Float64bits(real(ref.Complex()))
	}
	return 0
}

func ValueComplexImaginaryBits(value any) uint64 {
	ref, ok := reflectedValue(value)
	if !ok {
		return 0
	}
	if ref.Kind() == reflect.Complex64 {
		return uint64(math.Float32bits(float32(imag(ref.Complex()))))
	}
	if ref.Kind() == reflect.Complex128 {
		return math.Float64bits(imag(ref.Complex()))
	}
	return 0
}

func ValueStringText(value any) string {
	ref, ok := reflectedValue(value)
	if !ok || ref.Kind() != reflect.String {
		return ""
	}
	return ref.String()
}

func ValueLength(value any) int {
	ref, ok := reflectedValue(value)
	if !ok || ref.Kind() != reflect.Slice {
		return 0
	}
	return ref.Len()
}

func ValueCapacity(value any) int {
	ref, ok := reflectedValue(value)
	if !ok || ref.Kind() != reflect.Slice {
		return 0
	}
	return ref.Cap()
}

func ValueIndex(value any, index int) any {
	ref, ok := reflectedValue(value)
	if !ok || ref.Kind() != reflect.Slice || index < 0 || index >= ref.Len() {
		return reflect.Value{}
	}
	return Clone(ref.Index(index))
}

func (registry *Registry) Call(id CallID, arguments []ValueRef) CallResult {
	function, err := registry.lookup(id)
	if err != nil {
		return CallResult{Err: err}
	}
	if err := validateCall(function.Type(), arguments, false); err != nil {
		return CallResult{Err: fmt.Errorf("native call %d: %w", id, err)}
	}
	return invoke(function, arguments, false)
}

func (registry *Registry) CallSlice(id CallID, arguments []ValueRef) CallResult {
	function, err := registry.lookup(id)
	if err != nil {
		return CallResult{Err: err}
	}
	if err := validateCall(function.Type(), arguments, true); err != nil {
		return CallResult{Err: fmt.Errorf("native call %d: %w", id, err)}
	}
	return invoke(function, arguments, true)
}

func CallMethod(receiver ValueRef, name string, arguments []ValueRef) CallResult {
	method, err := lookupMethod(receiver, name)
	if err != nil {
		return CallResult{Err: err}
	}
	if err := validateCall(method.Type(), arguments, false); err != nil {
		return CallResult{Err: fmt.Errorf("native method %s: %w", name, err)}
	}
	return invoke(method, arguments, false)
}

func CallMethodSlice(receiver ValueRef, name string, arguments []ValueRef) CallResult {
	method, err := lookupMethod(receiver, name)
	if err != nil {
		return CallResult{Err: err}
	}
	if err := validateCall(method.Type(), arguments, true); err != nil {
		return CallResult{Err: fmt.Errorf("native method %s: %w", name, err)}
	}
	return invoke(method, arguments, true)
}

func (registry *Registry) lookup(id CallID) (reflect.Value, error) {
	if registry == nil {
		return reflect.Value{}, fmt.Errorf("native registry is nil")
	}
	function, exists := registry.functions[id]
	if !exists {
		return reflect.Value{}, fmt.Errorf("native call %d is not registered", id)
	}
	return function, nil
}

func lookupMethod(receiver ValueRef, name string) (reflect.Value, error) {
	if !receiver.IsValid() {
		return reflect.Value{}, fmt.Errorf("native method %s has an invalid receiver", name)
	}
	method := receiver.MethodByName(name)
	if !method.IsValid() {
		return reflect.Value{}, fmt.Errorf("native method %s is not available on %s", name, receiver.Type())
	}
	return method, nil
}

func invocationValue(value any) (*Invocation, error) {
	ref, ok := value.(reflect.Value)
	if !ok || !ref.IsValid() || ref.Kind() != reflect.Pointer || ref.IsNil() {
		return nil, fmt.Errorf("native invocation is invalid")
	}
	invocation, ok := ref.Interface().(*Invocation)
	if !ok {
		return nil, fmt.Errorf("native invocation has type %s", ref.Type())
	}
	return invocation, nil
}

func reflectedValue(value any) (reflect.Value, bool) {
	ref, ok := value.(reflect.Value)
	return ref, ok && ref.IsValid()
}

func newValue(typeKey string) (reflect.Value, error) {
	typeValue, err := typeForKey(typeKey)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.New(typeValue).Elem(), nil
}

func typeForKey(typeKey string) (reflect.Type, error) {
	if len(typeKey) >= 2 && typeKey[:2] == "[]" {
		element, err := typeForKey(typeKey[2:])
		if err != nil {
			return nil, err
		}
		return reflect.SliceOf(element), nil
	}
	switch typeKey {
	case "bool":
		return reflect.TypeFor[bool](), nil
	case "int":
		return reflect.TypeFor[int](), nil
	case "int8":
		return reflect.TypeFor[int8](), nil
	case "int16":
		return reflect.TypeFor[int16](), nil
	case "int32":
		return reflect.TypeFor[int32](), nil
	case "int64":
		return reflect.TypeFor[int64](), nil
	case "uint":
		return reflect.TypeFor[uint](), nil
	case "uint8":
		return reflect.TypeFor[uint8](), nil
	case "uint16":
		return reflect.TypeFor[uint16](), nil
	case "uint32":
		return reflect.TypeFor[uint32](), nil
	case "uint64":
		return reflect.TypeFor[uint64](), nil
	case "uintptr":
		return reflect.TypeFor[uintptr](), nil
	case "float32":
		return reflect.TypeFor[float32](), nil
	case "float64":
		return reflect.TypeFor[float64](), nil
	case "complex64":
		return reflect.TypeFor[complex64](), nil
	case "complex128":
		return reflect.TypeFor[complex128](), nil
	case "string":
		return reflect.TypeFor[string](), nil
	case "interface":
		return reflect.TypeFor[any](), nil
	case "error":
		return reflect.TypeFor[error](), nil
	default:
		return nil, fmt.Errorf("unsupported native type %s", typeKey)
	}
}

func convertedValue(value reflect.Value, target reflect.Type, index int) (reflect.Value, error) {
	if value.IsValid() && value.Type() == reflect.TypeFor[nativeNil]() && nilableType(target) {
		return reflect.Zero(target), nil
	}
	if value.IsValid() && value.Type() == reflect.TypeFor[*CallbackToken]() && target.Kind() == reflect.Func {
		return MakeCallback(target, value.Interface().(*CallbackToken))
	}
	if value.IsValid() && value.Type() == reflect.TypeFor[*CallbackToken]() && target.Kind() == reflect.Interface {
		return MakeInterfaceProxy(target, value.Interface().(*CallbackToken))
	}
	if value.Type().AssignableTo(target) {
		return value, nil
	}
	if value.Type().ConvertibleTo(target) {
		return value.Convert(target), nil
	}
	return reflect.Value{}, fmt.Errorf("argument %d has type %s, want %s", index, value.Type(), target)
}

func nilableType(value reflect.Type) bool {
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return true
	default:
		return false
	}
}

func convertedArguments(function reflect.Type, arguments []ValueRef, slice bool) ([]ValueRef, error) {
	if slice {
		if !function.IsVariadic() {
			return nil, fmt.Errorf("CallSlice requires a variadic function")
		}
		if len(arguments) != function.NumIn() {
			return nil, fmt.Errorf("got %d arguments, want %d", len(arguments), function.NumIn())
		}
	} else if !function.IsVariadic() && len(arguments) != function.NumIn() {
		return nil, fmt.Errorf("got %d arguments, want %d", len(arguments), function.NumIn())
	} else if function.IsVariadic() && len(arguments) < function.NumIn()-1 {
		return nil, fmt.Errorf("got %d arguments, want at least %d", len(arguments), function.NumIn()-1)
	}
	result := make([]ValueRef, len(arguments))
	for index, argument := range arguments {
		if !argument.IsValid() {
			return nil, fmt.Errorf("argument %d is invalid", index)
		}
		target := invocationArgumentType(function, index, slice)
		converted, err := convertedValue(argument, target, index)
		if err != nil {
			return nil, err
		}
		result[index] = converted
	}
	return result, nil
}

func invocationArgumentType(function reflect.Type, index int, slice bool) reflect.Type {
	if !slice && function.IsVariadic() && index >= function.NumIn()-1 {
		return function.In(function.NumIn() - 1).Elem()
	}
	return function.In(index)
}

func validateCall(function reflect.Type, arguments []ValueRef, slice bool) error {
	if slice {
		return validateSliceCall(function, arguments)
	}
	if !function.IsVariadic() {
		if len(arguments) != function.NumIn() {
			return fmt.Errorf("got %d arguments, want %d", len(arguments), function.NumIn())
		}
		for index, argument := range arguments {
			if err := validateArgument(argument, function.In(index), index); err != nil {
				return err
			}
		}
		return nil
	}
	minimum := function.NumIn() - 1
	if len(arguments) < minimum {
		return fmt.Errorf("got %d arguments, want at least %d", len(arguments), minimum)
	}
	for index := 0; index < minimum; index++ {
		if err := validateArgument(arguments[index], function.In(index), index); err != nil {
			return err
		}
	}
	element := function.In(function.NumIn() - 1).Elem()
	for index := minimum; index < len(arguments); index++ {
		if err := validateArgument(arguments[index], element, index); err != nil {
			return err
		}
	}
	return nil
}

func validateSliceCall(function reflect.Type, arguments []ValueRef) error {
	if !function.IsVariadic() {
		return fmt.Errorf("CallSlice requires a variadic function")
	}
	if len(arguments) != function.NumIn() {
		return fmt.Errorf("got %d arguments, want %d", len(arguments), function.NumIn())
	}
	for index, argument := range arguments {
		if err := validateArgument(argument, function.In(index), index); err != nil {
			return err
		}
	}
	return nil
}

func validateArgument(argument ValueRef, parameter reflect.Type, index int) error {
	if !argument.IsValid() {
		return fmt.Errorf("argument %d is invalid", index)
	}
	if !argument.Type().AssignableTo(parameter) {
		return fmt.Errorf("argument %d has type %s, want %s", index, argument.Type(), parameter)
	}
	return nil
}

func invoke(function reflect.Value, arguments []ValueRef, slice bool) (result CallResult) {
	completed := false
	defer func() {
		if completed {
			return
		}
		recovered := recover()
		if signal, ok := recovered.(exitSignal); ok {
			result.Exited = true
			result.ExitCode = signal.code
			return
		}
		result.Panicked = true
		result.PanicValue = recovered
	}()
	if slice {
		result.Values = function.CallSlice(arguments)
	} else {
		result.Values = function.Call(arguments)
	}
	completed = true
	return result
}

func exitAdapter(code int) {
	panic(exitSignal{code: code})
}
