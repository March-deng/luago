package binchunk

type Prototype struct {
	//source file
	Source string
	//first line number
	FirstLine uint32
	//last line number
	LastLine uint32
	//fixed params num
	FixedParamsNum uint8
	IsVararg       uint8
	MaxStackSize   uint8
	Code           []uint32
	Constants      []interface{}
	Upvalues       []Upvalue
	Protos         []*Prototype
	//line for every code in codes
	LineInfo     []uint32
	LocalVars    []LocalVar
	UpvalueNames []string
}

type Upvalue struct {
	InStack uint8
	Idx     uint8
}

type LocalVar struct {
	Name     string
	StartPos uint32
	EndPos   uint32
}
