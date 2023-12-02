package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	inputLns := strings.Split(inputStr, "\n")

	actual := partOne(inputLns)
	assert.Equal(t, 8, actual)
}

func TestParseRound(t *testing.T) {
	cases := map[string]round{ // map input to expected output
		"5 red":                        {5, 0, 0},
		"5 green":                      {0, 5, 0},
		"5 blue":                       {0, 0, 5},
		"1 blue, 2 green, 3 red":       {3, 2, 1},
		"6 red, 9 blue":                {6, 0, 9},
		"123 green, 456 red, 789 blue": {456, 123, 789},
		"bad string":                   {0, 0, 0},
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			actual := parseRound(input)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParseGame(t *testing.T) {
	cases := map[string]game{ // map input to expected output
		"Game 1: 5 red": {1, []round{{5, 0, 0}}},
		"Game 2: 5 green; 5 blue": {2, []round{
			{0, 5, 0},
			{0, 0, 5}}},
		"Game 3: 1 blue, 2 green, 3 red; 3 blue, 2 green, 1 red; 1 blue, 1 green, 1 red": {3,
			[]round{
				{3, 2, 1},
				{1, 2, 3},
				{1, 1, 1}}},
		"Game 4: 6 red; 9 blue, 5 red; 6 green": {4,
			[]round{
				{6, 0, 0},
				{5, 0, 9},
				{0, 6, 0},
			}},
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			actual, err := parseGame(input)
			assert.Nil(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}
