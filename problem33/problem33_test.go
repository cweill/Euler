package problem33_test

import (
	. "euler/problem33"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem33", func() {

  It("answers", func(){
    Expect(DigitCancellingFractions()).To(Equal(100.0))
  })
})
