package main

import (
	"day16/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day16.Part1()
	}

	if os.Args[1] == "2" {
		day16.Part2()
	}
}
