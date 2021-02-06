package api

type (
	LuaType   int
	ArithOp   int
	CompareOp int
)

func (t LuaType) String() string {
	switch t {
	case LUA_TNONE:
		return "no value"
	case LUAT_TNIL:
		return "nil"
	case LUA_TBOOLEAN:
		return "boolean"
	case LUAT_TLIGHTUSERDATA:
		return "userdata"
	case LUAT_TNUMBER:
		return "number"
	case LUAT_TSTRING:
		return "string"
	case LUAT_TTABLE:
		return "table"
	case LUAT_TFUNCTION:
		return "function"
	case LUAT_TUSERDATA:
		return "userdata"
	case LUAT_TTHREAD:
		return "thread"
	}
	panic("type invalid!")
}
