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
	cases := map[string]int{ // map input to expected output
		"...........\n.S-------7.\n.|F-----7|.\n.||.....||.\n.||.....||.\n.|L-7.F-J|.\n.|..|.|..|.\n.L--J.L--J.\n...........":                                                                                                        4,
		"..........\n.S------7.\n.|F----7|.\n.||....||.\n.||....||.\n.|L-7F-J|.\n.|..||..|.\n.L--JL--J.\n..........":                                                                                                                 4,
		".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ...": 8,
		"FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L": 10,
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			inputLns := strings.Split(input, "\n")

			actual, matrix := partTwo(inputLns)
			if !assert.Equal(t, expected, actual) {
				debugPrint(matrix)
			}
		})
	}
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
		connections: utils.NewSet[utils.Coord](utils.Coord{2, 0}, utils.Coord{X: 1, Y: 1}),
		matrix:      &actualMatrix,
	}

	assert.Equal(t, &expectedTwoOne, actualTwoOne)
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
			16,
		},
		{"start from start other dir - loop",
			start.coords,
			utils.Coord{0, 3},
			16,
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
			startCell, err := matrix.GetByCoord(c.start)
			assert.Nil(t, err)
			actualLoopCoords, actualIsLoop := startCell.isLoop(c.firstStep)
			assert.Equal(t, c.expectedSteps, len(actualLoopCoords))
			assert.Equal(t, c.expectedSteps != 0, actualIsLoop)
		})
	}
}

func TestFindAndMarkLoop(t *testing.T) {
	inputStr := ".....\n.S-7.\n.|.|.\n.L-J.\n....."
	matrix, start := matrixFromInput(strings.Split(inputStr, "\n"))

	expectedLoopCoords := utils.NewSet[utils.Coord](
		utils.Coord{1, 1}, utils.Coord{2, 1}, utils.Coord{3, 1},
		utils.Coord{1, 2}, utils.Coord{3, 2},
		utils.Coord{1, 3}, utils.Coord{2, 3}, utils.Coord{3, 3},
	)

	findAndMarkLoop(start)
	for _, cell := range matrix.Flatten() {
		expectedMark := UNKNOWN
		if expectedLoopCoords.Contains(cell.coords) {
			expectedMark = LOOP
		}
		assert.Equal(t, expectedMark, cell.mark,
			"expected mark %s but got %s (cell at (%d,%d) with value '%s')",
			expectedMark.toString(), cell.mark.toString(),
			cell.coords.X, cell.coords.Y, cell.val)
	}
}

func TestRadiateAdacent(t *testing.T) {
	cases := []struct {
		name              string
		input             string
		start             utils.Coord
		loopCoords        utils.Set[utils.Coord]
		expectedVisited   utils.Set[utils.Coord]
		expectedAnyIsEdge bool
	}{
		{"small grid only edges",
			"..\n..",
			utils.Coord{0, 0},
			utils.NewSet[utils.Coord](),
			utils.NewSet[utils.Coord](
				utils.Coord{0, 0}, utils.Coord{1, 0},
				utils.Coord{0, 1}, utils.Coord{1, 1},
			),
			true,
		},
		{"ignore loop tiles",
			"...\n.L.\n.L.",
			utils.Coord{0, 0},
			utils.NewSet[utils.Coord](
				utils.Coord{1, 1}, utils.Coord{1, 2},
			),
			utils.NewSet[utils.Coord](
				utils.Coord{0, 0}, utils.Coord{1, 0}, utils.Coord{2, 0},
				utils.Coord{0, 1}, utils.Coord{2, 1},
				utils.Coord{0, 2}, utils.Coord{2, 2},
			),
			true,
		},
		{"internal",
			"LLLLL\nL...L\nLLLLL",
			utils.Coord{2, 1},
			utils.NewSet[utils.Coord](
				utils.Coord{0, 0}, utils.Coord{1, 0}, utils.Coord{2, 0}, utils.Coord{3, 0}, utils.Coord{4, 0},
				utils.Coord{0, 1}, utils.Coord{4, 1},
				utils.Coord{0, 2}, utils.Coord{1, 2}, utils.Coord{2, 2}, utils.Coord{3, 2}, utils.Coord{4, 2},
			),
			utils.NewSet[utils.Coord](
				utils.Coord{1, 1}, utils.Coord{2, 1}, utils.Coord{3, 1},
			),
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m, _ := matrixFromInput(strings.Split(c.input, "\n"))
			markCoords(&m, c.loopCoords, LOOP) // setup: mark specified cells as part of the loop

			start, err := m.GetByCoord(c.start)
			assert.Nil(t, err)

			actualVisited, actualAnyIsEdge := start.radiateAdjacent(utils.NewSet[utils.Coord]())
			assert.Equal(t, c.expectedVisited, actualVisited)
			assert.Equal(t, c.expectedAnyIsEdge, actualAnyIsEdge)
		})
	}
}
