package fib

import (
	"testing"
)

// TestFib ensures Fib is properly calculating the Fibonacci numbers
func TestFib(t *testing.T) {
	num := Fib(12)
	if num != 144 {
		t.Errorf("The 12th Fibonacci number should be 144, instead got %d", num)
	}
}

// TestUpToLimit ensures UpToLimit is returning the correct slice of Fibonacci numbers
func TestUpToLimit(t *testing.T) {
	fcor := []int{0, 1, 1, 2, 3, 5} // the desired slice
	futl := UpToLimit(12)           // slice recieved from UptoLimit

	for i := 0; i < len(futl); i++ {
		if fcor[i] != futl[i] {
			t.Error("Error retrieving entries whose sum does not exceed 12.\n")
			t.Errorf("\tcorrect entries: %v", fcor)
			t.Errorf("\tcorrect entries: %v", futl)
		}
	}

	// test the smallest case
	fcor = []int{0}
	futl = UpToLimit(0)

	for i := 0; i < len(futl); i++ {
		if fcor[i] != futl[i] {
			t.Error("Error retrieving entries whose sum does not exceed 12.\n")
			t.Errorf("\tcorrect entries: %v", fcor)
			t.Errorf("\tcorrect entries: %v", futl)
		}
	}
}
