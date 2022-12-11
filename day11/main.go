package main

import (
	day11 "day11/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day11.Part1()
	}

	if os.Args[1] == "2" {
		day11.Part2()
	}
}
