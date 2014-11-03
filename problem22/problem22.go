package problem22

import (
  "euler/utils"
)

func NameScore(name string, row int) int {
  return utils.WordValue(name) * row
}

func TotalNameScores(names []string) int {
  total :=0
  for i, name := range names {
    total += NameScore(name, i+1)
  }
  return total
}
