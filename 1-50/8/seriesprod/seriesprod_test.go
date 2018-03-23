package seriesprod

import (
	"testing"
)

func TestLargestSeriesProdSmall(t *testing.T) {
	prod := LargestSeriesProd("./twenty.txt", 4)
	exp := 630
	if prod != exp {
		t.Errorf("Incorrect product.\n\texpected: %d\n\trecieved: %d", exp, prod)
	}

	prod = LargestSeriesProd("./four.txt", 4)
	exp = 5832
	if prod != exp {
		t.Errorf("Incorrect product.\n\texpected: %d\n\trecieved: %d", exp, prod)
	}
}

// func TestLargestSeriesProd(t *testing.T) {
// 	prod := LargestSeriesProd("./../thousand.txt", 4)
// 	if prod != 5832 {
// 		t.Errorf("Incorrect product.\n\texpected: %d\n\trecieved: %d", 5832, prod)
// 	}
// }

func TestAdjCycle(t *testing.T) {
	adjque := []string{"4", "3", "2", "1"}
	exp := []string{"3", "2", "1", "0"}
	adjque = adjCycle(adjque, "0")

	if len(adjque) != 4 {
		t.Errorf("New adjque has improper length. want: 4, got: %d", len(adjque))
	}

	for i := range exp {
		if adjque[i] != exp[i] {
			t.Errorf("The %dth value does not match the expected value.\n\texpected: %s\n\trecieved: %s", i, exp[i], adjque[i])
		}
	}
}

func TestAdjProd(t *testing.T) {
	adjque := []string{"4", "3", "2", "1"}
	prod := 4 * 3 * 2 * 1

	if adjProd(adjque) != prod {
		t.Errorf("Products do not match, got %d instead of %d", adjProd(adjque), prod)
	}
}
