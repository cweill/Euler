package problem148

import (
  "strconv"
)

// http://fbarbuto.wordpress.com/2011/07/16/problem-148/
func NonDivisibleBySevenCountForRow(i int) int {
  // Convert to base 7
  num, _ := strconv.Atoi(strconv.FormatInt(int64(i), 7))

  // Get the product of each (digit+1). That's the trick
  base := 1
  prod := 1
  for base <= num {
    digit :=((num/base) % 10) + 1
    prod *= digit
    base *= 10
  }
  return prod
}

func NonDivisibleBySeven(maxn int) int {
  sum :=0
  for i:=0; i < maxn; i++ {
    sum += NonDivisibleBySevenCountForRow(i)
  }
  return sum
}
