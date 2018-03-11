package primefactor

import (
	"math"
)

// SieveOfEratosthenes finds all prime numbers up to the given limit n
func SieveOfEratosthenes(n int) []int {
	// create the sieve, initialize all values to true
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	var primes []int

	// for each prime in the sieve
	for num := 2; num <= int(math.Floor(math.Sqrt(float64(n)))); num++ {
		if sieve[num] == true {
			// mark off its' multiple
			for mult := int(math.Floor(math.Pow(float64(num), 2))); mult <= n; mult += num {
				sieve[mult] = false
			}
		}
	}

	for num, isPrime := range sieve {
		if isPrime {
			primes = append(primes, num)
		}
	}

	return primes[2:]
}
