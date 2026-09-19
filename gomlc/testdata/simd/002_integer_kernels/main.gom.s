TEXT ·_goml_m_inherent_i_std_p_simd_p_u8x16_i_std_p_simd_p_u8x16_i_splat(SB), 4, $0-24
    MOVBLZX value__0+0(FP), AX
    MOVQ AX, X0
    PUNPCKLBW X0, X0
    PSHUFLW $0, X0, X0
    PSHUFD $0, X0, X0
    MOVUPS X0, ret+8(FP)
    RET

TEXT ·_goml_m_trait__impl_i_std_p_si_ha54a999764fc0f89c31804f55a4f1184__i_select__mask(SB), 4, $0-56
    MOVUPS self__0+0(FP), X0
    MOVWLZX mask__0+16(FP), AX
    XORQ R8, R8
    MOVQ AX, BX
    SHRQ $0, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $0, BX
    ORQ BX, R8
    MOVQ AX, BX
    SHRQ $1, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $8, BX
    ORQ BX, R8
    MOVQ AX, BX
    SHRQ $2, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $16, BX
    ORQ BX, R8
    MOVQ AX, BX
    SHRQ $3, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $24, BX
    ORQ BX, R8
    MOVQ AX, BX
    SHRQ $4, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $32, BX
    ORQ BX, R8
    MOVQ AX, BX
    SHRQ $5, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $40, BX
    ORQ BX, R8
    MOVQ AX, BX
    SHRQ $6, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $48, BX
    ORQ BX, R8
    MOVQ AX, BX
    SHRQ $7, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $56, BX
    ORQ BX, R8
    XORQ R9, R9
    MOVQ AX, BX
    SHRQ $8, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $0, BX
    ORQ BX, R9
    MOVQ AX, BX
    SHRQ $9, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $8, BX
    ORQ BX, R9
    MOVQ AX, BX
    SHRQ $10, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $16, BX
    ORQ BX, R9
    MOVQ AX, BX
    SHRQ $11, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $24, BX
    ORQ BX, R9
    MOVQ AX, BX
    SHRQ $12, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $32, BX
    ORQ BX, R9
    MOVQ AX, BX
    SHRQ $13, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $40, BX
    ORQ BX, R9
    MOVQ AX, BX
    SHRQ $14, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $48, BX
    ORQ BX, R9
    MOVQ AX, BX
    SHRQ $15, BX
    ANDQ $1, BX
    NEGQ BX
    ANDQ $255, BX
    SHLQ $56, BX
    ORQ BX, R9
    MOVQ R8, X1
    MOVQ R9, X15
    PUNPCKLQDQ X15, X1
    MOVUPS other__0+18(FP), X2
    MOVUPS X1, X14
    MOVUPS X1, X15
    ANDPS X0, X14
    ANDNPS X2, X15
    MOVUPS X14, X3
    ORPS X15, X3
    MOVUPS X3, ret+40(FP)
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_h798988c669064201ad8c97c85c826892__i16x16_i_splat(SB), 4, $0-40
    MOVWLZX value__0+0(FP), AX
    MOVQ AX, X0
    VPBROADCASTW X0, Y0
    VMOVUPS Y0, ret+8(FP)
    VZEROUPPER
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_i32x4_i_std_p_simd_p_i32x4_i_mul(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    MOVUPS X0, X14
    PMULULQ X1, X14
    MOVUPS X0, X15
    PSRLQ $32, X15
    MOVUPS X1, X2
    PSRLQ $32, X2
    PMULULQ X15, X2
    PSHUFD $136, X14, X14
    PSHUFD $136, X2, X2
    PUNPCKLLQ X2, X14
    MOVUPS X14, X2
    MOVUPS X2, ret+32(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_i32x4_i_std_p_simd_p_i32x4_i_min(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    MOVUPS X1, X2
    PCMPGTL X0, X2
    MOVUPS X2, X14
    MOVUPS X2, X15
    ANDPS X0, X14
    ANDNPS X1, X15
    MOVUPS X14, X2
    ORPS X15, X2
    MOVUPS X2, ret+32(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_i32x4_i_std_p_simd_p_i32x4_i_bitnot(SB), 4, $0-32
    MOVUPS self__0+0(FP), X0
    PCMPEQL X15, X15
    XORPS X15, X0
    MOVUPS X0, ret+16(FP)
    RET

TEXT ·_goml_simd_avx2__goml_m_inherent_i_std_p_simd_p_i64x4_i_std_p_simd_p_i64x4_i_mul(SB), 4, $0-96
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS other__0+32(FP), Y1
    VPSRLQ $32, Y0, Y14
    VPMULUDQ Y1, Y14, Y14
    VPSRLQ $32, Y1, Y15
    VPMULUDQ Y0, Y15, Y15
    VPADDQ Y15, Y14, Y14
    VPSLLQ $32, Y14, Y14
    VPMULUDQ Y1, Y0, Y2
    VPADDQ Y14, Y2, Y2
    VMOVUPS Y2, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_simd_avx2__goml_m_inhere_ha0b334929cdf9bd4a4bd368fafaf5cc7_4_i_reduce__sum(SB), 4, $0-40
    VMOVUPS self__0+0(FP), Y0
    VMOVUPS X0, X14
    MOVQ X14, AX
    MOVQ AX, R8
    VMOVUPS X0, X14
    PSRLDQ $8, X14
    MOVQ X14, AX
    ADDQ AX, R8
    VEXTRACTF128 $1, Y0, X14
    MOVQ X14, AX
    ADDQ AX, R8
    VEXTRACTF128 $1, Y0, X14
    PSRLDQ $8, X14
    MOVQ X14, AX
    ADDQ AX, R8
    MOVQ R8, X1
    MOVQ X1, AX
    MOVQ AX, ret+32(FP)
    VZEROUPPER
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_mul(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    MOVUPS X0, X14
    PSRLQ $32, X14
    PMULULQ X1, X14
    MOVUPS X1, X15
    PSRLQ $32, X15
    PMULULQ X0, X15
    PADDQ X15, X14
    PSLLQ $32, X14
    MOVUPS X0, X2
    PMULULQ X1, X2
    PADDQ X14, X2
    MOVUPS X2, ret+32(FP)
    RET

TEXT ·_goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_max(SB), 4, $0-48
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    MOVUPS _goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_max_constant_0<>(SB), X14
    MOVUPS X0, X15
    XORPS X14, X15
    MOVUPS X1, X2
    XORPS X14, X2
    PCMPGTL X15, X2
    PSHUFD $160, X2, X14
    PSHUFD $245, X2, X2
    MOVUPS X0, X15
    PCMPEQL X1, X15
    PSHUFD $245, X15, X15
    PAND X15, X14
    POR X14, X2
    MOVUPS X2, X14
    MOVUPS X2, X15
    ANDPS X1, X14
    ANDNPS X0, X15
    MOVUPS X14, X2
    ORPS X15, X2
    MOVUPS X2, ret+32(FP)
    RET
DATA _goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_max_constant_0<>+0(SB)/8, $9223372039002259456
DATA _goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_max_constant_0<>+8(SB)/8, $9223372039002259456
GLOBL _goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_max_constant_0<>(SB), 24, $16

TEXT ·_goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_simd__le(SB), 4, $0-33
    MOVUPS self__0+0(FP), X0
    MOVUPS other__0+16(FP), X1
    MOVUPS _goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_simd__le_constant_0<>(SB), X14
    MOVUPS X1, X15
    XORPS X14, X15
    MOVUPS X0, X2
    XORPS X14, X2
    PCMPGTL X15, X2
    PSHUFD $160, X2, X14
    PSHUFD $245, X2, X2
    MOVUPS X1, X15
    PCMPEQL X0, X15
    PSHUFD $245, X15, X15
    PAND X15, X14
    POR X14, X2
    PCMPEQL X15, X15
    XORPS X15, X2
    MOVMSKPD X2, AX
    MOVB AX, ret+32(FP)
    RET
DATA _goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_simd__le_constant_0<>+0(SB)/8, $9223372039002259456
DATA _goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_simd__le_constant_0<>+8(SB)/8, $9223372039002259456
GLOBL _goml_m_inherent_i_std_p_simd_p_u64x2_i_std_p_simd_p_u64x2_i_simd__le_constant_0<>(SB), 24, $16

TEXT ·bytes_sse(SB), 4, $0-48
    MOVUPS a__0+0(FP), X0
    MOVUPS b__0+16(FP), X1
    MOVUPS bytes_sse_constant_0<>(SB), X14
    MOVUPS X1, X15
    XORPS X14, X15
    MOVUPS X0, X2
    XORPS X14, X2
    PCMPGTB X15, X2
    MOVUPS X0, X14
    PSRLW $8, X14
    MOVUPS X1, X15
    PSRLW $8, X15
    PMULLW X15, X14
    PSLLW $8, X14
    MOVUPS X0, X3
    PMULLW X1, X3
    MOVUPS bytes_sse_constant_1<>(SB), X15
    ANDPS X15, X3
    ORPS X14, X3
    MOVUPS X3, X4
    PADDUSB X1, X4
    MOVUPS X4, X3
    PSRLW $3, X3
    MOVUPS bytes_sse_constant_2<>(SB), X15
    ANDPS X15, X3
    XORPS X1, X0
    MOVUPS X2, X14
    MOVUPS X2, X15
    ANDPS X3, X14
    ANDNPS X0, X15
    MOVUPS X14, X1
    ORPS X15, X1
    MOVUPS X1, ret+32(FP)
    RET
DATA bytes_sse_constant_0<>+0(SB)/1, $128
DATA bytes_sse_constant_0<>+1(SB)/1, $128
DATA bytes_sse_constant_0<>+2(SB)/1, $128
DATA bytes_sse_constant_0<>+3(SB)/1, $128
DATA bytes_sse_constant_0<>+4(SB)/1, $128
DATA bytes_sse_constant_0<>+5(SB)/1, $128
DATA bytes_sse_constant_0<>+6(SB)/1, $128
DATA bytes_sse_constant_0<>+7(SB)/1, $128
DATA bytes_sse_constant_0<>+8(SB)/1, $128
DATA bytes_sse_constant_0<>+9(SB)/1, $128
DATA bytes_sse_constant_0<>+10(SB)/1, $128
DATA bytes_sse_constant_0<>+11(SB)/1, $128
DATA bytes_sse_constant_0<>+12(SB)/1, $128
DATA bytes_sse_constant_0<>+13(SB)/1, $128
DATA bytes_sse_constant_0<>+14(SB)/1, $128
DATA bytes_sse_constant_0<>+15(SB)/1, $128
GLOBL bytes_sse_constant_0<>(SB), 24, $16
DATA bytes_sse_constant_1<>+0(SB)/2, $255
DATA bytes_sse_constant_1<>+2(SB)/2, $255
DATA bytes_sse_constant_1<>+4(SB)/2, $255
DATA bytes_sse_constant_1<>+6(SB)/2, $255
DATA bytes_sse_constant_1<>+8(SB)/2, $255
DATA bytes_sse_constant_1<>+10(SB)/2, $255
DATA bytes_sse_constant_1<>+12(SB)/2, $255
DATA bytes_sse_constant_1<>+14(SB)/2, $255
GLOBL bytes_sse_constant_1<>(SB), 24, $16
DATA bytes_sse_constant_2<>+0(SB)/1, $31
DATA bytes_sse_constant_2<>+1(SB)/1, $31
DATA bytes_sse_constant_2<>+2(SB)/1, $31
DATA bytes_sse_constant_2<>+3(SB)/1, $31
DATA bytes_sse_constant_2<>+4(SB)/1, $31
DATA bytes_sse_constant_2<>+5(SB)/1, $31
DATA bytes_sse_constant_2<>+6(SB)/1, $31
DATA bytes_sse_constant_2<>+7(SB)/1, $31
DATA bytes_sse_constant_2<>+8(SB)/1, $31
DATA bytes_sse_constant_2<>+9(SB)/1, $31
DATA bytes_sse_constant_2<>+10(SB)/1, $31
DATA bytes_sse_constant_2<>+11(SB)/1, $31
DATA bytes_sse_constant_2<>+12(SB)/1, $31
DATA bytes_sse_constant_2<>+13(SB)/1, $31
DATA bytes_sse_constant_2<>+14(SB)/1, $31
DATA bytes_sse_constant_2<>+15(SB)/1, $31
GLOBL bytes_sse_constant_2<>(SB), 24, $16

TEXT ·_goml_simd_avx2_words_avx(SB), 4, $0-96
    VMOVUPS a__0+0(FP), Y0
    VMOVUPS b__0+32(FP), Y1
    VPSUBSW Y1, Y0, Y2
    VPMULLW Y1, Y2, Y0
    VPSRAW $15, Y0, Y15
    VXORPS Y15, Y0, Y1
    VPSUBW Y15, Y1, Y1
    VPSLLW $2, Y1, Y0
    VMOVUPS Y0, ret+64(FP)
    VZEROUPPER
    RET

TEXT ·_goml_m_inherent_i_std_p_simd__h5b1c4d30d5d63db324f03c858eef8d8f__p_simd_p_u64x2(SB), 4, $0-56
    MOVBLZX self__0+0(FP), AX
    MOVQ AX, X0
    PSHUFD $0, X0, X0
    MOVUPS _goml_m_inherent_i_std_p_simd__h5b1c4d30d5d63db324f03c858eef8d8f__p_simd_p_u64x2_constant_0<>(SB), X15
    PAND X15, X0
    PCMPEQL X15, X0
    MOVUPS if_true__0+8(FP), X1
    MOVUPS if_false__0+24(FP), X2
    MOVUPS X0, X14
    MOVUPS X0, X15
    ANDPS X1, X14
    ANDNPS X2, X15
    MOVUPS X14, X3
    ORPS X15, X3
    MOVUPS X3, ret+40(FP)
    RET
DATA _goml_m_inherent_i_std_p_simd__h5b1c4d30d5d63db324f03c858eef8d8f__p_simd_p_u64x2_constant_0<>+0(SB)/4, $1
DATA _goml_m_inherent_i_std_p_simd__h5b1c4d30d5d63db324f03c858eef8d8f__p_simd_p_u64x2_constant_0<>+4(SB)/4, $1
DATA _goml_m_inherent_i_std_p_simd__h5b1c4d30d5d63db324f03c858eef8d8f__p_simd_p_u64x2_constant_0<>+8(SB)/4, $2
DATA _goml_m_inherent_i_std_p_simd__h5b1c4d30d5d63db324f03c858eef8d8f__p_simd_p_u64x2_constant_0<>+12(SB)/4, $2
GLOBL _goml_m_inherent_i_std_p_simd__h5b1c4d30d5d63db324f03c858eef8d8f__p_simd_p_u64x2_constant_0<>(SB), 24, $16

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
