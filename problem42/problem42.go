package problem42

import (
  "euler/utils"
  "set"
)

func TriangleNumbers(maxn int) []int {
  var numbers []int
  for i := 1; i <= maxn; i++ {
    numbers = append(numbers, i * (i+1)/2)
  }
  return numbers
}

func CodedTriangleNumbers(names []string) int {
  count := 0
  set := mapset.NewSet()
  for _, triangle := range TriangleNumbers(100000) {
    set.Add(triangle)
  }
  for _, name := range names {
    if set.Contains(utils.WordValue(name)) {
      count++
    }
  }
  return count
}
