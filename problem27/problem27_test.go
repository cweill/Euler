package problem27_test

import (
	. "euler/problem27"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem27", func() {
  Describe("QuadraticPrimes", func() {
    It("is 41 for a=1", func () {
      Expect(QuadraticPrimes(0, 1, 100)).To(Equal(41))
    })

    It("is -126479 for a=79", func () {
      Expect(QuadraticPrimes(-79, 79, 2000)).To(Equal(-126479))
    })
  })

  It("answer", func() {
    Expect(QuadraticPrimes(-1000, 1000, 1000)).To(Equal(-59231))
  })
})
