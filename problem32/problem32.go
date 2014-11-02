package problem32

import (
  "sort"
  "strconv"
  "strings"
  "set"
)

// Sorts characters of a string
func sortString(w string) string {
  s := strings.Split(w, "")
  sort.Strings(s)
  return strings.Join(s, "")
}

func IsPandigital(a, b int) bool {
  str := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(a*b)
  if len(str) != 9{
    return false
  }
  str = sortString(str)
  return "123456789" == str
}

func PandigitalProduct() int{
  set := mapset.NewSet()
  sum := 0
  for i := 0; i <= 10000; i++ {
    for j := 0; j <= 10000; j++ {
      if i * j >= 10000 {
        break
      }
      if !set.Contains(i*j) && IsPandigital(i,j) {
        set.Add(i*j)
        sum += i*j
      }
    }
  }
  return sum
}

