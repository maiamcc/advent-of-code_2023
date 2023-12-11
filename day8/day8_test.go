package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)"

	actual := partOne(inputStr)
	assert.Equal(t, 6, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "input\ngoes\nhere"

	actual := partTwo(inputStr)
	assert.Equal(t, 123, actual)
}

func TestParseNodes(t *testing.T) {
	input := `AAA = (BBB, CCC)
BBB = (DDD, AAA)
CCC = (AAA, BBB)`

	// expected
	a := node{"AAA", nil, nil}
	b := node{"BBB", nil, nil}
	c := node{"CCC", nil, nil}
	d := node{"DDD", nil, nil}
	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &a
	c.left = &a
	c.right = &b
	expectedMap := map[string]*node{
		"AAA": &a,
		"BBB": &b,
		"CCC": &c,
		"DDD": &d,
	}

	actual := parseNodes(input)
	assert.Equal(t, expectedMap, actual)
}
