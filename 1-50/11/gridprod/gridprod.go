package gridprod

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// makeGrid reads from the given file and builds a grid of numbers
func makeGrid(fname string, dim int) ([][]int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v\n", fname, err)
	}
	defer file.Close()

	grid := make([][]int, dim, dim)
	for i := range grid {
		grid[i] = make([]int, dim, dim)
	}

	row := bufio.NewScanner(file) // scanner for the file
	for y := 0; y < len(grid); y++ {
		row.Scan()
		line := strings.Split(row.Text(), " ")
		for x := 0; x < len(grid[y]); x++ {
			dig, err := strconv.Atoi(string(line[x]))
			if err != nil {
				return nil, fmt.Errorf("couldn't populate cell with value: %v", err)
			}
			grid[y][x] = dig
		}
	}
	return grid, nil
}
