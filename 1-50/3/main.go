package main

import (
	prime "ProjectEuler/1-50/3/prime"
	"fmt"
)

func main() {
	num := 600851475143
	pfacs := prime.Factors(num)

	// find the largest prime factor returned
	fmt.Println(pfacs[len(pfacs)-1])
}
