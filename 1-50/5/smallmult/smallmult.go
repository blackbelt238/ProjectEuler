package smallmult

// SmallestMultiple finds the smallest number that can be evenly divided
//   by all positive integers between 1 and the given number, up to and including it
func SmallestMultiple(uptoinc int) int {
	snum := 0
	found := false

	// continue searching for a number until it is found
	for !found {
		snum++
		found = true

		// attempt to divide all possible values of i into snum
		for i := 2; i <= uptoinc; i++ {
			// if division cannot be done evenly, move on to another number
			if snum%i != 0 {
				found = false
				break
			}
		}
	}

	return snum
}
