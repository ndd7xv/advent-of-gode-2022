package day18

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
}

func NewCube(x, y, z int) Cube {
	return Cube{x, y, z}
}

func Part1() {
	input, _ := os.ReadFile("day18/inputs/input1.txt")

	cubes := []Cube{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		dimensions := strings.Split(line, ",")
		x, _ := strconv.Atoi(dimensions[0])
		y, _ := strconv.Atoi(dimensions[1])
		z, _ := strconv.Atoi(dimensions[2])

		cubes = append(cubes, NewCube(x, y, z))
	}

	totalArea := len(cubes) * 6

	for i := 0; i < len(cubes)-1; i++ {
		for j := i + 1; j < len(cubes); j++ {
			if SharesSide(cubes[i], cubes[j]) {
				totalArea -= 2
			}
		}
	}
	fmt.Println(totalArea)
}

func SharesSide(a, b Cube) bool {
	if a.x == b.x {
		if a.y == b.y {
			if math.Abs(float64(a.z-b.z)) == 1 {
				return true
			}
		} else if a.z == b.z {
			if math.Abs(float64(a.y-b.y)) == 1 {
				return true
			}
		}
	} else if a.y == b.y {
		if a.z == b.z {
			if math.Abs(float64(a.x-b.x)) == 1 {
				return true
			}
		}
	}
	return false
}
