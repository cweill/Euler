package problem44

import (
  "set"
)

func PentagonNumbers(maxn int) []int {
  var pentagons []int
  for i :=1; i<=maxn; i++ {
    pentagons = append(pentagons, i*(3*i-1)/2)
  }
  return pentagons
}

func MinDPentagonalPair() int {
  max := 10000
  pentagons := PentagonNumbers(2*max)
  set := mapset.NewSet()
  for _, pentagon := range pentagons {
    set.Add(pentagon)
  }
  j, k := 0, 1
  for k < max {
    pentj, pentk := pentagons[j], pentagons[k]
    if set.Contains(pentj + pentk) && set.Contains(pentk - pentj) {
      return pentk - pentj
    }
    j++
    k++
    if k == max {
      k = k - j + 1
      j = 0
    }
  }
  return -1
}
