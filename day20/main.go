package main

import (
	"day20/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day20.Part1()
	}

	if os.Args[1] == "2" {
		day20.Part2()
	}
}
