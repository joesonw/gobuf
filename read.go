package gobuf

import (
	"bytes"
	"math"
)

type Readable interface {
	Size() int
}

type Reader struct {
	Readable
	*Peeker
}

func NewRead(r Readable, p *Peeker) *Reader {
	return &Reader{
		Readable: r,
		Peeker:   p,
	}
}

// SkipRead advance read index
func (r *Reader) SkipRead(n int) {
	r.Peeker.index += n
}

// Available available bytes to read
func (r *Reader) Available() int {
	return r.Size() - r.ReaderIndex()
}

// Read io.Reader
func (r *Reader) Read(dst []byte) (n int, err error) {
	n, err = r.Peek(0, dst)
	if err == nil {
		r.SkipRead(n)
	}
	return
}

// ReadBool read a bool
func (r *Reader) ReadBool() (bool, error) {
	b, err := r.ReadByte()
	return b == 1, err
}

// PeekByte read a byte
func (r *Reader) ReadByte() (byte, error) {
	b, err := r.ReadBytes(1)
	if err != nil {
		return 0, err
	}

	return b[0], nil
}

// ReadBytes read given length of bytes
func (r *Reader) ReadBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// ReadString read given length of string
func (r *Reader) ReadString(n int) (string, error) {
	b, err := r.ReadBytes(n)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ReadUint8 read uint8
func (r *Reader) ReadUint8() (uint8, error) {
	b, err := r.ReadByte()
	return b, err
}

// ReadUint16 read uint16
func (r *Reader) ReadUint16() (uint16, error) {
	b, err := r.ReadBytes(2)
	if err != nil {
		return 0, err
	}

	return r.Order().Uint16(b), nil
}

// ReadUint32 read uint32
func (r *Reader) ReadUint32() (uint32, error) {
	b, err := r.ReadBytes(4)
	if err != nil {
		return 0, err
	}

	return r.Order().Uint32(b), nil
}

// ReadUint64 read uint64
func (r *Reader) ReadUint64() (uint64, error) {
	b, err := r.ReadBytes(8)
	if err != nil {
		return 0, err
	}

	return r.Order().Uint64(b), nil
}

// ReadInt8 read int8
func (r *Reader) ReadInt8() (int8, error) {
	b, err := r.ReadByte()
	return int8(b), err
}

// ReadInt16 read int16
func (r *Reader) ReadInt16() (int16, error) {
	b, err := r.ReadBytes(2)
	if err != nil {
		return 0, err
	}

	return int16(r.Order().Uint16(b)), nil
}

// ReadInt32 read int32
func (r *Reader) ReadInt32() (int32, error) {
	b, err := r.ReadBytes(4)
	if err != nil {
		return 0, err
	}

	return int32(r.Order().Uint32(b)), nil
}

// ReadInt64 read int64
func (r *Reader) ReadInt64() (int64, error) {
	b, err := r.ReadBytes(8)
	if err != nil {
		return 0, err
	}

	return int64(r.Order().Uint64(b)), nil
}

// ReadFloat32 read float32
func (r *Reader) ReadFloat32() (float32, error) {
	u, err := r.ReadUint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(u), nil
}

// ReadFloat64 read float64
func (r *Reader) ReadFloat64() (float64, error) {
	u, err := r.ReadUint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(u), nil
}

//nolint:gocritic
// ReadUntil read until any delimiter matches, than skip delimiter and return. otherwise, return false
func (r *Reader) ReadUntil(delims ...[]byte) ([]byte, bool, error) {
	// short circuit
	short := true
	for _, delim := range delims {
		// if any delim is shorter than available data, means we can short, no need to short
		if len(delim) <= r.Available() {
			short = false
		}
	}

	if short {
		return nil, false, nil
	}

	for _, delim := range delims {
		out, err := r.PeekBytes(len(delim))
		if err != nil {
			return nil, false, err
		}
		if bytes.Equal(out, delim) {
			r.SkipRead(len(delim))
			return []byte{}, true, nil
		}
	}

	var out []byte
	index := r.ReaderIndex()
	offset := 0
	size := r.Size()

	for index < size {
		b, err := r.PeekByte(offset)
		if err != nil {
			return nil, false, err
		}
		out = append(out, b)

		for _, delim := range delims {
			length := len(delim)
			start := len(out) - length
			if bytes.Equal(out[start:], delim) {
				r.SkipRead(len(out))
				return out[:start], true, nil
			}
		}
		offset++
		index++
	}
	return nil, false, nil
}
