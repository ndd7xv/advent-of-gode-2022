package main

import (
	day07 "day07/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day07.Part1()
	}

	if os.Args[1] == "2" {
		day07.Part2()
	}
}
