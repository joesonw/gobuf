package gobuf

import "errors"

var (
	ErrOutOfSpace = errors.New("not enough space to write")
)
