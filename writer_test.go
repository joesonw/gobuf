package gobuf

import (
	"bytes"
	"encoding/binary"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Writer", func() {
	It("should should writer", func() {

		writer := bytes.NewBuffer(nil)
		w := Write(writer, binary.LittleEndian)

		ExpectSizeError(5)(w.Write([]byte("hello")))
		Expect(writer.String()).To(Equal("hello"))

		ExpectSizeError(6)(w.Write([]byte(" world")))
		Expect(writer.String()).To(Equal("hello world"))
	})
})
