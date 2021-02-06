package state

import "github.com/March-deng/luago/pkg/api"

func (l *luaState) Compare(idx1, idx2 int, op api.CompareOp) bool {
	a, b := l.stack.get(idx1), l.stack.get(idx2)
	switch op {
	case api.LUA_OPEQ:
		return _eq(a, b)
	case api.LUA_OPLT:
		return _lt(a, b)
	case api.LUA_OPLE:
		return _le(a, b)
	default:
		panic("invalid compare op!")
	}
	return false
}

func _eq(a, b luaValue) bool {
	switch x := a.(type) {
	case nil:
		return b == nil
	case bool:
		y, ok := b.(bool)
		return ok && x == y
	case string:
		y, ok := b.(string)
		return ok && x == y
	case int64:
		y, ok := b.(int64)
		return ok && x == y
	case float64:
		y, ok := b.(float64)
		return ok && x == y
	default:
		return a == b
	}
}

func _lt(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x < y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x < y
		case float64:
			return float64(x) < y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x < float64(y)
		case float64:
			return x < y
		}
	}
	panic("comparison error!")
}

func _le(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x <= y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x <= y
		case float64:
			return float64(x) <= y
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x <= float64(y)
		case float64:
			return x <= y
		}
	}
	panic("comparison error!")
}
