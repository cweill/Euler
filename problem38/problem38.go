package problem38

import (
  "strconv"
  "strings"
  "sort"
)

// Sorts characters of a string
func sortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func isPandigital(str string) bool {
  if len(str) != 9{
    return false
  }
  str = sortString(str)
  return "123456789" == str
}

func ConcatenateProducts(n int) int {
  i := 1
  result := ""
  for {
    result+= strconv.Itoa(n*i)
    if isPandigital(result) || len(result) > 9 {
      break
    }
    i++
  }
  res, _ := strconv.Atoi(result)
  return res
}

func PandigitalMultiples() int{
  best := 0
  for i:=0;i<99999; i++ {
    current := ConcatenateProducts(i)
    if current > best && isPandigital(strconv.Itoa(current)){
      best = current
    }
  }
  return best
}
