package problem57_test

import (
	. "euler/problem57"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
)

var _ = Describe("Problem57", func() {
	Describe("Expansion", func() {
		It("works", func() {
			mem1, mem2 := make([]*big.Int, 100), make([]*big.Int, 100)
			num, den := Expansion(3, mem1, mem2)
			Expect(num).To(Equal(big.NewInt(17)))
			Expect(den).To(Equal(big.NewInt(12)))

			num, den = Expansion(5, mem1, mem2)
			Expect(num).To(Equal(big.NewInt(99)))
			Expect(den).To(Equal(big.NewInt(70)))
		})
	})

	It("answers", func() {
		Expect(SquareRootConvergents()).To(Equal(153))
	})
})
