package main

import (
    _goml_context "context"
    _goml_os "os"
    _goml_sync "sync"
    _goml_time "time"
)

var _goml_time_timer_mutex _goml_sync.Mutex = _goml_sync.Mutex{}

var _goml_time_timer_registry map[int64]*_goml_time.Timer = make(map[int64]*_goml_time.Timer)

var _goml_time_timer_next_id int64 = 1

func _goml_time_timer_fire(id int64, ready chan struct{}) struct{} {
    _goml_time_timer_mutex.Lock()
    var timer *_goml_time.Timer = _goml_time_timer_registry[id]
    if timer != nil {
        delete(_goml_time_timer_registry, id)
        close(ready)
    }
    _goml_time_timer_mutex.Unlock()
    return struct{}{}
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

func _goml_task_context(id int64) _goml_context.Context {
    var scope *_goml_task_scope_state = _goml_task_lookup(id)
    if scope != nil {
        return scope.ctx
    }
    var ctx _goml_context.Context
    var cancel _goml_context.CancelFunc
    ctx, cancel = _goml_context.WithCancel(_goml_context.Background())
    cancel()
    return ctx
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

func _goml_runtime_std_time_timer_new(nanoseconds int64) Tuple2_5int64_14Receiver_4unit {
    var ready chan struct{} = make(chan struct{})
    _goml_time_timer_mutex.Lock()
    var id int64 = _goml_time_timer_next_id
    _goml_time_timer_next_id = _goml_time_timer_next_id + 1
    var timer *_goml_time.Timer = _goml_time.AfterFunc(_goml_time.Duration(nanoseconds), func() {
        _goml_time_timer_fire(id, ready)
    })
    _goml_time_timer_registry[id] = timer
    _goml_time_timer_mutex.Unlock()
    return Tuple2_5int64_14Receiver_4unit{
        _0: id,
        _1: ready,
    }
}

func _goml_runtime_std_time_timer_stop(id int64) bool {
    _goml_time_timer_mutex.Lock()
    var timer *_goml_time.Timer = _goml_time_timer_registry[id]
    if timer == nil {
        _goml_time_timer_mutex.Unlock()
        return false
    }
    var stopped bool = timer.Stop()
    if stopped {
        delete(_goml_time_timer_registry, id)
    }
    _goml_time_timer_mutex.Unlock()
    return stopped
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

func _goml_runtime_std_task_scope_cancel(scope_id int64) struct{} {
    var scope *_goml_task_scope_state = _goml_task_lookup(scope_id)
    if scope != nil {
        scope.cancel()
    }
    return struct{}{}
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

func _goml_runtime_std_task_scope_done(scope_id int64) <-chan struct{} {
    return _goml_task_context(scope_id).Done()
}

type _goml_vec_string struct {
    items []string
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

type _goml_vec_Tuple2_6string_6string struct {
    items []Tuple2_6string_6string
}

type _goml_vec_uint32 struct {
    items []uint32
}

type ref_int_x struct {
    value int
}

type ref__goml_m_Option_____o__q__x struct {
    value _goml_m_Option_____o__q_
}

func ref___goml_m_Ref__10Option_____o__q_(value _goml_m_Option_____o__q_) *ref__goml_m_Option_____o__q__x {
    return &ref__goml_m_Option_____o__q__x{
        value: value,
    }
}

func ref_get___goml_m_Ref__10Option_____o__q_(reference *ref__goml_m_Option_____o__q__x) _goml_m_Option_____o__q_ {
    return reference.value
}

func ref_set___goml_m_Ref__10Option_____o__q_(reference *ref__goml_m_Option_____o__q__x, value _goml_m_Option_____o__q_) struct{} {
    reference.value = value
    return struct{}{}
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

type Tuple3_4bool_6string_6string struct {
    _0 bool
    _1 string
    _2 string
}

type Tuple2_4bool_6string struct {
    _0 bool
    _1 string
}

type Tuple3_4bool_10Vec_5uint8_6string struct {
    _0 bool
    _1 *_goml_vec_uint8
    _2 string
}

type Tuple6_4bool_10Vec_5uint8_3int_4bool_3int_6string struct {
    _0 bool
    _1 *_goml_vec_uint8
    _2 int
    _3 bool
    _4 int
    _5 string
}

type Tuple5_4bool_3int_4bool_3int_6string struct {
    _0 bool
    _1 int
    _2 bool
    _3 int
    _4 string
}

type Tuple9_4bool_3int_5int64_6uint32_5int64_3int_4bool_3int_6string struct {
    _0 bool
    _1 int
    _2 int64
    _3 uint32
    _4 int64
    _5 int
    _6 bool
    _7 int
    _8 string
}

type Tuple3_4bool_11Vec_6string_6string struct {
    _0 bool
    _1 *_goml_vec_string
    _2 string
}

type Tuple6_4bool_11Vec_6string_3int_4bool_3int_6string struct {
    _0 bool
    _1 *_goml_vec_string
    _2 int
    _3 bool
    _4 int
    _5 string
}

type Tuple2_6string_6string struct {
    _0 string
    _1 string
}

type Tuple5_4bool_3int_10Vec_5uint8_10Vec_5uint8_6string struct {
    _0 bool
    _1 int
    _2 *_goml_vec_uint8
    _3 *_goml_vec_uint8
    _4 string
}

type Tuple6_4bool_3int_10Vec_5uint8_10Vec_5uint8_6string_4bool struct {
    _0 bool
    _1 int
    _2 *_goml_vec_uint8
    _3 *_goml_vec_uint8
    _4 string
    _5 bool
}

type Tuple3_4bool_3int_6string struct {
    _0 bool
    _1 int
    _2 string
}

type Tuple4_4bool_3int_6string_4bool struct {
    _0 bool
    _1 int
    _2 string
    _3 bool
}

type Tuple2_5int64_14Receiver_4unit struct {
    _0 int64
    _1 <-chan struct{}
}

type Tuple3_4bool_4char_3int struct {
    _0 bool
    _1 rune
    _2 int
}

type Tuple2_4unit_4bool struct {
    _0 struct{}
    _1 bool
}

type Tuple2_4bool_4char struct {
    _0 bool
    _1 rune
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

type _goml_m_std_p_task_p_CancelToken struct {
    value _goml_m_std_p_internal_p_task_p_CancelToken
}

type _goml_m_std_p_task_p_Scope struct {
    handle _goml_m_std_p_internal_p_task_p_ScopeHandle
}

type _goml_m_std_p_task_p_ConcurrencyLimit struct {
    slots chan struct{}
}

type _goml_m_std_p_bytes_p_BoundsError struct {
    offset_value int
    needed_value int
    length_value int
}

type _goml_m_std_p_bytes_p_Builder struct {
    values *_goml_vec_uint8
}

type _goml_m_std_p_bytes_p_Bytes struct {
    values *_goml_vec_uint8
}

type _goml_m_std_p_utf8_p_Utf8Error struct {
    valid_up_to_value int
    error_length_value Option__isize
}

type _goml_m_std_p_io_p_ErrorDetails struct {
    kind_value _goml_m_std_p_io_p_ErrorKind
    operation_value string
    context_value Option__string
    raw_os_code_value Option__isize
    message_value string
}

type _goml_m_std_p_io_p_Error struct {
    details _goml_m_std_p_io_p_ErrorDetails
}

type _goml_m_std_p_io_p_Cursor struct {
    data _goml_m_std_p_bytes_p_Bytes
    offset *ref_int_x
}

type _goml_m_std_p_io_p_Stdout struct {}

type _goml_m_std_p_io_p_Stderr struct {}

type _goml_m_std_p_io_p_Stdin struct {}

type _goml_m_std_p_time_p_Error struct {
    details _goml_m_std_p_io_p_ErrorDetails
}

type _goml_m_std_p_time_p_Duration struct {
    nanoseconds int64
}

type _goml_m_std_p_time_p_Timer struct {
    id int64
    ready <-chan struct{}
}

type _goml_m_std_p_time_p_Instant struct {
    nanoseconds int64
}

type _goml_m_std_p_time_p_SystemTime struct {
    unix_nanoseconds int64
}

type _goml_m_std_p_task_p_Task____isize struct {
    result *ref_Option__isize_x
    ready chan struct{}
}

type FnIterator__u8 struct {
    next_fn func() Option__u8
}

type closure_env_std_io_read_stdin_to_string_0 struct {}

type closure_env_main_1 struct {}

type closure_env_main_2 struct {}

type closure_env_main_3 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Cursor_4 struct {}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_5 struct {}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_6 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Stdin_7 struct {}

type closure_env_std_task_scope_T_8 struct {
    result_0 *ref__goml_m_Option_____o__q__x
    body_1 func(_goml_m_std_p_task_p_Scope) struct{}
    handle_2 _goml_m_std_p_internal_p_task_p_ScopeHandle
}

type closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_9 struct {
    result_0 *ref_Option__isize_x
    body_1 func(_goml_m_std_p_task_p_CancelToken) int
    token_2 _goml_m_std_p_task_p_CancelToken
    ready_3 chan struct{}
}

type closure_env_inherent_Slice_Slice_T_iter_T_u8_10 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 []uint8
}

type Ordering uint8

type _goml_m_std_p_io_p_ErrorKind uint8

type _goml_m_std_p_task_p_WaitResult_____o__q_ uint8

type Option__u8 uint16

type _goml_m_Option____Slice_l_u8_r_ struct {
    _p0 []uint8
    _tag uint8
}

type _goml_m_Option____MutSlice_l_u8_r_ struct {
    _p0 []uint8
    _tag uint8
}

type _goml_m_Result____string____std_p_utf8_p_Utf8Error struct {
    _p1 _goml_m_std_p_utf8_p_Utf8Error
    _p0 string
    _tag uint8
}

type Option__isize struct {
    _p0 int
    _tag uint8
}

type _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error struct {
    _p0 _goml_m_std_p_utf8_p_Utf8Error
    _tag uint8
}

type _goml_m_Result____isize____std_p_bytes_p_BoundsError struct {
    _p1 _goml_m_std_p_bytes_p_BoundsError
    _p0 int
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

type _goml_m_Result____std_p_bytes_p_Bytes____std_p_io_p_Error interface {
    is_goml_m_Result____std_p_bytes_p_Bytes____std_p_io_p_Error()
}

type _goml_m_Result____std_p_bytes_p_Bytes____std_p_io_p_Error_Ok struct {
    _0 _goml_m_std_p_bytes_p_Bytes
}

func (_ _goml_m_Result____std_p_bytes_p_Bytes____std_p_io_p_Error_Ok) is_goml_m_Result____std_p_bytes_p_Bytes____std_p_io_p_Error() {}

type _goml_m_Result____std_p_bytes_p_Bytes____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result____std_p_bytes_p_Bytes____std_p_io_p_Error_Err) is_goml_m_Result____std_p_bytes_p_Bytes____std_p_io_p_Error() {}

type _goml_m_Result____string____std_p_io_p_Error interface {
    is_goml_m_Result____string____std_p_io_p_Error()
}

type _goml_m_Result____string____std_p_io_p_Error_Ok struct {
    _0 string
}

func (_ _goml_m_Result____string____std_p_io_p_Error_Ok) is_goml_m_Result____string____std_p_io_p_Error() {}

type _goml_m_Result____string____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result____string____std_p_io_p_Error_Err) is_goml_m_Result____string____std_p_io_p_Error() {}

type _goml_m_Result_____o__q_____std_p_io_p_Error interface {
    is_goml_m_Result_____o__q_____std_p_io_p_Error()
}

type _goml_m_Result_____o__q_____std_p_io_p_Error_Ok struct {
    _0 struct{}
}

func (_ _goml_m_Result_____o__q_____std_p_io_p_Error_Ok) is_goml_m_Result_____o__q_____std_p_io_p_Error() {}

type _goml_m_Result_____o__q_____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result_____o__q_____std_p_io_p_Error_Err) is_goml_m_Result_____o__q_____std_p_io_p_Error() {}

type _goml_m_Result____isize____std_p_io_p_Error interface {
    is_goml_m_Result____isize____std_p_io_p_Error()
}

type _goml_m_Result____isize____std_p_io_p_Error_Ok struct {
    _0 int
}

func (_ _goml_m_Result____isize____std_p_io_p_Error_Ok) is_goml_m_Result____isize____std_p_io_p_Error() {}

type _goml_m_Result____isize____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result____isize____std_p_io_p_Error_Err) is_goml_m_Result____isize____std_p_io_p_Error() {}

type _goml_m_Result____Slice_l_u8_r_____std_p_io_p_Error interface {
    is_goml_m_Result____Slice_l_u8_r_____std_p_io_p_Error()
}

type _goml_m_Result____Slice_l_u8_r_____std_p_io_p_Error_Ok struct {
    _0 []uint8
}

func (_ _goml_m_Result____Slice_l_u8_r_____std_p_io_p_Error_Ok) is_goml_m_Result____Slice_l_u8_r_____std_p_io_p_Error() {}

type _goml_m_Result____Slice_l_u8_r_____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result____Slice_l_u8_r_____std_p_io_p_Error_Err) is_goml_m_Result____Slice_l_u8_r_____std_p_io_p_Error() {}

type _goml_m_Result____Option____string____std_p_io_p_Error interface {
    is_goml_m_Result____Option____string____std_p_io_p_Error()
}

type _goml_m_Result____Option____string____std_p_io_p_Error_Ok struct {
    _0 Option__string
}

func (_ _goml_m_Result____Option____string____std_p_io_p_Error_Ok) is_goml_m_Result____Option____string____std_p_io_p_Error() {}

type _goml_m_Result____Option____string____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result____Option____string____std_p_io_p_Error_Err) is_goml_m_Result____Option____string____std_p_io_p_Error() {}

type _goml_m_Option____std_p_time_p_Duration struct {
    _p0 _goml_m_std_p_time_p_Duration
    _tag uint8
}

type _goml_m_Result____std_p_time_p_SystemTime____std_p_time_p_Error interface {
    is_goml_m_Result____std_p_time_p_SystemTime____std_p_time_p_Error()
}

type _goml_m_Result____std_p_time_p_SystemTime____std_p_time_p_Error_Ok struct {
    _0 _goml_m_std_p_time_p_SystemTime
}

func (_ _goml_m_Result____std_p_time_p_SystemTime____std_p_time_p_Error_Ok) is_goml_m_Result____std_p_time_p_SystemTime____std_p_time_p_Error() {}

type _goml_m_Result____std_p_time_p_SystemTime____std_p_time_p_Error_Err struct {
    _0 _goml_m_std_p_time_p_Error
}

func (_ _goml_m_Result____std_p_time_p_SystemTime____std_p_time_p_Error_Err) is_goml_m_Result____std_p_time_p_SystemTime____std_p_time_p_Error() {}

type _goml_m_Result____std_p_time_p_Duration____std_p_time_p_Error interface {
    is_goml_m_Result____std_p_time_p_Duration____std_p_time_p_Error()
}

type _goml_m_Result____std_p_time_p_Duration____std_p_time_p_Error_Ok struct {
    _0 _goml_m_std_p_time_p_Duration
}

func (_ _goml_m_Result____std_p_time_p_Duration____std_p_time_p_Error_Ok) is_goml_m_Result____std_p_time_p_Duration____std_p_time_p_Error() {}

type _goml_m_Result____std_p_time_p_Duration____std_p_time_p_Error_Err struct {
    _0 _goml_m_std_p_time_p_Error
}

func (_ _goml_m_Result____std_p_time_p_Duration____std_p_time_p_Error_Err) is_goml_m_Result____std_p_time_p_Duration____std_p_time_p_Error() {}

type _goml_m_Option_____o__q_ uint8

const (
    _goml_m_Option_____o__q__None _goml_m_Option_____o__q_ = 0
    _goml_m_Option_____o__q__Some _goml_m_Option_____o__q_ = 1
)

type _goml_m_Result____Option____string____std_p_utf8_p_Utf8Error interface {
    is_goml_m_Result____Option____string____std_p_utf8_p_Utf8Error()
}

type _goml_m_Result____Option____string____std_p_utf8_p_Utf8Error_Ok struct {
    _0 Option__string
}

func (_ _goml_m_Result____Option____string____std_p_utf8_p_Utf8Error_Ok) is_goml_m_Result____Option____string____std_p_utf8_p_Utf8Error() {}

type _goml_m_Result____Option____string____std_p_utf8_p_Utf8Error_Err struct {
    _0 _goml_m_std_p_utf8_p_Utf8Error
}

func (_ _goml_m_Result____Option____string____std_p_utf8_p_Utf8Error_Err) is_goml_m_Result____Option____string____std_p_utf8_p_Utf8Error() {}

type Option__char uint64

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

func _goml_m_std_p_internal_p_task_p_spawn(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle, body__0 func() struct{}) struct{} {
    var t0 int64 = scope__0.id
    _goml_runtime_std_task_scope_spawn(t0, body__0)
    return struct{}{}
}

func _goml_m_std_p_internal_p_task_p_run(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle, body__0 func() struct{}) struct{} {
    var t0 int64 = scope__0.id
    _goml_runtime_std_task_scope_run(t0, body__0)
    var t1 int64 = scope__0.id
    _goml_runtime_std_task_scope_finish(t1)
    return struct{}{}
}

func _goml_m_std_p_internal_p_task_p_cancel(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle) struct{} {
    var t0 int64 = scope__0.id
    _goml_runtime_std_task_scope_cancel(t0)
    return struct{}{}
}

func _goml_m_std_p_internal_p_task_p_token(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle) _goml_m_std_p_internal_p_task_p_CancelToken {
    var t0 int64 = scope__0.id
    var t1 _goml_m_std_p_internal_p_task_p_CancelToken = _goml_m_std_p_internal_p_task_p_CancelToken{
        id: t0,
    }
    return t1
}

func _goml_m_inherent_i_std_p_inter_h25a0a7c69f0f67d3e38fa37e5f2fc20e_celToken_i_done(self__0 _goml_m_std_p_internal_p_task_p_CancelToken) <-chan struct{} {
    var t0 int64 = self__0.id
    var inline0 <-chan struct{} = _goml_runtime_std_task_scope_done(t0)
    return inline0
}

func _goml_m_std_p_internal_p_host_p_timer__new(value__0 int64) Tuple2_5int64_14Receiver_4unit {
    var t0 Tuple2_5int64_14Receiver_4unit = _goml_runtime_std_time_timer_new(value__0)
    return t0
}

func _goml_m_std_p_internal_p_host_p_timer__stop(id__0 int64) bool {
    var t0 bool = _goml_runtime_std_time_timer_stop(id__0)
    return t0
}

func _goml_m_inherent_i_std_p_time__h80a59d7ad7a13ffcd2fc533934df5aaa_om__nanoseconds(value__0 int64) _goml_m_std_p_time_p_Duration {
    var t0 bool = value__0 < 0
    var jp0 int64
    if t0 {
        jp0 = 0
    } else {
        jp0 = value__0
    }
    var t1 _goml_m_std_p_time_p_Duration = _goml_m_std_p_time_p_Duration{
        nanoseconds: jp0,
    }
    return t1
}

func _goml_m_inherent_i_std_p_time__h3ebf4d33c01c8036463b8efe535bf1db_as__nanoseconds(self__0 _goml_m_std_p_time_p_Duration) int64 {
    var t0 int64 = self__0.nanoseconds
    return t0
}

func _goml_m_inherent_i_std_p_time_p_Timer_i_std_p_time_p_Timer_i_new(duration__0 _goml_m_std_p_time_p_Duration) _goml_m_std_p_time_p_Timer {
    var t0 int64
    var inline1 int64 = duration__0.nanoseconds
    t0 = inline1
    var mtmp0 Tuple2_5int64_14Receiver_4unit
    var inline0 Tuple2_5int64_14Receiver_4unit = _goml_runtime_std_time_timer_new(t0)
    mtmp0 = inline0
    var x0 int64 = mtmp0._0
    var x1 <-chan struct{} = mtmp0._1
    var t1 _goml_m_std_p_time_p_Timer = _goml_m_std_p_time_p_Timer{
        id: x0,
        ready: x1,
    }
    return t1
}

func main0() struct{} {
    var t0 _goml_m_std_p_time_p_Duration
    var inline35 _goml_m_std_p_time_p_Duration = _goml_m_std_p_time_p_Duration{
        nanoseconds: 0,
    }
    t0 = inline35
    var timer__0 _goml_m_std_p_time_p_Timer = _goml_m_inherent_i_std_p_time_p_Timer_i_std_p_time_p_Timer_i_new(t0)
    var t1 <-chan struct{}
    var inline34 <-chan struct{} = timer__0.ready
    t1 = inline34
    var _goml_m_______1_i_select__open bool
    select {
    case _, _goml_m_______1_i_select__open = <-t1:
        if _goml_m_______1_i_select__open {}
        var inline31 int = 1
        var inline32 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline31)
        _goml_runtime_core_string_println(inline32)
    }
    var t2 _goml_m_std_p_time_p_Duration
    var inline28 int64 = 60
    var inline29 int64 = inline28 * 1000000000
    var inline30 _goml_m_std_p_time_p_Duration = _goml_m_inherent_i_std_p_time__h80a59d7ad7a13ffcd2fc533934df5aaa_om__nanoseconds(inline29)
    t2 = inline30
    var stopped__0 _goml_m_std_p_time_p_Timer
    var inline23 int64 = _goml_m_inherent_i_std_p_time__h3ebf4d33c01c8036463b8efe535bf1db_as__nanoseconds(t2)
    var inline24 Tuple2_5int64_14Receiver_4unit = _goml_m_std_p_internal_p_host_p_timer__new(inline23)
    var inline25 int64 = inline24._0
    var inline26 <-chan struct{} = inline24._1
    var inline27 _goml_m_std_p_time_p_Timer = _goml_m_std_p_time_p_Timer{
        id: inline25,
        ready: inline26,
    }
    stopped__0 = inline27
    var t3 bool
    var inline21 int64 = stopped__0.id
    var inline22 bool = _goml_m_std_p_internal_p_host_p_timer__stop(inline21)
    t3 = inline22
    var inline19 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline19)
    var t4 <-chan struct{}
    var inline18 <-chan struct{} = stopped__0.ready
    t4 = inline18
    var _goml_m_______0_i_select__open bool
    select {
    case _, _goml_m_______0_i_select__open = <-t4:
        if _goml_m_______0_i_select__open {}
        var inline12 int = 2
        var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline12)
        _goml_runtime_core_string_println(inline13)
    default:
        var inline15 int = 3
        var inline16 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline15)
        _goml_runtime_core_string_println(inline16)
    }
    var t5 closure_env_main_2 = closure_env_main_2{}
    var t6 func(_goml_m_std_p_task_p_Scope) struct{} = func(p0 _goml_m_std_p_task_p_Scope) struct{} {
        return _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(t5, p0)
    }
    var inline6 _goml_m_std_p_internal_p_task_p_ScopeHandle = _goml_m_std_p_internal_p_task_p_root__scope()
    var inline7 *ref__goml_m_Option_____o__q__x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Option_l__o__q__r_(_goml_m_Option_____o__q__None)
    var inline8 closure_env_std_task_scope_T_8 = closure_env_std_task_scope_T_8{
        result_0: inline7,
        body_1: t6,
        handle_2: inline6,
    }
    var inline9 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h8c7921d612fbf78d91587ed4377816c8_e__T__8_i_apply(inline8)
    }
    _goml_m_std_p_internal_p_task_p_run(inline6, inline9)
    _goml_m_std_p_task_p_completed__scope__value____T___o__q_(inline7)
    var t7 closure_env_main_3 = closure_env_main_3{}
    var t8 func(_goml_m_std_p_task_p_Scope) struct{} = func(p0 _goml_m_std_p_task_p_Scope) struct{} {
        return _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(t7, p0)
    }
    var inline0 _goml_m_std_p_internal_p_task_p_ScopeHandle = _goml_m_std_p_internal_p_task_p_root__scope()
    var inline1 *ref__goml_m_Option_____o__q__x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Option_l__o__q__r_(_goml_m_Option_____o__q__None)
    var inline2 closure_env_std_task_scope_T_8 = closure_env_std_task_scope_T_8{
        result_0: inline1,
        body_1: t8,
        handle_2: inline0,
    }
    var inline3 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h8c7921d612fbf78d91587ed4377816c8_e__T__8_i_apply(inline2)
    }
    _goml_m_std_p_internal_p_task_p_run(inline0, inline3)
    _goml_m_std_p_task_p_completed__scope__value____T___o__q_(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T___o__q_(capacity__0 int) chan struct{} {
    var t0 chan struct{} = func(p0 int) chan struct{} {
        return make(chan struct{}, p0)
    }(capacity__0)
    return t0
}

func _goml_m_inherent_i_std_p_task__h53c24a9f0a217bb05f370161a3b86100_oin____T__isize(self__0 _goml_m_std_p_task_p_Task____isize) int {
    var jp0 int
    Loop_loop_expr0:
    for {
        var t0 chan struct{} = self__0.ready
        var inline1 Tuple2_4unit_4bool = func(p0 chan struct{}) Tuple2_4unit_4bool {
            var value struct{}
            var ok bool
            value, ok = <-p0
            return Tuple2_4unit_4bool{
                _0: value,
                _1: ok,
            }
        }(t0)
        var inline3 bool = inline1._1
        if inline3 {} else {}
        var t1 *ref_Option__isize_x = self__0.result
        var mtmp1 Option__isize
        var inline0 Option__isize = ref_get__Ref_13Option__isize(t1)
        mtmp1 = inline0
        switch mtmp1._tag {
        case 1:
            var x0 int = mtmp1._p0
            jp0 = x0
            break Loop_loop_expr0
        default:
            continue
        }
    }
    return jp0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Option_l__o__q__r_(value__0 _goml_m_Option_____o__q_) *ref__goml_m_Option_____o__q__x {
    var t0 *ref__goml_m_Option_____o__q__x = ref___goml_m_Ref__10Option_____o__q_(value__0)
    return t0
}

func _goml_m_std_p_task_p_completed__scope__value____T___o__q_(result__0 *ref__goml_m_Option_____o__q__x) struct{} {
    Loop_loop_expr0:
    for {
        var mtmp0 _goml_m_Option_____o__q_
        var inline0 _goml_m_Option_____o__q_ = ref_get___goml_m_Ref__10Option_____o__q_(result__0)
        mtmp0 = inline0
        switch mtmp0 {
        case _goml_m_Option_____o__q__Some:
            break Loop_loop_expr0
        default:
            continue
        }
    }
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Option_l_isize_r_(value__0 Option__isize) *ref_Option__isize_x {
    var t0 *ref_Option__isize_x = ref__Ref_13Option__isize(value__0)
    return t0
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T___o__q_(self__0 chan struct{}) <-chan struct{} {
    var t0 <-chan struct{} = func(p0 chan struct{}) <-chan struct{} {
        return p0
    }(self__0)
    return t0
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

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1, ___0 _goml_m_std_p_task_p_CancelToken) int {
    return 42
}

func _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(env0 closure_env_main_2, scope__0 _goml_m_std_p_task_p_Scope) struct{} {
    var t0 closure_env_main_1 = closure_env_main_1{}
    var t1 func(_goml_m_std_p_task_p_CancelToken) int = func(p0 _goml_m_std_p_task_p_CancelToken) int {
        return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t0, p0)
    }
    var work__0 _goml_m_std_p_task_p_Task____isize
    var inline4 *ref_Option__isize_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Option_l_isize_r_(Option__isize{
        _tag: 0,
    })
    var inline5 chan struct{} = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T___o__q_(0)
    var inline6 _goml_m_std_p_internal_p_task_p_ScopeHandle = scope__0.handle
    var inline7 _goml_m_std_p_internal_p_task_p_CancelToken = _goml_m_std_p_internal_p_task_p_token(inline6)
    var inline8 _goml_m_std_p_task_p_CancelToken = _goml_m_std_p_task_p_CancelToken{
        value: inline7,
    }
    var inline9 _goml_m_std_p_internal_p_task_p_ScopeHandle = scope__0.handle
    var inline10 closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_9 = closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_9{
        result_0: inline4,
        body_1: t1,
        token_2: inline8,
        ready_3: inline5,
    }
    var inline11 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h2edb97eafd7aca63bdc79db1f12910bf_size__9_i_apply(inline10)
    }
    _goml_m_std_p_internal_p_task_p_spawn(inline9, inline11)
    var inline13 _goml_m_std_p_task_p_Task____isize = _goml_m_std_p_task_p_Task____isize{
        result: inline4,
        ready: inline5,
    }
    work__0 = inline13
    var t2 <-chan struct{}
    var inline2 chan struct{} = work__0.ready
    var inline3 <-chan struct{} = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T___o__q_(inline2)
    t2 = inline3
    var _goml_m_______0_i_select__open bool
    select {
    case _, _goml_m_______0_i_select__open = <-t2:
        if _goml_m_______0_i_select__open {}
        var t3 int = _goml_m_inherent_i_std_p_task__h53c24a9f0a217bb05f370161a3b86100_oin____T__isize(work__0)
        var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t3)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    }
}

func _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(env0 closure_env_main_3, scope__0 _goml_m_std_p_task_p_Scope) struct{} {
    var cancel__0 _goml_m_std_p_task_p_CancelToken
    var inline7 _goml_m_std_p_internal_p_task_p_ScopeHandle = scope__0.handle
    var inline8 _goml_m_std_p_internal_p_task_p_CancelToken = _goml_m_std_p_internal_p_task_p_token(inline7)
    var inline9 _goml_m_std_p_task_p_CancelToken = _goml_m_std_p_task_p_CancelToken{
        value: inline8,
    }
    cancel__0 = inline9
    var inline5 _goml_m_std_p_internal_p_task_p_ScopeHandle = scope__0.handle
    _goml_m_std_p_internal_p_task_p_cancel(inline5)
    var t0 <-chan struct{}
    var inline3 _goml_m_std_p_internal_p_task_p_CancelToken = cancel__0.value
    var inline4 <-chan struct{} = _goml_m_inherent_i_std_p_inter_h25a0a7c69f0f67d3e38fa37e5f2fc20e_celToken_i_done(inline3)
    t0 = inline4
    var _goml_m_______0_i_select__open bool
    select {
    case _, _goml_m_______0_i_select__open = <-t0:
        if _goml_m_______0_i_select__open {}
        var inline0 int = 4
        var inline1 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    }
}

func _goml_m_inherent_i_closure__en_h8c7921d612fbf78d91587ed4377816c8_e__T__8_i_apply(env0 closure_env_std_task_scope_T_8) struct{} {
    var result__0 *ref__goml_m_Option_____o__q__x = env0.result_0
    var body__0 func(_goml_m_std_p_task_p_Scope) struct{} = env0.body_1
    var handle__0 _goml_m_std_p_internal_p_task_p_ScopeHandle = env0.handle_2
    var t0 _goml_m_std_p_task_p_Scope = _goml_m_std_p_task_p_Scope{
        handle: handle__0,
    }
    body__0(t0)
    var t2 _goml_m_Option_____o__q_ = _goml_m_Option_____o__q__Some
    ref_set___goml_m_Ref__10Option_____o__q_(result__0, t2)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_h2edb97eafd7aca63bdc79db1f12910bf_size__9_i_apply(env0 closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_9) struct{} {
    var result__0 *ref_Option__isize_x = env0.result_0
    var body__0 func(_goml_m_std_p_task_p_CancelToken) int = env0.body_1
    var token__0 _goml_m_std_p_task_p_CancelToken = env0.token_2
    var ready__0 chan struct{} = env0.ready_3
    var t0 int = body__0(token__0)
    var t1 Option__isize = Option__isize{
        _p0: t0,
        _tag: 1,
    }
    ref_set__Ref_13Option__isize(result__0, t1)
    func(p0 chan struct{}) struct{} {
        close(p0)
        return struct{}{}
    }(ready__0)
    return struct{}{}
}

func main() {
    main0()
}
