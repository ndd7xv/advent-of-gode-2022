package main

import (
	"day04/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day04.Part1()
	}

	if os.Args[1] == "2" {
		day04.Part2()
	}
}
