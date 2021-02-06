package api

type LuaVM interface {
	LuaState
	//Program counter
	PC() int
	AddPc(n int)
	Fetch() uint32
	GetConst(idx int)
	GetRK(rk int)
}
