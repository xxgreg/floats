
// func {{.Name}}(dst, a, b, c, d *float64, n int)
// {{.Doc}}
TEXT ·{{.Name}}(SB), NOSPLIT|NOFRAME, $0-48
        MOVD    dst(FP), R1
        MOVD    a+8(FP), R2
        MOVD    b+16(FP), R3
        MOVD    c+24(FP), R4
        MOVD    d+32(FP), R5
        MOVD    n+40(FP), R6

        LSL     $3, R6, R7
        ADD     R1, R7, R7            // end = dst + n*8

Loop:

        VLD1.P (R2), [V0.D2, V1.D2, V2.D2, V3.D2]
        VLD1.P (R3), [V4.D2, V5.D2, V6.D2, V7.D2]
        VLD1.P (R4), [V8.D2, V9.D2, V10.D2, V11.D2]
        VLD1.P (R5), [V12.D2, V13.D2, V14.D2, V15.D2]
{{.Body}}
        VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R1)

        CMP     R7, R1                  // if dst >= end goto Done
        BGE     Done
        JMP     Loop

Done:
    RET
