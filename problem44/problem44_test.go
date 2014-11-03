package problem44_test

import (
	. "euler/problem44"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem44", func() {
  Describe("PentagonNumbers", func() {
    It("works", func() {
      pentagons := []int {1, 5, 12, 22, 35, 51, 70, 92, 117, 145}
      Expect(PentagonNumbers(10)).To(Equal(pentagons))
    })
  })

  It("answer", func() {
    Expect(MinDPentagonalPair()).To(Equal(5482660))
  })
})
