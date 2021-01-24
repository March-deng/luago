package binchunk

type header struct {
	//magic number
	signature [4]byte
	//lua version
	version int8
	//0
	format int8
	//another magic number
	luacdata        [6]byte
	cintSize        uint8
	sizetSize       uint8
	instructionSize uint8
	luaIntegerSize  uint8
	luaNumberSize   uint8
	luacInt         int64
	luacNum         float64
}

func buildHeader() header {
	return header{}
}
