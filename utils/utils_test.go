package utils_test

import (
	. "euler/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Describe("Sieve", func() {
		It("gives empty for less than 1", func() {
			var primes []int
			Expect(Sieve(1)).To(Equal(primes))
		})

		It("gives prime numbers under 20", func() {
			primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
			Expect(Sieve(20)).To(Equal(primes))
		})
	})
})
