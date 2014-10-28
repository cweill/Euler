package sieve

import "math"

func Sieve(n int) []int {
	isPrimeArray := make([]bool, n)
	for i := 0; i < len(isPrimeArray); i++ {
		isPrimeArray[i] = true
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if isPrimeArray[i] {
			j := i * i
			counter := 1
			for j < n {
				isPrimeArray[j] = false
				j = i*i + i*counter
				counter++
			}
		}
	}
	var primes []int
	for i := 2; i < len(isPrimeArray); i++ {
		if isPrimeArray[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
