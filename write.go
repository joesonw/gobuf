package gobuf

import (
	"encoding/binary"
	"math"
)

type Writable interface {
	WriteAt(at int, src []byte) (n int, err error)
	Order() binary.ByteOrder
}

type Writer struct {
	Writable
	index int
}

func NewWriter(w Writable) *Writer {
	return &Writer{
		Writable: w,
	}
}

func (w *Writer) WriteIndex() int {
	return w.index
}

func (w *Writer) Write(src []byte) (n int, err error) {
	size := w.WriteIndex() + len(src)
	n, err = w.WriteAt(w.WriteIndex(), src)
	if err != nil {
		return
	}
	w.index = size
	return
}

// WriteBool write a bool value as byte
func (w *Writer) WriteBool(val bool) error {
	if val {
		return w.WriteByte(byte(1))
	}
	return w.WriteByte(byte(0))
}

// WriteByte write a byte into buffer
func (w *Writer) WriteByte(b byte) error {
	_, err := w.Write([]byte{b})
	return err
}

// WriteBytes write bytes into buffer
func (w *Writer) WriteBytes(b []byte) error {
	_, err := w.Write(b)
	return err
}

// WriteString write a string into buffer
func (w *Writer) WriteString(s string) error {
	return w.WriteBytes([]byte(s))
}

// WriteUint8 write a uint8 into buffer
func (w *Writer) WriteUint8(val uint8) error {
	return w.WriteByte(val)
}

// WriteUint16 write a uint16 into buffer
func (w *Writer) WriteUint16(val uint16) error {
	b := make([]byte, 2)
	w.Order().PutUint16(b, val)
	return w.WriteBytes(b)
}

// WriteUint32 write a uint32 into buffer
func (w *Writer) WriteUint32(val uint32) error {
	b := make([]byte, 4)
	w.Order().PutUint32(b, val)
	return w.WriteBytes(b)
}

// WriteUint64 write a uint64 into buffer
func (w *Writer) WriteUint64(val uint64) error {
	b := make([]byte, 8)
	w.Order().PutUint64(b, val)
	return w.WriteBytes(b)
}

// WriteUint8 write a int8 into buffer
func (w *Writer) WriteInt8(val int8) error {
	return w.WriteByte(byte(val))
}

// WriteInt16 write a int16 into buffer
func (w *Writer) WriteInt16(val int16) error {
	b := make([]byte, 2)
	w.Order().PutUint16(b, uint16(val))
	return w.WriteBytes(b)
}

// WriteInt32 write a int32 into buffer
func (w *Writer) WriteInt32(val int32) error {
	b := make([]byte, 4)
	w.Order().PutUint32(b, uint32(val))
	return w.WriteBytes(b)
}

// WriteInt64 write a int64 into buffer
func (w *Writer) WriteInt64(val int64) error {
	b := make([]byte, 8)
	w.Order().PutUint64(b, uint64(val))
	return w.WriteBytes(b)
}

// WriteFloat32 write a float32 into buffer
func (w *Writer) WriteFloat32(val float32) error {
	return w.WriteUint32(math.Float32bits(val))
}

// WriteFloat64 write a float64 into buffer
func (w *Writer) WriteFloat64(val float64) error {
	return w.WriteUint64(math.Float64bits(val))
}
