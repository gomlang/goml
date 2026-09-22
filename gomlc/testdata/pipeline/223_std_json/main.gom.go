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

func _goml_runtime_core_string_len(s string) int {
    return int(len(s))
}

func _goml_runtime_core_string_get(s string, i int) rune {
    return rune(s[i])
}

func _goml_runtime_core_string_byte_get(s string, i int) uint8 {
    return s[i]
}

func _goml_runtime_core_string_byte_slice(s string, start int, end int) string {
    return s[start:end]
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_char_to_string(x rune) string {
    return string(x)
}

func _goml_runtime_core_char_from_uint32(value uint32) Tuple2_4bool_4char {
    return Tuple2_4bool_4char{
        _0: true,
        _1: rune(value),
    }
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_string struct {
    items []string
}

type _goml_vec_uint8 struct {
    items []uint8
}

func vec_new__Vec_5uint8() *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: nil,
    }
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

func vec_reserve__Vec_5uint8(vec *_goml_vec_uint8, additional int) struct{} {
    if additional < 0 {
        panic("negative vector capacity")
    }
    var length int = len(vec.items)
    var required int = length + additional
    if required < length {
        panic("vector capacity overflow")
    }
    if required > cap(vec.items) {
        var next_capacity int = cap(vec.items) * 2
        if next_capacity < required {
            next_capacity = required
        }
        var next_items []uint8 = make([]uint8, length, next_capacity)
        copy(next_items, vec.items)
        vec.items = next_items
    }
    return struct{}{}
}

type _goml_vec_Tuple2_6string_6string struct {
    items []Tuple2_6string_6string
}

type _goml_vec__goml_m_std_p_unicode_p_Range struct {
    items []_goml_m_std_p_unicode_p_Range
}

type _goml_vec__goml_m_std_p_unicode_p_RangeTable struct {
    items []_goml_m_std_p_unicode_p_RangeTable
}

type _goml_vec__goml_m_std_p_unicode_p_CaseMapping struct {
    items []_goml_m_std_p_unicode_p_CaseMapping
}

type _goml_vec_int struct {
    items []int
}

type _goml_vec_Slice_5uint8 struct {
    items [][]uint8
}

type _goml_vec__goml_m_Tuple2__6string__18std_p_serde_p_Schema struct {
    items []Tuple2_6string_28_goml_m_std_p_serde_p_Schema
}

type _goml_vec__goml_m_std_p_serde_p_VariantSchema struct {
    items []_goml_m_std_p_serde_p_VariantSchema
}

type _goml_vec__goml_m_Tuple2__6string__17std_p_serde_p_Value struct {
    items []Tuple2_6string_27_goml_m_std_p_serde_p_Value
}

type _goml_vec__goml_m_std_p_serde_p_Value struct {
    items []_goml_m_std_p_serde_p_Value
}

type _goml_vec__goml_m_std_p_serde_p_ValueDeserializeFrame struct {
    items []_goml_m_std_p_serde_p_ValueDeserializeFrame
}

type _goml_vec__goml_m_Tuple2__17std_p_serde_p_Value__17std_p_serde_p_Value struct {
    items []Tuple2_27_goml_m_std_p_serde_p_Value_27_goml_m_std_p_serde_p_Value
}

type _goml_vec__goml_m_std_p_serde_p_ValueSerializeFrame struct {
    items []_goml_m_std_p_serde_p_ValueSerializeFrame
}

type _goml_vec__goml_m_std_p_text_p_LineIndexWideChar struct {
    items []_goml_m_std_p_text_p_LineIndexWideChar
}

type _goml_vec__goml_m_std_p_text_p_ReplacementNode struct {
    items []_goml_m_std_p_text_p_ReplacementNode
}

type _goml_vec__goml_m_std_p_json_p_JsonDeserializeFrame struct {
    items []_goml_m_std_p_json_p_JsonDeserializeFrame
}

type _goml_vec__goml_m_std_p_json_p_Value struct {
    items []_goml_m_std_p_json_p_Value
}

func vec_new___goml_m_Vec__16std_p_json_p_Value() *_goml_vec__goml_m_std_p_json_p_Value {
    return &_goml_vec__goml_m_std_p_json_p_Value{
        items: nil,
    }
}

func vec_push___goml_m_Vec__16std_p_json_p_Value(vec *_goml_vec__goml_m_std_p_json_p_Value, elem _goml_m_std_p_json_p_Value) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get___goml_m_Vec__16std_p_json_p_Value(vec *_goml_vec__goml_m_std_p_json_p_Value, index int) _goml_m_std_p_json_p_Value {
    return vec.items[index]
}

func vec_len___goml_m_Vec__16std_p_json_p_Value(vec *_goml_vec__goml_m_std_p_json_p_Value) int {
    return int(len(vec.items))
}

type _goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value struct {
    items []Tuple2_6string_26_goml_m_std_p_json_p_Value
}

func vec_new___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value() *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value {
    return &_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value{
        items: nil,
    }
}

func vec_push___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value(vec *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value, elem Tuple2_6string_26_goml_m_std_p_json_p_Value) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value(vec *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value, index int) Tuple2_6string_26_goml_m_std_p_json_p_Value {
    return vec.items[index]
}

func vec_len___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value(vec *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value) int {
    return int(len(vec.items))
}

type _goml_vec__goml_m_std_p_json_p_JsonSerializeFrame struct {
    items []_goml_m_std_p_json_p_JsonSerializeFrame
}

type _goml_vec_uint32 struct {
    items []uint32
}

type _goml_vec__goml_m_std_p_serde_p_Schema struct {
    items []_goml_m_std_p_serde_p_Schema
}

type ref_int_x struct {
    value int
}

func ref__Ref_3int(value int) *ref_int_x {
    return &ref_int_x{
        value: value,
    }
}

func ref_get__Ref_3int(reference *ref_int_x) int {
    return reference.value
}

func ref_set__Ref_3int(reference *ref_int_x, value int) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_bool_x struct {
    value bool
}

type ref__goml_m_Option____std_p_utf8_p_Utf8Error_x struct {
    value _goml_m_Option____std_p_utf8_p_Utf8Error
}

type ref_string_x struct {
    value string
}

type ref_Option__isize_x struct {
    value Option__isize
}

type ref__goml_m_Option____std_p_serde_p_Value_x struct {
    value _goml_m_Option____std_p_serde_p_Value
}

type ref_Option__string_x struct {
    value Option__string
}

type hashmap_char_bool_x_entry struct {
    active bool
    key rune
    value bool
}

type hashmap_char_bool_x struct {
    indices map[rune]int
    entries []hashmap_char_bool_x_entry
    len int
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

type Tuple2_12Slice_5uint8_12Slice_5uint8 struct {
    _0 []uint8
    _1 []uint8
}

type Tuple2_4char_3int struct {
    _0 rune
    _1 int
}

type Tuple2_3int_3int struct {
    _0 int
    _1 int
}

type Tuple3_4bool_4char_3int struct {
    _0 bool
    _1 rune
    _2 int
}

type Tuple2_29_goml_m_std_p_io_p_PipeReader_29_goml_m_std_p_io_p_PipeWriter struct {
    _0 _goml_m_std_p_io_p_PipeReader
    _1 _goml_m_std_p_io_p_PipeWriter
}

type Tuple2_28_goml_m_std_p_io_p_PipeState_4bool struct {
    _0 _goml_m_std_p_io_p_PipeState
    _1 bool
}

type Tuple2_4unit_4bool struct {
    _0 struct{}
    _1 bool
}

type Tuple2_51_goml_m_Result____isize____std_p_io_p_TransferError_4bool struct {
    _0 _goml_m_Result____isize____std_p_io_p_TransferError
    _1 bool
}

type Tuple4_4bool_3int_3int_4bool struct {
    _0 bool
    _1 int
    _2 int
    _3 bool
}

type Tuple3_4bool_6uint64_6string struct {
    _0 bool
    _1 uint64
    _2 string
}

type Tuple3_4bool_4uint_6string struct {
    _0 bool
    _1 uint
    _2 string
}

type Tuple2_4bool_7float64 struct {
    _0 bool
    _1 float64
}

type Tuple2_4bool_6uint64 struct {
    _0 bool
    _1 uint64
}

type Tuple2_10Vec_5uint8_3int struct {
    _0 *_goml_vec_uint8
    _1 int
}

type Tuple2_4bool_4bool struct {
    _0 bool
    _1 bool
}

type Tuple2_7float64_7float64 struct {
    _0 float64
    _1 float64
}

type Tuple2_7float32_7float32 struct {
    _0 float32
    _1 float32
}

type Tuple2_6string_28_goml_m_std_p_serde_p_Schema struct {
    _0 string
    _1 _goml_m_std_p_serde_p_Schema
}

type Tuple2_6string_27_goml_m_std_p_serde_p_Value struct {
    _0 string
    _1 _goml_m_std_p_serde_p_Value
}

type Tuple3_6string_3int_50Vec_44Tuple2_6string_27_goml_m_std_p_serde_p_Value struct {
    _0 string
    _1 int
    _2 *_goml_vec__goml_m_Tuple2__6string__17std_p_serde_p_Value
}

type Tuple2_4int8_10Vec_5uint8 struct {
    _0 int8
    _1 *_goml_vec_uint8
}

type Tuple2_27_goml_m_std_p_serde_p_Value_27_goml_m_std_p_serde_p_Value struct {
    _0 _goml_m_std_p_serde_p_Value
    _1 _goml_m_std_p_serde_p_Value
}

type Tuple2_3int_4char struct {
    _0 int
    _1 rune
}

type Tuple4_3int_3int_6uint32_4bool struct {
    _0 int
    _1 int
    _2 uint32
    _3 bool
}

type Tuple2_34_goml_m_std_p_text_p_QuotedElement_3int struct {
    _0 _goml_m_std_p_text_p_QuotedElement
    _1 int
}

type Tuple2_34_goml_m_std_p_text_p_QuotedElement_6string struct {
    _0 _goml_m_std_p_text_p_QuotedElement
    _1 string
}

type Tuple3_4bool_4bool_4bool struct {
    _0 bool
    _1 bool
    _2 bool
}

type Tuple3_4bool_4bool_14Option__string struct {
    _0 bool
    _1 bool
    _2 Option__string
}

type Tuple2_6string_26_goml_m_std_p_json_p_Value struct {
    _0 string
    _1 _goml_m_std_p_json_p_Value
}

type Tuple2_4bool_4char struct {
    _0 bool
    _1 rune
}

type Tuple2_13Option__isize_13Option__isize struct {
    _0 Option__isize
    _1 Option__isize
}

type Tuple2_6uint64_4bool struct {
    _0 uint64
    _1 bool
}

type Tuple2_4bool_3int struct {
    _0 bool
    _1 int
}

type Tuple2_6string_4bool struct {
    _0 string
    _1 bool
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

type _goml_m_std_p_num_p_ParseBoolError struct {
    input_value string
}

type _goml_m_std_p_num_p_ParseIntError struct {
    details _goml_m_std_p_io_p_ErrorDetails
    input_value string
    radix_value int
}

type _goml_m_std_p_num_p_ParseFloatError struct {
    details _goml_m_std_p_io_p_ErrorDetails
    input_value string
    range_value bool
}

type _goml_m_std_p_serde_p_VariantSchema struct {
    index int
    name string
    kind int
    fields *_goml_vec__goml_m_Tuple2__6string__18std_p_serde_p_Schema
}

type _goml_m_std_p_serde_p_ValueDeserializer struct {
    current *ref__goml_m_Option____std_p_serde_p_Value_x
    frames *_goml_vec__goml_m_std_p_serde_p_ValueDeserializeFrame
}

type _goml_m_std_p_serde_p_ValueSerializer struct {
    frames *_goml_vec__goml_m_std_p_serde_p_ValueSerializeFrame
    result *ref__goml_m_Option____std_p_serde_p_Value_x
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

type _goml_m_std_p_json_p_JsonDeserializer struct {
    parser _goml_m_std_p_json_p_JsonParser
    frames *_goml_vec__goml_m_std_p_json_p_JsonDeserializeFrame
}

type _goml_m_std_p_json_p_JsonParser struct {
    input string
    index *ref_int_x
}

type _goml_m_std_p_json_p_JsonSerializer struct {
    builder _goml_m_std_p_text_p_StringBuilder
    frames *_goml_vec__goml_m_std_p_json_p_JsonSerializeFrame
    root_written *ref_bool_x
}

type FnIterator__char struct {
    next_fn func() Option__char
}

type _goml_m_FnIterator____Slice_l_u8_r_ struct {
    next_fn func() _goml_m_Option____Slice_l_u8_r_
}

type FnIterator__string struct {
    next_fn func() Option__string
}

type _goml_m_FnIterator_____o_isize_c_char_q_ struct {
    next_fn func() _goml_m_Option_____o_isize_c_char_q_
}

type FnIterator__isize struct {
    next_fn func() Option__isize
}

type _goml_m_FnIterator____Result____string____std_p_text_p_ReplaceError struct {
    next_fn func() _goml_m_Option____Result____string____std_p_text_p_ReplaceError
}

type FnIterator__u8 struct {
    next_fn func() Option__u8
}

type _goml_m_FnIterator_____o_string_c_string_q_ struct {
    next_fn func() _goml_m_Option_____o_string_c_string_q_
}

type closure_env_std_bytes_split_iterator_0 struct {
    done_0 *ref_bool_x
    offset_1 *ref_int_x
    finder_2 _goml_m_std_p_bytes_p_Finder
    input_3 []uint8
    remaining_4 *ref_int_x
    after_5 bool
}

type closure_env_std_bytes_lines_iter_1 struct {
    parts_0 _goml_m_FnIterator____Slice_l_u8_r_
}

type closure_env_std_bytes_find_char_2 struct {
    character_0 rune
}

type closure_env_std_bytes_rfind_char_3 struct {
    character_0 rune
}

type closure_env_std_bytes_find_any_4 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_bytes_rfind_any_5 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_bytes_fields_by_iter_6 struct {
    offset_0 *ref_int_x
    input_1 []uint8
    separator_2 func(rune) bool
}

type closure_env_std_bytes_fields_iter_7 struct {}

type closure_env_std_bytes_fields_8 struct {}

type closure_env_std_bytes_trim_space_9 struct {}

type closure_env_std_bytes_trim_chars_10 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_bytes_trim_start_chars_11 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_bytes_trim_end_chars_12 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_bytes_case_with_13 struct {
    special_0 _goml_m_std_p_unicode_p_SpecialCase
    kind_1 _goml_m_std_p_unicode_p_Case
}

type closure_env_std_bytes_to_case_14 struct {
    kind_0 _goml_m_std_p_unicode_p_Case
}

type closure_env_std_io_read_stdin_to_string_15 struct {}

type closure_env_trait_impl_std_io_WriteAt_std_io_Cursor_write_at_16 struct {}

type closure_env_trait_impl_std_io_Write_std_io_PipeWriter_write_17 struct {}

type closure_env_std_num_parse_int_bits_18 struct {}

type closure_env_std_num_parse_uint_bits_19 struct {}

type closure_env_std_num_complex_component_20 struct {}

type closure_env_std_num_complex_component_21 struct {}

type closure_env_std_num_complex_components_22 struct {}

type closure_env_std_num_complex_components_23 struct {}

type closure_env_std_num_parse_complex_f32_24 struct {}

type closure_env_std_serde_parse_int_25 struct {}

type closure_env_std_serde_parse_uint_26 struct {}

type closure_env_std_serde_parse_float32_27 struct {}

type closure_env_std_serde_parse_float64_28 struct {}

type closure_env_trait_impl_std_ser_h041490d264fda8600ecbc53893ece034_ialize_int64_29 struct {}

type closure_env_trait_impl_std_ser_h7dcea04f8b1142fe863a42d6b087ae74_alize_uint64_30 struct {}

type closure_env_std_text_fields_by_iter_31 struct {
    offset_0 *ref_int_x
    value_1 string
    separator_2 func(rune) bool
}

type closure_env_std_text_fields_iter_32 struct {}

type closure_env_std_text_fields_33 struct {}

type closure_env_std_text_trim_space_34 struct {}

type closure_env_std_text_trim_chars_35 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_trim_start_chars_36 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_trim_end_chars_37 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_find_any_38 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_rfind_any_39 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_split_iterator_40 struct {
    done_0 *ref_bool_x
    offset_1 *ref_int_x
    separator_2 string
    value_3 string
    finder_4 _goml_m_std_p_text_p_Finder
    after_5 bool
}

type closure_env_std_text_lines_inclusive_iter_41 struct {
    parts_0 FnIterator__string
}

type closure_env_std_text_lines_iter_42 struct {
    parts_0 FnIterator__string
}

type closure_env_std_text_match_positions_43 struct {
    done_0 *ref_bool_x
    remaining_1 *ref_int_x
    finder_2 _goml_m_std_p_text_p_Finder
    offset_3 *ref_int_x
    value_4 string
}

type closure_env_std_text_case_checked_44 struct {
    kind_0 _goml_m_std_p_unicode_p_Case
}

type closure_env_std_text_case_with_checked_45 struct {
    special_0 _goml_m_std_p_unicode_p_SpecialCase
    kind_1 _goml_m_std_p_unicode_p_Case
}

type closure_env_std_text_unquote_46 struct {}

type closure_env_inherent_std_text_Replacer_std_text_Replacer_chunks_47 struct {
    done_0 *ref_bool_x
    cursor_1 _goml_m_std_p_text_p_ReplacementCursor
}

type closure_env_std_json_direct_json_int_48 struct {
    deserializer_0 _goml_m_std_p_json_p_JsonDeserializer
}

type closure_env_std_json_direct_json_uint_49 struct {
    deserializer_0 _goml_m_std_p_json_p_JsonDeserializer
}

type closure_env_trait_impl_std_ser_hd9a833df1dfd469aecf7b96c56fd99d5_lize_float32_50 struct {
    self_0 _goml_m_std_p_json_p_JsonDeserializer
}

type closure_env_trait_impl_std_ser_h0ccc965a377ead99aabd75125b739533_lize_float64_51 struct {
    self_0 _goml_m_std_p_json_p_JsonDeserializer
}

type closure_env_trait_impl_std_ser_hff9775fbd1a7df2f40012af8474c8705_ariant_field_52 struct {}

type closure_env_std_json_parse_json_value_53 struct {}

type closure_env_inherent_string_string_chars_54 struct {
    self_0 string
    index_1 *ref_int_x
}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Cursor_55 struct {}

type closure_env_std_io_trait_default_BufRead_discard_Self_std_io_Cursor_56 struct {
    discarded_0 *ref_int_x
}

type closure_env_std_io_trait_default_BufRead_discard_Self_std_io_Cursor_57 struct {
    discarded_0 *ref_int_x
}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_58 struct {}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_59 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Stdin_60 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_StringReader_61 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_PipeReader_62 struct {}

type closure_env_inherent_string_string_char_indices_63 struct {
    index_0 *ref_int_x
    self_1 string
}

type closure_env_goml_builtin_range_64 struct {
    current_0 *ref_int_x
    end_1 int
}

type closure_env_inherent_Slice_Slice_T_iter_T_u8_65 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 []uint8
}

type closure_env_inherent_Slice_Slice_T_iter_T_string_string_66 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 []Tuple2_6string_6string
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

type _goml_m_std_p_num_p_FormatIntError struct {
    _p0 int
    _p1 int
    _tag uint8
}

type _goml_m_std_p_num_p_IntWidthError interface {
    is_goml_m_std_p_num_p_IntWidthError()
}

type InvalidWidth struct {
    _0 int
}

func (_ InvalidWidth) is_goml_m_std_p_num_p_IntWidthError() {}

type Parse struct {
    _0 _goml_m_std_p_num_p_ParseIntError
}

func (_ Parse) is_goml_m_std_p_num_p_IntWidthError() {}

type _goml_m_std_p_num_p_IntWidthError_OutOfRange struct {
    _0 int
}

func (_ _goml_m_std_p_num_p_IntWidthError_OutOfRange) is_goml_m_std_p_num_p_IntWidthError() {}

type _goml_m_std_p_num_p_Rounding uint8

type _goml_m_std_p_num_p_FloatConversionError uint8

type _goml_m_std_p_num_p_FloatNotation uint8

type _goml_m_std_p_num_p_PowerFormat struct {
    _p0 int
    _p1 bool
    _tag uint8
}

type _goml_m_std_p_num_p_FormatFloatError struct {
    _p0 int
    _tag uint8
}

type _goml_m_std_p_num_p_DecimalStyle struct {
    _p0 bool
    _tag uint8
}

type _goml_m_std_p_num_p_ParseComplexError uint8

type _goml_m_std_p_num_p_ComplexFormat struct {
    _p1 int
    _p0 _goml_m_std_p_num_p_FloatNotation
    _p2 bool
    _tag uint8
}

type _goml_m_std_p_serde_p_FieldKey interface {
    is_goml_m_std_p_serde_p_FieldKey()
}

type Position struct {
    _0 int
}

func (_ Position) is_goml_m_std_p_serde_p_FieldKey() {}

type _goml_m_std_p_serde_p_FieldKey_Name struct {
    _0 string
}

func (_ _goml_m_std_p_serde_p_FieldKey_Name) is_goml_m_std_p_serde_p_FieldKey() {}

type _goml_m_std_p_serde_p_VariantKey interface {
    is_goml_m_std_p_serde_p_VariantKey()
}

type Index struct {
    _0 int
}

func (_ Index) is_goml_m_std_p_serde_p_VariantKey() {}

type _goml_m_std_p_serde_p_VariantKey_Name struct {
    _0 string
}

func (_ _goml_m_std_p_serde_p_VariantKey_Name) is_goml_m_std_p_serde_p_VariantKey() {}

type _goml_m_std_p_serde_p_Value interface {
    is_goml_m_std_p_serde_p_Value()
}

type _goml_m_std_p_serde_p_Value_Unit struct {}

func (_ _goml_m_std_p_serde_p_Value_Unit) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Bool struct {
    _0 bool
}

func (_ _goml_m_std_p_serde_p_Value_Bool) is_goml_m_std_p_serde_p_Value() {}

type Int struct {
    _0 int
}

func (_ Int) is_goml_m_std_p_serde_p_Value() {}

type Int8 struct {
    _0 int8
}

func (_ Int8) is_goml_m_std_p_serde_p_Value() {}

type Int16 struct {
    _0 int16
}

func (_ Int16) is_goml_m_std_p_serde_p_Value() {}

type Int32 struct {
    _0 int32
}

func (_ Int32) is_goml_m_std_p_serde_p_Value() {}

type Int64 struct {
    _0 int64
}

func (_ Int64) is_goml_m_std_p_serde_p_Value() {}

type Uint struct {
    _0 uint
}

func (_ Uint) is_goml_m_std_p_serde_p_Value() {}

type Uint8 struct {
    _0 uint8
}

func (_ Uint8) is_goml_m_std_p_serde_p_Value() {}

type Uint16 struct {
    _0 uint16
}

func (_ Uint16) is_goml_m_std_p_serde_p_Value() {}

type Uint32 struct {
    _0 uint32
}

func (_ Uint32) is_goml_m_std_p_serde_p_Value() {}

type Uint64 struct {
    _0 uint64
}

func (_ Uint64) is_goml_m_std_p_serde_p_Value() {}

type Float32 struct {
    _0 float32
}

func (_ Float32) is_goml_m_std_p_serde_p_Value() {}

type Float64 struct {
    _0 float64
}

func (_ Float64) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Number struct {
    _0 string
}

func (_ _goml_m_std_p_serde_p_Value_Number) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_String struct {
    _0 string
}

func (_ _goml_m_std_p_serde_p_Value_String) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Char struct {
    _0 rune
}

func (_ _goml_m_std_p_serde_p_Value_Char) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Binary struct {
    _0 *_goml_vec_uint8
}

func (_ _goml_m_std_p_serde_p_Value_Binary) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Map struct {
    _0 *_goml_vec__goml_m_Tuple2__17std_p_serde_p_Value__17std_p_serde_p_Value
}

func (_ _goml_m_std_p_serde_p_Value_Map) is_goml_m_std_p_serde_p_Value() {}

type Extension struct {
    _0 int8
    _1 *_goml_vec_uint8
}

func (_ Extension) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Sequence struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Value
}

func (_ _goml_m_std_p_serde_p_Value_Sequence) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Tuple struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Value
}

func (_ _goml_m_std_p_serde_p_Value_Tuple) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Optional struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Value
}

