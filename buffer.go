package gobuf

import (
	"encoding/binary"
	"io"
)

type Buffer struct {
	*Peeker
	*Reader
	*Writer
	size    int
	mem     Memory
	options []OptionFunc
	order   binary.ByteOrder
}

func New(buf []byte, options ...OptionFunc) *Buffer {
	b := &Buffer{
		mem:     NewSliceMemory(buf),
		options: options,
		order:   binary.LittleEndian,
	}
	b.size = b.mem.Length()
	for _, option := range options {
		option(b, buf)
	}
	b.Peeker = NewPeeker(b)
	b.Reader = NewRead(b, b.Peeker)
	b.Writer = NewWriter(b)
	return b
}

func (buf *Buffer) PeekAt(at int, dst []byte) (n int, err error) {
	length := buf.mem.Length()
	if n > length {
		return 0, io.EOF
	}

	n = len(dst)
	if at+n > length {
		n = buf.mem.Length() - at
	}

	err = buf.mem.Read(at, dst[:n])
	return
}

func (buf *Buffer) WriteAt(at int, src []byte) (n int, err error) {
	size := at + len(src)

	err = buf.mem.Write(at, src)
	if err != nil {
		return
	}

	if size > buf.size {
		buf.size = size
	}

	return len(src), nil
}

func (buf *Buffer) Size() int {
	return buf.size
}

func (buf *Buffer) Bytes() []byte {
	return buf.mem.Bytes()
}

func (buf *Buffer) Order() binary.ByteOrder {
	return buf.order
}
