package problem22

func NameScore(name string, row int) int {
  sum := 0
  for i:= 0; i< len(name);i++ {
    sum += int(name[i] - 64)
  }
  return sum * row
}

func TotalNameScores(names []string) int {
  total :=0
  for i, name := range names {
    total += NameScore(name, i+1)
  }
  return total
}