func (_ _goml_m_std_p_serde_p_Value_Optional) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Struct struct {
    _0 *_goml_vec__goml_m_Tuple2__6string__17std_p_serde_p_Value
}

func (_ _goml_m_std_p_serde_p_Value_Struct) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Value_Variant struct {
    _0 int
    _1 string
    _2 int
    _3 *_goml_vec__goml_m_Tuple2__6string__17std_p_serde_p_Value
}

func (_ _goml_m_std_p_serde_p_Value_Variant) is_goml_m_std_p_serde_p_Value() {}

type _goml_m_std_p_serde_p_Schema interface {
    is_goml_m_std_p_serde_p_Schema()
}

type Any struct {}

func (_ Any) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_Schema_Unit struct {}

func (_ _goml_m_std_p_serde_p_Schema_Unit) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_Schema_Bool struct {}

func (_ _goml_m_std_p_serde_p_Schema_Bool) is_goml_m_std_p_serde_p_Schema() {}

type Integer struct {
    _0 bool
    _1 int
}

func (_ Integer) is_goml_m_std_p_serde_p_Schema() {}

type Float struct {
    _0 int
}

func (_ Float) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_Schema_String struct {}

func (_ _goml_m_std_p_serde_p_Schema_String) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_Schema_Char struct {}

func (_ _goml_m_std_p_serde_p_Schema_Char) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_Schema_Sequence struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Schema
}

func (_ _goml_m_std_p_serde_p_Schema_Sequence) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_Schema_Tuple struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Schema
}

func (_ _goml_m_std_p_serde_p_Schema_Tuple) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_Schema_Optional struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Schema
}

func (_ _goml_m_std_p_serde_p_Schema_Optional) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_Schema_Struct struct {
    _0 *_goml_vec__goml_m_Tuple2__6string__18std_p_serde_p_Schema
}

func (_ _goml_m_std_p_serde_p_Schema_Struct) is_goml_m_std_p_serde_p_Schema() {}

type Enum struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_VariantSchema
}

func (_ Enum) is_goml_m_std_p_serde_p_Schema() {}

type _goml_m_std_p_serde_p_ValueDeserializeFrame interface {
    is_goml_m_std_p_serde_p_ValueDeserializeFrame()
}

type _goml_m_std_p_serde_p_ValueDeserializeFrame_Optional struct {}

func (_ _goml_m_std_p_serde_p_ValueDeserializeFrame_Optional) is_goml_m_std_p_serde_p_ValueDeserializeFrame() {}

type _goml_m_std_p_serde_p_ValueDeserializeFrame_Sequence struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Value
    _1 *ref_int_x
}

func (_ _goml_m_std_p_serde_p_ValueDeserializeFrame_Sequence) is_goml_m_std_p_serde_p_ValueDeserializeFrame() {}

type _goml_m_std_p_serde_p_ValueDeserializeFrame_Tuple struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Value
    _1 *ref_int_x
}

func (_ _goml_m_std_p_serde_p_ValueDeserializeFrame_Tuple) is_goml_m_std_p_serde_p_ValueDeserializeFrame() {}

type _goml_m_std_p_serde_p_ValueDeserializeFrame_Map struct {
    _0 *_goml_vec__goml_m_Tuple2__17std_p_serde_p_Value__17std_p_serde_p_Value
    _1 *ref_int_x
    _2 *ref_int_x
}

func (_ _goml_m_std_p_serde_p_ValueDeserializeFrame_Map) is_goml_m_std_p_serde_p_ValueDeserializeFrame() {}

type _goml_m_std_p_serde_p_ValueDeserializeFrame_Struct struct {
    _0 *_goml_vec__goml_m_Tuple2__6string__17std_p_serde_p_Value
    _1 *ref_int_x
}

func (_ _goml_m_std_p_serde_p_ValueDeserializeFrame_Struct) is_goml_m_std_p_serde_p_ValueDeserializeFrame() {}

type _goml_m_std_p_serde_p_ValueDeserializeFrame_Variant struct {
    _0 int
    _1 *_goml_vec__goml_m_Tuple2__6string__17std_p_serde_p_Value
    _2 *ref_int_x
    _3 *ref_bool_x
}

func (_ _goml_m_std_p_serde_p_ValueDeserializeFrame_Variant) is_goml_m_std_p_serde_p_ValueDeserializeFrame() {}

type _goml_m_std_p_serde_p_ValueSerializeFrame interface {
    is_goml_m_std_p_serde_p_ValueSerializeFrame()
}

type _goml_m_std_p_serde_p_ValueSerializeFrame_Optional struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Value
    _1 bool
    _2 *ref_bool_x
}

func (_ _goml_m_std_p_serde_p_ValueSerializeFrame_Optional) is_goml_m_std_p_serde_p_ValueSerializeFrame() {}

type _goml_m_std_p_serde_p_ValueSerializeFrame_Sequence struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Value
    _1 int
    _2 *ref_bool_x
}

func (_ _goml_m_std_p_serde_p_ValueSerializeFrame_Sequence) is_goml_m_std_p_serde_p_ValueSerializeFrame() {}

type _goml_m_std_p_serde_p_ValueSerializeFrame_Tuple struct {
    _0 *_goml_vec__goml_m_std_p_serde_p_Value
    _1 int
    _2 *ref_bool_x
}

func (_ _goml_m_std_p_serde_p_ValueSerializeFrame_Tuple) is_goml_m_std_p_serde_p_ValueSerializeFrame() {}

type _goml_m_std_p_serde_p_ValueSerializeFrame_Map struct {
    _0 *_goml_vec__goml_m_Tuple2__17std_p_serde_p_Value__17std_p_serde_p_Value
    _1 int
    _2 *ref_int_x
    _3 *ref__goml_m_Option____std_p_serde_p_Value_x
}

func (_ _goml_m_std_p_serde_p_ValueSerializeFrame_Map) is_goml_m_std_p_serde_p_ValueSerializeFrame() {}

type _goml_m_std_p_serde_p_ValueSerializeFrame_Struct struct {
    _0 *_goml_vec__goml_m_Tuple2__6string__17std_p_serde_p_Value
    _1 int
    _2 *ref_Option__string_x
}

func (_ _goml_m_std_p_serde_p_ValueSerializeFrame_Struct) is_goml_m_std_p_serde_p_ValueSerializeFrame() {}

type _goml_m_std_p_serde_p_ValueSerializeFrame_Variant struct {
    _0 int
    _1 string
    _2 int
    _3 *_goml_vec__goml_m_Tuple2__6string__17std_p_serde_p_Value
    _4 int
    _5 *ref_Option__string_x
}

func (_ _goml_m_std_p_serde_p_ValueSerializeFrame_Variant) is_goml_m_std_p_serde_p_ValueSerializeFrame() {}

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

type _goml_m_std_p_json_p_JsonDeserializeFrame interface {
    is_goml_m_std_p_json_p_JsonDeserializeFrame()
}

type _goml_m_std_p_json_p_JsonDeserializeFrame_Optional struct {}

func (_ _goml_m_std_p_json_p_JsonDeserializeFrame_Optional) is_goml_m_std_p_json_p_JsonDeserializeFrame() {}

type _goml_m_std_p_json_p_JsonDeserializeFrame_Sequence struct {
    _0 *ref_bool_x
    _1 *ref_bool_x
    _2 *ref_int_x
}

func (_ _goml_m_std_p_json_p_JsonDeserializeFrame_Sequence) is_goml_m_std_p_json_p_JsonDeserializeFrame() {}

type _goml_m_std_p_json_p_JsonDeserializeFrame_Tuple struct {
    _0 int
    _1 *ref_int_x
}

func (_ _goml_m_std_p_json_p_JsonDeserializeFrame_Tuple) is_goml_m_std_p_json_p_JsonDeserializeFrame() {}

type _goml_m_std_p_json_p_JsonDeserializeFrame_Map struct {
    _0 *ref_bool_x
    _1 *ref_bool_x
    _2 *ref_int_x
    _3 *ref_Option__string_x
}

func (_ _goml_m_std_p_json_p_JsonDeserializeFrame_Map) is_goml_m_std_p_json_p_JsonDeserializeFrame() {}

type _goml_m_std_p_json_p_JsonDeserializeFrame_Struct struct {
    _0 string
    _1 *ref_bool_x
    _2 *ref_bool_x
    _3 *ref_Option__string_x
}

func (_ _goml_m_std_p_json_p_JsonDeserializeFrame_Struct) is_goml_m_std_p_json_p_JsonDeserializeFrame() {}

type _goml_m_std_p_json_p_JsonDeserializeFrame_Variant struct {
    _0 string
    _1 string
    _2 int
    _3 *ref_int_x
    _4 *ref_int_x
    _5 *ref_int_x
    _6 *ref_bool_x
    _7 *ref_bool_x
    _8 *ref_bool_x
}

func (_ _goml_m_std_p_json_p_JsonDeserializeFrame_Variant) is_goml_m_std_p_json_p_JsonDeserializeFrame() {}

type _goml_m_std_p_json_p_Value interface {
    is_goml_m_std_p_json_p_Value()
}

type Object struct {
    _0 *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value
}

func (_ Object) is_goml_m_std_p_json_p_Value() {}

type _goml_m_std_p_json_p_Value_Array struct {
    _0 *_goml_vec__goml_m_std_p_json_p_Value
}

func (_ _goml_m_std_p_json_p_Value_Array) is_goml_m_std_p_json_p_Value() {}

type _goml_m_std_p_json_p_Value_String struct {
    _0 string
}

func (_ _goml_m_std_p_json_p_Value_String) is_goml_m_std_p_json_p_Value() {}

type _goml_m_std_p_json_p_Value_Number struct {
    _0 string
}

func (_ _goml_m_std_p_json_p_Value_Number) is_goml_m_std_p_json_p_Value() {}

type _goml_m_std_p_json_p_Value_Bool struct {
    _0 bool
}

func (_ _goml_m_std_p_json_p_Value_Bool) is_goml_m_std_p_json_p_Value() {}

type Null struct {}

func (_ Null) is_goml_m_std_p_json_p_Value() {}

type _goml_m_std_p_json_p_JsonSerializeFrame interface {
    is_goml_m_std_p_json_p_JsonSerializeFrame()
}

type _goml_m_std_p_json_p_JsonSerializeFrame_Optional struct {
    _0 *ref_bool_x
}

func (_ _goml_m_std_p_json_p_JsonSerializeFrame_Optional) is_goml_m_std_p_json_p_JsonSerializeFrame() {}

type _goml_m_std_p_json_p_JsonSerializeFrame_Array struct {
    _0 int
    _1 *ref_int_x
    _2 *ref_bool_x
}

func (_ _goml_m_std_p_json_p_JsonSerializeFrame_Array) is_goml_m_std_p_json_p_JsonSerializeFrame() {}

type _goml_m_std_p_json_p_JsonSerializeFrame_Map struct {
    _0 int
    _1 *ref_int_x
    _2 *ref_int_x
}

func (_ _goml_m_std_p_json_p_JsonSerializeFrame_Map) is_goml_m_std_p_json_p_JsonSerializeFrame() {}

type _goml_m_std_p_json_p_JsonSerializeFrame_Struct struct {
    _0 int
    _1 *ref_int_x
    _2 *ref_bool_x
}

func (_ _goml_m_std_p_json_p_JsonSerializeFrame_Struct) is_goml_m_std_p_json_p_JsonSerializeFrame() {}

type _goml_m_std_p_json_p_JsonSerializeFrame_Variant struct {
    _0 int
    _1 int
    _2 *ref_int_x
    _3 *ref_bool_x
}

func (_ _goml_m_std_p_json_p_JsonSerializeFrame_Variant) is_goml_m_std_p_json_p_JsonSerializeFrame() {}

type Option__char uint64

type _goml_m_Option____std_p_unicode_p_RangeTable struct {
    _p0 _goml_m_std_p_unicode_p_RangeTable
    _tag uint8
}

type _goml_m_Result____std_p_unicode_p_RangeTable____std_p_unicode_p_TableError struct {
    _p0 _goml_m_std_p_unicode_p_RangeTable
    _p1 _goml_m_std_p_unicode_p_TableError
    _tag uint8
}

type Option__isize struct {
    _p0 int
    _tag uint8
}

type _goml_m_Result____std_p_unicode_p_SpecialCase____std_p_unicode_p_TableError struct {
    _p0 _goml_m_std_p_unicode_p_SpecialCase
    _p1 _goml_m_std_p_unicode_p_TableError
    _tag uint8
}

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

type _goml_m_Option_____o_Slice_l_u8_r__c_Slice_l_u8_r__q_ interface {
    is_goml_m_Option_____o_Slice_l_u8_r__c_Slice_l_u8_r__q_()
}

type _goml_m_Option_____o_Slice_l_u8_r__c_Slice_l_u8_r__q__None struct {}

func (_ _goml_m_Option_____o_Slice_l_u8_r__c_Slice_l_u8_r__q__None) is_goml_m_Option_____o_Slice_l_u8_r__c_Slice_l_u8_r__q_() {}

type _goml_m_Option_____o_Slice_l_u8_r__c_Slice_l_u8_r__q__Some struct {
    _0 Tuple2_12Slice_5uint8_12Slice_5uint8
}

func (_ _goml_m_Option_____o_Slice_l_u8_r__c_Slice_l_u8_r__q__Some) is_goml_m_Option_____o_Slice_l_u8_r__c_Slice_l_u8_r__q_() {}

type _goml_m_Result____isize____std_p_bytes_p_TransformError struct {
    _p1 _goml_m_std_p_bytes_p_TransformError
    _p0 int
    _tag uint8
}

type _goml_m_Result_____o__q_____std_p_bytes_p_TransformError struct {
    _p0 _goml_m_std_p_bytes_p_TransformError
    _tag uint8
}

type _goml_m_Result____Vec_l_Slice_l_u8_r__r_____std_p_bytes_p_TransformError struct {
    _p1 _goml_m_std_p_bytes_p_TransformError
    _p0 *_goml_vec_Slice_5uint8
    _tag uint8
}

type _goml_m_Result____std_p_bytes_p_Bytes____std_p_bytes_p_TransformError struct {
    _p1 _goml_m_std_p_bytes_p_TransformError
    _p0 _goml_m_std_p_bytes_p_Bytes
    _tag uint8
}

type _goml_m_Result____string____std_p_utf8_p_Utf8Error struct {
    _p1 _goml_m_std_p_utf8_p_Utf8Error
    _p0 string
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

type _goml_m_Result_____o_char_c_isize_q_____std_p_utf8_p_Utf8Error struct {
    _p1 _goml_m_std_p_utf8_p_Utf8Error
    _p0 Tuple2_4char_3int
    _tag uint8
}

type _goml_m_Option____std_p_utf8_p_Utf8Error struct {
    _p0 _goml_m_std_p_utf8_p_Utf8Error
    _tag uint8
}

type _goml_m_Result____Option____char____std_p_utf8_p_DecoderError struct {
    _p1 _goml_m_std_p_utf8_p_DecoderError
    _p0 Option__char
    _tag uint8
}

type _goml_m_Result_____o__q_____std_p_utf8_p_DecoderError struct {
    _p0 _goml_m_std_p_utf8_p_DecoderError
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

type _goml_m_Result_____o__q_____std_p_io_p_TransferError interface {
    is_goml_m_Result_____o__q_____std_p_io_p_TransferError()
}

type _goml_m_Result_____o__q_____std_p_io_p_TransferError_Ok struct {
    _0 struct{}
}

func (_ _goml_m_Result_____o__q_____std_p_io_p_TransferError_Ok) is_goml_m_Result_____o__q_____std_p_io_p_TransferError() {}

type _goml_m_Result_____o__q_____std_p_io_p_TransferError_Err struct {
    _0 _goml_m_std_p_io_p_TransferError
}

func (_ _goml_m_Result_____o__q_____std_p_io_p_TransferError_Err) is_goml_m_Result_____o__q_____std_p_io_p_TransferError() {}

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

type _goml_m_Result____std_p_io_p_Cursor____std_p_io_p_Error interface {
    is_goml_m_Result____std_p_io_p_Cursor____std_p_io_p_Error()
}

type _goml_m_Result____std_p_io_p_Cursor____std_p_io_p_Error_Ok struct {
    _0 _goml_m_std_p_io_p_Cursor
}

func (_ _goml_m_Result____std_p_io_p_Cursor____std_p_io_p_Error_Ok) is_goml_m_Result____std_p_io_p_Cursor____std_p_io_p_Error() {}

type _goml_m_Result____std_p_io_p_Cursor____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result____std_p_io_p_Cursor____std_p_io_p_Error_Err) is_goml_m_Result____std_p_io_p_Cursor____std_p_io_p_Error() {}

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

type _goml_m_Result____std_p_io_p_ScanStep____std_p_io_p_Error interface {
    is_goml_m_Result____std_p_io_p_ScanStep____std_p_io_p_Error()
}

type _goml_m_Result____std_p_io_p_ScanStep____std_p_io_p_Error_Ok struct {
    _0 _goml_m_std_p_io_p_ScanStep
}

func (_ _goml_m_Result____std_p_io_p_ScanStep____std_p_io_p_Error_Ok) is_goml_m_Result____std_p_io_p_ScanStep____std_p_io_p_Error() {}

type _goml_m_Result____std_p_io_p_ScanStep____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result____std_p_io_p_ScanStep____std_p_io_p_Error_Err) is_goml_m_Result____std_p_io_p_ScanStep____std_p_io_p_Error() {}

