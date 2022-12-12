package day12

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day12/inputs/input1.txt")

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	grid := make([][]int, len(lines))

	var endX, endY int

	contenders := make(map[Coordinate]bool)

	for r, line := range lines {
		grid[r] = make([]int, len(line))
		for c, ch := range line {
			grid[r][c] = int(ch) - 'a'
			if ch == 'S' {
				grid[r][c] = 0
			}
			if ch == 'E' {
				endX, endY = r, c
				grid[r][c] = 'z' - 'a' + 1
				// fmt.Printf("Coordinate of E is %d, %d\n", endX, endY)
			}
			if grid[r][c] == 0 {
				contenders[Coordinate{r, c}] = true
			}
		}
	}

	answer := math.MaxFloat64
	for k := range contenders {
		// fmt.Printf("Comparing %v\n", k)
		shortestPath := float64(findShortestPath(grid, k.x, k.y)[endX][endY])
		if shortestPath != -1 {
			answer = math.Min(answer, shortestPath)
		}
	}

	fmt.Printf("%f\n", answer)
}
