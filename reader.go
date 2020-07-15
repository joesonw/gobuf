package gobuf

import (
	"encoding/binary"
	"io"
)

type IOReader struct {
	*Peeker
	*Reader
	reader io.Reader
	read   int
	mem    Memory
	order  binary.ByteOrder
}

func Read(r io.Reader, order binary.ByteOrder, memory Memory) *IOReader {
	rd := &IOReader{
		reader: r,
		mem:    memory,
		read:   memory.Length(),
		order:  order,
	}
	rd.Peeker = NewPeeker(rd)
	rd.Reader = NewRead(rd, rd.Peeker)
	return rd
}

func (r *IOReader) PeekAt(at int, dst []byte) (n int, err error) {
	n = len(dst)

	if at > r.read || (at+n) > r.read {
		amount := at + n - r.read
		b := make([]byte, amount)
		_, err = r.reader.Read(b)
		if err != nil {
			return
		}
		r.read += amount

		err = r.mem.Write(r.read-amount, b)
		if err != nil {
			return
		}
	}

	length := r.mem.Length()
	n = len(dst)
	if at+n > length {
		n = r.mem.Length() - at
	}

	err = r.mem.Read(at, dst[:n])
	return
}

func (r *IOReader) Order() binary.ByteOrder {
	return r.order
}
