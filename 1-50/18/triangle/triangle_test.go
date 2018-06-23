package triangle

import (
	"testing"
)

func TestCreateTriangle(t *testing.T) {
	tri, _ := CreateTriangle("./../small.txt")
	exp := [][]int{
		{3},
		{7, 4},
		{2, 4, 6},
		{8, 5, 9, 3},
	}

	for i, row := range tri {
		for j, ele := range row {
			expele := exp[i][j] // the expected element
			if expele != ele {
				t.Errorf("unexpected element found: %d should be %d\n\tgot: %v\n\texp: %v", ele, expele, tri, exp)
			}
		}
	}
}
