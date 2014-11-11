package problem63

import (
	"math/big"
)

func PowerfulDigitCounts() int {
	count := 0
	for i := 1; i <= 20; i++ {
		j := 1
		for {
			number := big.NewInt(0).Exp(big.NewInt(int64(i)), big.NewInt(int64(j)), nil)
			digits := len(number.String())
			if digits == j {
				count++
			} else {
				break
			}
			j++
		}
	}
	return count
}
