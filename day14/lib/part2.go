package day14

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day14/inputs/input1.txt")

	walls := []Wall{}
	maxY := math.MinInt

	minXCoord := math.MaxInt

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		traces := strings.Split(line, " -> ")
		coordinates := []Coordinate{}

		for _, point := range traces {
			c := NewCoordinate(point)
			coordinates = append(coordinates, c)
			if c.x < minXCoord {
				minXCoord = c.x
			}
			if c.y > maxY {
				maxY = c.y
			}
		}

		walls = append(walls, Wall{coordinates})
	}

	maxY += 2

	grid := [][]int{}
	for i := 0; i <= maxY; i++ {
		grid = append(grid, []int{})
		for j := 0; j <= 1000; j++ {
			grid[i] = append(grid[i], 0)
		}
	}

	for _, wall := range walls {
		for i := 1; i < len(wall.coordinates); i++ {
			current := wall.coordinates[i]
			prev := wall.coordinates[i-1]

			if current.x == prev.x {
				min := math.Min(float64(current.y), float64(prev.y))
				max := math.Max(float64(current.y), float64(prev.y))
				for y := min; y <= max; y++ {
					grid[int(y)][current.x] = 1
				}
			} else {
				min := math.Min(float64(current.x), float64(prev.x))
				max := math.Max(float64(current.x), float64(prev.x))
				for x := min; x <= max; x++ {
					grid[current.y][int(x)] = 1
				}
			}
		}
	}

	for i := 0; i <= 1000; i++ {
		grid[maxY][i] = 1
	}

	fmt.Println(simulateInfiniteSand(grid, maxY))
}

func simulateInfiniteSand(grid [][]int, maxY int) int {
	count := 0
	for {
		if grid[0][500] == 1 {
			return count
		}
		sand := Coordinate{500, 0}
		for {
			if (grid)[sand.y+1][sand.x] == 0 {
				sand.y++
			} else if (grid)[sand.y+1][sand.x-1] == 0 {
				sand.x--
				sand.y++
			} else if (grid)[sand.y+1][sand.x+1] == 0 {
				sand.x++
				sand.y++
			} else {
				(grid)[sand.y][sand.x] = 1
				break
			}
		}
		count++
	}
}
