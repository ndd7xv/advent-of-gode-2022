package main

import (
	"day10/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day10.Part1()
	}

	if os.Args[1] == "2" {
		day10.Part2()
	}
}
