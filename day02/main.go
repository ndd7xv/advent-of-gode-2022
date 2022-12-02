package main

import (
	"day02/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day02.Part1()
	}

	if os.Args[1] == "2" {
		day02.Part2()
	}
}
