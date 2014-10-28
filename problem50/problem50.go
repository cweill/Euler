package main

import (
	"euler/sieve"
	"fmt"
	"set"
)

func consecutivePrimeSum(max int) int {
	primes_arr := sieve.Sieve(max)
	primes := mapset.NewSet()
	for i := 0; i < len(primes_arr); i++ {
		primes.Add(primes_arr[i])
	}
	best_prime := 0
	best_terms := 0
	for i := 0; i < len(primes_arr); i++ {
		total := 0
		terms := 0
		last_prime := 0
		for j := 0; j < len(primes_arr); j++ {
			prime := primes_arr[j]
			if prime >= primes_arr[i] {
				total += prime
				if total > max {
					break
				}
				terms++
				if primes.Contains(total) {
					last_prime = total
					if terms >= best_terms {
						best_terms = terms
						best_prime = last_prime
					}
				}
			}
		}
	}
	return best_prime
}

func main() {
	fmt.Println(consecutivePrimeSum(1000000))
}
