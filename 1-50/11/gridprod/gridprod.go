package gridprod

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// adjProd takes in a set of adjacent digits and returns their sum
func adjProd(adj []int) int {
	prod := 1
	for _, dig := range adj {
		prod *= dig
	}
	return prod
}

// calcProds takes a coordinate pair as a seed, checking respective products for right, down, and diagonal
func calcProds(grid [][]int, y, x int) int {
	var gprod, p int // greatest calculated product, temporary product

	// right
	p = calcDirection(grid, y, x, 0, 1)
	if p > gprod {
		gprod = p
	}

	// down
	p = calcDirection(grid, y, x, 1, 0)
	if p > gprod {
		gprod = p
	}

	// diagonal
	p = calcDirection(grid, y, x, 1, 1)
	if p > gprod {
		gprod = p
	}

	return gprod
}

// calcDirection
func calcDirection(grid [][]int, y, x, ydir, xdir int) int {
	// don't try a direction if it will result in
	if ydir > 0 {
		ydir = 1
		if !(y+3 < len(grid)) {
			return 0
		}
	} else {
		ydir = 0
	}
	if xdir > 0 {
		xdir = 1
		if !(x+3 < len(grid)) {
			return 0
		}
	} else {
		xdir = 0
	}

	adj := make([]int, 0, 4) // slice to hold adjacent digits
	for i := 0; i < 4; i++ {
		adj = append(adj, grid[y][x])
		x += xdir
		y += ydir
	}

	return adjProd(adj)
}

// makeGrid reads from the given file and builds a grid of numbers
func makeGrid(fname string, dim int) ([][]int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", fname, err)
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
