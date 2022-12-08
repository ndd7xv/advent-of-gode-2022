package main

import (
	"day23/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day23.Part1()
	}

	if os.Args[1] == "2" {
		day23.Part2()
	}
}
