package main

import (
	"day06/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day06.Part1()
	}

	if os.Args[1] == "2" {
		day06.Part2()
	}
}
