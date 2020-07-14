package gobuf

import "math"

// WriteBool write a bool value as byte
func (buf *Buffer) WriteBool(val bool) error {
	if val {
		return buf.WriteByte(byte(1))
	}
	return buf.WriteByte(byte(0))
}

// WriteByte write a byte into buffer
func (buf *Buffer) WriteByte(b byte) error {
	_, err := buf.Write([]byte{b})
	return err
}

// WriteBytes write bytes into buffer
func (buf *Buffer) WriteBytes(b []byte) error {
	_, err := buf.Write(b)
	return err
}

// WriteString write a string into buffer
func (buf *Buffer) WriteString(s string) error {
	return buf.WriteBytes([]byte(s))
}

// WriteUint8 write a uint8 into buffer
func (buf *Buffer) WriteUint8(val uint8) error {
	return buf.WriteByte(val)
}

// WriteUint16 write a uint16 into buffer
func (buf *Buffer) WriteUint16(val uint16) error {
	b := make([]byte, 2)
	buf.order.PutUint16(b, val)
	return buf.WriteBytes(b)
}

// WriteUint32 write a uint32 into buffer
func (buf *Buffer) WriteUint32(val uint32) error {
	b := make([]byte, 4)
	buf.order.PutUint32(b, val)
	return buf.WriteBytes(b)
}

// WriteUint64 write a uint64 into buffer
func (buf *Buffer) WriteUint64(val uint64) error {
	b := make([]byte, 8)
	buf.order.PutUint64(b, val)
	return buf.WriteBytes(b)
}

// WriteUint8 write a int8 into buffer
func (buf *Buffer) WriteInt8(val int8) error {
	return buf.WriteByte(byte(val))
}

// WriteInt16 write a int16 into buffer
func (buf *Buffer) WriteInt16(val int16) error {
	b := make([]byte, 2)
	buf.order.PutUint16(b, uint16(val))
	return buf.WriteBytes(b)
}

// WriteInt32 write a int32 into buffer
func (buf *Buffer) WriteInt32(val int32) error {
	b := make([]byte, 4)
	buf.order.PutUint32(b, uint32(val))
	return buf.WriteBytes(b)
}

// WriteInt64 write a int64 into buffer
func (buf *Buffer) WriteInt64(val int64) error {
	b := make([]byte, 8)
	buf.order.PutUint64(b, uint64(val))
	return buf.WriteBytes(b)
}

// WriteFloat32 write a float32 into buffer
func (buf *Buffer) WriteFloat32(val float32) error {
	return buf.WriteUint32(math.Float32bits(val))
}

// WriteFloat64 write a float64 into buffer
func (buf *Buffer) WriteFloat64(val float64) error {
	return buf.WriteUint64(math.Float64bits(val))
}
