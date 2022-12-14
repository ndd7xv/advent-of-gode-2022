package day14

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

type Wall struct {
	coordinates []Coordinate
}

func NewCoordinate(s string) Coordinate {
	unparsed := strings.Split(s, ",")
	x, _ := strconv.Atoi(unparsed[0])
	y, _ := strconv.Atoi(unparsed[1])
	return Coordinate{
		x,
		y,
	}
}

func Part1() {
	input, _ := os.ReadFile("day14/inputs/input1.txt")

	walls := []Wall{}
	minX := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt

	fmt.Printf("%s\n", string(input))

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		traces := strings.Split(line, " -> ")
		coordinates := []Coordinate{}

		for _, point := range traces {
			c := NewCoordinate(point)
			coordinates = append(coordinates, c)
			if c.x < minX {
				minX = c.x
			}
			if c.x > maxX {
				maxX = c.x
			}
			if c.y > maxY {
				maxY = c.y
			}
		}

		walls = append(walls, Wall{coordinates})
	}

	grid := [][]int{}
	for i := 0; i <= maxY; i++ {
		grid = append(grid, []int{})
		for j := 0; j <= maxX-minX; j++ {
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
					grid[int(y)][int(float64(current.x)-float64(minX))] = 1
				}
			} else {
				min := math.Min(float64(current.x), float64(prev.x))
				max := math.Max(float64(current.x), float64(prev.x))
				for x := min; x <= max; x++ {
					grid[current.y][int(x-float64(minX))] = 1
				}
			}
		}
	}

	for r := range grid {
		for c := range grid[0] {
			fmt.Printf("%d", grid[r][c])
		}
		fmt.Println()
	}

	count := 0
	for simulateSand(grid, minX, maxY) {
		count++
	}
	fmt.Println(count)
}

func simulateSand(grid [][]int, minX int, maxY int) bool {
	sand := Coordinate{500 - minX, 0}
	for {
		if sand.y+1 == len(grid) || sand.x-1 < 0 || sand.x+1 == len(grid[0]) {
			return false
		} else if (grid)[sand.y+1][sand.x] == 0 {
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

	return true
}
