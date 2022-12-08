package main

import (
	"day24/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day24.Part1()
	}

	if os.Args[1] == "2" {
		day24.Part2()
	}
}
