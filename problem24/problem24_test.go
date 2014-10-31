package problem24_test

import (
	. "euler/problem24"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem24", func() {
	Describe("NextPermutation", func() {
		It("is 8344112666 after 8342666411", func() {
			Expect(NextPermutation("8342666411")).To(Equal("8344112666"))
		})

		It("permutes 012", func() {
			permutations := []string{"012", "021", "102", "120", "201", "210"}
			for i := 0; i < len(permutations); i++ {
				if i == len(permutations)-1 {
					Expect(NextPermutation(permutations[i])).To(Equal(permutations[i]))
				} else {
					Expect(NextPermutation(permutations[i])).To(Equal(permutations[i+1]))
				}
			}
		})
	})

	Describe("LexicographicPermutations", func() {
		It("outputs 012 in lexicographic order", func() {
			permutations := []string{"012", "021", "102", "120", "201", "210"}
			Expect(LexicographicPermutations("012")).To(Equal(permutations))
		})
	})

	It("answer", func() {
		Expect(LexicographicPermutations("0123456789")[999999]).To(Equal("2783915460"))
	})
})
