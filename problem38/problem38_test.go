package problem38_test

import (
	. "euler/problem38"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem38", func() {
  Describe("ConcatenateProducts", func(){
    It("is 192384576 for 192", func(){
      Expect(ConcatenateProducts(192)).To(Equal(192384576))
    })

    It("is 918273645 for 9", func(){
      Expect(ConcatenateProducts(9)).To(Equal(918273645))
    })
  })

  It("answers", func(){
    Expect(PandigitalMultiples()).To(Equal(932718654))
  })
})
