"".Add STEXT nosplit size=103 args=0x10 locals=0x18 funcid=0x0
	0x0000 00000 (.\main.go:3)	TEXT	"".Add(SB), NOSPLIT|ABIInternal, $24-16
	0x0000 00000 (.\main.go:3)	SUBQ	$24, SP
	0x0004 00004 (.\main.go:3)	MOVQ	BP, 16(SP)
	0x0009 00009 (.\main.go:3)	LEAQ	16(SP), BP
	0x000e 00014 (.\main.go:3)	FUNCDATA	$0, gclocals路1a65e721a2ccc325b382662e7ffee780(SB)
	0x000e 00014 (.\main.go:3)	FUNCDATA	$1, gclocals路69c1753bd5f81501d95132d08af04464(SB)
	0x000e 00014 (.\main.go:3)	FUNCDATA	$5, "".Add.arginfo1(SB)
	0x000e 00014 (.\main.go:3)	MOVQ	AX, "".a+32(SP)
	0x0013 00019 (.\main.go:3)	MOVQ	BX, "".b+40(SP)
	0x0018 00024 (.\main.go:3)	MOVQ	$0, "".~r2(SP)
	0x0020 00032 (.\main.go:4)	MOVQ	$0, "".c+8(SP)
	0x0029 00041 (.\main.go:5)	MOVQ	"".a+32(SP), CX
	0x002e 00046 (.\main.go:5)	TESTB	AL, (CX)
	0x0030 00048 (.\main.go:5)	MOVQ	(CX), CX
	0x0033 00051 (.\main.go:5)	ADDQ	"".b+40(SP), CX
	0x0038 00056 (.\main.go:5)	MOVQ	CX, "".c+8(SP)
	0x003d 00061 (.\main.go:6)	MOVQ	"".a+32(SP), CX
	0x0042 00066 (.\main.go:6)	TESTB	AL, (CX)
	0x0044 00068 (.\main.go:6)	MOVQ	$8, (CX)
	0x004b 00075 (.\main.go:7)	MOVQ	$9, "".b+40(SP)
	0x0054 00084 (.\main.go:8)	MOVQ	"".c+8(SP), AX
	0x0059 00089 (.\main.go:8)	MOVQ	AX, "".~r2(SP)
	0x005d 00093 (.\main.go:8)	MOVQ	16(SP), BP
	0x0062 00098 (.\main.go:8)	ADDQ	$24, SP
	0x0066 00102 (.\main.go:8)	RET
	0x0000 48 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 89  H...H.l$.H.l$.H.
	0x0010 44 24 20 48 89 5c 24 28 48 c7 04 24 00 00 00 00  D$ H.\$(H..$....
	0x0020 48 c7 44 24 08 00 00 00 00 48 8b 4c 24 20 84 01  H.D$.....H.L$ ..
	0x0030 48 8b 09 48 03 4c 24 28 48 89 4c 24 08 48 8b 4c  H..H.L$(H.L$.H.L
	0x0040 24 20 84 01 48 c7 01 08 00 00 00 48 c7 44 24 28  $ ..H......H.D$(
	0x0050 09 00 00 00 48 8b 44 24 08 48 89 04 24 48 8b 6c  ....H.D$.H..$H.l
	0x0060 24 10 48 83 c4 18 c3                             $.H....
"".main STEXT size=93 args=0x0 locals=0x30 funcid=0x0
	0x0000 00000 (.\main.go:11)	TEXT	"".main(SB), ABIInternal, $48-0
	0x0000 00000 (.\main.go:11)	CMPQ	SP, 16(R14)
	0x0004 00004 (.\main.go:11)	PCDATA	$0, $-2
	0x0004 00004 (.\main.go:11)	JLS	86
	0x0006 00006 (.\main.go:11)	PCDATA	$0, $-1
	0x0006 00006 (.\main.go:11)	SUBQ	$48, SP
	0x000a 00010 (.\main.go:11)	MOVQ	BP, 40(SP)
	0x000f 00015 (.\main.go:11)	LEAQ	40(SP), BP
	0x0014 00020 (.\main.go:11)	FUNCDATA	$0, gclocals路33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (.\main.go:11)	FUNCDATA	$1, gclocals路33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (.\main.go:12)	MOVQ	$0, "".num1+32(SP)
	0x001d 00029 (.\main.go:12)	MOVQ	$0, "".num2+24(SP)
	0x0026 00038 (.\main.go:13)	MOVQ	$3, "".num1+32(SP)
	0x002f 00047 (.\main.go:13)	MOVQ	$4, "".num2+24(SP)
	0x0038 00056 (.\main.go:15)	LEAQ	"".num1+32(SP), AX
	0x003d 00061 (.\main.go:15)	MOVL	$4, BX
	0x0042 00066 (.\main.go:15)	PCDATA	$1, $0
	0x0042 00066 (.\main.go:15)	CALL	"".Add(SB)
	0x0047 00071 (.\main.go:15)	MOVQ	AX, "".sum+16(SP)
	0x004c 00076 (.\main.go:17)	MOVQ	40(SP), BP
	0x0051 00081 (.\main.go:17)	ADDQ	$48, SP
	0x0055 00085 (.\main.go:17)	RET
	0x0056 00086 (.\main.go:17)	NOP
	0x0056 00086 (.\main.go:11)	PCDATA	$1, $-1
	0x0056 00086 (.\main.go:11)	PCDATA	$0, $-2
	0x0056 00086 (.\main.go:11)	CALL	runtime.morestack_noctxt(SB)
	0x005b 00091 (.\main.go:11)	PCDATA	$0, $-1
	0x005b 00091 (.\main.go:11)	JMP	0
	0x0000 49 3b 66 10 76 50 48 83 ec 30 48 89 6c 24 28 48  I;f.vPH..0H.l$(H
	0x0010 8d 6c 24 28 48 c7 44 24 20 00 00 00 00 48 c7 44  .l$(H.D$ ....H.D
	0x0020 24 18 00 00 00 00 48 c7 44 24 20 03 00 00 00 48  $.....H.D$ ....H
	0x0030 c7 44 24 18 04 00 00 00 48 8d 44 24 20 bb 04 00  .D$.....H.D$ ...
	0x0040 00 00 e8 00 00 00 00 48 89 44 24 10 48 8b 6c 24  .......H.D$.H.l$
	0x0050 28 48 83 c4 30 c3 e8 00 00 00 00 eb a3           (H..0........
	rel 67+4 t=7 "".Add+0
	rel 87+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals路1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals路69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
"".Add.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
gclocals路33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
