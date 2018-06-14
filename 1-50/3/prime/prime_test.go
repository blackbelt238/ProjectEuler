package prime

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

func TestFactorize(t *testing.T) {
	pfacs := Factorize(147)
	exp := map[int]int{3: 1, 7: 2}

	if len(pfacs) != 2 {
		t.Errorf("calc: %v\n exp: %v", pfacs, exp)
	}
	for k, v := range pfacs {
		if exp[k] != v {
			t.Errorf("calc: %v\n exp: %v", pfacs, exp)
		}
	}
}

// TestGetPrime ensures that GetPrime can find the n-th prime
func TestGetPrime(t *testing.T) {
	six := GetPrime(6)
	if six != 13 {
		t.Errorf("The sixth prime is 13, but instead got %d", six)
	}

	twenty5 := GetPrime(25)
	if twenty5 != 97 {
		t.Errorf("The 25th prime is 97, but instead got %d", twenty5)
	}

	cincociento := GetPrime(500)
	if cincociento != 3571 {
		t.Errorf("The 500th prime is 3581, but instead got %d", cincociento)
	}
}

func TestIsPrime(t *testing.T) {
	var num int

	num = 10
	if IsPrime(num) {
		t.Errorf("%d was found to be prime, but isnt", num)
	}

	num = 13 // 6th prime
	if !IsPrime(num) {
		t.Errorf("%d was found to be not prime, but is", num)
	}

	num = 97 // 25th prime
	if !IsPrime(num) {
		t.Errorf("%d was found to be not prime, but is", num)
	}

	num = 3580
	if IsPrime(num) {
		t.Errorf("%d was found to be prime, but isn't", num)
	}

	num = 3581 // 501st prime
	if !IsPrime(num) {
		t.Errorf("%d was found to be not prime, but is", num)
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
