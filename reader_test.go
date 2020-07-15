package gobuf

import (
	"encoding/binary"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reader", func() {
	It("should should read", func() {
		mem := NewSliceMemory([]byte("hello"), FixedGrow(5))
		reader := strings.NewReader(" world")
		r := Read(reader, binary.LittleEndian, mem)

		b := make([]byte, 11)
		ExpectSizeError(4)(r.Read(b[:4]))
		Expect(string(b[:4])).To(Equal("hell"))

		ExpectSizeError(5)(r.Read(b[4:9]))
		Expect(string(b[0:9])).To(Equal("hello wor"))

		ExpectSizeError(2)(r.Read(b[9:]))
		Expect(string(b)).To(Equal("hello world"))
	})
})
