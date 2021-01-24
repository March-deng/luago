package utils

import (
	"fmt"

	"github.com/March-deng/luago/pkg/vm"

	"github.com/March-deng/luago/pkg/binchunk"
)

func List(p *binchunk.Prototype) {
	printHeader(p)
	printCode(p)
	printDetail(p)
	for _, sp := range p.Protos {
		List(sp)
	}

}

func printHeader(p *binchunk.Prototype) {
	funcType := "main"
	if p.FirstLine > 0 {
		funcType = "function"
	}
	varargFlag := ""
	if p.IsVararg > 0 {
		varargFlag = "+"
	}
	fmt.Printf("\n%s <%s:%d,%d> (%d instructions)\n", funcType, p.Source, p.FirstLine, p.LastLine, len(p.Code))

	fmt.Printf("%d%s params, %d slots, %d upvalues, ", p.FixedParamsNum, varargFlag, p.MaxStackSize, len(p.Upvalues))

	fmt.Printf("%d locals, %d constants, %d functions\n", len(p.LocalVars), len(p.Constants), len(p.Protos))
}

func printCode(p *binchunk.Prototype) {
	for pc, c := range p.Code {
		line := "-"
		if len(p.LineInfo) != 0 {
			line = fmt.Sprintf("%d", p.LineInfo[pc])
		}
		i := vm.Instruction(c)
		fmt.Printf("\t%d\t[%s]\t%s \t", pc+1, line, i.OpName())
		printOperands(i)
		fmt.Printf("\n")
		//fmt.Printf("\t%d\t[%s]\t0x%08X\n", pc+1, line, c)
	}
}

func printOperands(i vm.Instruction) {
	switch i.OpMode() {
	case vm.IABC:
		a, b, c := i.ABC()
		fmt.Printf("%d", a)
		if i.BMode() != vm.OpArgN {
			if b > 0xff {
				fmt.Printf(" %d", -1-b&0xff)
			} else {
				fmt.Printf(" %d", b)
			}
		}
		if i.CMode() != vm.OpArgN {
			if c > 0xff {
				fmt.Printf(" %d", -1-c&0xff)
			} else {
				fmt.Printf(" %d", c)
			}
		}
	case vm.IABx:
		a, bx := i.ABx()
		fmt.Printf("%d", a)
		if i.BMode() == vm.OpArgK {
			fmt.Printf(" %d", -1-bx)
		} else if i.BMode() == vm.OpArgU {
			fmt.Printf(" %d", bx)
		}
	case vm.IAsBx:
		a, sbx := i.AsBx()
		fmt.Printf("%d %d", a, sbx)
	case vm.IAx:
		ax := i.Ax()
		fmt.Printf("%d", -1-ax)
	}
}
func printDetail(p *binchunk.Prototype) {
	fmt.Printf("constants (%d):\n", len(p.Constants))
	for i, k := range p.Constants {
		fmt.Printf("\t%d\t%s\n", i+1, constantsToString(k))
	}
	fmt.Printf("locals (%d):\n", len(p.LocalVars))
	for i, localVar := range p.LocalVars {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i, localVar.Name, localVar.StartPos+1, localVar.EndPos+1)
	}
	fmt.Printf("upvalues (%d):\n", len(p.Upvalues))
	for i, upval := range p.Upvalues {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i, upvalName(p, i), upval.InStack, upval.Idx)
	}
}

func constantsToString(k interface{}) string {
	switch k.(type) {
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("%v", k)
	}
}

func upvalName(p *binchunk.Prototype, idx int) string {
	if len(p.UpvalueNames) != 0 {
		return p.UpvalueNames[idx]
	}
	return "-"
}
