package day03

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part2() {
	input, _ := ioutil.ReadFile("inputs/input1.txt")
	score := 0

	var elves []string

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		elves = append(elves, line)
		if len(elves) == 3 {
			shared := findShared(elves)
			score += getPriority(shared)
			elves = []string{}
		}
	}

	fmt.Println(score)
}

func findShared(arr []string) rune {

	m := make(map[rune]int)

	for _, item := range arr {
		seen := make(map[rune]bool)
		for _, ch := range item {
			if _, ok := m[ch]; ok {
				if _, ok := seen[ch]; !ok {
					m[ch]++
					if m[ch] == 3 {
						return ch
					}
					seen[ch] = true
				}
			} else {
				m[ch] = 1
				seen[ch] = true
			}
		}
	}

	panic("Should've found shared item!")
}