type _goml_m_Option_____o_isize_c_isize_q_ struct {
    _p0 Tuple2_3int_3int
    _tag uint8
}

type _goml_m_Option_____o_char_c_isize_q_ struct {
    _p0 Tuple2_4char_3int
    _tag uint8
}

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

type _goml_m_Option____std_p_io_p_PipeState interface {
    is_goml_m_Option____std_p_io_p_PipeState()
}

type _goml_m_Option____std_p_io_p_PipeState_None struct {}

func (_ _goml_m_Option____std_p_io_p_PipeState_None) is_goml_m_Option____std_p_io_p_PipeState() {}

type _goml_m_Option____std_p_io_p_PipeState_Some struct {
    _0 _goml_m_std_p_io_p_PipeState
}

func (_ _goml_m_Option____std_p_io_p_PipeState_Some) is_goml_m_Option____std_p_io_p_PipeState() {}

type _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_Error interface {
    is_goml_m_Result____std_p_io_p_PipeState____std_p_io_p_Error()
}

type _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_Error_Ok struct {
    _0 _goml_m_std_p_io_p_PipeState
}

func (_ _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_Error_Ok) is_goml_m_Result____std_p_io_p_PipeState____std_p_io_p_Error() {}

type _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_Error_Err struct {
    _0 _goml_m_std_p_io_p_Error
}

func (_ _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_Error_Err) is_goml_m_Result____std_p_io_p_PipeState____std_p_io_p_Error() {}

type _goml_m_Option_____o__q_ uint8

type _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_TransferError interface {
    is_goml_m_Result____std_p_io_p_PipeState____std_p_io_p_TransferError()
}

type _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_TransferError_Ok struct {
    _0 _goml_m_std_p_io_p_PipeState
}

func (_ _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_TransferError_Ok) is_goml_m_Result____std_p_io_p_PipeState____std_p_io_p_TransferError() {}

type _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_TransferError_Err struct {
    _0 _goml_m_std_p_io_p_TransferError
}

func (_ _goml_m_Result____std_p_io_p_PipeState____std_p_io_p_TransferError_Err) is_goml_m_Result____std_p_io_p_PipeState____std_p_io_p_TransferError() {}

type _goml_m_Option____Result____isize____std_p_io_p_TransferError struct {
    _p0 _goml_m_Result____isize____std_p_io_p_TransferError
    _tag uint8
}

type _goml_m_Result____Result____is_h95e4d74e2a903fcd2136f2199d9d1a4e_p_TransferError interface {
    is_goml_m_Result____Result_____hf5643c21937cec799584c47faa31a151_p_TransferError()
}

type _goml_m_Result____Result____is_h38a640c1674b26e78ca6b1182c6b023d_ransferError_Ok struct {
    _0 _goml_m_Result____isize____std_p_io_p_TransferError
}

func (_ _goml_m_Result____Result____is_h38a640c1674b26e78ca6b1182c6b023d_ransferError_Ok) is_goml_m_Result____Result_____hf5643c21937cec799584c47faa31a151_p_TransferError() {}

type _goml_m_Result____Result____is_h7f29888bc4595ed12102cd212f288b22_ansferError_Err struct {
    _0 _goml_m_std_p_io_p_TransferError
}

func (_ _goml_m_Result____Result____is_h7f29888bc4595ed12102cd212f288b22_ansferError_Err) is_goml_m_Result____Result_____hf5643c21937cec799584c47faa31a151_p_TransferError() {}

type _goml_m_Result____bool____std_p_num_p_ParseBoolError struct {
    _p1 _goml_m_std_p_num_p_ParseBoolError
    _p0 bool
    _tag uint8
}

type _goml_m_Result____Vec_l_u8_r_____std_p_num_p_FormatIntError struct {
    _p1 _goml_m_std_p_num_p_FormatIntError
    _p0 *_goml_vec_uint8
    _tag uint8
}

type _goml_m_Result____string____std_p_num_p_FormatIntError struct {
    _p1 _goml_m_std_p_num_p_FormatIntError
    _p0 string
    _tag uint8
}

type _goml_m_Result____isize____std_p_num_p_FormatIntError struct {
    _p1 _goml_m_std_p_num_p_FormatIntError
    _p0 int
    _tag uint8
}

type _goml_m_Result____isize____std_p_num_p_ParseIntError interface {
    is_goml_m_Result____isize____std_p_num_p_ParseIntError()
}

type _goml_m_Result____isize____std_p_num_p_ParseIntError_Ok struct {
    _0 int
}

func (_ _goml_m_Result____isize____std_p_num_p_ParseIntError_Ok) is_goml_m_Result____isize____std_p_num_p_ParseIntError() {}

type _goml_m_Result____isize____std_p_num_p_ParseIntError_Err struct {
    _0 _goml_m_std_p_num_p_ParseIntError
}

func (_ _goml_m_Result____isize____std_p_num_p_ParseIntError_Err) is_goml_m_Result____isize____std_p_num_p_ParseIntError() {}

type _goml_m_Result____isize____std_p_num_p_IntWidthError struct {
    _p1 _goml_m_std_p_num_p_IntWidthError
    _p0 int
    _tag uint8
}

type _goml_m_Result____i64____std_p_num_p_IntWidthError struct {
    _p1 _goml_m_std_p_num_p_IntWidthError
    _p0 int64
    _tag uint8
}

type _goml_m_Result____u64____std_p_num_p_IntWidthError struct {
    _p1 _goml_m_std_p_num_p_IntWidthError
    _p0 uint64
    _tag uint8
}

type _goml_m_Result____usize____std_p_num_p_ParseIntError interface {
    is_goml_m_Result____usize____std_p_num_p_ParseIntError()
}

type _goml_m_Result____usize____std_p_num_p_ParseIntError_Ok struct {
    _0 uint
}

func (_ _goml_m_Result____usize____std_p_num_p_ParseIntError_Ok) is_goml_m_Result____usize____std_p_num_p_ParseIntError() {}

type _goml_m_Result____usize____std_p_num_p_ParseIntError_Err struct {
    _0 _goml_m_std_p_num_p_ParseIntError
}

func (_ _goml_m_Result____usize____std_p_num_p_ParseIntError_Err) is_goml_m_Result____usize____std_p_num_p_ParseIntError() {}

type _goml_m_Result____usize____std_p_num_p_IntWidthError struct {
    _p1 _goml_m_std_p_num_p_IntWidthError
    _p0 uint
    _tag uint8
}

type _goml_m_Result____f32____std_p_num_p_ParseFloatError interface {
    is_goml_m_Result____f32____std_p_num_p_ParseFloatError()
}

type _goml_m_Result____f32____std_p_num_p_ParseFloatError_Ok struct {
    _0 float32
}

func (_ _goml_m_Result____f32____std_p_num_p_ParseFloatError_Ok) is_goml_m_Result____f32____std_p_num_p_ParseFloatError() {}

type _goml_m_Result____f32____std_p_num_p_ParseFloatError_Err struct {
    _0 _goml_m_std_p_num_p_ParseFloatError
}

func (_ _goml_m_Result____f32____std_p_num_p_ParseFloatError_Err) is_goml_m_Result____f32____std_p_num_p_ParseFloatError() {}

type _goml_m_Result____f64____std_p_num_p_ParseFloatError interface {
    is_goml_m_Result____f64____std_p_num_p_ParseFloatError()
}

type _goml_m_Result____f64____std_p_num_p_ParseFloatError_Ok struct {
    _0 float64
}

func (_ _goml_m_Result____f64____std_p_num_p_ParseFloatError_Ok) is_goml_m_Result____f64____std_p_num_p_ParseFloatError() {}

type _goml_m_Result____f64____std_p_num_p_ParseFloatError_Err struct {
    _0 _goml_m_std_p_num_p_ParseFloatError
}

func (_ _goml_m_Result____f64____std_p_num_p_ParseFloatError_Err) is_goml_m_Result____f64____std_p_num_p_ParseFloatError() {}

type Option__i64 struct {
    _p0 int64
    _tag uint8
}

