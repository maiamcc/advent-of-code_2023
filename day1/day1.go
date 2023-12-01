package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
)

func main() {
	inputLines := utils.MustReadFileAsLines("day1/input.txt")
	fmt.Println(inputLines[:5])
	fmt.Println("num lines:", len(inputLines))
}
