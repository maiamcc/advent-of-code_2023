package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"os"
	"regexp"
	"strings"
)

var nodeRe = regexp.MustCompile("([A-Z]{3}) = \\(([A-Z]{3}), ([A-Z]{3})\\)")

func main() {
	fullInput := utils.MustReadFileAsString("day8/input.txt")
	fmt.Println("The answer to Part One is:", partOne(fullInput))
	//fmt.Println("The answer to Part Two is:", partTwo(fullInput))
}

func partOne(fullInput string) int {
	dirs, nodeMap := parseInput(fullInput)

	curNode, ok := nodeMap["AAA"]
	if !ok {
		utils.LogfAndExit("couldn't find start node")
	}
	count := 0
	for {
		for _, dir := range dirs {
			curNode = curNode.step(dir)
			count += 1
			if curNode.label == "ZZZ" {
				// win condition
				return count
			}
		}
	}
}

func partTwo(fullInput string) int {
	return len(fullInput)
}

type node struct {
	label string
	left  *node
	right *node
}

func (n *node) step(dir string) *node {
	if dir == "L" {
		return n.left
	} else if dir == "R" {
		return n.right
	}
	utils.LogfAndExit("Unrecognized direction string '%s'\n", dir)
	return nil
}

func parseNode(input string, nodeMap map[string]*node) {
	submatches := nodeRe.FindStringSubmatch(input)
	if len(submatches) != 4 { // whole string match, + 3 submatches
		fmt.Printf("Error parsing node input '%s'\n", input)
		os.Exit(1)
	}
	selfLabel := submatches[1]
	leftLabel := submatches[2]
	rightLabel := submatches[3]

	leftNode, ok := nodeMap[leftLabel]
	if !ok {
		leftNode = &node{leftLabel, nil, nil}
		nodeMap[leftLabel] = leftNode
	}
	rightNode, ok := nodeMap[rightLabel]
	if !ok {
		rightNode = &node{rightLabel, nil, nil}
		nodeMap[rightLabel] = rightNode
	}
	selfNode, ok := nodeMap[selfLabel]
	if !ok {
		selfNode = &node{selfLabel, leftNode, rightNode}
		nodeMap[selfLabel] = selfNode
	} else {
		// node exists as created by some previous parsing,
		// it just doesn't have destination values populated
		selfNode.left = leftNode
		selfNode.right = rightNode
	}

	// nothing to return; node map was modified in place
}

func parseNodes(input string) map[string]*node {
	lns := strings.Split(input, "\n")
	nodeMap := make(map[string]*node)
	for _, ln := range lns {
		parseNode(ln, nodeMap)
	}
	return nodeMap
}

func parseInput(input string) (directions []string, nodeMap map[string]*node) {
	dirsAndNodes, err := utils.SplitIntoExpectedParts(input, "\n\n", 2)
	if err != nil {
		utils.LogfErrorAndExit(err, "splitting input")
	}
	return strings.Split(strings.TrimSpace(dirsAndNodes[0]), ""), parseNodes(strings.TrimSpace(dirsAndNodes[1]))
}
