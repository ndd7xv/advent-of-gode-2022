package main

import (
	"day13/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day13.Part1()
	}

	if os.Args[1] == "2" {
		day13.Part2()
	}
}
