package main

import (
	"day14/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day14.Part1()
	}

	if os.Args[1] == "2" {
		day14.Part2()
	}
}
