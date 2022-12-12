package day12

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func Part1() {
	input, _ := os.ReadFile("day12/inputs/input1.txt")

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	grid := make([][]int, len(lines))

	var startX, startY int
	var endX, endY int

	for r, line := range lines {
		grid[r] = make([]int, len(line))
		for c, ch := range line {
			grid[r][c] = int(ch) - 'a'
			if ch == 'S' {
				startX, startY = r, c
				grid[r][c] = 0
			}
			if ch == 'E' {
				endX, endY = r, c
				grid[r][c] = 'z' - 'a' + 1
				// fmt.Printf("Coordinate of E is %d, %d\n", endX, endY)
			}
		}
	}

	shortestPaths := findShortestPath(grid, startX, startY)

	fmt.Printf("%d\n", shortestPaths[endX][endY])
}

func findShortestPath(grid [][]int, x, y int) [][]int {

	unvisited := make(map[Coordinate]bool)

	// Initialize the answers grid, where answers[r][c] contains the least amount of steps to that
	// position (r, c)
	answers := make([][]int, len(grid))
	for row := range answers {
		answers[row] = make([]int, len(grid[0]))
		for col := range answers[row] {
			answers[row][col] = -1
			unvisited[Coordinate{row, col}] = true
		}
	}
	answers[x][y] = 0
	current := Coordinate{x, y}

	neighbors := []Coordinate{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	for len(unvisited) != 0 {
		for _, neighbor := range neighbors {
			neighborX, neighborY := current.x-neighbor.x, current.y-neighbor.y
			// if neighborX == 4 && neighborY == 5 {
			// 	fmt.Printf("%d", answers[neighborX][neighborY])
			// }
			if neighborX < 0 || neighborY < 0 || neighborX == len(grid) || neighborY == len(grid[0]) {
				continue
			}
			if grid[neighborX][neighborY]-grid[current.x][current.y] > 1 {
				continue
			} else if answers[neighborX][neighborY] == -1 {
				//fmt.Printf("Setting %d, %d to %d\n", neighborX, neighborY, answers[current.x][current.y]+1)
				answers[neighborX][neighborY] = answers[current.x][current.y] + 1
			} else {
				answers[neighborX][neighborY] = int(math.Min(float64(answers[neighborX][neighborY]), float64(answers[current.x][current.y]+1)))
			}
		}

		delete(unvisited, current)

		// for row := range answers {
		// 	for col := range answers[row] {
		// 		fmt.Printf("%d ", answers[row][col])
		// 	}
		// 	fmt.Println()
		// }

		// This should be turned into a heap for better completion
		if len(unvisited) != 0 {
			lowestVisited := -1
			newX, newY := -1, -1
			for k := range unvisited {
				// fmt.Printf("%v\n", k)
				if answers[k.x][k.y] != -1 && (lowestVisited == -1 || answers[k.x][k.y] < lowestVisited) {
					lowestVisited = answers[k.x][k.y]
					newX, newY = k.x, k.y
				}
			}
			//fmt.Printf("Not empty! New coordinates are %d %d\n", newX, newY)
			if newX == -1 || newY == -1 {
				break
			}
			current = Coordinate{newX, newY}
		}
	}

	return answers
}
