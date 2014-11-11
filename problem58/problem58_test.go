package problem58_test

import (
	. "euler/problem58"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem58", func() {
	Describe("PrimeCountForSpiral", func() {

		It("works", func() {
			Expect(PrimeCountForSpiral(1)).To(Equal(0))
		})

		It("works", func() {
			Expect(PrimeCountForSpiral(3)).To(Equal(3))
		})

		It("works", func() {
			Expect(PrimeCountForSpiral(5)).To(Equal(2))
		})

		It("works", func() {
			Expect(PrimeCountForSpiral(7)).To(Equal(3))
		})
	})

	It("answers", func() {
		Expect(SpiralPrimes()).To(Equal(26241))
	})
})
