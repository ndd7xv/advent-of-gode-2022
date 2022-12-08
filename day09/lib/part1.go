package day09

import (
	"fmt"
	"os"
	"strings"
)

func Part1() {
	input, _ := os.ReadFile("day09/inputs/input1.txt")

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		
	}
	fmt.Println("Part 1, ready to go!")
}
