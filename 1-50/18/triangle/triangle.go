package triangle

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Triangle represents a number of rows of integers whose length equals the row number
type Triangle [][]int

// CreateTriangle goes to the given file and creates a Triangle out of the contents
func CreateTriangle(fname string) (Triangle, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("opening Triangle file: %v", err)
	}
	defer file.Close()

	t := Triangle{} // start with an empty Triangle

	// scan through the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		elems := strings.Split(row, " ") // obtain all elements from the row
		trow := []int{}                  // the row of numbers to add to the Triangle

		// go through each element from the file and add it to trow
		for _, elem := range elems {
			ele, err := strconv.Atoi(elem)
			if err != nil {
				return nil, fmt.Errorf("converting Triangle elements to int: %v", err)
			}
			trow = append(trow, ele)
		}
		t = append(t, trow) // once a trow is built, add it to the Triangle
	}

	return t, nil
}

// MaxPathSum looks 2 rows ahead to determine the max path
func MaxPathSum(t Triangle) int {
	sum := t[0][0] // initialize the sum at the only number in the first row
	lgst := 0      // index of largest adjacent number for the current row

	// go through all rows in the triangle and find the path
	var llgst, rlgst int
	for row := 1; row < len(t); row++ {
		llgst, rlgst = largestLookaheadSums(t, row, lgst)

		// pick which adjacent number to use based on the lookahead sums
		if llgst < rlgst {
			lgst++
		}
		sum += t[row][lgst]
	}

	return sum
}

func largestLookaheadSums(t Triangle, row int, lgst int) (int, int) {
	llgst, rlgst := t[row][lgst], t[row][lgst+1] // begin the calculation with the 2 adjacent numbers on this row
	// if this is the last row, don't look ahead
	if row >= len(t)-1 {
		return llgst, rlgst
	}

	// looking ahead to the next row, determine the largest possible left sum
	llgst += t[row+1][lgst]
	if llgst < t[row][lgst]+t[row+1][lgst+1] {
		llgst = t[row][lgst] + t[row+1][lgst+1]
	}

	// looking ahead to the next row, determine the largest possible right sum
	rlgst += t[row+1][lgst+1]
	if rlgst < t[row][lgst+1]+t[row+1][lgst+2] {
		rlgst = t[row][lgst+1] + t[row+1][lgst+2]
	}

	return llgst, rlgst
}
