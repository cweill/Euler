package problem29

import (
  "set"
  "math/big"
)


func DistinctPowers(maxa, maxb int) int {
  set := mapset.NewSet()
  count :=0
  for a:= 2; a <= maxa; a++ {
    for b:= 2; b <= maxb ; b++ {
      var num big.Int
      num.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
      if !set.Contains(num.String()) {
        set.Add(num.String())
        count++
      }
    }
  }
  return count
}
