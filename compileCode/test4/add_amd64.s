TEXT ·max+0(SB),$0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    CMPQ AX, BX
    JLS cmp_else
    MOVQ AX, ret+16(FP)
    RET
    cmp_else:
    MOVQ BX, ret+16(FP)
    RET
