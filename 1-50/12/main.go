package main

import (
	tri "ProjectEuler/1-50/12/triangular"
	"fmt"
)

func main() {
	t := tri.CreateTriangular()
	ndiv := 500

	for t.NumFactors() <= ndiv {
		t.Next() // advance linearly, since the number of factors in a given triangular varies unpredictably
	}

	fmt.Printf("%d is the first triangular to have more than %d divisors\n", t.Val(), ndiv)
}
