TEXT ·sse_kernel(SB), 4, $0-56
    MOVUPS a__0+0(FP), X0
    MOVUPS b__0+16(FP), X1
    MOVSS scale__0+32(FP), X2
    MOVUPS X0, X3
    ADDPS X1, X3
    MOVUPS X3, X4
    MULPS X3, X4
    SUBPS X0, X4
    PSHUFD $0, X2, X2
    DIVPS X2, X4
    MOVUPS X0, X2
    CMPPS X1, X2, $1
    MOVUPS X2, X14
    MOVUPS X2, X15
    ANDPS X4, X14
    ANDNPS X1, X15
    MOVUPS X14, X0
    ORPS X15, X0
    MOVUPS X0, ret+40(FP)
    RET

TEXT ·_goml_simd_avx2_avx_kernel(SB), 4, $0-76
    VMOVUPS a__0+0(FP), Y0
    VMOVUPS b__0+32(FP), Y1
    MOVSS scale__0+64(FP), X2
    VADDPS Y1, Y0, Y3
    VMULPS Y3, Y3, Y4
    VSUBPS Y0, Y4, Y4
    VPBROADCASTD X2, Y2
    VDIVPS Y2, Y4, Y4
    VCMPPS $1, Y1, Y0, Y2
    VANDPS Y4, Y2, Y14
    VANDNPS Y1, Y2, Y15
    VORPS Y15, Y14, Y0
    VMOVUPS Y0, Y1
    VMOVUPS _goml_simd_avx2_avx_kernel_constant_0<>(SB), Y14
    VPERMPS Y1, Y14, Y15
    VADDPS Y15, Y1, Y1
    VMOVUPS _goml_simd_avx2_avx_kernel_constant_1<>(SB), Y14
    VPERMPS Y1, Y14, Y15
    VADDPS Y15, Y1, Y1
    VMOVUPS _goml_simd_avx2_avx_kernel_constant_2<>(SB), Y14
    VPERMPS Y1, Y14, Y15
    VADDPS Y15, Y1, Y1
    MOVSS X1, ret+72(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2_avx_kernel_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2_avx_kernel_constant_0<>+4(SB)/4, $0
DATA _goml_simd_avx2_avx_kernel_constant_0<>+8(SB)/4, $3
DATA _goml_simd_avx2_avx_kernel_constant_0<>+12(SB)/4, $2
DATA _goml_simd_avx2_avx_kernel_constant_0<>+16(SB)/4, $5
DATA _goml_simd_avx2_avx_kernel_constant_0<>+20(SB)/4, $4
DATA _goml_simd_avx2_avx_kernel_constant_0<>+24(SB)/4, $7
DATA _goml_simd_avx2_avx_kernel_constant_0<>+28(SB)/4, $6
GLOBL _goml_simd_avx2_avx_kernel_constant_0<>(SB), 24, $32
DATA _goml_simd_avx2_avx_kernel_constant_1<>+0(SB)/4, $2
DATA _goml_simd_avx2_avx_kernel_constant_1<>+4(SB)/4, $3
DATA _goml_simd_avx2_avx_kernel_constant_1<>+8(SB)/4, $0
DATA _goml_simd_avx2_avx_kernel_constant_1<>+12(SB)/4, $1
DATA _goml_simd_avx2_avx_kernel_constant_1<>+16(SB)/4, $6
DATA _goml_simd_avx2_avx_kernel_constant_1<>+20(SB)/4, $7
DATA _goml_simd_avx2_avx_kernel_constant_1<>+24(SB)/4, $4
DATA _goml_simd_avx2_avx_kernel_constant_1<>+28(SB)/4, $5
GLOBL _goml_simd_avx2_avx_kernel_constant_1<>(SB), 24, $32
DATA _goml_simd_avx2_avx_kernel_constant_2<>+0(SB)/4, $4
DATA _goml_simd_avx2_avx_kernel_constant_2<>+4(SB)/4, $5
DATA _goml_simd_avx2_avx_kernel_constant_2<>+8(SB)/4, $6
DATA _goml_simd_avx2_avx_kernel_constant_2<>+12(SB)/4, $7
DATA _goml_simd_avx2_avx_kernel_constant_2<>+16(SB)/4, $0
DATA _goml_simd_avx2_avx_kernel_constant_2<>+20(SB)/4, $1
DATA _goml_simd_avx2_avx_kernel_constant_2<>+24(SB)/4, $2
DATA _goml_simd_avx2_avx_kernel_constant_2<>+28(SB)/4, $3
GLOBL _goml_simd_avx2_avx_kernel_constant_2<>(SB), 24, $32

TEXT ·_goml_simd_avx2_masks(SB), 4, $0-9
    MOVBLZX a__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTD X0, Y0
    VMOVUPS _goml_simd_avx2_masks_constant_0<>(SB), Y15
    VPAND Y15, Y0, Y0
    VPCMPEQD Y15, Y0, Y0
    MOVBLZX b__0+1(FP), AX
    MOVQ AX, X1
    VPBROADCASTD X1, Y1
    VMOVUPS _goml_simd_avx2_masks_constant_1<>(SB), Y15
    VPAND Y15, Y1, Y1
    VPCMPEQD Y15, Y1, Y1
    VANDPS Y1, Y0, Y2
    VPCMPEQD Y15, Y15, Y15
    VXORPS Y15, Y0, Y0
    VORPS Y0, Y2, Y2
    VXORPS Y1, Y2, Y2
    VMOVMSKPS Y2, AX
    TESTL AX, AX
    SETNE AL
    MOVB AX, ret+8(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2_masks_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2_masks_constant_0<>+4(SB)/4, $2
DATA _goml_simd_avx2_masks_constant_0<>+8(SB)/4, $4
DATA _goml_simd_avx2_masks_constant_0<>+12(SB)/4, $8
DATA _goml_simd_avx2_masks_constant_0<>+16(SB)/4, $16
DATA _goml_simd_avx2_masks_constant_0<>+20(SB)/4, $32
DATA _goml_simd_avx2_masks_constant_0<>+24(SB)/4, $64
DATA _goml_simd_avx2_masks_constant_0<>+28(SB)/4, $128
GLOBL _goml_simd_avx2_masks_constant_0<>(SB), 24, $32
DATA _goml_simd_avx2_masks_constant_1<>+0(SB)/4, $1
DATA _goml_simd_avx2_masks_constant_1<>+4(SB)/4, $2
DATA _goml_simd_avx2_masks_constant_1<>+8(SB)/4, $4
DATA _goml_simd_avx2_masks_constant_1<>+12(SB)/4, $8
DATA _goml_simd_avx2_masks_constant_1<>+16(SB)/4, $16
DATA _goml_simd_avx2_masks_constant_1<>+20(SB)/4, $32
DATA _goml_simd_avx2_masks_constant_1<>+24(SB)/4, $64
DATA _goml_simd_avx2_masks_constant_1<>+28(SB)/4, $128
GLOBL _goml_simd_avx2_masks_constant_1<>(SB), 24, $32

TEXT ·_goml_m_trait__impl_i_std_p_si_hf1d5479f5b9171873ccf2609d58341d2__i_select__mask(SB), 4, $0-56
    MOVUPS self__0+0(FP), X0
    MOVBLZX mask__0+16(FP), AX
    MOVQ AX, X1
    PSHUFD $0, X1, X1
    MOVUPS _goml_m_trait__impl_i_std_p_si_hf1d5479f5b9171873ccf2609d58341d2__i_select__mask_constant_0<>(SB), X15
    PAND X15, X1
    PCMPEQL X15, X1
    MOVUPS other__0+20(FP), X2
    MOVUPS X1, X14
    MOVUPS X1, X15
    ANDPS X0, X14
    ANDNPS X2, X15
    MOVUPS X14, X3
    ORPS X15, X3
    MOVUPS X3, ret+40(FP)
    RET
DATA _goml_m_trait__impl_i_std_p_si_hf1d5479f5b9171873ccf2609d58341d2__i_select__mask_constant_0<>+0(SB)/4, $1
DATA _goml_m_trait__impl_i_std_p_si_hf1d5479f5b9171873ccf2609d58341d2__i_select__mask_constant_0<>+4(SB)/4, $2
DATA _goml_m_trait__impl_i_std_p_si_hf1d5479f5b9171873ccf2609d58341d2__i_select__mask_constant_0<>+8(SB)/4, $4
DATA _goml_m_trait__impl_i_std_p_si_hf1d5479f5b9171873ccf2609d58341d2__i_select__mask_constant_0<>+12(SB)/4, $8
GLOBL _goml_m_trait__impl_i_std_p_si_hf1d5479f5b9171873ccf2609d58341d2__i_select__mask_constant_0<>(SB), 24, $16

TEXT ·_goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask(SB), 4, $0-104
    VMOVUPS self__0+0(FP), Y0
    MOVBLZX mask__0+32(FP), AX
    MOVQ AX, X1
    VPBROADCASTD X1, Y1
    VMOVUPS _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>(SB), Y15
    VPAND Y15, Y1, Y1
    VPCMPEQD Y15, Y1, Y1
    VMOVUPS other__0+36(FP), Y2
    VANDPS Y0, Y1, Y14
    VANDNPS Y2, Y1, Y15
    VORPS Y15, Y14, Y3
    VMOVUPS Y3, ret+72(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>+0(SB)/4, $1
DATA _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>+4(SB)/4, $2
DATA _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>+8(SB)/4, $4
DATA _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>+12(SB)/4, $8
DATA _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>+16(SB)/4, $16
DATA _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>+20(SB)/4, $32
DATA _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>+24(SB)/4, $64
DATA _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>+28(SB)/4, $128
GLOBL _goml_simd_avx2__goml_m_trait__h99bff29e5e8961d94ea103b50e3be800__i_select__mask_constant_0<>(SB), 24, $32

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
