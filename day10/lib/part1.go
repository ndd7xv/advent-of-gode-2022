package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	input, _ := os.ReadFile("day10/inputs/test1.txt")

	X := 1
	cycle := 0
	triggered := make(map[int]bool)
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
			parse, _ := strconv.Atoi(commands[1])
			increment = parse
			cycle += 2
		}

		if _, ok := triggered[220]; cycle >= 220 && !ok {
			fmt.Println(X)
			sum += (X * 220)
			triggered[220] = true
		} else if _, ok := triggered[180]; cycle >= 180 && !ok {
			fmt.Println(X)
			sum += (X * 180)
			triggered[180] = true
		} else if _, ok := triggered[140]; cycle >= 140 && !ok {
			fmt.Println(X)
			sum += (X * 140)
			triggered[140] = true
		} else if _, ok := triggered[100]; cycle >= 100 && !ok {
			fmt.Println(X)
			sum += (X * 100)
			triggered[100] = true
		} else if _, ok := triggered[60]; cycle >= 60 && !ok {
			fmt.Println(X)
			sum += (X * 60)
			triggered[60] = true
		} else if _, ok := triggered[20]; cycle >= 20 && !ok {
			fmt.Println(X)
			sum += (X * 20)
			triggered[20] = true
		}

		X += increment
	}

	fmt.Println(sum)
}
