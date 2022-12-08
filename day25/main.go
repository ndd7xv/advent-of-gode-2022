package main

import (
	"day25/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day25.Part1()
	}

	if os.Args[1] == "2" {
		day25.Part2()
	}
}
