package collatz

// Sequence finds the Collatz sequence iteratively
func Sequence(seed int) []int {
	seq := make([]int, 0, 10)

	// work through the sequence until the finishing value is found
	for seed != 1 {
		seq = append(seq, seed)
		if seed%2 == 0 {
			seed = seed / 2
		} else {
			seed = 3*seed + 1
		}
	}
	seq = append(seq, 1)

	return seq
}

// SequenceLen finds the length of the Collatz sequence iteratively
func SequenceLen(seed int) int {
	seqlen := 0

	// work through the sequence until the finishing value is found
	for seed != 1 {
		seqlen++
		if seed%2 == 0 {
			seed = seed / 2
		} else {
			seed = 3*seed + 1
		}
	}
	seqlen++

	return seqlen
}
