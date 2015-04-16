package utils

import (
	"io/ioutil"
	"math"
	"math/big"
	"regexp"
	"set"
	"sort"
	"strconv"
	"strings"
)

func GCD(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
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

// AKS Algorithm
// http://en.wikipedia.org/wiki/AKS_primality_test#Algorithm

func IsPowerOfInteger(n int64) bool {
	for b := 2.0; b < math.Log2(float64(n)); b++ {
		a := math.Pow(float64(n), 1.0/b)
		if a-math.Trunc(a) < 0.00000001 {
			return true
		}
	}
	return false
}

func SmallestMultiplicativeOrder(n int64) int64 {
	maxk := int64(math.Floor(math.Pow(math.Log2(float64(n)), 2)))
	maxr := int64(math.Max(3, math.Ceil(math.Pow(math.Log2(float64(n)), 5))))
	nextR := true
	var r, k int64
	for r = 2.0; nextR && r < maxr; r++ {
		nextR = false
		for k = 1.0; !nextR && k <= maxk; k++ {
			remainder := big.NewInt(0).Exp(big.NewInt(n), big.NewInt(int64(k)), big.NewInt(int64(r)))
			nextR = (remainder.Int64() == 1 || remainder.Int64() == 0)
		}
	}
	return r - 1
}

// Agrawal, Kayal and Saxena
// Input: Integer n > 1
//
// if (n is has the form ab with b > 1) then output COMPOSITE
//
// r := 2
// while (r < n) {
//     if (gcd(n,r) is not 1) then output COMPOSITE
//     if (r is prime greater than 2) then {
//         let q be the largest factor of r-1
//         if (q > 4sqrt(r)log n) and (n(r-1)/q is not 1 (mod r)) then break
//     }
//     r := r+1
// }
//
// for a = 1 to 2sqrt(r)log n {
//     if ( (x-a)n is not (xn-a) (mod xr-1,n) ) then output COMPOSITE
// }
//
// output PRIME;

func IsPrime(n int64) bool {
	// (1) If n = a^b for integers a > 1 and b > 1, output composite.
	if IsPowerOfInteger(n) {
		return false
	}
	// (2) Find the smallest r such that O_r(n) > (log n)^2.
	r := SmallestMultiplicativeOrder(n)

	// (3) If 1 < gcd(a,n) < n for some a ≤ r, output composite.
	for a := r; a > 1; a-- {
		gcd := GCD(a, n)
		if gcd > 1 && gcd < n {
			return false
		}
	}

	// (4) If n ≤ r, output prime.
	if n <= r {
		return true
	}

	// (5)

	// (6) Output prime.
	return true
}