type _goml_m_Result_____o_bool_c_u64_q_____std_p_num_p_FloatConversionError struct {
    _p0 Tuple2_4bool_6uint64
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____i64____std_p_num_p_FloatConversionError struct {
    _p0 int64
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____u64____std_p_num_p_FloatConversionError struct {
    _p0 uint64
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____isize____std_p_num_p_FloatConversionError struct {
    _p0 int
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____i8____std_p_num_p_FloatConversionError struct {
    _p0 int8
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____i16____std_p_num_p_FloatConversionError struct {
    _p0 int16
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____i32____std_p_num_p_FloatConversionError struct {
    _p0 int32
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____usize____std_p_num_p_FloatConversionError struct {
    _p0 uint
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____u8____std_p_num_p_FloatConversionError struct {
    _p0 uint8
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____u16____std_p_num_p_FloatConversionError struct {
    _p0 uint16
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result____u32____std_p_num_p_FloatConversionError struct {
    _p0 uint32
    _p1 _goml_m_std_p_num_p_FloatConversionError
    _tag uint8
}

type _goml_m_Result_____o_Vec_l_u8_r__c_isize_q_____std_p_num_p_FormatFloatError struct {
    _p0 Tuple2_10Vec_5uint8_3int
    _p1 _goml_m_std_p_num_p_FormatFloatError
    _tag uint8
}

type _goml_m_Result____string____std_p_num_p_FormatFloatError struct {
    _p0 string
    _p1 _goml_m_std_p_num_p_FormatFloatError
    _tag uint8
}

type _goml_m_Result____f64____std_p_num_p_ParseComplexError struct {
    _p0 float64
    _p1 _goml_m_std_p_num_p_ParseComplexError
    _tag uint8
}

type _goml_m_Result_____o_f64_c_f64_q_____std_p_num_p_ParseComplexError struct {
    _p0 Tuple2_7float64_7float64
    _p1 _goml_m_std_p_num_p_ParseComplexError
    _tag uint8
}

type _goml_m_Result_____o_f32_c_f32_q_____std_p_num_p_ParseComplexError struct {
    _p0 Tuple2_7float32_7float32
    _p1 _goml_m_std_p_num_p_ParseComplexError
    _tag uint8
}

type _goml_m_Option____std_p_serde_p_FieldKey struct {
    _p0 _goml_m_std_p_serde_p_FieldKey
    _tag uint8
}

type _goml_m_Result____std_p_serde_p_Value____string struct {
    _p0 _goml_m_std_p_serde_p_Value
    _p1 string
    _tag uint8
}

type _goml_m_Result_____o_string_c__h0fd03a31a9a9f0946c70cfe7367b8f54_r__q_____string interface {
    is_goml_m_Result_____o_string__he5eedb2e62343642e9a52d0140c62a64_r__q_____string()
}

type _goml_m_Result_____o_string_c__h8dcc36dec0d2d2fc6ff4f8cac4666ec9_q_____string_Ok struct {
    _0 Tuple3_6string_3int_50Vec_44Tuple2_6string_27_goml_m_std_p_serde_p_Value
}

func (_ _goml_m_Result_____o_string_c__h8dcc36dec0d2d2fc6ff4f8cac4666ec9_q_____string_Ok) is_goml_m_Result_____o_string__he5eedb2e62343642e9a52d0140c62a64_r__q_____string() {}

type _goml_m_Result_____o_string_c__ha7cf7403e05526da4cceea823947098e______string_Err struct {
    _0 string
}

func (_ _goml_m_Result_____o_string_c__ha7cf7403e05526da4cceea823947098e______string_Err) is_goml_m_Result_____o_string__he5eedb2e62343642e9a52d0140c62a64_r__q_____string() {}

type _goml_m_Result_____o__q_____string struct {
    _p0 string
    _tag uint8
}

type Result__isize__string struct {
    _p1 string
    _p0 int
    _tag uint8
}

type Result__usize__string struct {
    _p1 string
    _p0 uint
    _tag uint8
}

type Result__f32__string struct {
    _p1 string
    _p0 float32
    _tag uint8
}

type Result__f64__string struct {
    _p1 string
    _p0 float64
    _tag uint8
}

type _goml_m_Option____std_p_serde_p_Value struct {
    _p0 _goml_m_std_p_serde_p_Value
    _tag uint8
}

type _goml_m_Option____std_p_serde_p_ValueDeserializeFrame struct {
    _p0 _goml_m_std_p_serde_p_ValueDeserializeFrame
    _tag uint8
}

type Result__string__string struct {
    _p0 string
    _tag uint8
}

type _goml_m_Result____Vec_l_u8_r_____string struct {
    _p1 string
    _p0 *_goml_vec_uint8
    _tag uint8
}

type Result__bool__string struct {
    _p1 string
    _p0 bool
    _tag uint8
}

type Result__i8__string struct {
    _p1 string
    _p0 int8
    _tag uint8
}

type Result__i16__string struct {
    _p1 string
    _p0 int16
    _tag uint8
}

type Result__i32__string struct {
    _p1 string
    _p0 int32
    _tag uint8
}

type Result__i64__string struct {
    _p1 string
    _p0 int64
    _tag uint8
}

type Result__u8__string struct {
    _p1 string
    _p0 uint8
    _tag uint8
}

type Result__u16__string struct {
    _p1 string
    _p0 uint16
    _tag uint8
}

type Result__u32__string struct {
    _p1 string
    _p0 uint32
    _tag uint8
}

type Result__u64__string struct {
    _p1 string
    _p0 uint64
    _tag uint8
}

type Result__char__string struct {
    _p1 string
    _p0 rune
    _tag uint8
}

type _goml_m_Result_____o_i8_c_Vec_l_u8_r__q_____string struct {
    _p0 Tuple2_4int8_10Vec_5uint8
    _p1 string
    _tag uint8
}

type Result__Option__isize__string struct {
    _p0 Option__isize
    _p1 string
    _tag uint8
}

type _goml_m_Result____Option____std_p_serde_p_FieldKey____string struct {
    _p0 _goml_m_Option____std_p_serde_p_FieldKey
    _p1 string
    _tag uint8
}

type _goml_m_Result____std_p_serde_p_VariantKey____string struct {
    _p0 _goml_m_std_p_serde_p_VariantKey
    _p1 string
    _tag uint8
}

type _goml_m_Option____std_p_serde_p_ValueSerializeFrame struct {
    _p0 _goml_m_std_p_serde_p_ValueSerializeFrame
    _tag uint8
}

type _goml_m_Result____Vec_l_string_r_____std_p_bytes_p_TransformError struct {
    _p1 _goml_m_std_p_bytes_p_TransformError
    _p0 *_goml_vec_string
    _tag uint8
}

type _goml_m_Option_____o_isize_c_char_q_ struct {
    _p0 Tuple2_3int_4char
    _tag uint8
}

type _goml_m_Option_____o_string_c_string_q_ struct {
    _p0 Tuple2_6string_6string
    _tag uint8
}

type _goml_m_Result____string____std_p_bytes_p_TransformError struct {
    _p0 string
    _p1 _goml_m_std_p_bytes_p_TransformError
    _tag uint8
}

type Option__u32 uint64

type _goml_m_Result_____o_std_p_tex_h8e435facaacabba5e764f2f239ae3f1c__p_UnquoteError struct {
    _p0 Tuple2_34_goml_m_std_p_text_p_QuotedElement_3int
    _p1 _goml_m_std_p_text_p_UnquoteError
    _tag uint8
}

type _goml_m_Result_____o_std_p_tex_h6608824cb0a3e66f3073db78fd881a91__p_UnquoteError struct {
    _p0 Tuple2_34_goml_m_std_p_text_p_QuotedElement_6string
    _p1 _goml_m_std_p_text_p_UnquoteError
    _tag uint8
}

type _goml_m_Result____std_p_bytes_p_Bytes____std_p_text_p_UnquoteError struct {
    _p1 _goml_m_std_p_text_p_UnquoteError
    _p0 _goml_m_std_p_bytes_p_Bytes
    _tag uint8
}

type _goml_m_Result____string____std_p_text_p_UnquoteError struct {
    _p0 string
    _p1 _goml_m_std_p_text_p_UnquoteError
    _tag uint8
}

type _goml_m_Result_____o__q_____std_p_text_p_ReplaceError struct {
    _p0 _goml_m_std_p_text_p_ReplaceError
    _tag uint8
}

type _goml_m_Result____Option____string____std_p_text_p_ReplaceError struct {
    _p0 Option__string
    _p1 _goml_m_std_p_text_p_ReplaceError
    _tag uint8
}

type _goml_m_Result_____o_isize_c_isize_q_____std_p_text_p_ReplaceError struct {
    _p0 Tuple2_3int_3int
    _p1 _goml_m_std_p_text_p_ReplaceError
    _tag uint8
}

type _goml_m_Result____std_p_text_p_Replacer____std_p_text_p_ReplaceError struct {
    _p0 _goml_m_std_p_text_p_Replacer
    _p1 _goml_m_std_p_text_p_ReplaceError
    _tag uint8
}

type _goml_m_Result____string____std_p_text_p_ReplaceError struct {
    _p0 string
    _p1 _goml_m_std_p_text_p_ReplaceError
    _tag uint8
}

type _goml_m_Option____Result____string____std_p_text_p_ReplaceError struct {
    _p0 _goml_m_Result____string____std_p_text_p_ReplaceError
    _tag uint8
}

type _goml_m_Result____FnIterator___h8cbf8748430d700ba9742992933c2390__p_ReplaceError struct {
    _p1 _goml_m_std_p_text_p_ReplaceError
    _p0 _goml_m_FnIterator____Result____string____std_p_text_p_ReplaceError
    _tag uint8
}

type _goml_m_Option____std_p_json_p_JsonDeserializeFrame struct {
    _p0 _goml_m_std_p_json_p_JsonDeserializeFrame
    _tag uint8
}

type _goml_m_Result_____o_bool_c_bool_c_bool_q_____string struct {
    _p1 string
    _p0 Tuple3_4bool_4bool_4bool
    _tag uint8
}

type _goml_m_Result_____o_bool_c_bool_c_Option____string_q_____string interface {
    is_goml_m_Result_____o_bool_c_bool_c_Option____string_q_____string()
}

type _goml_m_Result_____o_bool_c_bool_c_Option____string_q_____string_Ok struct {
    _0 Tuple3_4bool_4bool_14Option__string
}

func (_ _goml_m_Result_____o_bool_c_bool_c_Option____string_q_____string_Ok) is_goml_m_Result_____o_bool_c_bool_c_Option____string_q_____string() {}

type _goml_m_Result_____o_bool_c_bool_c_Option____string_q_____string_Err struct {
    _0 string
}

func (_ _goml_m_Result_____o_bool_c_bool_c_Option____string_q_____string_Err) is_goml_m_Result_____o_bool_c_bool_c_Option____string_q_____string() {}

type Result__Option__string__string struct {
    _p0 Option__string
    _p1 string
    _tag uint8
}

type _goml_m_Result____std_p_json_p_Value____string struct {
    _p0 _goml_m_std_p_json_p_Value
    _p1 string
    _tag uint8
}

type _goml_m_Option____std_p_json_p_Value struct {
    _p0 _goml_m_std_p_json_p_Value
    _tag uint8
}

type Option__bool uint8

type _goml_m_Option____Vec_l_std_p_json_p_Value_r_ struct {
    _value *_goml_vec__goml_m_std_p_json_p_Value
}

type _goml_m_Option____std_p_json_p_JsonSerializeFrame struct {
    _p0 _goml_m_std_p_json_p_JsonSerializeFrame
    _tag uint8
}

type _goml_m_Result____Slice_l_u8_r_____std_p_io_p_TransferError interface {
    is_goml_m_Result____Slice_l_u8_r_____std_p_io_p_TransferError()
}

type _goml_m_Result____Slice_l_u8_r_____std_p_io_p_TransferError_Ok struct {
    _0 []uint8
}

func (_ _goml_m_Result____Slice_l_u8_r_____std_p_io_p_TransferError_Ok) is_goml_m_Result____Slice_l_u8_r_____std_p_io_p_TransferError() {}

type _goml_m_Result____Slice_l_u8_r_____std_p_io_p_TransferError_Err struct {
    _0 _goml_m_std_p_io_p_TransferError
}

func (_ _goml_m_Result____Slice_l_u8_r_____std_p_io_p_TransferError_Err) is_goml_m_Result____Slice_l_u8_r_____std_p_io_p_TransferError() {}

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

func _goml_m_inherent_i_std_p_text_p_StringBuilder_i_std_p_text_p_StringBuilder_i_new() _goml_m_std_p_text_p_StringBuilder {
    var t0 *_goml_vec_uint8
    var inline0 *_goml_vec_uint8 = vec_new__Vec_5uint8()
    t0 = inline0
    var t1 _goml_m_std_p_text_p_StringBuilder = _goml_m_std_p_text_p_StringBuilder{
        values: t0,
    }
    return t1
}

func _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(self__0 _goml_m_std_p_text_p_StringBuilder, value__0 string) struct{} {
    var length__0 int
    var inline3 int = _goml_runtime_core_string_len(value__0)
    length__0 = inline3
    var t0 *_goml_vec_uint8 = self__0.values
    vec_reserve__Vec_5uint8(t0, length__0)
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t1 bool = for_index0 < length__0
        if t1 {
            var for_item0 int = for_index0
            var t2 int = for_index0 + 1
            for_index0 = t2
            var t3 *_goml_vec_uint8 = self__0.values
            var t4 uint8
            var inline1 uint8 = _goml_runtime_core_string_byte_get(value__0, for_item0)
            t4 = inline1
            vec_push__Vec_5uint8(t3, t4)
            continue
        } else {
            break Loop_loop0
        }
    }
    return struct{}{}
}

func _goml_m_inherent_i_std_p_text__he83b4afafc069a3c24c64018b13ca033_r_i_write__char(self__0 _goml_m_std_p_text_p_StringBuilder, value__0 rune) struct{} {
    var t0 string
    var inline0 string = char_to_string(value__0)
    t0 = inline0
    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(self__0, t0)
    return struct{}{}
}

func _goml_m_std_p_json_p_json__error(value__0 _goml_m_std_p_json_p_JsonParser, message__0 string) string {
    var t0 string = "" + message__0
    var t1 string = t0 + " at byte "
    var t2 *ref_int_x = value__0.index
    var t3 int
    var inline1 int = ref_get__Ref_3int(t2)
    t3 = inline1
    var t4 string
    var inline0 string = __goml_builtin_int_to_string(t3)
    t4 = inline0
    var t5 string = t1 + t4
    return t5
}

func _goml_m_std_p_json_p_skip__json__whitespace(value__0 _goml_m_std_p_json_p_JsonParser) struct{} {
    Loop_loop0:
    for {
        var t0 *ref_int_x = value__0.index
        var t1 int
        var inline11 int = ref_get__Ref_3int(t0)
        t1 = inline11
        var t2 string = value__0.input
        var t3 int
        var inline10 int = _goml_runtime_core_string_len(t2)
        t3 = inline10
        var t4 bool = t1 < t3
        var jp0 bool
        if t4 {
            var t9 string = value__0.input
            var t10 *ref_int_x = value__0.index
            var t11 int
            var inline9 int = ref_get__Ref_3int(t10)
            t11 = inline9
            var t12 uint8
            var inline8 uint8 = _goml_runtime_core_string_byte_get(t9, t11)
            t12 = inline8
            var inline2 bool = t12 == 9
            var inline3 bool
            if inline2 {
                inline3 = true
            } else {
                var inline7 bool = t12 == 10
                inline3 = inline7
            }
            var inline4 bool
            if inline3 {
                inline4 = true
            } else {
                var inline6 bool = t12 == 13
                inline4 = inline6
            }
            if inline4 {
                jp0 = true
            } else {
                var inline5 bool = t12 == 32
                jp0 = inline5
            }
        } else {
            jp0 = false
        }
        if jp0 {
            var t5 *ref_int_x = value__0.index
            var t6 *ref_int_x = value__0.index
            var t7 int
            var inline1 int = ref_get__Ref_3int(t6)
            t7 = inline1
            var t8 int = t7 + 1
            ref_set__Ref_3int(t5, t8)
            continue
        } else {
            break Loop_loop0
        }
    }
    return struct{}{}
}

func _goml_m_std_p_json_p_hex__digit(value__0 uint8) Option__u32 {
    var t0 bool = value__0 >= 48
    var jp0 bool
    if t0 {
        var t16 bool = value__0 <= 57
        jp0 = t16
    } else {
        jp0 = false
    }
    if jp0 {
        var t1 uint8 = value__0 - 48
        var t2 uint32 = uint32(uint8(t1))
        var t3 Option__u32 = Option__u32(uint64(uint32(t2)) + 1)
        return t3
    } else {
        var t4 bool = value__0 >= 65
        var jp1 bool
        if t4 {
            var t15 bool = value__0 <= 70
            jp1 = t15
        } else {
            jp1 = false
        }
        if jp1 {
            var t5 uint8 = value__0 - 65
            var t6 uint8 = t5 + 10
            var t7 uint32 = uint32(uint8(t6))
            var t8 Option__u32 = Option__u32(uint64(uint32(t7)) + 1)
            return t8
        } else {
            var t9 bool = value__0 >= 97
            var jp2 bool
            if t9 {
                var t14 bool = value__0 <= 102
                jp2 = t14
            } else {
                jp2 = false
            }
            if jp2 {
                var t10 uint8 = value__0 - 97
                var t11 uint8 = t10 + 10
                var t12 uint32 = uint32(uint8(t11))
                var t13 Option__u32 = Option__u32(uint64(uint32(t12)) + 1)
                return t13
            } else {
                return Option__u32(0)
            }
        }
    }
}

func _goml_m_std_p_json_p_parse__hex__quad(value__0 _goml_m_std_p_json_p_JsonParser) Result__u32__string {
    var t0 *ref_int_x = value__0.index
    var t1 int
    var inline19 int = ref_get__Ref_3int(t0)
    t1 = inline19
    var t2 int = t1 + 4
    var t3 string = value__0.input
    var t4 int
    var inline18 int = _goml_runtime_core_string_len(t3)
    t4 = inline18
    var t5 bool = t2 > t4
    if t5 {
        var t6 string
        var inline0 string = "incomplete unicode escape"
        var inline1 string = "" + inline0
        var inline2 string = inline1 + " at byte "
        var inline3 *ref_int_x = value__0.index
        var inline4 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline3)
        var inline5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline4)
        var inline6 string = inline2 + inline5
        t6 = inline6
        var t7 Result__u32__string = Result__u32__string{
            _p1: t6,
            _tag: 1,
        }
        return t7
    } else {
        var result__0_source int = 0
        var result__0 uint32 = uint32(int(result__0_source))
        var for_index0 int = 0
        var for_limit0 int = 4
        Loop_loop0:
        for {
            var t13 bool = for_index0 < for_limit0
            if t13 {
                var for_item0 int = for_index0
                var t14 int = for_index0 + 1
                for_index0 = t14
                var t15 string = value__0.input
                var t16 *ref_int_x = value__0.index
                var t17 int
                var inline17 int = ref_get__Ref_3int(t16)
                t17 = inline17
                var t18 int = t17 + for_item0
                var t19 uint8
                var inline16 uint8 = _goml_runtime_core_string_byte_get(t15, t18)
                t19 = inline16
                var mtmp0 Option__u32 = _goml_m_std_p_json_p_hex__digit(t19)
                switch mtmp0 != Option__u32(0) {
                case false:
                    var t20 string
                    var inline9 string = "invalid unicode escape"
                    var inline10 string = "" + inline9
                    var inline11 string = inline10 + " at byte "
                    var inline12 *ref_int_x = value__0.index
                    var inline13 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline12)
                    var inline14 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline13)
                    var inline15 string = inline11 + inline14
                    t20 = inline15
                    var t21 Result__u32__string = Result__u32__string{
                        _p1: t20,
                        _tag: 1,
                    }
                    return t21
                case true:
                    var x0 uint32 = uint32(uint64(mtmp0) - 1)
                    var t22 uint32 = result__0 * 16
                    var t23 uint32 = t22 + x0
                    result__0 = t23
                    continue
                default:
                    panic("non-exhaustive match")
                }
            } else {
                break Loop_loop0
            }
        }
        var t8 *ref_int_x = value__0.index
        var t9 *ref_int_x = value__0.index
        var t10 int
        var inline8 int = ref_get__Ref_3int(t9)
        t10 = inline8
        var t11 int = t10 + 4
        ref_set__Ref_3int(t8, t11)
        var t12 Result__u32__string = Result__u32__string{
            _p0: result__0,
            _tag: 0,
        }
        return t12
    }
}

func _goml_m_std_p_json_p_write__codepoint(value__0 _goml_m_std_p_json_p_JsonParser, builder__0 _goml_m_std_p_text_p_StringBuilder, codepoint__0 uint32) _goml_m_Result_____o__q_____string {
    var mtmp0 Option__char
    var inline9 Option__char = __goml_builtin_char_from_uint32(codepoint__0)
    mtmp0 = inline9
    switch mtmp0 != Option__char(0) {
    case false:
        var t0 string
        var inline0 string = "invalid unicode codepoint"
        var inline1 string = "" + inline0
        var inline2 string = inline1 + " at byte "
        var inline3 *ref_int_x = value__0.index
        var inline4 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline3)
        var inline5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline4)
        var inline6 string = inline2 + inline5
        t0 = inline6
        var t1 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
            _p0: t0,
            _tag: 1,
        }
        return t1
    case true:
        var x0 rune = rune(uint64(mtmp0) - 1)
        var inline7 string = _goml_m_inherent_i_char_i_char_i_to__string(x0)
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline7)
        var t2 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
            _tag: 0,
        }
        return t2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_json_p_parse__unicode__escape(value__0 _goml_m_std_p_json_p_JsonParser, builder__0 _goml_m_std_p_text_p_StringBuilder) _goml_m_Result_____o__q_____string {
    var mtmp0 Result__u32__string = _goml_m_std_p_json_p_parse__hex__quad(value__0)
    var jp0 uint32
    switch mtmp0._tag {
    case 0:
        var x2 uint32 = mtmp0._p0
        jp0 = x2
        var t0 bool = jp0 >= 55296
        var jp1 bool
        if t0 {
            var t39 bool = jp0 <= 56319
            jp1 = t39
        } else {
            jp1 = false
        }
        if jp1 {
            var t1 *ref_int_x = value__0.index
            var t2 int
            var inline27 int = ref_get__Ref_3int(t1)
            t2 = inline27
            var t3 int = t2 + 2
            var t4 string = value__0.input
            var t5 int
            var inline26 int = _goml_runtime_core_string_len(t4)
            t5 = inline26
            var t6 bool = t3 > t5
            var jp2 bool
            if t6 {
                jp2 = true
            } else {
                var t29 string = value__0.input
                var t30 *ref_int_x = value__0.index
                var t31 int
                var inline25 int = ref_get__Ref_3int(t30)
                t31 = inline25
                var t32 uint8
                var inline24 uint8 = _goml_runtime_core_string_byte_get(t29, t31)
                t32 = inline24
                var t33 bool = t32 != 92
                jp2 = t33
            }
            var jp3 bool
            if jp2 {
                jp3 = true
            } else {
                var t23 string = value__0.input
                var t24 *ref_int_x = value__0.index
                var t25 int
                var inline23 int = ref_get__Ref_3int(t24)
                t25 = inline23
                var t26 int = t25 + 1
                var t27 uint8
                var inline22 uint8 = _goml_runtime_core_string_byte_get(t23, t26)
                t27 = inline22
                var t28 bool = t27 != 117
                jp3 = t28
            }
            if jp3 {
                var t7 string
                var inline0 string = "missing low surrogate"
                var inline1 string = "" + inline0
                var inline2 string = inline1 + " at byte "
                var inline3 *ref_int_x = value__0.index
                var inline4 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline3)
                var inline5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline4)
                var inline6 string = inline2 + inline5
                t7 = inline6
                var t8 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                    _p0: t7,
                    _tag: 1,
                }
                return t8
            } else {
                var t9 *ref_int_x = value__0.index
                var t10 *ref_int_x = value__0.index
                var t11 int
                var inline21 int = ref_get__Ref_3int(t10)
                t11 = inline21
                var t12 int = t11 + 2
                ref_set__Ref_3int(t9, t12)
                var mtmp1 Result__u32__string = _goml_m_std_p_json_p_parse__hex__quad(value__0)
                var jp4 uint32
                switch mtmp1._tag {
                case 0:
                    var x0 uint32 = mtmp1._p0
                    jp4 = x0
                    var t13 bool = jp4 < 56320
                    var jp5 bool
                    if t13 {
                        jp5 = true
                    } else {
                        var t21 bool = jp4 > 57343
                        jp5 = t21
                    }
                    if jp5 {
                        var t14 string
                        var inline7 string = "invalid low surrogate"
                        var inline8 string = "" + inline7
                        var inline9 string = inline8 + " at byte "
                        var inline10 *ref_int_x = value__0.index
                        var inline11 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline10)
                        var inline12 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline11)
                        var inline13 string = inline9 + inline12
                        t14 = inline13
                        var t15 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                            _p0: t14,
                            _tag: 1,
                        }
                        return t15
                    } else {
                        var t16 uint32 = jp0 - 55296
                        var t17 uint32 = t16 * 1024
                        var t18 uint32 = 65536 + t17
                        var t19 uint32 = t18 + jp4
                        var t20 uint32 = t19 - 56320
                        var inline14 Option__char = char_from_u32(t20)
                        switch inline14 != Option__char(0) {
                        case false:
                            var inline15 string = _goml_m_std_p_json_p_json__error(value__0, "invalid unicode codepoint")
                            var inline16 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                                _p0: inline15,
                                _tag: 1,
                            }
                            return inline16
                        case true:
                            var inline17 rune = rune(uint64(inline14) - 1)
                            _goml_m_inherent_i_std_p_text__he83b4afafc069a3c24c64018b13ca033_r_i_write__char(builder__0, inline17)
                            var inline19 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                                _tag: 0,
                            }
                            return inline19
                        default:
                            panic("non-exhaustive match")
                        }
                    }
                case 1:
                    var x1 string = mtmp1._p1
                    var t22 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                        _p0: x1,
                        _tag: 1,
                    }
                    return t22
                default:
                    panic("non-exhaustive match")
                }
            }
        } else {
            var t34 bool = jp0 >= 56320
            var jp6 bool
            if t34 {
                var t38 bool = jp0 <= 57343
                jp6 = t38
            } else {
                jp6 = false
            }
            if jp6 {
                var t35 string
                var inline28 string = "unexpected low surrogate"
                var inline29 string = "" + inline28
                var inline30 string = inline29 + " at byte "
                var inline31 *ref_int_x = value__0.index
                var inline32 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline31)
                var inline33 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline32)
                var inline34 string = inline30 + inline33
                t35 = inline34
                var t36 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                    _p0: t35,
                    _tag: 1,
                }
                return t36
            } else {
                var t37 _goml_m_Result_____o__q_____string = _goml_m_std_p_json_p_write__codepoint(value__0, builder__0, jp0)
                return t37
            }
        }
    case 1:
        var x3 string = mtmp0._p1
        var t40 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
            _p0: x3,
            _tag: 1,
        }
        return t40
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_json_p_parse__json__string(value__0 _goml_m_std_p_json_p_JsonParser) Result__string__string {
    var t0 *ref_int_x = value__0.index
    var t1 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t0)
    var t2 string = value__0.input
    var t3 int = _goml_m_inherent_i_string_i_string_i_byte__len(t2)
    var t4 bool = t1 >= t3
    var jp0 bool
    if t4 {
        jp0 = true
    } else {
        var t92 string = value__0.input
        var t93 *ref_int_x = value__0.index
        var t94 int
        var inline33 int = ref_get__Ref_3int(t93)
        t94 = inline33
        var t95 uint8
        var inline32 uint8 = _goml_runtime_core_string_byte_get(t92, t94)
        t95 = inline32
        var t96 bool = t95 != 34
        jp0 = t96
    }
    if jp0 {
        var t86 string
        var inline23 string = "expected string"
        var inline24 string = "" + inline23
        var inline25 string = inline24 + " at byte "
        var inline26 *ref_int_x = value__0.index
        var inline27 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline26)
        var inline28 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline27)
        var inline29 string = inline25 + inline28
        t86 = inline29
        var t87 Result__string__string = Result__string__string{
            _p0: t86,
            _tag: 1,
        }
        return t87
    } else {
        var t88 *ref_int_x = value__0.index
        var t89 *ref_int_x = value__0.index
        var t90 int
        var inline31 int = ref_get__Ref_3int(t89)
        t90 = inline31
        var t91 int = t90 + 1
        ref_set__Ref_3int(t88, t91)
        var builder__0 _goml_m_std_p_text_p_StringBuilder = _goml_m_inherent_i_std_p_text_p_StringBuilder_i_std_p_text_p_StringBuilder_i_new()
        var t5 *ref_int_x = value__0.index
        var segment__0 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t5)
        Loop_loop0:
        for {
            var t8 *ref_int_x = value__0.index
            var t9 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t8)
            var t10 string = value__0.input
            var t11 int = _goml_m_inherent_i_string_i_string_i_byte__len(t10)
            var t12 bool = t9 < t11
            if t12 {
                var t13 string = value__0.input
                var t14 *ref_int_x = value__0.index
                var t15 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t14)
                var byte__0 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t13, t15)
                var t16 bool = byte__0 == 34
                if t16 {
                    var t17 *ref_int_x = value__0.index
                    var t18 int
                    var inline7 int = ref_get__Ref_3int(t17)
                    t18 = inline7
                    var t19 bool = segment__0 < t18
                    if t19 {
                        var t26 string = value__0.input
                        var t27 *ref_int_x = value__0.index
                        var t28 int
                        var inline6 int = ref_get__Ref_3int(t27)
                        t28 = inline6
                        var t29 string
                        var inline5 string = string_byte_slice(t26, segment__0, t28)
                        t29 = inline5
                        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, t29)
                    } else {}
                    var t20 *ref_int_x = value__0.index
                    var t21 *ref_int_x = value__0.index
                    var t22 int
                    var inline4 int = ref_get__Ref_3int(t21)
                    t22 = inline4
                    var t23 int = t22 + 1
                    ref_set__Ref_3int(t20, t23)
                    var t24 string
                    var inline0 *_goml_vec_uint8 = builder__0.values
                    var inline1 Tuple2_4bool_6string = string_from_utf8(inline0)
                    var inline2 string = inline1._1
                    t24 = inline2
                    var t25 Result__string__string = Result__string__string{
                        _p0: t24,
                        _tag: 0,
                    }
                    return t25
                } else {
                    var t31 bool = byte__0 == 92
                    if t31 {
                        var t32 *ref_int_x = value__0.index
                        var t33 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t32)
                        var t34 bool = segment__0 < t33
                        if t34 {
                            var t73 string = value__0.input
                            var t74 *ref_int_x = value__0.index
                            var t75 int
                            var inline22 int = ref_get__Ref_3int(t74)
                            t75 = inline22
                            var t76 string
                            var inline21 string = string_byte_slice(t73, segment__0, t75)
                            t76 = inline21
                            _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, t76)
                        } else {}
                        var t35 *ref_int_x = value__0.index
                        var t36 *ref_int_x = value__0.index
                        var t37 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t36)
                        var t38 int = t37 + 1
                        _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t35, t38)
                        var t39 *ref_int_x = value__0.index
                        var t40 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t39)
                        var t41 string = value__0.input
                        var t42 int = _goml_m_inherent_i_string_i_string_i_byte__len(t41)
                        var t43 bool = t40 >= t42
                        if t43 {
                            var t71 string
                            var inline14 string = "incomplete escape"
                            var inline15 string = "" + inline14
                            var inline16 string = inline15 + " at byte "
                            var inline17 *ref_int_x = value__0.index
                            var inline18 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline17)
                            var inline19 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline18)
                            var inline20 string = inline16 + inline19
                            t71 = inline20
                            var t72 Result__string__string = Result__string__string{
                                _p0: t71,
                                _tag: 1,
                            }
                            return t72
                        } else {
                            var t44 string = value__0.input
                            var t45 *ref_int_x = value__0.index
                            var t46 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t45)
                            var escape__0 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t44, t46)
                            var t47 *ref_int_x = value__0.index
                            var t48 *ref_int_x = value__0.index
                            var t49 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t48)
                            var t50 int = t49 + 1
                            _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t47, t50)
                            var t51 bool = escape__0 == 34
                            if t51 {
                                var inline8 rune = 34
                                var inline9 string = _goml_m_inherent_i_char_i_char_i_to__string(inline8)
                                _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline9)
                                var t52 *ref_int_x = value__0.index
                                var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                segment__0 = t53
                                continue
                            } else {
                                var t54 bool = escape__0 == 92
                                if t54 {
                                    var inline11 rune = 92
                                    var inline12 string = _goml_m_inherent_i_char_i_char_i_to__string(inline11)
                                    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline12)
                                    var t52 *ref_int_x = value__0.index
                                    var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                    segment__0 = t53
                                    continue
                                } else {
                                    var t55 bool = escape__0 == 47
                                    if t55 {
                                        _goml_m_inherent_i_std_p_text__he83b4afafc069a3c24c64018b13ca033_r_i_write__char(builder__0, 47)
                                        var t52 *ref_int_x = value__0.index
                                        var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                        segment__0 = t53
                                        continue
                                    } else {
                                        var t57 bool = escape__0 == 98
                                        if t57 {
                                            var mtmp0 Option__char = char_from_u32(8)
                                            switch mtmp0 != Option__char(0) {
                                            case true:
                                                var x0 rune = rune(uint64(mtmp0) - 1)
                                                _goml_m_inherent_i_std_p_text__he83b4afafc069a3c24c64018b13ca033_r_i_write__char(builder__0, x0)
                                                var t52 *ref_int_x = value__0.index
                                                var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                                segment__0 = t53
                                                continue
                                            default:
                                                var t52 *ref_int_x = value__0.index
                                                var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                                segment__0 = t53
                                                continue
                                            }
                                        } else {
                                            var t59 bool = escape__0 == 102
                                            if t59 {
                                                var mtmp1 Option__char = char_from_u32(12)
                                                switch mtmp1 != Option__char(0) {
                                                case true:
                                                    var x1 rune = rune(uint64(mtmp1) - 1)
                                                    _goml_m_inherent_i_std_p_text__he83b4afafc069a3c24c64018b13ca033_r_i_write__char(builder__0, x1)
                                                    var t52 *ref_int_x = value__0.index
                                                    var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                                    segment__0 = t53
                                                    continue
                                                default:
                                                    var t52 *ref_int_x = value__0.index
                                                    var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                                    segment__0 = t53
                                                    continue
                                                }
                                            } else {
                                                var t61 bool = escape__0 == 110
                                                if t61 {
                                                    _goml_m_inherent_i_std_p_text__he83b4afafc069a3c24c64018b13ca033_r_i_write__char(builder__0, 10)
                                                    var t52 *ref_int_x = value__0.index
                                                    var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                                    segment__0 = t53
                                                    continue
                                                } else {
                                                    var t63 bool = escape__0 == 114
                                                    if t63 {
                                                        _goml_m_inherent_i_std_p_text__he83b4afafc069a3c24c64018b13ca033_r_i_write__char(builder__0, 13)
                                                        var t52 *ref_int_x = value__0.index
                                                        var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                                        segment__0 = t53
                                                        continue
                                                    } else {
                                                        var t65 bool = escape__0 == 116
                                                        if t65 {
                                                            _goml_m_inherent_i_std_p_text__he83b4afafc069a3c24c64018b13ca033_r_i_write__char(builder__0, 9)
                                                            var t52 *ref_int_x = value__0.index
                                                            var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                                            segment__0 = t53
                                                            continue
                                                        } else {
                                                            var t67 bool = escape__0 == 117
                                                            if t67 {
                                                                var mtmp2 _goml_m_Result_____o__q_____string = _goml_m_std_p_json_p_parse__unicode__escape(value__0, builder__0)
                                                                switch mtmp2._tag {
                                                                case 0:
                                                                    var t52 *ref_int_x = value__0.index
                                                                    var t53 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t52)
                                                                    segment__0 = t53
                                                                    continue
                                                                case 1:
                                                                    var x2 string = mtmp2._p0
                                                                    var t68 Result__string__string = Result__string__string{
                                                                        _p0: x2,
                                                                        _tag: 1,
                                                                    }
                                                                    return t68
                                                                default:
                                                                    panic("non-exhaustive match")
                                                                }
                                                            } else {
                                                                var t69 string = _goml_m_std_p_json_p_json__error(value__0, "invalid escape")
                                                                var t70 Result__string__string = Result__string__string{
                                                                    _p0: t69,
                                                                    _tag: 1,
                                                                }
                                                                return t70
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    } else {
                        var t78 bool = byte__0 < 32
                        if t78 {
                            var t79 string = _goml_m_std_p_json_p_json__error(value__0, "unescaped control character")
                            var t80 Result__string__string = Result__string__string{
                                _p0: t79,
                                _tag: 1,
                            }
                            return t80
                        } else {
                            var t81 *ref_int_x = value__0.index
                            var t82 *ref_int_x = value__0.index
                            var t83 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t82)
                            var t84 int = t83 + 1
                            _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t81, t84)
                            continue
                        }
                    }
                }
            } else {
                break Loop_loop0
            }
        }
        var t6 string = _goml_m_std_p_json_p_json__error(value__0, "unterminated string")
        var t7 Result__string__string = Result__string__string{
            _p0: t6,
            _tag: 1,
        }
        return t7
    }
}

