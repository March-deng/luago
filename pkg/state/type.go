package state

type LuaType int

const (
	LUA_TNONE LuaType = iota - 1
	LUAT_TNIL
	LUA_TBOOLEAN
	LUAT_TLIGHTUSERDATA
	LUAT_TNUMBER
	LUAT_TSTRING
	LUAT_TTABLE
	LUAT_TFUNCTION
	LUAT_TUSERDATA
	LUAT_TTHREAD
)

func typeOf(val luaValue) LuaType {
	switch val.(type) {
	case nil:
		return LUAT_TNIL
	case bool:
		return LUA_TBOOLEAN
	case int64:
		return LUAT_TNUMBER
	case float64:
		return LUAT_TNUMBER
	case string:
		return LUAT_TSTRING
	default:
		panic("todo!")
	}
}
