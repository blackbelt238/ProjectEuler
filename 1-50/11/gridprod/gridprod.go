package gridprod

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LargestGridProd finds the biggest product of 4 adjacent numbers in the grid at the given location
func LargestGridProd(fname string, dim int) int {
	grid, err := makeGrid(fname, dim)
	if err != nil {
		fmt.Printf("Couldn't make a grid from %s: %v\n", fname, err)
		return 0
	}

	var gprod int // the greatest product found so far

	// for each square
	for i := range grid {
		for j := range grid[i] {
			prod := calcProds(grid, i, j)
			if prod > gprod {
				gprod = prod
			}
		}
	}
	return gprod
}

// adjProd takes in a slice of adjacent digits and returns their product
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

	// diagonal forward
	p = calcDirection(grid, y, x, 1, 1)
	if p > gprod {
		gprod = p
	}

	// diagonal backward
	p = calcDirection(grid, y, x, 1, -1)
	if p > gprod {
		gprod = p
	}

	return gprod
}

// calcDirection returns the product of a single direction
func calcDirection(grid [][]int, y, x, ydir, xdir int) int {
	// don't try to calculate if a direction will be out of range
	if ydir != 0 {
		if !(y+ydir*3 < len(grid)) || !(y+ydir*3 > -1) {
			return 0
		}
	}
	if xdir != 0 {
		if !(x+xdir*3 < len(grid[y])) || !(x+xdir*3 > -1) {
			return 0
		}
	}

	// build the adjacent digits slice to pass to adjProd
	adj := make([]int, 0, 4)
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

	// initialize the grid
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
