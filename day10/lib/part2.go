package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day10/inputs/input1.txt")

	X := 1
	cycle := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		increment := 0
		commands := strings.Split(line, " ")

		// Account for additional cycle taken in addx
		if commands[0] == "addx" {
			cycle = runCycle(cycle, X)
			increment, _ = strconv.Atoi(commands[1])
		}
		cycle = runCycle(cycle, X)

		X += increment
	}
}

func runCycle(cycle int, X int) int {
	cycle++
	printPixel(cycle, X)
	if (cycle)%40 == 0 {
		cycle = 0
	}
	return cycle
}

func printPixel(cycle int, X int) {
	if (cycle-1)-X == 1 || X-(cycle-1) == 1 || X == (cycle-1) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if (cycle)%40 == 0 {
		fmt.Println()
	}
}
