package main

import (
	"github.com/stretchr/testify/assert"
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
