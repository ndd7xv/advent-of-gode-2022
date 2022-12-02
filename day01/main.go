package main

import (
	"day01/lib"
	"os"
)

// Really liked the way this was solved here: https://github.com/mnml/aoc/blob/main/2022/01/1.go
func main() {
	if os.Args[1] == "1" {
		day01.Part1()
	}

	if os.Args[1] == "2" {
		day01.Part2()
	}
}
