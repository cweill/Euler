package problem36

import (
  "strconv"
)

func IsPalindrome(str string) bool {
  runes := []rune(str)
  i, j := 0, len(str) - 1
  for j > i {
    if runes[i] != runes[j] {
      return false
    }
    i++
    j--
  }
  return true
}

func DoubleBasePalindromes() int {
  sum := 0
  for i:=1; i < 1000000; i++ {
    if IsPalindrome(strconv.Itoa(i)) && IsPalindrome(strconv.FormatInt(int64(i), 2)) {
      sum += i
    }
  }
  return sum
}
