package main

import (
    _goml_os "os"
    _goml_reflect "reflect"
)

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

type _goml_vec_Ref_4unit struct {
    items []*ref_unit_x
}

func vec_new__Vec_9Ref_4unit() *_goml_vec_Ref_4unit {
    return &_goml_vec_Ref_4unit{
        items: nil,
    }
}

func vec_push__Vec_9Ref_4unit(vec *_goml_vec_Ref_4unit, elem *ref_unit_x) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_9Ref_4unit(vec *_goml_vec_Ref_4unit, index int) *ref_unit_x {
    return vec.items[index]
}

type _goml_vec_Ref_5Empty struct {
    items []*ref_Empty_x
}

func vec_new__Vec_10Ref_5Empty() *_goml_vec_Ref_5Empty {
    return &_goml_vec_Ref_5Empty{
        items: nil,
    }
}

func vec_push__Vec_10Ref_5Empty(vec *_goml_vec_Ref_5Empty, elem *ref_Empty_x) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_10Ref_5Empty(vec *_goml_vec_Ref_5Empty, index int) *ref_Empty_x {
    return vec.items[index]
}

type _goml_vec_Ref_6Nested struct {
    items []*ref_Nested_x
}

func vec_new__Vec_11Ref_6Nested() *_goml_vec_Ref_6Nested {
    return &_goml_vec_Ref_6Nested{
        items: nil,
    }
}

func vec_push__Vec_11Ref_6Nested(vec *_goml_vec_Ref_6Nested, elem *ref_Nested_x) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_11Ref_6Nested(vec *_goml_vec_Ref_6Nested, index int) *ref_Nested_x {
    return vec.items[index]
}

type _goml_vec_Ref_13Array_0_4unit struct {
    items []*ref_Array_0_4unit_x
}

func vec_new__Vec_19Ref_13Array_0_4unit() *_goml_vec_Ref_13Array_0_4unit {
    return &_goml_vec_Ref_13Array_0_4unit{
        items: nil,
    }
}

func vec_push__Vec_19Ref_13Array_0_4unit(vec *_goml_vec_Ref_13Array_0_4unit, elem *ref_Array_0_4unit_x) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_19Ref_13Array_0_4unit(vec *_goml_vec_Ref_13Array_0_4unit, index int) *ref_Array_0_4unit_x {
    return vec.items[index]
}

type _goml_vec_Ref_13Array_3_4unit struct {
    items []*ref_Array_3_4unit_x
}

func vec_new__Vec_19Ref_13Array_3_4unit() *_goml_vec_Ref_13Array_3_4unit {
    return &_goml_vec_Ref_13Array_3_4unit{
        items: nil,
    }
}

func vec_push__Vec_19Ref_13Array_3_4unit(vec *_goml_vec_Ref_13Array_3_4unit, elem *ref_Array_3_4unit_x) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_19Ref_13Array_3_4unit(vec *_goml_vec_Ref_13Array_3_4unit, index int) *ref_Array_3_4unit_x {
    return vec.items[index]
}

type _goml_vec_Ref_4bool struct {
    items []*ref_bool_x
}

func vec_new__Vec_9Ref_4bool() *_goml_vec_Ref_4bool {
    return &_goml_vec_Ref_4bool{
        items: nil,
    }
}

