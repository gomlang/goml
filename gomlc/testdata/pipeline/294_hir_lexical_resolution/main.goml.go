package main

import (
    _goml_fmt "fmt"
    _goml_debug "runtime/debug"
    _goml_context "context"
    _goml_os "os"
    _goml_sync "sync"
)

type _goml_panic_info struct {
    value any
    message string
    stack string
}

func (info *_goml_panic_info) Error() string {
    return info.message
}

func (info *_goml_panic_info) GomlPanicValue() any {
    return info.value
}

func (info *_goml_panic_info) GomlPanicMessage() string {
    return info.message
}

func (info *_goml_panic_info) GomlPanicStack() string {
    return info.stack
}

type _goml_panic_payload interface {
    GomlPanicValue() any
    GomlPanicMessage() string
    GomlPanicStack() string
}

type _goml_task_scope_state struct {
    mu _goml_sync.Mutex
    wg _goml_sync.WaitGroup
    state int
    ctx _goml_context.Context
    cancel _goml_context.CancelFunc
    panicked bool
    panic_value any
}

var _goml_task_registry_mutex _goml_sync.Mutex = _goml_sync.Mutex{}

var _goml_task_registry map[int64]*_goml_task_scope_state = make(map[int64]*_goml_task_scope_state)

var _goml_task_next_id int64 = 1

func _goml_task_lookup(id int64) *_goml_task_scope_state {
    _goml_task_registry_mutex.Lock()
    var scope *_goml_task_scope_state = _goml_task_registry[id]
    _goml_task_registry_mutex.Unlock()
    return scope
}

func _goml_task_record_panic(scope *_goml_task_scope_state, value any) {
    scope.mu.Lock()
    if !scope.panicked {
        scope.panicked = true
        scope.panic_value = value
    }
    scope.mu.Unlock()
    scope.cancel()
}

func _goml_task_run(scope *_goml_task_scope_state, body func() struct{}) {
    defer func() {
        var recovered any = recover()
        if recovered != nil {
            _goml_task_record_panic(scope, recovered)
        }
        scope.wg.Done()
    }()
    body()
}

func _goml_runtime_std_panic_raise(message string) struct{} {
    panic(message)
}

func _goml_runtime_std_panic_catch(body func() struct{}, caught func(string, string, func() struct{}) struct{}) struct{} {
    defer func() {
        var recovered any = recover()
        if recovered != nil {
            var info *_goml_panic_info
            switch value := recovered.(type) {
            case *_goml_panic_info:
                info = value
            case _goml_panic_payload:
                info = &_goml_panic_info{
                    value: value.GomlPanicValue(),
                    message: string([]rune(value.GomlPanicMessage())),
                    stack: string([]rune(value.GomlPanicStack())),
                }
            default:
                info = &_goml_panic_info{
                    value: recovered,
                    message: string([]rune(_goml_fmt.Sprint(recovered))),
                    stack: string([]rune(string(_goml_debug.Stack()))),
                }
            }
            caught(info.message, info.stack, func() struct{} {
                panic(info)
            })
        }
    }()
    body()
    return struct{}{}
}

