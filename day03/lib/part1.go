package day03

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func Part1() {
	input, _ := ioutil.ReadFile("inputs/input1.txt")

	score := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		m := make(map[rune]bool)

		first := line[:(len(line) / 2)]
		second := line[(len(line) / 2):]

		for _, ch := range first {
			m[ch] = true
		}

		for _, ch := range second {
			if m[ch] {
				score += getPriority(ch)
				break
			}
		}
	}

	fmt.Println(score)
}

func getPriority(b rune) int {
	if unicode.IsLower(b) {
		return int(b) - 96
	}
	return int(b) - 38
}
