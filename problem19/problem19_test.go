package problem19_test

import (
	. "euler/problem19"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem19", func() {
  It("answers", func(){
    Expect(CountingSundays()).To(Equal(171))
  })
})
