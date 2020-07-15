package gobuf

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Buffer", func() {
	It("should write", func() {
		buf := New(nil, WithAutoGrowMemory(FixedGrow(5)))
		Expect(buf.Size()).To(Equal(0))
		Expect(buf.WriterIndex()).To(Equal(0))
		Expect(buf.ReaderIndex()).To(Equal(0))
		Expect(buf.Available()).To(Equal(0))

		ExpectSizeError(11)(buf.Write([]byte("hello world")))
		Expect(buf.Size()).To(Equal(11))
		Expect(buf.Available()).To(Equal(11))
		Expect(buf.WriterIndex()).To(Equal(11))

		ExpectSizeError(3)(buf.Write([]byte("yes")))
		Expect(buf.Size()).To(Equal(14))
		Expect(buf.Available()).To(Equal(14))
		Expect(buf.WriterIndex()).To(Equal(14))
	})

	It("should skip", func() {
		buf := New(nil, WithAutoGrowMemory(FixedGrow(5)))
		_, _ = buf.Write([]byte("01234567890123456789"))
		buf.SkipRead(5)
		b := make([]byte, 5)
		ExpectSizeError(5)(buf.Peek(0, b))
		Expect(string(b)).To(Equal("56789"))
	})

	It("should read", func() {
		buf := New(nil, WithAutoGrowMemory(FixedGrow(5)))
		_, _ = buf.Write([]byte("01234567890123456789"))

		Expect(buf.Available()).To(Equal(20))
		Expect(buf.Size()).To(Equal(20))
		Expect(buf.ReaderIndex()).To(Equal(0))

		b := make([]byte, 5)
		ExpectSizeError(5)(buf.Peek(0, b))
		Expect(string(b)).To(Equal("01234"))
		Expect(buf.Available()).To(Equal(20))
		Expect(buf.Size()).To(Equal(20))
		Expect(buf.ReaderIndex()).To(Equal(0))

		b = make([]byte, 4)
		ExpectSizeError(4)(buf.Peek(1, b))
		Expect(string(b)).To(Equal("1234"))
		Expect(buf.Available()).To(Equal(20))
		Expect(buf.Size()).To(Equal(20))
		Expect(buf.ReaderIndex()).To(Equal(0))

		b = make([]byte, 5)
		ExpectSizeError(5)(buf.Read(b))
		Expect(string(b)).To(Equal("01234"))
		Expect(buf.Available()).To(Equal(15))
		Expect(buf.Size()).To(Equal(20))
		Expect(buf.ReaderIndex()).To(Equal(5))

		b = make([]byte, 20)
		ExpectSizeError(15)(buf.Read(b))
		Expect(string(b)).To(Equal("567890123456789\x00\x00\x00\x00\x00"))
		Expect(buf.Available()).To(Equal(0))
		Expect(buf.Size()).To(Equal(20))
		Expect(buf.ReaderIndex()).To(Equal(20))
	})
})
