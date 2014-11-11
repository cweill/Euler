package problem63_test

import (
	. "euler/problem63"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem63", func() {
	It("answers", func() {
		Expect(PowerfulDigitCounts()).To(Equal(49))
	})
})
