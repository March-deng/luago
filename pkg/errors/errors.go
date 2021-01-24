package errors

import "errors"

var (
	ErrNotEnoughData = errors.New("not enough data left")

	ErrMalFormat = errors.New("chunk malformat")
)
