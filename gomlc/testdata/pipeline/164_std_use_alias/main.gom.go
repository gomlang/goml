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

type _goml_vec__goml_m_std_p_text_p_LineIndexWideChar struct {
    items []_goml_m_std_p_text_p_LineIndexWideChar
}

type _goml_vec__goml_m_std_p_text_p_ReplacementNode struct {
    items []_goml_m_std_p_text_p_ReplacementNode
}

type _goml_vec_uint32 struct {
    items []uint32
}

type ref_int_x struct {
    value int
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

type _goml_synthetic__goml_option__bool interface {
    is_goml_synthetic__goml_option__bool()
}

type _goml_synthetic__goml_option__bool_Some struct {
    _0 bool
}

func (_ _goml_synthetic__goml_option__bool_Some) is_goml_synthetic__goml_option__bool() {}

type _goml_synthetic__goml_option__bool_None struct {}

func (_ _goml_synthetic__goml_option__bool_None) is_goml_synthetic__goml_option__bool() {}

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

type Tuple2_10Vec_5uint8_3int struct {
    _0 *_goml_vec_uint8
    _1 int
}

type Tuple2_4bool_4char struct {
    _0 bool
    _1 rune
}

type Tuple2_13Option__isize_13Option__isize struct {
    _0 Option__isize
    _1 Option__isize
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

type closure_env_std_text_fields_by_iter_18 struct {
    offset_0 *ref_int_x
    value_1 string
    separator_2 func(rune) bool
}

type closure_env_std_text_fields_iter_19 struct {}

type closure_env_std_text_fields_20 struct {}

type closure_env_std_text_trim_space_21 struct {}

type closure_env_std_text_trim_chars_22 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_trim_start_chars_23 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_trim_end_chars_24 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_find_any_25 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_rfind_any_26 struct {
    set_0 *hashmap_char_bool_x
}

type closure_env_std_text_split_iterator_27 struct {
    done_0 *ref_bool_x
    offset_1 *ref_int_x
    separator_2 string
    value_3 string
    finder_4 _goml_m_std_p_text_p_Finder
    after_5 bool
}

type closure_env_std_text_lines_inclusive_iter_28 struct {
    parts_0 FnIterator__string
}

type closure_env_std_text_lines_iter_29 struct {
    parts_0 FnIterator__string
}

type closure_env_std_text_match_positions_30 struct {
    done_0 *ref_bool_x
    remaining_1 *ref_int_x
    finder_2 _goml_m_std_p_text_p_Finder
    offset_3 *ref_int_x
    value_4 string
}

type closure_env_std_text_case_checked_31 struct {
    kind_0 _goml_m_std_p_unicode_p_Case
}

type closure_env_std_text_case_with_checked_32 struct {
    special_0 _goml_m_std_p_unicode_p_SpecialCase
    kind_1 _goml_m_std_p_unicode_p_Case
}

type closure_env_std_text_unquote_33 struct {}

type closure_env_inherent_std_text_Replacer_std_text_Replacer_chunks_34 struct {
    done_0 *ref_bool_x
    cursor_1 _goml_m_std_p_text_p_ReplacementCursor
}

type closure_env_inherent_string_string_chars_35 struct {
    self_0 string
    index_1 *ref_int_x
}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Cursor_36 struct {}

type closure_env_std_io_trait_default_BufRead_discard_Self_std_io_Cursor_37 struct {
    discarded_0 *ref_int_x
}

type closure_env_std_io_trait_default_BufRead_discard_Self_std_io_Cursor_38 struct {
    discarded_0 *ref_int_x
}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_39 struct {}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_40 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Stdin_41 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_StringReader_42 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_PipeReader_43 struct {}

type closure_env_inherent_string_string_char_indices_44 struct {
    index_0 *ref_int_x
    self_1 string
}

type closure_env_goml_builtin_range_45 struct {
    current_0 *ref_int_x
    end_1 int
}

type closure_env_inherent_Slice_Slice_T_iter_T_u8_46 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 []uint8
}

type closure_env_inherent_Slice_Slice_T_iter_T_string_string_47 struct {
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
