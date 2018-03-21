package pythag

import (
	"testing"
)

// TestIsTriplet determines if IsTriplet properly detects Pythagorean triplets
func TestIsTriplet(t *testing.T) {
	if !IsTriplet(3, 4, 5) {
		t.Errorf("The given triplet is a Pythagorean triplet, but was determined to not be.")
	}

	if IsTriplet(3, 4, 6) {
		t.Errorf("The given triplet is not a Pythagorean triplet, but was determined to be.")
	}
}
