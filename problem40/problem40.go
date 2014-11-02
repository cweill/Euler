package problem40

import (
  "strconv"
  "bytes"
)

func ChampernownesConstant(maxn int) string {
  var buffer bytes.Buffer
  buffer.WriteString("0.")
  for i := 1; i <= maxn ; i++ {
    buffer.WriteString(strconv.Itoa(i))
  }
  return buffer.String()
}

func MultipleOfChampernownesConstantDigits() int {
  c := ChampernownesConstant(1000000)
  product, base := 1, 1
  for base <= 1000000 {
    i, _ := strconv.Atoi(c[base+1:base+2])
    product *= i
    base *= 10
  }
  return product
}
