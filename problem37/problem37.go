package problem37

import (
  "set"
  "euler/utils"
)

func IsTruncatablePrimeLeftToRight(n int, primes mapset.Set) bool {
  base := 1
  for base < n {
    base *= 10
    if !primes.Contains(n % base) {
      return false
    }
  }
  return true
}

func IsTruncatablePrimeRightToLeft(n int, primes mapset.Set) bool {
  base := 1
  for base < n {
    if !primes.Contains(n / base) {
      return false
    }
    base *= 10
  }
  return true
}

func TruncatablePrimes() int {
  primes := utils.Sieve(1000000)
  primeSet := utils.PrimeSet(1000000)
  sum, count := 0, 0
  for _, prime := range primes {
    if prime > 7 {
      if IsTruncatablePrimeLeftToRight(prime, primeSet) &&
          IsTruncatablePrimeRightToLeft(prime, primeSet) {
        sum += prime
        count++
        if count == 11 {
          return sum
        }
      }
    }
  }
  return -1
}
