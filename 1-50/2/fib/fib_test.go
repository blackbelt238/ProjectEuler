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

// TestUpToMax ensures UpToMax is functioning as expected in a small case
func TestUpToMax(t *testing.T) {
	fcor := []int{0, 1, 1, 2, 3, 5} // the desired slice
	futm := UpToMax(5)              // slice recieved from UptoLimit

	if len(fcor) != len(futm) {
		t.Error("Lists are not the same length.\n")
		t.Errorf("\tcorrect entries: %v", fcor)
		t.Errorf("\trecieved entries: %v", futm)
	}

	for i := 0; i < len(futm); i++ {
		if fcor[i] != futm[i] {
			t.Error("Error retrieving entries whose sum does not exceed 12.\n")
			t.Errorf("\tcorrect entries: %v", fcor)
			t.Errorf("\trecieved entries: %v", futm)
		}
	}
}

// TestUpToMaxSmallestCase ensures UpToMax functions as expected in its' smallest case
func TestUpToMaxSmallestCase(t *testing.T) {
	fcor := []int{0, 1, 1} // the desired slice
	futm := UpToMax(1)     // slice recieved from UptoLimit

	if len(fcor) != len(futm) {
		t.Error("Lists are not the same length.\n")
		t.Errorf("\tcorrect entries: %v", fcor)
		t.Errorf("\trecieved entries: %v", futm)
	}

	for i := 0; i < len(futm); i++ {
		if fcor[i] != futm[i] {
			t.Error("Error retrieving entries whose sum does not exceed 12.\n")
			t.Errorf("\tcorrect entries: %v", fcor)
			t.Errorf("\trecieved entries: %v", futm)
		}
	}
}

// TestSumEven replicates the desired behavior of main to ensure the proper result is being calculated
func TestSumEven(t *testing.T) {
	fnums := UpToMax(5) // [0, 1, 1, 2, 3, 5]

	sumEven := 0 // sum of all even numbers in fnums
	for _, num := range fnums {
		if num%2 == 0 {
			sumEven += num
		}
	}

	if sumEven != 2 {
		t.Errorf("The sum of even numbers should be 2, got %d", sumEven)
	}

	fnums = UpToMax(144) // [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144]

	sumEven = 0 // sum of all even numbers in fnums
	for _, num := range fnums {
		if num%2 == 0 {
			sumEven += num
		}
	}

	if sumEven != 188 {
		t.Errorf("The sum of even numbers should be 188, got %d", sumEven)
	}
}
