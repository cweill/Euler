package problem56_test

import (
	. "euler/problem56"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
)

var _ = Describe("Problem56", func() {
	Describe("CountDigits", func() {
		It("works", func() {
			Expect(CountDigits(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(100), nil))).To(Equal(1))
		})
	})

	It("answers", func() {
		Expect(PowerfulDigitSum()).To(Equal(972))
	})
})
