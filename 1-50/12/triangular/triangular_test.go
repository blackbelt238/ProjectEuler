package triangular

import (
	"testing"
)

func TestFind(t *testing.T) {
	tri := CreateTriangular()
	tri.Find(7)

	if tri.Val() != 28 {
		t.Errorf("finding the 7th triangular number yielded %d, should be 28", tri.Val())
	}

	tri.Find(tri.N() - 4)
	if tri.Val() != 6 {
		t.Errorf("finding the 3rd triangular number yielded %d, should be 6", tri.Val())
	}
}

func TestNumFactors(t *testing.T) {
	tri := CreateTriangular()
	tri.Find(7)

	nf := tri.NumFactors()
	if nf != 6 {
		t.Errorf("triangle %d should have 6 divisors, got %d", tri.Val(), nf)
	}
}
