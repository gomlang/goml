package main

import (
    _goml_os "os"
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

type Sink struct {
    prefix string
}

type Ordering uint8

type _goml_m_dyn_____d_dyn_k_dynCount_u_Label__vtable struct {
    _goml_m_Label_i_label func(any) string
    _goml_m_Count_i_count func(any) int
}

type _goml_m_dyn_____d_dyn_k_dynCount_u_Label struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynCount_u_Label__vtable
}

type _goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r___vtable struct {
    _goml_m_First_i_value func(any) int
}

type _goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r_ struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r___vtable
}

type _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r___vtable struct {
    _goml_m_Consumer_i_consume func(any, string) string
}

type _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r_ struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r___vtable
}

type _goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r___vtable struct {
    _goml_m_Consumer_i_consume func(any, int) string
}

type _goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r_ struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r___vtable
}

type _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label__vtable struct {
    _goml_m_Consumer_i_consume func(any, string) string
    _goml_m_Label_i_label func(any) string
}

type _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label__vtable
}

type _goml_m_dyn_____d_dyn_k_dynChild_l_string_r__u_Root_l_string_r___vtable struct {
    _goml_m_Root_i_root func(any) string
    _goml_m_Child_i_child func(any) string
}

type _goml_m_dyn_____d_dyn_k_dynChild_l_string_r__u_Root_l_string_r_ struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynChild_l_string_r__u_Root_l_string_r___vtable
}

type _goml_m_dyn_____d_dyn_k_dynFam_h46fb78443fc145ca82f4caf7b7cfa239_ring_r___vtable struct {
    _goml_m_First_i_value func(any) int
    _goml_m_Second_i_value func(any) string
}

type _goml_m_dyn_____d_dyn_k_dynFam_hd40cf3de0a5a94b4a60c87cac34baed7_m_x3d_string_r_ struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynFam_h46fb78443fc145ca82f4caf7b7cfa239_ring_r___vtable
}

type _goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent__vtable struct {
    _goml_m_NamedChild_i_name func(any) string
    _goml_m_NamedParent_i_name func(any) string
}

type _goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent__vtable
}