func _goml_runtime_core_bool_to_string(x bool) string {
    if x {
        return "true"
    } else {
        return "false"
    }
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

func _goml_runtime_std_task_scope_new(parent int64) int64 {
    var parent_context _goml_context.Context = _goml_context.Background()
    var parent_missing bool = false
    if parent != 0 {
        var parent_scope *_goml_task_scope_state = _goml_task_lookup(parent)
        if parent_scope != nil {
            parent_context = parent_scope.ctx
        } else {
            parent_missing = true
        }
    }
    var ctx _goml_context.Context
    var cancel _goml_context.CancelFunc
    ctx, cancel = _goml_context.WithCancel(parent_context)
    if parent_missing {
        cancel()
    }
    _goml_task_registry_mutex.Lock()
    var id int64 = _goml_task_next_id
    _goml_task_next_id = _goml_task_next_id + 1
    var scope *_goml_task_scope_state = &_goml_task_scope_state{
        state: 0,
        ctx: ctx,
        cancel: cancel,
    }
    _goml_task_registry[id] = scope
    _goml_task_registry_mutex.Unlock()
    return id
}

func _goml_runtime_std_task_scope_run(scope_id int64, body func() struct{}) struct{} {
    var scope *_goml_task_scope_state = _goml_task_lookup(scope_id)
    func() {
        defer func() {
            var recovered any = recover()
            if recovered != nil {
                _goml_task_record_panic(scope, recovered)
            }
        }()
        body()
    }()
    return struct{}{}
}

func _goml_runtime_std_task_scope_spawn(scope_id int64, body func() struct{}) bool {
    var scope *_goml_task_scope_state = _goml_task_lookup(scope_id)
    if scope == nil {
        panic("task scope is closed")
    }
    scope.mu.Lock()
    if scope.state != 0 {
        scope.mu.Unlock()
        panic("task scope is closed")
    }
    scope.wg.Add(1)
    scope.mu.Unlock()
    go _goml_task_run(scope, body)
    return true
}

func _goml_runtime_std_task_scope_finish(scope_id int64) struct{} {
    var scope *_goml_task_scope_state = _goml_task_lookup(scope_id)
    if scope == nil {
        return struct{}{}
    }
    scope.mu.Lock()
    if scope.state == 0 {
        scope.state = 1
    }
    scope.mu.Unlock()
    scope.wg.Wait()
    scope.cancel()
    scope.mu.Lock()
    var panicked bool = scope.panicked
    var panic_value any = scope.panic_value
    scope.state = 2
    scope.mu.Unlock()
    _goml_task_registry_mutex.Lock()
    delete(_goml_task_registry, scope_id)
    _goml_task_registry_mutex.Unlock()
    if panicked {
        panic(panic_value)
    }
    return struct{}{}
}

func array_get__Array_2_3int(arr [2]int, index int) int {
    return arr[index]
}

func array_get__Array_3_5uint8(arr [3]uint8, index int) uint8 {
    return arr[index]
}

type _goml_vec_uint8 struct {
    items []uint8
}

func vec_with_capacity__Vec_5uint8(capacity int) *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: make([]uint8, 0, capacity),
    }
}

func vec_push__Vec_5uint8(vec *_goml_vec_uint8, elem uint8) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_5uint8(vec *_goml_vec_uint8, index int) uint8 {
    return vec.items[index]
}

func vec_len__Vec_5uint8(vec *_goml_vec_uint8) int {
    return int(len(vec.items))
}

type _goml_vec_uint32 struct {
    items []uint32
}

type ref_Option__isize_x struct {
    value Option__isize
}

func ref__Ref_13Option__isize(value Option__isize) *ref_Option__isize_x {
    return &ref_Option__isize_x{
        value: value,
    }
}

func ref_get__Ref_13Option__isize(reference *ref_Option__isize_x) Option__isize {
    return reference.value
}

func ref_set__Ref_13Option__isize(reference *ref_Option__isize_x, value Option__isize) struct{} {
    reference.value = value
    return struct{}{}
}

type ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x struct {
    value _goml_m_Option____Result____isize____std_p_panic_p_Panic
}

func ref___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(value _goml_m_Option____Result____isize____std_p_panic_p_Panic) *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x {
    return &ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x{
        value: value,
    }
}

func ref_get___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(reference *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x) _goml_m_Option____Result____isize____std_p_panic_p_Panic {
    return reference.value
}

func ref_set___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(reference *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x, value _goml_m_Option____Result____isize____std_p_panic_p_Panic) struct{} {
    reference.value = value
    return struct{}{}
}

type Tuple2_4unit_4bool struct {
    _0 struct{}
    _1 bool
}

type Tuple2_4bool_6string struct {
    _0 bool
    _1 string
}

type FloatNatural struct {
    words *_goml_vec_uint32
}

type ParsedFloat struct {
    valid bool
    negative bool
    special int
    numerator FloatNatural
    decimal_exponent int
    binary_exponent int
    hexadecimal bool
    significant_digits int
}

type _goml_m_std_p_internal_p_task_p_ScopeHandle struct {
    id int64
}

