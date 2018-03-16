package smallmult

import (
	"testing"
)

// TestSmallestMultiple ensures SmallestMultiple is properly functioning
func TestSmallestMultiple(t *testing.T) {
	exp := 2520
	rec := SmallestMultiple(10)

	if exp != rec {
		t.Errorf("Recieved unexpedted number. Got %d instead of %d\n", rec, exp)
	}
}
