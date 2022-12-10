package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	input, _ := os.ReadFile("day10/inputs/input1.txt")

	haveReached := make(map[int]bool)
	thresholds := []int{220, 180, 140, 100, 60, 20}

	X := 1
	cycle := 0
	sum := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		increment := 0
		commands := strings.Split(line, " ")
		if commands[0] == "noop" {
			cycle++
		} else {
			increment, _ = strconv.Atoi(commands[1])
			cycle += 2
		}

		for _, check := range thresholds {
			if _, ok := haveReached[check]; cycle >= check && !ok {
				sum += (X * check)
				haveReached[check] = true
				break
			}
		}

		X += increment
	}

	fmt.Println(sum)
}
