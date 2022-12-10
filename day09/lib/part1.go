package day09

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Coordinate represents a point where a knot in the rope could be
type Coordinate struct {
	x int
	y int
}

func Part1() {
	input, _ := os.ReadFile("day09/inputs/input1.txt")

	head := Coordinate{0, 0}
	tail := Coordinate{0, 0}

	// A set of locations the tail has been
	tailLocations := make(map[Coordinate]bool)
	tailLocations[tail] = true

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
			head = Coordinate{head.x + dX, head.y + dY}
			head, tail = ReadjustTail(head, tail)
			tailLocations[tail] = true
		}
	}
	fmt.Println(len(tailLocations))
}

// ReadjustTail checks two adjacent knot coordinates to determine if the successive knot should be
// moved to be adjacent with the knot before it.
func ReadjustTail(head, tail Coordinate) (
	Coordinate,
	Coordinate,
) {
	dX := head.x - tail.x
	dY := head.y - tail.y
	if (head.x != tail.x && head.y != tail.y) && ((-2 >= dX || dX >= 2) || (-2 >= dY || dY >= 2)) {
		if dX > 0 {
			tail.x++
		} else {
			tail.x--
		}
		if dY > 0 {
			tail.y++
		} else {
			tail.y--
		}
	} else if -2 >= dX || dX >= 2 {
		if dX > 0 {
			tail.x++
		} else {
			tail.x--
		}
	} else if -2 >= dY || dY >= 2 {
		if dY > 0 {
			tail.y++
		} else {
			tail.y--
		}
	}

	return head, tail
}
