package binchunk


type header struct {
	//magic number
	signature [4]byte
	//lua version
	version byte
	//0
	format byte
	//another magic number
	luacdata [6]byte
	cintSize int8
	sizetSize int8
	instructionSize int8
	luaIntegerSize int8
	luaNumberSize int8
	luacInt int64
	luacNum float64
}