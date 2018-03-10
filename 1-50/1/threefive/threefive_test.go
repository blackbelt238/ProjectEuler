package threefive

import (
	"testing"
)

// TestThreeAndFiveBelow10 tests the ThreeAndFive function using Euler's sample input
func TestThreeAndFiveBelow10(t *testing.T) {
	val := ThreeAndFive(10)
	if val != 23 {
		t.Errorf("threeAndFive below 10 should be 23, instead found %d", val)
	}
}
