package collatz

import (
	"testing"
)

func TestSequence(t *testing.T) {
	c := Sequence(13)
	exp := []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}

	if len(c) != len(exp) {
		t.Errorf("Incorrect sequence.\n\trec: %v\n\texp: %v", c, exp)
	}
	for i, num := range c {
		if num != exp[i] {
			t.Errorf("Incorrect sequence.\n\trec: %v\n\texp: %v", c, exp)
		}
	}
}

func TestSequenceLen(t *testing.T) {
	l := SequenceLen(13)
	exp := 10

	if l != exp {
		t.Errorf("Incorrect sequence length: got %d, should be %d", l, exp)
	}
}
