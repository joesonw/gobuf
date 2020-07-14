package gobuf

import "math"

// PeekBool peek a bool
func (buf *Buffer) PeekBool(offset ...int) (bool, error) {
	b, err := buf.PeekByte(offset...)
	return b == 1, err
}

// PeekByte peek a byte
func (buf *Buffer) PeekByte(offset ...int) (byte, error) {
	b, err := buf.PeekBytes(1, offset...)
	if err != nil {
		return 0, err
	}

	return b[0], nil
}

// PeekBytes peek given length of bytes
func (buf *Buffer) PeekBytes(n int, offset ...int) ([]byte, error) {
	o := 0
	if len(offset) > 0 {
		o = offset[0]
	}
	b := make([]byte, n)
	_, err := buf.Peek(o, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// PeekString peek given length of string
func (buf *Buffer) PeekString(n int, offset ...int) (string, error) {
	b, err := buf.PeekBytes(n, offset...)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PeekUint8 peek uint8
func (buf *Buffer) PeekUint8(offset ...int) (uint8, error) {
	b, err := buf.PeekByte(offset...)
	if err != nil {
		return 0, err
	}

	return b, nil
}

// PeekUint16 peek uint16
func (buf *Buffer) PeekUint16(offset ...int) (uint16, error) {
	b, err := buf.PeekBytes(2, offset...)
	if err != nil {
		return 0, err
	}

	return buf.order.Uint16(b), nil
}

// PeekUint32 peek uint32
func (buf *Buffer) PeekUint32(offset ...int) (uint32, error) {
	b, err := buf.PeekBytes(4, offset...)
	if err != nil {
		return 0, err
	}

	return buf.order.Uint32(b), nil
}

// PeekUint64 peek uint64
func (buf *Buffer) PeekUint64(offset ...int) (uint64, error) {
	b, err := buf.PeekBytes(8, offset...)
	if err != nil {
		return 0, err
	}

	return buf.order.Uint64(b), nil
}

// PeekInt8 peek int8
func (buf *Buffer) PeekInt8(offset ...int) (int8, error) {
	b, err := buf.PeekByte(offset...)
	if err != nil {
		return 0, err
	}

	return int8(b), nil
}

// PeekInt16 peek int16
func (buf *Buffer) PeekInt16(offset ...int) (int16, error) {
	b, err := buf.PeekBytes(2, offset...)
	if err != nil {
		return 0, err
	}

	return int16(buf.order.Uint16(b)), nil
}

// PeekInt32 peek int32
func (buf *Buffer) PeekInt32(offset ...int) (int32, error) {
	b, err := buf.PeekBytes(4, offset...)
	if err != nil {
		return 0, err
	}

	return int32(buf.order.Uint32(b)), nil
}

// PeekInt64 peek int64
func (buf *Buffer) PeekInt64(offset ...int) (int64, error) {
	b, err := buf.PeekBytes(8, offset...)
	if err != nil {
		return 0, err
	}

	return int64(buf.order.Uint64(b)), nil
}

// PeekFloat32 peek float32
func (buf *Buffer) PeekFloat32(offset ...int) (float32, error) {
	u, err := buf.PeekUint32(offset...)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(u), nil
}

// PeekFloat64 peek float64
func (buf *Buffer) PeekFloat64(offset ...int) (float64, error) {
	u, err := buf.PeekUint64(offset...)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(u), nil
}
