package gobuf

import "errors"

var (
	ErrOutOfSpace    = errors.New("not enough space to write")
	ErrNotEnoughData = errors.New("not enough data to read")
)
