
// func {{.Name}}(dst, a, b *float64, n int)
// {{.Doc}}
TEXT ·{{.Name}}(SB), NOSPLIT|NOFRAME, $0-32

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
{{.Body}}
    VMOVUPD  Y0, (AX)(R8*8)
    VMOVUPD  Y1, 32(AX)(R8*8)

    ADDQ    $8, R8               // i += 8
    CMPQ    R8, R9               // if i > n-8 goto Done
    JGT     Done
    JMP     Loop

Done:
    VZEROUPPER
    RET