type _goml_m_std_p_internal_p_task_p_CancelToken struct {
    id int64
}

type _goml_m_std_p_panic_p_Panic struct {
    message string
    stack string
    rethrow func() struct{}
}

type _goml_m_std_p_task_p_CancelToken struct {
    value _goml_m_std_p_internal_p_task_p_CancelToken
}

type _goml_m_std_p_task_p_Scope struct {
    handle _goml_m_std_p_internal_p_task_p_ScopeHandle
}

type _goml_m_std_p_task_p_ConcurrencyLimit struct {
    slots chan struct{}
}

type _goml_m_std_p_task_p_Task____isize struct {
    result *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
    ready chan struct{}
}

type closure_env_main_0 struct {
    value_0 int
}

type closure_env_main_1 struct {}

type closure_env_constructor_2 struct {}

type closure_env_store_3 struct {}

type closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_4 struct {
    task_0 _goml_m_std_p_task_p_Task____isize
    token_1 _goml_m_std_p_task_p_CancelToken
    body_2 func(_goml_m_std_p_task_p_CancelToken) int
}

type closure_env_std_task_run_scope_T_isize_5 struct {
    result_0 *ref_Option__isize_x
    body_1 func(_goml_m_std_p_task_p_Scope) int
    handle_2 _goml_m_std_p_internal_p_task_p_ScopeHandle
}

type closure_env_inherent_std_task_Task_std_task_Task_T_run_T_isize_6 struct {
    body_0 func(_goml_m_std_p_task_p_CancelToken) int
    token_1 _goml_m_std_p_task_p_CancelToken
}

type closure_env_std_panic_catch_T_isize_7 struct {
    result_0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
    body_1 func() int
}

type closure_env_std_panic_catch_T_isize_8 struct {
    result_0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
}

type Ordering uint8

type _goml_m_Result____isize____std_p_panic_p_Panic interface {
    is_goml_m_Result____isize____std_p_panic_p_Panic()
}

type Ok struct {
    _0 int
}

func (_ Ok) is_goml_m_Result____isize____std_p_panic_p_Panic() {}

type Err struct {
    _0 _goml_m_std_p_panic_p_Panic
}

func (_ Err) is_goml_m_Result____isize____std_p_panic_p_Panic() {}

type _goml_m_Option____Result____isize____std_p_panic_p_Panic struct {
    _p0 _goml_m_Result____isize____std_p_panic_p_Panic
    _tag uint8
}

type Message__isize struct {
    _p0 int
    _tag uint8
}

type Message__string struct {
    _p0 string
    _tag uint8
}

type Message__bool uint8

type _goml_m_Storage____isize_____d_const__value_k_2 struct {
    _p0 [2]int
    _tag uint8
}

type _goml_m_Option_____o__q_ uint8

const (
    _goml_m_Option_____o__q__None _goml_m_Option_____o__q_ = 0
    _goml_m_Option_____o__q__Some _goml_m_Option_____o__q_ = 1
)

type Option__isize struct {
    _p0 int
    _tag uint8
}

