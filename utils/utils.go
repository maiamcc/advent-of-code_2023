package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LogfErrorAndExit(err error, msg string, a ...interface{}) {
	msgFormatted := fmt.Sprintf(msg, a...)
	if err != nil {
		fmt.Printf("%s\n\t%v", msgFormatted, err)
	} else {
		fmt.Printf("%s\n\t", msgFormatted)
	}
	os.Exit(1)
}

func MustReadFileAsString(name string) string {
	bytes, err := os.ReadFile(name)
	if err != nil {
		LogfErrorAndExit(err, "Error opening file %s", name)
	}
	return string(bytes)
}

func MustReadFileAsLines(name string) []string {
	s := MustReadFileAsString(name)
	return strings.Split(s, "\n")
}

func MustAtoI(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		LogfErrorAndExit(err, "converting string to int")
	}
	return i
}

func MustRune(s string) rune {
	runes := []rune(s)
	if len(runes) != 1 {
		fmt.Printf("string %s should have contained a single rune but was multiple", s)
	}
	return runes[0]
}

func SplitIntoExpectedParts(s string, sep string, expectedParts int) ([]string, error) {
	parts := strings.Split(s, sep)
	if len(parts) != expectedParts {
		return nil, fmt.Errorf("expected %d parts when splitting string '%s' on sep '%s' but got %d: %+v",
			expectedParts, s, sep, len(parts), parts)
	}
	return parts, nil
}

func NumResultFromRe(s string, re *regexp.Regexp) int {
	res := re.FindStringSubmatch(s)
	if len(res) == 0 {
		// no match
		return 0
	}
	return MustAtoI(res[1])
}

func StringsToInts(s []string) ([]int, error) {
	var res []int
	for _, num := range s {
		// account for wonky spaces / empty
		num = strings.TrimSpace(num)
		if num == "" {
			continue
		}
		i, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

func MustStringsToInts(s []string) []int {
	ints, err := StringsToInts(s)
	if err != nil {
		LogfErrorAndExit(err, "converting []string to []int")
	}
	return ints
}

func Rng(start int, end int) []int {
	if end < start {
		fmt.Printf("the range function doesn't work like that")
		os.Exit(1)
	}
	// behaves like the python range function: inclusive of `start`, exclusive of `end`
	a := make([]int, end-start)
	for i := range a {
		a[i] = start + i
	}
	return a
}
