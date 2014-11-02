package utils

import (
	"math"
	"set"
	"sort"
	"strings"
	"strconv"
)

func Sieve(n int) []int {
	isPrimeArray := make([]bool, n)
	for i := 0; i < len(isPrimeArray); i++ {
		isPrimeArray[i] = true
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if isPrimeArray[i] {
			j := i * i
			counter := 1
			for j < n {
				isPrimeArray[j] = false
				j = i*i + i*counter
				counter++
			}
		}
	}
	var primes []int
	for i := 2; i < len(isPrimeArray); i++ {
		if isPrimeArray[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func PrimeSet(maxn int) mapset.Set {
  set := mapset.NewSet()
  for _, prime := range Sieve(maxn) {
    set.Add(prime)
  }
  return set
}

// Sorts characters of a string
func SortString(w string) string {
  s := strings.Split(w, "")
  sort.Strings(s)
  return strings.Join(s, "")
}

func IsPandigital(num, n int) bool {
	str := strconv.Itoa(num)
  if len(str) != n{
    return false
  }
  str = SortString(str)
  comp := ""
  for i:= 1; i<=n ; i++ {
  	comp += strconv.Itoa(i)
  }
  return comp == str
}
