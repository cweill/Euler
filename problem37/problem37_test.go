package problem37_test

import (
	. "euler/problem37"
  "euler/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem37", func() {
  Describe("IsTruncatablePrimeLeftToRight", func() {
    It("works", func() {
      Expect(IsTruncatablePrimeLeftToRight(3797, utils.PrimeSet(4000))).To(BeTrue())
    })

    It("works", func() {
      Expect(IsTruncatablePrimeLeftToRight(1234, utils.PrimeSet(4000))).To(BeFalse())
    })
  })

  Describe("IsTruncatablePrimeRightToLeft", func() {
    It("works", func() {
      Expect(IsTruncatablePrimeRightToLeft(3797, utils.PrimeSet(4000))).To(BeTrue())
    })

    It("works", func() {
      Expect(IsTruncatablePrimeRightToLeft(1234, utils.PrimeSet(4000))).To(BeFalse())
    })
  })

  It("answers", func() {
    Expect(TruncatablePrimes()).To(Equal(748317))
  })
})
