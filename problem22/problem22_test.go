package problem22_test

import (
	. "euler/problem22"
  "io/ioutil"
  "strings"
  "sort"
  "regexp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problem22", func() {
  Describe("NameScore", func () {
    It("is 49714 for COLIN (938)", func () {
      Expect(NameScore("COLIN", 938)).To(Equal(49714))
    })
  })

  It("answers", func(){
    content, _ := ioutil.ReadFile("p022_names.txt")
    re := regexp.MustCompile("(\"|\\s)+")
    names := strings.Split(re.ReplaceAllString(string(content), ""), ",")
    sort.Strings(names)
    Expect(TotalNameScores(names)).To(Equal(871198282))
  })
})
