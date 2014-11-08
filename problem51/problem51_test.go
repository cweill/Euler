package problem51_test

import (
	. "euler/problem51"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem51", func() {
  Describe("PrimeDigitReplacement", func() {
    It("is 13 for 2", func() {
      Expect(PrimeDigitReplacement(6)).To(Equal(13))
    })

    It("is 56003 for 7", func() {
      Expect(PrimeDigitReplacement(7)).To(Equal(56003))
    })
  })

  It("answers", func() {
    Expect(PrimeDigitReplacement(8)).To(Equal(121313))
  })
})
