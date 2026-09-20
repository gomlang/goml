package main

import (
    _goml_context "context"
    _goml_os "os"
    _goml_sync "sync"
)

type _goml_task_scope_state struct {
    mu _goml_sync.Mutex
    wg _goml_sync.WaitGroup
    state int
    ctx _goml_context.Context
    cancel _goml_context.CancelFunc
    panicked bool
    panic_value any
}

func _goml_runtime_core_bool_to_string(x bool) string {
    if x {
        return "true"
    } else {
        return "false"
    }
}

func _goml_runtime_std_env_args() *_goml_vec_string {
    return &_goml_vec_string{
        items: _goml_os.Args,
    }
}

func _goml_runtime_std_io_println(value string) struct{} {
    _goml_os.Stdout.WriteString(value + "\n")
    return struct{}{}
}

type _goml_vec_string struct {
    items []string
}

func vec_len__Vec_6string(vec *_goml_vec_string) int {
    return int(len(vec.items))
}

type _goml_vec_uint8 struct {
    items []uint8
}

type _goml_vec_Tuple2_6string_6string struct {
    items []Tuple2_6string_6string
}

type _goml_vec_int struct {
    items []int
}

type _goml_vec__goml_m_std_p_text_p_LineIndexWideChar struct {
    items []_goml_m_std_p_text_p_LineIndexWideChar
}

type _goml_vec_uint32 struct {
    items []uint32
}

type ref_int_x struct {
    value int
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

type Tuple2_13Option__isize_13Option__isize struct {
    _0 Option__isize
    _1 Option__isize
}

type Tuple2_3int_4char struct {
    _0 int
    _1 rune
}

type Tuple2_4char_3int struct {
    _0 rune
    _1 int
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

type _goml_m_std_p_text_p_LineColumn struct {
    line int
    column int
}

type _goml_m_std_p_text_p_LineIndexWideChar struct {
    start int
    end int
}

type _goml_m_std_p_text_p_LineIndex struct {
    line_starts *_goml_vec_int
    line_ends *_goml_vec_int
    wide_offsets *_goml_vec_int
    wide_chars *_goml_vec__goml_m_std_p_text_p_LineIndexWideChar
    source_length int
}

type _goml_m_std_p_text_p_StringBuilder struct {
    values *_goml_vec_uint8
}

type _goml_m_std_p_env_p_VarError struct {
    details _goml_m_std_p_io_p_ErrorDetails
}

type _goml_m_FnIterator_____o_isize_c_char_q_ struct {
    next_fn func() _goml_m_Option_____o_isize_c_char_q_
}

type FnIterator__u8 struct {
    next_fn func() Option__u8
}

type FnIterator__char struct {
    next_fn func() Option__char
}

type FnIterator__isize struct {
    next_fn func() Option__isize
}

type closure_env_std_io_read_stdin_to_string_0 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Cursor_1 struct {}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_2 struct {}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_3 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Stdin_4 struct {}

type closure_env_inherent_string_string_char_indices_5 struct {
    index_0 *ref_int_x
    self_1 string
}

type closure_env_inherent_string_string_chars_6 struct {
    self_0 string
    index_1 *ref_int_x
}

type closure_env_goml_builtin_range_7 struct {
    current_0 *ref_int_x
    end_1 int
}

type closure_env_inherent_Slice_Slice_T_iter_T_u8_8 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 []uint8
}

type FrozenVec__u8 struct {
    values *_goml_vec_uint8
}

type Ordering uint8

type _goml_m_std_p_io_p_ErrorKind uint8

type _goml_m_std_p_text_p_PositionEncoding uint8

type Option__u8 uint16

type _goml_m_Option____Slice_l_u8_r_ struct {
    _p0 []uint8
    _tag uint8
}

type _goml_m_Option____MutSlice_l_u8_r_ struct {
    _p0 []uint8
    _tag uint8
}

type _goml_m_Option____std_p_bytes_p_FrozenBytes struct {
    _p0 _goml_m_std_p_bytes_p_FrozenBytes
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

type _goml_m_Option_____o_isize_c_char_q_ struct {
    _p0 Tuple2_3int_4char
    _tag uint8
}

type _goml_m_Option_____o_string_c_string_q_ struct {
    _p0 Tuple2_6string_6string
    _tag uint8
}

type _goml_m_Result____string____std_p_env_p_VarError interface {
    is_goml_m_Result____string____std_p_env_p_VarError()
}

type _goml_m_Result____string____std_p_env_p_VarError_Ok struct {
    _0 string
}

func (_ _goml_m_Result____string____std_p_env_p_VarError_Ok) is_goml_m_Result____string____std_p_env_p_VarError() {}

type _goml_m_Result____string____std_p_env_p_VarError_Err struct {
    _0 _goml_m_std_p_env_p_VarError
}

func (_ _goml_m_Result____string____std_p_env_p_VarError_Err) is_goml_m_Result____string____std_p_env_p_VarError() {}

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

type _goml_m_Option_____o_char_c_isize_q_ struct {
    _p0 Tuple2_4char_3int
    _tag uint8
}

type Option__char uint64

func _goml_m_std_p_internal_p_host_p_args() *_goml_vec_string {
    var t0 *_goml_vec_string = _goml_runtime_std_env_args()
    return t0
}

func _goml_m_std_p_internal_p_host_p_println(value__0 string) struct{} {
    _goml_runtime_std_io_println(value__0)
    return struct{}{}
}

func main0() struct{} {
    var t0 *_goml_vec_string
    var inline4 *_goml_vec_string = _goml_m_std_p_internal_p_host_p_args()
    t0 = inline4
    var t1 int
    var inline3 int = vec_len__Vec_6string(t0)
    t1 = inline3
    var t2 bool = t1 > 0
    var t3 string
    var inline2 string = _goml_runtime_core_bool_to_string(t2)
    t3 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_m_std_p_internal_p_host_p_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
