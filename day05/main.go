package main

import (
	"day05/lib"
	"os"
)

// In retrospect it might've been better to just manipulate the input as strings instead of
// importing a stack package, but I learned how to do that so :^)
func main() {
	if os.Args[1] == "1" {
		day05.Part1()
	}

	if os.Args[1] == "2" {
		day05.Part2()
	}
}
