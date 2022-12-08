package day08

import (
	"fmt"
	"os"
	"strings"
)

func Part1() {
	input, _ := os.ReadFile("day08/inputs/input1.txt")

	lines := strings.Split(string(input), "\n")
	rows := len(lines)
	columns := len(lines[0])

	grid := make([][][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([][]int, columns)
		for c := 0; c < columns; c++ {
			grid[i][c] = make([]int, 2)
		}
	}

	for r, line := range lines {
		for c, ch := range line {
			grid[r][c][0] = int(ch - '0')
		}
	}

	visible := 0
	for r := range lines {
		tallestFromLeft := -1
		tallestFromRight := -1
		for c := range lines[0] {
			if grid[r][c][0] > tallestFromLeft {
				if grid[r][c][1] != 1 {
					visible++
					grid[r][c][1] = 1
				}
				tallestFromLeft = grid[r][c][0]
			}
			if grid[r][len(lines[0])-1-c][0] > tallestFromRight {
				if grid[r][len(lines[0])-1-c][1] != 1 {
					visible++
					grid[r][len(lines[0])-1-c][1] = 1
				}
				tallestFromRight = grid[r][len(lines[0])-1-c][0]
			}
		}
	}

	for c := range lines[0] {
		tallestFromBottom := -1
		tallestFromTop := -1
		for r := range lines {
			if grid[r][c][0] > tallestFromTop {
				if grid[r][c][1] != 1 {
					visible++
					grid[r][c][1] = 1
				}
				tallestFromTop = grid[r][c][0]
			}
			if grid[len(lines)-1-r][c][0] > tallestFromBottom {
				if grid[len(lines)-1-r][c][1] != 1 {
					visible++
					grid[len(lines)-1-r][c][1] = 1
				}
				tallestFromBottom = grid[len(lines)-1-r][c][0]
			}
		}
	}
	fmt.Println(visible)
}
