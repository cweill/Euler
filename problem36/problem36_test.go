package problem36_test

import (
	. "euler/problem36"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem36", func() {
  Describe("IsPalindrome", func() {
    It("works for 585", func () {
      Expect(IsPalindrome("585")).To(BeTrue())
    })

    It("works for 1001001001", func () {
      Expect(IsPalindrome("1001001001")).To(BeTrue())
    })
  })

  It("answers", func () {
    Expect(DoubleBasePalindromes()).To(Equal(872187))
  })
})
