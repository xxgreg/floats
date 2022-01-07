//go:build amd64 && !noasm

#include "textflag.h"

// func add8(dst, a, b *float64, n int)
TEXT ·add8(SB), NOSPLIT|NOFRAME, $0-32
    MOVQ    dst(FP), AX
    MOVQ    a+8(FP), BX
    MOVQ    b+16(FP), CX
    MOVQ    n+24(FP), DX

    MOVQ    $0, R8              // i = 0

    MOVQ    DX, R9
    SHRQ    $3, R9
    SHLQ    $3, R9              // end = (n/8)*8   which is: (n>>3)<<3

Loop:
    VMOVAPD  (BX)(R8*8), Y0
    VMOVAPD  32(BX)(R8*8), Y1

    VADDPD   (CX)(R8*8), Y0, Y0
    VADDPD   32(CX)(R8*8), Y1, Y1

    VMOVAPD  Y0, (AX)(R8*8)
    VMOVAPD  Y1, 32(AX)(R8*8)

    ADDQ    $8, R8               // i += 8
    CMPQ    R8, R9               // if i >= end goto Done
    JGE     Done
    JMP     Loop

Done:
    VZEROUPPER
    RET

// func mul8(dst, a, b *float64, n int)
TEXT ·mul8(SB), NOSPLIT|NOFRAME, $0-32
    MOVQ    dst(FP), AX
    MOVQ    a+8(FP), BX
    MOVQ    b+16(FP), CX
    MOVQ    n+24(FP), DX

    MOVQ    $0, R8              // i = 0

    MOVQ    DX, R9
    SHRQ    $3, R9
    SHLQ    $3, R9              // end = (n/8)*8   which is: (n>>3)<<3

Loop:
    VMOVAPD  (BX)(R8*8), Y0
    VMOVAPD  32(BX)(R8*8), Y1

    VMULPD   (CX)(R8*8), Y0, Y0
    VMULPD   32(CX)(R8*8), Y1, Y1

    VMOVAPD  Y0, (AX)(R8*8)
    VMOVAPD  Y1, 32(AX)(R8*8)

    ADDQ    $8, R8               // i += 8
    CMPQ    R8, R9               // if i >= end goto Done
    JGE     Done
    JMP     Loop

Done:
    VZEROUPPER
    RET
