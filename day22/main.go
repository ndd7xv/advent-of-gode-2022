package main

import (
	"day22/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day22.Part1()
	}

	if os.Args[1] == "2" {
		day22.Part2()
	}
}
