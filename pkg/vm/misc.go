package vm

import (
	"github.com/March-deng/luago/pkg/api"
)

func move(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Copy(b, a)
}

func jmp(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	vm.AddPc(sBx)
	if a != 0 {
		panic("todo!")
	}
}
