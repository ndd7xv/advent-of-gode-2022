package day17

import (
	"fmt"
	"math"
	"os"
	"strings"
)

/*
####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##
*/

type Coordinate struct {
	x, y int
}

func Part1() {
	input, _ := os.ReadFile("day17/inputs/input1.txt")

	rockCount := 0

	grid := [15000][7]int{}
	for i := 0; i < 7; i++ {
		grid[0][i] = 1
	}

	max := 0
	lowerLeft := Coordinate{2, max + 3}

	pattern := strings.Trim(string(input), "\n")
	for index := 0; ; index++ {
		direction := pattern[index%len(pattern)]
		if rockCount == 5000 {
			break
		}

		lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
		result := Lower((rockCount)%5, lowerLeft, &grid)

		if result > 0 {
			// fmt.Println("Settled!")
			max = int(math.Max(float64(max), float64(result)))
			rockCount++
			lowerLeft = Coordinate{2, max + 3}
		} else {
			lowerLeft.y--
		}
	}
	// for r := 10; r >= 0; r-- {
	// 	for c := 0; c < len(grid[0]); c++ {
	// 		if grid[r][c] == 1 {
	// 			fmt.Printf("#")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(max)
}

func Shift(direction byte, rock int, lowerLeft Coordinate, grid *[15000][7]int) Coordinate {

	newLeft := Coordinate{lowerLeft.x, lowerLeft.y}
	if direction == '<' {
		// fmt.Println("Moving left")
		newLeft.x--
	} else {
		// fmt.Println("Moving right")
		newLeft.x++
	}
	switch rock {
	case 0:
		for i := newLeft.x; i < newLeft.x+4; i++ {
			if i >= 7 || i < 0 || grid[newLeft.y+1][i] == 1 {
				return lowerLeft
			}
		}
		return newLeft
	case 1:
		if newLeft.x < 0 || newLeft.x+2 >= 7 || grid[newLeft.y+1][newLeft.x+1] == 1 || grid[newLeft.y+3][newLeft.x+1] == 1 || grid[newLeft.y+2][newLeft.x] == 1 || grid[newLeft.y+2][newLeft.x+2] == 1 {
			return lowerLeft
		}
		return newLeft
	case 2:
		for i := newLeft.x; i < newLeft.x+3; i++ {
			if i >= 7 || i < 0 || grid[newLeft.y+1][i] == 1 {
				return lowerLeft
			}
		}
		if grid[newLeft.y+2][newLeft.x+2] == 1 || grid[newLeft.y+3][newLeft.x+2] == 1 {
			return lowerLeft
		}
		return newLeft
	case 3:
		for j := newLeft.y + 1; j < newLeft.y+5; j++ {
			if newLeft.x >= 7 || newLeft.x < 0 || grid[j][newLeft.x] == 1 {
				return lowerLeft
			}
		}
		return newLeft
	case 4:
		if newLeft.x+1 >= 7 || newLeft.x < 0 || grid[newLeft.y+1][newLeft.x] == 1 || grid[newLeft.y+2][newLeft.x] == 1 || grid[newLeft.y+1][newLeft.x+1] == 1 || grid[newLeft.y+2][newLeft.x+1] == 1 {
			return lowerLeft
		}
		return newLeft
	}

	return lowerLeft
}

func Lower(rock int, lowerLeft Coordinate, grid *[15000][7]int) int {
	switch rock {
	case 0:
		for i := lowerLeft.x; i < lowerLeft.x+4; i++ {
			if grid[lowerLeft.y][i] == 1 {
				for j := lowerLeft.x; j < lowerLeft.x+4; j++ {
					grid[lowerLeft.y+1][j] = 1
				}
				return lowerLeft.y + 1
			}
		}
	case 1:
		if grid[lowerLeft.y][lowerLeft.x+1] == 1 || grid[lowerLeft.y+1][lowerLeft.x] == 1 || grid[lowerLeft.y+1][lowerLeft.x+2] == 1 {
			grid[lowerLeft.y+1][lowerLeft.x+1] = 1
			grid[lowerLeft.y+2][lowerLeft.x] = 1
			grid[lowerLeft.y+2][lowerLeft.x+1] = 1
			grid[lowerLeft.y+2][lowerLeft.x+2] = 1
			grid[lowerLeft.y+3][lowerLeft.x+1] = 1
			return lowerLeft.y + 3
		}
	case 2:
		for i := lowerLeft.x; i < lowerLeft.x+3; i++ {
			if grid[lowerLeft.y][i] == 1 {
				for j := lowerLeft.x; j < lowerLeft.x+3; j++ {
					grid[lowerLeft.y+1][j] = 1
				}
				grid[lowerLeft.y+2][lowerLeft.x+2] = 1
				grid[lowerLeft.y+3][lowerLeft.x+2] = 1
				return lowerLeft.y + 3
			}
		}
	case 3:
		if grid[lowerLeft.y][lowerLeft.x] == 1 {
			grid[lowerLeft.y+1][lowerLeft.x] = 1
			grid[lowerLeft.y+2][lowerLeft.x] = 1
			grid[lowerLeft.y+3][lowerLeft.x] = 1
			grid[lowerLeft.y+4][lowerLeft.x] = 1
			return lowerLeft.y + 4
		}
	case 4:
		if grid[lowerLeft.y][lowerLeft.x] == 1 || grid[lowerLeft.y][lowerLeft.x+1] == 1 {
			grid[lowerLeft.y+1][lowerLeft.x] = 1
			grid[lowerLeft.y+1][lowerLeft.x+1] = 1
			grid[lowerLeft.y+2][lowerLeft.x] = 1
			grid[lowerLeft.y+2][lowerLeft.x+1] = 1
			return lowerLeft.y + 2
		}
	}
	return 0
}
