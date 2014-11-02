package problem34_test

import (
	. "euler/problem34"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem34", func() {
  Describe("FactorialDigit", func() {
    It("works", func() {
      Expect(IsSumFactorialOfDigits(145)).To(BeTrue())
    })
  })

  It("answers", func() {
    Expect(DigitFactorials()).To(Equal(40730))
  })
})
