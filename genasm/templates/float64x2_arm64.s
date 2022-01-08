
// func {{.Name}}(dst, a, b *float64, n int)
// {{.Doc}}
TEXT ·{{.Name}}(SB), NOSPLIT|NOFRAME, $0-32

        MOVD    dst(FP), R1
        MOVD    a+8(FP), R2
        MOVD    b+16(FP), R3
        MOVD    n+24(FP), R4

        LSL     $3, R4, R5
        ADD     R1, R5, R5            // end = dst + n*8

Loop:

        // Arm64 Neon SIMD uses 2xfloat64 in 128bit registers. These are unrolled 4x.
        VLD1.P (R2), [V0.D2, V1.D2, V2.D2, V3.D2]
        VLD1.P (R3), [V4.D2, V5.D2, V6.D2, V7.D2]
{{.Body}}
        VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R1)

        CMP     R5, R1                // if dst >= end goto Done
        BGE     Done
        JMP     Loop

Done:
    RET
