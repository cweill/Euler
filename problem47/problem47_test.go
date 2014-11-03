package problem47_test

import (
	. "euler/problem47"
  "euler/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem47", func() {
  Describe("DistinctPrimeFactorCount", func() {
    max := 100
    primeSet := utils.PrimeSet(max)
    primes := utils.Sieve(max)

    It("is 2 for 14 and 15", func(){
      Expect(DistinctPrimeFactorCount(14,primes,primeSet)).To(Equal(2))
      Expect(DistinctPrimeFactorCount(15,primes,primeSet)).To(Equal(2))
    })

    It("is 3 for 644, 645, and 646", func(){
      Expect(DistinctPrimeFactorCount(644,primes,primeSet)).To(Equal(3))
      Expect(DistinctPrimeFactorCount(645,primes,primeSet)).To(Equal(3))
      Expect(DistinctPrimeFactorCount(646,primes,primeSet)).To(Equal(3))
    })
  })

  Describe("NDistinctPrimeFactors", func() {
    It("is 14 for 2", func(){
      Expect(NDistinctPrimeFactors(2)).To(Equal(14))
    })

    It("is 644 for 3", func(){
      Expect(NDistinctPrimeFactors(3)).To(Equal(644))
    })
  })

  It("answers", func(){
    Expect(NDistinctPrimeFactors(4)).To(Equal(134043))
  })
})