func _goml_m_dyn_____d_dyn_k_dynChi_hc4219ff85cd5e4edeecd7f9c461cc71e_____Root_i_root(self any) string {
    return _goml_m_trait__impl_i_Root_i__l_string_r__x40_Sink_i_root(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynChi_h487eb1bac7ecfdb2e4fb5c97d454d2cb___Child_i_child(self any) string {
    return _goml_m_trait__impl_i_Child_i__l_string_r__x40_Sink_i_child(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynChi_hf0508c19f8ba6b317c0993d5169addaa__vtable____Sink() *_goml_m_dyn_____d_dyn_k_dynChild_l_string_r__u_Root_l_string_r___vtable {
    return &_goml_m_dyn_____d_dyn_k_dynChild_l_string_r__u_Root_l_string_r___vtable{
        _goml_m_Root_i_root: _goml_m_dyn_____d_dyn_k_dynChi_hc4219ff85cd5e4edeecd7f9c461cc71e_____Root_i_root,
        _goml_m_Child_i_child: _goml_m_dyn_____d_dyn_k_dynChi_h487eb1bac7ecfdb2e4fb5c97d454d2cb___Child_i_child,
    }
}

func _goml_m_dyn_____d_dyn_k_dynCon_ha54fb8b5be72e8a9302a5f16b7f3eb2c_sumer_i_consume(self any, p0 int) string {
    return _goml_m_trait__impl_i_Consumer_i__l_isize_r__x40_Sink_i_consume(self.(Sink), p0)
}

func _goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r_____vtable____Sink() *_goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r___vtable {
    return &_goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r___vtable{
        _goml_m_Consumer_i_consume: _goml_m_dyn_____d_dyn_k_dynCon_ha54fb8b5be72e8a9302a5f16b7f3eb2c_sumer_i_consume,
    }
}

func _goml_m_dyn_____d_dyn_k_dynCon_h1be1ecd1fa7ab8a64f6154f3ec465e05_sumer_i_consume(self any, p0 string) string {
    return _goml_m_trait__impl_i_Consumer_i__l_string_r__x40_Sink_i_consume(self.(Sink), p0)
}

func _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r_____vtable____Sink() *_goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r___vtable {
    return &_goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r___vtable{
        _goml_m_Consumer_i_consume: _goml_m_dyn_____d_dyn_k_dynCon_h1be1ecd1fa7ab8a64f6154f3ec465e05_sumer_i_consume,
    }
}

func _goml_m_dyn_____d_dyn_k_dynCon_h31aa7bb43d859cf4123c595578227ee7_sumer_i_consume(self any, p0 string) string {
    return _goml_m_trait__impl_i_Consumer_i__l_string_r__x40_Sink_i_consume(self.(Sink), p0)
}

func _goml_m_dyn_____d_dyn_k_dynCon_he10298e4aa11ce4f14200c6f820fd4d6___Label_i_label(self any) string {
    return _goml_m_trait__impl_i_Label_i_Sink_i_label(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label____vtable____Sink() *_goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label__vtable {
    return &_goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label__vtable{
        _goml_m_Consumer_i_consume: _goml_m_dyn_____d_dyn_k_dynCon_h31aa7bb43d859cf4123c595578227ee7_sumer_i_consume,
        _goml_m_Label_i_label: _goml_m_dyn_____d_dyn_k_dynCon_he10298e4aa11ce4f14200c6f820fd4d6___Label_i_label,
    }
}

func _goml_m_dyn_____d_dyn_k_dynCount_u_Label____wrap____Sink____Label_i_label(self any) string {
    return _goml_m_trait__impl_i_Label_i_Sink_i_label(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynCount_u_Label____wrap____Sink____Count_i_count(self any) int {
    return _goml_m_trait__impl_i_Count_i_Sink_i_count(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynCount_u_Label____vtable____Sink() *_goml_m_dyn_____d_dyn_k_dynCount_u_Label__vtable {
    return &_goml_m_dyn_____d_dyn_k_dynCount_u_Label__vtable{
        _goml_m_Label_i_label: _goml_m_dyn_____d_dyn_k_dynCount_u_Label____wrap____Sink____Label_i_label,
        _goml_m_Count_i_count: _goml_m_dyn_____d_dyn_k_dynCount_u_Label____wrap____Sink____Count_i_count,
    }
}

func _goml_m_dyn_____d_dyn_k_dynFam_hb8410b29dac75e30699ec1ed63c4506a___First_i_value(self any) int {
    return _goml_m_trait__impl_i_First_i_Sink_i_value(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynFam_h8c023ee0f963fa81b3f82d6aff22749a__Second_i_value(self any) string {
    return _goml_m_trait__impl_i_Second_i_Sink_i_value(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynFam_h60f1288ae0fc67db8ffad1408254e3f8__vtable____Sink() *_goml_m_dyn_____d_dyn_k_dynFam_h46fb78443fc145ca82f4caf7b7cfa239_ring_r___vtable {
    return &_goml_m_dyn_____d_dyn_k_dynFam_h46fb78443fc145ca82f4caf7b7cfa239_ring_r___vtable{
        _goml_m_First_i_value: _goml_m_dyn_____d_dyn_k_dynFam_hb8410b29dac75e30699ec1ed63c4506a___First_i_value,
        _goml_m_Second_i_value: _goml_m_dyn_____d_dyn_k_dynFam_h8c023ee0f963fa81b3f82d6aff22749a__Second_i_value,
    }
}

func _goml_m_dyn_____d_dyn_k_dynFir_had62bca62cea7af7e6fb3c9427149d93___First_i_value(self any) int {
    return _goml_m_trait__impl_i_First_i_Sink_i_value(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r_____vtable____Sink() *_goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r___vtable {
    return &_goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r___vtable{
        _goml_m_First_i_value: _goml_m_dyn_____d_dyn_k_dynFir_had62bca62cea7af7e6fb3c9427149d93___First_i_value,
    }
}

func _goml_m_dyn_____d_dyn_k_dynNam_h41043ac0941a4cf641edd813498d743a_medChild_i_name(self any) string {
    return _goml_m_trait__impl_i_NamedChild_i_Sink_i_name(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynNam_ha6c9c9580026e0192cd43540a5786598_edParent_i_name(self any) string {
    return _goml_m_trait__impl_i_NamedParent_i_Sink_i_name(self.(Sink))
}

func _goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent____vtable____Sink() *_goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent__vtable {
    return &_goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent__vtable{
        _goml_m_NamedChild_i_name: _goml_m_dyn_____d_dyn_k_dynNam_h41043ac0941a4cf641edd813498d743a_medChild_i_name,
        _goml_m_NamedParent_i_name: _goml_m_dyn_____d_dyn_k_dynNam_ha6c9c9580026e0192cd43540a5786598_edParent_i_name,
    }
}

func _goml_m_trait__impl_i_Consumer_i__l_string_r__x40_Sink_i_consume(self__0 Sink, value__0 string) string {
    var t0 string = self__0.prefix
    var t1 string = t0 + value__0
    return t1
}

func _goml_m_trait__impl_i_Consumer_i__l_isize_r__x40_Sink_i_consume(self__0 Sink, value__0 int) string {
    var t0 string = self__0.prefix
    var t1 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t1 = inline0
    var t2 string = t0 + t1
    return t2
}

func _goml_m_trait__impl_i_Label_i_Sink_i_label(self__0 Sink) string {
    var t0 string = self__0.prefix
    return t0
}

func _goml_m_trait__impl_i_Count_i_Sink_i_count(self__0 Sink) int {
    return 7
}

func _goml_m_trait__impl_i_Root_i__l_string_r__x40_Sink_i_root(self__0 Sink) string {
    var t0 string = self__0.prefix
    return t0
}

func _goml_m_trait__impl_i_Child_i__l_string_r__x40_Sink_i_child(self__0 Sink) string {
    return "child"
}

func _goml_m_trait__impl_i_First_i_Sink_i_value(self__0 Sink) int {
    return 42
}

func _goml_m_trait__impl_i_Second_i_Sink_i_value(self__0 Sink) string {
    var t0 string = self__0.prefix
    return t0
}

func _goml_m_trait__impl_i_NamedParent_i_Sink_i_name(self__0 Sink) string {
    return "parent"
}

func _goml_m_trait__impl_i_NamedChild_i_Sink_i_name(self__0 Sink) string {
    return "child"
}

func main0() struct{} {
    var t0 Sink = Sink{
        prefix: "text:",
    }
    var text__0 _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r_ = _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r_{
        data: t0,
        vtable: _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r_____vtable____Sink(),
    }
    var t1 Sink = Sink{
        prefix: "number:",
    }
    var number__0 _goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r_ = _goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r_{
        data: t1,
        vtable: _goml_m_dyn_____d_dyn_k_dynConsumer_l_isize_r_____vtable____Sink(),
    }
    var t2 string = text__0.vtable._goml_m_Consumer_i_consume(text__0.data, "ok")
    println__T_string(t2)
    var t3 string = number__0.vtable._goml_m_Consumer_i_consume(number__0.data, 42)
    println__T_string(t3)
    var t4 Sink = Sink{
        prefix: "both:",
    }
    var both__0 _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label = erase__T_Sink(t4)
    var t5 string = both__0.vtable._goml_m_Consumer_i_consume(both__0.data, "ok")
    var inline28 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline28)
    var t6 string = both__0.vtable._goml_m_Label_i_label(both__0.data)
    var inline26 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t6)
    _goml_runtime_core_string_println(inline26)
    var t7 Sink = Sink{
        prefix: "count:",
    }
    var combined__0 _goml_m_dyn_____d_dyn_k_dynCount_u_Label = _goml_m_dyn_____d_dyn_k_dynCount_u_Label{
        data: t7,
        vtable: _goml_m_dyn_____d_dyn_k_dynCount_u_Label____vtable____Sink(),
    }
    var t8 string = combined__0.vtable._goml_m_Label_i_label(combined__0.data)
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t8)
    _goml_runtime_core_string_println(inline24)
    var t9 int = combined__0.vtable._goml_m_Count_i_count(combined__0.data)
    var inline22 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline22)
    var t10 _goml_m_dyn_____d_dyn_k_dynCount_u_Label
    t10 = combined__0
    var t11 string = t10.vtable._goml_m_Label_i_label(t10.data)
    var inline20 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t11)
    _goml_runtime_core_string_println(inline20)
    var t12 Sink = Sink{
        prefix: "inherited",
    }
    var child__0 _goml_m_dyn_____d_dyn_k_dynChild_l_string_r__u_Root_l_string_r_
    var inline19 _goml_m_dyn_____d_dyn_k_dynChild_l_string_r__u_Root_l_string_r_ = _goml_m_dyn_____d_dyn_k_dynChild_l_string_r__u_Root_l_string_r_{
        data: t12,
        vtable: _goml_m_dyn_____d_dyn_k_dynChi_hf0508c19f8ba6b317c0993d5169addaa__vtable____Sink(),
    }
    child__0 = inline19
    var t13 string = child__0.vtable._goml_m_Root_i_root(child__0.data)
    var inline17 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t13)
    _goml_runtime_core_string_println(inline17)
    var t14 string = child__0.vtable._goml_m_Child_i_child(child__0.data)
    var inline15 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t14)
    _goml_runtime_core_string_println(inline15)
    var t15 string = child__0.vtable._goml_m_Root_i_root(child__0.data)
    var inline13 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t15)
    _goml_runtime_core_string_println(inline13)
    var t16 string = text__0.vtable._goml_m_Consumer_i_consume(text__0.data, "ufcs")
    var inline11 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t16)
    _goml_runtime_core_string_println(inline11)
    var t17 Sink = Sink{
        prefix: "associated",
    }
    var family__0 _goml_m_dyn_____d_dyn_k_dynFam_hd40cf3de0a5a94b4a60c87cac34baed7_m_x3d_string_r_ = _goml_m_dyn_____d_dyn_k_dynFam_hd40cf3de0a5a94b4a60c87cac34baed7_m_x3d_string_r_{
        data: t17,
        vtable: _goml_m_dyn_____d_dyn_k_dynFam_h60f1288ae0fc67db8ffad1408254e3f8__vtable____Sink(),
    }
    var t18 int = family__0.vtable._goml_m_First_i_value(family__0.data)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t18)
    _goml_runtime_core_string_println(inline9)
    var t19 string = family__0.vtable._goml_m_Second_i_value(family__0.data)
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t19)
    _goml_runtime_core_string_println(inline7)
    var t20 Sink = Sink{
        prefix: "",
    }
    var named__0 _goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent = _goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent{
        data: t20,
        vtable: _goml_m_dyn_____d_dyn_k_dynNamedChild_u_NamedParent____vtable____Sink(),
    }
    var t21 string = named__0.vtable._goml_m_NamedChild_i_name(named__0.data)
    var inline5 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t21)
    _goml_runtime_core_string_println(inline5)
    var t22 string = named__0.vtable._goml_m_NamedParent_i_name(named__0.data)
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t22)
    _goml_runtime_core_string_println(inline3)
    var t23 Sink = Sink{
        prefix: "",
    }
    var t24 _goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r_ = _goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r_{
        data: t23,
        vtable: _goml_m_dyn_____d_dyn_k_dynFirst_l_Item_x3d_isize_r_____vtable____Sink(),
    }
    var t25 int
    var inline2 int = t24.vtable._goml_m_First_i_value(t24.data)
    t25 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t25)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func erase__T_Sink(value__0 Sink) _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label {
    var t0 _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label = _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label{
        data: value__0,
        vtable: _goml_m_dyn_____d_dyn_k_dynConsumer_l_string_r__u_Label____vtable____Sink(),
    }
    return t0
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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
