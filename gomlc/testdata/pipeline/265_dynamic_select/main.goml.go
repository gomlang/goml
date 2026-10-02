package main

import (
    _goml_os "os"
    _goml_reflect "reflect"
)

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

type _goml_vec__goml_m_std_p_channel_p_Operation____isize struct {
    items []_goml_m_std_p_channel_p_Operation____isize
}

func vec_new___goml_m_Vec__30std_p_channel_p_Operation____isize() *_goml_vec__goml_m_std_p_channel_p_Operation____isize {
    return &_goml_vec__goml_m_std_p_channel_p_Operation____isize{
        items: nil,
    }
}

func vec_len___goml_m_Vec__30std_p_channel_p_Operation____isize(vec *_goml_vec__goml_m_std_p_channel_p_Operation____isize) int {
    return int(len(vec.items))
}

type _goml_vec_int struct {
    items []int
}

func vec_new__Vec_3int() *_goml_vec_int {
    return &_goml_vec_int{
        items: nil,
    }
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

func vec_len__Vec_3int(vec *_goml_vec_int) int {
    return int(len(vec.items))
}

type _goml_vec_Receiver_3int struct {
    items []<-chan int
}

func vec_new__Vec_13Receiver_3int() *_goml_vec_Receiver_3int {
    return &_goml_vec_Receiver_3int{
        items: nil,
    }
}

func vec_push__Vec_13Receiver_3int(vec *_goml_vec_Receiver_3int, elem <-chan int) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_len__Vec_13Receiver_3int(vec *_goml_vec_Receiver_3int) int {
    return int(len(vec.items))
}

type _goml_vec_Sender_3int struct {
    items []chan<- int
}

func vec_new__Vec_11Sender_3int() *_goml_vec_Sender_3int {
    return &_goml_vec_Sender_3int{
        items: nil,
    }
}

func vec_push__Vec_11Sender_3int(vec *_goml_vec_Sender_3int, elem chan<- int) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_len__Vec_11Sender_3int(vec *_goml_vec_Sender_3int) int {
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

type Tuple4_10Slice_3int_21Slice_13Receiver_3int_19Slice_11Sender_3int_10Slice_3int struct {
    _0 []int
    _1 []<-chan int
    _2 []chan<- int
    _3 []int
}

type Tuple5_3int_3int_3int_4bool_4bool struct {
    _0 int
    _1 int
    _2 int
    _3 bool
    _4 bool
}

type Tuple2_3int_4bool struct {
    _0 int
    _1 bool
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

type _goml_m_std_p_channel_p_PreparedOperations____isize struct {
    kinds *_goml_vec_int
    receivers *_goml_vec_Receiver_3int
    senders *_goml_vec_Sender_3int
    values *_goml_vec_int
}

type Ordering uint8

type _goml_m_std_p_channel_p_SelectError struct{}

type _goml_m_std_p_channel_p_Selection____isize interface {
    is_goml_m_std_p_channel_p_Selection____isize()
}

type Received struct {
    _0 int
    _1 Option__isize
}

func (_ Received) is_goml_m_std_p_channel_p_Selection____isize() {}

type Sent struct {
    _0 int
}

func (_ Sent) is_goml_m_std_p_channel_p_Selection____isize() {}

type Option__isize struct {
    _p0 int
    _tag uint8
}

type _goml_m_std_p_channel_p_Operation____isize interface {
    is_goml_m_std_p_channel_p_Operation____isize()
}

type Receive struct {
    _0 <-chan int
}

func (_ Receive) is_goml_m_std_p_channel_p_Operation____isize() {}

type Send struct {
    _0 chan<- int
    _1 int
}

func (_ Send) is_goml_m_std_p_channel_p_Operation____isize() {}

type _goml_m_Option____std_p_channel_p_Selection____isize struct {
    _p0 _goml_m_std_p_channel_p_Selection____isize
    _tag uint8
}

type _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError struct {
    _p0 _goml_m_Option____std_p_channel_p_Selection____isize
    _tag uint8
}

type _goml_m_Result____std_p_channe_h5400d2fb6c865f5b36aaa9b0f9e787f7_l_p_SelectError struct {
    _p0 _goml_m_std_p_channel_p_Selection____isize
    _tag uint8
}

func print_selection(value__0 _goml_m_std_p_channel_p_Selection____isize) struct{} {
    switch value__0.(type) {
    case Received:
        var x0 int = value__0.(Received)._0
        var x1 Option__isize = value__0.(Received)._1
        var inline4 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x0)
        _goml_runtime_core_string_println(inline4)
        var t0 int
        var inline2 int = -1
        switch x1._tag {
        case 0:
            t0 = inline2
        case 1:
            var inline3 int = x1._p0
            t0 = inline3
        default:
            panic("non-exhaustive match")
        }
        var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    case Sent:
        var x2 int = value__0.(Sent)._0
        var inline6 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x2)
        _goml_runtime_core_string_println(inline6)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var first__0 chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(1)
    var second__0 chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(1)
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_send____T__isize(first__0, 10)
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_send____T__isize(second__0, 20)
    var t0 <-chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T__isize(first__0)
    var t1 _goml_m_std_p_channel_p_Operation____isize = Receive{
        _0: t0,
    }
    var t2 <-chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T__isize(second__0)
    var t3 _goml_m_std_p_channel_p_Operation____isize = Receive{
        _0: t2,
    }
    var t4 [2]_goml_m_std_p_channel_p_Operation____isize = [2]_goml_m_std_p_channel_p_Operation____isize{t1, t3}
    var operations__0 *_goml_vec__goml_m_std_p_channel_p_Operation____isize = func(values [2]_goml_m_std_p_channel_p_Operation____isize) *_goml_vec__goml_m_std_p_channel_p_Operation____isize {
        var storage struct {
            vector _goml_vec__goml_m_std_p_channel_p_Operation____isize
            values [2]_goml_m_std_p_channel_p_Operation____isize
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t4)
    var t5 int = _goml_m_inherent_i_Vec_i_Vec_l_h6161506801e1e3b7949d66d7b76c6b61_tion_l_isize_r_(operations__0)
    var t6 []_goml_m_std_p_channel_p_Operation____isize = operations__0.items[0:t5]
    var mtmp0 _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError = _goml_m_std_p_channel_p_try__select__priority____T__isize(t6)
    switch mtmp0._tag {
    case 0:
        var x5 _goml_m_Option____std_p_channel_p_Selection____isize = mtmp0._p0
        switch x5._tag {
        case 1:
            var x6 _goml_m_std_p_channel_p_Selection____isize = x5._p0
            switch x6.(type) {
            case Received:
                var inline20 int = x6.(Received)._0
                var inline21 Option__isize = x6.(Received)._1
                println__T_isize(inline20)
                var inline23 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(inline21, -1)
                println__T_isize(inline23)
            case Sent:
                var inline25 int = x6.(Sent)._0
                println__T_isize(inline25)
            default:
                panic("non-exhaustive match")
            }
        default:
            var inline27 int = -1
            var inline28 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline27)
            _goml_runtime_core_string_println(inline28)
        }
    default:
        var inline30 int = -1
        var inline31 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline30)
        _goml_runtime_core_string_println(inline31)
    }
    var t7 int = _goml_m_inherent_i_Vec_i_Vec_l_h6161506801e1e3b7949d66d7b76c6b61_tion_l_isize_r_(operations__0)
    var t8 []_goml_m_std_p_channel_p_Operation____isize = operations__0.items[0:t7]
    var mtmp1 _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError = _goml_m_std_p_channel_p_try__select____T__isize(t8)
    switch mtmp1._tag {
    case 0:
        var x3 _goml_m_Option____std_p_channel_p_Selection____isize = mtmp1._p0
        switch x3._tag {
        case 1:
            var x4 _goml_m_std_p_channel_p_Selection____isize = x3._p0
            switch x4.(type) {
            case Received:
                var inline7 int = x4.(Received)._0
                var inline8 Option__isize = x4.(Received)._1
                println__T_isize(inline7)
                var inline10 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(inline8, -1)
                println__T_isize(inline10)
            case Sent:
                var inline12 int = x4.(Sent)._0
                println__T_isize(inline12)
            default:
                panic("non-exhaustive match")
            }
        default:
            var inline14 int = -2
            var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline14)
            _goml_runtime_core_string_println(inline15)
        }
    default:
        var inline17 int = -2
        var inline18 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline17)
        _goml_runtime_core_string_println(inline18)
    }
    var empty__0 chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(0)
    var target__0 chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(1)
    var t9 <-chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T__isize(empty__0)
    var t10 _goml_m_std_p_channel_p_Operation____isize = Receive{
        _0: t9,
    }
    var t11 chan<- int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_sender____T__isize(target__0)
    var t12 _goml_m_std_p_channel_p_Operation____isize = Send{
        _0: t11,
        _1: 30,
    }
    var t13 [2]_goml_m_std_p_channel_p_Operation____isize = [2]_goml_m_std_p_channel_p_Operation____isize{t10, t12}
    var mixed__0 *_goml_vec__goml_m_std_p_channel_p_Operation____isize = func(values [2]_goml_m_std_p_channel_p_Operation____isize) *_goml_vec__goml_m_std_p_channel_p_Operation____isize {
        var storage struct {
            vector _goml_vec__goml_m_std_p_channel_p_Operation____isize
            values [2]_goml_m_std_p_channel_p_Operation____isize
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t13)
    var t14 int = _goml_m_inherent_i_Vec_i_Vec_l_h6161506801e1e3b7949d66d7b76c6b61_tion_l_isize_r_(mixed__0)
    var t15 []_goml_m_std_p_channel_p_Operation____isize = mixed__0.items[0:t14]
    var mtmp2 _goml_m_Result____std_p_channe_h5400d2fb6c865f5b36aaa9b0f9e787f7_l_p_SelectError = _goml_m_std_p_channel_p_select____T__isize(t15)
    switch mtmp2._tag {
    case 0:
        var x2 _goml_m_std_p_channel_p_Selection____isize = mtmp2._p0
        print_selection(x2)
    case 1:
        var inline4 int = -3
        var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline4)
        _goml_runtime_core_string_println(inline5)
    default:
        panic("non-exhaustive match")
    }
    var t16 Option__isize = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_recv____T__isize(target__0)
    var t17 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t16, -4)
    println__T_isize(t17)
    var t18 <-chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T__isize(empty__0)
    var t19 _goml_m_std_p_channel_p_Operation____isize = Receive{
        _0: t18,
    }
    var t20 [1]_goml_m_std_p_channel_p_Operation____isize = [1]_goml_m_std_p_channel_p_Operation____isize{t19}
    var none__0 *_goml_vec__goml_m_std_p_channel_p_Operation____isize = func(values [1]_goml_m_std_p_channel_p_Operation____isize) *_goml_vec__goml_m_std_p_channel_p_Operation____isize {
        var storage struct {
            vector _goml_vec__goml_m_std_p_channel_p_Operation____isize
            values [1]_goml_m_std_p_channel_p_Operation____isize
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t20)
    var t21 int = _goml_m_inherent_i_Vec_i_Vec_l_h6161506801e1e3b7949d66d7b76c6b61_tion_l_isize_r_(none__0)
    var t22 []_goml_m_std_p_channel_p_Operation____isize = none__0.items[0:t21]
    var mtmp3 _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError = _goml_m_std_p_channel_p_try__select____T__isize(t22)
    var jp0 int
    switch mtmp3._tag {
    case 0:
        var x1 _goml_m_Option____std_p_channel_p_Selection____isize = mtmp3._p0
        switch x1._tag {
        case 0:
            jp0 = 40
        default:
            jp0 = -5
        }
    default:
        jp0 = -5
    }
    println__T_isize(jp0)
    var no_operations__0 *_goml_vec__goml_m_std_p_channel_p_Operation____isize
    var inline3 *_goml_vec__goml_m_std_p_channel_p_Operation____isize = vec_new___goml_m_Vec__30std_p_channel_p_Operation____isize()
    no_operations__0 = inline3
    var t23 int
    var inline2 int = vec_len___goml_m_Vec__30std_p_channel_p_Operation____isize(no_operations__0)
    t23 = inline2
    var t24 []_goml_m_std_p_channel_p_Operation____isize = no_operations__0.items[0:t23]
    var mtmp4 _goml_m_Result____std_p_channe_h5400d2fb6c865f5b36aaa9b0f9e787f7_l_p_SelectError = _goml_m_std_p_channel_p_select____T__isize(t24)
    var jp1 int
    switch mtmp4._tag {
    case 1:
        var x0 _goml_m_std_p_channel_p_SelectError = _goml_m_std_p_channel_p_SelectError{}
        switch x0 {
        case _goml_m_std_p_channel_p_SelectError{}:
            jp1 = 50
        default:
            panic("non-exhaustive match")
        }
    default:
        jp1 = -6
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(jp1)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(self__0 Option__isize, fallback__0 int) int {
    switch self__0._tag {
    case 0:
        return fallback__0
    case 1:
        var x0 int = self__0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(capacity__0 int) chan int {
    var t0 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(capacity__0)
    return t0
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_send____T__isize(self__0 chan int, value__0 int) struct{} {
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(self__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_receiver____T__isize(self__0 chan int) <-chan int {
    var t0 <-chan int = func(p0 chan int) <-chan int {
        return p0
    }(self__0)
    return t0
}

func _goml_m_std_p_channel_p_try__select__priority____T__isize(operations__0 []_goml_m_std_p_channel_p_Operation____isize) _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError {
    var t0 int
    var inline5 int = len(operations__0)
    t0 = inline5
    var t1 bool = t0 == 0
    if t1 {
        var t2 _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError = _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError{
            _tag: 1,
        }
        return t2
    } else {
        var prepared__0 _goml_m_std_p_channel_p_PreparedOperations____isize = _goml_m_std_p_channel_p_prepare____T__isize(operations__0)
        var mtmp0 Tuple4_10Slice_3int_21Slice_13Receiver_3int_19Slice_11Sender_3int_10Slice_3int = _goml_m_std_p_channel_p_prepared__slices____T__isize(prepared__0)
        var x0 []int = mtmp0._0
        var x1 []<-chan int = mtmp0._1
        var x2 []chan<- int = mtmp0._2
        var x3 []int = mtmp0._3
        var mtmp1 Tuple5_3int_3int_3int_4bool_4bool = func(p0 []int, p1 []<-chan int, p2 []chan<- int, p3 []int) Tuple5_3int_3int_3int_4bool_4bool {
            var index int = 0
            var receive_index int = 0
            var send_index int = 0
            var value int
            var open bool
            var selected bool
            var cases []_goml_reflect.SelectCase = make([]_goml_reflect.SelectCase, 2)
            cases[1] = _goml_reflect.SelectCase{
                Dir: _goml_reflect.SelectDefault,
            }
            var chosen int
            var received _goml_reflect.Value
            for {
                if index >= len(p0) {
                    break
                }
                if p0[index] == 0 {
                    cases[0] = _goml_reflect.SelectCase{
                        Dir: _goml_reflect.SelectRecv,
                        Chan: _goml_reflect.ValueOf(p1[receive_index]),
                    }
                    chosen, received, open = _goml_reflect.Select(cases)
                    if chosen == 0 {
                        value = received.Interface().(int)
                        selected = true
                    }
                    receive_index = receive_index + 1
                } else {
                    cases[0] = _goml_reflect.SelectCase{
                        Dir: _goml_reflect.SelectSend,
                        Chan: _goml_reflect.ValueOf(p2[send_index]),
                        Send: _goml_reflect.ValueOf(p3[send_index]),
                    }
                    chosen, received, open = _goml_reflect.Select(cases)
                    if chosen == 0 {
                        selected = true
                    }
                    send_index = send_index + 1
                }
                if selected {
                    return Tuple5_3int_3int_3int_4bool_4bool{
                        _0: index,
                        _1: p0[index],
                        _2: value,
                        _3: open,
                        _4: true,
                    }
                }
                index = index + 1
            }
            return Tuple5_3int_3int_3int_4bool_4bool{
                _0: -1,
                _1: -1,
                _2: value,
                _3: false,
                _4: false,
            }
        }(x0, x1, x2, x3)
        var x4 int = mtmp1._0
        var x5 int = mtmp1._1
        var x6 int = mtmp1._2
        var x7 bool = mtmp1._3
        var x8 bool = mtmp1._4
        var jp0 _goml_m_Option____std_p_channel_p_Selection____isize
        if x8 {
            var t4 _goml_m_std_p_channel_p_Selection____isize
            var inline0 bool = x5 == 0
            if inline0 {
                var inline1 Option__isize
                if x7 {
                    var inline3 Option__isize = Option__isize{
                        _p0: x6,
                        _tag: 1,
                    }
                    inline1 = inline3
                } else {
                    inline1 = Option__isize{
                        _tag: 0,
                    }
                }
                var inline2 _goml_m_std_p_channel_p_Selection____isize = Received{
                    _0: x4,
                    _1: inline1,
                }
                t4 = inline2
            } else {
                var inline4 _goml_m_std_p_channel_p_Selection____isize = Sent{
                    _0: x4,
                }
                t4 = inline4
            }
            var t5 _goml_m_Option____std_p_channel_p_Selection____isize = _goml_m_Option____std_p_channel_p_Selection____isize{
                _p0: t4,
                _tag: 1,
            }
            jp0 = t5
        } else {
            jp0 = _goml_m_Option____std_p_channel_p_Selection____isize{
                _tag: 0,
            }
        }
        var t3 _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError = _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError{
            _p0: jp0,
            _tag: 0,
        }
        return t3
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_h6161506801e1e3b7949d66d7b76c6b61_tion_l_isize_r_(self__0 *_goml_vec__goml_m_std_p_channel_p_Operation____isize) int {
    var t0 int = vec_len___goml_m_Vec__30std_p_channel_p_Operation____isize(self__0)
    return t0
}

func _goml_m_std_p_channel_p_try__select____T__isize(operations__0 []_goml_m_std_p_channel_p_Operation____isize) _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError {
    var t0 int
    var inline5 int = len(operations__0)
    t0 = inline5
    var t1 bool = t0 == 0
    if t1 {
        var t2 _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError = _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError{
            _tag: 1,
        }
        return t2
    } else {
        var prepared__0 _goml_m_std_p_channel_p_PreparedOperations____isize = _goml_m_std_p_channel_p_prepare____T__isize(operations__0)
        var mtmp0 Tuple4_10Slice_3int_21Slice_13Receiver_3int_19Slice_11Sender_3int_10Slice_3int = _goml_m_std_p_channel_p_prepared__slices____T__isize(prepared__0)
        var x0 []int = mtmp0._0
        var x1 []<-chan int = mtmp0._1
        var x2 []chan<- int = mtmp0._2
        var x3 []int = mtmp0._3
        var mtmp1 Tuple5_3int_3int_3int_4bool_4bool = func(p0 []int, p1 []<-chan int, p2 []chan<- int, p3 []int) Tuple5_3int_3int_3int_4bool_4bool {
            var cases []_goml_reflect.SelectCase = make([]_goml_reflect.SelectCase, len(p0) + 1)
            var index int = 0
            var receive_index int = 0
            var send_index int = 0
            for {
                if index >= len(p0) {
                    break
                }
                if p0[index] == 0 {
                    cases[index] = _goml_reflect.SelectCase{
                        Dir: _goml_reflect.SelectRecv,
                        Chan: _goml_reflect.ValueOf(p1[receive_index]),
                    }
                    receive_index = receive_index + 1
                } else {
                    cases[index] = _goml_reflect.SelectCase{
                        Dir: _goml_reflect.SelectSend,
                        Chan: _goml_reflect.ValueOf(p2[send_index]),
                        Send: _goml_reflect.ValueOf(p3[send_index]),
                    }
                    send_index = send_index + 1
                }
                index = index + 1
            }
            cases[len(p0)] = _goml_reflect.SelectCase{
                Dir: _goml_reflect.SelectDefault,
            }
            var chosen int
            var received _goml_reflect.Value
            var open bool
            chosen, received, open = _goml_reflect.Select(cases)
            var value int
            if chosen == len(p0) {
                return Tuple5_3int_3int_3int_4bool_4bool{
                    _0: -1,
                    _1: -1,
                    _2: value,
                    _3: false,
                    _4: false,
                }
            }
            if p0[chosen] == 0 {
                value = received.Interface().(int)
            }
            return Tuple5_3int_3int_3int_4bool_4bool{
                _0: chosen,
                _1: p0[chosen],
                _2: value,
                _3: open,
                _4: true,
            }
        }(x0, x1, x2, x3)
        var x4 int = mtmp1._0
        var x5 int = mtmp1._1
        var x6 int = mtmp1._2
        var x7 bool = mtmp1._3
        var x8 bool = mtmp1._4
        var jp0 _goml_m_Option____std_p_channel_p_Selection____isize
        if x8 {
            var t4 _goml_m_std_p_channel_p_Selection____isize
            var inline0 bool = x5 == 0
            if inline0 {
                var inline1 Option__isize
                if x7 {
                    var inline3 Option__isize = Option__isize{
                        _p0: x6,
                        _tag: 1,
                    }
                    inline1 = inline3
                } else {
                    inline1 = Option__isize{
                        _tag: 0,
                    }
                }
                var inline2 _goml_m_std_p_channel_p_Selection____isize = Received{
                    _0: x4,
                    _1: inline1,
                }
                t4 = inline2
            } else {
                var inline4 _goml_m_std_p_channel_p_Selection____isize = Sent{
                    _0: x4,
                }
                t4 = inline4
            }
            var t5 _goml_m_Option____std_p_channel_p_Selection____isize = _goml_m_Option____std_p_channel_p_Selection____isize{
                _p0: t4,
                _tag: 1,
            }
            jp0 = t5
        } else {
            jp0 = _goml_m_Option____std_p_channel_p_Selection____isize{
                _tag: 0,
            }
        }
        var t3 _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError = _goml_m_Result____Option____st_ha315a4af527ac57871bea61a2ab15462_l_p_SelectError{
            _p0: jp0,
            _tag: 0,
        }
        return t3
    }
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_sender____T__isize(self__0 chan int) chan<- int {
    var t0 chan<- int = func(p0 chan int) chan<- int {
        return p0
    }(self__0)
    return t0
}

func _goml_m_std_p_channel_p_select____T__isize(operations__0 []_goml_m_std_p_channel_p_Operation____isize) _goml_m_Result____std_p_channe_h5400d2fb6c865f5b36aaa9b0f9e787f7_l_p_SelectError {
    var t0 int
    var inline5 int = len(operations__0)
    t0 = inline5
    var t1 bool = t0 == 0
    if t1 {
        var t2 _goml_m_Result____std_p_channe_h5400d2fb6c865f5b36aaa9b0f9e787f7_l_p_SelectError = _goml_m_Result____std_p_channe_h5400d2fb6c865f5b36aaa9b0f9e787f7_l_p_SelectError{
            _tag: 1,
        }
        return t2
    } else {
        var prepared__0 _goml_m_std_p_channel_p_PreparedOperations____isize = _goml_m_std_p_channel_p_prepare____T__isize(operations__0)
        var mtmp0 Tuple4_10Slice_3int_21Slice_13Receiver_3int_19Slice_11Sender_3int_10Slice_3int = _goml_m_std_p_channel_p_prepared__slices____T__isize(prepared__0)
        var x0 []int = mtmp0._0
        var x1 []<-chan int = mtmp0._1
        var x2 []chan<- int = mtmp0._2
        var x3 []int = mtmp0._3
        var mtmp1 Tuple5_3int_3int_3int_4bool_4bool = func(p0 []int, p1 []<-chan int, p2 []chan<- int, p3 []int) Tuple5_3int_3int_3int_4bool_4bool {
            var cases []_goml_reflect.SelectCase = make([]_goml_reflect.SelectCase, len(p0))
            var index int = 0
            var receive_index int = 0
            var send_index int = 0
            for {
                if index >= len(p0) {
                    break
                }
                if p0[index] == 0 {
                    cases[index] = _goml_reflect.SelectCase{
                        Dir: _goml_reflect.SelectRecv,
                        Chan: _goml_reflect.ValueOf(p1[receive_index]),
                    }
                    receive_index = receive_index + 1
                } else {
                    cases[index] = _goml_reflect.SelectCase{
                        Dir: _goml_reflect.SelectSend,
                        Chan: _goml_reflect.ValueOf(p2[send_index]),
                        Send: _goml_reflect.ValueOf(p3[send_index]),
                    }
                    send_index = send_index + 1
                }
                index = index + 1
            }
            var chosen int
            var received _goml_reflect.Value
            var open bool
            chosen, received, open = _goml_reflect.Select(cases)
            var value int
            if p0[chosen] == 0 {
                value = received.Interface().(int)
            }
            return Tuple5_3int_3int_3int_4bool_4bool{
                _0: chosen,
                _1: p0[chosen],
                _2: value,
                _3: open,
                _4: true,
            }
        }(x0, x1, x2, x3)
        var x4 int = mtmp1._0
        var x5 int = mtmp1._1
        var x6 int = mtmp1._2
        var x7 bool = mtmp1._3
        var t3 _goml_m_std_p_channel_p_Selection____isize
        var inline0 bool = x5 == 0
        if inline0 {
            var inline1 Option__isize
            if x7 {
                var inline3 Option__isize = Option__isize{
                    _p0: x6,
                    _tag: 1,
                }
                inline1 = inline3
            } else {
                inline1 = Option__isize{
                    _tag: 0,
                }
            }
            var inline2 _goml_m_std_p_channel_p_Selection____isize = Received{
                _0: x4,
                _1: inline1,
            }
            t3 = inline2
        } else {
            var inline4 _goml_m_std_p_channel_p_Selection____isize = Sent{
                _0: x4,
            }
            t3 = inline4
        }
        var t4 _goml_m_Result____std_p_channe_h5400d2fb6c865f5b36aaa9b0f9e787f7_l_p_SelectError = _goml_m_Result____std_p_channe_h5400d2fb6c865f5b36aaa9b0f9e787f7_l_p_SelectError{
            _p0: t3,
            _tag: 0,
        }
        return t4
    }
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_recv____T__isize(self__0 chan int) Option__isize {
    var mtmp0 Tuple2_3int_4bool = func(p0 chan int) Tuple2_3int_4bool {
        var value int
        var ok bool
        value, ok = <-p0
        return Tuple2_3int_4bool{
            _0: value,
            _1: ok,
        }
    }(self__0)
    var x0 int = mtmp0._0
    var x1 bool = mtmp0._1
    if x1 {
        var t0 Option__isize = Option__isize{
            _p0: x0,
            _tag: 1,
        }
        return t0
    } else {
        return Option__isize{
            _tag: 0,
        }
    }
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_std_p_channel_p_prepare____T__isize(operations__0 []_goml_m_std_p_channel_p_Operation____isize) _goml_m_std_p_channel_p_PreparedOperations____isize {
    var t0 int
    var inline11 int = len(operations__0)
    t0 = inline11
    var kinds__0 *_goml_vec_int
    var inline10 *_goml_vec_int = vec_with_capacity__Vec_3int(t0)
    kinds__0 = inline10
    var receivers__0 *_goml_vec_Receiver_3int
    var inline9 *_goml_vec_Receiver_3int = vec_new__Vec_13Receiver_3int()
    receivers__0 = inline9
    var senders__0 *_goml_vec_Sender_3int
    var inline8 *_goml_vec_Sender_3int = vec_new__Vec_11Sender_3int()
    senders__0 = inline8
    var values__0 *_goml_vec_int
    var inline7 *_goml_vec_int = vec_new__Vec_3int()
    values__0 = inline7
    var for_limit0 int = len(operations__0)
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t2 bool = for_index0 < for_limit0
        if t2 {
            var for_item0 _goml_m_std_p_channel_p_Operation____isize = operations__0[for_index0]
            var t3 int = for_index0 + 1
            for_index0 = t3
            switch for_item0.(type) {
            case Receive:
                var x0 <-chan int = for_item0.(Receive)._0
                var inline1 int = 0
                vec_push__Vec_3int(kinds__0, inline1)
                vec_push__Vec_13Receiver_3int(receivers__0, x0)
                continue
            case Send:
                var x1 chan<- int = for_item0.(Send)._0
                var x2 int = for_item0.(Send)._1
                var inline5 int = 1
                vec_push__Vec_3int(kinds__0, inline5)
                vec_push__Vec_11Sender_3int(senders__0, x1)
                vec_push__Vec_3int(values__0, x2)
                continue
            default:
                panic("non-exhaustive match")
            }
        } else {
            break Loop_loop0
        }
    }
    var t1 _goml_m_std_p_channel_p_PreparedOperations____isize = _goml_m_std_p_channel_p_PreparedOperations____isize{
        kinds: kinds__0,
        receivers: receivers__0,
        senders: senders__0,
        values: values__0,
    }
    return t1
}

func _goml_m_std_p_channel_p_prepared__slices____T__isize(prepared__0 _goml_m_std_p_channel_p_PreparedOperations____isize) Tuple4_10Slice_3int_21Slice_13Receiver_3int_19Slice_11Sender_3int_10Slice_3int {
    var t0 *_goml_vec_int = prepared__0.kinds
    var t1 *_goml_vec_int = prepared__0.kinds
    var t2 int
    var inline3 int = vec_len__Vec_3int(t1)
    t2 = inline3
    var t3 []int = t0.items[0:t2]
    var t4 *_goml_vec_Receiver_3int = prepared__0.receivers
    var t5 *_goml_vec_Receiver_3int = prepared__0.receivers
    var t6 int
    var inline2 int = vec_len__Vec_13Receiver_3int(t5)
    t6 = inline2
    var t7 []<-chan int = t4.items[0:t6]
    var t8 *_goml_vec_Sender_3int = prepared__0.senders
    var t9 *_goml_vec_Sender_3int = prepared__0.senders
    var t10 int
    var inline1 int = vec_len__Vec_11Sender_3int(t9)
    t10 = inline1
    var t11 []chan<- int = t8.items[0:t10]
    var t12 *_goml_vec_int = prepared__0.values
    var t13 *_goml_vec_int = prepared__0.values
    var t14 int
    var inline0 int = vec_len__Vec_3int(t13)
    t14 = inline0
    var t15 []int = t12.items[0:t14]
    var t16 Tuple4_10Slice_3int_21Slice_13Receiver_3int_19Slice_11Sender_3int_10Slice_3int = Tuple4_10Slice_3int_21Slice_13Receiver_3int_19Slice_11Sender_3int_10Slice_3int{
        _0: t3,
        _1: t7,
        _2: t11,
        _3: t15,
    }
    return t16
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

func main() {
    main0()
}
