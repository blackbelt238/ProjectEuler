package palindrome

import (
	"math"
	"strconv"
)

// IsPalindromic determines if the given value is a palindromic number
func IsPalindromic(val int) bool {
	vstr := strconv.FormatInt(int64(val), 10) // string representation of val
	digs := make([]int, len(vstr))            // val's digits
	var digsr []int                           // val's digits in reverse order

	// populate the list of val's digits and its' reverse
	for i, dig := range vstr {
		digs[i] = int(dig)
	}
	digsr = reverse(digs)

	for i, j := 0, len(digs)-1; i <= j; i, j = i+1, j-1 {
		if digs[i] != digsr[i] || digs[j] != digsr[j] {
			return false
		}
	}

	return true
}

// reverse returns a reflected version of s
func reverse(s []int) []int {
	sr := make([]int, len(s))

	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		sr[i] = s[j]
		sr[j] = s[i]
	}
	return sr
}
