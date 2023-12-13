package main

import (
	"github.com/maiamcc/advent-of-code_2023/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ.LJ"
	inputLns := strings.Split(inputStr, "\n")

	actual := partOne(inputLns)
	assert.Equal(t, 8, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "input\ngoes\nhere"
	inputLns := strings.Split(inputStr, "\n")

	actual := partTwo(inputLns)
	assert.Equal(t, 123, actual)
}

func TestMatrixFromInput(t *testing.T) {
	inputStr := "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ.LJ"
	actualMatrix, actualStart := matrixFromInput(strings.Split(inputStr, "\n"))

	// just spotchecking
	assert.Equal(t, 5, len(actualMatrix.Cells))
	assert.Equal(t, 5, len(actualMatrix.Cells[0]))

	actualTwoOne, err := actualMatrix.Get(2, 1)
	assert.Nil(t, err)
	expectedTwoOne := pipeCell{
		coords:      utils.Coord{X: 2, Y: 1},
		val:         "J",
		connections: utils.NewSet[utils.Coord](utils.Coord{2, 0}, utils.Coord{1, 1}),
		matrix:      &actualMatrix,
	}

	assert.Equal(t, expectedTwoOne, actualTwoOne)
	assert.Equal(t, utils.Coord{0, 2}, actualStart.coords)
	assert.Equal(t, utils.NewSet[utils.Coord](), actualStart.connections)
}

func TestIsLoop(t *testing.T) {
	inputStr := "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ---"
	matrix, start := matrixFromInput(strings.Split(inputStr, "\n"))

	cases := []struct {
		name          string
		start         utils.Coord
		firstStep     utils.Coord
		expectedSteps int // 0 implies we expect !isLoop
	}{
		{"start from period - no loop",
			utils.Coord{0, 1},
			utils.Coord{1, 1},
			0,
		},
		{"start from start one dir - loop",
			start.coords,
			utils.Coord{1, 2},
			15,
		},
		{"start from start other dir - loop",
			start.coords,
			utils.Coord{0, 3},
			15,
		},
		{"start from start wrong dir - no loop",
			start.coords,
			utils.Coord{0, 1},
			0,
		},
		{"start promising but dead end - no loop",
			utils.Coord{4, 4},
			utils.Coord{3, 4},
			0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			startCellGeneric, err := matrix.GetByCoord(c.start)
			assert.Nil(t, err)
			startCell := startCellGeneric.(pipeCell)
			actualSteps, actualIsLoop := startCell.isLoop(c.firstStep)
			assert.Equal(t, c.expectedSteps, actualSteps)
			assert.Equal(t, c.expectedSteps != 0, actualIsLoop)
		})
	}
}
