package problem18_test

import (
	. "euler/problem18"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem18", func() {
	It("works for an example", func() {
		sum := MaximumPathSum([][]int{
			[]int{3},
			[]int{7, 4},
			[]int{2, 4, 6},
			[]int{8, 5, 9, 3},
		})
		Expect(sum).To(Equal(23))
	})

	It("Answer", func() {
		pyramid := [][]int{
			[]int{75},
			[]int{95, 64},
			[]int{17, 47, 82},
			[]int{18, 35, 87, 10},
			[]int{20, 4, 82, 47, 65},
			[]int{19, 1, 23, 75, 3, 34},
			[]int{88, 2, 77, 73, 7, 63, 67},
			[]int{99, 65, 4, 28, 6, 16, 70, 92},
			[]int{41, 41, 26, 56, 83, 40, 80, 70, 33},
			[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
			[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
			[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
			[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
			[]int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
			[]int{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
		}
		Expect(MaximumPathSum(pyramid)).To(Equal(1074))
	})
})
