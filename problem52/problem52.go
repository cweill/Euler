package problem52

import (
  "euler/utils"
)

func PermutedMultiples(n int) int {
  max := 1000000
  for i := 1; i < max ; i++ {
    for j := 1; j <= n; j++ {
      if !utils.HasSameDigits(i, i * j) {
        break
      }
      if j == n {
        return i
      }
    }
  }
  return -1
}
