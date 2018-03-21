package palindrome

import (
	"math"
	"sort"
	"strconv"
)

// IsPalindromic determines if the given decimal value is a palindromic number
func IsPalindromic(val int, base int) bool {
	vstr := strconv.FormatInt(int64(val), base) // string representation of val
	for i, j := 0, len(vstr)-1; i <= j; i, j = i+1, j-1 {
		if vstr[i] != vstr[j] {
			return false
		}
	}
	return true
}

// Products returns all palindromic numbers that can be made
//   by multiplying numbers of a given number of digits
func Products(ndigs int) []int {
	var pals []int // palindromic numbers
	var prod int   // product of 2 numbers

	// calculate all applicable products and determine if they are palindromic
	for i := int(math.Pow(float64(10), float64(ndigs-1))); i < int(math.Pow(float64(10), float64(ndigs))); i++ {
		for j := int(math.Pow(float64(10), float64(ndigs-1))); j < int(math.Pow(float64(10), float64(ndigs))); j++ {
			prod = i * j
			if IsPalindromic(prod, 10) {
				pals = append(pals, prod)
			}
		}
	}

	// sort palindromes in ascending order
	sort.Slice(pals, func(i, j int) bool {
		return pals[i] < pals[j]
	})
	return pals
}
