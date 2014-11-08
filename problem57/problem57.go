package problem57

import (
	"math/big"
)

func Expansion(n int, mem1, mem2 []*big.Int) (*big.Int, *big.Int) {
	if n == 0 {
		return big.NewInt(1), big.NewInt(1)
	}
	if n == 1 {
		return big.NewInt(3), big.NewInt(2)
	}
	if mem1[n] != nil {
		return mem1[n], mem2[n]
	}
	num1, den1 := Expansion(n-1, mem1, mem2)
	num2, den2 := Expansion(n-2, mem1, mem2)
	result1, result2 := big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(2), num1), num2),
		big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(2), den1), den2)
	mem1[n] = result1
	mem2[n] = result2
	return result1, result2
}

func SquareRootConvergents() int {
	count := 0
	mem1, mem2 := make([]*big.Int, 1000), make([]*big.Int, 1000)
	for i := 0; i < 1000; i++ {
		num, den := Expansion(i, mem1, mem2)
		if len(num.String()) > len(den.String()) {
			count++
		}
	}
	return count
}
