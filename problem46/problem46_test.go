package problem46_test

import (
	. "euler/problem46"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem46", func() {
  It("answers", func(){
    Expect(GoldbachsOtherConjecture()).To(Equal(5777))
  })
})
