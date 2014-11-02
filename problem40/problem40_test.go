package problem40_test

import (
	. "euler/problem40"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem40", func() {
Describe("ChampernownesConstant", func() {
It("works", func() {
  Expect(ChampernownesConstant(21)).To(Equal("0.123456789101112131415161718192021"))
})
})

It("answers", func(){
  Expect(MultipleOfChampernownesConstantDigits()).To(Equal(210))
})
})
