package gobuf

import (
	"io"
)

// Memory used to store raw []byte
type Memory interface {
	// Write write at given location, with b
	Write(at int, src []byte) error

	// Read read from given location, to b
	Read(at int, dst []byte) error

	// Compact into one single byte array
	Bytes() []byte

	// Length total length
	Length() int

	// Reset reset memory
	Reset()
}

type SliceMemory struct {
	buf  []byte
	grow Grow
}

func NewSliceMemory(buf []byte, grow Grow) *SliceMemory {
	return &SliceMemory{
		buf:  buf,
		grow: grow,
	}
}

func (m *SliceMemory) Write(at int, src []byte) error {
	end := at + len(src)
	if end > cap(m.buf) {
		if m.grow == nil {
			return ErrOutOfSpace
		}
		finalCap := m.grow(cap(m.buf), end)
		newBuf := make([]byte, finalCap)
		copy(newBuf, m.buf)
		m.buf = newBuf
	}

	copy(m.buf[at:end], src)
	return nil
}

func (m *SliceMemory) Read(at int, dst []byte) error {
	end := at + len(dst)
	if end > cap(m.buf) {
		return io.EOF
	}

	copy(dst, m.buf[at:end])
	return nil
}

func (m *SliceMemory) Bytes() []byte {
	out := make([]byte, len(m.buf))
	copy(out, m.buf)
	return out
}

func (m *SliceMemory) Length() int {
	return len(m.buf)
}

func (m *SliceMemory) Reset() {
	m.buf = make([]byte, m.grow(0, 1))
}

type listMemoryNode struct {
	buf  []byte
	next *listMemoryNode
}

type ListMemory struct {
	start *listMemoryNode
	grow  Grow
}

func NewListMemory(buf []byte, grow Grow) *ListMemory {
	return &ListMemory{
		start: &listMemoryNode{
			buf: buf,
		},
		grow: grow,
	}
}

func (m *ListMemory) Write(at int, src []byte) error {
	total := len(src)
	end := total + at
	capacity := 0
	node := m.start
	prev := node
	wrote := 0
	index := 0

	for wrote < total {
		// reached end of list, allocate new node
		if node == nil {
			newCapacity := m.grow(capacity, end)
			node = &listMemoryNode{
				buf: make([]byte, newCapacity-capacity),
			}
			prev.next = node
			capacity = newCapacity
		}

		nodeCapacity := cap(node.buf)
		capacity += nodeCapacity

		// skip until cursor is at the correct node
		if (index + nodeCapacity) < at {
			index += nodeCapacity
			prev = node
			node = node.next
			continue
		}

		offset := at - index
		if offset < 0 {
			offset = 0
		}

		available := nodeCapacity - offset
		remain := total - wrote
		// current node is enough to write rest data
		if remain <= available {
			copy(node.buf[offset:], src[wrote:])
			break
		}

		copy(node.buf[offset:], src[wrote:wrote+available])
		wrote += available
		index += nodeCapacity
		prev = node
		node = node.next
	}

	return nil
}

func (m *ListMemory) Read(at int, dst []byte) error {
	total := len(dst)
	node := m.start
	read := 0
	index := 0
	length := 0

	for read < total {
		if node == nil {
			return io.EOF
		}

		nodeCapacity := cap(node.buf)
		length += nodeCapacity

		// skip until cursor is at the correct node
		if (index + nodeCapacity) < at {
			index += nodeCapacity
			node = node.next
			continue
		}

		offset := at - index
		if offset < 0 {
			offset = 0
		}

		available := nodeCapacity - offset
		remain := total - read
		// current node is enough to read rest data
		if remain <= available {
			copy(dst[read:], node.buf[offset:offset+remain])
			break
		}

		copy(dst[read:], node.buf[offset:])
		read += available
		index += nodeCapacity
		node = node.next
	}

	return nil
}

func (m *ListMemory) Bytes() []byte {
	length := m.Length()

	out := make([]byte, length)
	index := 0
	node := m.start
	for node != nil {
		copy(out[index:], node.buf)
		index += cap(node.buf)
		node = node.next
	}

	return out
}

func (m *ListMemory) Length() int {
	length := 0
	node := m.start
	for node != nil {
		length += cap(node.buf)
		node = node.next
	}

	return length
}

func (m *ListMemory) Reset() {
	m.start = &listMemoryNode{
		buf: make([]byte, m.grow(0, 1)),
	}
}
