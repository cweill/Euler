package problem23

import (
	"euler/problem21"
	"set"
)

func AbundantNumbers(maxn int) []int {
	var abundant []int
	for n := 1; n <= maxn; n++ {
		iDivisorSum := problem21.SumProperDivisors(n)
		if iDivisorSum > n {
			abundant = append(abundant, n)
		}
	}
	return abundant
}

func IsNotSumOfAbundantNumber(n int, abundantNumbers []int, abundantNumbersSet mapset.Set) bool {
	abundant := true
	for j := 0; j < len(abundantNumbers); j++ {
		number := n - abundantNumbers[j]
		if number < 0 {
			break
		} else if abundantNumbersSet.Contains(number) {
			abundant = false
			break
		}
	}
	return abundant
}

func SumOfNonAbundantNumbers() int {
	limit := 28123
	abundantNumbers := AbundantNumbers(limit)
	abundantNumbersSet := mapset.NewSet()
	for i := 0; i < len(abundantNumbers); i++ {
		abundantNumbersSet.Add(abundantNumbers[i])
	}
	sum := 0
	for i := 1; i <= limit; i++ {
		if IsNotSumOfAbundantNumber(i, abundantNumbers, abundantNumbersSet) {
			sum += i
		}
	}
	return sum
}
