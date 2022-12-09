package day08

import (
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func Part1() {
	input, _ := os.ReadFile("day08/inputs/input1.txt")
	lines := strings.Split(string(input), "\n")

	seen := make(map[Coordinate]bool)
	grid := make([][]int, len(lines))

	for r, line := range lines {
		grid[r] = make([]int, len(lines[0]))
		for c, ch := range line {
			grid[r][c] = int(ch - '0')
		}
	}

	for r := range lines {
		tallestFromLeft := -1
		tallestFromRight := -1
		for c := range lines[0] {
			if grid[r][c] > tallestFromLeft {
				seen[Coordinate{r, c}] = true
				tallestFromLeft = grid[r][c]
			}
			if grid[r][len(lines[0])-1-c] > tallestFromRight {
				seen[Coordinate{r, len(lines[0]) - 1 - c}] = true
				tallestFromRight = grid[r][len(lines[0])-1-c]
			}
		}
	}

	for c := range lines[0] {
		tallestFromBottom := -1
		tallestFromTop := -1
		for r := range lines {
			if grid[r][c] > tallestFromTop {
				seen[Coordinate{r, c}] = true
				tallestFromTop = grid[r][c]
			}
			if grid[len(lines)-1-r][c] > tallestFromBottom {
				seen[Coordinate{len(lines) - 1 - r, c}] = true
				tallestFromBottom = grid[len(lines)-1-r][c]
			}
		}
	}

	fmt.Println(len(seen))
}