func vec_push__Vec_9Ref_4bool(vec *_goml_vec_Ref_4bool, elem *ref_bool_x) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_9Ref_4bool(vec *_goml_vec_Ref_4bool, index int) *ref_bool_x {
    return vec.items[index]
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

type ref_unit_x struct {
    value struct{}
    identity uint8
}

func ref__Ref_4unit(value struct{}) *ref_unit_x {
    return &ref_unit_x{
        value: value,
    }
}

func ref_set__Ref_4unit(reference *ref_unit_x, value struct{}) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_4unit(a *ref_unit_x, b *ref_unit_x) bool {
    return a == b
}

func ptr_hash__Ref_4unit(reference *ref_unit_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type ref_Empty_x struct {
    value Empty
    identity uint8
}

func ref__Ref_5Empty(value Empty) *ref_Empty_x {
    return &ref_Empty_x{
        value: value,
    }
}

func ref_set__Ref_5Empty(reference *ref_Empty_x, value Empty) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_5Empty(a *ref_Empty_x, b *ref_Empty_x) bool {
    return a == b
}

func ptr_hash__Ref_5Empty(reference *ref_Empty_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type ref_Nested_x struct {
    value Nested
    identity uint8
}

func ref__Ref_6Nested(value Nested) *ref_Nested_x {
    return &ref_Nested_x{
        value: value,
    }
}

func ref_set__Ref_6Nested(reference *ref_Nested_x, value Nested) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_6Nested(a *ref_Nested_x, b *ref_Nested_x) bool {
    return a == b
}

func ptr_hash__Ref_6Nested(reference *ref_Nested_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type ref_Array_0_4unit_x struct {
    value [0]struct{}
    identity uint8
}

func ref__Ref_13Array_0_4unit(value [0]struct{}) *ref_Array_0_4unit_x {
    return &ref_Array_0_4unit_x{
        value: value,
    }
}

func ref_set__Ref_13Array_0_4unit(reference *ref_Array_0_4unit_x, value [0]struct{}) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_13Array_0_4unit(a *ref_Array_0_4unit_x, b *ref_Array_0_4unit_x) bool {
    return a == b
}

func ptr_hash__Ref_13Array_0_4unit(reference *ref_Array_0_4unit_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type ref_Array_3_4unit_x struct {
    value [3]struct{}
    identity uint8
}

func ref__Ref_13Array_3_4unit(value [3]struct{}) *ref_Array_3_4unit_x {
    return &ref_Array_3_4unit_x{
        value: value,
    }
}

func ref_set__Ref_13Array_3_4unit(reference *ref_Array_3_4unit_x, value [3]struct{}) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_13Array_3_4unit(a *ref_Array_3_4unit_x, b *ref_Array_3_4unit_x) bool {
    return a == b
}

func ptr_hash__Ref_13Array_3_4unit(reference *ref_Array_3_4unit_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type ref_bool_x struct {
    value bool
}

func ref__Ref_4bool(value bool) *ref_bool_x {
    return &ref_bool_x{
        value: value,
    }
}

func ref_set__Ref_4bool(reference *ref_bool_x, value bool) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_4bool(a *ref_bool_x, b *ref_bool_x) bool {
    return a == b
}

func ptr_hash__Ref_4bool(reference *ref_bool_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type hashmap_Ref_4unit_int_x_entry struct {
    active bool
    key *ref_unit_x
    value int
}

type hashmap_Ref_4unit_int_x struct {
    buckets map[uint64][]hashmap_Ref_4unit_int_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_9Ref_4unit_3int() *hashmap_Ref_4unit_int_x {
    return &hashmap_Ref_4unit_int_x{
        buckets: make(map[uint64][]hashmap_Ref_4unit_int_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_9Ref_4unit_3int(m *hashmap_Ref_4unit_int_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_9Ref_4unit_3int(m *hashmap_Ref_4unit_int_x, key *ref_unit_x) (int, bool, int, uint64) {
    if m == nil {
        var zero int
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l__o__q__r__i_hash(key)
    var bucket []hashmap_Ref_4unit_int_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_4unit_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l__o__q__r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero int
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_9Ref_4unit_3int(m *hashmap_Ref_4unit_int_x, key *ref_unit_x) Option__isize {
    var value int
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_9Ref_4unit_3int(m, key)
    if ok {
        return Option__isize{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__isize{
        _tag: 0,
    }
}

func hashmap_set__HashMap_9Ref_4unit_3int(m *hashmap_Ref_4unit_int_x, key *ref_unit_x, value int) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l__o__q__r__i_hash(key)
    var bucket []hashmap_Ref_4unit_int_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_4unit_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l__o__q__r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_4unit_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_4unit_int_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_Ref_5Empty_int_x_entry struct {
    active bool
    key *ref_Empty_x
    value int
}

type hashmap_Ref_5Empty_int_x struct {
    buckets map[uint64][]hashmap_Ref_5Empty_int_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_10Ref_5Empty_3int() *hashmap_Ref_5Empty_int_x {
    return &hashmap_Ref_5Empty_int_x{
        buckets: make(map[uint64][]hashmap_Ref_5Empty_int_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_10Ref_5Empty_3int(m *hashmap_Ref_5Empty_int_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_10Ref_5Empty_3int(m *hashmap_Ref_5Empty_int_x, key *ref_Empty_x) (int, bool, int, uint64) {
    if m == nil {
        var zero int
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Empty_r__i_hash(key)
    var bucket []hashmap_Ref_5Empty_int_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_5Empty_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Empty_r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero int
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_10Ref_5Empty_3int(m *hashmap_Ref_5Empty_int_x, key *ref_Empty_x) Option__isize {
    var value int
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_10Ref_5Empty_3int(m, key)
    if ok {
        return Option__isize{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__isize{
        _tag: 0,
    }
}

func hashmap_set__HashMap_10Ref_5Empty_3int(m *hashmap_Ref_5Empty_int_x, key *ref_Empty_x, value int) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Empty_r__i_hash(key)
    var bucket []hashmap_Ref_5Empty_int_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_5Empty_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Empty_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_5Empty_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_5Empty_int_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_Ref_6Nested_int_x_entry struct {
    active bool
    key *ref_Nested_x
    value int
}

type hashmap_Ref_6Nested_int_x struct {
    buckets map[uint64][]hashmap_Ref_6Nested_int_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_11Ref_6Nested_3int() *hashmap_Ref_6Nested_int_x {
    return &hashmap_Ref_6Nested_int_x{
        buckets: make(map[uint64][]hashmap_Ref_6Nested_int_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_11Ref_6Nested_3int(m *hashmap_Ref_6Nested_int_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_11Ref_6Nested_3int(m *hashmap_Ref_6Nested_int_x, key *ref_Nested_x) (int, bool, int, uint64) {
    if m == nil {
        var zero int
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Nested_r__i_hash(key)
    var bucket []hashmap_Ref_6Nested_int_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_6Nested_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Nested_r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero int
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_11Ref_6Nested_3int(m *hashmap_Ref_6Nested_int_x, key *ref_Nested_x) Option__isize {
    var value int
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_11Ref_6Nested_3int(m, key)
    if ok {
        return Option__isize{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__isize{
        _tag: 0,
    }
}

func hashmap_set__HashMap_11Ref_6Nested_3int(m *hashmap_Ref_6Nested_int_x, key *ref_Nested_x, value int) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Nested_r__i_hash(key)
    var bucket []hashmap_Ref_6Nested_int_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_6Nested_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Nested_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_6Nested_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_6Nested_int_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_Ref_13Array_0_4unit_int_x_entry struct {
    active bool
    key *ref_Array_0_4unit_x
    value int
}

type hashmap_Ref_13Array_0_4unit_int_x struct {
    buckets map[uint64][]hashmap_Ref_13Array_0_4unit_int_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_19Ref_13Array_0_4unit_3int() *hashmap_Ref_13Array_0_4unit_int_x {
    return &hashmap_Ref_13Array_0_4unit_int_x{
        buckets: make(map[uint64][]hashmap_Ref_13Array_0_4unit_int_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_19Ref_13Array_0_4unit_3int(m *hashmap_Ref_13Array_0_4unit_int_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_19Ref_13Array_0_4unit_3int(m *hashmap_Ref_13Array_0_4unit_int_x, key *ref_Array_0_4unit_x) (int, bool, int, uint64) {
    if m == nil {
        var zero int
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l__l__o__q__x3b_0_r__r__i_hash(key)
    var bucket []hashmap_Ref_13Array_0_4unit_int_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_13Array_0_4unit_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l__l__o__q__x3b_0_r__r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero int
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_19Ref_13Array_0_4unit_3int(m *hashmap_Ref_13Array_0_4unit_int_x, key *ref_Array_0_4unit_x) Option__isize {
    var value int
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_19Ref_13Array_0_4unit_3int(m, key)
    if ok {
        return Option__isize{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__isize{
        _tag: 0,
    }
}

func hashmap_set__HashMap_19Ref_13Array_0_4unit_3int(m *hashmap_Ref_13Array_0_4unit_int_x, key *ref_Array_0_4unit_x, value int) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l__l__o__q__x3b_0_r__r__i_hash(key)
    var bucket []hashmap_Ref_13Array_0_4unit_int_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_13Array_0_4unit_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l__l__o__q__x3b_0_r__r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_13Array_0_4unit_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_13Array_0_4unit_int_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_Ref_13Array_3_4unit_int_x_entry struct {
    active bool
    key *ref_Array_3_4unit_x
    value int
}

type hashmap_Ref_13Array_3_4unit_int_x struct {
    buckets map[uint64][]hashmap_Ref_13Array_3_4unit_int_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_19Ref_13Array_3_4unit_3int() *hashmap_Ref_13Array_3_4unit_int_x {
    return &hashmap_Ref_13Array_3_4unit_int_x{
        buckets: make(map[uint64][]hashmap_Ref_13Array_3_4unit_int_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_19Ref_13Array_3_4unit_3int(m *hashmap_Ref_13Array_3_4unit_int_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_19Ref_13Array_3_4unit_3int(m *hashmap_Ref_13Array_3_4unit_int_x, key *ref_Array_3_4unit_x) (int, bool, int, uint64) {
    if m == nil {
        var zero int
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l__l__o__q__x3b_3_r__r__i_hash(key)
    var bucket []hashmap_Ref_13Array_3_4unit_int_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_13Array_3_4unit_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l__l__o__q__x3b_3_r__r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero int
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_19Ref_13Array_3_4unit_3int(m *hashmap_Ref_13Array_3_4unit_int_x, key *ref_Array_3_4unit_x) Option__isize {
    var value int
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_19Ref_13Array_3_4unit_3int(m, key)
    if ok {
        return Option__isize{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__isize{
        _tag: 0,
    }
}

func hashmap_set__HashMap_19Ref_13Array_3_4unit_3int(m *hashmap_Ref_13Array_3_4unit_int_x, key *ref_Array_3_4unit_x, value int) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l__l__o__q__x3b_3_r__r__i_hash(key)
    var bucket []hashmap_Ref_13Array_3_4unit_int_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_13Array_3_4unit_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l__l__o__q__x3b_3_r__r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_13Array_3_4unit_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_13Array_3_4unit_int_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_Ref_4bool_int_x_entry struct {
    active bool
    key *ref_bool_x
    value int
}

type hashmap_Ref_4bool_int_x struct {
    buckets map[uint64][]hashmap_Ref_4bool_int_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_9Ref_4bool_3int() *hashmap_Ref_4bool_int_x {
    return &hashmap_Ref_4bool_int_x{
        buckets: make(map[uint64][]hashmap_Ref_4bool_int_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_9Ref_4bool_3int(m *hashmap_Ref_4bool_int_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_9Ref_4bool_3int(m *hashmap_Ref_4bool_int_x, key *ref_bool_x) (int, bool, int, uint64) {
    if m == nil {
        var zero int
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_bool_r__i_hash(key)
    var bucket []hashmap_Ref_4bool_int_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_4bool_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_bool_r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero int
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_9Ref_4bool_3int(m *hashmap_Ref_4bool_int_x, key *ref_bool_x) Option__isize {
    var value int
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_9Ref_4bool_3int(m, key)
    if ok {
        return Option__isize{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__isize{
        _tag: 0,
    }
}

func hashmap_set__HashMap_9Ref_4bool_3int(m *hashmap_Ref_4bool_int_x, key *ref_bool_x, value int) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_bool_r__i_hash(key)
    var bucket []hashmap_Ref_4bool_int_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_4bool_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_bool_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_4bool_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_4bool_int_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type Tuple2_4unit_5Empty struct {
    _0 struct{}
    _1 Empty
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

type Empty struct {}

type Nested struct {
    value Tuple2_4unit_5Empty
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

func main0() struct{} {
    _goml_m_distinct____T___o__q_(struct{}{})
    var t0 Empty = Empty{}
    distinct__T_Empty(t0)
    var t1 Empty = Empty{}
    var t2 Tuple2_4unit_5Empty = Tuple2_4unit_5Empty{
        _0: struct{}{},
        _1: t1,
    }
    var t3 Nested = Nested{
        value: t2,
    }
    distinct__T_Nested(t3)
    var empty__0 [0]struct{} = [0]struct{}{}
    _goml_m_distinct____T___l__o__q__x3b_0_r_(empty__0)
    var t4 [3]struct{} = [3]struct{}{struct{}{}, struct{}{}, struct{}{}}
    _goml_m_distinct____T___l__o__q__x3b_3_r_(t4)
    distinct__T_bool(false)
    return struct{}{}
}

func _goml_m_distinct____T___o__q_(value__0 struct{}) struct{} {
    var refs__0 *_goml_vec_Ref_4unit
    var inline20 *_goml_vec_Ref_4unit = vec_new__Vec_9Ref_4unit()
    refs__0 = inline20
    var keys__0 *hashmap_Ref_4unit_int_x
    var inline19 *hashmap_Ref_4unit_int_x = hashmap_new__HashMap_9Ref_4unit_3int()
    keys__0 = inline19
    var for_index0 int = 0
    var for_limit0 int = 128
    Loop_loop0:
    for {
        var t14 bool = for_index0 < for_limit0
        if t14 {
            var for_item0 int = for_index0
            var t15 int = for_index0 + 1
            for_index0 = t15
            var reference__0 *ref_unit_x
            var inline18 *ref_unit_x = ref__Ref_4unit(value__0)
            reference__0 = inline18
            vec_push__Vec_9Ref_4unit(refs__0, reference__0)
            hashmap_set__HashMap_9Ref_4unit_3int(keys__0, reference__0, for_item0)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline15 int = hashmap_len__HashMap_9Ref_4unit_3int(keys__0)
    t0 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline13)
    var t1 *ref_unit_x = vec_get__Vec_9Ref_4unit(refs__0, 0)
    var t2 *ref_unit_x = vec_get__Vec_9Ref_4unit(refs__0, 1)
    var t3 bool = ptr_eq__Ref_4unit(t1, t2)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline11)
    var t4 *ref_unit_x = vec_get__Vec_9Ref_4unit(refs__0, 0)
    var t5 *ref_unit_x = vec_get__Vec_9Ref_4unit(refs__0, 0)
    var t6 bool = ptr_eq__Ref_4unit(t4, t5)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline9)
    var t7 *ref_unit_x = vec_get__Vec_9Ref_4unit(refs__0, 127)
    var t8 Option__isize = hashmap_get__HashMap_9Ref_4unit_3int(keys__0, t7)
    var t9 int
    var inline7 int = -1
    switch t8._tag {
    case 0:
        t9 = inline7
    case 1:
        var inline8 int = t8._p0
        t9 = inline8
    default:
        panic("non-exhaustive match")
    }
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline5)
    var t10 *ref_unit_x = vec_get__Vec_9Ref_4unit(refs__0, 0)
    ref_set__Ref_4unit(t10, value__0)
    var t11 *ref_unit_x = vec_get__Vec_9Ref_4unit(refs__0, 0)
    var t12 Option__isize = hashmap_get__HashMap_9Ref_4unit_3int(keys__0, t11)
    var t13 int
    var inline2 int = -1
    switch t12._tag {
    case 0:
        t13 = inline2
    case 1:
        var inline3 int = t12._p0
        t13 = inline3
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t13)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func distinct__T_Empty(value__0 Empty) struct{} {
    var refs__0 *_goml_vec_Ref_5Empty
    var inline20 *_goml_vec_Ref_5Empty = vec_new__Vec_10Ref_5Empty()
    refs__0 = inline20
    var keys__0 *hashmap_Ref_5Empty_int_x
    var inline19 *hashmap_Ref_5Empty_int_x = hashmap_new__HashMap_10Ref_5Empty_3int()
    keys__0 = inline19
    var for_index0 int = 0
    var for_limit0 int = 128
    Loop_loop0:
    for {
        var t14 bool = for_index0 < for_limit0
        if t14 {
            var for_item0 int = for_index0
            var t15 int = for_index0 + 1
            for_index0 = t15
            var reference__0 *ref_Empty_x
            var inline18 *ref_Empty_x = ref__Ref_5Empty(value__0)
            reference__0 = inline18
            vec_push__Vec_10Ref_5Empty(refs__0, reference__0)
            hashmap_set__HashMap_10Ref_5Empty_3int(keys__0, reference__0, for_item0)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline15 int = hashmap_len__HashMap_10Ref_5Empty_3int(keys__0)
    t0 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline13)
    var t1 *ref_Empty_x = vec_get__Vec_10Ref_5Empty(refs__0, 0)
    var t2 *ref_Empty_x = vec_get__Vec_10Ref_5Empty(refs__0, 1)
    var t3 bool = ptr_eq__Ref_5Empty(t1, t2)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline11)
    var t4 *ref_Empty_x = vec_get__Vec_10Ref_5Empty(refs__0, 0)
    var t5 *ref_Empty_x = vec_get__Vec_10Ref_5Empty(refs__0, 0)
    var t6 bool = ptr_eq__Ref_5Empty(t4, t5)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline9)
    var t7 *ref_Empty_x = vec_get__Vec_10Ref_5Empty(refs__0, 127)
    var t8 Option__isize = hashmap_get__HashMap_10Ref_5Empty_3int(keys__0, t7)
    var t9 int
    var inline7 int = -1
    switch t8._tag {
    case 0:
        t9 = inline7
    case 1:
        var inline8 int = t8._p0
        t9 = inline8
    default:
        panic("non-exhaustive match")
    }
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline5)
    var t10 *ref_Empty_x = vec_get__Vec_10Ref_5Empty(refs__0, 0)
    ref_set__Ref_5Empty(t10, value__0)
    var t11 *ref_Empty_x = vec_get__Vec_10Ref_5Empty(refs__0, 0)
    var t12 Option__isize = hashmap_get__HashMap_10Ref_5Empty_3int(keys__0, t11)
    var t13 int
    var inline2 int = -1
    switch t12._tag {
    case 0:
        t13 = inline2
    case 1:
        var inline3 int = t12._p0
        t13 = inline3
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t13)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func distinct__T_Nested(value__0 Nested) struct{} {
    var refs__0 *_goml_vec_Ref_6Nested
    var inline20 *_goml_vec_Ref_6Nested = vec_new__Vec_11Ref_6Nested()
    refs__0 = inline20
    var keys__0 *hashmap_Ref_6Nested_int_x
    var inline19 *hashmap_Ref_6Nested_int_x = hashmap_new__HashMap_11Ref_6Nested_3int()
    keys__0 = inline19
    var for_index0 int = 0
    var for_limit0 int = 128
    Loop_loop0:
    for {
        var t14 bool = for_index0 < for_limit0
        if t14 {
            var for_item0 int = for_index0
            var t15 int = for_index0 + 1
            for_index0 = t15
            var reference__0 *ref_Nested_x
            var inline18 *ref_Nested_x = ref__Ref_6Nested(value__0)
            reference__0 = inline18
            vec_push__Vec_11Ref_6Nested(refs__0, reference__0)
            hashmap_set__HashMap_11Ref_6Nested_3int(keys__0, reference__0, for_item0)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline15 int = hashmap_len__HashMap_11Ref_6Nested_3int(keys__0)
    t0 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline13)
    var t1 *ref_Nested_x = vec_get__Vec_11Ref_6Nested(refs__0, 0)
    var t2 *ref_Nested_x = vec_get__Vec_11Ref_6Nested(refs__0, 1)
    var t3 bool = ptr_eq__Ref_6Nested(t1, t2)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline11)
    var t4 *ref_Nested_x = vec_get__Vec_11Ref_6Nested(refs__0, 0)
    var t5 *ref_Nested_x = vec_get__Vec_11Ref_6Nested(refs__0, 0)
    var t6 bool = ptr_eq__Ref_6Nested(t4, t5)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline9)
    var t7 *ref_Nested_x = vec_get__Vec_11Ref_6Nested(refs__0, 127)
    var t8 Option__isize = hashmap_get__HashMap_11Ref_6Nested_3int(keys__0, t7)
    var t9 int
    var inline7 int = -1
    switch t8._tag {
    case 0:
        t9 = inline7
    case 1:
        var inline8 int = t8._p0
        t9 = inline8
    default:
        panic("non-exhaustive match")
    }
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline5)
    var t10 *ref_Nested_x = vec_get__Vec_11Ref_6Nested(refs__0, 0)
    ref_set__Ref_6Nested(t10, value__0)
    var t11 *ref_Nested_x = vec_get__Vec_11Ref_6Nested(refs__0, 0)
    var t12 Option__isize = hashmap_get__HashMap_11Ref_6Nested_3int(keys__0, t11)
    var t13 int
    var inline2 int = -1
    switch t12._tag {
    case 0:
        t13 = inline2
    case 1:
        var inline3 int = t12._p0
        t13 = inline3
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t13)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_distinct____T___l__o__q__x3b_0_r_(value__0 [0]struct{}) struct{} {
    var refs__0 *_goml_vec_Ref_13Array_0_4unit
    var inline20 *_goml_vec_Ref_13Array_0_4unit = vec_new__Vec_19Ref_13Array_0_4unit()
    refs__0 = inline20
    var keys__0 *hashmap_Ref_13Array_0_4unit_int_x
    var inline19 *hashmap_Ref_13Array_0_4unit_int_x = hashmap_new__HashMap_19Ref_13Array_0_4unit_3int()
    keys__0 = inline19
    var for_index0 int = 0
    var for_limit0 int = 128
    Loop_loop0:
    for {
        var t14 bool = for_index0 < for_limit0
        if t14 {
            var for_item0 int = for_index0
            var t15 int = for_index0 + 1
            for_index0 = t15
            var reference__0 *ref_Array_0_4unit_x
            var inline18 *ref_Array_0_4unit_x = ref__Ref_13Array_0_4unit(value__0)
            reference__0 = inline18
            vec_push__Vec_19Ref_13Array_0_4unit(refs__0, reference__0)
            hashmap_set__HashMap_19Ref_13Array_0_4unit_3int(keys__0, reference__0, for_item0)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline15 int = hashmap_len__HashMap_19Ref_13Array_0_4unit_3int(keys__0)
    t0 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline13)
    var t1 *ref_Array_0_4unit_x = vec_get__Vec_19Ref_13Array_0_4unit(refs__0, 0)
    var t2 *ref_Array_0_4unit_x = vec_get__Vec_19Ref_13Array_0_4unit(refs__0, 1)
    var t3 bool = ptr_eq__Ref_13Array_0_4unit(t1, t2)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline11)
    var t4 *ref_Array_0_4unit_x = vec_get__Vec_19Ref_13Array_0_4unit(refs__0, 0)
    var t5 *ref_Array_0_4unit_x = vec_get__Vec_19Ref_13Array_0_4unit(refs__0, 0)
    var t6 bool = ptr_eq__Ref_13Array_0_4unit(t4, t5)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline9)
    var t7 *ref_Array_0_4unit_x = vec_get__Vec_19Ref_13Array_0_4unit(refs__0, 127)
    var t8 Option__isize = hashmap_get__HashMap_19Ref_13Array_0_4unit_3int(keys__0, t7)
    var t9 int
    var inline7 int = -1
    switch t8._tag {
    case 0:
        t9 = inline7
    case 1:
        var inline8 int = t8._p0
        t9 = inline8
    default:
        panic("non-exhaustive match")
    }
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline5)
    var t10 *ref_Array_0_4unit_x = vec_get__Vec_19Ref_13Array_0_4unit(refs__0, 0)
    ref_set__Ref_13Array_0_4unit(t10, value__0)
    var t11 *ref_Array_0_4unit_x = vec_get__Vec_19Ref_13Array_0_4unit(refs__0, 0)
    var t12 Option__isize = hashmap_get__HashMap_19Ref_13Array_0_4unit_3int(keys__0, t11)
    var t13 int
    var inline2 int = -1
    switch t12._tag {
    case 0:
        t13 = inline2
    case 1:
        var inline3 int = t12._p0
        t13 = inline3
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t13)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_distinct____T___l__o__q__x3b_3_r_(value__0 [3]struct{}) struct{} {
    var refs__0 *_goml_vec_Ref_13Array_3_4unit
    var inline20 *_goml_vec_Ref_13Array_3_4unit = vec_new__Vec_19Ref_13Array_3_4unit()
    refs__0 = inline20
    var keys__0 *hashmap_Ref_13Array_3_4unit_int_x
    var inline19 *hashmap_Ref_13Array_3_4unit_int_x = hashmap_new__HashMap_19Ref_13Array_3_4unit_3int()
    keys__0 = inline19
    var for_index0 int = 0
    var for_limit0 int = 128
    Loop_loop0:
    for {
        var t14 bool = for_index0 < for_limit0
        if t14 {
            var for_item0 int = for_index0
            var t15 int = for_index0 + 1
            for_index0 = t15
            var reference__0 *ref_Array_3_4unit_x
            var inline18 *ref_Array_3_4unit_x = ref__Ref_13Array_3_4unit(value__0)
            reference__0 = inline18
            vec_push__Vec_19Ref_13Array_3_4unit(refs__0, reference__0)
            hashmap_set__HashMap_19Ref_13Array_3_4unit_3int(keys__0, reference__0, for_item0)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline15 int = hashmap_len__HashMap_19Ref_13Array_3_4unit_3int(keys__0)
    t0 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline13)
    var t1 *ref_Array_3_4unit_x = vec_get__Vec_19Ref_13Array_3_4unit(refs__0, 0)
    var t2 *ref_Array_3_4unit_x = vec_get__Vec_19Ref_13Array_3_4unit(refs__0, 1)
    var t3 bool = ptr_eq__Ref_13Array_3_4unit(t1, t2)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline11)
    var t4 *ref_Array_3_4unit_x = vec_get__Vec_19Ref_13Array_3_4unit(refs__0, 0)
    var t5 *ref_Array_3_4unit_x = vec_get__Vec_19Ref_13Array_3_4unit(refs__0, 0)
    var t6 bool = ptr_eq__Ref_13Array_3_4unit(t4, t5)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline9)
    var t7 *ref_Array_3_4unit_x = vec_get__Vec_19Ref_13Array_3_4unit(refs__0, 127)
    var t8 Option__isize = hashmap_get__HashMap_19Ref_13Array_3_4unit_3int(keys__0, t7)
    var t9 int
    var inline7 int = -1
    switch t8._tag {
    case 0:
        t9 = inline7
    case 1:
        var inline8 int = t8._p0
        t9 = inline8
    default:
        panic("non-exhaustive match")
    }
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline5)
    var t10 *ref_Array_3_4unit_x = vec_get__Vec_19Ref_13Array_3_4unit(refs__0, 0)
    ref_set__Ref_13Array_3_4unit(t10, value__0)
    var t11 *ref_Array_3_4unit_x = vec_get__Vec_19Ref_13Array_3_4unit(refs__0, 0)
    var t12 Option__isize = hashmap_get__HashMap_19Ref_13Array_3_4unit_3int(keys__0, t11)
    var t13 int
    var inline2 int = -1
    switch t12._tag {
    case 0:
        t13 = inline2
    case 1:
        var inline3 int = t12._p0
        t13 = inline3
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t13)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func distinct__T_bool(value__0 bool) struct{} {
    var refs__0 *_goml_vec_Ref_4bool
    var inline20 *_goml_vec_Ref_4bool = vec_new__Vec_9Ref_4bool()
    refs__0 = inline20
    var keys__0 *hashmap_Ref_4bool_int_x
    var inline19 *hashmap_Ref_4bool_int_x = hashmap_new__HashMap_9Ref_4bool_3int()
    keys__0 = inline19
    var for_index0 int = 0
    var for_limit0 int = 128
    Loop_loop0:
    for {
        var t14 bool = for_index0 < for_limit0
        if t14 {
            var for_item0 int = for_index0
            var t15 int = for_index0 + 1
            for_index0 = t15
            var reference__0 *ref_bool_x
            var inline18 *ref_bool_x = ref__Ref_4bool(value__0)
            reference__0 = inline18
            vec_push__Vec_9Ref_4bool(refs__0, reference__0)
            hashmap_set__HashMap_9Ref_4bool_3int(keys__0, reference__0, for_item0)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline15 int = hashmap_len__HashMap_9Ref_4bool_3int(keys__0)
    t0 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline13)
    var t1 *ref_bool_x = vec_get__Vec_9Ref_4bool(refs__0, 0)
    var t2 *ref_bool_x = vec_get__Vec_9Ref_4bool(refs__0, 1)
    var t3 bool = ptr_eq__Ref_4bool(t1, t2)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline11)
    var t4 *ref_bool_x = vec_get__Vec_9Ref_4bool(refs__0, 0)
    var t5 *ref_bool_x = vec_get__Vec_9Ref_4bool(refs__0, 0)
    var t6 bool = ptr_eq__Ref_4bool(t4, t5)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline9)
    var t7 *ref_bool_x = vec_get__Vec_9Ref_4bool(refs__0, 127)
    var t8 Option__isize = hashmap_get__HashMap_9Ref_4bool_3int(keys__0, t7)
    var t9 int
    var inline7 int = -1
    switch t8._tag {
    case 0:
        t9 = inline7
    case 1:
        var inline8 int = t8._p0
        t9 = inline8
    default:
        panic("non-exhaustive match")
    }
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline5)
    var t10 *ref_bool_x = vec_get__Vec_9Ref_4bool(refs__0, 0)
    ref_set__Ref_4bool(t10, value__0)
    var t11 *ref_bool_x = vec_get__Vec_9Ref_4bool(refs__0, 0)
    var t12 Option__isize = hashmap_get__HashMap_9Ref_4bool_3int(keys__0, t11)
    var t13 int
    var inline2 int = -1
    switch t12._tag {
    case 0:
        t13 = inline2
    case 1:
        var inline3 int = t12._p0
        t13 = inline3
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t13)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
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

func _goml_m_trait__impl_i_PartialEq_i_Ref_l__o__q__r__i_eq(self__0 *ref_unit_x, other__0 *ref_unit_x) bool {
    var t0 bool = ptr_eq__Ref_4unit(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l__o__q__r__i_hash(self__0 *ref_unit_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_4unit(self__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l_Empty_r__i_eq(self__0 *ref_Empty_x, other__0 *ref_Empty_x) bool {
    var t0 bool = ptr_eq__Ref_5Empty(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l_Empty_r__i_hash(self__0 *ref_Empty_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_5Empty(self__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l_Nested_r__i_eq(self__0 *ref_Nested_x, other__0 *ref_Nested_x) bool {
    var t0 bool = ptr_eq__Ref_6Nested(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l_Nested_r__i_hash(self__0 *ref_Nested_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_6Nested(self__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l__l__o__q__x3b_0_r__r__i_eq(self__0 *ref_Array_0_4unit_x, other__0 *ref_Array_0_4unit_x) bool {
    var t0 bool = ptr_eq__Ref_13Array_0_4unit(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l__l__o__q__x3b_0_r__r__i_hash(self__0 *ref_Array_0_4unit_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_13Array_0_4unit(self__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l__l__o__q__x3b_3_r__r__i_eq(self__0 *ref_Array_3_4unit_x, other__0 *ref_Array_3_4unit_x) bool {
    var t0 bool = ptr_eq__Ref_13Array_3_4unit(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l__l__o__q__x3b_3_r__r__i_hash(self__0 *ref_Array_3_4unit_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_13Array_3_4unit(self__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l_bool_r__i_eq(self__0 *ref_bool_x, other__0 *ref_bool_x) bool {
    var t0 bool = ptr_eq__Ref_4bool(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l_bool_r__i_hash(self__0 *ref_bool_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_4bool(self__0)
    return t0
}

func main() {
    main0()
}
