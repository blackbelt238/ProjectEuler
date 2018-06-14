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
