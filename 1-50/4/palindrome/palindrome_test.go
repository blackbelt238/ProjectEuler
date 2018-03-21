package palindrome

import (
	"testing"
)

// TestIsPalindromic ensures the IsPalindromic function is properly identifying palindromic numbers
func TestIsPalindromic(t *testing.T) {
	val := 2
	if !IsPalindromic(val, 10) {
		t.Errorf("%d is palindromic, yet was determined not to be.\n", val)
	}

	val = 21
	if IsPalindromic(val, 10) {
		t.Errorf("%d is not palindromic, yet was determined to be.\n", val)
	}

	val = 212
	if !IsPalindromic(val, 10) {
		t.Errorf("%d is palindromic, yet was determined not to be.\n", val)
	}

	val = 2112
	if !IsPalindromic(val, 10) {
		t.Errorf("%d is palindromic, yet was determined not to be.\n", val)
	}

	val = 21121
	if IsPalindromic(val, 10) {
		t.Errorf("%d is not palindromic, yet was determined to be.\n", val)
	}
}

func TestIsPalindromicBinary(t *testing.T) {
	val := 1
	if !IsPalindromic(val, 2) {
		t.Errorf("%b is palindromic, yet was determined not to be.\n", val)
	}

	val = 11 // 1011
	if IsPalindromic(val, 2) {
		t.Errorf("%b is not palindromic, yet was determined to be.\n", val)
	}

	val = 101 // 1100101
	if IsPalindromic(val, 2) {
		t.Errorf("%b is not palindromic, yet was determined to be.\n", val)
	}
}

// TestProducts ensures that Products is properly determining all palindromic products for a given number of digits
func TestProducts(t *testing.T) {
	lgst := 9
	pals := Products(1)
	lgstrec := pals[len(pals)-1]
	if lgstrec != lgst {
		t.Errorf("largest palindrome in list should be %d, instead got %d\n", lgst, lgstrec)
	}

	lgst = 9009
	pals = Products(2)
	lgstrec = pals[len(pals)-1]
	if lgstrec != lgst {
		t.Errorf("largest palindrome in list should be %d, instead got %d\n", lgst, lgstrec)
	}
}
