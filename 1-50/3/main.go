package main

import (
	primefactor "ProjectEuler/1-50/3/primefactor"
	"fmt"
)

func main() {
	num := 600851475143
	pfacs := primefactor.PrimeFactors(num)

	// find the largest prime factor returned
	fmt.Println(pfacs[len(pfacs)-1])
}
