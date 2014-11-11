package utils

import (
	"io/ioutil"
	"math"
	"regexp"
	"set"
	"sort"
	"strconv"
	"strings"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

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

func IsPandigital(num, n uint) bool {
	min := []uint{0, 1, 12, 123, 1234, 12345, 123456, 1234567, 12345678, 123456789}
	if num >= min[n] {
		digits := 0
		for ; num > 0; num /= 10 {
			k := num % 10
			// already seen digit k
			if (digits&(1<<k))>>k == 1 {
				return false
			}
			// Mark kth bit as true
			digits |= 1 << k
		}
		return digits == (1<<(n+1))-2
	}
	return false
}
func ReadStringsFromFile(filename string) []string {
	content, _ := ioutil.ReadFile(filename)
	re := regexp.MustCompile("(\"|\\s)+")
	return strings.Split(re.ReplaceAllString(string(content), ""), ",")
}

func WordValue(word string) int {
	sum := 0
	for i := 0; i < len(word); i++ {
		sum += int(word[i] - 64)
	}
	return sum
}

func HasSameDigits(this, that int) bool {
	sthis, sthat := strconv.Itoa(this), strconv.Itoa(that)
	sthis = SortString(sthis)
	sthat = SortString(sthat)
	return sthis == sthat
}
