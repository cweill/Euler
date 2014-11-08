package problem55

import (
	"math/big"
	"strconv"
)

func Reverse(str string) string {
	runes := []rune(str)
	i, j := 0, len(runes)-1
	for i < j {
		temp := runes[i]
		runes[i] = runes[j]
		runes[j] = temp
		i++
		j--
	}
	return string(runes)
}

func IsPalindrome(str string) bool {
	runes := []rune(str)
	i, j := 0, len(runes)-1
	for i < j {
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func ReverseAndAdd(n *big.Int) *big.Int {
	str := Reverse(n.String())
	num, _ := strconv.Atoi(str)
	return n.Add(big.NewInt(int64(num)), n)
}

func IsLychrel(n int) bool {
	count := 0
	num := big.NewInt(int64(n))
	for count < 50 {
		newNum := ReverseAndAdd(num)
		if IsPalindrome(newNum.String()) {
			return false
		}
		count++
	}
	return true
}

func LychrelNumbers() int {
	count := 0
	for i := 0; i < 10000; i++ {
		if IsLychrel(i) {
			count++
		}
	}
	return count
}
