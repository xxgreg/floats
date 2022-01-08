//go:build amd64 && !noasm

#include "textflag.h"

// func add8(dst, a, b *float64, n int)
// n is the number of elements to add, and must be a multiple of 8.
TEXT ·add8(SB), NOSPLIT|NOFRAME, $0-32

    MOVQ    dst(FP), AX
    MOVQ    a+8(FP), BX
    MOVQ    b+16(FP), CX
    MOVQ    n+24(FP), DX

    MOVQ    $0, R8              // i = 0

    MOVQ    DX, R9              // n-8
    SUBQ    $8, R9

Loop:
    VMOVUPD  (BX)(R8*8), Y0
    VMOVUPD  32(BX)(R8*8), Y1

    VADDPD   (CX)(R8*8), Y0, Y0
    VADDPD   32(CX)(R8*8), Y1, Y1

    VMOVUPD  Y0, (AX)(R8*8)
    VMOVUPD  Y1, 32(AX)(R8*8)

    ADDQ    $8, R8               // i += 8
    CMPQ    R8, R9               // if i > n-8 goto Done
    JGT     Done
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

    MOVQ    $0, R8               // i = 0

    MOVQ    DX, R9               // n-8
    SUBQ    $8, R9

Loop:
    VMOVUPD  (BX)(R8*8), Y0
    VMOVUPD  32(BX)(R8*8), Y1

    VMULPD   (CX)(R8*8), Y0, Y0
    VMULPD   32(CX)(R8*8), Y1, Y1

    VMOVUPD  Y0, (AX)(R8*8)
    VMOVUPD  Y1, 32(AX)(R8*8)

    ADDQ    $8, R8               // i += 8
    CMPQ    R8, R9               // if i > n-8 goto Done
    JGT     Done
    JMP     Loop

Done:
    VZEROUPPER
    RET

// func add8_4(dst, a, b, c, d *float64, n int)
TEXT ·add8_4(SB), NOSPLIT|NOFRAME, $0-48
    MOVQ    dst(FP), AX
    MOVQ    a+8(FP), BX
    MOVQ    b+16(FP), CX
    MOVQ    c+24(FP), R10
    MOVQ    d+32(FP), R11
    MOVQ    n+40(FP), DX

    MOVQ    $0, R8               // i = 0

    MOVQ    DX, R9               // n-8
    SUBQ    $8, R9

Loop:
    VMOVUPD  (BX)(R8*8), Y0      // load a[i:i+8]
    VMOVUPD  32(BX)(R8*8), Y1

    VMOVUPD  (CX)(R8*8), Y2      // load b[i:i+8]
    VMOVUPD  32(CX)(R8*8), Y3

    VMOVUPD  (R10)(R8*8), Y4     // load c[i:i+8]
    VMOVUPD  32(R10)(R8*8), Y5

    VMOVUPD  (R11)(R8*8), Y6     // load d[i:i+8]
    VMOVUPD  32(R11)(R8*8), Y7

    VADDPD   Y0, Y2, Y0          // + b[i:i+8]
    VADDPD   Y1, Y3, Y1

    VADDPD   Y0, Y4, Y0          // + c[i:i+8]
    VADDPD   Y1, Y5, Y1

    VADDPD   Y0, Y6, Y0          // + d[i:i+8]
    VADDPD   Y1, Y7, Y1

    VMOVUPD  Y0, (AX)(R8*8)      // store dst[i:i+8]
    VMOVUPD  Y1, 32(AX)(R8*8)

    ADDQ    $8, R8               // i += 8
    CMPQ    R8, R9               // if i > n-8 goto Done
    JGT     Done
    JMP     Loop

Done:
    VZEROUPPER
    RET