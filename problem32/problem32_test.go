package problem32_test

import (
	. "euler/problem32"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem32", func() {
  Describe("Problem32", func() {
    It("works", func() {
      Expect(IsPandigital(39, 186)).To(BeTrue())
    })
  })

  It("answers", func() {
    Expect(PandigitalProduct()).To(Equal(45228))
  })
})
