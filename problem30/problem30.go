package problem30

import (
  "math"
)

func DigitNPowers(power int) int{
  sum := 0
  for n := 2; n < 500000 ; n++ {
    base := 1
    digitsum := 0
    for base < n {
      num := n / base % 10
      digitsum += int(math.Pow(float64(num), float64(power)))
      base *= 10
    }
    if digitsum == n {
      sum += n
    }
  }
  return sum
}
