
// func {{.Name}}(dst, a, b, c, d *float64, n int)
// {{.Doc}}
TEXT ·{{.Name}}(SB), NOSPLIT|NOFRAME, $0-48
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
{{.Body}}
    VMOVUPD  Y0, (AX)(R8*8)      // store dst[i:i+8]
    VMOVUPD  Y1, 32(AX)(R8*8)

    ADDQ    $8, R8               // i += 8
    CMPQ    R8, R9               // if i > n-8 goto Done
    JGT     Done
    JMP     Loop

Done:
    VZEROUPPER
    RET

