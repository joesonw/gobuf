package gobuf

import (
	. "github.com/onsi/gomega"
)

func ExpectSizeError(expected int) func(int, error) {
	return func(n int, err error) {
		Expect(err).To(BeNil())
		Expect(n).To(Equal(expected))
	}
}
