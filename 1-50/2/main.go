package main

import (
	fib "ProjectEuler/1-50/2/fib"
	"fmt"
)

func main() {
	fnums := fib.UpToMax(4000000)

	sumEven := 0 // sum of all even numbers in fnums
	for _, num := range fnums {
		if num%2 == 0 {
			sumEven += num
		}
	}

	fmt.Println(sumEven)
}
