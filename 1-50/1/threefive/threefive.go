package threefive

// ThreeAndFive finds the sum of all natural numbers below num that are multiples of 3 or 5
func ThreeAndFive(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	return sum
}
