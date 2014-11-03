package problem43_test

import (
	. "euler/problem43"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem43", func() {
  Describe("IsSubstringDivisible", func() {
    It("works", func () {
      Expect(IsSubstringDivisible("1406357289")).To(BeTrue())
    })
  })

  It("answers", func() {
    Expect(SubstringDivisibility()).To(Equal(16695334890))
  })
})
