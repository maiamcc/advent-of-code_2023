package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
)

func main() {
	inputLines := utils.MustReadFileAsLines("day9/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	total := 0
	for _, ln := range inputLines {
		ints := utils.MustSplitStringToInts(ln, " ")
		extrapolated := extrapolateRightElem(ints)
		total += extrapolated
	}
	return total
}

func partTwo(inputLines []string) int {
	total := 0
	for _, ln := range inputLines {
		ints := utils.MustSplitStringToInts(ln, " ")
		extrapolated := extrapolateLeftElem(ints)
		total += extrapolated
	}
	return total
}

func diffList(input []int) []int {
	var res []int
	for i := 1; i < len(input); i++ { // index of the second elem of a pair
		res = append(res, input[i]-input[i-1])
	}
	return res
}

func allDiffLists(input []int) [][]int {
	var diffLists [][]int
	curArr := input
	for !isZerosArray(curArr) {
		curArr = diffList(curArr)
		diffLists = append(diffLists, curArr)
	}
	return diffLists
}

func diffListsAndOrig(input []int) [][]int {
	return append([][]int{input}, allDiffLists(input)...)
}

func isZerosArray(arr []int) bool {
	// an empty list will also return true here, which is not strictly accurate
	// but breaks the loop above so whatever that's fine.
	for _, elem := range arr {
		if elem != 0 {
			return false
		}
	}
	return true
}

func extrapolateRightElem(input []int) int {
	diffLists := diffListsAndOrig(input)
	if !isZerosArray(utils.LastElem(diffLists)) {
		utils.LogfAndExit("last difflist was not a zero array!\n\tgot: %v", utils.LastElem(diffLists))
	}
	extrapolated := 0                          // we know that the last diff list is all zeros
	for i := len(diffLists) - 2; i >= 0; i-- { // loop indexes backwards (not counting last list, which is all zeros)
		extrapolated = utils.LastElem(diffLists[i]) + extrapolated
	}
	return extrapolated
}

func extrapolateLeftElem(input []int) int {
	diffLists := diffListsAndOrig(input)
	if !isZerosArray(utils.LastElem(diffLists)) {
		utils.LogfAndExit("last difflist was not a zero array!\n\tgot: %v", utils.LastElem(diffLists))
	}
	extrapolated := 0                          // we know that the last diff list is all zeros
	for i := len(diffLists) - 2; i >= 0; i-- { // loop indexes backwards (not counting last list, which is all zeros)
		extrapolated = diffLists[i][0] - extrapolated
	}
	return extrapolated
}
