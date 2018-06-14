package main

import (
	tri "ProjectEuler/1-50/12/triangular"
	"fmt"
)

func main() {
	t := tri.CreateTriangular()
	ndiv := 400

	for t.NumFactors() <= ndiv {
		t.Next() // advance linearly (only "works" until ~100)
	}

	fmt.Printf("%d is the first triangular to have more than %d divisors\n", t.Val(), ndiv)
}
