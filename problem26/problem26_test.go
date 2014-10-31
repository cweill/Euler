package problem26_test

import (
	. "euler/problem26"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem26", func() {
	It("Answers", func() {
		Expect(ReciprocalCycles()).To(Equal(983))
	})
})
