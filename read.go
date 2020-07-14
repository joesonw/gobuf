package gobuf

import (
	"bytes"
	"math"
)

// ReadBool read a bool
func (buf *Buffer) ReadBool() (bool, error) {
	b, err := buf.ReadByte()
	return b == 1, err
}

// PeekByte read a byte
func (buf *Buffer) ReadByte() (byte, error) {
	b, err := buf.ReadBytes(1)
	if err != nil {
		return 0, err
	}

	return b[0], nil
}

// ReadBytes read given length of bytes
func (buf *Buffer) ReadBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := buf.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// ReadString read given length of string
func (buf *Buffer) ReadString(n int) (string, error) {
	b, err := buf.ReadBytes(n)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ReadUint8 read uint8
func (buf *Buffer) ReadUint8() (uint8, error) {
	b, err := buf.ReadByte()
	return b, err
}

// ReadUint16 read uint16
func (buf *Buffer) ReadUint16() (uint16, error) {
	b, err := buf.ReadBytes(2)
	if err != nil {
		return 0, err
	}

	return buf.order.Uint16(b), nil
}

// ReadUint32 read uint32
func (buf *Buffer) ReadUint32() (uint32, error) {
	b, err := buf.ReadBytes(4)
	if err != nil {
		return 0, err
	}

	return buf.order.Uint32(b), nil
}

// ReadUint64 read uint64
func (buf *Buffer) ReadUint64() (uint64, error) {
	b, err := buf.ReadBytes(8)
	if err != nil {
		return 0, err
	}

	return buf.order.Uint64(b), nil
}

// ReadInt8 read int8
func (buf *Buffer) ReadInt8() (int8, error) {
	b, err := buf.ReadByte()
	return int8(b), err
}

// ReadInt16 read int16
func (buf *Buffer) ReadInt16() (int16, error) {
	b, err := buf.ReadBytes(2)
	if err != nil {
		return 0, err
	}

	return int16(buf.order.Uint16(b)), nil
}

// ReadInt32 read int32
func (buf *Buffer) ReadInt32() (int32, error) {
	b, err := buf.ReadBytes(4)
	if err != nil {
		return 0, err
	}

	return int32(buf.order.Uint32(b)), nil
}

// ReadInt64 read int64
func (buf *Buffer) ReadInt64() (int64, error) {
	b, err := buf.ReadBytes(8)
	if err != nil {
		return 0, err
	}

	return int64(buf.order.Uint64(b)), nil
}

// ReadFloat32 read float32
func (buf *Buffer) ReadFloat32() (float32, error) {
	u, err := buf.ReadUint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(u), nil
}

// ReadFloat64 read float64
func (buf *Buffer) ReadFloat64() (float64, error) {
	u, err := buf.ReadUint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(u), nil
}

//nolint:gocritic
// ReadUntil read until delimiter, than skip delimiter and return. otherwise, return false
func (buf *Buffer) ReadUntil(delim []byte) ([]byte, bool, error) {
	length := len(delim)
	// short circuit
	if buf.Available() < length {
		return nil, false, nil
	}

	out, err := buf.PeekBytes(length)
	if err != nil {
		return nil, false, err
	}
	if bytes.Equal(out, delim) {
		buf.readIndex += length
		return []byte{}, true, nil
	}

	index := buf.readIndex + length
	offset := length
	size := buf.size

	for index < size {
		b, err := buf.PeekByte(offset)
		if err != nil {
			return nil, false, err
		}
		out = append(out, b)

		if bytes.Equal(out[offset:], delim) {
			buf.readIndex += offset + 1
			return out[:offset], true, nil
		}
		offset++
		index++
	}
	return nil, false, nil
}
