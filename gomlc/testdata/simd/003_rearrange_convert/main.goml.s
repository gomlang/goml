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

TEXT ·_goml_simd_avx2_words_widen(SB), 4, $0-64
    VMOVUPS a__0+0(FP), X0
    VMOVUPS b__0+16(FP), Y1
    VPMOVSXWD X0, Y2
    VPADDD Y1, Y2, Y2
    PXOR X0, X0
    VMOVUPS X2, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    POR X14, X0
    VMOVUPS X2, X14
    PSRLDQ $4, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $2, X14
    POR X14, X0
    VMOVUPS X2, X14
    PSRLDQ $8, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $4, X14
    POR X14, X0
    VMOVUPS X2, X14
    PSRLDQ $12, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $6, X14
    POR X14, X0
    VEXTRACTF128 $1, Y2, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $8, X14
    POR X14, X0
    VEXTRACTF128 $1, Y2, X14
    PSRLDQ $4, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $10, X14
    POR X14, X0
    VEXTRACTF128 $1, Y2, X14
    PSRLDQ $8, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $12, X14
    POR X14, X0
    VEXTRACTF128 $1, Y2, X14
    PSRLDQ $12, X14
    MOVQ X14, AX
    MOVL AX, AX
    ANDL $65535, AX
    MOVQ AX, X14
    PSLLDQ $14, X14
    POR X14, X0
    VMOVUPS X0, ret+48(FP)
    VZEROUPPER
    RET

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

TEXT ·_goml_simd_avx2_double_convert(SB), 4, $0-64
    VMOVUPS a__0+0(FP), X0
    VMOVUPS b__0+16(FP), Y1
    VCVTPS2PD X0, Y2
    VADDPD Y1, Y2, Y2
    VCVTPD2PSY Y2, X0
    VMOVUPS X0, ret+48(FP)
    VZEROUPPER
    RET

TEXT ·double_bits(SB), 4, $0-32
    MOVUPS a__0+0(FP), X0
    MOVUPS X0, X1
    MOVQ $9223372036854775808, AX
    MOVQ AX, X0
    PSHUFD $68, X0, X0
    XORPS X0, X1
    MOVUPS X1, X0
    MOVUPS X0, ret+16(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_u8x16_i_std_p_simd_p_u8x16_i_splat(SB), 4, $0-24
    MOVBLZX value__0+0(FP), AX
    MOVQ AX, X0
    PUNPCKLBW X0, X0
    PSHUFLW $0, X0, X0
    PSHUFD $0, X0, X0
    MOVUPS X0, ret+8(FP)
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
