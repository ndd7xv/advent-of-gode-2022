package main

import (
	day08 "day08/lib"
	"os"
)

func main() {
	if os.Args[1] == "1" {
		day08.Part1()
	}

	// not super impressed with what I see on Reddit but I learned I can return things in a tuple
	// like fashion:
	// https://gist.github.com/foxwhite25/4aa4dff56631555a2f929ba927fe57ae#file-solution-go-L42-L47
	if os.Args[1] == "2" {
		day08.Part2()
	}
}
