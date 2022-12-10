package day09

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day09/inputs/input1.txt")

	rope := make([]Coordinate, 10)
	for i := range rope {
		rope[i] = Coordinate{0, 0}
	}

	tailLocations := make(map[Coordinate]bool)
	tailLocations[Coordinate{0, 0}] = true

	directionUpdates := make(map[string]Coordinate)
	directionUpdates["U"] = Coordinate{0, 1}
	directionUpdates["D"] = Coordinate{0, -1}
	directionUpdates["L"] = Coordinate{-1, 0}
	directionUpdates["R"] = Coordinate{1, 0}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		words := strings.Split(line, " ")
		direction := words[0]
		steps, _ := strconv.Atoi(words[1])

		for i := 0; i < steps; i++ {
			dX, dY := directionUpdates[direction].x, directionUpdates[direction].y
			rope[0] = Coordinate{rope[0].x + dX, rope[0].y + dY}
			readjustRope(rope)
			tailLocations[rope[len(rope)-1]] = true
		}
	}
	fmt.Println(len(tailLocations))
}

func readjustRope(rope []Coordinate) {
	for j := 1; j < len(rope); j++ {
		rope[j-1], rope[j] = ReadjustTail(rope[j-1], rope[j])
	}
}
