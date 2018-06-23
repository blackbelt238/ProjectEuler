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
