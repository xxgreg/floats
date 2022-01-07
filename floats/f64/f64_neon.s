//go:build arm64 && !noasm

#include "textflag.h"

// func add8(dst, a, b *float64, n int64)
TEXT ·add8(SB), NOSPLIT|NOFRAME, $0-32

        MOVD    dst(FP), R1
        MOVD    a+8(FP), R2
        MOVD    b+16(FP), R3
        MOVD    n+24(FP), R4

        LSR     $3, R4, R5
        LSL     $6, R5, R5
        ADD     R1, R5, R5            // end = dst + (n/8)*8*8    which is: (n>>3)<<6

Loop:

        // Arm64 Neon SIMD uses 2xfloat64 in 128bit registers.
        VLD1.P (R2), [V0.D2, V1.D2, V2.D2, V3.D2]
        VLD1.P (R3), [V4.D2, V5.D2, V6.D2, V7.D2]
        WORD   $0x4E64D400            // fadd v0.2d, v0.2d, v4.2d  See: https://armconverter.com/ LLDB BigEndian
        WORD   $0x4E65D421            // fadd v1.2d, v1.2d, v5.2d
        WORD   $0x4E66D442            // fadd v2.2d, v2.2d, v6.2d
        WORD   $0x4E67D463            // fadd v3.2d, v3.2d, v7.2d
        VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R1)

        CMP     R5, R1                // if dst >= end goto Done
        BGE     Done
        JMP     Loop

Done:
    RET

// func mul8(dst, a, b *float64, n int64)
TEXT ·mul8(SB), NOSPLIT|NOFRAME, $0-32

        MOVD    dst(FP), R1
        MOVD    a+8(FP), R2
        MOVD    b+16(FP), R3
        MOVD    n+24(FP), R4

        LSR     $3, R4, R5
        LSL     $6, R5, R5
        ADD     R1, R5, R5              // end = dst + (n/8)*8*8    which is: (n>>3)<<6

Loop:

        // Arm64 Neon SIMD uses 2xfloat64 in 128bit registers.
        VLD1.P (R2), [V0.D2, V1.D2, V2.D2, V3.D2]
        VLD1.P (R3), [V4.D2, V5.D2, V6.D2, V7.D2]
        WORD   $0x6E64DC00              // fmul v0.2d, v0.2d, v4.2d  See: https://armconverter.com/ LLDB BigEndian
        WORD   $0x6E65DC21              // fmul v1.2d, v1.2d, v5.2d
        WORD   $0x6E66DC42              // fmul v2.2d, v2.2d, v6.2d
        WORD   $0x6E67DC63              // fmul v3.2d, v3.2d, v7.2d
        VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R1)

        CMP     R5, R1                  // if dst >= end goto Done
        BGE     Done
        JMP     Loop

Done:
    RET
