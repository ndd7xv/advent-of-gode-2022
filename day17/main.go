package main

import (
	"day17/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day17.Part1()
	}

	if os.Args[1] == "2" {
		day17.Part2()
	}
}
