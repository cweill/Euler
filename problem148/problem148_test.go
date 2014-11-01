package problem148_test

import (
	. "euler/problem148"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem148", func() {
  Describe("NonDivisibleBySevenCountForRow", func(){
    It("is 960 for row 8992", func() {
      Expect(NonDivisibleBySevenCountForRow(8992)).To(Equal(960))
    })
  })

  Describe("NonDivisibleBySeven", func() {
    It("is 28 for row 7", func () {
      Expect(NonDivisibleBySeven(7)).To(Equal(28))
    })

    It("is 2361 for row 100", func () {
      Expect(NonDivisibleBySeven(100)).To(Equal(2361))
    })
  })

  It("answer", func () {
    Expect(NonDivisibleBySeven(1000000000)).To(Equal(2129970655314432))
  })
})
