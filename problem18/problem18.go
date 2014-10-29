package problem18

func MaximumPathSumHelper(pyramid [][]int, column, level int, mem [][]int) int {
	if level == len(pyramid)-1 {
		return pyramid[level][column]
	}
	if mem[level][column] != 0 {
		return mem[level][column]
	}
	down := MaximumPathSumHelper(pyramid, column, level+1, mem)
	right := MaximumPathSumHelper(pyramid, column+1, level+1, mem)
	var max int
	if down > right {
		max = down
	} else {
		max = right
	}
	sum := pyramid[level][column] + max
	mem[level][column] = sum
	return sum
}

func MaximumPathSum(pyramid [][]int) int {
	mem := make([][]int, len(pyramid))
	for i := 0; i < len(pyramid); i++ {
		mem[i] = make([]int, len(pyramid[i]))
	}
	return MaximumPathSumHelper(pyramid, 0, 0, mem)
}
