TEXT ·_goml_m_inherent_i_std_p_simd_p_f32x4_i_std_p_simd_p_f32x4_i_add(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    ADDPS X1, X0
    MOVUPS X0, ret+32(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_f32x4_i_std_p_simd_p_f32x4_i_mul(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    MOVUPS X0, X2
    MULPS X1, X2
    MOVUPS X2, ret+32(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_f32x4_i_std_p_simd_p_f32x4_i_sub(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    SUBPS X1, X0
    MOVUPS X0, ret+32(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_f32x4_i_std_p_simd_p_f32x4_i_splat(SB), 4, $0-24
    MOVSS value__0+0(FP), X0
    PSHUFD $0, X0, X0
    MOVUPS X0, ret+8(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_f32x4_i_std_p_simd_p_f32x4_i_div(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    DIVPS X1, X0
    MOVUPS X0, ret+32(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_f32x4_i_std_p_simd_p_f32x4_i_simd__lt(SB), 4, $0-33
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    MOVUPS X0, X2
    CMPPS X1, X2, $1
    MOVMSKPS X2, AX
    MOVB AX, ret+32(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd__h62612dde994cfc616b451f5df4971037__p_simd_p_f32x4(SB), 4, $0-56
    MOVBLZX self__0+0(FP), AX
    MOVQ AX, X0
    PSHUFD $0, X0, X0
    MOVUPS _goml_m_inherent_i_std_p_simd__h62612dde994cfc616b451f5df4971037__p_simd_p_f32x4_constant_0<>(SB), X15
    PAND X15, X0
    PCMPEQL X15, X0
    MOVUPS if_true__0+4(FP), X1
    MOVUPS if_false__0+20(FP), X2
    MOVUPS X0, X14
    MOVUPS X0, X15
    ANDPS X1, X14
    ANDNPS X2, X15
    MOVUPS X14, X3
    ORPS X15, X3
    MOVUPS X3, ret+40(FP)
    RET
DATA _goml_m_inherent_i_std_p_simd__h62612dde994cfc616b451f5df4971037__p_simd_p_f32x4_constant_0<>+0(SB)/4, $1
DATA _goml_m_inherent_i_std_p_simd__h62612dde994cfc616b451f5df4971037__p_simd_p_f32x4_constant_0<>+4(SB)/4, $2
DATA _goml_m_inherent_i_std_p_simd__h62612dde994cfc616b451f5df4971037__p_simd_p_f32x4_constant_0<>+8(SB)/4, $4
DATA _goml_m_inherent_i_std_p_simd__h62612dde994cfc616b451f5df4971037__p_simd_p_f32x4_constant_0<>+12(SB)/4, $8
GLOBL _goml_m_inherent_i_std_p_simd__h62612dde994cfc616b451f5df4971037__p_simd_p_f32x4_constant_0<>(SB), 24, $16

TEXT ·_goml_simd_avx2__goml_m_inherent_i_std_p_simd_p_f32x8_i_std_p_simd_p_f32x8_i_add(SB), 4, $0-96
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VADDPS Y1, Y0, Y0
    VMOVUPS Y0, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inherent_i_std_p_simd_p_f32x8_i_std_p_simd_p_f32x8_i_mul(SB), 4, $0-96
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VMULPS Y1, Y0, Y2
    VMOVUPS Y2, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inherent_i_std_p_simd_p_f32x8_i_std_p_simd_p_f32x8_i_sub(SB), 4, $0-96
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VSUBPS Y1, Y0, Y0
    VMOVUPS Y0, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_hbf308f0da68aabac491becc7e3ab5729_p_f32x8_i_splat(SB), 4, $0-40
    MOVSS value__0+0(FP), X0
    VPBROADCASTD X0, Y0
    VMOVUPS Y0, ret+8(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inherent_i_std_p_simd_p_f32x8_i_std_p_simd_p_f32x8_i_div(SB), 4, $0-96
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VDIVPS Y1, Y0, Y0
    VMOVUPS Y0, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_hc1d5ed47edfb6dd864fe75f13dda8164_32x8_i_simd__lt(SB), 4, $0-65
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VCMPPS $1, Y1, Y0, Y2
    VMOVMSKPS Y2, AX
    MOVB AX, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8(SB), 4, $0-104
    MOVBLZX self__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTD X0, Y0
    VMOVUPS _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>(SB), Y15
    VPAND Y15, Y0, Y0
    VPCMPEQD Y15, Y0, Y0
    VMOVUPS if_true__0+4(FP), Y1
    VMOVUPS if_false__0+36(FP), Y2
    VANDPS Y1, Y0, Y14
    VANDNPS Y2, Y0, Y15
    VORPS Y15, Y14, Y3
    VMOVUPS Y3, ret+72(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_h9040d6dbbca5b39b4c91c70938d8ec59__p_simd_p_f32x8_constant_0<>(SB), 24, $32

TEXT ·_goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum(SB), 4, $0-36
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS Y0, Y1
    VMOVUPS _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>(SB), Y14
    VPERMPS Y1, Y14, Y15
    VADDPS Y15, Y1, Y1
    VMOVUPS _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>(SB), Y14
    VPERMPS Y1, Y14, Y15
    VADDPS Y15, Y1, Y1
    VMOVUPS _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>(SB), Y14
    VPERMPS Y1, Y14, Y15
    VADDPS Y15, Y1, Y1
    MOVSS X1, ret+32(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>+4(SB)/4, $0
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>+8(SB)/4, $3
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>+12(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>+16(SB)/4, $5
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>+20(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>+24(SB)/4, $7
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>+28(SB)/4, $6
GLOBL _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_0<>(SB), 24, $32
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>+0(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>+4(SB)/4, $3
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>+8(SB)/4, $0
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>+12(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>+16(SB)/4, $6
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>+20(SB)/4, $7
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>+24(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>+28(SB)/4, $5
GLOBL _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_1<>(SB), 24, $32
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>+0(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>+4(SB)/4, $5
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>+8(SB)/4, $6
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>+12(SB)/4, $7
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>+16(SB)/4, $0
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>+20(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>+24(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>+28(SB)/4, $3
GLOBL _goml_simd_avx2__goml_m_inhere_hb2b8ea87034a6f88e9726aaa88fe6075_8_i_reduce__sum_constant_2<>(SB), 24, $32

TEXT ·_goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand(SB), 4, $0-9
    MOVBLZX self__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTD X0, Y0
    VMOVUPS _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>(SB), Y15
    VPAND Y15, Y0, Y0
    VPCMPEQD Y15, Y0, Y0
    MOVBLZX other__0+1(FP), AX
    MOVQ AX, X1
    VPBROADCASTD X1, Y1
    VMOVUPS _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>(SB), Y15
    VPAND Y15, Y1, Y1
    VPCMPEQD Y15, Y1, Y1
    VANDPS Y1, Y0, Y0
    VMOVMSKPS Y0, AX
    MOVB AX, ret+8(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_0<>(SB), 24, $32
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_h92d9adefb93045451d9d8c98adbea12c_sk32x8_i_bitand_constant_1<>(SB), 24, $32

TEXT ·_goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot(SB), 4, $0-9
    MOVBLZX self__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTD X0, Y0
    VMOVUPS _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>(SB), Y15
    VPAND Y15, Y0, Y0
    VPCMPEQD Y15, Y0, Y0
    VPCMPEQD Y15, Y15, Y15
    VXORPS Y15, Y0, Y0
    VMOVMSKPS Y0, AX
    MOVB AX, ret+8(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_hdd4088432312936781d8012b60a2da7f_sk32x8_i_bitnot_constant_0<>(SB), 24, $32

TEXT ·_goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor(SB), 4, $0-9
    MOVBLZX self__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTD X0, Y0
    VMOVUPS _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>(SB), Y15
    VPAND Y15, Y0, Y0
    VPCMPEQD Y15, Y0, Y0
    MOVBLZX other__0+1(FP), AX
    MOVQ AX, X1
    VPBROADCASTD X1, Y1
    VMOVUPS _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>(SB), Y15
    VPAND Y15, Y1, Y1
    VPCMPEQD Y15, Y1, Y1
    VORPS Y1, Y0, Y0
    VMOVMSKPS Y0, AX
    MOVB AX, ret+8(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_0<>(SB), 24, $32
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_h996a1e133a8266bc7f98c019cf31874e_ask32x8_i_bitor_constant_1<>(SB), 24, $32

TEXT ·_goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor(SB), 4, $0-9
    MOVBLZX self__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTD X0, Y0
    VMOVUPS _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>(SB), Y15
    VPAND Y15, Y0, Y0
    VPCMPEQD Y15, Y0, Y0
    MOVBLZX other__0+1(FP), AX
    MOVQ AX, X1
    VPBROADCASTD X1, Y1
    VMOVUPS _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>(SB), Y15
    VPAND Y15, Y1, Y1
    VPCMPEQD Y15, Y1, Y1
    VXORPS Y1, Y0, Y0
    VMOVMSKPS Y0, AX
    MOVB AX, ret+8(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_0<>(SB), 24, $32
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_hf3b459cf4c540ececd9cf68a703d53c8_sk32x8_i_bitxor_constant_1<>(SB), 24, $32

TEXT ·_goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any(SB), 4, $0-9
    MOVBLZX self__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTD X0, Y0
    VMOVUPS _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>(SB), Y15
    VPAND Y15, Y0, Y0
    VPCMPEQD Y15, Y0, Y0
    VMOVMSKPS Y0, AX
    TESTL AX, AX
    SETNE AL
    MOVB AX, ret+8(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_inhere_h4eea6dab7814e00f769953c5cfb8a238__mask32x8_i_any_constant_0<>(SB), 24, $32

TEXT ·_goml_simd_avx2_supported(SB), 4, $0-1
    MOVB $0, ret+0(FP)
    XORL AX, AX
    CPUID
    CMPL AX, $7
    JCS unsupported
    MOVL $1, AX
    CPUID
    ANDL $0x1c000000, CX
    CMPL CX, $0x1c000000
    JNE unsupported
    XORL CX, CX
    XGETBV
    ANDL $6, AX
    CMPL AX, $6
    JNE unsupported
    MOVL $7, AX
    XORL CX, CX
    CPUID
    TESTL $0x20, BX
    JE unsupported
    MOVB $1, ret+0(FP)
unsupported:
    RET
