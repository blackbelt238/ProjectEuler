package prime

import (
	"math"
)

// Factorize finds the prime factorization of a number
func Factorize(num int) map[int]int {
	pfac := make(map[int]int)
	primes := SieveOfEratosthenes(int(math.Sqrt(float64(num))))

	for !IsPrime(num) {
		for _, prime := range primes {
			if num%prime == 0 {
				pfac[prime]++
				num = num / prime
				break
			}
		}
	}
	pfac[num]++

	return pfac
}

// Factors returns all the prime factors for the given number
func Factors(num int) []int {
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

// GetPrime finds the n-th prime
func GetPrime(n int) int {
	var prime, num int
	for i := 2; num < n; i++ {
		if IsPrime(i) {
			prime = i
			num++
		}
	}
	return prime
}

// IsPrime determines if the given number is prime
func IsPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// SieveOfEratosthenes finds all prime numbers up to the given limit
func SieveOfEratosthenes(limit int) []int {
	// create the sieve, initialize all values to true
	sieve := make([]bool, limit+1)
	var primes []int

	// for each prime in the sieve
	for num := 2; num <= int(math.Sqrt(float64(limit))); num++ {
		if sieve[num] == false {
			// mark off its' multiple as composite
			for mult := num * num; mult <= limit; mult += num {
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
