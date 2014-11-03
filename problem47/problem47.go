package problem47

import (
  "euler/utils"
  "set"
)

func DistinctPrimeFactorCount(n int, primes []int, primeSet mapset.Set) int {
  if primeSet.Contains(n) {
    return 0
  }
  count := 0
  for _, prime := range primes {
    if prime > n {
      break
    }
    if n % prime == 0 {
      count++
      n = n / prime
    }
  }
  return count
}

func NDistinctPrimeFactors(n int) int {
  max := 1000000
  primeSet := utils.PrimeSet(max)
  primes := utils.Sieve(max)
  for i:=2; i<= max; i++ {
    correct := true
    for j := 0; j < n; j++ {
      count := DistinctPrimeFactorCount(i + j, primes, primeSet)
      if count != n {
        correct = false
        break
      }
    }
    if correct {
      return i
    }
  }
  return -1
}
