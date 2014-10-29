package problem17_test

import (
	. "euler/problem17"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem17", func() {
	It("314", func() {
		Expect(NumberLetterCount(314)).To(Equal(23))
	})

	It("115", func() {
		Expect(NumberLetterCount(115)).To(Equal(20))
	})

	It("Answers", func() {
		total := 0
		for i := 1; i <= 1000; i++ {
			total += NumberLetterCount(i)
		}
		fmt.Println(total)
	})
})
