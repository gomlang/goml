TEXT ·_goml_m_inherent_i_std_p_simd_p_f64x2_i_std_p_simd_p_f64x2_i_splat(SB), 4, $0-24
    MOVSD value__0+0(FP), X0
    PSHUFD $68, X0, X0
    MOVUPS X0, ret+8(FP)
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_h3bdddb550d8be6e65f616b40ae26c68b_p_f64x4_i_splat(SB), 4, $0-40
    MOVSD value__0+0(FP), X0
    VPBROADCASTQ X0, Y0
    VMOVUPS Y0, ret+8(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2_fma__goml_m_in_hc4bab3c933044b7ff22af4f65b1c3002_64x4_i_mul__add(SB), 4, $0-128
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VMOVUPS addend__0+64(FP), Y2
    VMOVUPS Y0, Y3
    VFMADD213PD Y2, Y1, Y3
    VMOVUPS Y3, ret+96(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask(SB), 4, $0-104
    VMOVUPS self__0+0(FP), Y0
    MOVBLZX mask__0+32(FP), AX
    MOVQ AX, X1
    VPBROADCASTD X1, Y1
    VMOVUPS _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>(SB), Y15
    VPAND Y15, Y1, Y1
    VPCMPEQD Y15, Y1, Y1
    VMOVUPS other__0+40(FP), Y2
    VANDPS Y0, Y1, Y14
    VANDNPS Y2, Y1, Y15
    VORPS Y15, Y14, Y3
    VMOVUPS Y3, ret+72(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>+4(SB)/4, $1
DATA _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>+8(SB)/4, $2
DATA _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>+12(SB)/4, $2
DATA _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>+16(SB)/4, $4
DATA _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>+20(SB)/4, $4
DATA _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>+24(SB)/4, $8
DATA _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>+28(SB)/4, $8
GLOBL _goml_simd_avx2__goml_m_trait__he138df98c53aa348801ec381be0cda07__i_select__mask_constant_0<>(SB), 24, $32

TEXT ·double_sse(SB), 4, $0-40
    MOVUPS a__0+0(FP), X0
    MOVUPS b__0+16(FP), X1
    MOVUPS X0, X2
    MULPD X1, X2
    SUBPD X0, X2
    DIVPD X1, X2
    MOVUPS X2, X0
    PSHUFD $78, X0, X15
    ADDPD X15, X0
    MOVSD X0, ret+32(FP)
    RET

TEXT ·_goml_simd_avx2_double_select(SB), 4, $0-96
    VMOVUPS a__0+0(FP), Y0
    VMOVUPS b__0+32(FP), Y1
    VCMPPD $2, Y1, Y0, Y2
    VANDPS Y0, Y2, Y14
    VANDNPS Y1, Y2, Y15
    VORPS Y15, Y14, Y3
    VMOVUPS _goml_simd_avx2_double_select_constant_0<>(SB), Y15
    VXORPS Y15, Y3, Y2
    VMOVUPS _goml_simd_avx2_double_select_constant_1<>(SB), Y15
    VANDPS Y15, Y2, Y3
    VSQRTPD Y3, Y3
    VMOVUPS Y3, Y12
    VMOVUPS Y0, Y13
    VMINPD Y13, Y12, Y2
    VCMPPD $0, Y13, Y12, Y14
    VORPS Y13, Y12, Y15
    VXORPS Y2, Y15, Y15
    VANDPS Y14, Y15, Y15
    VXORPS Y15, Y2, Y2
    VCMPPD $3, Y13, Y13, Y14
    VXORPS Y2, Y12, Y15
    VANDPS Y14, Y15, Y15
    VXORPS Y15, Y2, Y2
    VMOVUPS Y2, Y12
    VMOVUPS Y1, Y13
    VMAXPD Y13, Y12, Y0
    VCMPPD $0, Y13, Y12, Y14
    VANDPS Y13, Y12, Y15
    VXORPS Y0, Y15, Y15
    VANDPS Y14, Y15, Y15
    VXORPS Y15, Y0, Y0
    VCMPPD $3, Y13, Y13, Y14
    VXORPS Y0, Y12, Y15
    VANDPS Y14, Y15, Y15
    VXORPS Y15, Y0, Y0
    VMOVUPS Y0, ret+64(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2_double_select_constant_0<>+0(SB)/8, $9223372036854775808
DATA _goml_simd_avx2_double_select_constant_0<>+8(SB)/8, $9223372036854775808
DATA _goml_simd_avx2_double_select_constant_0<>+16(SB)/8, $9223372036854775808
DATA _goml_simd_avx2_double_select_constant_0<>+24(SB)/8, $9223372036854775808
GLOBL _goml_simd_avx2_double_select_constant_0<>(SB), 24, $32
DATA _goml_simd_avx2_double_select_constant_1<>+0(SB)/8, $9223372036854775807
DATA _goml_simd_avx2_double_select_constant_1<>+8(SB)/8, $9223372036854775807
DATA _goml_simd_avx2_double_select_constant_1<>+16(SB)/8, $9223372036854775807
DATA _goml_simd_avx2_double_select_constant_1<>+24(SB)/8, $9223372036854775807
GLOBL _goml_simd_avx2_double_select_constant_1<>(SB), 24, $32

TEXT ·_goml_simd_avx2_rounding(SB), 4, $0-32
    VMOVUPS a__0+0(FP), X0
    VROUNDPS $9, X0, X1
    VROUNDPS $10, X0, X2
    VADDPS X2, X1, X1
    VROUNDPS $11, X0, X2
    VSUBPS X2, X1, X1
    VMOVUPS _goml_simd_avx2_rounding_constant_0<>(SB), X15
    VANDPS X15, X0, X12
    VROUNDPS $11, X12, X13
    VSUBPS X13, X12, X14
    VMOVUPS _goml_simd_avx2_rounding_constant_1<>(SB), X15
    VCMPPS $2, X14, X15, X14
    VMOVUPS _goml_simd_avx2_rounding_constant_2<>(SB), X15
    VANDPS X15, X14, X14
    VADDPS X14, X13, X2
    VMOVUPS _goml_simd_avx2_rounding_constant_3<>(SB), X15
    VANDPS X0, X15, X15
    VORPS X15, X2, X2
    VADDPS X2, X1, X1
    VROUNDPS $8, X0, X0
    VADDPS X0, X1, X1
    VMOVUPS X1, ret+16(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2_rounding_constant_0<>+0(SB)/4, $2147483647
DATA _goml_simd_avx2_rounding_constant_0<>+4(SB)/4, $2147483647
DATA _goml_simd_avx2_rounding_constant_0<>+8(SB)/4, $2147483647
DATA _goml_simd_avx2_rounding_constant_0<>+12(SB)/4, $2147483647
GLOBL _goml_simd_avx2_rounding_constant_0<>(SB), 24, $16
DATA _goml_simd_avx2_rounding_constant_1<>+0(SB)/4, $1056964608
DATA _goml_simd_avx2_rounding_constant_1<>+4(SB)/4, $1056964608
DATA _goml_simd_avx2_rounding_constant_1<>+8(SB)/4, $1056964608
DATA _goml_simd_avx2_rounding_constant_1<>+12(SB)/4, $1056964608
GLOBL _goml_simd_avx2_rounding_constant_1<>(SB), 24, $16
DATA _goml_simd_avx2_rounding_constant_2<>+0(SB)/4, $1065353216
DATA _goml_simd_avx2_rounding_constant_2<>+4(SB)/4, $1065353216
DATA _goml_simd_avx2_rounding_constant_2<>+8(SB)/4, $1065353216
DATA _goml_simd_avx2_rounding_constant_2<>+12(SB)/4, $1065353216
GLOBL _goml_simd_avx2_rounding_constant_2<>(SB), 24, $16
DATA _goml_simd_avx2_rounding_constant_3<>+0(SB)/4, $2147483648
DATA _goml_simd_avx2_rounding_constant_3<>+4(SB)/4, $2147483648
DATA _goml_simd_avx2_rounding_constant_3<>+8(SB)/4, $2147483648
DATA _goml_simd_avx2_rounding_constant_3<>+12(SB)/4, $2147483648
GLOBL _goml_simd_avx2_rounding_constant_3<>(SB), 24, $16

TEXT ·_goml_simd_fma_fma_sse(SB), 4, $0-64
    VMOVUPS a__0+0(FP), X0
    VMOVUPS b__0+16(FP), X1
    VMOVUPS c__0+32(FP), X2
    VMOVUPS X0, X3
    VFMADD213PS X2, X1, X3
    VMOVUPS X3, ret+48(FP)
    VZEROUPPER
    RET

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
TEXT ·_goml_simd_fma_supported(SB), 4, $0-1
    MOVB $0, ret+0(FP)
    XORL AX, AX
    CPUID
    CMPL AX, $1
    JCS unsupported
    MOVL $1, AX
    CPUID
    ANDL $0x1c001000, CX
    CMPL CX, $0x1c001000
    JNE unsupported
    XORL CX, CX
    XGETBV
    ANDL $6, AX
    CMPL AX, $6
    JNE unsupported
    MOVB $1, ret+0(FP)
unsupported:
    RET
