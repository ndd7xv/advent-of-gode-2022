package main

import (
	"day15/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day15.Part1()
	}

	if os.Args[1] == "2" {
		day15.Part2()
	}
}
