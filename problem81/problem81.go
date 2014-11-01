package problem81

func MinimalPathSumHelper(matrix, mem [][]int, x, y int) int {
  if x == len(matrix) - 1 && y == len(matrix[x]) - 1 {
    return matrix[x][y]
  }
  if mem[x][y] != 0 {
    return mem[x][y]
  }
  down, right, here := 0, 0, matrix[x][y]
  if x < len(matrix) - 1 {
    right = MinimalPathSumHelper(matrix, mem, x + 1, y)
  }
  if y < len(matrix[x]) - 1 {
    down = MinimalPathSumHelper(matrix, mem, x, y + 1)
  }
  sum := here
  if right != 0 && down > right || down == 0 {
    sum += right
  } else {
    sum += down
  }
  mem[x][y] = sum
  return sum
}

func MinimalPathSum(matrix[][]int) int {
  mem := make([][]int, len(matrix))
  for i, _ := range mem {
    mem[i] = make([]int, len(matrix[i]))
  }
  return MinimalPathSumHelper(matrix, mem, 0, 0)
}
