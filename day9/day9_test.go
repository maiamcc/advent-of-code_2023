package main

import (
	"github.com/maiamcc/advent-of-code_2023/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"
	inputLns := strings.Split(inputStr, "\n")

	actual := partOne(inputLns)
	assert.Equal(t, 114, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"
	inputLns := strings.Split(inputStr, "\n")

	actual := partTwo(inputLns)
	assert.Equal(t, 2, actual)
}

func TestDiffList(t *testing.T) {
	cases := map[string][]int{ // map input (as string) to expected output
		"1 2 3 4 5":   {1, 1, 1, 1},
		"5 4 3 2 1":   {-1, -1, -1, -1},
		"6 6 106 200": {0, 100, 94},
		"5":           nil,
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			ints := utils.MustSplitStringToInts(input, " ")
			assert.Equal(t, expected, diffList(ints))
		})
	}
}

func TestDiffAllList(t *testing.T) {
	cases := map[string][][]int{ // map input (as string) to expected output
		"1 2 3 4 5":            {{1, 1, 1, 1}, {0, 0, 0}},
		"5 4 3 2 1":            {{-1, -1, -1, -1}, {0, 0, 0}},
		"6 6 106 200":          {{0, 100, 94}, {100, -6}, {-106}, nil},
		"10 13 16 21 30 45 68": {{3, 3, 5, 9, 15, 23}, {0, 2, 4, 6, 8}, {2, 2, 2, 2}, {0, 0, 0}},
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			ints := utils.MustSplitStringToInts(input, " ")
			assert.Equal(t, expected, allDiffLists(ints))
		})
	}
}

func TestExtrapolateRightElem(t *testing.T) {
	cases := map[string]int{ // map input (as string) to expected output
		"0 3 6 9 12 15":     18,
		"1 3 6 10 15 21":    28,
		"10 13 16 21 30 45": 68,
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			ints := utils.MustSplitStringToInts(input, " ")
			assert.Equal(t, expected, extrapolateRightElem(ints))
		})
	}
}
