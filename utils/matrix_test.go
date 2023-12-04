package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	input := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
	}
	expectedCells := [][]Cell{
		[]Cell{{0, 0, "a"}, {1, 0, "b"}, {2, 0, "c"}},
		[]Cell{{0, 1, "d"}, {1, 1, "e"}, {2, 1, "f"}},
		[]Cell{{0, 2, "g"}, {1, 2, "h"}, {2, 2, "i"}},
		[]Cell{{0, 3, "j"}, {1, 3, "k"}, {2, 3, "l"}},
	}
	actual, err := NewMatrix(input)
	assert.Nil(t, err)
	assert.Equal(t, actual.Cells, expectedCells)
	assert.Equal(t, actual.NumRows, 4)
	assert.Equal(t, actual.NumCols, 3)
}

func TestNewMatrixUnequalRows(t *testing.T) {
	input := []string{
		"abc",
		"def",
		"ghij",
	}
	_, err := NewMatrix(input)
	assert.Error(t, err)
}

func TestMatrix_Get(t *testing.T) {
	matrix := MustMatrix(
		[]string{
			"abc",
			"def",
			"ghi",
			"jkl",
		})
	cases := []struct {
		x           int
		y           int
		expectedVal string // empty string --> expect an error
	}{
		{0, 0, "a"},
		{1, 1, "e"},
		{2, 3, "l"},
		{0, 6, ""},
		{3, 1, ""},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("(%d, %d)", c.x, c.y), func(t *testing.T) {
			actual, err := matrix.Get(c.x, c.y)
			if c.expectedVal != "" {
				assert.Nil(t, err)
				assert.Equal(t, Cell{c.x, c.y, c.expectedVal}, actual)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestMatrix_Flatten(t *testing.T) {
	matrix := MustMatrix(
		[]string{
			"abc",
			"def",
		})
	expected := []Cell{
		{0, 0, "a"}, {1, 0, "b"}, {2, 0, "c"},
		{0, 1, "d"}, {1, 1, "e"}, {2, 1, "f"},
	}

	actual := matrix.Flatten()
	assert.Equal(t, expected, actual)
}
