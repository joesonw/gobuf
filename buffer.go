package gobuf

import "encoding/binary"

type Buffer struct {
	mem        Memory
	size       int
	writeIndex int
	readIndex  int
	options    []OptionFunc
	order      binary.ByteOrder
}

func New(buf []byte, options ...OptionFunc) *Buffer {
	buffer := &Buffer{
		mem:     newSliceMemory(buf),
		options: options,
		order:   binary.LittleEndian,
	}
	for _, option := range options {
		option(buffer, buf)
	}
	buffer.size = len(buf)
	return buffer
}

func (buf *Buffer) Clone() *Buffer {
	b := New(buf.Bytes(), buf.options...)
	b.writeIndex = buf.writeIndex
	b.readIndex = buf.readIndex
	b.size = buf.size
	b.order = buf.order
	return b
}

func (buf *Buffer) WriteIndex() int {
	return buf.writeIndex
}

func (buf *Buffer) ReadIndex() int {
	return buf.readIndex
}

func (buf *Buffer) Available() int {
	return buf.size - buf.readIndex
}

func (buf *Buffer) Size() int {
	return buf.size
}

func (buf *Buffer) Reset() {
	buf.writeIndex = 0
	buf.readIndex = 0
	buf.size = 0
}

func (buf *Buffer) Write(src []byte) (n int, err error) {
	size := buf.writeIndex + len(src)

	err = buf.mem.Write(buf.writeIndex, src)
	if err != nil {
		return
	}

	if size > buf.size {
		buf.size = size
	}
	buf.writeIndex = size
	return len(src), nil
}

func (buf *Buffer) Read(dst []byte) (n int, err error) {
	n, err = buf.Peek(0, dst)
	if err == nil {
		buf.readIndex += n
	}
	return
}

func (buf *Buffer) Peek(offset int, dst []byte) (n int, err error) {
	index := buf.readIndex + offset
	n = len(dst)
	if n+index > buf.mem.Length() {
		n = buf.mem.Length() - index
	}

	err = buf.mem.Read(index, dst[:n])
	return
}

func (buf *Buffer) SkipRead(n int) {
	index := buf.readIndex + n
	if index > buf.mem.Length() {
		index = buf.mem.Length()
	}
	buf.readIndex = index
}

func (buf *Buffer) Bytes() []byte {
	return buf.mem.Bytes()
}

func (buf *Buffer) Order() binary.ByteOrder {
	return buf.order
}
