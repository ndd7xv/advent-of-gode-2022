package day02

import (
	"fmt"
	"os"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("inputs/input1.txt")

	score := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		opponent := choice(line[0])
		you := result(line[2])

		score += you

		switch you {
		case 0:
			switch opponent {
			case "rock":
				score += 3
			case "paper":
				score++
			case "scissors":
				score += 2
			}
		case 3:
			switch opponent {
			case "rock":
				score++
			case "paper":
				score += 2
			case "scissors":
				score += 3
			}
		case 6:
			switch opponent {
			case "rock":
				score += 2
			case "paper":
				score += 3
			case "scissors":
				score++
			}
		}
	}
	fmt.Println(score)
}

func result(c byte) int {
	switch c {
	case 'X':
		return 0
	case 'Y':
		return 3
	case 'Z':
		return 6
	}
	panic("invalid input!")
}
