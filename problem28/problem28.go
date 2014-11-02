package problem28

func CaculateSpiralDiagonalSum(size int) int {
  sum := 1
  for i := 3; i<=size ; i +=2 {
    sum += 4 * i * i - 6 * (i - 1)
  }
  return sum
}
