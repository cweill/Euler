package problem51

import (
  "euler/utils"
  "strconv"
)

func PrimeDigitReplacement(num int) int {
  max := 1000000
  primes := utils.Sieve(max)
  primeSet := utils.PrimeSet(max)
  digits := []rune("0123456789")
  for i := 5; i < len(primes); i++ {
    prime := primes[i]
    primeString := []rune(strconv.Itoa(prime))
    for j := 0; j < len(primeString) - 1; j++ {
      for k := j; k < len(primeString) - 1; k++ {
        for l := k; l < len(primeString) - 1; l++ {
          count := 0
          primeString = []rune(strconv.Itoa(prime))
          smallestPrime := 0
          for digit := 0; digit < len(digits); digit++ {
            if j != 0 || digits[digit] != '0' {
              primeString[j] = digits[digit]
              primeString[k] = digits[digit]
              primeString[l] = digits[digit]

              newNum, _ := strconv.Atoi(string(primeString))
              if smallestPrime == 0 {
                smallestPrime = newNum
              }
              if primeSet.Contains(newNum) {
                count++
                if count == num {
                  return smallestPrime
                }
              }
            }
          }
        }
      }
    }
  }
  return -1
}
