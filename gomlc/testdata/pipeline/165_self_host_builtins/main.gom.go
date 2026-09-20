package main

import (
    _goml_context "context"
    _goml_errors "errors"
    _goml_io "io"
    _goml_os "os"
    _goml_sync "sync"
    _goml_syscall "syscall"
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

func _goml_runtime_error_kind(err error) int {
    if _goml_os.IsNotExist(err) {
        return 1
    }
    if _goml_os.IsPermission(err) {
        return 2
    }
    if _goml_os.IsExist(err) {
        return 3
    }
    if _goml_os.IsTimeout(err) {
        return 6
    }
    if _goml_errors.Is(err, _goml_syscall.EINVAL) {
        return 4
    }
    if _goml_errors.Is(err, _goml_syscall.ENAMETOOLONG) {
        return 4
    }
    if _goml_errors.Is(err, _goml_syscall.EINTR) {
        return 7
    }
    if _goml_errors.Is(err, _goml_io.ErrUnexpectedEOF) {
        return 8
    }
    if _goml_errors.Is(err, _goml_io.EOF) {
        return 8
    }
    if _goml_errors.Is(err, _goml_io.ErrShortWrite) {
        return 9
    }
    if _goml_errors.Is(err, _goml_syscall.EPIPE) {
        return 10
    }
    if _goml_errors.Is(err, _goml_syscall.EAGAIN) {
        return 11
    }
    return 0
}

func _goml_runtime_error_raw_code(err error) (bool, int) {
    switch host_error := err.(type) {
    case *_goml_os.PathError:
        switch errno := host_error.Err.(type) {
        case _goml_syscall.Errno:
            return true, int(errno)
        }
    case *_goml_os.LinkError:
        switch errno := host_error.Err.(type) {
        case _goml_syscall.Errno:
            return true, int(errno)
        }
    case *_goml_os.SyscallError:
        switch errno := host_error.Err.(type) {
        case _goml_syscall.Errno:
            return true, int(errno)
        }
    }
    return false, 0
}

func _goml_runtime_core_string_len(s string) int {
    return int(len(s))
}

func _goml_runtime_core_string_byte_get(s string, i int) uint8 {
    return s[i]
}

func _goml_runtime_core_string_byte_slice(s string, start int, end int) string {
    return s[start:end]
}

func _goml_runtime_core_string_to_bytes(s string) *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: []byte(s),
    }
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_char_from_uint32(value uint32) Tuple2_4bool_4char {
    return Tuple2_4bool_4char{
        _0: true,
        _1: rune(value),
    }
}

func _goml_runtime_std_fs_read_bytes_v2(path string) Tuple6_4bool_10Vec_5uint8_3int_4bool_3int_6string {
    var data []uint8
    var err error
    var has_raw_code bool
    var raw_code int
    data, err = _goml_os.ReadFile(path)
    if err != nil {
        has_raw_code, raw_code = _goml_runtime_error_raw_code(err)
        return Tuple6_4bool_10Vec_5uint8_3int_4bool_3int_6string{
            _0: false,
            _1: &_goml_vec_uint8{
                items: nil,
            },
            _2: _goml_runtime_error_kind(err),
            _3: has_raw_code,
            _4: raw_code,
            _5: err.Error(),
        }
    }
    return Tuple6_4bool_10Vec_5uint8_3int_4bool_3int_6string{
        _0: true,
        _1: &_goml_vec_uint8{
            items: data,
        },
        _2: 0,
        _3: false,
        _4: 0,
        _5: "",
    }
}

func _goml_runtime_std_fs_write_bytes_v2(path string, data *_goml_vec_uint8) Tuple5_4bool_3int_4bool_3int_6string {
    var has_raw_code bool
    var raw_code int
    var err error = _goml_os.WriteFile(path, data.items, 0644)
    if err != nil {
        has_raw_code, raw_code = _goml_runtime_error_raw_code(err)
        return Tuple5_4bool_3int_4bool_3int_6string{
            _0: false,
            _1: _goml_runtime_error_kind(err),
            _2: has_raw_code,
            _3: raw_code,
            _4: err.Error(),
        }
    }
    return Tuple5_4bool_3int_4bool_3int_6string{
        _0: true,
        _1: 0,
        _2: false,
        _3: 0,
        _4: "",
    }
}

func _goml_runtime_std_fs_create_dir_all(path string) Tuple2_4bool_6string {
    var err error = _goml_os.MkdirAll(path, 0755)
    if err != nil {
        return Tuple2_4bool_6string{
            _0: false,
            _1: err.Error(),
        }
    }
    return Tuple2_4bool_6string{
        _0: true,
        _1: "",
    }
}

func _goml_runtime_std_io_println(value string) struct{} {
    _goml_os.Stdout.WriteString(value + "\n")
    return struct{}{}
}

