package problem21

func ProperDivisors(n int) []int {
	var divisors []int
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func SumProperDivisors(n int) int {
	nDivisors := ProperDivisors(n)
	sum := 0
	for i := 0; i < len(nDivisors); i++ {
		sum += nDivisors[i]
	}
	return sum
}

func IsAmicable(a int) bool {
	b := SumProperDivisors(a)
	if b == a {
		return false
	}
	potentiala := SumProperDivisors(b)
	return potentiala == a
}

func SumAmicables(maxn int) int {
	sum := 0
	for i := 1; i < 10000; i++ {
		if IsAmicable(i) {
			sum += i
		}
	}
	return sum
}
