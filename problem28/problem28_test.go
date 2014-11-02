package problem28_test

import (
	. "euler/problem28"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem28", func() {
  Describe("CaculateSpiralDiagonalSum", func() {
    It("is 1 for 1", func(){
      Expect(CaculateSpiralDiagonalSum(1)).To(Equal(1))
    })

    It("is 1 for 3", func(){
      Expect(CaculateSpiralDiagonalSum(3)).To(Equal(25))
    })

    It("is 1 for 5", func(){
      Expect(CaculateSpiralDiagonalSum(5)).To(Equal(101))
    })
  })

  It("answers", func(){
    Expect(CaculateSpiralDiagonalSum(1001)).To(Equal(669171001))
  })
})
