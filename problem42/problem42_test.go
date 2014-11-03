package problem42_test

import (
	. "euler/problem42"
  "euler/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem42", func() {
  Describe("TriangleNumbers", func() {
    It("works", func() {
      triangles := []int { 1, 3, 6, 10, 15, 21, 28, 36, 45, 55 }
      Expect(TriangleNumbers(10)).To(Equal(triangles))
    })
  })

  It("answers", func() {
    words := utils.ReadStringsFromFile("p042_words.txt")
    Expect(CodedTriangleNumbers(words)).To(Equal(162))
  })
})
