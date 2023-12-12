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
		{SimpleCell{0, 0, "a"}, SimpleCell{1, 0, "b"}, SimpleCell{2, 0, "c"}},
		{SimpleCell{0, 1, "d"}, SimpleCell{1, 1, "e"}, SimpleCell{2, 1, "f"}},
		{SimpleCell{0, 2, "g"}, SimpleCell{1, 2, "h"}, SimpleCell{2, 2, "i"}},
		{SimpleCell{0, 3, "j"}, SimpleCell{1, 3, "k"}, SimpleCell{2, 3, "l"}},
	}
	actual, err := NewMatrix(input, NewSimpleCell)
	assert.Nil(t, err)
	assert.Equal(t, expectedCells, actual.Cells)
	assert.Equal(t, 4, actual.NumRows)
	assert.Equal(t, 3, actual.NumCols)
}

func TestNewMatrixUnequalRows(t *testing.T) {
	input := []string{
		"abc",
		"def",
		"ghij",
	}
	_, err := NewMatrix(input, NewSimpleCell)
	assert.Error(t, err)
}

func TestMatrix_Get(t *testing.T) {
	matrix := MustSimpleCellMatrix(
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
		{1, -1, ""},
		{-1, 1, ""},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("(%d, %d)", c.x, c.y), func(t *testing.T) {
			actual, err := matrix.Get(c.x, c.y)
			if c.expectedVal != "" {
				assert.Nil(t, err)
				assert.Equal(t, SimpleCell{c.x, c.y, c.expectedVal}, actual)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestMatrix_Flatten(t *testing.T) {
	matrix := MustSimpleCellMatrix(
		[]string{
			"abc",
			"def",
		})
	expected := []Cell{
		SimpleCell{0, 0, "a"}, SimpleCell{1, 0, "b"}, SimpleCell{2, 0, "c"},
		SimpleCell{0, 1, "d"}, SimpleCell{1, 1, "e"}, SimpleCell{2, 1, "f"},
	}

	actual := matrix.Flatten()
	assert.Equal(t, expected, actual)
}