func _goml_m_std_p_json_p_parse__digits(value__0 _goml_m_std_p_json_p_JsonParser) bool {
    var t0 *ref_int_x = value__0.index
    var start__0 int
    var inline9 int = ref_get__Ref_3int(t0)
    start__0 = inline9
    Loop_loop0:
    for {
        var t4 *ref_int_x = value__0.index
        var t5 int
        var inline8 int = ref_get__Ref_3int(t4)
        t5 = inline8
        var t6 string = value__0.input
        var t7 int
        var inline7 int = _goml_runtime_core_string_len(t6)
        t7 = inline7
        var t8 bool = t5 < t7
        var jp0 bool
        if t8 {
            var t13 string = value__0.input
            var t14 *ref_int_x = value__0.index
            var t15 int
            var inline6 int = ref_get__Ref_3int(t14)
            t15 = inline6
            var t16 uint8
            var inline5 uint8 = _goml_runtime_core_string_byte_get(t13, t15)
            t16 = inline5
            var inline3 bool = t16 >= 48
            if inline3 {
                var inline4 bool = t16 <= 57
                jp0 = inline4
            } else {
                jp0 = false
            }
        } else {
            jp0 = false
        }
        if jp0 {
            var t9 *ref_int_x = value__0.index
            var t10 *ref_int_x = value__0.index
            var t11 int
            var inline2 int = ref_get__Ref_3int(t10)
            t11 = inline2
            var t12 int = t11 + 1
            ref_set__Ref_3int(t9, t12)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t1 *ref_int_x = value__0.index
    var t2 int
    var inline0 int = ref_get__Ref_3int(t1)
    t2 = inline0
    var t3 bool = t2 > start__0
    return t3
}

func _goml_m_std_p_json_p_parse__json__number__text(value__0 _goml_m_std_p_json_p_JsonParser) Result__string__string {
    var t0 *ref_int_x = value__0.index
    var start__0 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t0)
    var t1 string = value__0.input
    var t2 *ref_int_x = value__0.index
    var t3 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t2)
    var t4 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t1, t3)
    var t5 bool = t4 == 45
    if t5 {
        var t103 *ref_int_x = value__0.index
        var t104 *ref_int_x = value__0.index
        var t105 int
        var inline33 int = ref_get__Ref_3int(t104)
        t105 = inline33
        var t106 int = t105 + 1
        ref_set__Ref_3int(t103, t106)
    } else {}
    var t6 *ref_int_x = value__0.index
    var t7 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t6)
    var t8 string = value__0.input
    var t9 int = _goml_m_inherent_i_string_i_string_i_byte__len(t8)
    var t10 bool = t7 >= t9
    if t10 {
        var t77 string
        var inline1 string = "incomplete number"
        var inline2 string = "" + inline1
        var inline3 string = inline2 + " at byte "
        var inline4 *ref_int_x = value__0.index
        var inline5 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline4)
        var inline6 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline5)
        var inline7 string = inline3 + inline6
        t77 = inline7
        var t78 Result__string__string = Result__string__string{
            _p0: t77,
            _tag: 1,
        }
        return t78
    } else {
        var t79 string = value__0.input
        var t80 *ref_int_x = value__0.index
        var t81 int
        var inline31 int = ref_get__Ref_3int(t80)
        t81 = inline31
        var t82 uint8
        var inline30 uint8 = _goml_runtime_core_string_byte_get(t79, t81)
        t82 = inline30
        var t83 bool = t82 == 48
        if t83 {
            var t84 *ref_int_x = value__0.index
            var t85 *ref_int_x = value__0.index
            var t86 int
            var inline22 int = ref_get__Ref_3int(t85)
            t86 = inline22
            var t87 int = t86 + 1
            ref_set__Ref_3int(t84, t87)
            var t88 *ref_int_x = value__0.index
            var t89 int
            var inline20 int = ref_get__Ref_3int(t88)
            t89 = inline20
            var t90 string = value__0.input
            var t91 int
            var inline19 int = _goml_runtime_core_string_len(t90)
            t91 = inline19
            var t92 bool = t89 < t91
            var jp3 bool
            if t92 {
                var t95 string = value__0.input
                var t96 *ref_int_x = value__0.index
                var t97 int
                var inline18 int = ref_get__Ref_3int(t96)
                t97 = inline18
                var t98 uint8
                var inline17 uint8 = _goml_runtime_core_string_byte_get(t95, t97)
                t98 = inline17
                var inline15 bool = t98 >= 48
                if inline15 {
                    var inline16 bool = t98 <= 57
                    jp3 = inline16
                } else {
                    jp3 = false
                }
            } else {
                jp3 = false
            }
            if jp3 {
                var t93 string
                var inline8 string = "invalid leading zero"
                var inline9 string = "" + inline8
                var inline10 string = inline9 + " at byte "
                var inline11 *ref_int_x = value__0.index
                var inline12 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline11)
                var inline13 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline12)
                var inline14 string = inline10 + inline13
                t93 = inline14
                var t94 Result__string__string = Result__string__string{
                    _p0: t93,
                    _tag: 1,
                }
                return t94
            } else {
                var t11 *ref_int_x = value__0.index
                var t12 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t11)
                var t13 string = value__0.input
                var t14 int = _goml_m_inherent_i_string_i_string_i_byte__len(t13)
                var t15 bool = t12 < t14
                var jp0 bool
                if t15 {
                    var t72 string = value__0.input
                    var t73 *ref_int_x = value__0.index
                    var t74 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t73)
                    var t75 uint8
                    var inline0 uint8 = _goml_runtime_core_string_byte_get(t72, t74)
                    t75 = inline0
                    var t76 bool = t75 == 46
                    jp0 = t76
                } else {
                    jp0 = false
                }
                if jp0 {
                    var t64 *ref_int_x = value__0.index
                    var t65 *ref_int_x = value__0.index
                    var t66 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t65)
                    var t67 int = t66 + 1
                    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t64, t67)
                    var t68 bool = _goml_m_std_p_json_p_parse__digits(value__0)
                    var t69 bool = !t68
                    if t69 {
                        var t70 string = _goml_m_std_p_json_p_json__error(value__0, "missing fraction digits")
                        var t71 Result__string__string = Result__string__string{
                            _p0: t70,
                            _tag: 1,
                        }
                        return t71
                    } else {
                        var t16 *ref_int_x = value__0.index
                        var t17 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t16)
                        var t18 string = value__0.input
                        var t19 int = _goml_m_inherent_i_string_i_string_i_byte__len(t18)
                        var t20 bool = t17 < t19
                        var jp1 bool
                        if t20 {
                            var t54 string = value__0.input
                            var t55 *ref_int_x = value__0.index
                            var t56 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t55)
                            var t57 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t54, t56)
                            var t58 bool = t57 == 101
                            if t58 {
                                jp1 = true
                            } else {
                                var t59 string = value__0.input
                                var t60 *ref_int_x = value__0.index
                                var t61 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t60)
                                var t62 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t59, t61)
                                var t63 bool = t62 == 69
                                jp1 = t63
                            }
                        } else {
                            jp1 = false
                        }
                        if jp1 {
                            var t26 *ref_int_x = value__0.index
                            var t27 *ref_int_x = value__0.index
                            var t28 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t27)
                            var t29 int = t28 + 1
                            _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t26, t29)
                            var t30 *ref_int_x = value__0.index
                            var t31 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t30)
                            var t32 string = value__0.input
                            var t33 int = _goml_m_inherent_i_string_i_string_i_byte__len(t32)
                            var t34 bool = t31 < t33
                            var jp2 bool
                            if t34 {
                                var t44 string = value__0.input
                                var t45 *ref_int_x = value__0.index
                                var t46 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t45)
                                var t47 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t44, t46)
                                var t48 bool = t47 == 43
                                if t48 {
                                    jp2 = true
                                } else {
                                    var t49 string = value__0.input
                                    var t50 *ref_int_x = value__0.index
                                    var t51 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t50)
                                    var t52 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t49, t51)
                                    var t53 bool = t52 == 45
                                    jp2 = t53
                                }
                            } else {
                                jp2 = false
                            }
                            if jp2 {
                                var t39 *ref_int_x = value__0.index
                                var t40 *ref_int_x = value__0.index
                                var t41 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t40)
                                var t42 int = t41 + 1
                                _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t39, t42)
                            } else {}
                            var t35 bool = _goml_m_std_p_json_p_parse__digits(value__0)
                            var t36 bool = !t35
                            if t36 {
                                var t37 string = _goml_m_std_p_json_p_json__error(value__0, "missing exponent digits")
                                var t38 Result__string__string = Result__string__string{
                                    _p0: t37,
                                    _tag: 1,
                                }
                                return t38
                            } else {
                                var t21 string = value__0.input
                                var t22 *ref_int_x = value__0.index
                                var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                                var t24 string = _goml_m_inherent_i_string_i_string_i_byte__slice(t21, start__0, t23)
                                var t25 Result__string__string = Result__string__string{
                                    _p0: t24,
                                    _tag: 0,
                                }
                                return t25
                            }
                        } else {
                            var t21 string = value__0.input
                            var t22 *ref_int_x = value__0.index
                            var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                            var t24 string = _goml_m_inherent_i_string_i_string_i_byte__slice(t21, start__0, t23)
                            var t25 Result__string__string = Result__string__string{
                                _p0: t24,
                                _tag: 0,
                            }
                            return t25
                        }
                    }
                } else {
                    var t16 *ref_int_x = value__0.index
                    var t17 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t16)
                    var t18 string = value__0.input
                    var t19 int = _goml_m_inherent_i_string_i_string_i_byte__len(t18)
                    var t20 bool = t17 < t19
                    var jp1 bool
                    if t20 {
                        var t54 string = value__0.input
                        var t55 *ref_int_x = value__0.index
                        var t56 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t55)
                        var t57 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t54, t56)
                        var t58 bool = t57 == 101
                        if t58 {
                            jp1 = true
                        } else {
                            var t59 string = value__0.input
                            var t60 *ref_int_x = value__0.index
                            var t61 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t60)
                            var t62 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t59, t61)
                            var t63 bool = t62 == 69
                            jp1 = t63
                        }
                    } else {
                        jp1 = false
                    }
                    if jp1 {
                        var t26 *ref_int_x = value__0.index
                        var t27 *ref_int_x = value__0.index
                        var t28 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t27)
                        var t29 int = t28 + 1
                        _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t26, t29)
                        var t30 *ref_int_x = value__0.index
                        var t31 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t30)
                        var t32 string = value__0.input
                        var t33 int = _goml_m_inherent_i_string_i_string_i_byte__len(t32)
                        var t34 bool = t31 < t33
                        var jp2 bool
                        if t34 {
                            var t44 string = value__0.input
                            var t45 *ref_int_x = value__0.index
                            var t46 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t45)
                            var t47 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t44, t46)
                            var t48 bool = t47 == 43
                            if t48 {
                                jp2 = true
                            } else {
                                var t49 string = value__0.input
                                var t50 *ref_int_x = value__0.index
                                var t51 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t50)
                                var t52 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t49, t51)
                                var t53 bool = t52 == 45
                                jp2 = t53
                            }
                        } else {
                            jp2 = false
                        }
                        if jp2 {
                            var t39 *ref_int_x = value__0.index
                            var t40 *ref_int_x = value__0.index
                            var t41 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t40)
                            var t42 int = t41 + 1
                            _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t39, t42)
                        } else {}
                        var t35 bool = _goml_m_std_p_json_p_parse__digits(value__0)
                        var t36 bool = !t35
                        if t36 {
                            var t37 string = _goml_m_std_p_json_p_json__error(value__0, "missing exponent digits")
                            var t38 Result__string__string = Result__string__string{
                                _p0: t37,
                                _tag: 1,
                            }
                            return t38
                        } else {
                            var t21 string = value__0.input
                            var t22 *ref_int_x = value__0.index
                            var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                            var t24 string = _goml_m_inherent_i_string_i_string_i_byte__slice(t21, start__0, t23)
                            var t25 Result__string__string = Result__string__string{
                                _p0: t24,
                                _tag: 0,
                            }
                            return t25
                        }
                    } else {
                        var t21 string = value__0.input
                        var t22 *ref_int_x = value__0.index
                        var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                        var t24 string = _goml_m_inherent_i_string_i_string_i_byte__slice(t21, start__0, t23)
                        var t25 Result__string__string = Result__string__string{
                            _p0: t24,
                            _tag: 0,
                        }
                        return t25
                    }
                }
            }
        } else {
            var t99 bool = _goml_m_std_p_json_p_parse__digits(value__0)
            var t100 bool = !t99
            if t100 {
                var t101 string
                var inline23 string = "expected number"
                var inline24 string = "" + inline23
                var inline25 string = inline24 + " at byte "
                var inline26 *ref_int_x = value__0.index
                var inline27 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline26)
                var inline28 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline27)
                var inline29 string = inline25 + inline28
                t101 = inline29
                var t102 Result__string__string = Result__string__string{
                    _p0: t101,
                    _tag: 1,
                }
                return t102
            } else {
                var t11 *ref_int_x = value__0.index
                var t12 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t11)
                var t13 string = value__0.input
                var t14 int = _goml_m_inherent_i_string_i_string_i_byte__len(t13)
                var t15 bool = t12 < t14
                var jp0 bool
                if t15 {
                    var t72 string = value__0.input
                    var t73 *ref_int_x = value__0.index
                    var t74 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t73)
                    var t75 uint8
                    var inline0 uint8 = _goml_runtime_core_string_byte_get(t72, t74)
                    t75 = inline0
                    var t76 bool = t75 == 46
                    jp0 = t76
                } else {
                    jp0 = false
                }
                if jp0 {
                    var t64 *ref_int_x = value__0.index
                    var t65 *ref_int_x = value__0.index
                    var t66 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t65)
                    var t67 int = t66 + 1
                    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t64, t67)
                    var t68 bool = _goml_m_std_p_json_p_parse__digits(value__0)
                    var t69 bool = !t68
                    if t69 {
                        var t70 string = _goml_m_std_p_json_p_json__error(value__0, "missing fraction digits")
                        var t71 Result__string__string = Result__string__string{
                            _p0: t70,
                            _tag: 1,
                        }
                        return t71
                    } else {
                        var t16 *ref_int_x = value__0.index
                        var t17 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t16)
                        var t18 string = value__0.input
                        var t19 int = _goml_m_inherent_i_string_i_string_i_byte__len(t18)
                        var t20 bool = t17 < t19
                        var jp1 bool
                        if t20 {
                            var t54 string = value__0.input
                            var t55 *ref_int_x = value__0.index
                            var t56 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t55)
                            var t57 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t54, t56)
                            var t58 bool = t57 == 101
                            if t58 {
                                jp1 = true
                            } else {
                                var t59 string = value__0.input
                                var t60 *ref_int_x = value__0.index
                                var t61 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t60)
                                var t62 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t59, t61)
                                var t63 bool = t62 == 69
                                jp1 = t63
                            }
                        } else {
                            jp1 = false
                        }
                        if jp1 {
                            var t26 *ref_int_x = value__0.index
                            var t27 *ref_int_x = value__0.index
                            var t28 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t27)
                            var t29 int = t28 + 1
                            _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t26, t29)
                            var t30 *ref_int_x = value__0.index
                            var t31 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t30)
                            var t32 string = value__0.input
                            var t33 int = _goml_m_inherent_i_string_i_string_i_byte__len(t32)
                            var t34 bool = t31 < t33
                            var jp2 bool
                            if t34 {
                                var t44 string = value__0.input
                                var t45 *ref_int_x = value__0.index
                                var t46 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t45)
                                var t47 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t44, t46)
                                var t48 bool = t47 == 43
                                if t48 {
                                    jp2 = true
                                } else {
                                    var t49 string = value__0.input
                                    var t50 *ref_int_x = value__0.index
                                    var t51 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t50)
                                    var t52 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t49, t51)
                                    var t53 bool = t52 == 45
                                    jp2 = t53
                                }
                            } else {
                                jp2 = false
                            }
                            if jp2 {
                                var t39 *ref_int_x = value__0.index
                                var t40 *ref_int_x = value__0.index
                                var t41 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t40)
                                var t42 int = t41 + 1
                                _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t39, t42)
                            } else {}
                            var t35 bool = _goml_m_std_p_json_p_parse__digits(value__0)
                            var t36 bool = !t35
                            if t36 {
                                var t37 string = _goml_m_std_p_json_p_json__error(value__0, "missing exponent digits")
                                var t38 Result__string__string = Result__string__string{
                                    _p0: t37,
                                    _tag: 1,
                                }
                                return t38
                            } else {
                                var t21 string = value__0.input
                                var t22 *ref_int_x = value__0.index
                                var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                                var t24 string = _goml_m_inherent_i_string_i_string_i_byte__slice(t21, start__0, t23)
                                var t25 Result__string__string = Result__string__string{
                                    _p0: t24,
                                    _tag: 0,
                                }
                                return t25
                            }
                        } else {
                            var t21 string = value__0.input
                            var t22 *ref_int_x = value__0.index
                            var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                            var t24 string = _goml_m_inherent_i_string_i_string_i_byte__slice(t21, start__0, t23)
                            var t25 Result__string__string = Result__string__string{
                                _p0: t24,
                                _tag: 0,
                            }
                            return t25
                        }
                    }
                } else {
                    var t16 *ref_int_x = value__0.index
                    var t17 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t16)
                    var t18 string = value__0.input
                    var t19 int = _goml_m_inherent_i_string_i_string_i_byte__len(t18)
                    var t20 bool = t17 < t19
                    var jp1 bool
                    if t20 {
                        var t54 string = value__0.input
                        var t55 *ref_int_x = value__0.index
                        var t56 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t55)
                        var t57 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t54, t56)
                        var t58 bool = t57 == 101
                        if t58 {
                            jp1 = true
                        } else {
                            var t59 string = value__0.input
                            var t60 *ref_int_x = value__0.index
                            var t61 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t60)
                            var t62 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t59, t61)
                            var t63 bool = t62 == 69
                            jp1 = t63
                        }
                    } else {
                        jp1 = false
                    }
                    if jp1 {
                        var t26 *ref_int_x = value__0.index
                        var t27 *ref_int_x = value__0.index
                        var t28 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t27)
                        var t29 int = t28 + 1
                        _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t26, t29)
                        var t30 *ref_int_x = value__0.index
                        var t31 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t30)
                        var t32 string = value__0.input
                        var t33 int = _goml_m_inherent_i_string_i_string_i_byte__len(t32)
                        var t34 bool = t31 < t33
                        var jp2 bool
                        if t34 {
                            var t44 string = value__0.input
                            var t45 *ref_int_x = value__0.index
                            var t46 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t45)
                            var t47 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t44, t46)
                            var t48 bool = t47 == 43
                            if t48 {
                                jp2 = true
                            } else {
                                var t49 string = value__0.input
                                var t50 *ref_int_x = value__0.index
                                var t51 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t50)
                                var t52 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t49, t51)
                                var t53 bool = t52 == 45
                                jp2 = t53
                            }
                        } else {
                            jp2 = false
                        }
                        if jp2 {
                            var t39 *ref_int_x = value__0.index
                            var t40 *ref_int_x = value__0.index
                            var t41 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t40)
                            var t42 int = t41 + 1
                            _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t39, t42)
                        } else {}
                        var t35 bool = _goml_m_std_p_json_p_parse__digits(value__0)
                        var t36 bool = !t35
                        if t36 {
                            var t37 string = _goml_m_std_p_json_p_json__error(value__0, "missing exponent digits")
                            var t38 Result__string__string = Result__string__string{
                                _p0: t37,
                                _tag: 1,
                            }
                            return t38
                        } else {
                            var t21 string = value__0.input
                            var t22 *ref_int_x = value__0.index
                            var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                            var t24 string = _goml_m_inherent_i_string_i_string_i_byte__slice(t21, start__0, t23)
                            var t25 Result__string__string = Result__string__string{
                                _p0: t24,
                                _tag: 0,
                            }
                            return t25
                        }
                    } else {
                        var t21 string = value__0.input
                        var t22 *ref_int_x = value__0.index
                        var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                        var t24 string = _goml_m_inherent_i_string_i_string_i_byte__slice(t21, start__0, t23)
                        var t25 Result__string__string = Result__string__string{
                            _p0: t24,
                            _tag: 0,
                        }
                        return t25
                    }
                }
            }
        }
    }
}

