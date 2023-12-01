package utils

import (
	"fmt"
	"os"
	"strings"
)

func LogfErrorAndExit(err error, msg string, a ...any) {
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
