package problem81_test

import (
	. "euler/problem81"
  "io/ioutil"
  "strings"
  "strconv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem81", func() {
  Describe("MinimalPathSum", func() {
    It("works", func(){
      matrix := [][]int{
        []int { 331,},
      }
      Expect(MinimalPathSum(matrix)).To(Equal(331))
    })

    It("works", func(){
      matrix := [][]int{
        []int {  37,},
        []int {  331,},
      }
      Expect(MinimalPathSum(matrix)).To(Equal(368))
    })

    It("works", func(){
      matrix := [][]int{
        []int { 956, 331,},
      }
      Expect(MinimalPathSum(matrix)).To(Equal(1287))
    })

    It("works", func(){
      matrix := [][]int{
        []int { 131 ,201 ,630 ,537 ,805 ,},
        []int { 673, 96 ,803 ,699 ,732 ,},
        []int { 234 ,342 ,746 ,497 ,524 ,},
        []int { 103, 965, 422, 121, 37 ,},
        []int { 18, 150, 111, 956, 331,},
      }
      Expect(MinimalPathSum(matrix)).To(Equal(2427))
    })
  })

  It("answers", func(){
    content, err := ioutil.ReadFile("p081_matrix.txt")
    if err != nil {
        //Do something
    }
    lines := strings.Split(string(content), "\n")
    matrix := make([][]int, len(lines) -1)
    for i:=0;i<len(lines) -1; i++ {
      ints := strings.Split(lines[i], ",")
      matrix[i] = make([]int, len(ints))
      for j, str := range ints {
        matrix[i][j], _ = strconv.Atoi(str)
      }
    }
    Expect(MinimalPathSum(matrix)).To(Equal(427337))
  })
})