func _goml_m_std_p_json_p_parse__json__literal(value__0 _goml_m_std_p_json_p_JsonParser, expected__0 string, result__0 _goml_m_std_p_json_p_Value) _goml_m_Result____std_p_json_p_Value____string {
    var t0 *ref_int_x = value__0.index
    var t1 int
    var inline16 int = ref_get__Ref_3int(t0)
    t1 = inline16
    var t2 int
    var inline15 int = _goml_runtime_core_string_len(expected__0)
    t2 = inline15
    var t3 int = t1 + t2
    var t4 string = value__0.input
    var t5 int
    var inline14 int = _goml_runtime_core_string_len(t4)
    t5 = inline14
    var t6 bool = t3 <= t5
    var jp0 bool
    if t6 {
        var t15 string = value__0.input
        var t16 *ref_int_x = value__0.index
        var t17 int
        var inline13 int = ref_get__Ref_3int(t16)
        t17 = inline13
        var t18 *ref_int_x = value__0.index
        var t19 int
        var inline12 int = ref_get__Ref_3int(t18)
        t19 = inline12
        var t20 int
        var inline11 int = _goml_runtime_core_string_len(expected__0)
        t20 = inline11
        var t21 int = t19 + t20
        var t22 string
        var inline10 string = string_byte_slice(t15, t17, t21)
        t22 = inline10
        var t23 bool = t22 == expected__0
        jp0 = t23
    } else {
        jp0 = false
    }
    if jp0 {
        var t7 *ref_int_x = value__0.index
        var t8 *ref_int_x = value__0.index
        var t9 int
        var inline2 int = ref_get__Ref_3int(t8)
        t9 = inline2
        var t10 int
        var inline1 int = _goml_runtime_core_string_len(expected__0)
        t10 = inline1
        var t11 int = t9 + t10
        ref_set__Ref_3int(t7, t11)
        var t12 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
            _p0: result__0,
            _tag: 0,
        }
        return t12
    } else {
        var t13 string
        var inline3 string = "invalid literal"
        var inline4 string = "" + inline3
        var inline5 string = inline4 + " at byte "
        var inline6 *ref_int_x = value__0.index
        var inline7 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline6)
        var inline8 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline7)
        var inline9 string = inline5 + inline8
        t13 = inline9
        var t14 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
            _p1: t13,
            _tag: 1,
        }
        return t14
    }
}

func _goml_m_std_p_json_p_parse__json__array(value__0 _goml_m_std_p_json_p_JsonParser) _goml_m_Result____std_p_json_p_Value____string {
    var t0 *ref_int_x = value__0.index
    var t1 *ref_int_x = value__0.index
    var t2 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t1)
    var t3 int = t2 + 1
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t0, t3)
    _goml_m_std_p_json_p_skip__json__whitespace(value__0)
    var result__0 *_goml_vec__goml_m_std_p_json_p_Value = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__std_p_json_p_Value()
    var t4 *ref_int_x = value__0.index
    var t5 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t4)
    var t6 string = value__0.input
    var t7 int = _goml_m_inherent_i_string_i_string_i_byte__len(t6)
    var t8 bool = t5 < t7
    var jp0 bool
    if t8 {
        var t53 string = value__0.input
        var t54 *ref_int_x = value__0.index
        var t55 int
        var inline30 int = ref_get__Ref_3int(t54)
        t55 = inline30
        var t56 uint8
        var inline29 uint8 = _goml_runtime_core_string_byte_get(t53, t55)
        t56 = inline29
        var t57 bool = t56 == 93
        jp0 = t57
    } else {
        jp0 = false
    }
    if jp0 {
        var t47 *ref_int_x = value__0.index
        var t48 *ref_int_x = value__0.index
        var t49 int
        var inline28 int = ref_get__Ref_3int(t48)
        t49 = inline28
        var t50 int = t49 + 1
        ref_set__Ref_3int(t47, t50)
        var t51 _goml_m_std_p_json_p_Value = _goml_m_std_p_json_p_Value_Array{
            _0: result__0,
        }
        var t52 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
            _p0: t51,
            _tag: 0,
        }
        return t52
    } else {
        Loop_loop0:
        for {
            var t11 *ref_int_x = value__0.index
            var t12 int
            var inline26 int = ref_get__Ref_3int(t11)
            t12 = inline26
            var t13 string = value__0.input
            var t14 int
            var inline25 int = _goml_runtime_core_string_len(t13)
            t14 = inline25
            var t15 bool = t12 < t14
            if t15 {
                var mtmp0 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse__json__value(value__0)
                var jp1 _goml_m_std_p_json_p_Value
                switch mtmp0._tag {
                case 0:
                    var x0 _goml_m_std_p_json_p_Value = mtmp0._p0
                    jp1 = x0
                    vec_push___goml_m_Vec__16std_p_json_p_Value(result__0, jp1)
                    _goml_m_std_p_json_p_skip__json__whitespace(value__0)
                    var t16 *ref_int_x = value__0.index
                    var t17 int
                    var inline23 int = ref_get__Ref_3int(t16)
                    t17 = inline23
                    var t18 string = value__0.input
                    var t19 int
                    var inline22 int = _goml_runtime_core_string_len(t18)
                    t19 = inline22
                    var t20 bool = t17 >= t19
                    if t20 {
                        var t21 string
                        var inline0 string = "unterminated array"
                        var inline1 string = "" + inline0
                        var inline2 string = inline1 + " at byte "
                        var inline3 *ref_int_x = value__0.index
                        var inline4 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline3)
                        var inline5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline4)
                        var inline6 string = inline2 + inline5
                        t21 = inline6
                        var t22 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                            _p1: t21,
                            _tag: 1,
                        }
                        return t22
                    } else {
                        var t23 string = value__0.input
                        var t24 *ref_int_x = value__0.index
                        var t25 int
                        var inline21 int = ref_get__Ref_3int(t24)
                        t25 = inline21
                        var t26 uint8
                        var inline20 uint8 = _goml_runtime_core_string_byte_get(t23, t25)
                        t26 = inline20
                        var t27 bool = t26 == 93
                        if t27 {
                            var t28 *ref_int_x = value__0.index
                            var t29 *ref_int_x = value__0.index
                            var t30 int
                            var inline8 int = ref_get__Ref_3int(t29)
                            t30 = inline8
                            var t31 int = t30 + 1
                            ref_set__Ref_3int(t28, t31)
                            var t32 _goml_m_std_p_json_p_Value = _goml_m_std_p_json_p_Value_Array{
                                _0: result__0,
                            }
                            var t33 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                                _p0: t32,
                                _tag: 0,
                            }
                            return t33
                        } else {
                            var t34 string = value__0.input
                            var t35 *ref_int_x = value__0.index
                            var t36 int
                            var inline19 int = ref_get__Ref_3int(t35)
                            t36 = inline19
                            var t37 uint8
                            var inline18 uint8 = _goml_runtime_core_string_byte_get(t34, t36)
                            t37 = inline18
                            var t38 bool = t37 == 44
                            if t38 {
                                var t39 *ref_int_x = value__0.index
                                var t40 *ref_int_x = value__0.index
                                var t41 int
                                var inline10 int = ref_get__Ref_3int(t40)
                                t41 = inline10
                                var t42 int = t41 + 1
                                ref_set__Ref_3int(t39, t42)
                                _goml_m_std_p_json_p_skip__json__whitespace(value__0)
                                continue
                            } else {
                                var t44 string
                                var inline11 string = "expected array separator"
                                var inline12 string = "" + inline11
                                var inline13 string = inline12 + " at byte "
                                var inline14 *ref_int_x = value__0.index
                                var inline15 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline14)
                                var inline16 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline15)
                                var inline17 string = inline13 + inline16
                                t44 = inline17
                                var t45 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                                    _p1: t44,
                                    _tag: 1,
                                }
                                return t45
                            }
                        }
                    }
                case 1:
                    var x1 string = mtmp0._p1
                    var t46 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                        _p1: x1,
                        _tag: 1,
                    }
                    return t46
                default:
                    panic("non-exhaustive match")
                }
            } else {
                break Loop_loop0
            }
        }
        var t9 string = _goml_m_std_p_json_p_json__error(value__0, "unterminated array")
        var t10 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
            _p1: t9,
            _tag: 1,
        }
        return t10
    }
}

func _goml_m_std_p_json_p_parse__json__object(value__0 _goml_m_std_p_json_p_JsonParser) _goml_m_Result____std_p_json_p_Value____string {
    var t0 *ref_int_x = value__0.index
    var t1 *ref_int_x = value__0.index
    var t2 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t1)
    var t3 int = t2 + 1
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t0, t3)
    _goml_m_std_p_json_p_skip__json__whitespace(value__0)
    var result__0 *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value = _goml_m_inherent_i_Vec_i_Vec_l_h57e42dbef834e6b8ee6cf77cf9eb9d23_json_p_Value_q_()
    var t4 *ref_int_x = value__0.index
    var t5 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t4)
    var t6 string = value__0.input
    var t7 int = _goml_m_inherent_i_string_i_string_i_byte__len(t6)
    var t8 bool = t5 < t7
    var jp0 bool
    if t8 {
        var t71 string = value__0.input
        var t72 *ref_int_x = value__0.index
        var t73 int
        var inline33 int = ref_get__Ref_3int(t72)
        t73 = inline33
        var t74 uint8
        var inline32 uint8 = _goml_runtime_core_string_byte_get(t71, t73)
        t74 = inline32
        var t75 bool = t74 == 125
        jp0 = t75
    } else {
        jp0 = false
    }
    if jp0 {
        var t65 *ref_int_x = value__0.index
        var t66 *ref_int_x = value__0.index
        var t67 int
        var inline31 int = ref_get__Ref_3int(t66)
        t67 = inline31
        var t68 int = t67 + 1
        ref_set__Ref_3int(t65, t68)
        var t69 _goml_m_std_p_json_p_Value = Object{
            _0: result__0,
        }
        var t70 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
            _p0: t69,
            _tag: 0,
        }
        return t70
    } else {
        Loop_loop0:
        for {
            var t11 *ref_int_x = value__0.index
            var t12 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t11)
            var t13 string = value__0.input
            var t14 int = _goml_m_inherent_i_string_i_string_i_byte__len(t13)
            var t15 bool = t12 < t14
            if t15 {
                var mtmp0 Result__string__string = _goml_m_std_p_json_p_parse__json__string(value__0)
                var jp1 string
                switch mtmp0._tag {
                case 0:
                    var x2 string = mtmp0._p0
                    jp1 = x2
                    _goml_m_std_p_json_p_skip__json__whitespace(value__0)
                    var t16 *ref_int_x = value__0.index
                    var t17 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t16)
                    var t18 string = value__0.input
                    var t19 int = _goml_m_inherent_i_string_i_string_i_byte__len(t18)
                    var t20 bool = t17 >= t19
                    var jp2 bool
                    if t20 {
                        jp2 = true
                    } else {
                        var t59 string = value__0.input
                        var t60 *ref_int_x = value__0.index
                        var t61 int
                        var inline29 int = ref_get__Ref_3int(t60)
                        t61 = inline29
                        var t62 uint8
                        var inline28 uint8 = _goml_runtime_core_string_byte_get(t59, t61)
                        t62 = inline28
                        var t63 bool = t62 != 58
                        jp2 = t63
                    }
                    if jp2 {
                        var t53 string
                        var inline19 string = "expected object colon"
                        var inline20 string = "" + inline19
                        var inline21 string = inline20 + " at byte "
                        var inline22 *ref_int_x = value__0.index
                        var inline23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline22)
                        var inline24 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline23)
                        var inline25 string = inline21 + inline24
                        t53 = inline25
                        var t54 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                            _p1: t53,
                            _tag: 1,
                        }
                        return t54
                    } else {
                        var t55 *ref_int_x = value__0.index
                        var t56 *ref_int_x = value__0.index
                        var t57 int
                        var inline27 int = ref_get__Ref_3int(t56)
                        t57 = inline27
                        var t58 int = t57 + 1
                        ref_set__Ref_3int(t55, t58)
                        _goml_m_std_p_json_p_skip__json__whitespace(value__0)
                        var mtmp1 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse__json__value(value__0)
                        var jp3 _goml_m_std_p_json_p_Value
                        switch mtmp1._tag {
                        case 0:
                            var x0 _goml_m_std_p_json_p_Value = mtmp1._p0
                            jp3 = x0
                            var t21 Tuple2_6string_26_goml_m_std_p_json_p_Value = Tuple2_6string_26_goml_m_std_p_json_p_Value{
                                _0: jp1,
                                _1: jp3,
                            }
                            _goml_m_inherent_i_Vec_i_Vec_l_h153048c8bd06f0dfabad32cddaecb150_json_p_Value_q_(result__0, t21)
                            _goml_m_std_p_json_p_skip__json__whitespace(value__0)
                            var t22 *ref_int_x = value__0.index
                            var t23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t22)
                            var t24 string = value__0.input
                            var t25 int = _goml_m_inherent_i_string_i_string_i_byte__len(t24)
                            var t26 bool = t23 >= t25
                            if t26 {
                                var t27 string
                                var inline0 string = "unterminated object"
                                var inline1 string = "" + inline0
                                var inline2 string = inline1 + " at byte "
                                var inline3 *ref_int_x = value__0.index
                                var inline4 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline3)
                                var inline5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline4)
                                var inline6 string = inline2 + inline5
                                t27 = inline6
                                var t28 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                                    _p1: t27,
                                    _tag: 1,
                                }
                                return t28
                            } else {
                                var t29 string = value__0.input
                                var t30 *ref_int_x = value__0.index
                                var t31 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t30)
                                var t32 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(t29, t31)
                                var t33 bool = t32 == 125
                                if t33 {
                                    var t34 *ref_int_x = value__0.index
                                    var t35 *ref_int_x = value__0.index
                                    var t36 int
                                    var inline8 int = ref_get__Ref_3int(t35)
                                    t36 = inline8
                                    var t37 int = t36 + 1
                                    ref_set__Ref_3int(t34, t37)
                                    var t38 _goml_m_std_p_json_p_Value = Object{
                                        _0: result__0,
                                    }
                                    var t39 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                                        _p0: t38,
                                        _tag: 0,
                                    }
                                    return t39
                                } else {
                                    var t40 string = value__0.input
                                    var t41 *ref_int_x = value__0.index
                                    var t42 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t41)
                                    var t43 uint8
                                    var inline18 uint8 = _goml_runtime_core_string_byte_get(t40, t42)
                                    t43 = inline18
                                    var t44 bool = t43 == 44
                                    if t44 {
                                        var t45 *ref_int_x = value__0.index
                                        var t46 *ref_int_x = value__0.index
                                        var t47 int
                                        var inline10 int = ref_get__Ref_3int(t46)
                                        t47 = inline10
                                        var t48 int = t47 + 1
                                        ref_set__Ref_3int(t45, t48)
                                        _goml_m_std_p_json_p_skip__json__whitespace(value__0)
                                        continue
                                    } else {
                                        var t50 string
                                        var inline11 string = "expected object separator"
                                        var inline12 string = "" + inline11
                                        var inline13 string = inline12 + " at byte "
                                        var inline14 *ref_int_x = value__0.index
                                        var inline15 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline14)
                                        var inline16 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline15)
                                        var inline17 string = inline13 + inline16
                                        t50 = inline17
                                        var t51 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                                            _p1: t50,
                                            _tag: 1,
                                        }
                                        return t51
                                    }
                                }
                            }
                        case 1:
                            var x1 string = mtmp1._p1
                            var t52 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                                _p1: x1,
                                _tag: 1,
                            }
                            return t52
                        default:
                            panic("non-exhaustive match")
                        }
                    }
                case 1:
                    var x3 string = mtmp0._p0
                    var t64 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                        _p1: x3,
                        _tag: 1,
                    }
                    return t64
                default:
                    panic("non-exhaustive match")
                }
            } else {
                break Loop_loop0
            }
        }
        var t9 string = _goml_m_std_p_json_p_json__error(value__0, "unterminated object")
        var t10 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
            _p1: t9,
            _tag: 1,
        }
        return t10
    }
}

