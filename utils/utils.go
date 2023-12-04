package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LogfErrorAndExit(err error, msg string, a ...interface{}) {
	msgFormatted := fmt.Sprintf(msg, a...)
	fmt.Printf("%s\n\t%v", msgFormatted, err)
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
