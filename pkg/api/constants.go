package api

const (
	LUA_TNONE = iota - 1
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

const (
	LUA_OPADD  = iota // +
	LUA_OPSUB         // -
	LUA_OPMUL         // *
	LUA_OPMOD         // %
	LUA_OPPOW         // ^
	LUA_OPDIV         // /
	LUA_OPIDIV        // //
	LUA_OPBAND        // &
	LUA_OPBOR         // |
	LUA_OPBXOR        // ~
	LUA_OPSHL         // <<
	LUA_OPSHR         // >>
	LUA_OPUNM         // - (unary minus)
	LUA_OPBNOT        // ~

)

const (
	LUA_OPEQ = iota // ==
	LUA_OPLT        // <
	LUA_OPLE        // >
)
