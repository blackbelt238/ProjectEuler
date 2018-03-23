package main

import (
	prime "ProjectEuler/1-50/3/prime"
	"fmt"
)

func main() {
	primes := prime.SieveOfEratosthenes(2000000)
	var sum int
	for _, prime := range primes {
		sum += prime
	}
	fmt.Println(sum)
}
