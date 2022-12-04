package main

import (
	"day04/lib"
	"os"
)

// None of the alternatives I've seen in my quick peruse have really impressed me, but they do
// seem better than my spaghetti code. I've seen uses of modulus that are interesting too.
//
// https://github.com/tudorpavel/advent-2022/blob/master/day04/main.go
// https://github.com/mnml/aoc/blob/main/2022/02/1.go
func main() {
	if os.Args[1] == "1" {
		day04.Part1()
	}

	if os.Args[1] == "2" {
		day04.Part2()
	}
}
