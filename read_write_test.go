package gobuf

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReadWrite", func() {
	It("should write/read bool", func() {
		b := New(nil, WithAutoGrowMemory(FixedGrow(32)))

		Expect(b.WriteBool(true)).To(BeNil())
		Expect(b.WriteBool(false)).To(BeNil())

		r, err := b.PeekBool()
		Expect(err).To(BeNil())
		Expect(r).To(BeTrue())

		r, err = b.PeekBool(1)
		Expect(err).To(BeNil())
		Expect(r).To(BeFalse())

		r, err = b.ReadBool()
		Expect(err).To(BeNil())
		Expect(r).To(BeTrue())

		r, err = b.ReadBool()
		Expect(err).To(BeNil())
		Expect(r).To(BeFalse())
	})

	//nolint:dupl
	It("should write/read uints", func() {
		b := New(nil, WithAutoGrowMemory(FixedGrow(32)))

		Expect(b.WriteUint8(8)).To(BeNil())
		Expect(b.WriteUint16(16)).To(BeNil())
		Expect(b.WriteUint32(32)).To(BeNil())
		Expect(b.WriteUint64(64)).To(BeNil())

		u8, err := b.PeekUint8()
		Expect(err).To(BeNil())
		Expect(u8).To(Equal(uint8(8)))

		u16, err := b.PeekUint16(1)
		Expect(err).To(BeNil())
		Expect(u16).To(Equal(uint16(16)))

		u32, err := b.PeekUint32(3)
		Expect(err).To(BeNil())
		Expect(u32).To(Equal(uint32(32)))

		u64, err := b.PeekUint64(7)
		Expect(err).To(BeNil())
		Expect(u64).To(Equal(uint64(64)))

		u8, err = b.ReadUint8()
		Expect(err).To(BeNil())
		Expect(u8).To(Equal(uint8(8)))

		u16, err = b.ReadUint16()
		Expect(err).To(BeNil())
		Expect(u16).To(Equal(uint16(16)))

		u32, err = b.ReadUint32()
		Expect(err).To(BeNil())
		Expect(u32).To(Equal(uint32(32)))

		u64, err = b.ReadUint64()
		Expect(err).To(BeNil())
		Expect(u64).To(Equal(uint64(64)))
	})

	//nolint:dupl
	It("should write/read ints", func() {
		b := New(nil, WithAutoGrowMemory(FixedGrow(32)))

		Expect(b.WriteInt8(8)).To(BeNil())
		Expect(b.WriteInt16(16)).To(BeNil())
		Expect(b.WriteInt32(32)).To(BeNil())
		Expect(b.WriteInt64(64)).To(BeNil())

		i8, err := b.PeekInt8()
		Expect(err).To(BeNil())
		Expect(i8).To(Equal(int8(8)))

		i16, err := b.PeekInt16(1)
		Expect(err).To(BeNil())
		Expect(i16).To(Equal(int16(16)))

		i32, err := b.PeekInt32(3)
		Expect(err).To(BeNil())
		Expect(i32).To(Equal(int32(32)))

		i64, err := b.PeekInt64(7)
		Expect(err).To(BeNil())
		Expect(i64).To(Equal(int64(64)))

		i8, err = b.ReadInt8()
		Expect(err).To(BeNil())
		Expect(i8).To(Equal(int8(8)))

		i16, err = b.ReadInt16()
		Expect(err).To(BeNil())
		Expect(i16).To(Equal(int16(16)))

		i32, err = b.ReadInt32()
		Expect(err).To(BeNil())
		Expect(i32).To(Equal(int32(32)))

		i64, err = b.ReadInt64()
		Expect(err).To(BeNil())
		Expect(i64).To(Equal(int64(64)))
	})

	It("should write/read floats", func() {
		b := New(nil, WithAutoGrowMemory(FixedGrow(32)))

		Expect(b.WriteFloat32(32.32)).To(BeNil())
		Expect(b.WriteFloat64(64.64)).To(BeNil())

		f32, err := b.PeekFloat32()
		Expect(err).To(BeNil())
		Expect(f32).To(Equal(float32(32.32)))

		f64, err := b.PeekFloat64(4)
		Expect(err).To(BeNil())
		Expect(f64).To(Equal(float64(64.64)))

		f32, err = b.ReadFloat32()
		Expect(err).To(BeNil())
		Expect(f32).To(Equal(float32(32.32)))

		f64, err = b.ReadFloat64()
		Expect(err).To(BeNil())
		Expect(f64).To(Equal(float64(64.64)))
	})

	It("should write/read strings", func() {
		b := New(nil, WithAutoGrowMemory(FixedGrow(32)))

		Expect(b.WriteString("hello world")).To(BeNil())

		s, err := b.PeekString(5)
		Expect(err).To(BeNil())
		Expect(s).To(Equal("hello"))

		s, err = b.PeekString(5, 6)
		Expect(err).To(BeNil())
		Expect(s).To(Equal("world"))

		s, err = b.ReadString(5)
		Expect(err).To(BeNil())
		Expect(s).To(Equal("hello"))

		s, err = b.ReadString(6)
		Expect(err).To(BeNil())
		Expect(s).To(Equal(" world"))
	})

	It("should read until", func() {
		b := New(nil, WithAutoGrowMemory(FixedGrow(32)))
		Expect(b.WriteString("hello world tomorrow")).To(BeNil())

		read, ok, err := b.ReadUntil([]byte("hello world\ttomorrow "))
		Expect(err).To(BeNil())
		Expect(ok).To(BeFalse())
		Expect(string(read)).To(Equal(""))

		read, ok, err = b.ReadUntil([]byte(" "), []byte("\t"))
		Expect(err).To(BeNil())
		Expect(ok).To(BeTrue())
		Expect(string(read)).To(Equal("hello"))

		read, ok, err = b.ReadUntil([]byte(" "), []byte("\t"))
		Expect(err).To(BeNil())
		Expect(ok).To(BeTrue())
		Expect(string(read)).To(Equal("world"))

		read, ok, err = b.ReadUntil([]byte(" "))
		Expect(err).To(BeNil())
		Expect(ok).To(BeFalse())
		Expect(string(read)).To(Equal(""))
	})
})
