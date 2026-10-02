package main

import (
    _goml_fmt "fmt"
    _goml_debug "runtime/debug"
    _goml_context "context"
    _goml_os "os"
    _goml_sync "sync"
    _goml_time "time"
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

type _goml_vec_int struct {
    items []int
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

type ref_int_x struct {
    value int
}

type ref__goml_m_Option____std_p_utf8_p_Utf8Error_x struct {
    value _goml_m_Option____std_p_utf8_p_Utf8Error
}

type ref_bool_x struct {
    value bool
}

type ref_string_x struct {
    value string
}

type ref_Option__isize_x struct {
    value Option__isize
}

type Tuple2_5int64_14Receiver_4unit struct {
    _0 int64
    _1 <-chan struct{}
}

type Tuple2_4unit_4bool struct {
    _0 struct{}
    _1 bool
}

type Tuple2_4bool_6string struct {
    _0 bool
    _1 string
}

type Tuple2_3int_3int struct {
    _0 int
    _1 int
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

type _goml_m_std_p_unicode_p_Range struct {
    low uint32
    high uint32
    stride uint32
}

type _goml_m_std_p_unicode_p_RangeTable struct {
    data string
}

type _goml_m_std_p_unicode_p_CaseMapping struct {
    from rune
    upper rune
    lower rune
    title rune
}

type _goml_m_std_p_unicode_p_SpecialCase struct {
    data string
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

type _goml_m_std_p_bytes_p_FrozenBytes struct {
    values FrozenVec__u8
    offset int
    length int
}

type _goml_m_std_p_bytes_p_Finder struct {
    pattern *_goml_vec_uint8
    fallback *_goml_vec_int
}

type _goml_m_std_p_utf8_p_Utf8Error struct {
    valid_up_to_value int
    error_length_value Option__isize
}

type _goml_m_std_p_utf8_p_Decoder struct {
    pending *_goml_vec_uint8
    offset *ref_int_x
    failure *ref__goml_m_Option____std_p_utf8_p_Utf8Error_x
    finished *ref_bool_x
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

type _goml_m_std_p_io_p_TransferError struct {
    transferred_value uint
    error_value _goml_m_std_p_io_p_Error
}

type _goml_m_std_p_io_p_Cursor struct {
    data _goml_m_std_p_bytes_p_Bytes
    offset *ref_int_x
    limit int
}

type _goml_m_std_p_io_p_Discard struct {}

type _goml_m_std_p_io_p_PeekError struct {
    available_value []uint8
    error_value _goml_m_std_p_io_p_Error
}

type _goml_m_std_p_io_p_Stdout struct {}

type _goml_m_std_p_io_p_Stderr struct {}

type _goml_m_std_p_io_p_Stdin struct {}

type _goml_m_std_p_io_p_StringReader struct {
    value *ref_string_x
    offset *ref_int_x
    previous_char *ref_Option__isize_x
}

type _goml_m_std_p_io_p_CopyError struct {
    read_value uint64
    written_value uint64
    cause _goml_m_std_p_io_p_Error
}

type _goml_m_std_p_io_p_ReadFragment struct {
    bytes_value []uint8
    end_value _goml_m_std_p_io_p_FragmentEnd
    consumed_value int
}

type _goml_m_std_p_io_p_FragmentError struct {
    bytes_value []uint8
    error_value _goml_m_std_p_io_p_Error
}

type _goml_m_std_p_io_p_PipePacket struct {
    input []uint8
    consumed int
    reply chan _goml_m_Result____isize____std_p_io_p_TransferError
}

type _goml_m_std_p_io_p_PipeState struct {
    read_closed _goml_m_Option____std_p_io_p_Error
    write_closed _goml_m_Option____Option____std_p_io_p_Error
    active _goml_m_Option____std_p_io_p_PipePacket
    changed chan struct{}
}

type _goml_m_std_p_io_p_PipeReader struct {
    state chan _goml_m_std_p_io_p_PipeState
}

type _goml_m_std_p_io_p_PipeWriter struct {
    state chan _goml_m_std_p_io_p_PipeState
}

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
    result *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
    ready chan struct{}
}

type closure_env_main_0 struct {}

type closure_env_main_1 struct {}

type closure_env_main_2 struct {}

type closure_env_std_task_scope_T_3 struct {
    result_0 *ref__goml_m_Option_____o__q__x
    body_1 func(_goml_m_std_p_task_p_Scope) struct{}
    handle_2 _goml_m_std_p_internal_p_task_p_ScopeHandle
}

type closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_4 struct {
    body_0 func(_goml_m_std_p_task_p_CancelToken) int
    token_1 _goml_m_std_p_task_p_CancelToken
}

type closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_5 struct {
    body_0 func(_goml_m_std_p_task_p_CancelToken) int
    token_1 _goml_m_std_p_task_p_CancelToken
    result_2 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
    ready_3 chan struct{}
}

type closure_env_std_panic_catch_T_isize_6 struct {
    result_0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
    body_1 func() int
}

type closure_env_std_panic_catch_T_isize_7 struct {
    result_0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x
}

type FrozenVec__u8 struct {
    values *_goml_vec_uint8
}

type Ordering uint8

type _goml_m_std_p_unicode_p_TableError struct {
    _p0 int
    _tag uint8
}

type _goml_m_std_p_unicode_p_Case uint8

type _goml_m_std_p_bytes_p_TransformError struct {
    _p0 int
    _tag uint8
}

type _goml_m_std_p_utf8_p_Prefix struct {
    _p1 int
    _p0 rune
    _tag uint8
}

type _goml_m_std_p_utf8_p_DecoderError interface {
    is_goml_m_std_p_utf8_p_DecoderError()
}

type InvalidUtf8 struct {
    _0 _goml_m_std_p_utf8_p_Utf8Error
}

func (_ InvalidUtf8) is_goml_m_std_p_utf8_p_DecoderError() {}

type Finished struct {}

func (_ Finished) is_goml_m_std_p_utf8_p_DecoderError() {}

type OffsetOverflow struct {}

func (_ OffsetOverflow) is_goml_m_std_p_utf8_p_DecoderError() {}

type _goml_m_std_p_io_p_ErrorKind uint8

type _goml_m_std_p_io_p_ScanStep interface {
    is_goml_m_std_p_io_p_ScanStep()
}

type More struct {
    _0 int
}

func (_ More) is_goml_m_std_p_io_p_ScanStep() {}

type Token struct {
    _0 int
    _1 int
    _2 int
}

func (_ Token) is_goml_m_std_p_io_p_ScanStep() {}

type Emit struct {
    _0 int
    _1 _goml_m_std_p_bytes_p_Bytes
}

func (_ Emit) is_goml_m_std_p_io_p_ScanStep() {}

type Final struct {
    _0 _goml_m_Option_____o_isize_c_isize_q_
}

func (_ Final) is_goml_m_std_p_io_p_ScanStep() {}

type _goml_m_std_p_io_p_SeekFrom struct {
    _p0 int
    _tag uint8
}

type _goml_m_std_p_io_p_FragmentEnd uint8

type _goml_m_Option_____o__q_ uint8

const (
    _goml_m_Option_____o__q__None _goml_m_Option_____o__q_ = 0
    _goml_m_Option_____o__q__Some _goml_m_Option_____o__q_ = 1
)

type _goml_m_Result____isize____std_p_panic_p_Panic interface {
    is_goml_m_Result____isize____std_p_panic_p_Panic()
}

type _goml_m_Result____isize____std_p_panic_p_Panic_Ok struct {
    _0 int
}

func (_ _goml_m_Result____isize____std_p_panic_p_Panic_Ok) is_goml_m_Result____isize____std_p_panic_p_Panic() {}

type _goml_m_Result____isize____std_p_panic_p_Panic_Err struct {
    _0 _goml_m_std_p_panic_p_Panic
}

func (_ _goml_m_Result____isize____std_p_panic_p_Panic_Err) is_goml_m_Result____isize____std_p_panic_p_Panic() {}

type _goml_m_Option____Result____isize____std_p_panic_p_Panic struct {
    _p0 _goml_m_Result____isize____std_p_panic_p_Panic
    _tag uint8
}

type _goml_m_Option_____o_isize_c_isize_q_ struct {
    _p0 Tuple2_3int_3int
    _tag uint8
}

type Option__isize struct {
    _p0 int
    _tag uint8
}

type _goml_m_Option____std_p_utf8_p_Utf8Error struct {
    _p0 _goml_m_std_p_utf8_p_Utf8Error
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

type _goml_m_Result____isize____std_p_io_p_TransferError interface {
    is_goml_m_Result____isize____std_p_io_p_TransferError()
}

type _goml_m_Result____isize____std_p_io_p_TransferError_Ok struct {
    _0 int
}

func (_ _goml_m_Result____isize____std_p_io_p_TransferError_Ok) is_goml_m_Result____isize____std_p_io_p_TransferError() {}

type _goml_m_Result____isize____std_p_io_p_TransferError_Err struct {
    _0 _goml_m_std_p_io_p_TransferError
}

func (_ _goml_m_Result____isize____std_p_io_p_TransferError_Err) is_goml_m_Result____isize____std_p_io_p_TransferError() {}

type _goml_m_Option____std_p_io_p_Error interface {
    is_goml_m_Option____std_p_io_p_Error()
}

type _goml_m_Option____std_p_io_p_Error_None struct {}

func (_ _goml_m_Option____std_p_io_p_Error_None) is_goml_m_Option____std_p_io_p_Error() {}

type _goml_m_Option____std_p_io_p_Error_Some struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Option____std_p_io_p_Error_Some) is_goml_m_Option____std_p_io_p_Error() {}

type _goml_m_Option____Option____std_p_io_p_Error struct {
    _p0 _goml_m_Option____std_p_io_p_Error
    _tag uint8
}

type _goml_m_Option____std_p_io_p_PipePacket struct {
    _p0 _goml_m_std_p_io_p_PipePacket
    _tag uint8
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
    var t5 closure_env_main_1 = closure_env_main_1{}
    var t6 func(_goml_m_std_p_task_p_Scope) struct{} = func(p0 _goml_m_std_p_task_p_Scope) struct{} {
        return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t5, p0)
    }
    var inline6 _goml_m_std_p_internal_p_task_p_ScopeHandle = _goml_m_std_p_internal_p_task_p_root__scope()
    var inline7 *ref__goml_m_Option_____o__q__x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Option_l__o__q__r_(_goml_m_Option_____o__q__None)
    var inline8 closure_env_std_task_scope_T_3 = closure_env_std_task_scope_T_3{
        result_0: inline7,
        body_1: t6,
        handle_2: inline6,
    }
    var inline9 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h78f773e7d4b2cfe5d41554b5624d043f_e__T__3_i_apply(inline8)
    }
    _goml_m_std_p_internal_p_task_p_run(inline6, inline9)
    _goml_m_std_p_task_p_completed__scope__value____T___o__q_(inline7)
    var t7 closure_env_main_2 = closure_env_main_2{}
    var t8 func(_goml_m_std_p_task_p_Scope) struct{} = func(p0 _goml_m_std_p_task_p_Scope) struct{} {
        return _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(t7, p0)
    }
    var inline0 _goml_m_std_p_internal_p_task_p_ScopeHandle = _goml_m_std_p_internal_p_task_p_root__scope()
    var inline1 *ref__goml_m_Option_____o__q__x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Option_l__o__q__r_(_goml_m_Option_____o__q__None)
    var inline2 closure_env_std_task_scope_T_3 = closure_env_std_task_scope_T_3{
        result_0: inline1,
        body_1: t8,
        handle_2: inline0,
    }
    var inline3 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h78f773e7d4b2cfe5d41554b5624d043f_e__T__3_i_apply(inline2)
    }
    _goml_m_std_p_internal_p_task_p_run(inline0, inline3)
    _goml_m_std_p_task_p_completed__scope__value____T___o__q_(inline1)
    return struct{}{}
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

func _goml_m_std_p_internal_p_host_p_timer__new(value__0 int64) Tuple2_5int64_14Receiver_4unit {
    var t0 Tuple2_5int64_14Receiver_4unit = _goml_runtime_std_time_timer_new(value__0)
    return t0
}

func _goml_m_inherent_i_std_p_time__h3ebf4d33c01c8036463b8efe535bf1db_as__nanoseconds(self__0 _goml_m_std_p_time_p_Duration) int64 {
    var t0 int64 = self__0.nanoseconds
    return t0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_std_p_internal_p_host_p_timer__stop(id__0 int64) bool {
    var t0 bool = _goml_runtime_std_time_timer_stop(id__0)
    return t0
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

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Option_l__o__q__r_(value__0 _goml_m_Option_____o__q_) *ref__goml_m_Option_____o__q__x {
    var t0 *ref__goml_m_Option_____o__q__x = ref___goml_m_Ref__10Option_____o__q_(value__0)
    return t0
}

func _goml_m_std_p_internal_p_task_p_run(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle, body__0 func() struct{}) struct{} {
    var t0 int64 = scope__0.id
    _goml_runtime_std_task_scope_run(t0, body__0)
    var t1 int64 = scope__0.id
    _goml_runtime_std_task_scope_finish(t1)
    return struct{}{}
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

func _goml_m_inherent_i_Ref_i_Ref_l_h36f65a4e1ce11b8891da0ef266111b3e_c_p_Panic_r__r_(value__0 _goml_m_Option____Result____isize____std_p_panic_p_Panic) *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x {
    var t0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = ref___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(value__0)
    return t0
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T___o__q_(capacity__0 int) chan struct{} {
    var t0 chan struct{} = func(p0 int) chan struct{} {
        return make(chan struct{}, p0)
    }(capacity__0)
    return t0
}

func _goml_m_std_p_internal_p_task_p_token(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle) _goml_m_std_p_internal_p_task_p_CancelToken {
    var t0 int64 = scope__0.id
    var t1 _goml_m_std_p_internal_p_task_p_CancelToken = _goml_m_std_p_internal_p_task_p_CancelToken{
        id: t0,
    }
    return t1
}

func _goml_m_std_p_internal_p_task_p_spawn(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle, body__0 func() struct{}) struct{} {
    var t0 int64 = scope__0.id
    _goml_runtime_std_task_scope_spawn(t0, body__0)
    return struct{}{}
}

func _goml_m_std_p_panic_p_resume(info__0 _goml_m_std_p_panic_p_Panic) struct{} {
    var t0 func() struct{} = info__0.rethrow
    t0()
    panic("unreachable")
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T___o__q_(self__0 chan struct{}) <-chan struct{} {
    var t0 <-chan struct{} = func(p0 chan struct{}) <-chan struct{} {
        return p0
    }(self__0)
    return t0
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

func _goml_m_std_p_internal_p_task_p_cancel(scope__0 _goml_m_std_p_internal_p_task_p_ScopeHandle) struct{} {
    var t0 int64 = scope__0.id
    _goml_runtime_std_task_scope_cancel(t0)
    return struct{}{}
}

func _goml_m_inherent_i_std_p_inter_h25a0a7c69f0f67d3e38fa37e5f2fc20e_celToken_i_done(self__0 _goml_m_std_p_internal_p_task_p_CancelToken) <-chan struct{} {
    var t0 int64 = self__0.id
    var inline0 <-chan struct{} = _goml_runtime_std_task_scope_done(t0)
    return inline0
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

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0, ___0 _goml_m_std_p_task_p_CancelToken) int {
    return 42
}

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1, scope__0 _goml_m_std_p_task_p_Scope) struct{} {
    var t0 closure_env_main_0 = closure_env_main_0{}
    var t1 func(_goml_m_std_p_task_p_CancelToken) int = func(p0 _goml_m_std_p_task_p_CancelToken) int {
        return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t0, p0)
    }
    var work__0 _goml_m_std_p_task_p_Task____isize
    var inline13 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = _goml_m_inherent_i_Ref_i_Ref_l_h36f65a4e1ce11b8891da0ef266111b3e_c_p_Panic_r__r_(_goml_m_Option____Result____isize____std_p_panic_p_Panic{
        _tag: 0,
    })
    var inline14 chan struct{} = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T___o__q_(0)
    var inline15 _goml_m_std_p_internal_p_task_p_ScopeHandle = scope__0.handle
    var inline16 _goml_m_std_p_internal_p_task_p_CancelToken = _goml_m_std_p_internal_p_task_p_token(inline15)
    var inline17 _goml_m_std_p_task_p_CancelToken = _goml_m_std_p_task_p_CancelToken{
        value: inline16,
    }
    var inline18 _goml_m_std_p_internal_p_task_p_ScopeHandle = scope__0.handle
    var inline19 closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_5 = closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_5{
        body_0: t1,
        token_1: inline17,
        result_2: inline13,
        ready_3: inline14,
    }
    var inline20 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h928008e92c5c30e99d0f6f5d21779a91_size__5_i_apply(inline19)
    }
    _goml_m_std_p_internal_p_task_p_spawn(inline18, inline20)
    var inline22 _goml_m_std_p_task_p_Task____isize = _goml_m_std_p_task_p_Task____isize{
        result: inline13,
        ready: inline14,
    }
    work__0 = inline22
    var t2 <-chan struct{}
    var inline11 chan struct{} = work__0.ready
    var inline12 <-chan struct{} = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T___o__q_(inline11)
    t2 = inline12
    var _goml_m_______0_i_select__open bool
    select {
    case _, _goml_m_______0_i_select__open = <-t2:
        if _goml_m_______0_i_select__open {}
        var t3 int
        var inline2 chan struct{} = work__0.ready
        _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_recv____T___o__q_(inline2)
        var inline4 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = work__0.result
        var inline5 _goml_m_Option____Result____isize____std_p_panic_p_Panic = _goml_m_inherent_i_Ref_i_Ref_l_h80255cab1de9a17c566649cce0d96ab2_c_p_Panic_r__r_(inline4)
        switch inline5._tag {
        case 0:
            _goml_m_std_p_panic_p_raise("task completed without a result")
            panic("unreachable")
        case 1:
            var inline7 _goml_m_Result____isize____std_p_panic_p_Panic = inline5._p0
            switch inline7.(type) {
            case _goml_m_Result____isize____std_p_panic_p_Panic_Ok:
                var inline8 int = inline7.(_goml_m_Result____isize____std_p_panic_p_Panic_Ok)._0
                t3 = inline8
                var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t3)
                _goml_runtime_core_string_println(inline0)
                return struct{}{}
            case _goml_m_Result____isize____std_p_panic_p_Panic_Err:
                var inline9 _goml_m_std_p_panic_p_Panic = inline7.(_goml_m_Result____isize____std_p_panic_p_Panic_Err)._0
                _goml_m_std_p_panic_p_resume(inline9)
                panic("unreachable")
            default:
                panic("non-exhaustive match")
            }
        default:
            panic("non-exhaustive match")
        }
    }
}

func _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(env0 closure_env_main_2, scope__0 _goml_m_std_p_task_p_Scope) struct{} {
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

func _goml_m_inherent_i_closure__en_h78f773e7d4b2cfe5d41554b5624d043f_e__T__3_i_apply(env0 closure_env_std_task_scope_T_3) struct{} {
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

func _goml_m_inherent_i_closure__en_h06a158548ef8dc096a001bc90e690a25_size__4_i_apply(env0 closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_4) int {
    var body__0 func(_goml_m_std_p_task_p_CancelToken) int = env0.body_0
    var token__0 _goml_m_std_p_task_p_CancelToken = env0.token_1
    var t0 int = body__0(token__0)
    return t0
}

func _goml_m_inherent_i_closure__en_h928008e92c5c30e99d0f6f5d21779a91_size__5_i_apply(env0 closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_5) struct{} {
    var body__0 func(_goml_m_std_p_task_p_CancelToken) int = env0.body_0
    var token__0 _goml_m_std_p_task_p_CancelToken = env0.token_1
    var result__0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = env0.result_2
    var ready__0 chan struct{} = env0.ready_3
    var t0 closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_4 = closure_env_inherent_std_task_Scope_std_task_Scope_spawn_T_isize_4{
        body_0: body__0,
        token_1: token__0,
    }
    var t1 func() int = func() int {
        return _goml_m_inherent_i_closure__en_h06a158548ef8dc096a001bc90e690a25_size__4_i_apply(t0)
    }
    var outcome__0 _goml_m_Result____isize____std_p_panic_p_Panic
    var inline4 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = _goml_m_inherent_i_Ref_i_Ref_l_h36f65a4e1ce11b8891da0ef266111b3e_c_p_Panic_r__r_(_goml_m_Option____Result____isize____std_p_panic_p_Panic{
        _tag: 0,
    })
    var inline5 closure_env_std_panic_catch_T_isize_6 = closure_env_std_panic_catch_T_isize_6{
        result_0: inline4,
        body_1: t1,
    }
    var inline6 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h97b804dd1298dda5454e9c6ab33fd749_size__6_i_apply(inline5)
    }
    var inline7 closure_env_std_panic_catch_T_isize_7 = closure_env_std_panic_catch_T_isize_7{
        result_0: inline4,
    }
    var inline8 func(string, string, func() struct{}) struct{} = func(p0 string, p1 string, p2 func() struct{}) struct{} {
        return _goml_m_inherent_i_closure__en_h0a05d186557f06626dc1823ff592bceb_size__7_i_apply(inline7, p0, p1, p2)
    }
    _goml_runtime_std_panic_catch(inline6, inline8)
    var inline10 _goml_m_Option____Result____isize____std_p_panic_p_Panic = _goml_m_inherent_i_Ref_i_Ref_l_h80255cab1de9a17c566649cce0d96ab2_c_p_Panic_r__r_(inline4)
    switch inline10._tag {
    case 0:
        _goml_m_std_p_panic_p_raise("panic boundary returned without a result")
        panic("unreachable")
    case 1:
        var inline12 _goml_m_Result____isize____std_p_panic_p_Panic = inline10._p0
        outcome__0 = inline12
        var t2 _goml_m_Option____Result____isize____std_p_panic_p_Panic = _goml_m_Option____Result____isize____std_p_panic_p_Panic{
            _p0: outcome__0,
            _tag: 1,
        }
        ref_set___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(result__0, t2)
        func(p0 chan struct{}) struct{} {
            close(p0)
            return struct{}{}
        }(ready__0)
        switch outcome__0.(type) {
        case _goml_m_Result____isize____std_p_panic_p_Panic_Err:
            var x0 _goml_m_std_p_panic_p_Panic = outcome__0.(_goml_m_Result____isize____std_p_panic_p_Panic_Err)._0
            var inline0 func() struct{} = x0.rethrow
            inline0()
            panic("unreachable")
        default:
            return struct{}{}
        }
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_closure__en_h97b804dd1298dda5454e9c6ab33fd749_size__6_i_apply(env0 closure_env_std_panic_catch_T_isize_6) struct{} {
    var result__0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = env0.result_0
    var body__0 func() int = env0.body_1
    var t0 int = body__0()
    var t1 _goml_m_Result____isize____std_p_panic_p_Panic = _goml_m_Result____isize____std_p_panic_p_Panic_Ok{
        _0: t0,
    }
    var t2 _goml_m_Option____Result____isize____std_p_panic_p_Panic = _goml_m_Option____Result____isize____std_p_panic_p_Panic{
        _p0: t1,
        _tag: 1,
    }
    ref_set___goml_m_Ref__40Option____Result____isize____std_p_panic_p_Panic(result__0, t2)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_h0a05d186557f06626dc1823ff592bceb_size__7_i_apply(env0 closure_env_std_panic_catch_T_isize_7, message__0 string, stack__0 string, rethrow__0 func() struct{}) struct{} {
    var result__0 *ref__goml_m_Option____Result____isize____std_p_panic_p_Panic_x = env0.result_0
    var t0 _goml_m_std_p_panic_p_Panic = _goml_m_std_p_panic_p_Panic{
        message: message__0,
        stack: stack__0,
        rethrow: rethrow__0,
    }
    var t1 _goml_m_Result____isize____std_p_panic_p_Panic = _goml_m_Result____isize____std_p_panic_p_Panic_Err{
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
