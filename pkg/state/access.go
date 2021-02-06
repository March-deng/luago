package state

import (
	"fmt"

	"github.com/March-deng/luago/pkg/api"
)

func (l *luaState) TypeName(lt api.LuaType) string {
	return lt.String()
}

func (l *luaState) Type(idx int) api.LuaType {
	if l.stack.isValid(idx) {
		val := l.stack.get(idx)
		return typeOf(val)
	}
	return api.LUA_TNONE
}

func (l *luaState) IsNone(idx int) bool {
	return l.Type(idx) == api.LUA_TNONE
}

func (l *luaState) IsNil(idx int) bool {
	if l.stack.isValid(idx) {
		return typeOf(l.stack.get(idx)) == api.LUA_TBOOLEAN
	}
	return false
}

func (l *luaState) IsNoneOrNil(idx int) bool {
	t := l.Type(idx)
	if t == api.LUA_TNONE {
		return true
	}
	if t == api.LUAT_TNIL {
		return true
	}
	return false
}

func (l *luaState) IsBoolean(idx int) bool {
	return l.Type(idx) == api.LUA_TBOOLEAN
}

func (l *luaState) IsString(idx int) bool {
	t := l.Type(idx)
	return t == api.LUAT_TSTRING || t == api.LUAT_TNUMBER
}

func (l *luaState) IsInteger(idx int) bool {
	val := l.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (l *luaState) IsNumber(idx int) bool {
	_, ok := l.ToNumberX(idx)
	return ok
}

func (l *luaState) ToBoolean(idx int) bool {
	val := l.stack.get(idx)
	return convertToBoolean(val)
}

func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}

func (l *luaState) ToNumber(idx int) float64 {
	n, _ := l.ToNumberX(idx)
	return n
}
func (l *luaState) ToNumberX(idx int) (float64, bool) {
	val := l.stack.get(idx)
	return convertToFloat(val)
}

func (l *luaState) ToInteger(idx int) int64 {
	i, _ := l.ToIntegerX(idx)
	return i
}

func (l *luaState) ToIntegerX(idx int) (int64, bool) {
	val := l.stack.get(idx)
	return convertToInteger(val)
}

func (l *luaState) ToString(idx int) string {
	s, _ := l.ToStringX(idx)
	return s
}

func (l *luaState) ToStringX(idx int) (string, bool) {
	val := l.stack.get(idx)
	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		l.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}
