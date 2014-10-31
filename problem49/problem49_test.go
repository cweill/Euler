package problem49_test

import (
	. "euler/problem49"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem49", func() {
  It("Answers", func() {
    Expect(PrimePermutations()).To(Equal("296962999629"))
  })
})
