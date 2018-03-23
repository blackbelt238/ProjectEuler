package seriesprod

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// LargestSeriesProd uses the number in the given file to determine the largest product of n adjacent digits
func LargestSeriesProd(fname string, numdigs int) int {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("could not open %s: %v\n", fname, err)
		return 0
	}
	defer file.Close()

	adjque := make([]string, numdigs, numdigs) // adjacent digits
	var gprod int                              // the greatest prod found so far

	var numstr string              // the line of numbers to check against
	line := bufio.NewScanner(file) // scanner for the file
	for line.Scan() {              // for each line in the file, process the line
		numstr = line.Text()
		adjque, gprod = adjProcess(adjque, numstr, gprod)
	}
	return gprod
}

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
