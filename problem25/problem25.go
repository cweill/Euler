package problem25

import (
	"math/big"
)

func FibonacciHelper(n int, mem []*big.Int) *big.Int {
	if n == 1 || n == 2 {
		return big.NewInt(1)
	}
	if mem[n] != nil {
		return mem[n]
	}
	result := big.NewInt(0)
	result.Add(FibonacciHelper(n-1, mem), FibonacciHelper(n-2, mem))
	mem[n] = result
	return result
}

func Fibonacci(n int) *big.Int {
	mem := make([]*big.Int, n+1)
	return FibonacciHelper(n, mem)
}

func FirstFibTermWithNDigits(n int) int {
	mem := make([]*big.Int, n*10+1)
	for i := 1; i <= n*10; i++ {
		numberString := FibonacciHelper(i, mem).String()
		if len(numberString) == n {
			return i
		}
	}
	return -1
}
