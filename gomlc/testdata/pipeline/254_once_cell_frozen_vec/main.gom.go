package main

import (
    _goml_os "os"
    _goml_sync "sync"
    _goml_runtime_pkg "runtime"
)

func _goml_once_cell_goroutine_id() uint64 {
    var buffer []uint8 = make([]uint8, 64)
    var length int = _goml_runtime_pkg.Stack(buffer, false)
    var index int = 10
    var result uint64 = 0
    for {
        if index >= length {
            break
        }
        if buffer[index] < 48 || buffer[index] > 57 {
            break
        }
        result = result * 10 + uint64(buffer[index] - 48)
        index = index + 1
    }
    return result
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

func once_cell_new__OnceCell__FrozenVec__isize() *OnceCell__FrozenVec__isize {
    var cell *OnceCell__FrozenVec__isize = &OnceCell__FrozenVec__isize{}
    cell.cond = _goml_sync.NewCond(&cell.mutex)
    return cell
}

func once_cell_get_or_init__OnceCell__FrozenVec__isize(cell *OnceCell__FrozenVec__isize, init func() FrozenVec__isize) FrozenVec__isize {
    cell.mutex.Lock()
    for {
        if cell.state == 2 {
            cell.mutex.Unlock()
            return cell.value
        }
        var goroutine uint64 = _goml_once_cell_goroutine_id()
        if cell.state == 1 {
            if cell.owner == goroutine {
                cell.mutex.Unlock()
                panic("recursive OnceCell initialization: " + cell.name)
            }
            cell.cond.Wait()
            continue
        }
        cell.state = 1
        cell.owner = goroutine
        cell.mutex.Unlock()
        var completed bool = false
        defer func() {
            if !completed {
                cell.mutex.Lock()
                cell.state = 0
                cell.owner = 0
                cell.cond.Broadcast()
                cell.mutex.Unlock()
            }
        }()
        var initialized FrozenVec__isize = init()
        cell.mutex.Lock()
        cell.value = initialized
        cell.state = 2
        cell.owner = 0
        cell.cond.Broadcast()
        completed = true
        cell.mutex.Unlock()
        return initialized
    }
}

type _goml_vec_int struct {
    items []int
}

func vec_with_capacity__Vec_3int(capacity int) *_goml_vec_int {
    return &_goml_vec_int{
        items: make([]int, 0, capacity),
    }
}

func vec_push__Vec_3int(vec *_goml_vec_int, elem int) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_3int(vec *_goml_vec_int, index int) int {
    return vec.items[index]
}

func vec_set__Vec_3int(vec *_goml_vec_int, index int, value int) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_3int(vec *_goml_vec_int) int {
    return int(len(vec.items))
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

type closure_env_values_0 struct {}

type FrozenVec__isize struct {
    values *_goml_vec_int
}

type OnceCell__FrozenVec__isize struct {
    mutex _goml_sync.Mutex
    cond *_goml_sync.Cond
    state int
    owner uint64
    value FrozenVec__isize
    name string
}

type Ordering uint8

var VALUES *OnceCell__FrozenVec__isize = func() *OnceCell__FrozenVec__isize {
    var cell *OnceCell__FrozenVec__isize = once_cell_new__OnceCell__FrozenVec__isize()
    cell.name = "VALUES"
    return cell
}()

func values() FrozenVec__isize {
    var t0 closure_env_values_0 = closure_env_values_0{}
    var t1 func() FrozenVec__isize = func() FrozenVec__isize {
        return _goml_m_inherent_i_closure__env__values__0_i_closure__env__values__0_i_apply(t0)
    }
    var inline0 FrozenVec__isize = once_cell_get_or_init__OnceCell__FrozenVec__isize(VALUES, t1)
    return inline0
}

func main0() struct{} {
    var frozen__0 FrozenVec__isize
    var inline25 closure_env_values_0 = closure_env_values_0{}
    var inline26 func() FrozenVec__isize = func() FrozenVec__isize {
        return _goml_m_inherent_i_closure__env__values__0_i_closure__env__values__0_i_apply(inline25)
    }
    var inline27 FrozenVec__isize = _goml_m_inherent_i_OnceCell_i__hc919cf300d97b8e399cb9d4664fdfe6e_nVec_l_isize_r_(VALUES, inline26)
    frozen__0 = inline27
    var copy__0 *_goml_vec_int
    var inline23 *_goml_vec_int = frozen__0.values
    var inline24 *_goml_vec_int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_copy____T__isize(inline23)
    copy__0 = inline24
    var inline20 int = 0
    var inline21 int = 9
    vec_set__Vec_3int(copy__0, inline20, inline21)
    var t0 int
    var inline17 int = 0
    var inline18 *_goml_vec_int = frozen__0.values
    var inline19 int = vec_get__Vec_3int(inline18, inline17)
    t0 = inline19
    var t1 string
    var inline16 string = __goml_builtin_int_to_string(t0)
    t1 = inline16
    var inline14 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline14)
    var t2 int
    var inline12 int = 0
    var inline13 int = vec_get__Vec_3int(copy__0, inline12)
    t2 = inline13
    var t3 string
    var inline11 string = __goml_builtin_int_to_string(t2)
    t3 = inline11
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline9)
    var t4 FrozenVec__isize
    var inline6 closure_env_values_0 = closure_env_values_0{}
    var inline7 func() FrozenVec__isize = func() FrozenVec__isize {
        return _goml_m_inherent_i_closure__env__values__0_i_closure__env__values__0_i_apply(inline6)
    }
    var inline8 FrozenVec__isize = _goml_m_inherent_i_OnceCell_i__hc919cf300d97b8e399cb9d4664fdfe6e_nVec_l_isize_r_(VALUES, inline7)
    t4 = inline8
    var t5 int
    var inline3 int = 1
    var inline4 *_goml_vec_int = t4.values
    var inline5 int = vec_get__Vec_3int(inline4, inline3)
    t5 = inline5
    var t6 string
    var inline2 string = __goml_builtin_int_to_string(t5)
    t6 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t6)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_OnceCell_i__hc919cf300d97b8e399cb9d4664fdfe6e_nVec_l_isize_r_(self__0 *OnceCell__FrozenVec__isize, init__0 func() FrozenVec__isize) FrozenVec__isize {
    var t0 FrozenVec__isize = once_cell_get_or_init__OnceCell__FrozenVec__isize(self__0, init__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_copy____T__isize(self__0 *_goml_vec_int) *_goml_vec_int {
    var t0 int
    var inline3 int = vec_len__Vec_3int(self__0)
    t0 = inline3
    var result__0 *_goml_vec_int
    var inline2 *_goml_vec_int = vec_with_capacity__Vec_3int(t0)
    result__0 = inline2
    var index__0 int = 0
    Loop_loop0:
    for {
        var t1 int
        var inline1 int = vec_len__Vec_3int(self__0)
        t1 = inline1
        var t2 bool = index__0 < t1
        if t2 {
            var t3 int = vec_get__Vec_3int(self__0, index__0)
            vec_push__Vec_3int(result__0, t3)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t4 int = compound_old0 + compound_value0
            index__0 = t4
            continue
        } else {
            break Loop_loop0
        }
    }
    return result__0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func _goml_m_inherent_i_closure__env__values__0_i_closure__env__values__0_i_apply(env0 closure_env_values_0) FrozenVec__isize {
    var t0 [3]int = [3]int{1, 2, 3}
    var t1 *_goml_vec_int = func(values [3]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [3]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t0)
    var inline0 *_goml_vec_int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_copy____T__isize(t1)
    var inline1 FrozenVec__isize = FrozenVec__isize{
        values: inline0,
    }
    return inline1
}

func main() {
    main0()
}
