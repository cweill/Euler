package main

import (
	"fmt"
	"strconv"
)

func isPandigital(n int) bool {
	if n < 123456789 || n > 9876543210 {
		return false
	}
	nString := strconv.Itoa(n)
	characterCheck := make([]bool, 10)
	if len(nString) < 10 {
		nString = "0" + nString
	}
	for i := 0; i < 10; i++ {
		charNum, _ := strconv.Atoi(nString[i : i+1])
		if characterCheck[charNum] {
			return false
		}
		characterCheck[charNum] = true
	}
	fmt.Println(nString)
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	start := 1
	for start+3 <= 10 {
		sub, _ := strconv.Atoi(nString[start : start+3])
		if sub%primes[start-1] != 0 {
			return false
		}
		start++
	}
	return true
}

func main() {
	count := 0
	for i := 2000000000; i <= 10000000000; i++ {
		if isPandigital(i) {
			count++
		}
	}
	fmt.Println(count)
}
