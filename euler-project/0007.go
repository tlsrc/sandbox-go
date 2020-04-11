package eulerproject

import (
	"math"
)

//Problem0007 solves problem 7
func Problem0007() int {
	primes := []int{2, 3, 5, 7, 11, 13}
	i := 14
	for len(primes) < 10_001 {
		if isPrime(primes, i) {
			primes = append(primes, i)
		}
		i++
	}
	return primes[len(primes)-1]
}

func isPrime(knownPrimes []int, candidate int) bool {
	for _, p := range knownPrimes {
		if p > int(math.Sqrt(float64(candidate)))+1 {
			return true
		}
		if candidate%p == 0 {
			return false
		}
	}
	return true
}
