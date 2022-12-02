package day02

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part1() {
	input, _ := ioutil.ReadFile("inputs/input1.txt")

	score := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		opponent := choice(line[0])
		you := choice(line[2])

		switch you {
		case "rock":
			switch opponent {
			case "rock":
				score += 3
			case "scissors":
				score += 6
			}
			score++
		case "paper":
			switch opponent {
			case "paper":
				score += 3
			case "rock":
				score += 6
			}
			score += 2
		case "scissors":
			switch opponent {
			case "scissors":
				score += 3
			case "paper":
				score += 6
			}
			score += 3
		}
	}
	fmt.Println(score)
}

func choice(c byte) string {
	switch c {
	case 'A', 'X':
		return "rock"
	case 'B', 'Y':
		return "paper"
	case 'C', 'Z':
		return "scissors"
	}
	panic("invalid input!")
}
