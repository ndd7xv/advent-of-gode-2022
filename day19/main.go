package main

import (
	"day19/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day19.Part1()
	}

	if os.Args[1] == "2" {
		day19.Part2()
	}
}
