package ssq

// SumSquareDifference finds the difference between the sqs and ssq for all natural numbers up to and including uptoinc
func SumSquareDifference(uptoinc int) int {
	ssq := SSQ(uptoinc)
	sqs := SQS(uptoinc)

	return sqs - ssq
}

// SSQ finds the sum of the squares for all natural numbers up to and including uptoinc
func SSQ(uptoinc int) int {
	var sum int
	// sum up the squares
	for i := 1; i <= uptoinc; i++ {
		sum += i * i
	}
	return sum
}

// SQS finds the square of the sum for all natural numbers up to and including uptoinc
func SQS(uptoinc int) int {
	var sum int
	// sum the numbers
	for i := 1; i <= uptoinc; i++ {
		sum += i
	}
	return sum * sum // return the square of the sum
}
