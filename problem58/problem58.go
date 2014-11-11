package problem58

import (
	"euler/utils"
)

func PrimeCountForSpiral(n int) int {
	count := 0
	if utils.IsPrime(n*n - n + 1) {
		count++
	}
	if utils.IsPrime(n*n - 2*(n-1)) {
		count++
	}
	if utils.IsPrime(n*n - 3*(n-1)) {
		count++
	}
	return count
}

func SpiralPrimes() int {
	primeCount := 0
	spiral := 3
	for {
		count := PrimeCountForSpiral(spiral)
		primeCount += count
		if primeCount*100/(spiral*2-1) < 10 {
			return spiral
		}
		spiral += 2
	}
	return primeCount
}
