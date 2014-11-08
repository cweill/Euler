package problem53

import (
	"math/big"
)

func Factorial(n int64, mem []*big.Int) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	if mem[n] != nil {
		return mem[n]
	}
	var result *big.Int
	result = big.NewInt(n).Mul(big.NewInt(n), Factorial(n-1, mem))
	mem[n] = result
	return result
}

func Choose(n, r int64, mem []*big.Int) *big.Int {
	return big.NewInt(n).Div(Factorial(n, mem), big.NewInt(n).Mul(Factorial(r, mem), Factorial(n-r, mem)))
}

func CombinatoricSelections() int {
	count := 0
	mem := make([]*big.Int, 101)
	var n, r int64
	for n = 1; n <= 100; n++ {
		for r = 1; r <= n; r++ {
			if Choose(n, r, mem).Cmp(big.NewInt(1000000)) > 0 {
				count++
			}
		}
	}
	return count
}
