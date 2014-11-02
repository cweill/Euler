package problem29_test

import (
	. "euler/problem29"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem29", func() {
  Describe("DistinctPowers", func() {
    It("works", func(){
      Expect(DistinctPowers(5, 5)).To(Equal(15))
    })
  })

    It("answers", func(){
      Expect(DistinctPowers(100, 100)).To(Equal(9183))
    })
})
