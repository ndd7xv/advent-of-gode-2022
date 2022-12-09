package day09

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day09/inputs/input1.txt")

	tailLocations := make(map[Coordinate]bool)
	rope := make([]Coordinate, 10)
	for i := range rope {
		rope[i] = Coordinate{0, 0}
	}
	tailLocations[Coordinate{0, 0}] = true

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		words := strings.Split(line, " ")
		direction := words[0]
		steps, _ := strconv.Atoi(words[1])

		fmt.Printf("%s %s\n", direction, words[1])
		switch direction {
		case "U":
			for i := 0; i < steps; i++ {
				rope[0] = Coordinate{rope[0].x, rope[0].y + 1}
				for j := 1; j < len(rope); j++ {
					rope[j-1], rope[j] = ReadjustTail(rope[j-1], rope[j])
				}
				tailLocations[rope[len(rope)-1]] = true
			}
		case "D":
			for i := 0; i < steps; i++ {
				rope[0] = Coordinate{rope[0].x, rope[0].y - 1}
				for j := 1; j < len(rope); j++ {
					rope[j-1], rope[j] = ReadjustTail(rope[j-1], rope[j])
				}
				tailLocations[rope[len(rope)-1]] = true
			}
		case "L":
			for i := 0; i < steps; i++ {
				rope[0] = Coordinate{rope[0].x - 1, rope[0].y}
				for j := 1; j < len(rope); j++ {
					rope[j-1], rope[j] = ReadjustTail(rope[j-1], rope[j])
				}
				tailLocations[rope[len(rope)-1]] = true
			}
		case "R":
			for i := 0; i < steps; i++ {
				rope[0] = Coordinate{rope[0].x + 1, rope[0].y}
				for j := 1; j < len(rope); j++ {
					rope[j-1], rope[j] = ReadjustTail(rope[j-1], rope[j])
				}
				tailLocations[rope[len(rope)-1]] = true
			}
		}
	}
	fmt.Println(len(tailLocations))
}
