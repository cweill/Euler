package problem41_test

import (
	. "euler/problem41"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem41", func() {
  It("answers", func(){
    Expect(PandigitalPrime()).To(Equal(7652413))
  })
})
