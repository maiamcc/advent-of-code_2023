package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"strconv"
	"testing"
)

func main() {
	inputLines := utils.MustReadFileAsLines("day1/input.txt")

	total := 0
	for i, ln := range inputLines {
		numStr := fmt.Sprintf("%s%s", ln[0], ln[len(ln)-1])
		num, err := strconv.Atoi(numStr)
		if err != nil {
			utils.LogfErrorAndExit(err, "converting ln at index %d ('%s')", i, ln)
		}
		total += num
	}
	fmt.Printf("Total: %d", total)
}

// getNumericChars returns an array of all the numeric characters in given string s
// (still stored as strings))
func getNumericChars(s string) []string {

}

func TestGetNumericChars(t *testing.T) {
	cases := map[string][]string{ // map input to expected output
		"1hello2":      {"1", "2"},
		"hel1lo2world": {"1", "2"},
		"321stuff":     {"3", "2", "1"},
	}

	for k, v := range cases {
		actual := getNumericChars(k)
		assert.Equal()
	}
}
