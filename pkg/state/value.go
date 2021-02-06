package state

import (
	"github.com/March-deng/luago/pkg/api"
	"github.com/March-deng/luago/pkg/number"
)

type luaValue interface{}

func typeOf(val luaValue) api.LuaType {
	switch val.(type) {
	case nil:
		return api.LUAT_TNIL
	case bool:
		return api.LUA_TBOOLEAN
	case int64:
		return api.LUAT_TNUMBER
	case float64:
		return api.LUAT_TNUMBER
	case string:
		return api.LUAT_TSTRING
	default:
		panic("todo!")
	}
}

func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	case string:
		return number.ParseFloat(x)
	default:
		return 0, false
	}
}

func convertToInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case int64:
		return x, true
	case float64:
		return number.FloatToInteger(x)
	case string:
		return number.ParseInteger(x)
	default:
		return 0, false
	}
}

func _stringToInteger(s string) (int64, bool) {
	if i, ok := number.ParseInteger(s); ok {
		return i, ok
	}
	if f, ok := number.ParseFloat(s); ok {
		return number.FloatToInteger(f)
	}
	return 0, false
}
