package gobuf

import "encoding/binary"

type OptionFunc func(b *Buffer, buf []byte)

func WithMemory(m Memory) OptionFunc {
	return func(b *Buffer, buf []byte) {
		b.mem = m
	}
}

func WithAutoGrowMemory(grow Grow) OptionFunc {
	return func(b *Buffer, buf []byte) {
		m := newSliceMemory(buf)
		m.grow = grow
		b.mem = m
	}
}

func WithLinkedListMemory(grow Grow) OptionFunc {
	return func(b *Buffer, buf []byte) {
		b.mem = newLinkedListMemory(buf, grow)
	}
}

func WithLittleEndian() OptionFunc {
	return func(b *Buffer, buf []byte) {
		b.order = binary.LittleEndian
	}
}

func WithBigEndian() OptionFunc {
	return func(b *Buffer, buf []byte) {
		b.order = binary.BigEndian
	}
}
