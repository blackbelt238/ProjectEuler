package gridprod

import (
	"fmt"
	"testing"
)

func TestAdjProd(t *testing.T) {
	adjque := []int{26, 63, 78, 14}
	if adjProd(adjque) != 1788696 {
		t.Errorf("Improper product calculation, got %d, want %d", adjProd(adjque), 1788696)
	}
}

func TestCalcProds(t *testing.T) {
	grid, _ := makeGrid("./../grid.txt", 4)
	p := calcProds(grid, 0, 0)
	if p != 1651104 {
		t.Errorf("Improper product calculated, got %d, want %d", p, 1651104)
	}
}

func TestLargestGridProduct(t *testing.T) {
	p := LargestGridProd("./ezgrid.txt", 4)
	if p != 144 {
		t.Errorf("Improper product calculated, got %d, want %d", p, 144)
	}
}

// TestMakeGrid ensures makeGrid is properly parsing the file and generating a matrix
func TestMakeGrid(t *testing.T) {
	grid, err := makeGrid("./../grid.txt", 4)
	if err != nil {
		t.Errorf("Making grid: %v", err)
	}
	gexp := [][]int{{8, 2, 22, 97}, {49, 49, 99, 40}, {81, 49, 31, 73}, {52, 70, 95, 23}}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != gexp[i][j] {
				t.Errorf("grid unexpected\n\texpected: %v\n\trecieved: %v", gexp, grid)
			}
		}
	}
}

// printGrid prints the given grid. NOTE: keep for debugging purposes
func printGrid(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("|%2d|", grid[i][j])
		}
		fmt.Println()
	}
}
