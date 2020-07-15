package gobuf

import (
	"encoding/binary"
	"io"
)

type IOWriter struct {
	*Writer
	writer io.Writer
	order  binary.ByteOrder
}

func Write(w io.Writer, order binary.ByteOrder) *IOWriter {
	writer := &IOWriter{
		writer: w,
		order:  order,
	}
	writer.Writer = NewWriter(writer)
	return writer
}

func (w *IOWriter) Order() binary.ByteOrder {
	return w.order
}

func (w *IOWriter) WriteSome(src []byte) (n int, err error) {
	return w.writer.Write(src)
}
