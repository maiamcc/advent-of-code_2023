package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"os"
	"regexp"
	"strings"
)

var nodeRe = regexp.MustCompile("([A-Z0-9]{3}) = \\(([A-Z0-9]{3}), ([A-Z0-9]{3})\\)")

func main() {
	fullInput := utils.MustReadFileAsString("day8/input.txt")
	fmt.Println("The answer to Part One is:", partOne(fullInput))
	fmt.Println("The answer to Part Two is:", partTwo(fullInput))
}

func partOne(fullInput string) int {
	dirs, nm := parseInput(fullInput)

	count, _ := stepsToEndNodes(dirs, nm, "AAA", func(n *node) bool {
		return n.label == "ZZZ"
	})
	return count
}

func partTwo(fullInput string) int {
	dirs, nodes := parseInput(fullInput)

	startNodes := nodes.findStartNodes()
	var stepCounts []int
	for _, startNode := range startNodes {
		numSteps, _ := stepsToEndNodes(dirs, nodes, startNode.label, func(n *node) bool {
			return n.isTerminalNode()
		})
		// Look, this is janky but we can use the lcm of the number of steps because for
		// every nodeA -> nodeZ in n steps, nodeZ -> nodeZ cycles in n steps as well
		// (which isn't a constraint in the problem but is a constraint I found in my
		// input, grumble grumble
		stepCounts = append(stepCounts, numSteps)
	}
	lcm := LCM(stepCounts[0], stepCounts[1], stepCounts[2:]...)
	return lcm
}

func stepsToEndNodes(dirs []string, nodes nodeMap, startNodeLabel string, endCond func(n *node) bool) (int, *node) {
	curNode, ok := nodes[startNodeLabel]
	if !ok {
		utils.LogfAndExit("couldn't find start node %s", startNodeLabel)
	}

	count := 0
	for {
		for _, dir := range dirs {
			curNode = curNode.step(dir)
			count += 1
			if endCond(curNode) {
				// win condition
				return count, curNode
			}
		}
	}
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

func (n *node) isTerminalNode() bool {
	// is this too expensive? can cache it.
	return string(n.label[len(n.label)-1]) == "Z"
}

func (n *node) isStartNode() bool {
	// is this too expensive? can cache it.
	return string(n.label[len(n.label)-1]) == "A"
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

func parseInput(input string) (directions []string, nm nodeMap) {
	dirsAndNodes, err := utils.SplitIntoExpectedParts(input, "\n\n", 2)
	if err != nil {
		utils.LogfErrorAndExit(err, "splitting input")
	}
	return strings.Split(strings.TrimSpace(dirsAndNodes[0]), ""), parseNodes(strings.TrimSpace(dirsAndNodes[1]))
}

// blurgh nodeMap should be its own type alias

type nodeMap map[string]*node

func (nm nodeMap) findStartNodes() []*node {
	var res []*node
	for _, n := range nm {
		if n.isStartNode() {
			res = append(res, n)
		}
	}
	return res
}

func (nm nodeMap) findEndNodes() []*node {
	var res []*node
	for _, n := range nm {
		if n.isTerminalNode() {
			res = append(res, n)
		}
	}
	return res
}

// Shamelessly stolen from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
