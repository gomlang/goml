TEXT ·_goml_simd_avx2__goml_m_inhere_h4c1feee106c01c1c4f6820cb164c4b5a_f32x4_i_to__f64(SB), 4, $0-48
    VMOVUPS self__0+0(FP), X0
    VCVTPS2PD X0, Y1
    VMOVUPS Y1, ret+16(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_hbf308f0da68aabac491becc7e3ab5729_p_f32x8_i_splat(SB), 4, $0-40
    MOVSS value__0+0(FP), X0
    VPBROADCASTD X0, Y0
    VMOVUPS Y0, ret+8(FP)
    VZEROUPPER
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_f64x2_i_std_p_simd_p_f64x2_i_to__bits(SB), 4, $0-32
    MOVUPS self__0+0(FP), X0
    MOVUPS X0, X1
    MOVUPS X1, ret+16(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_f64x2_i_std_p_simd_p_f64x2_i_from__bits(SB), 4, $0-32
    MOVUPS value__0+0(FP), X0
    MOVUPS X0, X1
    MOVUPS X1, ret+16(FP)
    RET

TEXT ·_goml_simd_avx2__goml_m_inherent_i_std_p_simd_p_f64x4_i_std_p_simd_p_f64x4_i_add(SB), 4, $0-96
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VADDPD Y1, Y0, Y0
    VMOVUPS Y0, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_hc1a8801d14877eba1f367724106405e6_f64x4_i_to__f32(SB), 4, $0-48
    VMOVUPS self__0+0(FP), Y0
    VCVTPD2PSY Y0, X1
    VMOVUPS X1, ret+32(FP)
    VZEROUPPER
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_u8x16_i_std_p_simd_p_u8x16_i_splat(SB), 4, $0-24
    MOVBLZX value__0+0(FP), AX
    MOVQ AX, X0
    PUNPCKLBW X0, X0
    PSHUFLW $0, X0, X0
    PSHUFD $0, X0, X0
    MOVUPS X0, ret+8(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_i16x8_i_std_p_simd_p_i16x8_i_splat(SB), 4, $0-24
    MOVWLZX value__0+0(FP), AX
    MOVQ AX, X0
    PSHUFLW $0, X0, X0
    PSHUFD $0, X0, X0
    MOVUPS X0, ret+8(FP)
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_h4f568bffef1abd73c3ab50c32ba9ae51_i16x8_i_to__i32(SB), 4, $0-48
    VMOVUPS self__0+0(FP), X0
    VPMOVSXWD X0, Y1
    VMOVUPS Y1, ret+16(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_h7e54adbd4efb4e3a51b52447f975fdf6_p_i32x8_i_splat(SB), 4, $0-40
    MOVL value__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTD X0, Y0
    VMOVUPS Y0, ret+8(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inherent_i_std_p_simd_p_i32x8_i_std_p_simd_p_i32x8_i_add(SB), 4, $0-96
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VPADDD Y1, Y0, Y0
    VMOVUPS Y0, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_hea4a7d4a5b82e3193a5f8fcb19c2a782_i32x8_i_to__i16(SB), 4, $0-48
    VMOVUPS self__0+0(FP), Y0
    PXOR X1, X1
    VMOVUPS X0, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    POR X14, X1
    VMOVUPS X0, X14
    PSRLDQ $4, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $2, X14
    POR X14, X1
    VMOVUPS X0, X14
    PSRLDQ $8, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $4, X14
    POR X14, X1
    VMOVUPS X0, X14
    PSRLDQ $12, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $6, X14
    POR X14, X1
    VEXTRACTF128 $1, Y0, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $8, X14
    POR X14, X1
    VEXTRACTF128 $1, Y0, X14
    PSRLDQ $4, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $10, X14
    POR X14, X1
    VEXTRACTF128 $1, Y0, X14
    PSRLDQ $8, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $12, X14
    POR X14, X1
    VEXTRACTF128 $1, Y0, X14
    PSRLDQ $12, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $14, X14
    POR X14, X1
    VMOVUPS X1, ret+32(FP)
    VZEROUPPER
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_splat(SB), 4, $0-24
    MOVQ value__0+0(FP), AX
    MOVQ AX, X0
    PSHUFD $68, X0, X0
    MOVUPS X0, ret+8(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_bitxor(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    XORPS X1, X0
    MOVUPS X0, ret+32(FP)
    RET

TEXT ·bytes_reverse(SB), 4, $0-48
    MOVUPS a__0+0(FP), X0
    MOVUPS b__0+16(FP), X1
    PXOR X2, X2
    MOVUPS X0, X14
    PSRLDQ $15, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $14, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $1, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $13, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $2, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $12, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $3, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $11, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $4, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $10, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $5, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $9, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $6, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $8, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $7, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $7, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $8, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $6, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $9, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $5, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $10, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $4, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $11, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $3, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $12, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $2, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $13, X14
    POR X14, X2
    MOVUPS X0, X14
    PSRLDQ $1, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $14, X14
    POR X14, X2
    MOVUPS X0, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $15, X14
    POR X14, X2
    PXOR X0, X0
    MOVUPS X2, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    POR X14, X0
    MOVUPS X1, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $1, X14
    POR X14, X0
    MOVUPS X2, X14
    PSRLDQ $1, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $2, X14
    POR X14, X0
    MOVUPS X1, X14
    PSRLDQ $1, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $3, X14
    POR X14, X0
    MOVUPS X2, X14
    PSRLDQ $2, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $4, X14
    POR X14, X0
    MOVUPS X1, X14
    PSRLDQ $2, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $5, X14
    POR X14, X0
    MOVUPS X2, X14
    PSRLDQ $3, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $6, X14
    POR X14, X0
    MOVUPS X1, X14
    PSRLDQ $3, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $7, X14
    POR X14, X0
    MOVUPS X2, X14
    PSRLDQ $4, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $8, X14
    POR X14, X0
    MOVUPS X1, X14
    PSRLDQ $4, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $9, X14
    POR X14, X0
    MOVUPS X2, X14
    PSRLDQ $5, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $10, X14
    POR X14, X0
    MOVUPS X1, X14
    PSRLDQ $5, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $11, X14
    POR X14, X0
    MOVUPS X2, X14
    PSRLDQ $6, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $12, X14
    POR X14, X0
    MOVUPS X1, X14
    PSRLDQ $6, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $13, X14
    POR X14, X0
    MOVUPS X2, X14
    PSRLDQ $7, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $14, X14
    POR X14, X0
    MOVUPS X1, X14
    PSRLDQ $7, X14
    MOVQ X14, AX
    ANDL $255, AX
    MOVQ AX, X14
    PSLLDQ $15, X14
    POR X14, X0
    MOVUPS X0, ret+32(FP)
    RET

TEXT ·_goml_simd_avx2_floats_permute(SB), 4, $0-64
    VMOVUPS a__0+0(FP), Y0
    VMOVUPS _goml_simd_avx2_floats_permute_constant_0<>(SB), Y14
    VPERMPS Y0, Y14, Y1
    VMOVUPS Y1, ret+32(FP)
    VZEROUPPER
    RET
DATA _goml_simd_avx2_floats_permute_constant_0<>+0(SB)/4, $7
DATA _goml_simd_avx2_floats_permute_constant_0<>+4(SB)/4, $1
DATA _goml_simd_avx2_floats_permute_constant_0<>+8(SB)/4, $6
DATA _goml_simd_avx2_floats_permute_constant_0<>+12(SB)/4, $2
DATA _goml_simd_avx2_floats_permute_constant_0<>+16(SB)/4, $5
DATA _goml_simd_avx2_floats_permute_constant_0<>+20(SB)/4, $3
DATA _goml_simd_avx2_floats_permute_constant_0<>+24(SB)/4, $4
DATA _goml_simd_avx2_floats_permute_constant_0<>+28(SB)/4, $0
GLOBL _goml_simd_avx2_floats_permute_constant_0<>(SB), 24, $32

TEXT ·numbers_convert(SB), 4, $0-32
    MOVUPS a__0+0(FP), X0
    CVTTPS2PL X0, X1
    MOVUPS numbers_convert_constant_0<>(SB), X14
    CMPPS X0, X14, $2
    MOVUPS numbers_convert_constant_1<>(SB), X15
    XORPS X1, X15
    ANDPS X14, X15
    XORPS X15, X1
    MOVUPS X0, X14
    CMPPS X0, X14, $3
    ANDNPS X1, X14
    MOVUPS X14, X1
    CVTPL2PS X1, X0
    MOVUPS X0, ret+16(FP)
    RET
DATA numbers_convert_constant_0<>+0(SB)/4, $1325400064
DATA numbers_convert_constant_0<>+4(SB)/4, $1325400064
DATA numbers_convert_constant_0<>+8(SB)/4, $1325400064
DATA numbers_convert_constant_0<>+12(SB)/4, $1325400064
GLOBL numbers_convert_constant_0<>(SB), 24, $16
DATA numbers_convert_constant_1<>+0(SB)/4, $2147483647
DATA numbers_convert_constant_1<>+4(SB)/4, $2147483647
DATA numbers_convert_constant_1<>+8(SB)/4, $2147483647
DATA numbers_convert_constant_1<>+12(SB)/4, $2147483647
GLOBL numbers_convert_constant_1<>(SB), 24, $16

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
