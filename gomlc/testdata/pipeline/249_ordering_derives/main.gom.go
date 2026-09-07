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

func _goml_runtime_core_int_hash(x int) uint64 {
    return uint64(x)
}

func _goml_runtime_core_string_hash(s string) uint64 {
    var h uint64 = 14695981039346656037
    var i int = 0
    for {
        if i >= int(len(s)) {
            break
        }
        h = h * 1099511628211 + uint64(s[i])
        i = i + 1
    }
    return h
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

func array_get__Array_2_3int(arr [2]int, index int) int {
    return arr[index]
}

func array_get__Array_3_3int(arr [3]int, index int) int {
    return arr[index]
}

type _goml_vec_int struct {
    items []int
}

func vec_get__Vec_3int(vec *_goml_vec_int, index int) int {
    return vec.items[index]
}

func vec_len__Vec_3int(vec *_goml_vec_int) int {
    return int(len(vec.items))
}

type _goml_vec_uint8 struct {
    items []uint8
}

type _goml_vec_uint32 struct {
    items []uint32
}

type hashmap_Vec_3int_string_x_entry struct {
    active bool
    key *_goml_vec_int
    value string
}

type hashmap_Vec_3int_string_x struct {
    buckets map[uint64][]hashmap_Vec_3int_string_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_8Vec_3int_6string() *hashmap_Vec_3int_string_x {
    return &hashmap_Vec_3int_string_x{
        buckets: make(map[uint64][]hashmap_Vec_3int_string_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_lookup__HashMap_8Vec_3int_6string(m *hashmap_Vec_3int_string_x, key *_goml_vec_int) (string, bool, int, uint64) {
    if m == nil {
        var zero string
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Vec_l_isize_r__i_hash(key)
    var bucket []hashmap_Vec_3int_string_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Vec_3int_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Vec_l_isize_r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero string
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_8Vec_3int_6string(m *hashmap_Vec_3int_string_x, key *_goml_vec_int) Option__string {
    var value string
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_8Vec_3int_6string(m, key)
    if ok {
        return Option__string{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__string{
        _tag: 0,
    }
}

func hashmap_set__HashMap_8Vec_3int_6string(m *hashmap_Vec_3int_string_x, key *_goml_vec_int, value string) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Vec_l_isize_r__i_hash(key)
    var bucket []hashmap_Vec_3int_string_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Vec_3int_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Vec_l_isize_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Vec_3int_string_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Vec_3int_string_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_Tuple2_3int_6string_string_x_entry struct {
    active bool
    key Tuple2_3int_6string
    value string
}

type hashmap_Tuple2_3int_6string_string_x struct {
    buckets map[uint64][]hashmap_Tuple2_3int_6string_string_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_19Tuple2_3int_6string_6string() *hashmap_Tuple2_3int_6string_string_x {
    return &hashmap_Tuple2_3int_6string_string_x{
        buckets: make(map[uint64][]hashmap_Tuple2_3int_6string_string_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_lookup__HashMap_19Tuple2_3int_6string_6string(m *hashmap_Tuple2_3int_6string_string_x, key Tuple2_3int_6string) (string, bool, int, uint64) {
    if m == nil {
        var zero string
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i__o_isize_c_string_q__i_hash(key)
    var bucket []hashmap_Tuple2_3int_6string_string_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Tuple2_3int_6string_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i__o_isize_c_string_q__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero string
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_19Tuple2_3int_6string_6string(m *hashmap_Tuple2_3int_6string_string_x, key Tuple2_3int_6string) Option__string {
    var value string
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_19Tuple2_3int_6string_6string(m, key)
    if ok {
        return Option__string{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__string{
        _tag: 0,
    }
}

func hashmap_set__HashMap_19Tuple2_3int_6string_6string(m *hashmap_Tuple2_3int_6string_string_x, key Tuple2_3int_6string, value string) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i__o_isize_c_string_q__i_hash(key)
    var bucket []hashmap_Tuple2_3int_6string_string_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Tuple2_3int_6string_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i__o_isize_c_string_q__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Tuple2_3int_6string_string_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Tuple2_3int_6string_string_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_Array_2_3int_string_x_entry struct {
    active bool
    key [2]int
    value string
}

type hashmap_Array_2_3int_string_x struct {
    buckets map[uint64][]hashmap_Array_2_3int_string_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_12Array_2_3int_6string() *hashmap_Array_2_3int_string_x {
    return &hashmap_Array_2_3int_string_x{
        buckets: make(map[uint64][]hashmap_Array_2_3int_string_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_lookup__HashMap_12Array_2_3int_6string(m *hashmap_Array_2_3int_string_x, key [2]int) (string, bool, int, uint64) {
    if m == nil {
        var zero string
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i__l_isize_x3b_2_r__i_hash(key)
    var bucket []hashmap_Array_2_3int_string_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Array_2_3int_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i__l_isize_x3b_2_r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero string
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_12Array_2_3int_6string(m *hashmap_Array_2_3int_string_x, key [2]int) Option__string {
    var value string
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_12Array_2_3int_6string(m, key)
    if ok {
        return Option__string{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__string{
        _tag: 0,
    }
}

func hashmap_set__HashMap_12Array_2_3int_6string(m *hashmap_Array_2_3int_string_x, key [2]int, value string) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i__l_isize_x3b_2_r__i_hash(key)
    var bucket []hashmap_Array_2_3int_string_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Array_2_3int_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i__l_isize_x3b_2_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Array_2_3int_string_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Array_2_3int_string_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type Tuple2_3int_6string struct {
    _0 int
    _1 string
}

type Tuple2_3int_3int struct {
    _0 int
    _1 int
}

type Tuple2_3int_7float64 struct {
    _0 int
    _1 float64
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

type Version struct {
    major int
    minor int
}

type MaybeNumber struct {
    value float64
}

type NoTraits struct {}

type GenericPair__isize struct {
    first int
    second int
    nested *_goml_vec_int
}

type Ordering uint8

const (
    Less Ordering = 0
    Equal Ordering = 1
    Greater Ordering = 2
)

type Level struct {
    _p0 int
    _p1 int
    _tag uint8
}

type PartialLevel struct {
    _p0 float64
    _tag uint8
}

type Option__Ordering struct {
    _p0 Ordering
    _tag uint8
}

type Phantom__NoTraits uint8

const (
    First Phantom__NoTraits = 0
    Second Phantom__NoTraits = 1
)

type Option__isize struct {
    _p0 int
    _tag uint8
}

type Result__isize__string struct {
    _p1 string
    _p0 int
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

func _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_isize_i_cmp(self__0 int, other__0 int) Ordering {
    var t0 bool = self__0 < other__0
    if t0 {
        return Less
    } else {
        var t1 bool = self__0 > other__0
        if t1 {
            return Greater
        } else {
            return Equal
        }
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Version_i_partial__cmp(self__0 Version, other__0 Version) Option__Ordering {
    var t0 int = self__0.major
    var t1 int = other__0.major
    var commute_field0 Ordering
    var inline3 bool = t0 < t1
    var inline4 Ordering
    if inline3 {
        inline4 = Less
    } else {
        var inline5 bool = t0 > t1
        if inline5 {
            inline4 = Greater
        } else {
            inline4 = Equal
        }
    }
    commute_field0 = inline4
    switch commute_field0 {
    case Equal:
        var t2 int = self__0.minor
        var t3 int = other__0.minor
        var commute_field1 Ordering
        var inline0 bool = t2 < t3
        var inline1 Ordering
        if inline0 {
            inline1 = Less
        } else {
            var inline2 bool = t2 > t3
            if inline2 {
                inline1 = Greater
            } else {
                inline1 = Equal
            }
        }
        commute_field1 = inline1
        switch commute_field1 {
        case Equal:
            var t4 Option__Ordering = Option__Ordering{
                _p0: Equal,
                _tag: 1,
            }
            return t4
        default:
            var t5 Option__Ordering = Option__Ordering{
                _p0: commute_field1,
                _tag: 1,
            }
            return t5
        }
    default:
        var t6 Option__Ordering = Option__Ordering{
            _p0: commute_field0,
            _tag: 1,
        }
        return t6
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Version_i_lt(default_arg0 Version, default_arg1 Version) bool {
    var inline0 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Version_i_partial__cmp(default_arg0, default_arg1)
    switch inline0._tag {
    case 0:
        return false
    case 1:
        var inline1 Ordering = inline0._p0
        var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(inline1, Less)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_Version_i_cmp(self__0 Version, other__0 Version) Ordering {
    var t0 int = self__0.major
    var t1 int = other__0.major
    var _goml_m__i_derive1__ordering____0 Ordering
    var inline2 bool = t0 < t1
    if inline2 {
        _goml_m__i_derive1__ordering____0 = Less
    } else {
        var inline3 bool = t0 > t1
        if inline3 {
            _goml_m__i_derive1__ordering____0 = Greater
        } else {
            _goml_m__i_derive1__ordering____0 = Equal
        }
    }
    var t2 bool
    switch _goml_m__i_derive1__ordering____0 {
    case Less:
        t2 = false
    case Equal:
        t2 = true
    case Greater:
        t2 = false
    default:
        panic("non-exhaustive match")
    }
    if t2 {
        var t3 int = self__0.minor
        var t4 int = other__0.minor
        var _goml_m__i_derive0__ordering____0 Ordering
        var inline0 bool = t3 < t4
        if inline0 {
            _goml_m__i_derive0__ordering____0 = Less
        } else {
            var inline1 bool = t3 > t4
            if inline1 {
                _goml_m__i_derive0__ordering____0 = Greater
            } else {
                _goml_m__i_derive0__ordering____0 = Equal
            }
        }
        var t5 bool
        switch _goml_m__i_derive0__ordering____0 {
        case Less:
            t5 = false
        case Equal:
            t5 = true
        case Greater:
            t5 = false
        default:
            panic("non-exhaustive match")
        }
        if t5 {
            return Equal
        } else {
            return _goml_m__i_derive0__ordering____0
        }
    } else {
        return _goml_m__i_derive1__ordering____0
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Level_i_partial__cmp(self__0 Level, other__0 Level) Option__Ordering {
    var jp0 int
    switch self__0._tag {
    case 0:
        jp0 = 0
    case 1:
        jp0 = 1
    case 2:
        jp0 = 2
    default:
        panic("non-exhaustive match")
    }
    var jp1 int
    switch other__0._tag {
    case 0:
        jp1 = 0
    case 1:
        jp1 = 1
    case 2:
        jp1 = 2
    default:
        panic("non-exhaustive match")
    }
    var t0 bool = jp0 < jp1
    if t0 {
        var t1 Option__Ordering = Option__Ordering{
            _p0: Less,
            _tag: 1,
        }
        return t1
    } else {
        var t2 bool = jp0 > jp1
        if t2 {
            var t3 Option__Ordering = Option__Ordering{
                _p0: Greater,
                _tag: 1,
            }
            return t3
        } else {
            switch other__0._tag {
            case 0:
                switch self__0._tag {
                case 0:
                    var t4 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t4
                default:
                    var t5 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t5
                }
            case 1:
                var x0 int = other__0._p0
                var x1 int = other__0._p1
                switch self__0._tag {
                case 1:
                    var x2 int = self__0._p0
                    var x3 int = self__0._p1
                    var commute_field0 Ordering
                    var inline3 bool = x2 < x0
                    var inline4 Ordering
                    if inline3 {
                        inline4 = Less
                    } else {
                        var inline5 bool = x2 > x0
                        if inline5 {
                            inline4 = Greater
                        } else {
                            inline4 = Equal
                        }
                    }
                    commute_field0 = inline4
                    switch commute_field0 {
                    case Equal:
                        var commute_field1 Ordering
                        var inline0 bool = x3 < x1
                        var inline1 Ordering
                        if inline0 {
                            inline1 = Less
                        } else {
                            var inline2 bool = x3 > x1
                            if inline2 {
                                inline1 = Greater
                            } else {
                                inline1 = Equal
                            }
                        }
                        commute_field1 = inline1
                        switch commute_field1 {
                        case Equal:
                            var t6 Option__Ordering = Option__Ordering{
                                _p0: Equal,
                                _tag: 1,
                            }
                            return t6
                        default:
                            var t7 Option__Ordering = Option__Ordering{
                                _p0: commute_field1,
                                _tag: 1,
                            }
                            return t7
                        }
                    default:
                        var t8 Option__Ordering = Option__Ordering{
                            _p0: commute_field0,
                            _tag: 1,
                        }
                        return t8
                    }
                default:
                    var t9 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t9
                }
            case 2:
                var x4 int = other__0._p0
                var x5 int = other__0._p1
                switch self__0._tag {
                case 2:
                    var x6 int = self__0._p0
                    var x7 int = self__0._p1
                    var commute_field2 Ordering
                    var inline9 bool = x6 < x4
                    var inline10 Ordering
                    if inline9 {
                        inline10 = Less
                    } else {
                        var inline11 bool = x6 > x4
                        if inline11 {
                            inline10 = Greater
                        } else {
                            inline10 = Equal
                        }
                    }
                    commute_field2 = inline10
                    switch commute_field2 {
                    case Equal:
                        var commute_field3 Ordering
                        var inline6 bool = x7 < x5
                        var inline7 Ordering
                        if inline6 {
                            inline7 = Less
                        } else {
                            var inline8 bool = x7 > x5
                            if inline8 {
                                inline7 = Greater
                            } else {
                                inline7 = Equal
                            }
                        }
                        commute_field3 = inline7
                        switch commute_field3 {
                        case Equal:
                            var t10 Option__Ordering = Option__Ordering{
                                _p0: Equal,
                                _tag: 1,
                            }
                            return t10
                        default:
                            var t11 Option__Ordering = Option__Ordering{
                                _p0: commute_field3,
                                _tag: 1,
                            }
                            return t11
                        }
                    default:
                        var t12 Option__Ordering = Option__Ordering{
                            _p0: commute_field2,
                            _tag: 1,
                        }
                        return t12
                    }
                default:
                    var t13 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t13
                }
            default:
                panic("non-exhaustive match")
            }
        }
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Level_i_lt(default_arg0 Level, default_arg1 Level) bool {
    var inline0 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Level_i_partial__cmp(default_arg0, default_arg1)
    switch inline0._tag {
    case 0:
        return false
    case 1:
        var inline1 Ordering = inline0._p0
        var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(inline1, Less)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_Level_i_cmp(self__0 Level, other__0 Level) Ordering {
    var jp0 int
    switch self__0._tag {
    case 0:
        jp0 = 0
    case 1:
        jp0 = 1
    case 2:
        jp0 = 2
    default:
        panic("non-exhaustive match")
    }
    var jp1 int
    switch other__0._tag {
    case 0:
        jp1 = 0
    case 1:
        jp1 = 1
    case 2:
        jp1 = 2
    default:
        panic("non-exhaustive match")
    }
    var t0 bool = jp0 < jp1
    if t0 {
        return Less
    } else {
        var t1 bool = jp0 > jp1
        if t1 {
            return Greater
        } else {
            switch other__0._tag {
            case 0:
                switch self__0._tag {
                case 0:
                    return Equal
                default:
                    return Equal
                }
            case 1:
                var x0 int = other__0._p0
                var x1 int = other__0._p1
                switch self__0._tag {
                case 1:
                    var x2 int = self__0._p0
                    var x3 int = self__0._p1
                    var _goml_m__i_derive7__ordering____0 Ordering
                    var inline2 bool = x2 < x0
                    if inline2 {
                        _goml_m__i_derive7__ordering____0 = Less
                    } else {
                        var inline3 bool = x2 > x0
                        if inline3 {
                            _goml_m__i_derive7__ordering____0 = Greater
                        } else {
                            _goml_m__i_derive7__ordering____0 = Equal
                        }
                    }
                    var t2 bool
                    switch _goml_m__i_derive7__ordering____0 {
                    case Less:
                        t2 = false
                    case Equal:
                        t2 = true
                    case Greater:
                        t2 = false
                    default:
                        panic("non-exhaustive match")
                    }
                    if t2 {
                        var _goml_m__i_derive4__ordering____0 Ordering
                        var inline0 bool = x3 < x1
                        if inline0 {
                            _goml_m__i_derive4__ordering____0 = Less
                        } else {
                            var inline1 bool = x3 > x1
                            if inline1 {
                                _goml_m__i_derive4__ordering____0 = Greater
                            } else {
                                _goml_m__i_derive4__ordering____0 = Equal
                            }
                        }
                        var t3 bool
                        switch _goml_m__i_derive4__ordering____0 {
                        case Less:
                            t3 = false
                        case Equal:
                            t3 = true
                        case Greater:
                            t3 = false
                        default:
                            panic("non-exhaustive match")
                        }
                        if t3 {
                            return Equal
                        } else {
                            return _goml_m__i_derive4__ordering____0
                        }
                    } else {
                        return _goml_m__i_derive7__ordering____0
                    }
                default:
                    return Equal
                }
            case 2:
                var x4 int = other__0._p0
                var x5 int = other__0._p1
                switch self__0._tag {
                case 2:
                    var x6 int = self__0._p0
                    var x7 int = self__0._p1
                    var _goml_m__i_derive13__ordering____0 Ordering = _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_isize_i_cmp(x6, x4)
                    var t4 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(_goml_m__i_derive13__ordering____0, Equal)
                    if t4 {
                        var _goml_m__i_derive10__ordering____0 Ordering = _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_isize_i_cmp(x7, x5)
                        var t5 bool
                        switch _goml_m__i_derive10__ordering____0 {
                        case Less:
                            t5 = false
                        case Equal:
                            t5 = true
                        case Greater:
                            t5 = false
                        default:
                            panic("non-exhaustive match")
                        }
                        if t5 {
                            return Equal
                        } else {
                            return _goml_m__i_derive10__ordering____0
                        }
                    } else {
                        return _goml_m__i_derive13__ordering____0
                    }
                default:
                    return Equal
                }
            default:
                panic("non-exhaustive match")
            }
        }
    }
}

func _goml_m_trait__impl_i_PartialEq_i_MaybeNumber_i_eq(self__0 MaybeNumber, other__0 MaybeNumber) bool {
    var t0 float64 = self__0.value
    var t1 float64 = other__0.value
    var inline0 bool = t0 == t1
    return inline0
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_MaybeNumber_i_partial__cmp(self__0 MaybeNumber, other__0 MaybeNumber) Option__Ordering {
    var t0 float64 = self__0.value
    var t1 float64 = other__0.value
    var commute_field0 Ordering
    var inline0 bool = t0 < t1
    if inline0 {
        commute_field0 = Less
        switch commute_field0 {
        case Equal:
            var t2 Option__Ordering = Option__Ordering{
                _p0: Equal,
                _tag: 1,
            }
            return t2
        default:
            var t3 Option__Ordering = Option__Ordering{
                _p0: commute_field0,
                _tag: 1,
            }
            return t3
        }
    } else {
        var inline1 bool = t0 > t1
        if inline1 {
            commute_field0 = Greater
            switch commute_field0 {
            case Equal:
                var t2 Option__Ordering = Option__Ordering{
                    _p0: Equal,
                    _tag: 1,
                }
                return t2
            default:
                var t3 Option__Ordering = Option__Ordering{
                    _p0: commute_field0,
                    _tag: 1,
                }
                return t3
            }
        } else {
            var inline2 bool = t0 == t1
            if inline2 {
                commute_field0 = Equal
                switch commute_field0 {
                case Equal:
                    var t2 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t2
                default:
                    var t3 Option__Ordering = Option__Ordering{
                        _p0: commute_field0,
                        _tag: 1,
                    }
                    return t3
                }
            } else {
                return Option__Ordering{
                    _tag: 0,
                }
            }
        }
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_PartialLevel_i_partial__cmp(self__0 PartialLevel, other__0 PartialLevel) Option__Ordering {
    var jp0 int
    switch self__0._tag {
    case 0:
        jp0 = 0
    case 1:
        jp0 = 1
    default:
        panic("non-exhaustive match")
    }
    var jp1 int
    switch other__0._tag {
    case 0:
        jp1 = 0
    case 1:
        jp1 = 1
    default:
        panic("non-exhaustive match")
    }
    var t0 bool = jp0 < jp1
    if t0 {
        var t1 Option__Ordering = Option__Ordering{
            _p0: Less,
            _tag: 1,
        }
        return t1
    } else {
        var t2 bool = jp0 > jp1
        if t2 {
            var t3 Option__Ordering = Option__Ordering{
                _p0: Greater,
                _tag: 1,
            }
            return t3
        } else {
            switch other__0._tag {
            case 0:
                var x0 float64 = other__0._p0
                switch self__0._tag {
                case 0:
                    var x1 float64 = self__0._p0
                    var commute_field0 Ordering
                    var inline0 bool = x1 < x0
                    if inline0 {
                        commute_field0 = Less
                        switch commute_field0 {
                        case Equal:
                            var t4 Option__Ordering = Option__Ordering{
                                _p0: Equal,
                                _tag: 1,
                            }
                            return t4
                        default:
                            var t5 Option__Ordering = Option__Ordering{
                                _p0: commute_field0,
                                _tag: 1,
                            }
                            return t5
                        }
                    } else {
                        var inline1 bool = x1 > x0
                        if inline1 {
                            commute_field0 = Greater
                            switch commute_field0 {
                            case Equal:
                                var t4 Option__Ordering = Option__Ordering{
                                    _p0: Equal,
                                    _tag: 1,
                                }
                                return t4
                            default:
                                var t5 Option__Ordering = Option__Ordering{
                                    _p0: commute_field0,
                                    _tag: 1,
                                }
                                return t5
                            }
                        } else {
                            var inline2 bool = x1 == x0
                            if inline2 {
                                commute_field0 = Equal
                                switch commute_field0 {
                                case Equal:
                                    var t4 Option__Ordering = Option__Ordering{
                                        _p0: Equal,
                                        _tag: 1,
                                    }
                                    return t4
                                default:
                                    var t5 Option__Ordering = Option__Ordering{
                                        _p0: commute_field0,
                                        _tag: 1,
                                    }
                                    return t5
                                }
                            } else {
                                return Option__Ordering{
                                    _tag: 0,
                                }
                            }
                        }
                    }
                default:
                    var t6 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t6
                }
            case 1:
                switch self__0._tag {
                case 1:
                    var t7 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t7
                default:
                    var t8 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t8
                }
            default:
                panic("non-exhaustive match")
            }
        }
    }
}

func ordering_name(value__0 Ordering) string {
    switch value__0 {
    case Less:
        return "less"
    case Equal:
        return "equal"
    case Greater:
        return "greater"
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var first__0 Version = Version{
        major: 1,
        minor: 9,
    }
    var second__0 Version = Version{
        major: 2,
        minor: 0,
    }
    var t0 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Version_i_lt(first__0, second__0)
    var t1 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t0)
    println__T_string(t1)
    var t2 Ordering = _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_Version_i_cmp(first__0, second__0)
    var t3 string = ordering_name(t2)
    println__T_string(t3)
    var t4 Level = Level{
        _p0: 0,
        _p1: 0,
        _tag: 1,
    }
    var t5 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Level_i_lt(Level{
        _tag: 0,
    }, t4)
    var t6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t5)
    println__T_string(t6)
    var t7 Level = Level{
        _p0: 1,
        _p1: 9,
        _tag: 1,
    }
    var t8 Level = Level{
        _p0: 2,
        _p1: 0,
        _tag: 1,
    }
    var t9 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Level_i_lt(t7, t8)
    var t10 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t9)
    println__T_string(t10)
    var t11 Level = Level{
        _p0: 3,
        _p1: 1,
        _tag: 2,
    }
    var t12 Level = Level{
        _p0: 3,
        _p1: 1,
        _tag: 2,
    }
    var t13 Ordering = _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_Level_i_cmp(t11, t12)
    var t14 string = ordering_name(t13)
    println__T_string(t14)
    var zero__0 float64 = 0
    var t15 float64 = zero__0 / zero__0
    var nan__0 MaybeNumber = MaybeNumber{
        value: t15,
    }
    var t16 bool = _goml_m_trait__impl_i_PartialEq_i_MaybeNumber_i_eq(nan__0, nan__0)
    var t17 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t16)
    println__T_string(t17)
    var t18 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_MaybeNumber_i_partial__cmp(nan__0, nan__0)
    var t19 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__none____T__Ordering(t18)
    var t20 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t19)
    println__T_string(t20)
    var t21 [1]int = [1]int{3}
    var t22 *_goml_vec_int = func(values [1]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [1]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t21)
    var generic_first__0 GenericPair__isize = GenericPair__isize{
        first: 1,
        second: 2,
        nested: t22,
    }
    var t23 [1]int = [1]int{0}
    var t24 *_goml_vec_int = func(values [1]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [1]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t23)
    var generic_second__0 GenericPair__isize = GenericPair__isize{
        first: 1,
        second: 3,
        nested: t24,
    }
    var t25 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_GenericPair____isize_i_lt(generic_first__0, generic_second__0)
    var t26 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t25)
    println__T_string(t26)
    var phantom_first__0 Phantom__NoTraits = First
    var phantom_second__0 Phantom__NoTraits = Second
    var t27 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Phantom____NoTraits_i_lt(phantom_first__0, phantom_second__0)
    var t28 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t27)
    println__T_string(t28)
    var t29 float64 = zero__0 / zero__0
    var partial_nan__0 PartialLevel = PartialLevel{
        _p0: t29,
        _tag: 0,
    }
    var t30 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_PartialLevel_i_partial__cmp(partial_nan__0, partial_nan__0)
    var t31 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__none____T__Ordering(t30)
    var t32 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t31)
    println__T_string(t32)
    var t33 [2]int = [2]int{1, 2}
    var first_values__0 *_goml_vec_int = func(values [2]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [2]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t33)
    var t34 [2]int = [2]int{1, 3}
    var second_values__0 *_goml_vec_int = func(values [2]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [2]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t34)
    var t35 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Vec_l_isize_r__i_lt(first_values__0, second_values__0)
    var t36 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t35)
    println__T_string(t36)
    var t37 Option__isize = Option__isize{
        _p0: 2,
        _tag: 1,
    }
    var t38 Option__isize = Option__isize{
        _p0: 3,
        _tag: 1,
    }
    var t39 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Option____isize_i_lt(t37, t38)
    var t40 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t39)
    println__T_string(t40)
    var ok__0 Result__isize__string = Result__isize__string{
        _p0: 1,
        _tag: 0,
    }
    var error__0 Result__isize__string = Result__isize__string{
        _p1: "error",
        _tag: 1,
    }
    var t41 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Result____isize____string_i_lt(ok__0, error__0)
    var t42 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t41)
    println__T_string(t42)
    var t43 []int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_slice____T__isize(first_values__0, 0, 2)
    var t44 []int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_slice____T__isize(second_values__0, 0, 2)
    var t45 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Slice_l_isize_r__i_lt(t43, t44)
    var t46 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t45)
    println__T_string(t46)
    var values__0 *hashmap_Vec_3int_string_x = _goml_m_inherent_i_HashMap_i_H_hca0cffb759e2572d7286850a43f94208_r_____V__string()
    _goml_m_inherent_i_HashMap_i_H_hfc6d57fdfd8ace95a78adbd7155c7d4b_r_____V__string(values__0, first_values__0, "vector")
    var t47 [2]int = [2]int{1, 2}
    var t48 *_goml_vec_int = func(values [2]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [2]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t47)
    var t49 Option__string = _goml_m_inherent_i_HashMap_i_H_h0069152d187f39ec37f55a9b6f59774d_r_____V__string(values__0, t48)
    var t50 string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__string(t49, "missing")
    println__T_string(t50)
    var default_tuple__0 Tuple2_3int_6string = _goml_m_trait__impl_i_Default_i__o_isize_c_string_q__i_default()
    var t51 int = default_tuple__0._0
    var t52 int = 0
    var t53 bool = _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(t51, t52)
    var jp0 bool
    if t53 {
        var t84 string = default_tuple__0._1
        var t85 string = ""
        var inline27 bool = t84 == t85
        jp0 = inline27
    } else {
        jp0 = false
    }
    var t54 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(jp0)
    println__T_string(t54)
    var default_array__0 [3]int = _goml_m_trait__impl_i_Default_i__l_isize_x3b_3_r__i_default()
    var _eq_rhs0 [3]int = [3]int{0, 0, 0}
    var t55 int = array_get__Array_3_3int(default_array__0, 0)
    var t56 int = array_get__Array_3_3int(_eq_rhs0, 0)
    var t57 bool = _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(t55, t56)
    var jp1 bool
    if t57 {
        var t79 int = array_get__Array_3_3int(default_array__0, 1)
        var t80 int = array_get__Array_3_3int(_eq_rhs0, 1)
        var t81 bool
        var inline26 bool = t79 == t80
        t81 = inline26
        if t81 {
            var t82 int = array_get__Array_3_3int(default_array__0, 2)
            var t83 int = array_get__Array_3_3int(_eq_rhs0, 2)
            var inline25 bool = t82 == t83
            jp1 = inline25
        } else {
            jp1 = false
        }
    } else {
        jp1 = false
    }
    var t58 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(jp1)
    println__T_string(t58)
    var t59 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: 1,
        _1: 2,
    }
    var t60 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: 1,
        _1: 3,
    }
    var t61 bool = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i__o_isize_c_isize_q__i_lt(t59, t60)
    var t62 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t61)
    var inline23 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t62)
    _goml_runtime_core_string_println(inline23)
    var t63 [2]int = [2]int{1, 2}
    var t64 [2]int = [2]int{1, 3}
    var t65 Ordering = _goml_m_trait__impl_i_std_p_cmp_p_Ord_i__l_isize_x3b_2_r__i_cmp(t63, t64)
    var t66 string = ordering_name(t65)
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t66)
    _goml_runtime_core_string_println(inline21)
    var t67 float64 = zero__0 / zero__0
    var tuple_nan__0 Tuple2_3int_7float64 = Tuple2_3int_7float64{
        _0: 0.0,
        _1: t67,
    }
    var t68 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i__o_isize_c_f64_q__i_partial__cmp(tuple_nan__0, tuple_nan__0)
    var t69 bool
    var inline19 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__Ordering(t68)
    var inline20 bool = !inline19
    t69 = inline20
    var t70 string
    var inline18 string = _goml_runtime_core_bool_to_string(t69)
    t70 = inline18
    var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t70)
    _goml_runtime_core_string_println(inline16)
    var tuple_values__0 *hashmap_Tuple2_3int_6string_string_x
    var inline15 *hashmap_Tuple2_3int_6string_string_x = hashmap_new__HashMap_19Tuple2_3int_6string_6string()
    tuple_values__0 = inline15
    var t71 Tuple2_3int_6string = Tuple2_3int_6string{
        _0: 1,
        _1: "one",
    }
    var inline13 string = "tuple"
    hashmap_set__HashMap_19Tuple2_3int_6string_6string(tuple_values__0, t71, inline13)
    var t72 Tuple2_3int_6string = Tuple2_3int_6string{
        _0: 1,
        _1: "one",
    }
    var t73 Option__string
    var inline12 Option__string = hashmap_get__HashMap_19Tuple2_3int_6string_6string(tuple_values__0, t72)
    t73 = inline12
    var t74 string
    var inline10 string = "missing"
    switch t73._tag {
    case 0:
        t74 = inline10
    case 1:
        var inline11 string = t73._p0
        t74 = inline11
    default:
        panic("non-exhaustive match")
    }
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t74)
    _goml_runtime_core_string_println(inline8)
    var array_values__0 *hashmap_Array_2_3int_string_x
    var inline7 *hashmap_Array_2_3int_string_x = hashmap_new__HashMap_12Array_2_3int_6string()
    array_values__0 = inline7
    var t75 [2]int = [2]int{1, 2}
    var inline5 string = "array"
    hashmap_set__HashMap_12Array_2_3int_6string(array_values__0, t75, inline5)
    var t76 [2]int = [2]int{1, 2}
    var t77 Option__string
    var inline4 Option__string = hashmap_get__HashMap_12Array_2_3int_6string(array_values__0, t76)
    t77 = inline4
    var t78 string
    var inline2 string = "missing"
    switch t77._tag {
    case 0:
        t78 = inline2
    case 1:
        var inline3 string = t77._p0
        t78 = inline3
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t78)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(self__0 Ordering, other__0 Ordering) bool {
    if self__0 <= Greater {
        return other__0 == self__0
    } else {
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_lt____Self__isize(self__0 int, other__0 int) bool {
    var commute_field0 Ordering
    var inline0 bool = self__0 < other__0
    var inline1 Ordering
    if inline0 {
        inline1 = Less
    } else {
        var inline2 bool = self__0 > other__0
        if inline2 {
            inline1 = Greater
        } else {
            inline1 = Equal
        }
    }
    commute_field0 = inline1
    switch commute_field0 {
    case Less:
        return true
    case Equal:
        return false
    case Greater:
        return false
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(self__0 int, other__0 int) bool {
    var t0 bool = self__0 == other__0
    return t0
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__none____T__Ordering(self__0 Option__Ordering) bool {
    var t0 bool
    switch self__0._tag {
    case 0:
        t0 = false
    case 1:
        t0 = true
    default:
        panic("non-exhaustive match")
    }
    var t1 bool = !t0
    return t1
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_GenericPair____isize_i_lt(default_arg0 GenericPair__isize, default_arg1 GenericPair__isize) bool {
    var inline0 Option__Ordering = _goml_m_trait__impl_i_std_p_cm_hf6de147f382a0c7851836d0995f09328__i_partial__cmp(default_arg0, default_arg1)
    switch inline0._tag {
    case 0:
        return false
    case 1:
        var inline1 Ordering = inline0._p0
        var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(inline1, Less)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Phantom____NoTraits_i_lt(default_arg0 Phantom__NoTraits, default_arg1 Phantom__NoTraits) bool {
    var inline0 Option__Ordering = _goml_m_trait__impl_i_std_p_cm_h70a1f5151189c8228387cc52486880c2__i_partial__cmp(default_arg0, default_arg1)
    switch inline0._tag {
    case 0:
        return false
    case 1:
        var inline1 Ordering = inline0._p0
        var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(inline1, Less)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Vec_l_isize_r__i_lt(default_arg0 *_goml_vec_int, default_arg1 *_goml_vec_int) bool {
    var inline0 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Vec_l_isize_r__i_partial__cmp(default_arg0, default_arg1)
    switch inline0._tag {
    case 0:
        return false
    case 1:
        var inline1 Ordering = inline0._p0
        var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(inline1, Less)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Option____isize_i_lt(default_arg0 Option__isize, default_arg1 Option__isize) bool {
    var inline0 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Option____isize_i_partial__cmp(default_arg0, default_arg1)
    switch inline0._tag {
    case 0:
        return false
    case 1:
        var inline1 Ordering = inline0._p0
        var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(inline1, Less)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Result____isize____string_i_lt(default_arg0 Result__isize__string, default_arg1 Result__isize__string) bool {
    var inline0 Option__Ordering = _goml_m_trait__impl_i_std_p_cm_hc3f80c466e4f303ea6e82f1783f2c793__i_partial__cmp(default_arg0, default_arg1)
    switch inline0._tag {
    case 0:
        return false
    case 1:
        var inline1 Ordering = inline0._p0
        var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(inline1, Less)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_slice____T__isize(self__0 *_goml_vec_int, start__0 int, end__0 int) []int {
    var t0 []int = self__0.items[start__0:end__0]
    return t0
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Slice_l_isize_r__i_lt(default_arg0 []int, default_arg1 []int) bool {
    var inline0 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Slice_l_isize_r__i_partial__cmp(default_arg0, default_arg1)
    switch inline0._tag {
    case 0:
        return false
    case 1:
        var inline1 Ordering = inline0._p0
        var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_Ordering_i_eq(inline1, Less)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_HashMap_i_H_hca0cffb759e2572d7286850a43f94208_r_____V__string() *hashmap_Vec_3int_string_x {
    var t0 *hashmap_Vec_3int_string_x = hashmap_new__HashMap_8Vec_3int_6string()
    return t0
}

func _goml_m_inherent_i_HashMap_i_H_hfc6d57fdfd8ace95a78adbd7155c7d4b_r_____V__string(self__0 *hashmap_Vec_3int_string_x, key__0 *_goml_vec_int, value__0 string) struct{} {
    hashmap_set__HashMap_8Vec_3int_6string(self__0, key__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_H_h0069152d187f39ec37f55a9b6f59774d_r_____V__string(self__0 *hashmap_Vec_3int_string_x, key__0 *_goml_vec_int) Option__string {
    var t0 Option__string = hashmap_get__HashMap_8Vec_3int_6string(self__0, key__0)
    return t0
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__string(self__0 Option__string, fallback__0 string) string {
    switch self__0._tag {
    case 0:
        return fallback__0
    case 1:
        var x0 string = self__0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_Default_i__o_isize_c_string_q__i_default() Tuple2_3int_6string {
    var t0 int
    t0 = 0
    var t1 string
    t1 = ""
    var t2 Tuple2_3int_6string = Tuple2_3int_6string{
        _0: t0,
        _1: t1,
    }
    return t2
}

func _goml_m_trait__impl_i_Default_i__l_isize_x3b_3_r__i_default() [3]int {
    var t0 int
    t0 = 0
    var t1 int
    t1 = 0
    var t2 int
    t2 = 0
    var t3 [3]int = [3]int{t0, t1, t2}
    return t3
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i__o_isize_c_isize_q__i_lt(self Tuple2_3int_3int, other Tuple2_3int_3int) bool {
    var t0 int = self._0
    var t1 int = other._0
    var t2 bool
    var inline3 bool = t0 == t1
    t2 = inline3
    if t2 {
        var t3 int = self._1
        var t4 int = other._1
        var t5 bool
        var inline1 bool = t3 == t4
        t5 = inline1
        if t5 {
            return false
        } else {
            var t6 int = self._1
            var t7 int = other._1
            var inline0 bool = _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_lt____Self__isize(t6, t7)
            return inline0
        }
    } else {
        var t8 int = self._0
        var t9 int = other._0
        var inline2 bool = _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_lt____Self__isize(t8, t9)
        return inline2
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_Ord_i__l_isize_x3b_2_r__i_cmp(self [2]int, other [2]int) Ordering {
    var t0 int = array_get__Array_2_3int(self, 0)
    var t1 int = array_get__Array_2_3int(other, 0)
    var _structural_ordering_0 Ordering
    var inline2 bool = t0 < t1
    if inline2 {
        _structural_ordering_0 = Less
    } else {
        var inline3 bool = t0 > t1
        if inline3 {
            _structural_ordering_0 = Greater
        } else {
            _structural_ordering_0 = Equal
        }
    }
    var t2 bool
    switch _structural_ordering_0 {
    case Less:
        t2 = false
    case Equal:
        t2 = true
    case Greater:
        t2 = false
    default:
        panic("non-exhaustive match")
    }
    if t2 {
        var t3 int = array_get__Array_2_3int(self, 1)
        var t4 int = array_get__Array_2_3int(other, 1)
        var _structural_ordering_1 Ordering
        var inline0 bool = t3 < t4
        if inline0 {
            _structural_ordering_1 = Less
        } else {
            var inline1 bool = t3 > t4
            if inline1 {
                _structural_ordering_1 = Greater
            } else {
                _structural_ordering_1 = Equal
            }
        }
        var t5 bool
        switch _structural_ordering_1 {
        case Less:
            t5 = false
        case Equal:
            t5 = true
        case Greater:
            t5 = false
        default:
            panic("non-exhaustive match")
        }
        if t5 {
            return Equal
        } else {
            return _structural_ordering_1
        }
    } else {
        return _structural_ordering_0
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i__o_isize_c_f64_q__i_partial__cmp(self Tuple2_3int_7float64, other Tuple2_3int_7float64) Option__Ordering {
    var t0 int = self._0
    var t1 int = other._0
    var _structural_partial_ordering_0 Option__Ordering
    var commute_field0 Ordering
    var inline6 bool = t0 < t1
    var inline7 Ordering
    if inline6 {
        inline7 = Less
    } else {
        var inline9 bool = t0 > t1
        if inline9 {
            inline7 = Greater
        } else {
            inline7 = Equal
        }
    }
    var inline8 Option__Ordering = Option__Ordering{
        _p0: inline7,
        _tag: 1,
    }
    _structural_partial_ordering_0 = inline8
    commute_field0 = inline7
    var t2 bool
    switch commute_field0 {
    case Less:
        t2 = false
    case Equal:
        t2 = true
    case Greater:
        t2 = false
    default:
        panic("non-exhaustive match")
    }
    if t2 {
        var t3 float64 = self._1
        var t4 float64 = other._1
        var _structural_partial_ordering_1 Option__Ordering
        var commute_field1 Ordering
        var inline0 bool = t3 < t4
        if inline0 {
            var inline1 Option__Ordering = Option__Ordering{
                _p0: Less,
                _tag: 1,
            }
            _structural_partial_ordering_1 = inline1
            commute_field1 = Less
            var t5 bool
            switch commute_field1 {
            case Less:
                t5 = false
            case Equal:
                t5 = true
            case Greater:
                t5 = false
            default:
                panic("non-exhaustive match")
            }
            if t5 {
                var t6 Option__Ordering = Option__Ordering{
                    _p0: Equal,
                    _tag: 1,
                }
                return t6
            } else {
                return _structural_partial_ordering_1
            }
        } else {
            var inline2 bool = t3 > t4
            if inline2 {
                var inline3 Option__Ordering = Option__Ordering{
                    _p0: Greater,
                    _tag: 1,
                }
                _structural_partial_ordering_1 = inline3
                commute_field1 = Greater
                var t5 bool
                switch commute_field1 {
                case Less:
                    t5 = false
                case Equal:
                    t5 = true
                case Greater:
                    t5 = false
                default:
                    panic("non-exhaustive match")
                }
                if t5 {
                    var t6 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t6
                } else {
                    return _structural_partial_ordering_1
                }
            } else {
                var inline4 bool = t3 == t4
                if inline4 {
                    var inline5 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    _structural_partial_ordering_1 = inline5
                    commute_field1 = Equal
                    var t5 bool
                    switch commute_field1 {
                    case Less:
                        t5 = false
                    case Equal:
                        t5 = true
                    case Greater:
                        t5 = false
                    default:
                        panic("non-exhaustive match")
                    }
                    if t5 {
                        var t6 Option__Ordering = Option__Ordering{
                            _p0: Equal,
                            _tag: 1,
                        }
                        return t6
                    } else {
                        return _structural_partial_ordering_1
                    }
                } else {
                    return Option__Ordering{
                        _tag: 0,
                    }
                }
            }
        }
    } else {
        return _structural_partial_ordering_0
    }
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__Ordering(self__0 Option__Ordering) bool {
    switch self__0._tag {
    case 0:
        return false
    case 1:
        return true
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cm_hf6de147f382a0c7851836d0995f09328__i_partial__cmp(self__0 GenericPair__isize, other__0 GenericPair__isize) Option__Ordering {
    var t0 int = self__0.first
    var t1 int = other__0.first
    var commute_field0 Ordering
    var inline3 bool = t0 < t1
    var inline4 Ordering
    if inline3 {
        inline4 = Less
    } else {
        var inline5 bool = t0 > t1
        if inline5 {
            inline4 = Greater
        } else {
            inline4 = Equal
        }
    }
    commute_field0 = inline4
    switch commute_field0 {
    case Equal:
        var t2 int = self__0.second
        var t3 int = other__0.second
        var commute_field1 Ordering
        var inline0 bool = t2 < t3
        var inline1 Ordering
        if inline0 {
            inline1 = Less
        } else {
            var inline2 bool = t2 > t3
            if inline2 {
                inline1 = Greater
            } else {
                inline1 = Equal
            }
        }
        commute_field1 = inline1
        switch commute_field1 {
        case Equal:
            var t4 *_goml_vec_int = self__0.nested
            var t5 *_goml_vec_int = other__0.nested
            var mtmp0 Option__Ordering = _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Vec_l_isize_r__i_partial__cmp(t4, t5)
            switch mtmp0._tag {
            case 0:
                return Option__Ordering{
                    _tag: 0,
                }
            case 1:
                var x0 Ordering = mtmp0._p0
                switch x0 {
                case Equal:
                    var t6 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t6
                default:
                    var t7 Option__Ordering = Option__Ordering{
                        _p0: x0,
                        _tag: 1,
                    }
                    return t7
                }
            default:
                panic("non-exhaustive match")
            }
        default:
            var t8 Option__Ordering = Option__Ordering{
                _p0: commute_field1,
                _tag: 1,
            }
            return t8
        }
    default:
        var t9 Option__Ordering = Option__Ordering{
            _p0: commute_field0,
            _tag: 1,
        }
        return t9
    }
}

func _goml_m_trait__impl_i_std_p_cm_h70a1f5151189c8228387cc52486880c2__i_partial__cmp(self__0 Phantom__NoTraits, other__0 Phantom__NoTraits) Option__Ordering {
    var jp0 int
    switch self__0 {
    case First:
        jp0 = 0
    case Second:
        jp0 = 1
    default:
        panic("non-exhaustive match")
    }
    var jp1 int
    switch other__0 {
    case First:
        jp1 = 0
    case Second:
        jp1 = 1
    default:
        panic("non-exhaustive match")
    }
    var t0 bool = jp0 < jp1
    if t0 {
        var t1 Option__Ordering = Option__Ordering{
            _p0: Less,
            _tag: 1,
        }
        return t1
    } else {
        var t2 bool = jp0 > jp1
        if t2 {
            var t3 Option__Ordering = Option__Ordering{
                _p0: Greater,
                _tag: 1,
            }
            return t3
        } else {
            switch other__0 {
            case First:
                switch self__0 {
                case First:
                    var t4 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t4
                default:
                    var t5 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t5
                }
            case Second:
                switch self__0 {
                case Second:
                    var t6 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t6
                default:
                    var t7 Option__Ordering = Option__Ordering{
                        _p0: Equal,
                        _tag: 1,
                    }
                    return t7
                }
            default:
                panic("non-exhaustive match")
            }
        }
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Vec_l_isize_r__i_partial__cmp(self__0 *_goml_vec_int, other__0 *_goml_vec_int) Option__Ordering {
    var t0 int
    var inline12 int = vec_len__Vec_3int(self__0)
    t0 = inline12
    var t1 int
    var inline11 int = vec_len__Vec_3int(other__0)
    t1 = inline11
    var t2 bool = t0 < t1
    var jp0 int
    if t2 {
        var inline9 int = vec_len__Vec_3int(self__0)
        jp0 = inline9
    } else {
        var inline10 int = vec_len__Vec_3int(other__0)
        jp0 = inline10
    }
    var index__0 int = 0
    Loop_loop0:
    for {
        var t5 bool = index__0 < jp0
        if t5 {
            var t6 int = vec_get__Vec_3int(self__0, index__0)
            var t7 int = vec_get__Vec_3int(other__0, index__0)
            var commute_field0 Ordering
            var inline6 bool = t6 < t7
            var inline7 Ordering
            if inline6 {
                inline7 = Less
            } else {
                var inline8 bool = t6 > t7
                if inline8 {
                    inline7 = Greater
                } else {
                    inline7 = Equal
                }
            }
            commute_field0 = inline7
            switch commute_field0 {
            case Equal:
                var compound_old0 int = index__0
                var compound_value0 int = 1
                var t8 int = compound_old0 + compound_value0
                index__0 = t8
                continue
            default:
                var t10 Option__Ordering = Option__Ordering{
                    _p0: commute_field0,
                    _tag: 1,
                }
                return t10
            }
        } else {
            break Loop_loop0
        }
    }
    var t3 int
    var inline5 int = vec_len__Vec_3int(self__0)
    t3 = inline5
    var t4 int
    var inline4 int = vec_len__Vec_3int(other__0)
    t4 = inline4
    var inline0 bool = t3 < t4
    var inline1 Ordering
    if inline0 {
        inline1 = Less
    } else {
        var inline3 bool = t3 > t4
        if inline3 {
            inline1 = Greater
        } else {
            inline1 = Equal
        }
    }
    var inline2 Option__Ordering = Option__Ordering{
        _p0: inline1,
        _tag: 1,
    }
    return inline2
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Option____isize_i_partial__cmp(self__0 Option__isize, other__0 Option__isize) Option__Ordering {
    switch other__0._tag {
    case 0:
        switch self__0._tag {
        case 0:
            var t0 Option__Ordering = Option__Ordering{
                _p0: Equal,
                _tag: 1,
            }
            return t0
        case 1:
            var t1 Option__Ordering = Option__Ordering{
                _p0: Greater,
                _tag: 1,
            }
            return t1
        default:
            panic("non-exhaustive match")
        }
    case 1:
        var x0 int = other__0._p0
        switch self__0._tag {
        case 0:
            var t2 Option__Ordering = Option__Ordering{
                _p0: Less,
                _tag: 1,
            }
            return t2
        case 1:
            var x1 int = self__0._p0
            var inline0 bool = x1 < x0
            var inline1 Ordering
            if inline0 {
                inline1 = Less
            } else {
                var inline3 bool = x1 > x0
                if inline3 {
                    inline1 = Greater
                } else {
                    inline1 = Equal
                }
            }
            var inline2 Option__Ordering = Option__Ordering{
                _p0: inline1,
                _tag: 1,
            }
            return inline2
        default:
            panic("non-exhaustive match")
        }
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cm_hc3f80c466e4f303ea6e82f1783f2c793__i_partial__cmp(self__0 Result__isize__string, other__0 Result__isize__string) Option__Ordering {
    switch other__0._tag {
    case 0:
        var x0 int = other__0._p0
        switch self__0._tag {
        case 0:
            var x1 int = self__0._p0
            var inline0 bool = x1 < x0
            var inline1 Ordering
            if inline0 {
                inline1 = Less
            } else {
                var inline3 bool = x1 > x0
                if inline3 {
                    inline1 = Greater
                } else {
                    inline1 = Equal
                }
            }
            var inline2 Option__Ordering = Option__Ordering{
                _p0: inline1,
                _tag: 1,
            }
            return inline2
        case 1:
            var t0 Option__Ordering = Option__Ordering{
                _p0: Greater,
                _tag: 1,
            }
            return t0
        default:
            panic("non-exhaustive match")
        }
    case 1:
        var x2 string = other__0._p1
        switch self__0._tag {
        case 0:
            var t1 Option__Ordering = Option__Ordering{
                _p0: Less,
                _tag: 1,
            }
            return t1
        case 1:
            var x3 string = self__0._p1
            var inline4 bool = x3 < x2
            var inline5 Ordering
            if inline4 {
                inline5 = Less
            } else {
                var inline7 bool = x3 > x2
                if inline7 {
                    inline5 = Greater
                } else {
                    inline5 = Equal
                }
            }
            var inline6 Option__Ordering = Option__Ordering{
                _p0: inline5,
                _tag: 1,
            }
            return inline6
        default:
            panic("non-exhaustive match")
        }
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_PartialOrd_i_Slice_l_isize_r__i_partial__cmp(self__0 []int, other__0 []int) Option__Ordering {
    var t0 int
    var inline12 int = len(self__0)
    t0 = inline12
    var t1 int
    var inline11 int = len(other__0)
    t1 = inline11
    var t2 bool = t0 < t1
    var jp0 int
    if t2 {
        var inline9 int = len(self__0)
        jp0 = inline9
    } else {
        var inline10 int = len(other__0)
        jp0 = inline10
    }
    var index__0 int = 0
    Loop_loop0:
    for {
        var t5 bool = index__0 < jp0
        if t5 {
            var t6 int = self__0[index__0]
            var t7 int = other__0[index__0]
            var commute_field0 Ordering
            var inline6 bool = t6 < t7
            var inline7 Ordering
            if inline6 {
                inline7 = Less
            } else {
                var inline8 bool = t6 > t7
                if inline8 {
                    inline7 = Greater
                } else {
                    inline7 = Equal
                }
            }
            commute_field0 = inline7
            switch commute_field0 {
            case Equal:
                var compound_old0 int = index__0
                var compound_value0 int = 1
                var t8 int = compound_old0 + compound_value0
                index__0 = t8
                continue
            default:
                var t10 Option__Ordering = Option__Ordering{
                    _p0: commute_field0,
                    _tag: 1,
                }
                return t10
            }
        } else {
            break Loop_loop0
        }
    }
    var t3 int
    var inline5 int = len(self__0)
    t3 = inline5
    var t4 int
    var inline4 int = len(other__0)
    t4 = inline4
    var inline0 bool = t3 < t4
    var inline1 Ordering
    if inline0 {
        inline1 = Less
    } else {
        var inline3 bool = t3 > t4
        if inline3 {
            inline1 = Greater
        } else {
            inline1 = Equal
        }
    }
    var inline2 Option__Ordering = Option__Ordering{
        _p0: inline1,
        _tag: 1,
    }
    return inline2
}

func _goml_m_trait__impl_i_PartialEq_i_Vec_l_isize_r__i_eq(self__0 *_goml_vec_int, other__0 *_goml_vec_int) bool {
    var t0 int
    var inline3 int = vec_len__Vec_3int(self__0)
    t0 = inline3
    var t1 int
    var inline2 int = vec_len__Vec_3int(other__0)
    t1 = inline2
    var t2 bool = t0 != t1
    if t2 {
        return false
    } else {
        var index__0 int = 0
        Loop_loop0:
        for {
            var t3 int
            var inline1 int = vec_len__Vec_3int(self__0)
            t3 = inline1
            var t4 bool = index__0 < t3
            if t4 {
                var t5 int = vec_get__Vec_3int(self__0, index__0)
                var t6 int = vec_get__Vec_3int(other__0, index__0)
                var t7 bool
                var inline0 bool = t5 == t6
                t7 = inline0
                if t7 {
                    var compound_old0 int = index__0
                    var compound_value0 int = 1
                    var t8 int = compound_old0 + compound_value0
                    index__0 = t8
                    continue
                } else {
                    return false
                }
            } else {
                break Loop_loop0
            }
        }
        return true
    }
}

func _goml_m_trait__impl_i_Hash_i_Vec_l_isize_r__i_hash(self__0 *_goml_vec_int) uint64 {
    var value__0 uint64 = 14695981039346656037
    var index__0 int = 0
    Loop_loop0:
    for {
        var t0 int
        var inline1 int = vec_len__Vec_3int(self__0)
        t0 = inline1
        var t1 bool = index__0 < t0
        if t1 {
            var t2 uint64 = value__0 * 1099511628211
            var t3 int = vec_get__Vec_3int(self__0, index__0)
            var t4 uint64
            var inline0 uint64 = _goml_runtime_core_int_hash(t3)
            t4 = inline0
            var t5 uint64 = t2 + t4
            value__0 = t5
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t6 int = compound_old0 + compound_value0
            index__0 = t6
            continue
        } else {
            break Loop_loop0
        }
    }
    return value__0
}

func _goml_m_trait__impl_i_PartialEq_i__o_isize_c_string_q__i_eq(self Tuple2_3int_6string, other Tuple2_3int_6string) bool {
    var t0 int = self._0
    var t1 int = other._0
    var t2 bool
    var inline1 bool = t0 == t1
    t2 = inline1
    if t2 {
        var t3 string = self._1
        var t4 string = other._1
        var t5 bool
        var inline0 bool = t3 == t4
        t5 = inline0
        if t5 {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}

func _goml_m_trait__impl_i_Hash_i__o_isize_c_string_q__i_hash(self Tuple2_3int_6string) uint64 {
    var _structural_hash_0 uint64 = 14695981039346656037
    var t0 uint64 = _structural_hash_0 * 1099511628211
    var t1 int = self._0
    var t2 uint64
    var inline1 uint64 = _goml_runtime_core_int_hash(t1)
    t2 = inline1
    var _structural_hash_1 uint64 = t0 + t2
    var t3 uint64 = _structural_hash_1 * 1099511628211
    var t4 string = self._1
    var t5 uint64
    var inline0 uint64 = _goml_runtime_core_string_hash(t4)
    t5 = inline0
    var _structural_hash_2 uint64 = t3 + t5
    return _structural_hash_2
}

func _goml_m_trait__impl_i_PartialEq_i__l_isize_x3b_2_r__i_eq(self [2]int, other [2]int) bool {
    var t0 int = array_get__Array_2_3int(self, 0)
    var t1 int = array_get__Array_2_3int(other, 0)
    var t2 bool
    var inline1 bool = t0 == t1
    t2 = inline1
    if t2 {
        var t3 int = array_get__Array_2_3int(self, 1)
        var t4 int = array_get__Array_2_3int(other, 1)
        var t5 bool
        var inline0 bool = t3 == t4
        t5 = inline0
        if t5 {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}

func _goml_m_trait__impl_i_Hash_i__l_isize_x3b_2_r__i_hash(self [2]int) uint64 {
    var _structural_hash_0 uint64 = 14695981039346656037
    var t0 uint64 = _structural_hash_0 * 1099511628211
    var t1 int = array_get__Array_2_3int(self, 0)
    var t2 uint64
    var inline1 uint64 = _goml_runtime_core_int_hash(t1)
    t2 = inline1
    var _structural_hash_1 uint64 = t0 + t2
    var t3 uint64 = _structural_hash_1 * 1099511628211
    var t4 int = array_get__Array_2_3int(self, 1)
    var t5 uint64
    var inline0 uint64 = _goml_runtime_core_int_hash(t4)
    t5 = inline0
    var _structural_hash_2 uint64 = t3 + t5
    return _structural_hash_2
}

func main() {
    main0()
}
