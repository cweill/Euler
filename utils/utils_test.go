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

	Describe("IsPrime", func() {
		It("is true for 3367900313", func() {
			Expect(IsPrime(3367900313)).To(BeTrue())
		})

		It("is false for 1000000", func() {
			Expect(IsPrime(1000000)).To(BeFalse())
		})

		// It("is false for 7033313127171317417", func() {
		// 	Expect(IsPrime(7033313127171317417)).To(BeTrue())
		// })
	})

	Describe("GCD", func() {
		It("is 6 for 36, 42", func() {
			Expect(GCD(36, 42)).To(Equal(6))
		})

		It("is 1 for 17, 73", func() {
			Expect(GCD(17, 73)).To(Equal(1))
		})
	})

	Describe("SortString", func() {
		It("is `0123456789` for `2391085647`", func() {
			Expect(SortString("2391085647")).To(Equal("0123456789"))
		})
	})

	Describe("IsPandigital", func() {
		It("is `true` for `239185647`, 9", func() {
			Expect(IsPandigital(239185647, 9)).To(BeTrue())
		})

		It("is `true` for `2143`, 4", func() {
			Expect(IsPandigital(2143, 4)).To(BeTrue())
		})

		It("is `false` for `1111`, 4", func() {
			Expect(IsPandigital(1111, 4)).To(BeFalse())
		})

		It("is `false` for `239185647`, 9", func() {
			Expect(IsPandigital(239185647, 7)).To(BeFalse())
		})

		It("is `false` for `239185647`, 9", func() {
			Expect(IsPandigital(887645321, 9)).To(BeFalse())
		})
	})

	Describe("WordValue", func() {
		It("is 55 for SKY", func() {
			Expect(WordValue("SKY")).To(Equal(55))
		})
	})

	Describe("HasSameDigits", func() {
		It("is true for 125874 and 251748", func() {
			Expect(HasSameDigits(125874, 251748)).To(BeTrue())
		})

		It("is false for 125874 and 2511748", func() {
			Expect(HasSameDigits(125874, 2511748)).To(BeFalse())
		})
	})
})
