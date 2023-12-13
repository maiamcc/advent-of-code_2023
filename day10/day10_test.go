package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "input\ngoes\nhere"
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
