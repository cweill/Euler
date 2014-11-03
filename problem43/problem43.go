package problem43

import (
  "strconv"
  "permute"
)

func IsSubstringDivisible(n string) bool {
  primes := []int {1,2,3,5,7,11,13,17}
  for i := 1; i<8; i++ {
    num, _ := strconv.Atoi(n[i:i+3])
    if num % primes[i] != 0 {
      return false
    }
  }
  return true
}

func SubstringDivisibility() int {
  sum := 0
  last := "0123456789"
  current := permute.NextPermutation(last)
  for last != current {
    if IsSubstringDivisible(current) {
      num, _ := strconv.Atoi(current)
      sum += num
    }
    last = current
    current = permute.NextPermutation(last)
  }
  return sum
}
