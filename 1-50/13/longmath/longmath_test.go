package longmath

import (
	"testing"
)

// TestLongAdd runs LongAdd through a number of possible inputs to ensure valid calculations
func TestLongAdd(t *testing.T) {
	// adding 2 numbers of the same length whose sum is the same length
	n1, n2 := "12", "23"
	res := LongAdd(n1, n2)
	exp := "35"

	if res != exp {
		t.Errorf("incorrect sum of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	// adding 2 numbers of the same length whose sum is a different length
	n1, n2 = "2", "8"
	res = LongAdd(n1, n2)
	exp = "10"

	if res != exp {
		t.Errorf("incorrect sum of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	// adding 2 numbers of different lengths
	n1, n2 = "20", "8"
	res = LongAdd(n1, n2)
	exp = "28"

	if res != exp {
		t.Errorf("incorrect sum of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	n1, n2 = "8", "20"
	res = LongAdd(n1, n2)

	if res != exp {
		t.Errorf("incorrect sum of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	// adding 2 numbers of vastly different lengths
	n1, n2 = "2000", "8"
	res = LongAdd(n1, n2)
	exp = "2008"

	if res != exp {
		t.Errorf("incorrect sum of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	n1, n2 = "8", "2000"
	res = LongAdd(n1, n2)

	if res != exp {
		t.Errorf("incorrect sum of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}
}

func TestLongMult(t *testing.T) {
	n1, n2 := "2", "3"
	res := LongMult(n1, n2)
	exp := "6"

	if res != exp {
		t.Errorf("incorrect product of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	n1, n2 = "10", "20"
	res = LongMult(n1, n2)
	exp = "200"

	if res != exp {
		t.Errorf("incorrect product of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	n1, n2 = "612", "24"
	res = LongMult(n1, n2)
	exp = "14688"

	if res != exp {
		t.Errorf("incorrect product of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	n1, n2 = "423", "211"
	res = LongMult(n1, n2)
	exp = "89253"

	if res != exp {
		t.Errorf("incorrect product of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}

	n1, n2 = "674", "220002"
	res = LongMult(n1, n2)
	exp = "148281348"

	if res != exp {
		t.Errorf("incorrect product of %s and %s calculated: got %s, expected %s", n1, n2, res, exp)
	}
}
