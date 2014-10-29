package problem21_test

import (
	. "euler/problem21"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem21", func() {
	Describe("ProperDivisors", func() {
		It("is 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110 for 220", func() {
			divisors := []int{1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110}
			Expect(ProperDivisors(220)).To(Equal(divisors))
		})

		It("is 1, 2, 4, 71 and 142 for 284", func() {
			divisors := []int{1, 2, 4, 71, 142}
			Expect(ProperDivisors(284)).To(Equal(divisors))
		})
	})

	Describe("IsAmicable", func() {
		It("true for 220", func() {
			Expect(IsAmicable(220)).To(BeTrue())
		})

		It("true for 284", func() {
			Expect(IsAmicable(284)).To(BeTrue())
		})

		It("false for 283", func() {
			Expect(IsAmicable(283)).To(BeFalse())
		})

		It("false for 6", func() {
			Expect(IsAmicable(6)).To(BeFalse())
		})
	})

	It("Answer", func() {
		Expect(SumAmicables(10000)).To(Equal(31626))
	})
})