func _goml_m_std_p_json_p_parse__json__value(value__0 _goml_m_std_p_json_p_JsonParser) _goml_m_Result____std_p_json_p_Value____string {
    _goml_m_std_p_json_p_skip__json__whitespace(value__0)
    var t0 *ref_int_x = value__0.index
    var t1 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t0)
    var t2 string = value__0.input
    var t3 int
    var inline30 int = _goml_runtime_core_string_len(t2)
    t3 = inline30
    var t4 bool = t1 >= t3
    if t4 {
        var t5 string
        var inline0 string = "expected JSON value"
        var inline1 string = "" + inline0
        var inline2 string = inline1 + " at byte "
        var inline3 *ref_int_x = value__0.index
        var inline4 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline3)
        var inline5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline4)
        var inline6 string = inline2 + inline5
        t5 = inline6
        var t6 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
            _p1: t5,
            _tag: 1,
        }
        return t6
    } else {
        var t7 string = value__0.input
        var t8 *ref_int_x = value__0.index
        var t9 int
        var inline29 int = ref_get__Ref_3int(t8)
        t9 = inline29
        var mtmp0 uint8
        var inline28 uint8 = _goml_runtime_core_string_byte_get(t7, t9)
        mtmp0 = inline28
        switch mtmp0 {
        case 123:
            var t10 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse__json__object(value__0)
            return t10
        case 91:
            var t11 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse__json__array(value__0)
            return t11
        case 34:
            var t12 Result__string__string = _goml_m_std_p_json_p_parse__json__string(value__0)
            var t13 closure_env_std_json_parse_json_value_53 = closure_env_std_json_parse_json_value_53{}
            var t14 func(string) _goml_m_std_p_json_p_Value = func(p0 string) _goml_m_std_p_json_p_Value {
                return _goml_m_inherent_i_closure__en_h8131e8c15e6a17b3a38874e07820df59_lue__53_i_apply(t13, p0)
            }
            switch t12._tag {
            case 0:
                var inline7 string = t12._p0
                var inline8 _goml_m_std_p_json_p_Value = t14(inline7)
                var inline9 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                    _p0: inline8,
                    _tag: 0,
                }
                return inline9
            case 1:
                var inline10 string = t12._p0
                var inline11 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                    _p1: inline10,
                    _tag: 1,
                }
                return inline11
            default:
                panic("non-exhaustive match")
            }
        case 116:
            var t15 _goml_m_std_p_json_p_Value = _goml_m_std_p_json_p_Value_Bool{
                _0: true,
            }
            var t16 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse__json__literal(value__0, "true", t15)
            return t16
        case 102:
            var t17 _goml_m_std_p_json_p_Value = _goml_m_std_p_json_p_Value_Bool{
                _0: false,
            }
            var t18 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse__json__literal(value__0, "false", t17)
            return t18
        case 110:
            var t19 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse__json__literal(value__0, "null", Null{})
            return t19
        default:
            var t20 bool = mtmp0 == 45
            var jp0 bool
            if t20 {
                jp0 = true
            } else {
                var inline26 bool = mtmp0 >= 48
                if inline26 {
                    var inline27 bool = mtmp0 <= 57
                    jp0 = inline27
                } else {
                    jp0 = false
                }
            }
            if jp0 {
                var inline12 Result__string__string = _goml_m_std_p_json_p_parse__json__number__text(value__0)
                var inline13 string
                switch inline12._tag {
                case 0:
                    var inline16 string = inline12._p0
                    inline13 = inline16
                    var inline14 _goml_m_std_p_json_p_Value = _goml_m_std_p_json_p_Value_Number{
                        _0: inline13,
                    }
                    var inline15 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                        _p0: inline14,
                        _tag: 0,
                    }
                    return inline15
                case 1:
                    var inline17 string = inline12._p0
                    var inline18 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                        _p1: inline17,
                        _tag: 1,
                    }
                    return inline18
                default:
                    panic("non-exhaustive match")
                }
            } else {
                var t21 string
                var inline19 string = "unexpected JSON token"
                var inline20 string = "" + inline19
                var inline21 string = inline20 + " at byte "
                var inline22 *ref_int_x = value__0.index
                var inline23 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline22)
                var inline24 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline23)
                var inline25 string = inline21 + inline24
                t21 = inline25
                var t22 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                    _p1: t21,
                    _tag: 1,
                }
                return t22
            }
        }
    }
}

func _goml_m_std_p_json_p_parse(input__0 string) _goml_m_Result____std_p_json_p_Value____string {
    var parser__0 _goml_m_std_p_json_p_JsonParser
    var inline9 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    var inline10 _goml_m_std_p_json_p_JsonParser = _goml_m_std_p_json_p_JsonParser{
        input: input__0,
        index: inline9,
    }
    parser__0 = inline10
    var mtmp0 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse__json__value(parser__0)
    var jp0 _goml_m_std_p_json_p_Value
    switch mtmp0._tag {
    case 0:
        var x0 _goml_m_std_p_json_p_Value = mtmp0._p0
        jp0 = x0
        _goml_m_std_p_json_p_skip__json__whitespace(parser__0)
        var t0 *ref_int_x = parser__0.index
        var t1 int
        var inline8 int = ref_get__Ref_3int(t0)
        t1 = inline8
        var t2 int
        var inline7 int = _goml_runtime_core_string_len(input__0)
        t2 = inline7
        var t3 bool = t1 == t2
        if t3 {
            var t4 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                _p0: jp0,
                _tag: 0,
            }
            return t4
        } else {
            var t5 string
            var inline0 string = "trailing JSON data"
            var inline1 string = "" + inline0
            var inline2 string = inline1 + " at byte "
            var inline3 *ref_int_x = parser__0.index
            var inline4 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline3)
            var inline5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline4)
            var inline6 string = inline2 + inline5
            t5 = inline6
            var t6 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
                _p1: t5,
                _tag: 1,
            }
            return t6
        }
    case 1:
        var x1 string = mtmp0._p1
        var t7 _goml_m_Result____std_p_json_p_Value____string = _goml_m_Result____std_p_json_p_Value____string{
            _p1: x1,
            _tag: 1,
        }
        return t7
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_json_p_write__json__string(builder__0 _goml_m_std_p_text_p_StringBuilder, value__0 string) struct{} {
    var inline17 rune = 34
    var inline18 string = _goml_m_inherent_i_char_i_char_i_to__string(inline17)
    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline18)
    var start__0 int = 0
    var for_index0 int = 0
    var for_limit0 int
    var inline16 int = _goml_runtime_core_string_len(value__0)
    for_limit0 = inline16
    Loop_loop0:
    for {
        var t5 bool = for_index0 < for_limit0
        if t5 {
            var for_item0 int = for_index0
            var t6 int = for_index0 + 1
            for_index0 = t6
            var byte__0 uint8
            var inline15 uint8 = _goml_runtime_core_string_byte_get(value__0, for_item0)
            byte__0 = inline15
            var t7 bool = byte__0 == 34
            var jp0 bool
            if t7 {
                jp0 = true
            } else {
                var t36 bool = byte__0 == 92
                jp0 = t36
            }
            var jp1 bool
            if jp0 {
                jp1 = true
            } else {
                var t35 bool = byte__0 == 8
                jp1 = t35
            }
            var jp2 bool
            if jp1 {
                jp2 = true
            } else {
                var t34 bool = byte__0 == 9
                jp2 = t34
            }
            var jp3 bool
            if jp2 {
                jp3 = true
            } else {
                var t33 bool = byte__0 == 10
                jp3 = t33
            }
            var jp4 bool
            if jp3 {
                jp4 = true
            } else {
                var t32 bool = byte__0 == 12
                jp4 = t32
            }
            var jp5 bool
            if jp4 {
                jp5 = true
            } else {
                var t31 bool = byte__0 == 13
                jp5 = t31
            }
            var jp6 bool
            if jp5 {
                jp6 = true
            } else {
                var t30 bool = byte__0 < 32
                jp6 = t30
            }
            if jp6 {
                var t8 bool = start__0 < for_item0
                if t8 {
                    var t28 string
                    var inline14 string = string_byte_slice(value__0, start__0, for_item0)
                    t28 = inline14
                    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, t28)
                } else {}
                var t9 bool = byte__0 == 34
                if t9 {
                    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "\\\"")
                } else {
                    var t12 bool = byte__0 == 92
                    if t12 {
                        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "\\\\")
                    } else {
                        var t14 bool = byte__0 == 8
                        if t14 {
                            _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "\\b")
                        } else {
                            var t16 bool = byte__0 == 9
                            if t16 {
                                _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "\\t")
                            } else {
                                var t18 bool = byte__0 == 10
                                if t18 {
                                    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "\\n")
                                } else {
                                    var t20 bool = byte__0 == 12
                                    if t20 {
                                        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "\\f")
                                    } else {
                                        var t22 bool = byte__0 == 13
                                        if t22 {
                                            _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "\\r")
                                        } else {
                                            _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "\\u00")
                                            var t24 uint8 = byte__0 / 16
                                            var t25 rune
                                            var inline12 int = int(uint8(t24))
                                            var inline13 rune = _goml_m_inherent_i_string_i_string_i_get("0123456789abcdef", inline12)
                                            t25 = inline13
                                            var inline10 string = _goml_m_inherent_i_char_i_char_i_to__string(t25)
                                            _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline10)
                                            var t26 uint8 = byte__0 % 16
                                            var t27 rune
                                            var inline8 int = int(uint8(t26))
                                            var inline9 rune = _goml_m_inherent_i_string_i_string_i_get("0123456789abcdef", inline8)
                                            t27 = inline9
                                            var inline6 string = _goml_m_inherent_i_char_i_char_i_to__string(t27)
                                            _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline6)
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
                var t10 int = for_item0 + 1
                start__0 = t10
                continue
            } else {
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline5 int = _goml_runtime_core_string_len(value__0)
    t0 = inline5
    var t1 bool = start__0 < t0
    if t1 {
        var t2 int
        var inline4 int = _goml_runtime_core_string_len(value__0)
        t2 = inline4
        var t3 string
        var inline3 string = string_byte_slice(value__0, start__0, t2)
        t3 = inline3
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, t3)
    } else {}
    var inline0 rune = 34
    var inline1 string = _goml_m_inherent_i_char_i_char_i_to__string(inline0)
    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline1)
    return struct{}{}
}

func _goml_m_std_p_json_p_write__json__value(builder__0 _goml_m_std_p_text_p_StringBuilder, value__0 _goml_m_std_p_json_p_Value) struct{} {
    switch value__0.(type) {
    case Object:
        var x0 *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value = value__0.(Object)._0
        var inline9 rune = 123
        var inline10 string = _goml_m_inherent_i_char_i_char_i_to__string(inline9)
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline10)
        var index__0 int = 0
        var for_limit0 int = vec_len___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value(x0)
        var for_index0 int = 0
        Loop_loop0:
        for {
            var t0 bool = for_index0 < for_limit0
            if t0 {
                var for_item0 Tuple2_6string_26_goml_m_std_p_json_p_Value = vec_get___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value(x0, for_index0)
                var t1 int = for_index0 + 1
                for_index0 = t1
                var t2 bool = index__0 > 0
                if t2 {
                    var inline6 rune = 44
                    var inline7 string = _goml_m_inherent_i_char_i_char_i_to__string(inline6)
                    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline7)
                } else {}
                var t3 string = for_item0._0
                _goml_m_std_p_json_p_write__json__string(builder__0, t3)
                var inline3 rune = 58
                var inline4 string = _goml_m_inherent_i_char_i_char_i_to__string(inline3)
                _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline4)
                var t4 _goml_m_std_p_json_p_Value = for_item0._1
                _goml_m_std_p_json_p_write__json__value(builder__0, t4)
                var compound_old0 int = index__0
                var compound_value0 int = 1
                var t5 int = compound_old0 + compound_value0
                index__0 = t5
                continue
            } else {
                break Loop_loop0
            }
        }
        var inline0 rune = 125
        var inline1 string = _goml_m_inherent_i_char_i_char_i_to__string(inline0)
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline1)
        return struct{}{}
    case _goml_m_std_p_json_p_Value_Array:
        var x1 *_goml_vec__goml_m_std_p_json_p_Value = value__0.(_goml_m_std_p_json_p_Value_Array)._0
        var inline18 rune = 91
        var inline19 string = _goml_m_inherent_i_char_i_char_i_to__string(inline18)
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline19)
        var index__1 int = 0
        var for_limit1 int = vec_len___goml_m_Vec__16std_p_json_p_Value(x1)
        var for_index1 int = 0
        Loop_loop1:
        for {
            var t7 bool = for_index1 < for_limit1
            if t7 {
                var for_item1 _goml_m_std_p_json_p_Value = vec_get___goml_m_Vec__16std_p_json_p_Value(x1, for_index1)
                var t8 int = for_index1 + 1
                for_index1 = t8
                var t9 bool = index__1 > 0
                if t9 {
                    var inline15 rune = 44
                    var inline16 string = _goml_m_inherent_i_char_i_char_i_to__string(inline15)
                    _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline16)
                } else {}
                _goml_m_std_p_json_p_write__json__value(builder__0, for_item1)
                var compound_old1 int = index__1
                var compound_value1 int = 1
                var t10 int = compound_old1 + compound_value1
                index__1 = t10
                continue
            } else {
                break Loop_loop1
            }
        }
        var inline12 rune = 93
        var inline13 string = _goml_m_inherent_i_char_i_char_i_to__string(inline12)
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, inline13)
        return struct{}{}
    case _goml_m_std_p_json_p_Value_String:
        var x2 string = value__0.(_goml_m_std_p_json_p_Value_String)._0
        _goml_m_std_p_json_p_write__json__string(builder__0, x2)
        return struct{}{}
    case _goml_m_std_p_json_p_Value_Number:
        var x3 string = value__0.(_goml_m_std_p_json_p_Value_Number)._0
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, x3)
        return struct{}{}
    case _goml_m_std_p_json_p_Value_Bool:
        var x4 bool = value__0.(_goml_m_std_p_json_p_Value_Bool)._0
        var jp0 string
        if x4 {
            jp0 = "true"
        } else {
            jp0 = "false"
        }
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, jp0)
        return struct{}{}
    case Null:
        _goml_m_inherent_i_std_p_text__h0034629766b91c65ed1f7160ea470eda_i_write__string(builder__0, "null")
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_json_p_encode(value__0 _goml_m_std_p_json_p_Value) string {
    var builder__0 _goml_m_std_p_text_p_StringBuilder
    var inline3 *_goml_vec_uint8 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__u8()
    var inline4 _goml_m_std_p_text_p_StringBuilder = _goml_m_std_p_text_p_StringBuilder{
        values: inline3,
    }
    builder__0 = inline4
    _goml_m_std_p_json_p_write__json__value(builder__0, value__0)
    var inline0 *_goml_vec_uint8 = builder__0.values
    var inline1 Tuple2_4bool_6string = string_from_utf8(inline0)
    var inline2 string = inline1._1
    return inline2
}

func _goml_m_std_p_json_p_field(value__0 _goml_m_std_p_json_p_Value, name__0 string) _goml_m_Option____std_p_json_p_Value {
    switch value__0.(type) {
    case Object:
        var x0 *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value = value__0.(Object)._0
        var for_limit0 int = vec_len___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value(x0)
        var for_index0 int = 0
        Loop_loop0:
        for {
            var t0 bool = for_index0 < for_limit0
            if t0 {
                var for_item0 Tuple2_6string_26_goml_m_std_p_json_p_Value = vec_get___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value(x0, for_index0)
                var t1 int = for_index0 + 1
                for_index0 = t1
                var t2 string = for_item0._0
                var t3 bool = t2 == name__0
                if t3 {
                    var t4 _goml_m_std_p_json_p_Value = for_item0._1
                    var t5 _goml_m_Option____std_p_json_p_Value = _goml_m_Option____std_p_json_p_Value{
                        _p0: t4,
                        _tag: 1,
                    }
                    return t5
                } else {
                    continue
                }
            } else {
                break Loop_loop0
            }
        }
        return _goml_m_Option____std_p_json_p_Value{
            _tag: 0,
        }
    default:
        return _goml_m_Option____std_p_json_p_Value{
            _tag: 0,
        }
    }
}

func _goml_m_std_p_json_p_parse__json__int__text(value__0 string) Option__isize {
    var t0 int
    var inline5 int = _goml_runtime_core_string_len(value__0)
    t0 = inline5
    var t1 bool = t0 == 0
    if t1 {
        return Option__isize{
            _tag: 0,
        }
    } else {
        var t2 uint8
        var inline3 int = 0
        var inline4 uint8 = _goml_runtime_core_string_byte_get(value__0, inline3)
        t2 = inline4
        var negative__0 bool = t2 == 45
        var jp0 int
        if negative__0 {
            jp0 = 1
        } else {
            jp0 = 0
        }
        var index__0 int = jp0
        var result__0 int = 0
        var t3 int
        var inline2 int = _goml_runtime_core_string_len(value__0)
        t3 = inline2
        var t4 bool = index__0 == t3
        if t4 {
            return Option__isize{
                _tag: 0,
            }
        } else {
            Loop_loop0:
            for {
                var t7 int
                var inline1 int = _goml_runtime_core_string_len(value__0)
                t7 = inline1
                var t8 bool = index__0 < t7
                if t8 {
                    var byte__0 uint8
                    var inline0 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                    byte__0 = inline0
                    var t9 bool = byte__0 < 48
                    var jp2 bool
                    if t9 {
                        jp2 = true
                    } else {
                        var t16 bool = byte__0 > 57
                        jp2 = t16
                    }
                    if jp2 {
                        return Option__isize{
                            _tag: 0,
                        }
                    } else {
                        var t12 int = result__0 * 10
                        var t13 uint8 = byte__0 - 48
                        var t14 int = int(uint8(t13))
                        var t15 int = t12 + t14
                        result__0 = t15
                        var compound_old0 int = index__0
                        var compound_value0 int = 1
                        var t10 int = compound_old0 + compound_value0
                        index__0 = t10
                        continue
                    }
                } else {
                    break Loop_loop0
                }
            }
            var jp1 int
            if negative__0 {
                var t6 int = 0 - result__0
                jp1 = t6
            } else {
                jp1 = result__0
            }
            var t5 Option__isize = Option__isize{
                _p0: jp1,
                _tag: 1,
            }
            return t5
        }
    }
}

