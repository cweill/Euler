package problem45_test

import (
	. "euler/problem45"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem45", func() {
  It("answers", func(){
    Expect(TriangularPentagonalAndHexagonal(1)).To(Equal(40755))
    Expect(TriangularPentagonalAndHexagonal(165)).To(Equal(1533776805))
  })
})
