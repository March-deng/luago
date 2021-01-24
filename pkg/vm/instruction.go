package vm

type Instruction uint32

const (
	MAXARG_Bx  = 1<<18 - 1
	MAXARG_sBx = MAXARG_Bx >> 1
)

func (i Instruction) Opcode() int {
	return int(i & 0x3f)
}

func (i Instruction) ABC() (a, b, c int) {
	a = int(i >> 6 & 0xff)
	b = int(i >> 14 & 0x1ff)
	c = int(i >> 23 & 0x1ff)
	return
}

func (i Instruction) ABx() (a, bx int) {
	a = int(i >> 6 & 0xff)
	bx = int(i >> 14)
	return
}

func (i Instruction) AsBx() (s, sbx int) {
	a, bx := i.ABx()
	return a, bx - MAXARG_sBx
}

func (i Instruction) Ax() int {
	return int(i >> 6)
}

func (i Instruction) OpName() string {
	return opcodes[i.Opcode()].name
}

func (i Instruction) OpMode() byte {
	return opcodes[i.Opcode()].opMode
}

func (i Instruction) BMode() byte {
	return opcodes[i.Opcode()].argBMode
}

func (i Instruction) CMode() byte {
	return opcodes[i.Opcode()].argCMode
}
