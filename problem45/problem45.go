package problem45

import (
  "euler/problem44"
  "euler/problem42"
)

func TriangularPentagonalAndHexagonal(minp int) int{
  max := 100000
  triangles := problem42.TriangleNumbers(max)
  pentagons := problem44.PentagonNumbers(max)
  for p := minp; p < len(pentagons); p++{
    for t := p; t < len(triangles); t++{
      if triangles[t] > pentagons[p] {
        break
      } else if triangles[t] == pentagons[p] && t % 2 == 0 {
        return triangles[t]
      }
    }
  }
  return -1
}

