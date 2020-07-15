package gobuf

import (
	"encoding/binary"
	"math"
)

type Peekable interface {
	PeekAt(at int, dst []byte) (int, error)
	Order() binary.ByteOrder
}

type Peeker struct {
	Peekable
	index int
}

func NewPeeker(p Peekable) *Peeker {
	return &Peeker{
		Peekable: p,
	}
}

func (p *Peeker) ReadIndex() int {
	return p.index
}

func (p *Peeker) Peek(offset int, dst []byte) (n int, err error) {
	index := p.ReadIndex() + offset
	return p.PeekAt(index, dst)
}

// PeekBool peek a bool
func (p *Peeker) PeekBool(offset ...int) (bool, error) {
	b, err := p.PeekByte(offset...)
	return b == 1, err
}

// PeekByte peek a byte
func (p *Peeker) PeekByte(offset ...int) (byte, error) {
	b, err := p.PeekBytes(1, offset...)
	if err != nil {
		return 0, err
	}

	return b[0], nil
}

// PeekBytes peek given length of bytes
func (p *Peeker) PeekBytes(n int, offset ...int) ([]byte, error) {
	o := 0
	if len(offset) > 0 {
		o = offset[0]
	}
	b := make([]byte, n)
	_, err := p.Peek(o, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// PeekString peek given length of string
func (p *Peeker) PeekString(n int, offset ...int) (string, error) {
	b, err := p.PeekBytes(n, offset...)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PeekUint8 peek uint8
func (p *Peeker) PeekUint8(offset ...int) (uint8, error) {
	b, err := p.PeekByte(offset...)
	if err != nil {
		return 0, err
	}

	return b, nil
}

// PeekUint16 peek uint16
func (p *Peeker) PeekUint16(offset ...int) (uint16, error) {
	b, err := p.PeekBytes(2, offset...)
	if err != nil {
		return 0, err
	}

	return p.Order().Uint16(b), nil
}

// PeekUint32 peek uint32
func (p *Peeker) PeekUint32(offset ...int) (uint32, error) {
	b, err := p.PeekBytes(4, offset...)
	if err != nil {
		return 0, err
	}

	return p.Order().Uint32(b), nil
}

// PeekUint64 peek uint64
func (p *Peeker) PeekUint64(offset ...int) (uint64, error) {
	b, err := p.PeekBytes(8, offset...)
	if err != nil {
		return 0, err
	}

	return p.Order().Uint64(b), nil
}

// PeekInt8 peek int8
func (p *Peeker) PeekInt8(offset ...int) (int8, error) {
	b, err := p.PeekByte(offset...)
	if err != nil {
		return 0, err
	}

	return int8(b), nil
}

// PeekInt16 peek int16
func (p *Peeker) PeekInt16(offset ...int) (int16, error) {
	b, err := p.PeekBytes(2, offset...)
	if err != nil {
		return 0, err
	}

	return int16(p.Order().Uint16(b)), nil
}

// PeekInt32 peek int32
func (p *Peeker) PeekInt32(offset ...int) (int32, error) {
	b, err := p.PeekBytes(4, offset...)
	if err != nil {
		return 0, err
	}

	return int32(p.Order().Uint32(b)), nil
}

// PeekInt64 peek int64
func (p *Peeker) PeekInt64(offset ...int) (int64, error) {
	b, err := p.PeekBytes(8, offset...)
	if err != nil {
		return 0, err
	}

	return int64(p.Order().Uint64(b)), nil
}

// PeekFloat32 peek float32
func (p *Peeker) PeekFloat32(offset ...int) (float32, error) {
	u, err := p.PeekUint32(offset...)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(u), nil
}

// PeekFloat64 peek float64
func (p *Peeker) PeekFloat64(offset ...int) (float64, error) {
	u, err := p.PeekUint64(offset...)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(u), nil
}
