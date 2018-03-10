package threefive

import (
	"testing"
)

func TestThreeAndFiveBelow10(t *testing.T) {
	val := ThreeAndFive(10)
	if val != 23 {
		t.Errorf("threeAndFive below 10 should be 23, instead found %d", val)
	}
}
