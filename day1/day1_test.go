package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetNumericChars(t *testing.T) {
	cases := map[string][]string{ // map input to expected output
		"1hello2":      {"1", "2"},
		"hel1lo2world": {"1", "2"},
		"321stuff":     {"3", "2", "1"},
		"nonumbers":    {},
		"123":          {"1", "2", "3"},
	}

	for input, expected := range cases {
		actual := getNumericChars(input)
		assert.Equal(t, expected, actual, "for input: %s", input)
	}
}

type ExpectedNumStr struct {
	num int
	str string
}

func TestConsumeFirstNumber(t *testing.T) {
	cases := map[string]ExpectedNumStr{ // map input to expected output
		"1hello2":      {1, "hello2"},
		"hel1lo2world": {1, "lo2world"},
		"nonumbers":    {-1, "nonumbers"},
		"sixnine":      {6, "ixnine"}, // sigh
		"eightwo":      {8, "ightwo"}, // sigh
		"stuff5six":    {5, "six"},
		"stuffsix5":    {6, "ix5"}, // sigh
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			actualNum, actualRest, err := consumeFirstNumber(input)
			assert.Nil(t, err)
			assert.Equal(t, expected.num, actualNum)
			assert.Equal(t, expected.str, actualRest)
		})
	}
}

func TestGetAllNumbers(t *testing.T) {
	cases := map[string][]int{ // map input to expected output
		"hello2":       {2},
		"1hello2":      {1, 2},
		"hel1lo2world": {1, 2},
		"321stuff":     {3, 2, 1},
		"nonumbers":    {},
		"sixnine":      {6, 9},
		"eightwo":      {8, 2}, // this is the case that fucks you up!
		"stuff5six":    {5, 6},
		"stuffsix5":    {6, 5},
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			actual, err := getAllNumbers(input)
			assert.Nil(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestPartTwo(t *testing.T) {
	inputStr := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	inputLns := strings.Split(inputStr, "\n")

	actual := partTwo(inputLns)
	assert.Equal(t, 281, actual)
}
