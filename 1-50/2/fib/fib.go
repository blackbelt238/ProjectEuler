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

// UpToMax returns a list of all Fibonacci numbers that do not exceed max
func UpToMax(max int) []int {
	fibl := []int{0, 1} // list of Fibonacci numbers

	var next int
	for next <= max {
		next = fibl[len(fibl)-1] + fibl[len(fibl)-2] // calculate the next number to add
		fibl = append(fibl, next)                    // add next to the list
	}

	// return the constructed list without the final entry,
	//   since it exceeds the max
	return fibl[:len(fibl)-1]
}
