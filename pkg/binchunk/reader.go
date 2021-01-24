package binchunk

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"

	"github.com/March-deng/luago/pkg/errors"
)

type Reader struct {
	reader *bytes.Reader
}

func NewReader(data []byte) *Reader {
	return &Reader{
		reader: bytes.NewReader(data),
	}
}

func (r *Reader) checkHeader() error {
	bs, err := r.readBytes(4)
	if err != nil {
		return err
	}
	if string(bs) != LUA_SIGNATURE {
		return errors.ErrMalFormat
	}
	if r.readByte() != LUAC_VERSION {
		return errors.ErrMalFormat
	}
	if r.readByte() != LUAC_FORMAT {
		return errors.ErrMalFormat
	}
	bs, err = r.readBytes(6)
	if err != nil {
		return err
	}
	if string(bs) != LUAC_DATA {
		return errors.ErrMalFormat
	}

	if r.readByte() != CINT_SIZE {
		return errors.ErrMalFormat
	}

	if r.readByte() != CSIZET_SIZE {
		return errors.ErrMalFormat
	}

	if r.readByte() != INSTRUCTION_SIZE {
		return errors.ErrMalFormat
	}

	if r.readByte() != LUA_INTEGER_SIZE {
		return errors.ErrMalFormat
	}

	if r.readByte() != LUA_NUMBER_SIZE {
		return errors.ErrMalFormat
	}

	if r.readLuaInteger() != LUAC_INT {
		return errors.ErrMalFormat
	}

	if r.readLuaNumber() != LUAC_NUM {
		return errors.ErrMalFormat
	}
	return nil
}

func (r *Reader) readByte() byte {
	b, err := r.reader.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

func (r *Reader) readBytes(n uint) ([]byte, error) {
	bs := make([]byte, n)
	nn, err := r.reader.Read(bs)
	if err != nil {
		if nn == int(n) && err == io.EOF {
			return bs, nil
		}
		return nil, err
	}
	if nn != int(n) {
		return nil, errors.ErrNotEnoughData
	}
	return bs, nil
}
func (r *Reader) readUint32() uint32 {
	bs, err := r.readBytes(4)
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint32(bs)
}

func (r *Reader) readUint64() uint64 {
	bs, err := r.readBytes(8)
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(bs)
}

func (r *Reader) readLuaInteger() int64 {
	return int64(r.readUint64())
}
func (r *Reader) readLuaNumber() float64 {
	return math.Float64frombits(r.readUint64())
}

func (r *Reader) readString() string {
	size := uint(r.readByte())
	if size == 0 {
		return ""
	}

	if size == 0xff {
		size = uint(r.readUint64())
	}
	data, err := r.readBytes(size - 1)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (r *Reader) readProto(parent string) *Prototype {
	source := r.readString()
	if source == "" {
		source = parent
	}
	p := &Prototype{
		Source:         source,
		FirstLine:      r.readUint32(),
		LastLine:       r.readUint32(),
		FixedParamsNum: r.readByte(),
		IsVararg:       r.readByte(),
		MaxStackSize:   r.readByte(),
	}
	return p
}

func (r *Reader) readCode() []uint32 {
	code := make([]uint32, r.readUint32())
	for i := range code {
		code[i] = r.readUint32()
	}
	return code
}

func (r *Reader) readConstants() []interface{} {
	constants := make([]interface{}, r.readUint32())
	for i := range constants {
		constants[i] = r.readConstant()
	}
	return constants
}

func (r *Reader) readConstant() interface{} {
	switch r.readByte() {
	case TAG_BOOLEAN:
		return r.readByte() != 0
	case TAG_INTEGER:
		return r.readLuaInteger()
	case TAG_LONG_STR:
		return r.readString()
	case TAG_NUMBER:
		return r.readLuaNumber()
	case TAG_SHORT_STR:
		return r.readString()
	case TAG_NIL:
		return nil
	default:
		panic("corrupted")
	}
}

func (r *Reader) readUpvalues() []Upvalue {
	upvalues := make([]Upvalue, r.readUint32())
	for i := range upvalues {
		upvalues[i] = Upvalue{
			InStack: r.readByte(),
			Idx:     r.readByte(),
		}
	}
	return upvalues

}

func (r *Reader) readProtos(source string) []*Prototype {
	protos := make([]*Prototype, r.readUint32())

	for i := range protos {
		protos[i] = r.readProto(source)
	}
	return protos
}

func (r *Reader) readLineInfo() []uint32 {
	lineInfo := make([]uint32, r.readUint32())
	for i := range lineInfo {
		lineInfo[i] = r.readUint32()
	}
	return lineInfo
}

func (r *Reader) readLocalVars() []LocalVar {
	localVars := make([]LocalVar, r.readUint32())
	for i := range localVars {
		localVars[i] = LocalVar{
			Name:     r.readString(),
			StartPos: r.readUint32(),
			EndPos:   r.readUint32(),
		}
	}
	return localVars
}

func (r *Reader) readUpvalueNames() []string {
	names := make([]string, r.readUint32())
	for i := range names {
		names[i] = r.readString()
	}
	return names
}

func (r *Reader) readUpvaluesNames() []string {
	return nil
}
