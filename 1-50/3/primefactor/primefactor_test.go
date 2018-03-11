package primefactor

import (
	"testing"
)

// TestLargestPrime ensures that the largest number returned by SieveOfEratosthenes
//   is in fact the largest prime number under the given number
func TestLargestPrime(t *testing.T) {
	lgst := 29
	prec := SieveOfEratosthenes(30)
	lgstrec := prec[len(prec)-1]

	if lgst != lgstrec {
		t.Errorf("Largest recieved prime does not match the expected number:\n")
		t.Errorf("\texpected: %v\n", lgst)
		t.Errorf("\trecieved: %v\n", lgstrec)
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
}
