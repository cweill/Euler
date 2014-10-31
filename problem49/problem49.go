package problem49

import (
  "euler/utils"
  "set"
  "permute"
  "strconv"
)

func PrimePermutations() string {
  primes := utils.Sieve(10000)
  primesSet := mapset.NewSet()
  for _, prime := range primes {
    primesSet.Add(prime)
  }
  for _, prime := range primes {
    if prime >= 1000 && prime != 1487 {
      secondPrime := prime + 3330
      if primesSet.Contains(secondPrime) {
        thirdPrime := secondPrime + 3330
        if primesSet.Contains(thirdPrime) {
          permutations := permute.LexicographicPermutations(strconv.Itoa(prime))
          permutationsSet := mapset.NewSet()
          for _, permutation := range permutations {
            permutationsSet.Add(permutation)
          }
          if permutationsSet.Contains(strconv.Itoa(secondPrime)) && permutationsSet.Contains(strconv.Itoa(thirdPrime)) {
            return strconv.Itoa(prime) + strconv.Itoa(secondPrime) + strconv.Itoa(thirdPrime)
          }
        }
      }
    }
  }
  return "Nope"
}
