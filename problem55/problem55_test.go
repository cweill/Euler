package problem55_test

import (
	. "euler/problem55"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
)

var _ = Describe("Problem55", func() {
	Describe("Reverse", func() {
		It("works", func() {
			Expect(Reverse("123456")).To(Equal("654321"))
		})
	})

	Describe("IsPalindrome", func() {
		It("works", func() {
			Expect(IsPalindrome("123456")).To(BeFalse())
			Expect(IsPalindrome("4668731596684224866951378664")).To(BeTrue())
		})
	})

	Describe("ReverseAndAdd", func() {
		It("works", func() {
			Expect(ReverseAndAdd(big.NewInt(47))).To(Equal(big.NewInt(121)))
		})
	})

	Describe("IsLychrel", func() {
		It("is not for 349", func() {
			Expect(IsLychrel(349)).To(BeFalse())
		})

		It("is for 196", func() {
			Expect(IsLychrel(196)).To(BeTrue())
		})

		It("is for 10677", func() {
			Expect(IsLychrel(10677)).To(BeTrue())
		})

		It("is for 4994", func() {
			Expect(IsLychrel(4994)).To(BeTrue())
		})
	})

	It("answers", func() {
		Expect(LychrelNumbers()).To(Equal(249))
	})
})
