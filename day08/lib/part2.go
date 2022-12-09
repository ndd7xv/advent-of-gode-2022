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

	grid := make([][]int, len(lines))
	for r, line := range lines {
		grid[r] = make([]int, len(lines[0]))
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

// I contemplated (and shared with my friends on discord) the following:
// "i am waking up to remind myself a good way to make part 2 more efficient (maybe not memory wise) is store how far you can see from each direction in its own grid maybe; if next one is bigger then 1 but if smaller then add  next one's directional score and check if the tree it stopped on is smaller than current height; if so additive repeat"
// "as a more lucid nathan this doesn't make things much more efficient as i would've liked; in the case of calculating the trees visible from 9 to the right in 98888888888 things would technically still be n since we'd only be skipping over every other point;  a best case scenario of 987654321987654321 is probably a best case scenario"
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
