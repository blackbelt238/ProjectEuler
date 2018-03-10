package fib

// Fib calculates the n-th Fibonacci digit
func Fib(n int) int {
	if n == 0 { // n = 0
		return 0
	} else if n == 1 { // n = 1
		return 1
	} else {
		return Fib(n-1) + Fib(n-2) // Fn = Fn-1 + Fn-2
	}
}

// UpToLimit returns a list of all Fibonacci numbers whose sum does not exceed limit
func UpToLimit(limit int) []int {
	fibl := []int{0, 1} // list of Fibonacci numbers
	sum := 1            // sum of all numbers in fibl

	var next int
	for sum <= limit {
		next = fibl[len(fibl)-1] + fibl[len(fibl)-2] // calculate the next number to add
		fibl = append(fibl, next)                    // add next to the list
		sum += next                                  // update the sum
	}

	// return the constructed list without the final entry,
	//   since it caused the sum to exceed the limit
	return fibl[:len(fibl)-2]
}