func main0() struct{} {
    var mtmp0 _goml_m_Result____std_p_json_p_Value____string = _goml_m_std_p_json_p_parse("{\"name\":\"goml\",\"version\":1,\"stable\":true}")
    var jp0 _goml_m_std_p_json_p_Value
    switch mtmp0._tag {
    case 0:
        var x4 _goml_m_std_p_json_p_Value = mtmp0._p0
        jp0 = x4
        var mtmp1 _goml_m_Option____std_p_json_p_Value = _goml_m_std_p_json_p_field(jp0, "name")
        switch mtmp1._tag {
        case 0:
            var inline19 string = "missing name"
            var inline20 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline19)
            _goml_runtime_core_string_println(inline20)
            var mtmp2 _goml_m_Option____std_p_json_p_Value = _goml_m_std_p_json_p_field(jp0, "version")
            switch mtmp2._tag {
            case 0:
                var inline9 string = "missing version"
                var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline9)
                _goml_runtime_core_string_println(inline10)
            case 1:
                var x1 _goml_m_std_p_json_p_Value = mtmp2._p0
                var mtmp4 Option__isize
                switch x1.(type) {
                case _goml_m_std_p_json_p_Value_Number:
                    var inline17 string = x1.(_goml_m_std_p_json_p_Value_Number)._0
                    var inline18 Option__isize = _goml_m_std_p_json_p_parse__json__int__text(inline17)
                    mtmp4 = inline18
                default:
                    mtmp4 = Option__isize{
                        _tag: 0,
                    }
                }
                switch mtmp4._tag {
                case 0:
                    var inline12 string = "invalid version"
                    var inline13 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline12)
                    _goml_runtime_core_string_println(inline13)
                case 1:
                    var x2 int = mtmp4._p0
                    var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x2)
                    _goml_runtime_core_string_println(inline15)
                default:
                    panic("non-exhaustive match")
                }
            default:
                panic("non-exhaustive match")
            }
            var mtmp3 _goml_m_Option____std_p_json_p_Value = _goml_m_std_p_json_p_field(jp0, "stable")
            switch mtmp3._tag {
            case 0:
                var inline0 string = "missing stable"
                var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
                _goml_runtime_core_string_println(inline1)
                var t0 string = _goml_m_std_p_json_p_encode(jp0)
                println__T_string(t0)
                return struct{}{}
            case 1:
                var x0 _goml_m_std_p_json_p_Value = mtmp3._p0
                var commute_field0 bool
                switch x0.(type) {
                case _goml_m_std_p_json_p_Value_Bool:
                    var inline8 bool = x0.(_goml_m_std_p_json_p_Value_Bool)._0
                    commute_field0 = inline8
                    var inline6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(commute_field0)
                    _goml_runtime_core_string_println(inline6)
                    var t0 string = _goml_m_std_p_json_p_encode(jp0)
                    println__T_string(t0)
                    return struct{}{}
                default:
                    var inline3 string = "invalid stable"
                    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline3)
                    _goml_runtime_core_string_println(inline4)
                    var t0 string = _goml_m_std_p_json_p_encode(jp0)
                    println__T_string(t0)
                    return struct{}{}
                }
            default:
                panic("non-exhaustive match")
            }
        case 1:
            var x3 _goml_m_std_p_json_p_Value = mtmp1._p0
            var commute_field1 string
            switch x3.(type) {
            case _goml_m_std_p_json_p_Value_String:
                var inline27 string = x3.(_goml_m_std_p_json_p_Value_String)._0
                commute_field1 = inline27
                var inline25 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(commute_field1)
                _goml_runtime_core_string_println(inline25)
                var mtmp2 _goml_m_Option____std_p_json_p_Value = _goml_m_std_p_json_p_field(jp0, "version")
                switch mtmp2._tag {
                case 0:
                    var inline9 string = "missing version"
                    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline9)
                    _goml_runtime_core_string_println(inline10)
                case 1:
                    var x1 _goml_m_std_p_json_p_Value = mtmp2._p0
                    var mtmp4 Option__isize
                    switch x1.(type) {
                    case _goml_m_std_p_json_p_Value_Number:
                        var inline17 string = x1.(_goml_m_std_p_json_p_Value_Number)._0
                        var inline18 Option__isize = _goml_m_std_p_json_p_parse__json__int__text(inline17)
                        mtmp4 = inline18
                    default:
                        mtmp4 = Option__isize{
                            _tag: 0,
                        }
                    }
                    switch mtmp4._tag {
                    case 0:
                        var inline12 string = "invalid version"
                        var inline13 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline12)
                        _goml_runtime_core_string_println(inline13)
                    case 1:
                        var x2 int = mtmp4._p0
                        var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x2)
                        _goml_runtime_core_string_println(inline15)
                    default:
                        panic("non-exhaustive match")
                    }
                default:
                    panic("non-exhaustive match")
                }
                var mtmp3 _goml_m_Option____std_p_json_p_Value = _goml_m_std_p_json_p_field(jp0, "stable")
                switch mtmp3._tag {
                case 0:
                    var inline0 string = "missing stable"
                    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
                    _goml_runtime_core_string_println(inline1)
                    var t0 string = _goml_m_std_p_json_p_encode(jp0)
                    println__T_string(t0)
                    return struct{}{}
                case 1:
                    var x0 _goml_m_std_p_json_p_Value = mtmp3._p0
                    var commute_field0 bool
                    switch x0.(type) {
                    case _goml_m_std_p_json_p_Value_Bool:
                        var inline8 bool = x0.(_goml_m_std_p_json_p_Value_Bool)._0
                        commute_field0 = inline8
                        var inline6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(commute_field0)
                        _goml_runtime_core_string_println(inline6)
                        var t0 string = _goml_m_std_p_json_p_encode(jp0)
                        println__T_string(t0)
                        return struct{}{}
                    default:
                        var inline3 string = "invalid stable"
                        var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline3)
                        _goml_runtime_core_string_println(inline4)
                        var t0 string = _goml_m_std_p_json_p_encode(jp0)
                        println__T_string(t0)
                        return struct{}{}
                    }
                default:
                    panic("non-exhaustive match")
                }
            default:
                var inline22 string = "invalid name"
                var inline23 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline22)
                _goml_runtime_core_string_println(inline23)
                var mtmp2 _goml_m_Option____std_p_json_p_Value = _goml_m_std_p_json_p_field(jp0, "version")
                switch mtmp2._tag {
                case 0:
                    var inline9 string = "missing version"
                    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline9)
                    _goml_runtime_core_string_println(inline10)
                case 1:
                    var x1 _goml_m_std_p_json_p_Value = mtmp2._p0
                    var mtmp4 Option__isize
                    switch x1.(type) {
                    case _goml_m_std_p_json_p_Value_Number:
                        var inline17 string = x1.(_goml_m_std_p_json_p_Value_Number)._0
                        var inline18 Option__isize = _goml_m_std_p_json_p_parse__json__int__text(inline17)
                        mtmp4 = inline18
                    default:
                        mtmp4 = Option__isize{
                            _tag: 0,
                        }
                    }
                    switch mtmp4._tag {
                    case 0:
                        var inline12 string = "invalid version"
                        var inline13 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline12)
                        _goml_runtime_core_string_println(inline13)
                    case 1:
                        var x2 int = mtmp4._p0
                        var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x2)
                        _goml_runtime_core_string_println(inline15)
                    default:
                        panic("non-exhaustive match")
                    }
                default:
                    panic("non-exhaustive match")
                }
                var mtmp3 _goml_m_Option____std_p_json_p_Value = _goml_m_std_p_json_p_field(jp0, "stable")
                switch mtmp3._tag {
                case 0:
                    var inline0 string = "missing stable"
                    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
                    _goml_runtime_core_string_println(inline1)
                    var t0 string = _goml_m_std_p_json_p_encode(jp0)
                    println__T_string(t0)
                    return struct{}{}
                case 1:
                    var x0 _goml_m_std_p_json_p_Value = mtmp3._p0
                    var commute_field0 bool
                    switch x0.(type) {
                    case _goml_m_std_p_json_p_Value_Bool:
                        var inline8 bool = x0.(_goml_m_std_p_json_p_Value_Bool)._0
                        commute_field0 = inline8
                        var inline6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(commute_field0)
                        _goml_runtime_core_string_println(inline6)
                        var t0 string = _goml_m_std_p_json_p_encode(jp0)
                        println__T_string(t0)
                        return struct{}{}
                    default:
                        var inline3 string = "invalid stable"
                        var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline3)
                        _goml_runtime_core_string_println(inline4)
                        var t0 string = _goml_m_std_p_json_p_encode(jp0)
                        println__T_string(t0)
                        return struct{}{}
                    }
                default:
                    panic("non-exhaustive match")
                }
            }
        default:
            panic("non-exhaustive match")
        }
    case 1:
        var x5 string = mtmp0._p1
        var inline28 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x5)
        _goml_runtime_core_string_println(inline28)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_char_i_char_i_to__string(self__0 rune) string {
    var inline0 uint32 = uint32(rune(self__0))
    var inline1 bool = utf8_valid_scalar(inline0)
    if inline1 {
        var inline2 string = _goml_runtime_core_char_to_string(self__0)
        return inline2
    } else {
        _goml_runtime_core_string_get("", -1)
        return ""
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__u8() *_goml_vec_uint8 {
    var t0 *_goml_vec_uint8 = vec_new__Vec_5uint8()
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__get(self__0 string, index__0 int) uint8 {
    var t0 uint8 = _goml_runtime_core_string_byte_get(self__0, index__0)
    return t0
}

func string_from_utf8(bytes__0 *_goml_vec_uint8) Tuple2_4bool_6string {
    var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
    var x0 string = mtmp0._1
    var index__0 int = 0
    Loop_loop0:
    for {
        var t1 int
        var inline0 int = _goml_runtime_core_string_len(x0)
        t1 = inline0
        var t2 bool = index__0 < t1
        if t2 {
            var mtmp1 Tuple3_4bool_4char_3int = string_decode_utf8_at(x0, index__0)
            var x1 bool = mtmp1._0
            var x2 int = mtmp1._2
            if x1 {
                var compound_old0 int = index__0
                var t3 int = compound_old0 + x2
                index__0 = t3
                continue
            } else {
                var t5 Tuple2_4bool_6string = Tuple2_4bool_6string{
                    _0: false,
                    _1: "",
                }
                return t5
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 Tuple2_4bool_6string = Tuple2_4bool_6string{
        _0: true,
        _1: x0,
    }
    return t0
}

func _goml_m_inherent_i_isize_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func char_from_u32(value__0 uint32) Option__char {
    var inline0 bool = utf8_valid_scalar(value__0)
    if inline0 {
        var inline1 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var inline2 rune = inline1._1
        var inline3 Option__char = Option__char(uint64(uint32(inline2)) + 1)
        return inline3
    } else {
        return Option__char(0)
    }
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(value__0 int) *ref_int_x {
    var t0 *ref_int_x = ref__Ref_3int(value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(self__0 *ref_int_x) int {
    var t0 int = ref_get__Ref_3int(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(self__0 *ref_int_x, value__0 int) struct{} {
    ref_set__Ref_3int(self__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_string_i_string_i_byte__slice(self__0 string, start__0 int, end__0 int) string {
    var inline0 bool = string_is_char_boundary(self__0, start__0)
    var inline1 bool
    if inline0 {
        var inline4 bool = string_is_char_boundary(self__0, end__0)
        inline1 = inline4
    } else {
        inline1 = false
    }
    if inline1 {
        var inline2 string = _goml_runtime_core_string_byte_slice(self__0, start__0, end__0)
        return inline2
    } else {
        var inline3 string = _goml_runtime_core_string_byte_slice(self__0, -1, -1)
        return inline3
    }
}

func _goml_m_inherent_i_string_i_string_i_get(self__0 string, index__0 int) rune {
    var inline0 Tuple3_4bool_4char_3int = string_decode_utf8_at(self__0, index__0)
    var inline1 bool = inline0._0
    var inline2 rune = inline0._1
    if inline1 {
        return inline2
    } else {
        var inline3 rune = _goml_runtime_core_string_get("", -1)
        return inline3
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__std_p_json_p_Value() *_goml_vec__goml_m_std_p_json_p_Value {
    var t0 *_goml_vec__goml_m_std_p_json_p_Value = vec_new___goml_m_Vec__16std_p_json_p_Value()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_h57e42dbef834e6b8ee6cf77cf9eb9d23_json_p_Value_q_() *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value {
    var t0 *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value = vec_new___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_h153048c8bd06f0dfabad32cddaecb150_json_p_Value_q_(self__0 *_goml_vec__goml_m_Tuple2__6string__16std_p_json_p_Value, elem__0 Tuple2_6string_26_goml_m_std_p_json_p_Value) struct{} {
    vec_push___goml_m_Vec__33Tuple2__6string__16std_p_json_p_Value(self__0, elem__0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func char_to_string(value__0 rune) string {
    var t0 uint32 = uint32(rune(value__0))
    var t1 bool
    var inline0 bool = t0 <= 1114111
    if inline0 {
        var inline1 bool = t0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = t0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t1 = inline3
    } else {
        t1 = false
    }
    if t1 {
        var t2 string = _goml_runtime_core_char_to_string(value__0)
        return t2
    } else {
        _goml_runtime_core_string_get("", -1)
        return ""
    }
}

func string_decode_utf8_at(value__0 string, index__0 int) Tuple3_4bool_4char_3int {
    var length__0 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t63 bool = index__0 >= length__0
        jp0 = t63
    }
    if jp0 {
        var inline25 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline25
    } else {
        var t1 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
        var first__0 uint32 = uint32(uint8(t1))
        var t2 bool = first__0 < 128
        if t2 {
            var inline0 int = 1
            var inline1 Option__char = __goml_builtin_char_from_uint32(first__0)
            switch inline1 != Option__char(0) {
            case false:
                var inline2 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                return inline2
            case true:
                var inline3 rune = rune(uint64(inline1) - 1)
                var inline4 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: true,
                    _1: inline3,
                    _2: inline0,
                }
                return inline4
            default:
                panic("non-exhaustive match")
            }
        } else {
            var t3 bool = first__0 < 194
            if t3 {
                var inline5 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: false,
                    _1: 0,
                    _2: 0,
                }
                return inline5
            } else {
                var t4 bool = first__0 < 224
                if t4 {
                    var t5 int = length__0 - index__0
                    var t6 bool = t5 < 2
                    if t6 {
                        var inline15 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                            _0: false,
                            _1: 0,
                            _2: 0,
                        }
                        return inline15
                    } else {
                        var t7 int = index__0 + 1
                        var t8 uint8
                        var inline14 uint8 = _goml_runtime_core_string_byte_get(value__0, t7)
                        t8 = inline14
                        var second__0 uint32 = uint32(uint8(t8))
                        var t9 bool
                        var inline12 bool = second__0 < 128
                        if inline12 {
                            t9 = true
                        } else {
                            var inline13 bool = second__0 > 191
                            t9 = inline13
                        }
                        if t9 {
                            var inline6 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline6
                        } else {
                            var t10 uint32 = first__0 & 31
                            var t11 uint32 = t10 << 6
                            var t12 uint32 = second__0 & 63
                            var t13 uint32 = t11 | t12
                            var inline7 int = 2
                            var inline8 Option__char = __goml_builtin_char_from_uint32(t13)
                            switch inline8 != Option__char(0) {
                            case false:
                                var inline9 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return inline9
                            case true:
                                var inline10 rune = rune(uint64(inline8) - 1)
                                var inline11 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: true,
                                    _1: inline10,
                                    _2: inline7,
                                }
                                return inline11
                            default:
                                panic("non-exhaustive match")
                            }
                        }
                    }
                } else {
                    var t14 bool = first__0 < 240
                    if t14 {
                        var t15 int = length__0 - index__0
                        var t16 bool = t15 < 3
                        if t16 {
                            var inline24 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline24
                        } else {
                            var t17 int = index__0 + 1
                            var t18 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t17)
                            var second__1 uint32 = uint32(uint8(t18))
                            var t19 int = index__0 + 2
                            var t20 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t19)
                            var third__0 uint32 = uint32(uint8(t20))
                            var t21 bool = utf8_invalid_continuation(second__1)
                            var jp1 bool
                            if t21 {
                                jp1 = true
                            } else {
                                var inline22 bool = third__0 < 128
                                if inline22 {
                                    jp1 = true
                                } else {
                                    var inline23 bool = third__0 > 191
                                    jp1 = inline23
                                }
                            }
                            var jp2 bool
                            if jp1 {
                                jp2 = true
                            } else {
                                var t31 bool = first__0 == 224
                                if t31 {
                                    var t32 bool = second__1 < 160
                                    jp2 = t32
                                } else {
                                    jp2 = false
                                }
                            }
                            var jp3 bool
                            if jp2 {
                                jp3 = true
                            } else {
                                var t29 bool = first__0 == 237
                                if t29 {
                                    var t30 bool = second__1 >= 160
                                    jp3 = t30
                                } else {
                                    jp3 = false
                                }
                            }
                            if jp3 {
                                var inline16 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: false,
                                    _1: 0,
                                    _2: 0,
                                }
                                return inline16
                            } else {
                                var t22 uint32 = first__0 & 15
                                var t23 uint32 = t22 << 12
                                var t24 uint32 = second__1 & 63
                                var t25 uint32 = t24 << 6
                                var t26 uint32 = t23 | t25
                                var t27 uint32 = third__0 & 63
                                var t28 uint32 = t26 | t27
                                var inline17 int = 3
                                var inline18 Option__char = __goml_builtin_char_from_uint32(t28)
                                switch inline18 != Option__char(0) {
                                case false:
                                    var inline19 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return inline19
                                case true:
                                    var inline20 rune = rune(uint64(inline18) - 1)
                                    var inline21 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                        _0: true,
                                        _1: inline20,
                                        _2: inline17,
                                    }
                                    return inline21
                                default:
                                    panic("non-exhaustive match")
                                }
                            }
                        }
                    } else {
                        var t33 bool = first__0 < 245
                        if t33 {
                            var t34 int = length__0 - index__0
                            var t35 bool = t34 < 4
                            if t35 {
                                var t61 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return t61
                            } else {
                                var t36 int = index__0 + 1
                                var t37 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t36)
                                var second__2 uint32 = uint32(uint8(t37))
                                var t38 int = index__0 + 2
                                var t39 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t38)
                                var third__1 uint32 = uint32(uint8(t39))
                                var t40 int = index__0 + 3
                                var t41 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t40)
                                var fourth__0 uint32 = uint32(uint8(t41))
                                var t42 bool = utf8_invalid_continuation(second__2)
                                var jp4 bool
                                if t42 {
                                    jp4 = true
                                } else {
                                    var t60 bool = utf8_invalid_continuation(third__1)
                                    jp4 = t60
                                }
                                var jp5 bool
                                if jp4 {
                                    jp5 = true
                                } else {
                                    var t59 bool = utf8_invalid_continuation(fourth__0)
                                    jp5 = t59
                                }
                                var jp6 bool
                                if jp5 {
                                    jp6 = true
                                } else {
                                    var t57 bool = first__0 == 240
                                    if t57 {
                                        var t58 bool = second__2 < 144
                                        jp6 = t58
                                    } else {
                                        jp6 = false
                                    }
                                }
                                var jp7 bool
                                if jp6 {
                                    jp7 = true
                                } else {
                                    var t55 bool = first__0 == 244
                                    if t55 {
                                        var t56 bool = second__2 > 143
                                        jp7 = t56
                                    } else {
                                        jp7 = false
                                    }
                                }
                                if jp7 {
                                    var t43 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return t43
                                } else {
                                    var t44 uint32 = first__0 & 7
                                    var t45 uint32 = t44 << 18
                                    var t46 uint32 = second__2 & 63
                                    var t47 uint32 = t46 << 12
                                    var t48 uint32 = t45 | t47
                                    var t49 uint32 = third__1 & 63
                                    var t50 uint32 = t49 << 6
                                    var t51 uint32 = t48 | t50
                                    var t52 uint32 = fourth__0 & 63
                                    var t53 uint32 = t51 | t52
                                    var t54 Tuple3_4bool_4char_3int = utf8_valid_decode(t53, 4)
                                    return t54
                                }
                            }
                        } else {
                            var t62 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                            return t62
                        }
                    }
                }
            }
        }
    }
}

func __goml_builtin_int_to_string(value__0 int) string {
    var t0 int64 = int64(int(value__0))
    var inline0 bool = t0 < 0
    if inline0 {
        var inline1 uint64 = uint64(int64(t0))
        var inline2 uint64 = 0 - inline1
        var inline3 string = decimal_string(inline2)
        var inline4 string = "-" + inline3
        return inline4
    } else {
        var inline5 uint64 = uint64(int64(t0))
        var inline6 string = decimal_string(inline5)
        return inline6
    }
}

func __goml_builtin_char_from_uint32(value__0 uint32) Option__char {
    var t0 bool
    var inline0 bool = value__0 <= 1114111
    if inline0 {
        var inline1 bool = value__0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = value__0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t0 = inline3
    } else {
        t0 = false
    }
    if t0 {
        var mtmp0 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var x0 rune = mtmp0._1
        var t1 Option__char = Option__char(uint64(uint32(x0)) + 1)
        return t1
    } else {
        return Option__char(0)
    }
}

func string_byte_slice(value__0 string, start__0 int, end__0 int) string {
    var t0 bool = string_is_char_boundary(value__0, start__0)
    var jp0 bool
    if t0 {
        var t3 bool = string_is_char_boundary(value__0, end__0)
        jp0 = t3
    } else {
        jp0 = false
    }
    if jp0 {
        var t1 string = _goml_runtime_core_string_byte_slice(value__0, start__0, end__0)
        return t1
    } else {
        var t2 string = _goml_runtime_core_string_byte_slice(value__0, -1, -1)
        return t2
    }
}

func string_is_char_boundary(value__0 string, index__0 int) bool {
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t6 int
        var inline2 int = _goml_runtime_core_string_len(value__0)
        t6 = inline2
        var t7 bool = index__0 > t6
        jp0 = t7
    }
    if jp0 {
        return false
    } else {
        var t1 int
        var inline1 int = _goml_runtime_core_string_len(value__0)
        t1 = inline1
        var t2 bool = index__0 == t1
        if t2 {
            return true
        } else {
            var t3 uint8
            var inline0 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
            t3 = inline0
            var t4 uint8 = t3 & 192
            var t5 bool = t4 != 128
            return t5
        }
    }
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func utf8_valid_scalar(value__0 uint32) bool {
    var t0 bool = value__0 <= 1114111
    if t0 {
        var t1 bool = value__0 >= 55296
        var jp0 bool
        if t1 {
            var t3 bool = value__0 <= 57343
            jp0 = t3
        } else {
            jp0 = false
        }
        var t2 bool = !jp0
        return t2
    } else {
        return false
    }
}

func utf8_invalid_decode() Tuple3_4bool_4char_3int {
    var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
        _0: false,
        _1: 0,
        _2: 0,
    }
    return t0
}

func utf8_valid_decode(value__0 uint32, width__0 int) Tuple3_4bool_4char_3int {
    var commute_field0 rune
    var inline1 bool = utf8_valid_scalar(value__0)
    if inline1 {
        var inline2 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var inline3 rune = inline2._1
        commute_field0 = inline3
        var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: true,
            _1: commute_field0,
            _2: width__0,
        }
        return t0
    } else {
        var inline0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline0
    }
}

func utf8_invalid_continuation(value__0 uint32) bool {
    var t0 bool = value__0 < 128
    if t0 {
        return true
    } else {
        var t1 bool = value__0 > 191
        return t1
    }
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

func _goml_m_inherent_i_closure__en_h8131e8c15e6a17b3a38874e07820df59_lue__53_i_apply(env0 closure_env_std_json_parse_json_value_53, text__0 string) _goml_m_std_p_json_p_Value {
    var t0 _goml_m_std_p_json_p_Value = _goml_m_std_p_json_p_Value_String{
        _0: text__0,
    }
    return t0
}

func main() {
    main0()
}
