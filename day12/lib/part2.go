package day12

import (
	"fmt"
	"os"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day12/inputs/input.txt")

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	grid := make([][]int, len(lines))

	end := Coordinate{-1, -1}

	for r, line := range lines {
		grid[r] = make([]int, len(line))
		for c, ch := range line {
			grid[r][c] = int(ch) - 'a'
			if ch == 'S' {
				grid[r][c] = 0
			}
			if ch == 'E' {
				end = Coordinate{r, c}
				grid[r][c] = 26 // 'z' - 'a' + 1
			}
		}
	}

	fmt.Printf("%d\n", findShortestPath(grid, end, 0, true))
}
