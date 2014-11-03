package problem41

import (
  "euler/utils"
)

func PandigitalPrime()int{
  primes := utils.Sieve(987654321)
  min := 123456789
  for n := 9; n >=1 ; n-- {
    for i := len(primes)-1; i >=0 && primes[i] >= min ; i-- {
      prime := primes[i]
      if utils.IsPandigital(uint(prime), uint(n)) {
        return prime
      }
      min /= 10
    }
  }
  return -1
}
