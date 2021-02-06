package state

import (
	"github.com/March-deng/luago/pkg/api"
	"github.com/March-deng/luago/pkg/binchunk"
)

//lusState is the internal implementation of LusState
type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	//current program counter
	pc int
}

var _ api.LuaVM = &luaState{}

func New(size int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(20),
		proto: proto,
	}
}
