package problem62

import (
	"euler/utils"
	"strconv"
)

func CubicPermutations(n int) int {
	i := 0
	cubes := make(map[string][]int)
	for {
		i++
		key := utils.SortString(strconv.Itoa(i * i * i))
		if cubes[key] == nil {
			cubes[key] = []int{i * i * i}
		} else {
			cubes[key] = append(cubes[key], i*i*i)
		}
		if len(cubes[key]) == n {
			return cubes[key][0]
		}
	}
	return -1
}
