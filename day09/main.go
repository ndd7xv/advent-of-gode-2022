package main

import (
	day09 "day09/lib"
	"os"
)

func main() {
	// Apparently if dX or dY here was > 1 then the tail code just move to the previous position of
	// the head
	if os.Args[1] == "1" {
		day09.Part1()
	}

	if os.Args[1] == "2" {
		day09.Part2()
	}
}
