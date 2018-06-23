package triangle

import (
	"strings"
	"testing"
)

func TestCreateTriangle(t *testing.T) {
	tri, _ := CreateTriangle("./../small.txt")
	exp := Triangle{
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

func TestCreateTriangleNoFile(t *testing.T) {
	tri, err := CreateTriangle("./../404.txt")
	if tri != nil || !strings.Contains(err.Error(), "opening Triangle file") {
		t.Errorf("should error on nonexistant file")
	}
}

func TestCreateTriangleIntConversionError(t *testing.T) {
	tri, err := CreateTriangle("./../bad.txt")
	if tri != nil || !strings.Contains(err.Error(), "converting Triangle elements to int") {
		t.Errorf("should error on trying to convert a word to an int")
	}
}

func TestMaxPathSum(t *testing.T) {
	tri, _ := CreateTriangle("./../small.txt")
	sum := MaxPathSum(tri)
	if sum != 23 {
		t.Errorf("incorrect sum: got %d but 23 expected", sum)
	}
}
