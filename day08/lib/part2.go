package day08

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day08/inputs/input1.txt")

	lines := strings.Split(string(input), "\n")
	rows := len(lines)
	columns := len(lines[0])

	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, columns)

	}

	for r, line := range lines {
		for c, ch := range line {
			grid[r][c] = int(ch - '0')
		}
	}

	maxScenicScore := 0.0
	for r := range lines {
		for c := range lines[0] {
			maxScenicScore = math.Max(maxScenicScore, calculateScenicScore(grid, r, c))
		}
	}
	fmt.Println(maxScenicScore)
}

func calculateScenicScore(grid [][]int, r int, c int) float64 {
	score := 1
	height := grid[r][c]

	count := 0
	for i := c - 1; i >= 0; i-- {
		count++
		if grid[r][i] >= height {
			break
		}
	}
	score *= count

	count = 0
	for i := c + 1; i < len(grid[0]); i++ {
		count++
		if grid[r][i] >= height {
			break
		}
	}
	score *= count

	count = 0
	for i := r - 1; i >= 0; i-- {
		count++
		if grid[i][c] >= height {
			break
		}
	}
	score *= count

	count = 0
	for i := r + 1; i < len(grid); i++ {
		count++
		if grid[i][c] >= height {
			break
		}
	}
	score *= count

	return float64(score)
}
