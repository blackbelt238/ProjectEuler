package main

import (
	prime "ProjectEuler/1-50/3/prime"
	"fmt"
)

func main() {
	num := 600851475143
	pfacs := prime.Factorize(num) // perform prime factorization on num

	// determine the largest of the prime factors
	lfac := 0
	for k := range pfacs {
		if k > lfac {
			lfac = k
		}
	}

	fmt.Println(lfac)
}
