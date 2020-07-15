package gobuf

import (
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Memory", func() {
	Describe("SliceMemory", func() {
		It("should read", func() {
			m := NewSliceMemory([]byte("hello"), nil)

			b := make([]byte, 5)
			err := m.Read(0, b)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal("hello"))

			b = make([]byte, 3)
			err = m.Read(2, b)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal("llo"))

			b = make([]byte, 3)
			err = m.Read(3, b)
			Expect(err).To(Equal(io.EOF))
		})

		It("should write", func() {
			m := NewSliceMemory([]byte("hello"), nil)

			err := m.Write(0, []byte("world"))
			Expect(err).To(BeNil())
			Expect(string(m.buf)).To(Equal("world"))
			Expect(m.Length()).To(Equal(5))

			err = m.Write(2, []byte("aaa"))
			Expect(err).To(BeNil())
			Expect(string(m.buf)).To(Equal("woaaa"))

			err = m.Write(2, []byte("aaaa"))
			Expect(err).To(Equal(ErrOutOfSpace))
		})
	})

	Describe("SliceMemory canGrow", func() {
		It("should write", func() {
			m := NewSliceMemory([]byte("hello"), FixedGrow(5))

			err := m.Write(0, []byte("world"))
			Expect(err).To(BeNil())
			Expect(string(m.buf)).To(Equal("world"))

			err = m.Write(2, []byte("aaa"))
			Expect(err).To(BeNil())
			Expect(string(m.Bytes())).To(Equal("woaaa"))
			Expect(m.Length()).To(Equal(5))

			err = m.Write(0, []byte("hello world"))
			Expect(err).To(BeNil())
			Expect(string(m.buf)).To(Equal("hello world\x00\x00\x00\x00"))
			Expect(cap(m.buf)).To(Equal(15))
			Expect(m.Length()).To(Equal(15))

			m = NewSliceMemory([]byte("hello"), FixedGrow(5))
			err = m.Write(6, []byte("world"))
			Expect(err).To(BeNil())
			Expect(string(m.buf)).To(Equal("hello\x00world\x00\x00\x00\x00"))
			Expect(cap(m.buf)).To(Equal(15))

			err = m.Write(16, []byte("aaaa"))
			Expect(err).To(BeNil())
			Expect(string(m.buf)).To(Equal("hello\x00world\x00\x00\x00\x00\x00aaaa"))
			Expect(cap(m.buf)).To(Equal(20))
		})
	})

	Describe("LinkedListMemory", func() {
		It("should write", func() {
			m := NewLinkedListMemory([]byte("hello"), FixedGrow(5))

			err := m.Write(0, []byte("world"))
			Expect(err).To(BeNil())
			Expect(string(m.start.buf)).To(Equal("world"))
			Expect(m.Length()).To(Equal(5))

			err = m.Write(0, []byte("hello world"))
			Expect(err).To(BeNil())
			Expect(string(m.Bytes())).To(Equal("hello world\x00\x00\x00\x00"))
			Expect(len(m.start.buf)).To(Equal(5))
			Expect(string(m.start.buf)).To(Equal("hello"))
			Expect(len(m.start.next.buf)).To(Equal(10))
			Expect(string(m.start.next.buf)).To(Equal(" world\x00\x00\x00\x00"))
			Expect(m.Length()).To(Equal(15))

			err = m.Write(11, []byte("aaaa"))
			Expect(err).To(BeNil())
			Expect(string(m.Bytes())).To(Equal("hello worldaaaa"))
			Expect(len(m.start.buf)).To(Equal(5))
			Expect(string(m.start.buf)).To(Equal("hello"))
			Expect(len(m.start.next.buf)).To(Equal(10))
			Expect(string(m.start.next.buf)).To(Equal(" worldaaaa"))

			err = m.Write(0, []byte("01234567890123456789"))
			Expect(err).To(BeNil())
			Expect(string(m.Bytes())).To(Equal("01234567890123456789"))
			Expect(len(m.start.buf)).To(Equal(5))
			Expect(string(m.start.buf)).To(Equal("01234"))
			Expect(len(m.start.next.buf)).To(Equal(10))
			Expect(string(m.start.next.buf)).To(Equal("5678901234"))
			Expect(len(m.start.next.next.buf)).To(Equal(5))
			Expect(string(m.start.next.next.buf)).To(Equal("56789"))
			Expect(m.Length()).To(Equal(20))

			m = NewLinkedListMemory([]byte("hello"), FixedGrow(5))
			err = m.Write(6, []byte("world"))
			Expect(err).To(BeNil())
			Expect(string(m.Bytes())).To(Equal("hello\x00world\x00\x00\x00\x00"))
			Expect(len(m.start.buf)).To(Equal(5))
			Expect(string(m.start.buf)).To(Equal("hello"))
			Expect(len(m.start.next.buf)).To(Equal(10))
			Expect(string(m.start.next.buf)).To(Equal("\x00world\x00\x00\x00\x00"))
		})

		It("should read", func() {
			m := &LinkedListMemory{
				start: &linkedListMemoryNode{
					buf: []byte("01234"),
					next: &linkedListMemoryNode{
						buf: []byte("56789"),
						next: &linkedListMemoryNode{
							buf: []byte("0123456789"),
						},
					},
				},
			}

			b := make([]byte, 5)
			err := m.Read(0, b)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal("01234"))

			b = make([]byte, 30)
			err = m.Read(0, b)
			Expect(err).To(Equal(io.EOF))

			b = make([]byte, 20)
			err = m.Read(0, b)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal("01234567890123456789"))

			b = make([]byte, 5)
			err = m.Read(3, b)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal("34567"))

			b = make([]byte, 10)
			err = m.Read(3, b)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal("3456789012"))

			b = make([]byte, 5)
			err = m.Read(12, b)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal("23456"))
		})
	})
})
