package problem27

import (
  "euler/utils"
  "set"
)

func mostPrimesForAB(primesSet mapset.Set, a, b int) int {
  n := 0
  for {
    if primesSet.Contains(n * n + a * n + b) {
      n++
    } else {
      return n
    }
  }
  return -1
}

func QuadraticPrimes(mina, maxa, maxb int) int {
  bs := utils.Sieve(maxb)
  primes := utils.Sieve(maxb*maxb)
  primesSet := mapset.NewSet()
  for _, prime := range primes {
    primesSet.Add(prime)
  }
  besta, bestb, bestn := 1, 1, 0
  for a := mina; a <= maxa; a++ {
    for _, b := range bs {
      n := mostPrimesForAB(primesSet, a,b)
      if n > bestn {
        besta, bestb, bestn = a, b, n
      }
    }
    for _, b := range bs {
      b = -b
      n := mostPrimesForAB(primesSet, a,b)
      if n > bestn {
        besta, bestb, bestn = a, b, n
      }
    }
  }

  return besta * bestb
}
