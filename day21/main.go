package main

import (
	"day21/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day21.Part1()
	}

	if os.Args[1] == "2" {
		day21.Part2()
	}
}
