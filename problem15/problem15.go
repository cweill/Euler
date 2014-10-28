package main

import (
	"fmt"
)

func latticePathsHelper(n, x, y int, mem [][]int) int {
	if mem[x][y] != 0 {
		return mem[x][y]
	}
	if x == n && y == n {
		return 1
	}
	count := 0
	if y < n {
		count += latticePathsHelper(n, x, y+1, mem)
	}
	if x < n {
		count += latticePathsHelper(n, x+1, y, mem)
	}
	mem[x][y] = count
	return count
}

func latticePaths(n int) int {
	mem := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = make([]int, n+1)
	}
	return latticePathsHelper(n, 0, 0, mem)
}

func main() {
	fmt.Println(latticePaths(20))
}
