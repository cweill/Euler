package problem25_test

import (
	. "euler/problem25"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
)

var _ = Describe("Problem25", func() {
	Describe("Fibonacci", func() {
		It("returns the correct sequence", func() {
			fib := []int64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
			for i := 0; i < len(fib); i++ {
				Expect(Fibonacci(i + 1)).To(Equal(big.NewInt(fib[i])))
			}
		})
	})

	Describe("FirstFibTermWithNDigits", func() {
		It("is 12 for the first fib with 3 digits", func() {
			Expect(FirstFibTermWithNDigits(3)).To(Equal(12))
		})
	})

	It("Answer", func() {
		Expect(FirstFibTermWithNDigits(1000)).To(Equal(4782))
	})
})
