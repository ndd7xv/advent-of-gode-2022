package day09

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func Part1() {
	input, _ := os.ReadFile("day09/inputs/input1.txt")

	tailLocations := make(map[Coordinate]bool)
	head := Coordinate{0, 0}
	tail := Coordinate{0, 0}

	tailLocations[tail] = true

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
				head = Coordinate{head.x, head.y + 1}
				head, tail = ReadjustTail(head, tail)
				tailLocations[tail] = true
			}
		case "D":
			for i := 0; i < steps; i++ {
				head = Coordinate{head.x, head.y - 1}
				head, tail = ReadjustTail(head, tail)
				tailLocations[tail] = true
			}
		case "L":
			for i := 0; i < steps; i++ {
				head = Coordinate{head.x - 1, head.y}
				head, tail = ReadjustTail(head, tail)
				tailLocations[tail] = true
			}
		case "R":
			for i := 0; i < steps; i++ {
				head = Coordinate{head.x + 1, head.y}
				head, tail = ReadjustTail(head, tail)
				tailLocations[tail] = true
			}
		}
	}
	fmt.Println(len(tailLocations))
}

func ReadjustTail(head, tail Coordinate) (
	Coordinate,
	Coordinate,
) {

	fmt.Printf("Head %d %d | Tail %d %d... to", head.x, head.y, tail.x, tail.y)

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
	fmt.Printf(" Head %d %d | Tail %d %d\n", head.x, head.y, tail.x, tail.y)

	return head, tail
}
