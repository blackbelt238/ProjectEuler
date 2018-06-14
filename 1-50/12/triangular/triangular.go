package triangular

import (
	prime "ProjectEuler/1-50/3/prime"
)

// Triangular is an object used to interact with triangular numbers
type Triangular struct {
	val int // the current triangular number
	n   int // which triangular number val represents
}

// CreateTriangular returns a new Triangular starting at the 1st trianular number
func CreateTriangular() *Triangular {
	return &Triangular{
		val: 1,
		n:   1, // 1 is the 1st triangular number
	}
}

// Find determines the n-th triangular number
func (t *Triangular) Find(n int) int {
	// determine which direction to advance in
	adv := t.Next
	if t.n > n {
		adv = t.Prev
	}

	// advance until the n-th triangular is found
	for t.n != n {
		adv()
	}
	return t.val
}

// N is a getter for n
func (t *Triangular) N() int {
	return t.n
}

// Next determines the next triangular above val
func (t *Triangular) Next() {
	t.n++
	t.val += t.n
}

// NumFactors determines how many factors the current triangular has
func (t *Triangular) NumFactors() int {
	pfacs := prime.Factorize(t.val)
	numfacs := 1

	for _, v := range pfacs {
		numfacs *= v + 1
	}

	return numfacs
}

// Prev determines the previous triangular below val
func (t *Triangular) Prev() {
	t.val -= t.n
	t.n--
}

// Val is a getter for val
func (t *Triangular) Val() int {
	return t.val
}
