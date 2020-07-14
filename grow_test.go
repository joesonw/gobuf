package gobuf

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grow", func() {
	It("should grow fixed", func() {
		Expect(FixedGrow(5)(5, 19)).To(Equal(20))
		Expect(FixedGrow(5)(5, 21)).To(Equal(25))
		Expect(FixedGrow(5)(5, 6)).To(Equal(10))
	})

	It("should grow multiply", func() {
		Expect(MultiplyGrow(2.4)(5, 10)).To(Equal(12))
		Expect(MultiplyGrow(2.4)(5, 13)).To(Equal(29))
	})
})
