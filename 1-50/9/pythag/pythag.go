package pythag

import (
	"math"
)

// IsTriplet determines if the given number is a Pythagorean triplet
func IsTriplet(a, b, c int) bool {
	if int(math.Pow(float64(a), 2))+int(math.Pow(float64(b), 2)) == int(math.Pow(float64(c), 2)) {
		return true
	}
	return false
}
