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
		if commands[0] == "noop" {
			cycle++
		} else {
			parse, _ := strconv.Atoi(commands[1])
			increment = parse
			cycle++
			//fmt.Printf("\nDrawing Cycle %d where CRT is %d\n", cycle, X)
			if (cycle-1)-X == 1 || X-(cycle-1) == 1 || X == (cycle-1) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if (cycle)%40 == 0 {
				cycle = 0
				fmt.Println()
			}
			cycle++
		}

		//fmt.Printf("\nDrawing Cycle %d where CRT is %d\n", cycle, X)
		if (cycle-1)-X == 1 || X-(cycle-1) == 1 || X == (cycle-1) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if (cycle)%40 == 0 {
			cycle = 0
			fmt.Println()
		}

		X += increment
	}
}
