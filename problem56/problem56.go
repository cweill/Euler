package problem56

import (
	"math/big"
)

func CountDigits(n *big.Int) int {
	digits := []rune(n.String())
	count := 0
	for _, rune := range digits {
		// Ascii adjustment
		count += int(rune) - 48
	}
	return count
}

func PowerfulDigitSum() int {
	max := 0
	var i, j int64
	for i = 2; i < 100; i++ {
		for j = 2; j < 100; j++ {
			number := big.NewInt(0).Exp(big.NewInt(i), big.NewInt(j), nil)
			sum := CountDigits(number)
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
