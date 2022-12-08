package main

import (
	"day18/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day18.Part1()
	}

	if os.Args[1] == "2" {
		day18.Part2()
	}
}
