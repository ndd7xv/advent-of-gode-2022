package main

import (
	"day03/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day03.Part1()
	}

	if os.Args[1] == "2" {
		day03.Part2()
	}
}
