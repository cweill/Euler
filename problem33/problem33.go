package problem33

func gcd(a, b int) int {
  for b != 0 {
    temp := b
    b = a % b
    a = temp
  }
  return a
}

func DigitCancellingFractions () float64{
  var i, a, b, numprod, denoprod float64
  numprod, denoprod = 1, 1
  for i = 1; i <=9; i++ {
    for a = 1; a <= 9; a++ {
      for b = 1; b <= 9; b++ {
        numerator := a * 10 + i
        denominator := i * 10 + b
        if numerator / denominator != 1 && numerator / denominator == a / b {
          numprod *= numerator
          denoprod *= denominator
        }
      }
    }
  }
  return denoprod / float64(gcd(int(denoprod), int(numprod)))
}
