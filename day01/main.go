package main

import (
	"day01/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day01.Part1()
	}

	if os.Args[1] == "2" {
		day01.Part2()
	}
}
