package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)"

	actual := partOne(inputStr)
	assert.Equal(t, 6, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)\n"

	actual := partTwo(inputStr)
	assert.Equal(t, 6, actual)
}

func TestStartEndNodes(t *testing.T) {
	dirs, nodes := parseInput(utils.MustReadFileAsString("input.txt"))
	startNodes := nodes.findEndNodes()
	var endNodes []*node
	for _, n := range startNodes {
		count, endNode := stepsToEndNodes(dirs, nodes, n.label, func(n *node) bool {
			return n.isTerminalNode()
		})
		endNodes = append(endNodes, endNode)
		fmt.Printf("%s --> %s in %d steps\n", n.label, endNode.label, count)
	}
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
