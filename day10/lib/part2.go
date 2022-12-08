package day10

import (
	"fmt"
	"os"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day10/inputs/input1.txt")

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
	}

	fmt.Println("Part 2, ready to go!")
}
