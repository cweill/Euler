package main

import (
	"fmt"
	"math"
)

func integerRightTriangles(maxp int) int {
	results := make(map[int]int)
	for i := 0; i < maxp; i++ {
		results[i] = 0
	}
	for p := 0; p < maxp; p++ {
		for x := 0; x < p; x++ {
			y := x
			z := p - x - y
			for y <= z {
				if math.Pow(float64(x), 2)+math.Pow(float64(y), 2) == math.Pow(float64(z), 2) {
					results[p]++
				}
				y++
				z--
			}
		}
	}
	bestp := 0
	bestCount := 0
	for i := 0; i < maxp; i++ {
		if results[i] >= bestCount {
			bestCount = results[i]
			bestp = i
		}
	}
	return bestp
}

func main() {
	fmt.Println(integerRightTriangles(1000))
}
