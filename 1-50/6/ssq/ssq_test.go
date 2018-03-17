package ssq

import (
	"testing"
)

// TestSQS ensures that SQS is properly calculating the square-of-sum
func TestSQS(t *testing.T) {
	exp := 3025
	rec := SQS(10)

	if exp != rec {
		t.Errorf("the recieved number (%d) doesn't match the expected number (%d)\n", rec, exp)
	}
}

// TestSSQ ensures that SSQ is properly calculating the sum-of-squares
func TestSSQ(t *testing.T) {
	exp := 385
	rec := SSQ(10)

	if exp != rec {
		t.Errorf("the recieved number (%d) doesn't match the expected number (%d)\n", rec, exp)
	}
}

// TestSumSquareDifference ensures SumSquareDifference properly calculates SQS - SSQ
func TestSumSquareDifference(t *testing.T) {
	exp := 2640
	rec := SumSquareDifference(10)

	if exp != rec {
		t.Errorf("the recieved number (%d) doesn't match the expected number (%d)\n", rec, exp)
	}
}
