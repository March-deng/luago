package binchunk

type binaryChunk struct {
	header
	sizeUpvalues int8
	mainFunc     *Prototype
}

func Undump(data []byte) (*Prototype, error) {
	return nil, nil
}
