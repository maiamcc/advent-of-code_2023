package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
	inputLns := strings.Split(inputStr, "\n")

	actual := partOne(inputLns)
	assert.Equal(t, 4361, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
	inputLns := strings.Split(inputStr, "\n")

	actual := partTwo(inputLns)
	assert.Equal(t, 467835, actual)
}

func TestGetAdjacentCoords(t *testing.T) {
	actual := getAdjacentCoords(utils.Coord{0, 1})
	expected := []utils.Coord{
		{-1, 0}, {0, 0}, {1, 0},
		{-1, 1}, {1, 1},
		{-1, 2}, {0, 2}, {1, 2},
	}
	assert.ElementsMatch(t, expected, actual)
}

func TestGetAllAdjacentCoords(t *testing.T) {
	actual := getAllAdjacentCoords([]utils.Coord{{0, 1}, {1, 3}})
	expected := []utils.Coord{
		{-1, 0}, {0, 0}, {1, 0},
		{-1, 1}, {1, 1},
		{-1, 2}, {0, 2}, {1, 2}, {2, 2},
		{0, 3}, {2, 3},
		{0, 4}, {1, 4}, {2, 4},
	}
	assert.ElementsMatch(t, expected, actual)
}

func TestNumberCells_asInt(t *testing.T) {
	cases := map[string]int{ // map input to expected output
		"467":     467,
		"46a4":    -1, // indicates expected error
		"4":       4,
		"6177777": 6177777,
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			var cells numberCells
			for _, ch := range strings.Split(input, "") {
				cells = append(cells, utils.Cell{Val: ch}) // trivial cell with no coords
			}

			i, err := cells.asInt()
			if expected != -1 {
				assert.Nil(t, err)
				assert.Equal(t, expected, i)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestNumbersForRow(t *testing.T) {
	cases := map[string][]int{ // map input to expected output
		"467..114..": {467, 114},
		"...*......": {},
		"..35..633":  {35, 633},
		"......2...": {2},
		"6177777":    {6177777},
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			matrix := utils.MustMatrix([]string{input}) // trivial matrix
			cells := matrix.Cells[0]                    // first row the matrix represents the row we fed in
			numbers := numbersForRow(cells)

			actual := []int{} // no really, initialize it to empty so it matches the return value
			for _, numCells := range numbers {
				num, err := numCells.asInt()
				if err != nil {
					t.FailNow()
				}
				actual = append(actual, num)
			}
			assert.Equal(t, expected, actual)
		})
	}
}

func TestDedupeCoords(t *testing.T) {
	input := []utils.Coord{
		{3, 1},
		{6, 3},
		{3, 4},
		{50, 0},
		{0, 50},
		{3, 1}, // dupe
		{2, 2},
		{2, 2}, // dupe
		{3, 4}, // dupe
	}
	expected := []utils.Coord{
		{3, 1},
		{6, 3},
		{3, 4},
		{50, 0},
		{0, 50},
		{2, 2},
	}
	assert.ElementsMatch(t, expected, dedupeCoords(input))
}

func TestIsSymbol(t *testing.T) {
	matrix := utils.MustMatrix(
		[]string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
		})

	cases := map[utils.Coord]bool{ // map input to expected output
		{2, 0}:   false, // digit
		{1, 1}:   false, // period
		{3, 1}:   true,  // *
		{6, 3}:   true,  // #
		{-6, -9}: false, // invalid coords dont error, just return false
	}

	for input, expected := range cases {
		t.Run(fmt.Sprintf("(%d, %d)", input.X, input.Y), func(t *testing.T) {
			assert.Equal(t, expected, isSymbol(matrix, input))
		})
	}
}

func TestAnyIsSymbol(t *testing.T) {
	sourceStrings := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*..*...",
	}
	matrix := utils.MustMatrix(sourceStrings)

	// key is y index of row
	cases := map[int]bool{ // map input to expected output
		0: false,
		1: true,
		2: false,
		3: true,
		4: true,
	}

	for y, expected := range cases {
		t.Run(sourceStrings[y], func(t *testing.T) {
			coords := numberCells(matrix.Cells[y]).coords() // cheating, borrowing method off numberCells
			actual, err := anyIsSymbol(matrix, coords)
			assert.Nil(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}
