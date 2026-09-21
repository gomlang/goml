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

func _goml_runtime_core_bool_to_string(x bool) string {
    if x {
        return "true"
    } else {
        return "false"
    }
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

func once_cell_new__OnceCell__isize() *OnceCell__isize {
    var cell *OnceCell__isize = &OnceCell__isize{}
    cell.cond = _goml_sync.NewCond(&cell.mutex)
    return cell
}

func once_cell_get_or_init__OnceCell__isize(cell *OnceCell__isize, init func() int) int {
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
        var initialized int = init()
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

type _goml_vec_uint32 struct {
    items []uint32
}

type Tuple2_3int_4bool struct {
    _0 int
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

type closure_env_main_0 struct {}

type closure_env_main_1 struct {
    results_0 chan int
}

type closure_env_main_2 struct {}

type closure_env_main_3 struct {
    results_0 chan int
}

type OnceCell__isize struct {
    mutex _goml_sync.Mutex
    cond *_goml_sync.Cond
    state int
    owner uint64
    value int
    name string
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

var VALUE *OnceCell__isize = func() *OnceCell__isize {
    var cell *OnceCell__isize = once_cell_new__OnceCell__isize()
    cell.name = "VALUE"
    return cell
}()

func main0() struct{} {
    var results__0 chan int
    var inline15 int = 2
    var inline16 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline15)
    results__0 = inline16
    var t0 closure_env_main_1 = closure_env_main_1{
        results_0: results__0,
    }
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t0)
    }
    go t1()
    var t2 closure_env_main_3 = closure_env_main_3{
        results_0: results__0,
    }
    var t3 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(t2)
    }
    go t3()
    var t4 Option__isize
    var inline11 Tuple2_3int_4bool = func(p0 chan int) Tuple2_3int_4bool {
        var value int
        var ok bool
        value, ok = <-p0
        return Tuple2_3int_4bool{
            _0: value,
            _1: ok,
        }
    }(results__0)
    var inline12 int = inline11._0
    var inline13 bool = inline11._1
    if inline13 {
        var inline14 Option__isize = Option__isize{
            _p0: inline12,
            _tag: 1,
        }
        t4 = inline14
    } else {
        t4 = Option__isize{
            _tag: 0,
        }
    }
    var first__0 int
    var inline9 int = 0
    switch t4._tag {
    case 0:
        first__0 = inline9
    case 1:
        var inline10 int = t4._p0
        first__0 = inline10
    default:
        panic("non-exhaustive match")
    }
    var t5 Option__isize
    var inline5 Tuple2_3int_4bool = func(p0 chan int) Tuple2_3int_4bool {
        var value int
        var ok bool
        value, ok = <-p0
        return Tuple2_3int_4bool{
            _0: value,
            _1: ok,
        }
    }(results__0)
    var inline6 int = inline5._0
    var inline7 bool = inline5._1
    if inline7 {
        var inline8 Option__isize = Option__isize{
            _p0: inline6,
            _tag: 1,
        }
        t5 = inline8
    } else {
        t5 = Option__isize{
            _tag: 0,
        }
    }
    var second__0 int
    var inline3 int = 0
    switch t5._tag {
    case 0:
        second__0 = inline3
    case 1:
        var inline4 int = t5._p0
        second__0 = inline4
    default:
        panic("non-exhaustive match")
    }
    var t6 bool = first__0 == second__0
    var t7 string
    var inline2 string = _goml_runtime_core_bool_to_string(t6)
    t7 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0) int {
    return 41
}

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1) struct{} {
    var results__0 chan int = env0.results_0
    var t0 closure_env_main_0 = closure_env_main_0{}
    var t1 func() int = func() int {
        return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t0)
    }
    var t2 int
    var inline1 int = once_cell_get_or_init__OnceCell__isize(VALUE, t1)
    t2 = inline1
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(results__0, t2)
    return struct{}{}
}

func _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(env0 closure_env_main_2) int {
    return 42
}

func _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(env0 closure_env_main_3) struct{} {
    var results__0 chan int = env0.results_0
    var t0 closure_env_main_2 = closure_env_main_2{}
    var t1 func() int = func() int {
        return _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(t0)
    }
    var t2 int
    var inline1 int = once_cell_get_or_init__OnceCell__isize(VALUE, t1)
    t2 = inline1
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(results__0, t2)
    return struct{}{}
}

func main() {
    main0()
}
