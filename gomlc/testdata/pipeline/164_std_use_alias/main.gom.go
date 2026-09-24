package main

import (
    _goml_os "os"
)

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

type _goml_vec_uint32 struct {
    items []uint32
}

type _goml_vec_uint8 struct {
    items []uint8
}

type _goml_vec_int struct {
    items []int
}

type _goml_vec__goml_m_std_p_text_p_LineIndexWideChar struct {
    items []_goml_m_std_p_text_p_LineIndexWideChar
}

type _goml_vec__goml_m_std_p_text_p_ReplacementNode struct {
    items []_goml_m_std_p_text_p_ReplacementNode
}

type _goml_vec_Tuple2_6string_6string struct {
    items []Tuple2_6string_6string
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

type hashmap_uint8_int_x_entry struct {
    active bool
    key uint8
    value int
}

type hashmap_uint8_int_x struct {
    indices map[uint8]int
    entries []hashmap_uint8_int_x_entry
    len int
}

type Tuple2_6string_6string struct {
    _0 string
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

type _goml_m_std_p_text_p_Finder struct {
    pattern string
    fallback *_goml_vec_int
}

type _goml_m_std_p_text_p_ReplacementNode struct {
    edges *hashmap_uint8_int_x
    rule int
}

type _goml_m_std_p_text_p_Replacer struct {
    nodes *_goml_vec__goml_m_std_p_text_p_ReplacementNode
    rules *_goml_vec_Tuple2_6string_6string
}

type _goml_m_std_p_text_p_ReplacementCursor struct {
    replacer _goml_m_std_p_text_p_Replacer
    value string
    max_bytes int
    max_work int
    work *ref_int_x
    used *ref_int_x
    offset *ref_int_x
    ignore_empty *ref_bool_x
}

type _goml_m_std_p_env_p_VarError struct {
    details _goml_m_std_p_io_p_ErrorDetails
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

type _goml_m_std_p_utf8_p_DecoderError_InvalidUtf8 struct {
    _0 _goml_m_std_p_utf8_p_Utf8Error
}

func (_ _goml_m_std_p_utf8_p_DecoderError_InvalidUtf8) is_goml_m_std_p_utf8_p_DecoderError() {}

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

type _goml_m_std_p_text_p_PositionEncoding uint8

type _goml_m_std_p_text_p_UnquoteError struct {
    _p0 int
    _tag uint8
}

type _goml_m_std_p_text_p_QuoteContext uint8

type _goml_m_std_p_text_p_QuotedElement struct {
    _p1 rune
    _p0 uint8
    _tag uint8
}

type _goml_m_std_p_text_p_QuoteMode uint8

type _goml_m_std_p_text_p_ReplaceError struct {
    _p0 int
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

type Ok struct {
    _0 int
}

func (_ Ok) is_goml_m_Result____isize____std_p_io_p_TransferError() {}

type Err struct {
    _0 _goml_m_std_p_io_p_TransferError
}

func (_ Err) is_goml_m_Result____isize____std_p_io_p_TransferError() {}

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

func _goml_m_std_p_internal_p_host_p_println(value__0 string) struct{} {
    _goml_runtime_std_io_println(value__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_std_p_internal_p_host_p_args() *_goml_vec_string {
    var t0 *_goml_vec_string = _goml_runtime_std_env_args()
    return t0
}

func main() {
    main0()
}
