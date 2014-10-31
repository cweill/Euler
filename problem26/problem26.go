package problem26

import (
	"euler/utils"
	"math/big"
	"suffixarrayx"
)

func ReciprocalCycles() int {
	longest := 0
	best := -1
	primes := utils.Sieve(1000)
	for i := 0; i < len(primes); i++ {
		rat := big.NewRat(1, int64(primes[i]))
		sa := suffixarrayx.NewSuffixArrayX(rat.FloatString(2000))
		lnors := sa.LongestRepeatingNonOverlappingSubstring()
		if len(lnors) > longest {
			longest = len(lnors)
			best = primes[i]
		}
	}
	return best
}
