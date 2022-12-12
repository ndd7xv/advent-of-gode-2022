package day12

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// Coordinate represents a point in the grid
type Coordinate struct {
	x int
	y int
}

func Part1() {
	input, _ := os.ReadFile("day12/inputs/input.txt")

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	grid := make([][]int, len(lines))

	start := Coordinate{-1, -1}
	for r, line := range lines {
		grid[r] = make([]int, len(line))
		for c, ch := range line {
			grid[r][c] = int(ch) - 'a'
			if ch == 'S' {
				start = Coordinate{r, c}
				grid[r][c] = 0
			}
			if ch == 'E' {
				grid[r][c] = 26 // 'z' - 'a' + 1
			}
		}
	}

	shortestPaths := findShortestPath(grid, start, 26, false)

	fmt.Printf("%d\n", shortestPaths)
}

func findShortestPath(grid [][]int, start Coordinate, endValue int, goingDown bool) int {

	unvisited := make(map[Coordinate]bool)

	// Initialize the answers grid, where answers[r][c] contains the least amount of steps to that
	// position (r, c). Floats because I don't like casting my ints to floats and back to ints.
	answers := make([][]float64, len(grid))
	for row := range answers {
		answers[row] = make([]float64, len(grid[0]))
		for col := range answers[row] {
			answers[row][col] = -1.0
			unvisited[Coordinate{row, col}] = true
		}
	}
	answers[start.x][start.y] = 0
	// Because positions are unweighted, we can use BFS instead of Djikstra's and consequently a
	// FIFO queue instead of a min-heap.
	queue := []Coordinate{{start.x, start.y}}

	neighbors := []Coordinate{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	for len(queue) != 0 {
		current := queue[0]

		// Break early if we've reached the destination
		if grid[current.x][current.y] == endValue {
			return int(answers[current.x][current.y])
		}

		for _, neighbor := range neighbors {
			neighborX, neighborY := current.x-neighbor.x, current.y-neighbor.y
			if neighborX < 0 || neighborY < 0 || neighborX == len(grid) || neighborY == len(grid[0]) {
				continue
			}

			// To reduce time in part 2, I add goingDown, which allows us to find the path from
			// the end to the first 'a'; this requires checking the opposite of the requirement
			// that each next step must be at MAX the current height + 1(each next step must be at
			// MIN 1 step lower)
			stepDistance := grid[neighborX][neighborY] - grid[current.x][current.y]
			if goingDown {
				stepDistance = grid[current.x][current.y] - grid[neighborX][neighborY]
			}
			if stepDistance > 1 {
				continue
			} else if answers[neighborX][neighborY] == -1.0 {
				answers[neighborX][neighborY] = answers[current.x][current.y] + 1
				// Add the neighboring node to the queue if unvisited
				if _, ok := unvisited[Coordinate{neighborX, neighborY}]; ok {
					queue = append(queue, Coordinate{neighborX, neighborY})
				}
			} else {
				answers[neighborX][neighborY] = math.Min(answers[neighborX][neighborY], answers[current.x][current.y]+1)
			}
		}

		delete(unvisited, queue[0])
		queue = queue[1:]
	}

	return -1
}
