package problem34

func FactorialDigit(n int) int {
  facts := []int {
    1,
    1,
    2,
    6,
    24,
    120,
    720,
    5040,
    40320,
    362880,
  }
  return facts[n]
}

func IsSumFactorialOfDigits(n int) bool {
  base := 1
  sum := 0
  for base <= n {
    sum += FactorialDigit(n / base % 10)
    base *= 10
  }
  return sum == n
}

func DigitFactorials() int {
  total := 0
  for i := 3 ; i < 100000; i++ {
    if IsSumFactorialOfDigits(i) {
      total += i
    }
  }
  return total
}
