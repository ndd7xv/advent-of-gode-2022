package main

import (
	"day09/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day09.Part1()
	}

	if os.Args[1] == "2" {
		day09.Part2()
	}
}
