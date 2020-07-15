package benchmarks

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	humanize "github.com/dustin/go-humanize"
	"github.com/joesonw/gobuf"
)

var memoryTestTable = []struct {
	memory func() gobuf.Memory
	name   string
}{{
	func() gobuf.Memory { return gobuf.NewSliceMemory(nil, gobuf.FixedGrow(1024*256)) }, "Slice Grow:256K",
}, {
	func() gobuf.Memory { return gobuf.NewListMemory(nil, gobuf.FixedGrow(1024*256)) }, "List Grow:256K",
}, {
	func() gobuf.Memory { return gobuf.NewSliceMemory(nil, gobuf.FixedGrow(1024*1024)) }, "Slice Grow:1M",
}, {
	func() gobuf.Memory { return gobuf.NewListMemory(nil, gobuf.FixedGrow(1024*1024)) }, "List Grow:1M",
}}

var memoryTestSizes = []int{16, 256, 1024, 1024 * 4}

func memoryWriteGoBuffer(b *testing.B, buf *bytes.Buffer, size int) {
	for i := 0; i < b.N/size; i++ {
		body := make([]byte, size)
		_, err := buf.Write(body)
		if err != nil {
			b.Error(err)
			b.Fail()
		}
	}
}

func memoryWriteMemory(b *testing.B, mem gobuf.Memory, size int) {
	index := 0
	for i := 0; i < b.N/size; i++ {
		body := make([]byte, size)
		err := mem.Write(index, body)
		if err != nil {
			b.Error(err)
			b.Fail()
		}
		index += size
	}
}

func BenchmarkMemory(b *testing.B) {
	b.SetParallelism(1)
	for _, size := range memoryTestSizes {
		sizeName := strings.ReplaceAll(humanize.Bytes(uint64(size)), " ", "")

		b.Run(fmt.Sprintf("GoBuffer Size:%s", sizeName), func(b *testing.B) {

			b.Run("Write", func(b *testing.B) {
				buf := bytes.NewBuffer(nil)
				memoryWriteGoBuffer(b, buf, size)
			})

			b.Run("Read", func(b *testing.B) {
				b.StopTimer()
				buf := bytes.NewBuffer(nil)
				memoryWriteGoBuffer(b, buf, size)
				b.StartTimer()

				for i := 0; i < b.N/size; i++ {
					body := make([]byte, size)
					_, err := buf.Read(body)
					if err != nil {
						b.Error(err)
						b.Fail()
					}
				}
			})
		})

		for _, test := range memoryTestTable {
			b.Run(fmt.Sprintf("%s Size:%s", test.name, sizeName), func(b *testing.B) {

				b.Run("Write", func(b *testing.B) {
					mem := test.memory()
					memoryWriteMemory(b, mem, size)
				})

				b.Run("Read", func(b *testing.B) {
					b.StopTimer()
					mem := test.memory()
					memoryWriteMemory(b, mem, size)
					b.StartTimer()

					index := 0
					for i := 0; i < b.N/size; i++ {
						body := make([]byte, size)
						err := mem.Read(index, body)
						if err != nil {
							b.Error(err)
							b.Fail()
						}
						index += size
					}
				})
			})
		}
	}
}
