package main

import (
	"github.com/maiamcc/advent-of-code_2023/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
	inputLns := strings.Split(inputStr, "\n")

	actual := partOne(inputLns)
	assert.Equal(t, 123, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "input\ngoes\nhere"
	inputLns := strings.Split(inputStr, "\n")

	actual := partTwo(inputLns)
	assert.Equal(t, 123, actual)
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

func TestCoordsForSymbols(t *testing.T) {
	matrix := utils.MustMatrix(
		[]string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
		})
	expected := []utils.Coord{
		{3, 1}, // *
		{6, 3}, // #
		{3, 4}, // *
	}
	actual := coordsForSymbols(matrix)

	assert.ElementsMatch(t, expected, actual)
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
