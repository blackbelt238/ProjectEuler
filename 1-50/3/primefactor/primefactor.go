package primefactor

import (
	"math"
)

// PrimeFactors returns all the prime factors for the given number
func PrimeFactors(num int) []int {
	primes := SieveOfEratosthenes(int(math.Sqrt(float64(num)))) // all prime numbers up to sqrt(num)
	var pfacs []int                                             // prime factors for num

	for _, prime := range primes {
		// if the prime is a prime factor
		if num%prime == 0 {
			pfacs = append(pfacs, prime)
		}
	}

	return pfacs
}

// SieveOfEratosthenes finds all prime numbers up to the given limit n
func SieveOfEratosthenes(n int) []int {
	// create the sieve, initialize all values to true
	sieve := make([]bool, n+1)
	var primes []int

	// for each prime in the sieve
	for num := 2; num <= int(math.Sqrt(float64(n))); num++ {
		if sieve[num] == false {
			// mark off its' multiple as composite
			for mult := num * num; mult <= n; mult += num {
				sieve[mult] = true
			}
		}
	}

	// for each number that is not composite, add it to the list of primes
	for num, isComposite := range sieve {
		if !isComposite {
			primes = append(primes, num)
		}
	}

	return primes[2:]
}
