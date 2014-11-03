package problem46

import (
  "euler/utils"
  "set"
)

func GoldbachsOtherConjecture() int {
  max := 10000
  primes := utils.Sieve(max)
  primeSet := utils.PrimeSet(max)
  squares := mapset.NewSet()
  for i := 0; i <= max; i++ {
    squares.Add(i * i)
  }
  for n:= 9; n <= max; n+=2 {
    if !primeSet.Contains(n) {
      for _, prime := range primes {
        if prime > n {
          return n
        }
        if squares.Contains((n - prime) / 2) {
          break
        }
      }
    }
  }
  return -1
}
