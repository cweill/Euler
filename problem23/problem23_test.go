package problem23_test

import (
	. "euler/problem23"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem23", func() {
	Describe("AbundantNumbers", func() {
		It("is just [12] for 12", func() {
			Expect(AbundantNumbers(12)).To(Equal([]int{12}))
		})
	})

	It("Answer", func() {
		Expect(SumOfNonAbundantNumbers()).To(Equal(4179871))
	})
})
