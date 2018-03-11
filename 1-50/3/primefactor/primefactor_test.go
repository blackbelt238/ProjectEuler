package primefactor

import (
	"testing"
)

// TestLargestPrime ensures that the largest number returned by SieveOfEratosthenes
//   is in fact the largest prime number under the given number
func TestLargestPrime(t *testing.T) {
	lgst := 29 // expected largest prime
	prec := SieveOfEratosthenes(30)
	lgstrec := prec[len(prec)-1] // largest prime recieved

	if lgst != lgstrec {
		t.Errorf("Largest recieved prime does not match the expected number:\n")
		t.Errorf("\texpected: %v\n", lgst)
		t.Errorf("\trecieved: %v\n", lgstrec)
	}
}

// TestPrimeFactors ensures that PrimeFactors is returning all expected prime factors
func TestPrimeFactors(t *testing.T) {
	facexp := []int{5, 7, 13, 29} // expected prime factors
	facrec := PrimeFactors(13195) // recieved prime factors

	// immediately fail if the lists aren't the same length
	if len(facexp) != len(facrec) {
		t.Errorf("Length of recieved list does not match that which is expected:\n")
		t.Errorf("\texpected: %v\n", facexp)
		t.Errorf("\trecieved: %v\n", facrec)
	}

	// check equality of all prime factors
	for i := range facexp {
		if facexp[i] != facrec[i] {
			t.Errorf("Should have recieved %d, instead got %d\n", facexp[i], facrec[i])
		}
	}
}

// TestSieveOfEratosthenes ensures SieveOfEratosthenes is returning the expected list of primes
func TestSieveOfEratosthenes(t *testing.T) {
	pexp := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29} // the expected list of primes
	prec := SieveOfEratosthenes(30)                   // the recieved list of primes

	if len(pexp) != len(prec) {
		t.Errorf("Length of recieved list does not match that which is expected:\n")
		t.Errorf("\texpected: %v\n", pexp)
		t.Errorf("\trecieved: %v\n", prec)
	}

	for i := range pexp {
		if pexp[i] != prec[i] {
			t.Errorf("Should have recieved %d, instead got %d\n", pexp[i], prec[i])
		}
	}
}
