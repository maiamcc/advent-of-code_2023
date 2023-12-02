package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"strconv"
	"unicode"
)

func main() {
	inputLines := utils.MustReadFileAsLines("day1/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	total := 0
	for i, ln := range inputLines {
		numChars := getNumericChars(ln)
		numStr := fmt.Sprintf("%s%s", string(numChars[0]), string(numChars[len(numChars)-1]))
		num, err := strconv.Atoi(numStr)
		if err != nil {
			utils.LogfErrorAndExit(err, "converting ln at index %d ('%s')", i, ln)
		}
		total += num
	}
	return total
}

// getNumericChars returns an array of all the numeric characters in given string s
// (still stored as strings))
func getNumericChars(s string) []string {
	res := []string{}
	for _, ch := range s {
		if unicode.IsNumber(ch) {
			res = append(res, string(ch))
		}
	}
	return res
}

func partTwo(inputLines []string) int {
	//total := 0
	return len(inputLines)
}
