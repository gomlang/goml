//go:build !cgo && go1.26 && !go1.27

#include "textflag.h"

TEXT ·bootInit(SB),NOSPLIT|NOFRAME,$0
	PUSHQ BX
	SUBQ $80, SP
	MOVQ DI, BX
	MOVQ SI, ·setG(SB)
	MOVQ SP, DI
	CALL goml_cabi_pthread_attr_init(SB)
	TESTL AX, AX
	JNE boot_failure
	CALL goml_cabi_pthread_self(SB)
	MOVQ AX, DI
	MOVQ SP, SI
	CALL goml_cabi_pthread_getattr_np(SB)
	TESTL AX, AX
	JNE boot_failure
	MOVQ SP, DI
	LEAQ 56(SP), SI
	LEAQ 64(SP), DX
	CALL goml_cabi_pthread_attr_getstack(SB)
	TESTL AX, AX
	JNE boot_failure
	MOVQ 56(SP), AX
	MOVQ AX, 0(BX)
	MOVQ SP, DI
	CALL goml_cabi_pthread_attr_destroy(SB)
	ADDQ $80, SP
	POPQ BX
	RET
boot_failure:
	CALL goml_cabi_abort(SB)
	UD2

TEXT ·bootThreadStart(SB),NOSPLIT|NOFRAME,$0
	PUSHQ BX
	PUSHQ R12
	PUSHQ R13
	SUBQ $336, SP
	MOVQ DI, R12
	MOVQ $24, DI
	CALL goml_cabi_malloc(SB)
	TESTQ AX, AX
	JE thread_failure
	MOVQ AX, BX
	MOVQ 0(R12), AX
	MOVQ AX, 0(BX)
	MOVQ 8(R12), AX
	MOVQ AX, 8(BX)
	MOVQ 16(R12), AX
	MOVQ AX, 16(BX)
	LEAQ 64(SP), DI
	CALL goml_cabi_sigfillset(SB)
	MOVQ $2, DI
	LEAQ 64(SP), SI
	LEAQ 192(SP), DX
	CALL goml_cabi_pthread_sigmask(SB)
	TESTL AX, AX
	JNE thread_failure
	MOVQ SP, DI
	CALL goml_cabi_pthread_attr_init(SB)
	TESTL AX, AX
	JNE thread_failure
	MOVQ SP, DI
	MOVQ $1, SI
	CALL goml_cabi_pthread_attr_setdetachstate(SB)
	TESTL AX, AX
	JNE thread_failure
	MOVQ SP, DI
	LEAQ 320(SP), SI
	CALL goml_cabi_pthread_attr_getstacksize(SB)
	TESTL AX, AX
	JNE thread_failure
	MOVQ 0(BX), CX
	MOVQ 320(SP), AX
	MOVQ AX, 8(CX)
	LEAQ 328(SP), DI
	MOVQ SP, SI
	LEAQ ·threadEntry(SB), DX
	MOVQ BX, CX
	CALL goml_cabi_pthread_create(SB)
	MOVL AX, R13
	MOVQ SP, DI
	CALL goml_cabi_pthread_attr_destroy(SB)
	MOVQ $2, DI
	LEAQ 192(SP), SI
	XORQ DX, DX
	CALL goml_cabi_pthread_sigmask(SB)
	TESTL AX, AX
	JNE thread_failure
	TESTL R13, R13
	JNE thread_failure
	ADDQ $336, SP
	POPQ R13
	POPQ R12
	POPQ BX
	RET
thread_failure:
	CALL goml_cabi_abort(SB)
	UD2

TEXT ·threadEntry(SB),NOSPLIT|NOFRAME,$0
	PUSHQ BP
	PUSHQ BX
	PUSHQ R12
	PUSHQ R13
	PUSHQ R14
	PUSHQ R15
	SUBQ $8, SP
	MOVQ 0(DI), R12
	MOVQ 16(DI), BX
	CALL goml_cabi_free(SB)
	MOVQ R12, DI
	MOVQ ·setG(SB), AX
	CALL AX
	CALL BX
	XORQ AX, AX
	ADDQ $8, SP
	POPQ R15
	POPQ R14
	POPQ R13
	POPQ R12
	POPQ BX
	POPQ BP
	RET

TEXT ·bootNotify(SB),NOSPLIT|NOFRAME,$0
	XORQ AX, AX
	RET

TEXT ·bootSetenv(SB),NOSPLIT|NOFRAME,$0
	MOVQ 8(DI), SI
	MOVQ 0(DI), DI
	MOVQ $1, DX
	JMP goml_cabi_setenv(SB)

TEXT ·bootUnsetenv(SB),NOSPLIT|NOFRAME,$0
	MOVQ 0(DI), DI
	JMP goml_cabi_unsetenv(SB)

TEXT ·openABI(SB),NOSPLIT|NOFRAME,$0
	JMP goml_cabi_dlopen(SB)
TEXT ·symbolABI(SB),NOSPLIT|NOFRAME,$0
	JMP goml_cabi_dlsym(SB)
TEXT ·errorABI(SB),NOSPLIT|NOFRAME,$0
	JMP goml_cabi_dlerror(SB)
TEXT ·mallocABI(SB),NOSPLIT|NOFRAME,$0
	JMP goml_cabi_malloc(SB)
TEXT ·freeABI(SB),NOSPLIT|NOFRAME,$0
	JMP goml_cabi_free(SB)

TEXT ·invokeABI(SB),NOSPLIT|NOFRAME,$0
	PUSHQ BX
	MOVQ DI, BX
	SUBQ $1024, SP
	MOVQ 120(BX), CX
	XORQ AX, AX
copy_arguments:
	CMPQ AX, CX
	JGE arguments_ready
	MOVQ 128(BX)(AX*8), DX
	MOVQ DX, 0(SP)(AX*8)
	INCQ AX
	JMP copy_arguments
arguments_ready:
	MOVQ 8(BX), DI
	MOVQ 16(BX), SI
	MOVQ 24(BX), DX
	MOVQ 32(BX), CX
	MOVQ 40(BX), R8
	MOVQ 48(BX), R9
	MOVQ 56(BX), X0
	MOVQ 64(BX), X1
	MOVQ 72(BX), X2
	MOVQ 80(BX), X3
	MOVQ 88(BX), X4
	MOVQ 96(BX), X5
	MOVQ 104(BX), X6
	MOVQ 112(BX), X7
	MOVQ 0(BX), R11
	CALL R11
	MOVQ AX, 1152(BX)
	MOVQ X0, 1160(BX)
	XORQ AX, AX
	ADDQ $1024, SP
	POPQ BX
	RET

TEXT ·Pointer(SB),NOSPLIT,$0-16
	MOVQ value+0(FP), AX
	MOVQ AX, ret+8(FP)
	RET
