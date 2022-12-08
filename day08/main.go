package main

import (
	"day08/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day08.Part1()
	}

	if os.Args[1] == "2" {
		day08.Part2()
	}
}
