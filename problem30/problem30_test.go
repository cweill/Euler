package problem30_test

import (
	. "euler/problem30"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem30", func() {
  Describe("Problem30", func() {
    It("DigitNPowers", func() {
      Expect(DigitNPowers(4)).To(Equal(19316))
    })
  })

  It("answers", func() {
    Expect(DigitNPowers(5)).To(Equal(443839))
  })
})
