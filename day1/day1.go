package main

import (
	"errors"
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"regexp"
	"strconv"
	"unicode"
)

var numberStringsToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var numberRe = regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[1-9]")

func main() {
	inputLines := utils.MustReadFileAsLines("day1/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	total := 0
	for _, ln := range inputLines {
		numChars := getNumericChars(ln)
		numStr := fmt.Sprintf("%s%s", string(numChars[0]), string(numChars[len(numChars)-1]))
		num := utils.MustAtoI(numStr)
		total += num
	}
	return total
}

func partTwo(inputLines []string) int {
	total := 0
	for i, ln := range inputLines {
		nums, err := getAllNumbers(ln)
		if err != nil {
			utils.LogfErrorAndExit(err, "parsing line %d ('%s')", i, ln)
		}
		numStr := fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1])
		num, err := strconv.Atoi(numStr)
		if err != nil {
			utils.LogfErrorAndExit(err, "converting ln at index %d ('%s')", i, ln)
		}
		fmt.Printf("%s --> %v --> %d\n", ln, nums, num)
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

// numStrToInt converts a string representation of a digit (e.g. either "1" or "one") to an int
func numStrToInt(numStr string) (int, error) {
	if len(numStr) == 1 {
		// a single-character result is (hopefully!) just a digit
		return strconv.Atoi(numStr)
	} else {
		num, ok := numberStringsToInt[numStr]
		if !ok {
			return -1, errors.New(fmt.Sprintf("couldn't find int corresponding to string '%s'", numStr))
		}
		return num, nil
	}
}

// consumeFirstNumber finds the first representation of a number in the given
// string (may be numeric or spelled out, e.g. "1" or "one"), pops it from the
// string, and returns it as an int, along with the remaining string to the right
// of the popped num. Returns -1 to indicate that there are no more valid numbers
// in the string (or the string is exhausted)
func consumeFirstNumber(input string) (num int, rest string, err error) {
	idx := numberRe.FindStringIndex(input)
	if len(idx) == 0 {
		// no number representation found in this string
		return -1, input, nil
	}
	numStr := input[idx[0]:idx[1]]
	if len(numStr) == 1 {
		// a single-character result is (hopefully!) just a digit
		num = utils.MustAtoI(numStr)
	} else {
		var ok bool
		num, ok = numberStringsToInt[numStr]
		if !ok {
			return -1, "",
				errors.New(fmt.Sprintf("couldn't find int corresponding to string '%s'", numStr))
		}
	}
	return num, input[idx[1]:], nil
}

// getAllNumbers finds all representations of numbers (digits or strings) in order
// in the given string
func getAllNumbers(input string) ([]int, error) {
	res := []int{}
	num, rest, err := consumeFirstNumber(input)

	for num != -1 && err == nil {
		res = append(res, num)
		num, rest, err = consumeFirstNumber(rest)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
