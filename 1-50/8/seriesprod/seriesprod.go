package seriesprod

import (
	"strconv"
)

// adjCycle removes the digit at the front of the queue and adds the next one to the back
func adjCycle(adjque []string, next string) []string {
	return append(adjque[1:], next)
}

// adjProcess does all processing across the given number string
func adjProcess(adjque []string, numstr string, gprod int) ([]string, int) {
	var nprod int

	// for each digit in the number string
	for _, dig := range numstr {
		adjque = adjCycle(adjque, string(dig)) // cycle in next digit
		nprod = adjProd(adjque)                // find the new product
		if nprod > gprod {
			gprod = nprod
		}
	}
	return adjque, gprod
}

// adjProd takes in a set of adjacent digits and returns their sum
func adjProd(adjque []string) int {
	prod := 1
	for _, dig := range adjque {
		digint, err := strconv.Atoi(dig) // int representation of dig
		if err != nil {
			return 0 // cannot have a product of digits if there is an invalid digit
		}
		prod *= digint
	}
	return prod
}
