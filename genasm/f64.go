package main

var f64asm = []AsmFn{
	{
		Name: "add8",
		Doc: `
// Element-wise addition of a and b, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x2,
		Arch:   X86,
		Body: `
    VADDPD   (CX)(R8*8), Y0, Y0
    VADDPD   32(CX)(R8*8), Y1, Y1
`,
	},
	{
		Name: "add8",
		Doc: `
// Element-wise addition of a and b, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x2,
		Arch:   Arm64,
		Body: `
        WORD   $0x4E64D400            // fadd v0.2d, v0.2d, v4.2d  See: https://armconverter.com/ LLDB BigEndian
        WORD   $0x4E65D421            // fadd v1.2d, v1.2d, v5.2d
        WORD   $0x4E66D442            // fadd v2.2d, v2.2d, v6.2d
        WORD   $0x4E67D463            // fadd v3.2d, v3.2d, v7.2d
`,
	},
	{
		Name: "sub8",
		Doc: `
// Element-wise subtraction of a and b, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x2,
		Arch:   X86,
		Body: `
    VSUBPD   (CX)(R8*8), Y0, Y0
    VSUBPD   32(CX)(R8*8), Y1, Y1
`,
	},
	{
		Name: "sub8",
		Doc: `
	// Element-wise subtraction of a and b, storing the result in dst.
	// n must be a multiple of 8.`,
		FnKind: Float64x2,
		Arch:   Arm64,
		Body: `
	       WORD   $0x4EE4D400            // fsub v0.2d, v0.2d, v4.2d
	       WORD   $0x4EE5D421            // fsub v1.2d, v1.2d, v5.2d
	       WORD   $0x4EE6D442            // fsub v2.2d, v2.2d, v6.2d
	       WORD   $0x4EE7D463            // fsub v3.2d, v3.2d, v7.2d
	`,
	},
	{
		Name: "mul8",
		Doc: `
// Element-wise multiplication of a and b, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x2,
		Arch:   X86,
		Body: `
    VMULPD   (CX)(R8*8), Y0, Y0
    VMULPD   32(CX)(R8*8), Y1, Y1
`,
	},
	{
		Name: "mul8",
		Doc: `
// Element-wise multiplication of a and b, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x2,
		Arch:   Arm64,
		Body: `
        WORD   $0x6E64DC00              // fmul v0.2d, v0.2d, v4.2d  See: https://armconverter.com/ LLDB BigEndian
        WORD   $0x6E65DC21              // fmul v1.2d, v1.2d, v5.2d
        WORD   $0x6E66DC42              // fmul v2.2d, v2.2d, v6.2d
        WORD   $0x6E67DC63              // fmul v3.2d, v3.2d, v7.2d
`,
	},
	{
		Name: "add8_4",
		Doc: `
// Element-wise addition of a, b, c and d, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x4,
		Arch:   X86,
		Body: `
    VADDPD   Y0, Y2, Y0          // + b[i:i+8]
    VADDPD   Y1, Y3, Y1

    VADDPD   Y0, Y4, Y0          // + c[i:i+8]
    VADDPD   Y1, Y5, Y1

    VADDPD   Y0, Y6, Y0          // + d[i:i+8]
    VADDPD   Y1, Y7, Y1
`,
	},
	{
		Name: "add8_4",
		Doc: `
// Element-wise addition of a, b, c and d, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x4,
		Arch:   Arm64,
		Body: `
        WORD   $0x4E64D400            // fadd v0.2d, v0.2d, v4.2d
        WORD   $0x4E65D421            // fadd v1.2d, v1.2d, v5.2d
        WORD   $0x4E66D442            // fadd v2.2d, v2.2d, v6.2d
        WORD   $0x4E67D463            // fadd v3.2d, v3.2d, v7.2d

        WORD   $0x4E68D400            // fadd v0.2d, v0.2d, v8.2d
        WORD   $0x4E69D421            // fadd v1.2d, v1.2d, v9.2d
        WORD   $0x4E6AD442            // fadd v2.2d, v2.2d, v10.2d
        WORD   $0x4E6BD463            // fadd v3.2d, v3.2d, v11.2d

        WORD   $0x4E6CD400            // fadd v0.2d, v0.2d, v12.2d
        WORD   $0x4E6DD421            // fadd v1.2d, v1.2d, v13.2d
        WORD   $0x4E6ED442            // fadd v2.2d, v2.2d, v14.2d
        WORD   $0x4E6FD463            // fadd v3.2d, v3.2d, v15.2d
`,
	},
	{
		Name: "mul8_4",
		Doc: `
// Element-wise multiplication of a, b, c and d, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x4,
		Arch:   X86,
		Body: `
    VMULPD   Y0, Y2, Y0          // + b[i:i+8]
    VMULPD   Y1, Y3, Y1

    VMULPD   Y0, Y4, Y0          // + c[i:i+8]
    VMULPD   Y1, Y5, Y1

    VMULPD   Y0, Y6, Y0          // + d[i:i+8]
    VMULPD   Y1, Y7, Y1
`,
	},
	{
		Name: "mul8_4",
		Doc: `
// Element-wise addition of a, b, c and d, storing the result in dst.
// n must be a multiple of 8.`,
		FnKind: Float64x4,
		Arch:   Arm64,
		Body: `
    WORD $0x6E64DC00               // fmul v0.2d, v0.2d, v4.2d
    WORD $0x6E65DC21               // fmul v1.2d, v1.2d, v5.2d
    WORD $0x6E66DC42               // fmul v2.2d, v2.2d, v6.2d
    WORD $0x6E67DC63               // fmul v3.2d, v3.2d, v7.2d

    WORD $0x6E68DC00               // fmul v0.2d, v0.2d, v8.2d
    WORD $0x6E69DC21               // fmul v1.2d, v1.2d, v9.2d
    WORD $0x6E6ADC42               // fmul v2.2d, v2.2d, v10.2d
    WORD $0x6E6BDC63               // fmul v3.2d, v3.2d, v11.2d

    WORD $0x6E6CDC00               // fmul v0.2d, v0.2d, v12.2d
    WORD $0x6E6DDC21               // fmul v1.2d, v1.2d, v13.2d
    WORD $0x6E6EDC42               // fmul v2.2d, v2.2d, v14.2d
    WORD $0x6E6FDC63               // fmul v3.2d, v3.2d, v15.2d
`,
	},
}
