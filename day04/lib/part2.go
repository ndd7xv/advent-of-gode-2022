package day04

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := ioutil.ReadFile("inputs/test.txt")

	score := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		pairs := strings.Split(line, ",")
		firstRange := strings.Split(pairs[0], "-")
		secondRange := strings.Split(pairs[1], "-")

		firstStart, _ := strconv.Atoi(firstRange[0])
		firstEnd, _ := strconv.Atoi(firstRange[1])
		secondStart, _ := strconv.Atoi(secondRange[0])
		secondEnd, _ := strconv.Atoi(secondRange[1])

		if firstStart <= secondStart && secondStart <= firstEnd {
			score++
		} else if secondStart <= firstStart && firstStart <= secondEnd {
			score++
		}
	}

	fmt.Println(score)
}
