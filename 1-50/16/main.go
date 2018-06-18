package main

import (
	longmath "ProjectEuler/1-50/13/longmath"
	"fmt"
)

func main() {
	prod := "1"

	// obtain 2¹⁰⁰⁰ via long multiplication
	for i := 0; i < 1000; i++ {
		prod = longmath.LongMult(prod, "2")
	}

	// print the sum of the digits
	fmt.Println(sumdigs(prod))
}

func sumdigs(num string) string {
	sum := "0"

	for i := len(num) - 1; i >= 0; i-- {
		sum = longmath.LongAdd(sum, string(num[i]))
	}

	return sum
}
