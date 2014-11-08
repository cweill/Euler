package problem52_test

import (
	. "euler/problem52"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem52", func() {
  Describe("PermutedMultiples", func() {
    It("Problem52", func() {
      Expect(PermutedMultiples(2)).To(Equal(125874))
    })
  })

  It("answers", func() {
    Expect(PermutedMultiples(4)).To(Equal(142857))
  })
})
