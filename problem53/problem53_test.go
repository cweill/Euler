package problem53_test

import (
	. "euler/problem53"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
)

var _ = Describe("Problem53", func() {
	Describe("Factorial", func() {
		It("works", func() {
			mem := make([]*big.Int, 100)
			Expect(Factorial(2, mem)).To(Equal(big.NewInt(2)))
		})

		It("works", func() {
			mem := make([]*big.Int, 100)
			Expect(Factorial(5, mem)).To(Equal(big.NewInt(120)))
		})
	})

	Describe("Choose", func() {
		It("works", func() {
			mem := make([]*big.Int, 100)
			Expect(Choose(5, 3, mem)).To(Equal(big.NewInt(10)))
		})

		It("works", func() {
			mem := make([]*big.Int, 100)
			Expect(Choose(23, 10, mem)).To(Equal(big.NewInt(1144066)))
		})
	})

	It("answers", func() {
		Expect(CombinatoricSelections()).To(Equal(4075))
	})
})
