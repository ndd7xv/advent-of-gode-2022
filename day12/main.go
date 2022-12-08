package main

import (
	"day12/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day12.Part1()
	}

	if os.Args[1] == "2" {
		day12.Part2()
	}
}
