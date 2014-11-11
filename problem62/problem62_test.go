package problem62_test

import (
	. "euler/problem62"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem62", func() {
	Describe("CubicPermutations", func() {
		It("works", func() {
			Expect(CubicPermutations(3)).To(Equal(41063625))
		})
	})

	It("answers", func() {
		Expect(CubicPermutations(5)).To(Equal(127035954683))
	})
})