func main0() struct{} {
    var value__0 int = 1
    var value__1 int = 2
    var value__2 int = 3
    var t0 int = value__1 + value__2
    var inline29 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline29)
    var inline27 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(value__0)
    _goml_runtime_core_string_println(inline27)
    var computed__0 int = 4
    var inline25 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(computed__0)
    _goml_runtime_core_string_println(inline25)
    var inline23 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(value__0)
    _goml_runtime_core_string_println(inline23)
    var t1 closure_env_main_1 = closure_env_main_1{}
    var t2 func(_goml_m_std_p_task_p_Scope) int = func(p0 _goml_m_std_p_task_p_Scope) int {
        return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t1, p0)
    }
    var scoped__0 int
    var inline21 _goml_m_std_p_internal_p_task_p_ScopeHandle = _goml_m_std_p_internal_p_task_p_root__scope()
    var inline22 int = _goml_m_std_p_task_p_run__scope____T__isize(inline21, t2)
    scoped__0 = inline22
    var inline19 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(scoped__0)
    _goml_runtime_core_string_println(inline19)
    var inline17 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(value__0)
    _goml_runtime_core_string_println(inline17)
    var count__0 int = 6
    var compound_old0 int = count__0
    var compound_value0 int = 1
    var t3 int = compound_old0 + compound_value0
    count__0 = t3
    var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(count__0)
    _goml_runtime_core_string_println(inline15)
    var t5 [2]int = [2]int{9, 10}
    var t6 int
    var inline14 int = array_get__Array_2_3int(t5, 0)
    t6 = inline14
    var inline12 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t6)
    _goml_runtime_core_string_println(inline12)
    var t7 closure_env_constructor_2 = closure_env_constructor_2{}
    var constructor__0 func(string) Message__string = func(p0 string) Message__string {
        return _goml_m_inherent_i_closure__en_h7664649402e63076831a53b5acca4fde_ctor__2_i_apply(t7, p0)
    }
    var indirect__0 Message__string = constructor__0("bound")
    var empty__0 Message__bool = Message__bool(0)
    var x3 int = 11
    var inline10 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x3)
    _goml_runtime_core_string_println(inline10)
    switch indirect__0._tag {
    case 0:
    case 1:
        var x2 string = indirect__0._p0
        var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x2)
        _goml_runtime_core_string_println(inline8)
    default:
        panic("non-exhaustive match")
    }
    var jp0 bool
    switch empty__0 != Message__bool(0) {
    case false:
        jp0 = true
    case true:
        var x1 bool = empty__0 == Message__bool(2)
        jp0 = x1
    default:
        panic("non-exhaustive match")
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(jp0)
    _goml_runtime_core_string_println(inline6)
    var t8 [2]int = [2]int{12, 13}
    var t9 closure_env_store_3 = closure_env_store_3{}
    var store__0 func([2]int) _goml_m_Storage____isize_____d_const__value_k_2 = func(p0 [2]int) _goml_m_Storage____isize_____d_const__value_k_2 {
        return _goml_m_inherent_i_closure__env__store__3_i_closure__env__store__3_i_apply(t9, p0)
    }
    var t10 [2]int = [2]int{14, 15}
    var indirect_store__0 _goml_m_Storage____isize_____d_const__value_k_2 = store__0(t10)
    var t13 int = array_get__Array_2_3int(t8, 1)
    var inline4 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t13)
    _goml_runtime_core_string_println(inline4)
    switch indirect_store__0._tag {
    case 0:
    case 1:
        var x0 [2]int = indirect_store__0._p0
        var t12 int = array_get__Array_2_3int(x0, 0)
        var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t12)
        _goml_runtime_core_string_println(inline2)
    default:
        panic("non-exhaustive match")
    }
    var narrow__0 [3]uint8 = func(value uint8) [3]uint8 {
        var result [3]uint8
        for index := range result {
            result[index] = value
        }
        return result
    }(1)
    var t11 uint8 = array_get__Array_3_5uint8(narrow__0, 2)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_u8_i_to__string(t11)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_std_p_task_p_run__scope____T__isize(handle__0 _goml_m_std_p_internal_p_task_p_ScopeHandle, body__0 func(_goml_m_std_p_task_p_Scope) int) int {
    var result__0 *ref_Option__isize_x
    var inline7 *ref_Option__isize_x = ref__Ref_13Option__isize(Option__isize{
        _tag: 0,
    })
    result__0 = inline7
    var t0 closure_env_std_task_run_scope_T_isize_5 = closure_env_std_task_run_scope_T_isize_5{
        result_0: result__0,
        body_1: body__0,
        handle_2: handle__0,
    }
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h1b381f58657f78951ea5718f3d6a5402_size__5_i_apply(t0)
    }
    var inline3 int64 = handle__0.id
    _goml_m_std_p_internal_p_task_p_scope__run(inline3, t1)
    var inline5 int64 = handle__0.id
    _goml_m_std_p_internal_p_task_p_scope__finish(inline5)
    var inline0 Option__isize = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__Option_l_isize_r_(result__0)
    switch inline0._tag {
    case 0:
        _goml_m_std_p_panic_p_raise("scope completed without a result")
        panic("unreachable")
    case 1:
        var inline2 int = inline0._p0
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_internal_p_task_p_root__scope() _goml_m_std_p_internal_p_task_p_ScopeHandle {
    var t0 int64
    var inline0 int64 = 0
    var inline1 int64 = _goml_runtime_std_task_scope_new(inline0)
    t0 = inline1
    var t1 _goml_m_std_p_internal_p_task_p_ScopeHandle = _goml_m_std_p_internal_p_task_p_ScopeHandle{
        id: t0,
    }
    return t1
}

func _goml_m_inherent_i_std_p_task__h32a04edffc0d154c16cc452b495505e0_ing____T__isize() _goml_m_std_p_task_p_Task____isize {
    var t0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
    var inline2 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = ref___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(_goml_m_Option____Result____isize____std_p_panic_p_Panic{
        _tag: 0,
    })
    t0 = inline2
    var t1 chan struct{}
    var inline0 int = 0
    var inline1 chan struct{} = func(p0 int) chan struct{} {
        return make(chan struct{}, p0)
    }(inline0)
    t1 = inline1
    var t2 _goml_m_std_p_task_p_Task____isize = _goml_m_std_p_task_p_Task____isize{
        result: t0,
        ready: t1,
    }
    return t2
}

func _goml_m_inherent_i_std_p_task_p_Scope_i_std_p_task_p_Scope_i_token(self__0 _goml_m_std_p_task_p_Scope) _goml_m_std_p_task_p_CancelToken {
    var t0 _goml_m_std_p_internal_p_task_p_ScopeHandle = self__0.handle
    var t1 _goml_m_std_p_internal_p_task_p_CancelToken
    var inline0 int64 = t0.id
    var inline1 _goml_m_std_p_internal_p_task_p_CancelToken = _goml_m_std_p_internal_p_task_p_CancelToken{
        id: inline0,
    }
    t1 = inline1
    var t2 _goml_m_std_p_task_p_CancelToken = _goml_m_std_p_task_p_CancelToken{
        value: t1,
    }
    return t2
}

func _goml_m_std_p_internal_p_task_p_spawn(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle, body__0 func() struct{}) struct{} {
    var t0 int64 = scope__0.id
    _goml_runtime_std_task_scope_spawn(t0, body__0)
    return struct{}{}
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_recv____T___o__q_(self__0 chan struct{}) _goml_m_Option_____o__q_ {
    var mtmp0 Tuple2_4unit_4bool = func(p0 chan struct{}) Tuple2_4unit_4bool {
        var value struct{}
        var ok bool
        value, ok = <-p0
        return Tuple2_4unit_4bool{
            _0: value,
            _1: ok,
        }
    }(self__0)
    var x1 bool = mtmp0._1
    if x1 {
        var t0 _goml_m_Option_____o__q_ = _goml_m_Option_____o__q__Some
        return t0
    } else {
        return _goml_m_Option_____o__q__None
    }
}

func _goml_m_inherent_i_Ref_i_Ref_l_h80255cab1de9a17c566649cce0d96ab2_c_p_Panic_r__r_(self__0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x) _goml_m_Option____Result____isize____std_p_panic_p_Panic {
    var t0 _goml_m_Option____Result____isize____std_p_panic_p_Panic = ref_get___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(self__0)
    return t0
}

func _goml_m_std_p_panic_p_raise(message__0 string) struct{} {
    _goml_runtime_std_panic_raise(message__0)
    panic("unreachable")
}

func _goml_m_std_p_panic_p_resume(info__0 _goml_m_std_p_panic_p_Panic) struct{} {
    var t0 func() struct{} = info__0.rethrow
    t0()
    panic("unreachable")
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_u8_i_to__string(self__0 uint8) string {
    var inline0 uint64 = uint64(uint8(self__0))
    var inline1 string = decimal_string(inline0)
    return inline1
}

func _goml_m_std_p_panic_p_catch____T__isize(body__0 func() int) _goml_m_Result____isize____std_p_panic_p_Panic {
    var result__0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
    var inline3 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = ref___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(_goml_m_Option____Result____isize____std_p_panic_p_Panic{
        _tag: 0,
    })
    result__0 = inline3
    var t0 closure_env_std_panic_catch_T_isize_7 = closure_env_std_panic_catch_T_isize_7{
        result_0: result__0,
        body_1: body__0,
    }
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h0a05d186557f06626dc1823ff592bceb_size__7_i_apply(t0)
    }
    var t2 closure_env_std_panic_catch_T_isize_8 = closure_env_std_panic_catch_T_isize_8{
        result_0: result__0,
    }
    var t3 func(string, string, func() struct{}) struct{} = func(p0 string, p1 string, p2 func() struct{}) struct{} {
        return _goml_m_inherent_i_closure__en_hdafd69dd08870b3e61977a9b195040fa_size__8_i_apply(t2, p0, p1, p2)
    }
    _goml_runtime_std_panic_catch(t1, t3)
    var mtmp0 _goml_m_Option____Result____isize____std_p_panic_p_Panic
    var inline2 _goml_m_Option____Result____isize____std_p_panic_p_Panic = ref_get___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(result__0)
    mtmp0 = inline2
    switch mtmp0._tag {
    case 0:
        var inline0 string = "panic boundary returned without a result"
        _goml_runtime_std_panic_raise(inline0)
        panic("unreachable")
    case 1:
        var x0 _goml_m_Result____isize____std_p_panic_p_Panic = mtmp0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_Ref_i_Ref_l_h773c09d2b6237745ec73e218dcbd3298_c_p_Panic_r__r_(self__0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x, value__0 _goml_m_Option____Result____isize____std_p_panic_p_Panic) struct{} {
    ref_set___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(self__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_close____T___o__q_(self__0 chan struct{}) struct{} {
    func(p0 chan struct{}) struct{} {
        close(p0)
        return struct{}{}
    }(self__0)
    return struct{}{}
}

func signed_decimal_string(value__0 int64) string {
    var t0 bool = value__0 < 0
    if t0 {
        var t1 uint64 = uint64(int64(value__0))
        var t2 uint64 = 0 - t1
        var t3 string = decimal_string(t2)
        var t4 string = "-" + t3
        return t4
    } else {
        var t5 uint64 = uint64(int64(value__0))
        var t6 string = decimal_string(t5)
        return t6
    }
}

func _goml_m_std_p_internal_p_task_p_scope__run(scope__0 int64, body__0 func() struct{}) struct{} {
    _goml_runtime_std_task_scope_run(scope__0, body__0)
    return struct{}{}
}

func _goml_m_std_p_internal_p_task_p_scope__finish(scope__0 int64) struct{} {
    _goml_runtime_std_task_scope_finish(scope__0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__Option_l_isize_r_(self__0 *ref_Option__isize_x) Option__isize {
    var t0 Option__isize = ref_get__Ref_13Option__isize(self__0)
    return t0
}

func decimal_string(value__0 uint64) string {
    var t0 bool = value__0 == 0
    if t0 {
        return "0"
    } else {
        var reversed__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(20)
        var remaining__0 uint64 = value__0
        Loop_loop0:
        for {
            var t10 bool = remaining__0 > 0
            if t10 {
                var t11 uint64 = remaining__0 % 10
                var t12 uint8 = uint8(uint64(t11))
                var t13 uint8 = t12 + 48
                vec_push__Vec_5uint8(reversed__0, t13)
                var compound_old1 uint64 = remaining__0
                var compound_value1 uint64 = 10
                var t14 uint64 = compound_old1 / compound_value1
                remaining__0 = t14
                continue
            } else {
                break Loop_loop0
            }
        }
        var t1 int
        var inline3 int = vec_len__Vec_5uint8(reversed__0)
        t1 = inline3
        var bytes__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(t1)
        var offset__0 int = 0
        Loop_loop1:
        for {
            var t2 int
            var inline2 int = vec_len__Vec_5uint8(reversed__0)
            t2 = inline2
            var t3 bool = offset__0 < t2
            if t3 {
                var t4 int
                var inline1 int = vec_len__Vec_5uint8(reversed__0)
                t4 = inline1
                var t5 int = t4 - offset__0
                var t6 int = t5 - 1
                var t7 uint8 = vec_get__Vec_5uint8(reversed__0, t6)
                vec_push__Vec_5uint8(bytes__0, t7)
                var compound_old0 int = offset__0
                var compound_value0 int = 1
                var t8 int = compound_old0 + compound_value0
                offset__0 = t8
                continue
            } else {
                break Loop_loop1
            }
        }
        var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
        var x0 string = mtmp0._1
        return x0
    }
}

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0, _cancel__0 _goml_m_std_p_task_p_CancelToken) int {
    var value__0 int = env0.value_0
    return value__0
}

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1, scope__0 _goml_m_std_p_task_p_Scope) int {
    var value__0 int = 5
    var t0 closure_env_main_0 = closure_env_main_0{
        value_0: value__0,
    }
    var t1 func(_goml_m_std_p_task_p_CancelToken) int = func(p0 _goml_m_std_p_task_p_CancelToken) int {
        return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t0, p0)
    }
    var work__0 _goml_m_std_p_task_p_Task____isize
    var inline9 _goml_m_std_p_task_p_Task____isize = _goml_m_inherent_i_std_p_task__h32a04edffc0d154c16cc452b495505e0_ing____T__isize()
    var inline10 _goml_m_std_p_task_p_CancelToken = _goml_m_inherent_i_std_p_task_p_Scope_i_std_p_task_p_Scope_i_token(scope__0)
    var inline11 _goml_m_std_p_internal_p_task_p_ScopeHandle = scope__0.handle
    var inline12 closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_4 = closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_4{
        task_0: inline9,
        token_1: inline10,
        body_2: t1,
    }
    var inline13 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h06a158548ef8dc096a001bc90e690a25_size__4_i_apply(inline12)
    }
    _goml_m_std_p_internal_p_task_p_spawn(inline11, inline13)
    work__0 = inline9
    var inline0 chan struct{} = work__0.ready
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_recv____T___o__q_(inline0)
    var inline2 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = work__0.result
    var inline3 _goml_m_Option____Result____isize____std_p_panic_p_Panic = _goml_m_inherent_i_Ref_i_Ref_l_h80255cab1de9a17c566649cce0d96ab2_c_p_Panic_r__r_(inline2)
    switch inline3._tag {
    case 0:
        _goml_m_std_p_panic_p_raise("task completed without a result")
        panic("unreachable")
    case 1:
        var inline5 _goml_m_Result____isize____std_p_panic_p_Panic = inline3._p0
        switch inline5.(type) {
        case Ok:
            var inline6 int = inline5.(Ok)._0
            return inline6
        case Err:
            var inline7 _goml_m_std_p_panic_p_Panic = inline5.(Err)._0
            _goml_m_std_p_panic_p_resume(inline7)
            panic("unreachable")
        default:
            panic("non-exhaustive match")
        }
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_closure__en_h7664649402e63076831a53b5acca4fde_ctor__2_i_apply(env0 closure_env_constructor_2, ctor_arg_0 string) Message__string {
    var t0 Message__string = Message__string{
        _p0: ctor_arg_0,
        _tag: 1,
    }
    return t0
}

func _goml_m_inherent_i_closure__env__store__3_i_closure__env__store__3_i_apply(env0 closure_env_store_3, ctor_arg_0 [2]int) _goml_m_Storage____isize_____d_const__value_k_2 {
    var t0 _goml_m_Storage____isize_____d_const__value_k_2 = _goml_m_Storage____isize_____d_const__value_k_2{
        _p0: ctor_arg_0,
        _tag: 1,
    }
    return t0
}

func _goml_m_inherent_i_closure__en_h06a158548ef8dc096a001bc90e690a25_size__4_i_apply(env0 closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_4) struct{} {
    var task__0 _goml_m_std_p_task_p_Task____isize = env0.task_0
    var token__0 _goml_m_std_p_task_p_CancelToken = env0.token_1
    var body__0 func(_goml_m_std_p_task_p_CancelToken) int = env0.body_2
    var inline0 closure_env_inherent_std_task_Task_std_task_Task_T_run_T_isize_6 = closure_env_inherent_std_task_Task_std_task_Task_T_run_T_isize_6{
        body_0: body__0,
        token_1: token__0,
    }
    var inline1 func() int = func() int {
        return _goml_m_inherent_i_closure__en_h4ef4882a29fe8b3faf9424b454f70125_size__6_i_apply(inline0)
    }
    var inline2 _goml_m_Result____isize____std_p_panic_p_Panic = _goml_m_std_p_panic_p_catch____T__isize(inline1)
    var inline3 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = task__0.result
    var inline4 _goml_m_Option____Result____isize____std_p_panic_p_Panic = _goml_m_Option____Result____isize____std_p_panic_p_Panic{
        _p0: inline2,
        _tag: 1,
    }
    _goml_m_inherent_i_Ref_i_Ref_l_h773c09d2b6237745ec73e218dcbd3298_c_p_Panic_r__r_(inline3, inline4)
    var inline6 chan struct{} = task__0.ready
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_close____T___o__q_(inline6)
    switch inline2.(type) {
    case Err:
        var inline8 _goml_m_std_p_panic_p_Panic = inline2.(Err)._0
        _goml_m_std_p_panic_p_resume(inline8)
        panic("unreachable")
    default:
        return struct{}{}
    }
}

func _goml_m_inherent_i_closure__en_h1b381f58657f78951ea5718f3d6a5402_size__5_i_apply(env0 closure_env_std_task_run_scope_T_isize_5) struct{} {
    var result__0 *ref_Option__isize_x = env0.result_0
    var body__0 func(_goml_m_std_p_task_p_Scope) int = env0.body_1
    var handle__0 _goml_m_std_p_internal_p_task_p_ScopeHandle = env0.handle_2
    var t0 _goml_m_std_p_task_p_Scope = _goml_m_std_p_task_p_Scope{
        handle: handle__0,
    }
    var t1 int = body__0(t0)
    var t2 Option__isize = Option__isize{
        _p0: t1,
        _tag: 1,
    }
    ref_set__Ref_13Option__isize(result__0, t2)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_h4ef4882a29fe8b3faf9424b454f70125_size__6_i_apply(env0 closure_env_inherent_std_task_Task_std_task_Task_T_run_T_isize_6) int {
    var body__0 func(_goml_m_std_p_task_p_CancelToken) int = env0.body_0
    var token__0 _goml_m_std_p_task_p_CancelToken = env0.token_1
    var t0 int = body__0(token__0)
    return t0
}

func _goml_m_inherent_i_closure__en_h0a05d186557f06626dc1823ff592bceb_size__7_i_apply(env0 closure_env_std_panic_catch_T_isize_7) struct{} {
    var result__0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = env0.result_0
    var body__0 func() int = env0.body_1
    var t0 int = body__0()
    var t1 _goml_m_Result____isize____std_p_panic_p_Panic = Ok{
        _0: t0,
    }
    var t2 _goml_m_Option____Result____isize____std_p_panic_p_Panic = _goml_m_Option____Result____isize____std_p_panic_p_Panic{
        _p0: t1,
        _tag: 1,
    }
    ref_set___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(result__0, t2)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_hdafd69dd08870b3e61977a9b195040fa_size__8_i_apply(env0 closure_env_std_panic_catch_T_isize_8, message__0 string, stack__0 string, rethrow__0 func() struct{}) struct{} {
    var result__0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = env0.result_0
    var t0 _goml_m_std_p_panic_p_Panic = _goml_m_std_p_panic_p_Panic{
        message: message__0,
        stack: stack__0,
        rethrow: rethrow__0,
    }
    var t1 _goml_m_Result____isize____std_p_panic_p_Panic = Err{
        _0: t0,
    }
    var t2 _goml_m_Option____Result____isize____std_p_panic_p_Panic = _goml_m_Option____Result____isize____std_p_panic_p_Panic{
        _p0: t1,
        _tag: 1,
    }
    ref_set___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(result__0, t2)
    return struct{}{}
}

func main() {
    main0()
}
