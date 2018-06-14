package triangular

// Triangular is an object used to interact with triangular numbers
type Triangular struct {
	Val int // the current triangular number
	n   int // which triangular number Val represents
}

// CreateTriangular returns a new Triangular starting at the 1st trianular number
func CreateTriangular() *Triangular {
	return &Triangular{
		Val: 1,
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
	return t.Val
}

// Prev determines the previous triangular below Val
func (t *Triangular) Prev() {
	t.Val -= t.n
	t.n--
}

// Next determines the next triangular above Val
func (t *Triangular) Next() {
	t.n++
	t.Val += t.n
}

// NumFactors determines how many factors the current triangular has
func (t *Triangular) NumFactors() int {
	numfacs := 0

	for i := 1; i <= t.Val; i++ {
		if t.Val%i == 0 {
			numfacs++
		}
	}
	return numfacs
}