func _goml_runtime_std_io_eprint(value string) struct{} {
    _goml_os.Stderr.WriteString(value)
    return struct{}{}
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

type _goml_vec_string struct {
    items []string
}

type _goml_vec_Tuple2_6string_6string struct {
    items []Tuple2_6string_6string
}

type _goml_vec__goml_m_std_p_fs_p_DirEntry struct {
    items []_goml_m_std_p_fs_p_DirEntry
}

type _goml_vec_uint32 struct {
    items []uint32
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

type _goml_m_std_p_fs_p_Error struct {
    details _goml_m_std_p_io_p_ErrorDetails
}

type _goml_m_std_p_fs_p_Permissions struct {
    mode_value uint32
}

type _goml_m_std_p_fs_p_Metadata struct {
    file_type_value _goml_m_std_p_fs_p_FileType
    length_value int64
    permissions_value _goml_m_std_p_fs_p_Permissions
    modified_unix_nanoseconds_value int64
}

type _goml_m_std_p_fs_p_DirEntry struct {
    name_value string
    path_value string
    file_type_value _goml_m_std_p_fs_p_FileType
}

type FnIterator__u8 struct {
    next_fn func() Option__u8
}

type closure_env_std_io_read_stdin_to_string_0 struct {}

type closure_env_std_fs_read_file_structured_1 struct {
    path_0 string
}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Cursor_2 struct {}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_3 struct {}

type closure_env_std_io_trait_default_BufRead_read_line_Self_std_io_Cursor_4 struct {}

type closure_env_std_io_trait_default_Read_read_to_string_Self_std_io_Stdin_5 struct {}

type closure_env_inherent_Slice_Slice_T_iter_T_u8_6 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 []uint8
}

type Ordering uint8

type _goml_m_std_p_io_p_ErrorKind uint8

const (
    NotFound _goml_m_std_p_io_p_ErrorKind = 0
    PermissionDenied _goml_m_std_p_io_p_ErrorKind = 1
    AlreadyExists _goml_m_std_p_io_p_ErrorKind = 2
    InvalidInput _goml_m_std_p_io_p_ErrorKind = 3
    InvalidData _goml_m_std_p_io_p_ErrorKind = 4
    TimedOut _goml_m_std_p_io_p_ErrorKind = 5
    Interrupted _goml_m_std_p_io_p_ErrorKind = 6
    UnexpectedEof _goml_m_std_p_io_p_ErrorKind = 7
    WriteZero _goml_m_std_p_io_p_ErrorKind = 8
    BrokenPipe _goml_m_std_p_io_p_ErrorKind = 9
    WouldBlock _goml_m_std_p_io_p_ErrorKind = 10
    Unsupported _goml_m_std_p_io_p_ErrorKind = 11
    _goml_m_std_p_io_p_ErrorKind_Other _goml_m_std_p_io_p_ErrorKind = 12
)

type _goml_m_std_p_fs_p_FileType uint8

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

type _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error interface {
    is_goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error()
}

type _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Ok struct {
    _0 _goml_m_std_p_bytes_p_Bytes
}

func (_ _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Ok) is_goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error() {}

type _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Err struct {
    _0 _goml_m_std_p_fs_p_Error
}

func (_ _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Err) is_goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error() {}

type _goml_m_Result____string____std_p_fs_p_Error interface {
    is_goml_m_Result____string____std_p_fs_p_Error()
}

type _goml_m_Result____string____std_p_fs_p_Error_Ok struct {
    _0 string
}

func (_ _goml_m_Result____string____std_p_fs_p_Error_Ok) is_goml_m_Result____string____std_p_fs_p_Error() {}

type _goml_m_Result____string____std_p_fs_p_Error_Err struct {
    _0 _goml_m_std_p_fs_p_Error
}

func (_ _goml_m_Result____string____std_p_fs_p_Error_Err) is_goml_m_Result____string____std_p_fs_p_Error() {}

type _goml_m_Result_____o__q_____std_p_fs_p_Error interface {
    is_goml_m_Result_____o__q_____std_p_fs_p_Error()
}

type _goml_m_Result_____o__q_____std_p_fs_p_Error_Ok struct {
    _0 struct{}
}

func (_ _goml_m_Result_____o__q_____std_p_fs_p_Error_Ok) is_goml_m_Result_____o__q_____std_p_fs_p_Error() {}

type _goml_m_Result_____o__q_____std_p_fs_p_Error_Err struct {
    _0 _goml_m_std_p_fs_p_Error
}

func (_ _goml_m_Result_____o__q_____std_p_fs_p_Error_Err) is_goml_m_Result_____o__q_____std_p_fs_p_Error() {}

type _goml_m_Result____std_p_fs_p_Metadata____std_p_fs_p_Error interface {
    is_goml_m_Result____std_p_fs_p_Metadata____std_p_fs_p_Error()
}

type _goml_m_Result____std_p_fs_p_Metadata____std_p_fs_p_Error_Ok struct {
    _0 _goml_m_std_p_fs_p_Metadata
}

func (_ _goml_m_Result____std_p_fs_p_Metadata____std_p_fs_p_Error_Ok) is_goml_m_Result____std_p_fs_p_Metadata____std_p_fs_p_Error() {}

type _goml_m_Result____std_p_fs_p_Metadata____std_p_fs_p_Error_Err struct {
    _0 _goml_m_std_p_fs_p_Error
}

func (_ _goml_m_Result____std_p_fs_p_Metadata____std_p_fs_p_Error_Err) is_goml_m_Result____std_p_fs_p_Metadata____std_p_fs_p_Error() {}

type _goml_m_Result____bool____std_p_fs_p_Error interface {
    is_goml_m_Result____bool____std_p_fs_p_Error()
}

type _goml_m_Result____bool____std_p_fs_p_Error_Ok struct {
    _0 bool
}

func (_ _goml_m_Result____bool____std_p_fs_p_Error_Ok) is_goml_m_Result____bool____std_p_fs_p_Error() {}

type _goml_m_Result____bool____std_p_fs_p_Error_Err struct {
    _0 _goml_m_std_p_fs_p_Error
}

func (_ _goml_m_Result____bool____std_p_fs_p_Error_Err) is_goml_m_Result____bool____std_p_fs_p_Error() {}

type _goml_m_Result____Vec_l_string_r_____std_p_fs_p_Error interface {
    is_goml_m_Result____Vec_l_string_r_____std_p_fs_p_Error()
}

type _goml_m_Result____Vec_l_string_r_____std_p_fs_p_Error_Ok struct {
    _0 *_goml_vec_string
}

func (_ _goml_m_Result____Vec_l_string_r_____std_p_fs_p_Error_Ok) is_goml_m_Result____Vec_l_string_r_____std_p_fs_p_Error() {}

type _goml_m_Result____Vec_l_string_r_____std_p_fs_p_Error_Err struct {
    _0 _goml_m_std_p_fs_p_Error
}

func (_ _goml_m_Result____Vec_l_string_r_____std_p_fs_p_Error_Err) is_goml_m_Result____Vec_l_string_r_____std_p_fs_p_Error() {}

type _goml_m_Result____Vec_l_std_p_fs_p_DirEntry_r_____std_p_fs_p_Error interface {
    is_goml_m_Result____Vec_l_std_p_fs_p_DirEntry_r_____std_p_fs_p_Error()
}

type _goml_m_Result____Vec_l_std_p_fs_p_DirEntry_r_____std_p_fs_p_Error_Ok struct {
    _0 *_goml_vec__goml_m_std_p_fs_p_DirEntry
}

func (_ _goml_m_Result____Vec_l_std_p_fs_p_DirEntry_r_____std_p_fs_p_Error_Ok) is_goml_m_Result____Vec_l_std_p_fs_p_DirEntry_r_____std_p_fs_p_Error() {}

type _goml_m_Result____Vec_l_std_p_fs_p_DirEntry_r_____std_p_fs_p_Error_Err struct {
    _0 _goml_m_std_p_fs_p_Error
}

func (_ _goml_m_Result____Vec_l_std_p_fs_p_DirEntry_r_____std_p_fs_p_Error_Err) is_goml_m_Result____Vec_l_std_p_fs_p_DirEntry_r_____std_p_fs_p_Error() {}

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

func _goml_m_inherent_i_std_p_bytes_p_Bytes_i_std_p_bytes_p_Bytes_i_from__string(value__0 string) _goml_m_std_p_bytes_p_Bytes {
    var t0 *_goml_vec_uint8
    var inline0 *_goml_vec_uint8 = _goml_runtime_core_string_to_bytes(value__0)
    t0 = inline0
    var t1 _goml_m_std_p_bytes_p_Bytes = _goml_m_std_p_bytes_p_Bytes{
        values: t0,
    }
    return t1
}

func _goml_m_std_p_internal_p_host_p_println(value__0 string) struct{} {
    _goml_runtime_std_io_println(value__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_std_p_ut_he615d2de462723f24b2bd9e21c2fc9a4_o__string__utf8(self__0 _goml_m_std_p_bytes_p_Bytes) _goml_m_Result____string____std_p_utf8_p_Utf8Error {
    var t0 *_goml_vec_uint8
    var inline2 *_goml_vec_uint8 = self__0.values
    t0 = inline2
    var inline0 []uint8 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_as__slice____T__u8(t0)
    var inline1 _goml_m_Result____string____std_p_utf8_p_Utf8Error = _goml_m_std_p_utf8_p_decode__slice(inline0)
    return inline1
}

func _goml_m_inherent_i_std_p_utf8_p_Utf8Error_i_std_p_utf8_p_Utf8Error_i_new(valid_up_to__0 int, error_length__0 Option__isize) _goml_m_std_p_utf8_p_Utf8Error {
    var t0 _goml_m_std_p_utf8_p_Utf8Error = _goml_m_std_p_utf8_p_Utf8Error{
        valid_up_to_value: valid_up_to__0,
        error_length_value: error_length__0,
    }
    return t0
}

func _goml_m_std_p_utf8_p_continuation(value__0 uint8) bool {
    var t0 bool = value__0 >= 128
    if t0 {
        var t1 bool = value__0 <= 191
        return t1
    } else {
        return false
    }
}

func _goml_m_std_p_utf8_p_invalid(index__0 int) _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error {
    var t0 Option__isize = Option__isize{
        _p0: 1,
        _tag: 1,
    }
    var t1 _goml_m_std_p_utf8_p_Utf8Error
    var inline0 _goml_m_std_p_utf8_p_Utf8Error = _goml_m_std_p_utf8_p_Utf8Error{
        valid_up_to_value: index__0,
        error_length_value: t0,
    }
    t1 = inline0
    var t2 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error{
        _p0: t1,
        _tag: 1,
    }
    return t2
}

func _goml_m_std_p_utf8_p_validate(bytes__0 []uint8) _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error {
    var index__0 int = 0
    Loop_loop0:
    for {
        var t1 int = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_len____T__u8(bytes__0)
        var t2 bool = index__0 < t1
        if t2 {
            var first__0 uint8 = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_get____T__u8(bytes__0, index__0)
            var t3 bool = first__0 <= 127
            if t3 {
                var compound_old0 int = index__0
                var compound_value0 int = 1
                var t4 int = compound_old0 + compound_value0
                index__0 = t4
                continue
            } else {
                var t6 bool = first__0 >= 194
                var jp0 bool
                if t6 {
                    var t60 bool = first__0 <= 223
                    jp0 = t60
                } else {
                    jp0 = false
                }
                if jp0 {
                    var t7 int = index__0 + 1
                    var t8 int
                    var inline8 int = len(bytes__0)
                    t8 = inline8
                    var t9 bool = t7 >= t8
                    if t9 {
                        var inline0 _goml_m_std_p_utf8_p_Utf8Error = _goml_m_inherent_i_std_p_utf8_p_Utf8Error_i_std_p_utf8_p_Utf8Error_i_new(index__0, Option__isize{
                            _tag: 0,
                        })
                        var inline1 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error{
                            _p0: inline0,
                            _tag: 1,
                        }
                        return inline1
                    } else {
                        var t10 int = index__0 + 1
                        var t11 uint8
                        var inline7 uint8 = bytes__0[t10]
                        t11 = inline7
                        var t12 bool
                        var inline5 bool = t11 >= 128
                        if inline5 {
                            var inline6 bool = t11 <= 191
                            t12 = inline6
                        } else {
                            t12 = false
                        }
                        var t13 bool = !t12
                        if t13 {
                            var inline2 Option__isize = Option__isize{
                                _p0: 1,
                                _tag: 1,
                            }
                            var inline3 _goml_m_std_p_utf8_p_Utf8Error = _goml_m_inherent_i_std_p_utf8_p_Utf8Error_i_std_p_utf8_p_Utf8Error_i_new(index__0, inline2)
                            var inline4 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error{
                                _p0: inline3,
                                _tag: 1,
                            }
                            return inline4
                        } else {
                            var compound_old1 int = index__0
                            var compound_value1 int = 2
                            var t14 int = compound_old1 + compound_value1
                            index__0 = t14
                            continue
                        }
                    }
                } else {
                    var t16 bool = first__0 >= 224
                    var jp1 bool
                    if t16 {
                        var t59 bool = first__0 <= 239
                        jp1 = t59
                    } else {
                        jp1 = false
                    }
                    if jp1 {
                        var t17 int = index__0 + 2
                        var t18 int
                        var inline20 int = len(bytes__0)
                        t18 = inline20
                        var t19 bool = t17 >= t18
                        if t19 {
                            var inline9 _goml_m_std_p_utf8_p_Utf8Error = _goml_m_inherent_i_std_p_utf8_p_Utf8Error_i_std_p_utf8_p_Utf8Error_i_new(index__0, Option__isize{
                                _tag: 0,
                            })
                            var inline10 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error{
                                _p0: inline9,
                                _tag: 1,
                            }
                            return inline10
                        } else {
                            var t20 int = index__0 + 1
                            var second__0 uint8
                            var inline19 uint8 = bytes__0[t20]
                            second__0 = inline19
                            var t21 bool = first__0 == 224
                            var jp2 bool
                            if t21 {
                                var t29 bool = second__0 >= 160
                                if t29 {
                                    var t30 bool = second__0 <= 191
                                    jp2 = t30
                                } else {
                                    jp2 = false
                                }
                            } else {
                                var t31 bool = first__0 == 237
                                if t31 {
                                    var t32 bool = second__0 >= 128
                                    if t32 {
                                        var t33 bool = second__0 <= 159
                                        jp2 = t33
                                    } else {
                                        jp2 = false
                                    }
                                } else {
                                    var inline17 bool = second__0 >= 128
                                    if inline17 {
                                        var inline18 bool = second__0 <= 191
                                        jp2 = inline18
                                    } else {
                                        jp2 = false
                                    }
                                }
                            }
                            var t22 bool = !jp2
                            var jp3 bool
                            if t22 {
                                jp3 = true
                            } else {
                                var t25 int = index__0 + 2
                                var t26 uint8
                                var inline16 uint8 = bytes__0[t25]
                                t26 = inline16
                                var t27 bool
                                var inline14 bool = t26 >= 128
                                if inline14 {
                                    var inline15 bool = t26 <= 191
                                    t27 = inline15
                                } else {
                                    t27 = false
                                }
                                var t28 bool = !t27
                                jp3 = t28
                            }
                            if jp3 {
                                var inline11 Option__isize = Option__isize{
                                    _p0: 1,
                                    _tag: 1,
                                }
                                var inline12 _goml_m_std_p_utf8_p_Utf8Error = _goml_m_inherent_i_std_p_utf8_p_Utf8Error_i_std_p_utf8_p_Utf8Error_i_new(index__0, inline11)
                                var inline13 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error{
                                    _p0: inline12,
                                    _tag: 1,
                                }
                                return inline13
                            } else {
                                var compound_old2 int = index__0
                                var compound_value2 int = 3
                                var t23 int = compound_old2 + compound_value2
                                index__0 = t23
                                continue
                            }
                        }
                    } else {
                        var t34 bool = first__0 >= 240
                        var jp4 bool
                        if t34 {
                            var t58 bool = first__0 <= 244
                            jp4 = t58
                        } else {
                            jp4 = false
                        }
                        if jp4 {
                            var t35 int = index__0 + 3
                            var t36 int = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_len____T__u8(bytes__0)
                            var t37 bool = t35 >= t36
                            if t37 {
                                var inline21 _goml_m_std_p_utf8_p_Utf8Error = _goml_m_inherent_i_std_p_utf8_p_Utf8Error_i_std_p_utf8_p_Utf8Error_i_new(index__0, Option__isize{
                                    _tag: 0,
                                })
                                var inline22 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error{
                                    _p0: inline21,
                                    _tag: 1,
                                }
                                return inline22
                            } else {
                                var t38 int = index__0 + 1
                                var second__1 uint8 = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_get____T__u8(bytes__0, t38)
                                var t39 bool = first__0 == 240
                                var jp5 bool
                                if t39 {
                                    var t52 bool = second__1 >= 144
                                    if t52 {
                                        var t53 bool = second__1 <= 191
                                        jp5 = t53
                                    } else {
                                        jp5 = false
                                    }
                                } else {
                                    var t54 bool = first__0 == 244
                                    if t54 {
                                        var t55 bool = second__1 >= 128
                                        if t55 {
                                            var t56 bool = second__1 <= 143
                                            jp5 = t56
                                        } else {
                                            jp5 = false
                                        }
                                    } else {
                                        var inline23 bool = second__1 >= 128
                                        if inline23 {
                                            var inline24 bool = second__1 <= 191
                                            jp5 = inline24
                                        } else {
                                            jp5 = false
                                        }
                                    }
                                }
                                var t40 bool = !jp5
                                var jp6 bool
                                if t40 {
                                    jp6 = true
                                } else {
                                    var t48 int = index__0 + 2
                                    var t49 uint8 = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_get____T__u8(bytes__0, t48)
                                    var t50 bool = _goml_m_std_p_utf8_p_continuation(t49)
                                    var t51 bool = !t50
                                    jp6 = t51
                                }
                                var jp7 bool
                                if jp6 {
                                    jp7 = true
                                } else {
                                    var t44 int = index__0 + 3
                                    var t45 uint8 = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_get____T__u8(bytes__0, t44)
                                    var t46 bool = _goml_m_std_p_utf8_p_continuation(t45)
                                    var t47 bool = !t46
                                    jp7 = t47
                                }
                                if jp7 {
                                    var t41 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_std_p_utf8_p_invalid(index__0)
                                    return t41
                                } else {
                                    var compound_old3 int = index__0
                                    var compound_value3 int = 4
                                    var t42 int = compound_old3 + compound_value3
                                    index__0 = t42
                                    continue
                                }
                            }
                        } else {
                            var t57 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_std_p_utf8_p_invalid(index__0)
                            return t57
                        }
                    }
                }
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error{
        _tag: 0,
    }
    return t0
}

func _goml_m_std_p_utf8_p_decode__slice(bytes__0 []uint8) _goml_m_Result____string____std_p_utf8_p_Utf8Error {
    var mtmp0 _goml_m_Result_____o__q_____std_p_utf8_p_Utf8Error = _goml_m_std_p_utf8_p_validate(bytes__0)
    switch mtmp0._tag {
    case 0:
        var values__0 *_goml_vec_uint8 = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_to__vec____T__u8(bytes__0)
        var mtmp1 Tuple2_4bool_6string = string_from_utf8(values__0)
        var x0 bool = mtmp1._0
        var x1 string = mtmp1._1
        if x0 {
            var t0 _goml_m_Result____string____std_p_utf8_p_Utf8Error = _goml_m_Result____string____std_p_utf8_p_Utf8Error{
                _p0: x1,
                _tag: 0,
            }
            return t0
        } else {
            var t1 _goml_m_std_p_utf8_p_Utf8Error
            var inline0 int = 0
            var inline1 _goml_m_std_p_utf8_p_Utf8Error = _goml_m_std_p_utf8_p_Utf8Error{
                valid_up_to_value: inline0,
                error_length_value: Option__isize{
                    _tag: 0,
                },
            }
            t1 = inline1
            var t2 _goml_m_Result____string____std_p_utf8_p_Utf8Error = _goml_m_Result____string____std_p_utf8_p_Utf8Error{
                _p1: t1,
                _tag: 1,
            }
            return t2
        }
    case 1:
        var x2 _goml_m_std_p_utf8_p_Utf8Error = mtmp0._p0
        var t3 _goml_m_Result____string____std_p_utf8_p_Utf8Error = _goml_m_Result____string____std_p_utf8_p_Utf8Error{
            _p1: x2,
            _tag: 1,
        }
        return t3
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_io_p_kind__from__code(code__0 int) _goml_m_std_p_io_p_ErrorKind {
    switch code__0 {
    case 1:
        return NotFound
    case 2:
        return PermissionDenied
    case 3:
        return AlreadyExists
    case 4:
        return InvalidInput
    case 5:
        return InvalidData
    case 6:
        return TimedOut
    case 7:
        return Interrupted
    case 8:
        return UnexpectedEof
    case 9:
        return WriteZero
    case 10:
        return BrokenPipe
    case 11:
        return WouldBlock
    case 12:
        return Unsupported
    default:
        return _goml_m_std_p_io_p_ErrorKind_Other
    }
}

func _goml_m_inherent_i_std_p_io_p_ErrorDetails_i_std_p_io_p_ErrorDetails_i_new(kind__0 _goml_m_std_p_io_p_ErrorKind, operation__0 string, context__0 Option__string, raw_os_code__0 Option__isize, message__0 string) _goml_m_std_p_io_p_ErrorDetails {
    var t0 _goml_m_std_p_io_p_ErrorDetails = _goml_m_std_p_io_p_ErrorDetails{
        kind_value: kind__0,
        operation_value: operation__0,
        context_value: context__0,
        raw_os_code_value: raw_os_code__0,
        message_value: message__0,
    }
    return t0
}

func _goml_m_inherent_i_std_p_fs_p_Error_i_std_p_fs_p_Error_i_new(kind__0 _goml_m_std_p_io_p_ErrorKind, operation__0 string, path__0 Option__string, raw_os_code__0 Option__isize, message__0 string) _goml_m_std_p_fs_p_Error {
    var t0 _goml_m_std_p_io_p_ErrorDetails
    var inline0 _goml_m_std_p_io_p_ErrorDetails = _goml_m_std_p_io_p_ErrorDetails{
        kind_value: kind__0,
        operation_value: operation__0,
        context_value: path__0,
        raw_os_code_value: raw_os_code__0,
        message_value: message__0,
    }
    t0 = inline0
    var t1 _goml_m_std_p_fs_p_Error = _goml_m_std_p_fs_p_Error{
        details: t0,
    }
    return t1
}

func _goml_m_trait__impl_i_ToString_i_std_p_fs_p_Error_i_to__string(self__0 *_goml_m_std_p_fs_p_Error) string {
    var t0 _goml_m_std_p_io_p_ErrorDetails = (*self__0).details
    var inline0 Option__string = t0.context_value
    switch inline0._tag {
    case 0:
        var inline1 string = t0.operation_value
        var inline2 string = inline1 + ": "
        var inline3 string = t0.message_value
        var inline4 string = inline2 + inline3
        return inline4
    case 1:
        var inline5 string = inline0._p0
        var inline6 string = t0.operation_value
        var inline7 string = inline6 + " "
        var inline8 string = inline7 + inline5
        var inline9 string = inline8 + ": "
        var inline10 string = t0.message_value
        var inline11 string = inline9 + inline10
        return inline11
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_fs_p_read__bytes__structured(path__0 string) _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error {
    var mtmp0 Tuple6_4bool_10Vec_5uint8_3int_4bool_3int_6string
    var inline7 Tuple6_4bool_10Vec_5uint8_3int_4bool_3int_6string = _goml_runtime_std_fs_read_bytes_v2(path__0)
    mtmp0 = inline7
    var x0 bool = mtmp0._0
    var x1 *_goml_vec_uint8 = mtmp0._1
    var x2 int = mtmp0._2
    var x3 bool = mtmp0._3
    var x4 int = mtmp0._4
    var x5 string = mtmp0._5
    if x0 {
        var t0 _goml_m_std_p_bytes_p_Bytes
        var inline0 _goml_m_std_p_bytes_p_Bytes = _goml_m_std_p_bytes_p_Bytes{
            values: x1,
        }
        t0 = inline0
        var t1 _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error = _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Ok{
            _0: t0,
        }
        return t1
    } else {
        var t2 _goml_m_std_p_fs_p_Error
        var inline1 string = "read file"
        var inline2 _goml_m_std_p_io_p_ErrorKind = _goml_m_std_p_io_p_kind__from__code(x2)
        var inline3 Option__string = Option__string{
            _p0: path__0,
            _tag: 1,
        }
        var inline4 Option__isize
        if x3 {
            var inline6 Option__isize = Option__isize{
                _p0: x4,
                _tag: 1,
            }
            inline4 = inline6
        } else {
            inline4 = Option__isize{
                _tag: 0,
            }
        }
        var inline5 _goml_m_std_p_fs_p_Error = _goml_m_inherent_i_std_p_fs_p_Error_i_std_p_fs_p_Error_i_new(inline2, inline1, inline3, inline4, x5)
        t2 = inline5
        var t3 _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error = _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Err{
            _0: t2,
        }
        return t3
    }
}

func _goml_m_std_p_fs_p_write__bytes__structured(path__0 string, data__0 _goml_m_std_p_bytes_p_Bytes) _goml_m_Result_____o__q_____std_p_fs_p_Error {
    var t0 *_goml_vec_uint8
    var inline7 *_goml_vec_uint8 = data__0.values
    t0 = inline7
    var mtmp0 Tuple5_4bool_3int_4bool_3int_6string
    var inline6 Tuple5_4bool_3int_4bool_3int_6string = _goml_runtime_std_fs_write_bytes_v2(path__0, t0)
    mtmp0 = inline6
    var x0 bool = mtmp0._0
    var x1 int = mtmp0._1
    var x2 bool = mtmp0._2
    var x3 int = mtmp0._3
    var x4 string = mtmp0._4
    if x0 {
        var t1 _goml_m_Result_____o__q_____std_p_fs_p_Error = _goml_m_Result_____o__q_____std_p_fs_p_Error_Ok{
            _0: struct{}{},
        }
        return t1
    } else {
        var t2 _goml_m_std_p_fs_p_Error
        var inline0 string = "write file"
        var inline1 _goml_m_std_p_io_p_ErrorKind = _goml_m_std_p_io_p_kind__from__code(x1)
        var inline2 Option__string = Option__string{
            _p0: path__0,
            _tag: 1,
        }
        var inline3 Option__isize
        if x2 {
            var inline5 Option__isize = Option__isize{
                _p0: x3,
                _tag: 1,
            }
            inline3 = inline5
        } else {
            inline3 = Option__isize{
                _tag: 0,
            }
        }
        var inline4 _goml_m_std_p_fs_p_Error = _goml_m_inherent_i_std_p_fs_p_Error_i_std_p_fs_p_Error_i_new(inline1, inline0, inline2, inline3, x4)
        t2 = inline4
        var t3 _goml_m_Result_____o__q_____std_p_fs_p_Error = _goml_m_Result_____o__q_____std_p_fs_p_Error_Err{
            _0: t2,
        }
        return t3
    }
}

func _goml_m_std_p_fs_p_create__dir__all(path__0 string) _goml_m_Result_____o__q_____std_p_fs_p_Error {
    var mtmp0 Tuple2_4bool_6string
    var inline9 Tuple2_4bool_6string = _goml_runtime_std_fs_create_dir_all(path__0)
    mtmp0 = inline9
    var x0 bool = mtmp0._0
    var x1 string = mtmp0._1
    if x0 {
        var t0 _goml_m_Result_____o__q_____std_p_fs_p_Error = _goml_m_Result_____o__q_____std_p_fs_p_Error_Ok{
            _0: struct{}{},
        }
        return t0
    } else {
        var t1 _goml_m_std_p_fs_p_Error
        var inline0 string = "create directories"
        var inline1 int = 0
        var inline2 bool = false
        var inline3 int = 0
        var inline4 _goml_m_std_p_io_p_ErrorKind = _goml_m_std_p_io_p_kind__from__code(inline1)
        var inline5 Option__string = Option__string{
            _p0: path__0,
            _tag: 1,
        }
        var inline6 Option__isize
        if inline2 {
            var inline8 Option__isize = Option__isize{
                _p0: inline3,
                _tag: 1,
            }
            inline6 = inline8
        } else {
            inline6 = Option__isize{
                _tag: 0,
            }
        }
        var inline7 _goml_m_std_p_fs_p_Error = _goml_m_inherent_i_std_p_fs_p_Error_i_std_p_fs_p_Error_i_new(inline4, inline0, inline5, inline6, x1)
        t1 = inline7
        var t2 _goml_m_Result_____o__q_____std_p_fs_p_Error = _goml_m_Result_____o__q_____std_p_fs_p_Error_Err{
            _0: t1,
        }
        return t2
    }
}

func main0() struct{} {
    var t0 string = string_byte_slice("a你好z", 1, 7)
    _goml_m_std_p_io_p_println____T__string(t0)
    _goml_m_std_p_io_p_eprint____T__string("")
    _goml_m_std_p_io_p_eprintln____T__string("")
    var t1 _goml_m_Result_____o__q_____std_p_fs_p_Error = _goml_m_std_p_fs_p_create__dir__all("goml-self-host/nested")
    var t2 string
    switch t1.(type) {
    case _goml_m_Result_____o__q_____std_p_fs_p_Error_Ok:
        t2 = "ok"
    case _goml_m_Result_____o__q_____std_p_fs_p_Error_Err:
        var inline26 _goml_m_std_p_fs_p_Error = t1.(_goml_m_Result_____o__q_____std_p_fs_p_Error_Err)._0
        var inline27 string = _goml_m_trait__impl_i_ToString_i_std_p_fs_p_Error_i_to__string(&inline26)
        var inline28 string = "err " + inline27
        t2 = inline28
    default:
        panic("non-exhaustive match")
    }
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
    _goml_m_std_p_internal_p_host_p_println(inline24)
    var t3 _goml_m_Result_____o__q_____std_p_fs_p_Error
    var inline20 string = "goml-self-host/nested/output.txt"
    var inline21 string = "boot"
    var inline22 _goml_m_std_p_bytes_p_Bytes = _goml_m_inherent_i_std_p_bytes_p_Bytes_i_std_p_bytes_p_Bytes_i_from__string(inline21)
    var inline23 _goml_m_Result_____o__q_____std_p_fs_p_Error = _goml_m_std_p_fs_p_write__bytes__structured(inline20, inline22)
    t3 = inline23
    var t4 string
    switch t3.(type) {
    case _goml_m_Result_____o__q_____std_p_fs_p_Error_Ok:
        t4 = "ok"
    case _goml_m_Result_____o__q_____std_p_fs_p_Error_Err:
        var inline17 _goml_m_std_p_fs_p_Error = t3.(_goml_m_Result_____o__q_____std_p_fs_p_Error_Err)._0
        var inline18 string = _goml_m_trait__impl_i_ToString_i_std_p_fs_p_Error_i_to__string(&inline17)
        var inline19 string = "err " + inline18
        t4 = inline19
    default:
        panic("non-exhaustive match")
    }
    var inline15 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t4)
    _goml_m_std_p_internal_p_host_p_println(inline15)
    var t5 _goml_m_Result____string____std_p_fs_p_Error
    var inline6 string = "goml-self-host/nested/output.txt"
    var inline7 _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error = _goml_m_std_p_fs_p_read__bytes__structured(inline6)
    switch inline7.(type) {
    case _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Ok:
        var inline8 _goml_m_std_p_bytes_p_Bytes = inline7.(_goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Ok)._0
        var inline9 _goml_m_Result____string____std_p_utf8_p_Utf8Error = _goml_m_trait__impl_i_std_p_ut_he615d2de462723f24b2bd9e21c2fc9a4_o__string__utf8(inline8)
        var inline10 closure_env_std_fs_read_file_structured_1 = closure_env_std_fs_read_file_structured_1{
            path_0: inline6,
        }
        var inline11 func(_goml_m_std_p_utf8_p_Utf8Error) _goml_m_std_p_fs_p_Error = func(p0 _goml_m_std_p_utf8_p_Utf8Error) _goml_m_std_p_fs_p_Error {
            return _goml_m_inherent_i_closure__en_h3dd0cb03305303b9cfff8ae48a351c84_ured__1_i_apply(inline10, p0)
        }
        var inline12 _goml_m_Result____string____std_p_fs_p_Error = _goml_m_inherent_i_Result_i_Re_h48a6a59879d8846dfcb5cfdba927af0b_or____T__string(inline9, inline11)
        t5 = inline12
    case _goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Err:
        var inline13 _goml_m_std_p_fs_p_Error = inline7.(_goml_m_Result____std_p_bytes_p_Bytes____std_p_fs_p_Error_Err)._0
        var inline14 _goml_m_Result____string____std_p_fs_p_Error = _goml_m_Result____string____std_p_fs_p_Error_Err{
            _0: inline13,
        }
        t5 = inline14
    default:
        panic("non-exhaustive match")
    }
    var t6 string
    switch t5.(type) {
    case _goml_m_Result____string____std_p_fs_p_Error_Ok:
        var inline2 string = t5.(_goml_m_Result____string____std_p_fs_p_Error_Ok)._0
        t6 = inline2
    case _goml_m_Result____string____std_p_fs_p_Error_Err:
        var inline3 _goml_m_std_p_fs_p_Error = t5.(_goml_m_Result____string____std_p_fs_p_Error_Err)._0
        var inline4 string = _goml_m_trait__impl_i_ToString_i_std_p_fs_p_Error_i_to__string(&inline3)
        var inline5 string = "err " + inline4
        t6 = inline5
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t6)
    _goml_m_std_p_internal_p_host_p_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_isize_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_len____T__u8(self__0 []uint8) int {
    var t0 int = len(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_as__slice____T__u8(self__0 *_goml_vec_uint8) []uint8 {
    var t0 int
    var inline0 int = vec_len__Vec_5uint8(self__0)
    t0 = inline0
    var t1 []uint8 = self__0.items[0:t0]
    return t1
}

func _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_get____T__u8(self__0 []uint8, index__0 int) uint8 {
    var t0 uint8 = self__0[index__0]
    return t0
}

func _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_to__vec____T__u8(self__0 []uint8) *_goml_vec_uint8 {
    var t0 int
    var inline5 int = len(self__0)
    t0 = inline5
    var result__0 *_goml_vec_uint8
    var inline4 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(t0)
    result__0 = inline4
    var for_iter0 FnIterator__u8
    var inline3 FnIterator__u8 = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_iter____T__u8(self__0)
    for_iter0 = inline3
    Loop_loop0:
    for {
        var for_next0 Option__u8
        var inline1 func() Option__u8 = for_iter0.next_fn
        var inline2 Option__u8 = inline1()
        for_next0 = inline2
        switch for_next0 != Option__u8(0) {
        case false:
            break Loop_loop0
        case true:
            var x0 uint8 = uint8(uint16(for_next0) - 1)
            vec_push__Vec_5uint8(result__0, x0)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    return result__0
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

func _goml_m_inherent_i_string_i_string_i_byte__len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func _goml_m_inherent_i_Result_i_Re_h48a6a59879d8846dfcb5cfdba927af0b_or____T__string(self__0 _goml_m_Result____string____std_p_utf8_p_Utf8Error, map_fn__0 func(_goml_m_std_p_utf8_p_Utf8Error) _goml_m_std_p_fs_p_Error) _goml_m_Result____string____std_p_fs_p_Error {
    switch self__0._tag {
    case 0:
        var x0 string = self__0._p0
        var t0 _goml_m_Result____string____std_p_fs_p_Error = _goml_m_Result____string____std_p_fs_p_Error_Ok{
            _0: x0,
        }
        return t0
    case 1:
        var x1 _goml_m_std_p_utf8_p_Utf8Error = self__0._p1
        var t1 _goml_m_std_p_fs_p_Error = map_fn__0(x1)
        var t2 _goml_m_Result____string____std_p_fs_p_Error = _goml_m_Result____string____std_p_fs_p_Error_Err{
            _0: t1,
        }
        return t2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_io_p_println____T__string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_std_io_println(t0)
    return struct{}{}
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

func _goml_m_std_p_io_p_eprint____T__string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_std_io_eprint(t0)
    return struct{}{}
}

func _goml_m_std_p_io_p_eprintln____T__string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    var t1 string = t0 + "\n"
    _goml_runtime_std_io_eprint(t1)
    return struct{}{}
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_iter____T__u8(self__0 []uint8) FnIterator__u8 {
    var index__0 *ref_int_x = ref__Ref_3int(0)
    var len__0 int
    var inline1 int = len(self__0)
    len__0 = inline1
    var t0 closure_env_inherent_Slice_Slice_T_iter_T_u8_6 = closure_env_inherent_Slice_Slice_T_iter_T_u8_6{
        index_0: index__0,
        len_1: len__0,
        self_2: self__0,
    }
    var t1 func() Option__u8 = func() Option__u8 {
        return _goml_m_inherent_i_closure__en_h058fcc5d7e2318a71cafff326487f473___u8__6_i_apply(t0)
    }
    var inline0 FnIterator__u8 = FnIterator__u8{
        next_fn: t1,
    }
    return inline0
}

func utf8_invalid_decode() Tuple3_4bool_4char_3int {
    var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
        _0: false,
        _1: 0,
        _2: 0,
    }
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__get(self__0 string, index__0 int) uint8 {
    var t0 uint8 = _goml_runtime_core_string_byte_get(self__0, index__0)
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

func _goml_m_inherent_i_closure__en_h3dd0cb03305303b9cfff8ae48a351c84_ured__1_i_apply(env0 closure_env_std_fs_read_file_structured_1, error__0 _goml_m_std_p_utf8_p_Utf8Error) _goml_m_std_p_fs_p_Error {
    var path__0 string = env0.path_0
    var t0 Option__string = Option__string{
        _p0: path__0,
        _tag: 1,
    }
    var t1 string
    var inline3 string = "" + "invalid UTF-8 at byte "
    var inline4 int = error__0.valid_up_to_value
    var inline5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline4)
    var inline6 string = inline3 + inline5
    t1 = inline6
    var inline0 string = "read file"
    var inline1 _goml_m_std_p_io_p_ErrorDetails = _goml_m_inherent_i_std_p_io_p_ErrorDetails_i_std_p_io_p_ErrorDetails_i_new(InvalidData, inline0, t0, Option__isize{
        _tag: 0,
    }, t1)
    var inline2 _goml_m_std_p_fs_p_Error = _goml_m_std_p_fs_p_Error{
        details: inline1,
    }
    return inline2
}

func _goml_m_inherent_i_closure__en_h058fcc5d7e2318a71cafff326487f473___u8__6_i_apply(env0 closure_env_inherent_Slice_Slice_T_iter_T_u8_6) Option__u8 {
    var index__0 *ref_int_x = env0.index_0
    var len__0 int = env0.len_1
    var self__0 []uint8 = env0.self_2
    var current__0 int = ref_get__Ref_3int(index__0)
    var t0 bool = current__0 < len__0
    if t0 {
        var value__0 uint8
        var inline0 uint8 = self__0[current__0]
        value__0 = inline0
        var t1 int = current__0 + 1
        ref_set__Ref_3int(index__0, t1)
        var t2 Option__u8 = Option__u8(uint16(uint8(value__0)) + 1)
        return t2
    } else {
        return Option__u8(0)
    }
}

func main() {
    main0()
}
