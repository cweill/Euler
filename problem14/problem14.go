package main

import (
	"fmt"
)

func collatz(n int) int {
	if n == 1 {
		return 1
	}
	next := 0
	if n%2 == 0 {
		next = n / 2
	} else {
		next = 3*n + 1
	}
	return 1 + collatz(next)
}

func longestCollatz(maxn int) int {
	longest := 0
	best := 0
	for n := 1; n <= maxn; n++ {
		coll := collatz(n)
		if coll > longest {
			longest = coll
			best = n
		}
	}
	return best
}

func main() {
	fmt.Println(longestCollatz(1000000))
}
