package palindrome

import (
	"testing"
)

// TestIsPalindromic ensures the IsPalindromic function is properly identifying palindromic numbers
func TestIsPalindromic(t *testing.T) {
	val := 2
	if !IsPalindromic(val) {
		t.Errorf("%d is palindromic, yet was determined not to be.\n", val)
	}

	val = 21
	if IsPalindromic(val) {
		t.Errorf("%d is not palindromic, yet was determined to be.\n", val)
	}

	val = 212
	if !IsPalindromic(val) {
		t.Errorf("%d is palindromic, yet was determined not to be.\n", val)
	}

	val = 2112
	if !IsPalindromic(val) {
		t.Errorf("%d is palindromic, yet was determined not to be.\n", val)
	}
}

// TestProducts ensures that Products is properly determining all palindromic products for a given number of digits
func TestProducts(t *testing.T) {
	lgst := 9009
	pals := Products(2)
	if pals == nil {
		t.Errorf("palindromics cannot be nil.\n")
	}

	lgstrec := pals[len(pals)-1]
	if lgstrec != lgst {
		t.Errorf("largest palindrome in list should be %d, instead got %d\n", lgst, lgstrec)
	}
}

// TestReverse tests if reverse is properly reflecting the given int slice
func TestReverse(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}     // list to reverse
	lrexp := []int{5, 4, 3, 2, 1} // expected reverse
	lrrec := reverse(l)

	if len(lrexp) != len(lrrec) {
		t.Errorf("Recieved reverse not the expected length\n")
	}
}
