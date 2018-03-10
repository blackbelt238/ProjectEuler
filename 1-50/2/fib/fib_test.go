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
