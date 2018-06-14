package triangular

import (
	"testing"
)

func TestFind(t *testing.T) {
	tri := CreateTriangular()
	f := tri.Find(7)
	if f != 28 {
		t.Errorf("finding the 7th triangular number yielded %d, should be 28", f)
	}

	f = tri.Find(3)
	if f != 6 {
		t.Errorf("finding the 3rd triangular number yielded %d, should be 6", f)
	}
}

func TestNumFactors(t *testing.T) {
	tri := CreateTriangular()
	tri.Find(7)

	nf := tri.NumFactors()
	if nf != 6 {
		t.Errorf("triangle %d should have 6 divisors, got %d", tri.Val, nf)
	}
}
