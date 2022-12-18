package day17

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day17/inputs/input1.txt")

	rockCount := 0

	grid := [15000][7]int{}
	for i := 0; i < 7; i++ {
		grid[0][i] = 1
	}

	max := 0
	lowerLeft := Coordinate{2, max + 3}
	pattern := strings.Trim(string(input), "\n")

	// Run the first loop
	for _, direction := range pattern {
		lowerLeft = Shift(byte(direction), (rockCount)%5, lowerLeft, &grid)
		result := Lower((rockCount)%5, lowerLeft, &grid)
		if result > 0 {
			max = int(math.Max(float64(max), float64(result)))
			rockCount++
			lowerLeft = Coordinate{2, max + 3}
		} else {
			lowerLeft.y--
		}
	}
	start := -1

	for i := 0; ; i++ {
		direction := pattern[i%len(pattern)]
		lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
		result := Lower((rockCount)%5, lowerLeft, &grid)
		if result > 0 {
			max = int(math.Max(float64(max), float64(result)))
			rockCount++
			lowerLeft = Coordinate{2, max + 3}
			start = i + 1
			break
		} else {
			lowerLeft.y--
		}
	}

	firstMax := max
	firstRockCount := rockCount

	fmt.Printf("After loop 1: %d from %d rocks\n", firstMax, firstRockCount)
	fmt.Printf("%d %d to %d\n", lowerLeft.x, lowerLeft.y, max)

	// for j := 2; ; j++ {
	// 	for i := start; i != start-1; i = ((i + 1) % len(pattern)) {
	// 		direction := pattern[i]
	// 		lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
	// 		result := Lower((rockCount)%5, lowerLeft, &grid)
	// 		if result > 0 {
	// 			max = int(math.Max(float64(max), float64(result)))
	// 			rockCount++
	// 			lowerLeft = Coordinate{2, max + 3}
	// 			// fmt.Println("Clean!")
	// 		} else {
	// 			lowerLeft.y--
	// 			// fmt.Println("Dirty!")
	// 		}
	// 	}
	// 	fmt.Printf("After Loop %d: %d from %d rocks\n", j, max, rockCount)
	// }
	for i := start; i != start-1; i = ((i + 1) % len(pattern)) {
		direction := pattern[i]
		lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
		result := Lower((rockCount)%5, lowerLeft, &grid)
		if result > 0 {
			max = int(math.Max(float64(max), float64(result)))
			rockCount++
			lowerLeft = Coordinate{2, max + 3}
		} else {
			lowerLeft.y--
		}
	}
	direction := pattern[start-1]
	lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
	result := Lower((rockCount)%5, lowerLeft, &grid)
	if result > 0 {
		max = int(math.Max(float64(max), float64(result)))
		rockCount++
		lowerLeft = Coordinate{2, max + 3}
	} else {
		lowerLeft.y--
	}

	secondMax := max
	secondRockCount := rockCount

	fmt.Printf("After Loop 2: %d from %d rocks\n", secondMax, secondRockCount)
	fmt.Printf("%d %d to %d\n", lowerLeft.x, lowerLeft.y, secondMax)

	for i := start; i != start-1; i = ((i + 1) % len(pattern)) {
		direction := pattern[i]
		lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
		result := Lower((rockCount)%5, lowerLeft, &grid)
		if result > 0 {
			max = int(math.Max(float64(max), float64(result)))
			rockCount++
			lowerLeft = Coordinate{2, max + 3}
		} else {
			lowerLeft.y--
		}
	}
	direction = pattern[start-1]
	lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
	result = Lower((rockCount)%5, lowerLeft, &grid)
	if result > 0 {
		max = int(math.Max(float64(max), float64(result)))
		rockCount++
		lowerLeft = Coordinate{2, max + 3}
	} else {
		lowerLeft.y--
	}

	thirdMax := max
	thirdRockCount := rockCount

	fmt.Printf("After Loop 3: %d from %d rocks\n", thirdMax, thirdRockCount)
	heightPerLoop := thirdMax - secondMax
	rocksPerLoop := thirdRockCount - secondRockCount
	fmt.Printf("h/l: %d r/l: %d\n", heightPerLoop, rocksPerLoop)
	// Initial case: 5549 height from 3498 rocks
	loops := (1000000000000 - 5249) / rocksPerLoop
	height := 8390 + (heightPerLoop * loops)

	remainingRocks := 1000000000000 - 5249 - (loops * rocksPerLoop)
	fmt.Printf("Remaining Rocks: %d\n", remainingRocks)

	newGrid := [15000][7]int{}
	for i := thirdMax - 20; i <= thirdMax; i++ {
		for j := 0; j < 7; j++ {
			newGrid[i-(thirdMax-20)][j] = grid[i][j]
		}
	}
	grid = newGrid
	fmt.Printf("%d %d to %d\n", lowerLeft.x, lowerLeft.y, thirdMax)

	lastMax := 20
	lowerLeft.y = lowerLeft.y - (max - lastMax)
	fmt.Printf("Last Max: %d %d to %d\n", lowerLeft.x, lowerLeft.y, lastMax)

	for remainingRocks != 0 {
		for i := start; ; i = ((i + 1) % len(pattern)) {
			direction := pattern[i]
			lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
			result := Lower((rockCount)%5, lowerLeft, &grid)

			if result > 0 {
				// fmt.Println("Settled!")
				lastMax = int(math.Max(float64(lastMax), float64(result)))
				remainingRocks--
				lowerLeft = Coordinate{2, max + 3}
				break
			} else {
				lowerLeft.y--
			}
		}

	}
	height += (lastMax - 20)
	fmt.Printf("%d\n", height)

}

// 15971 after 10000 rocks
// Repeats after running the entire loop
//		Run the loop twice
// 		Compare the latest storeMax with the previous to get amountPerLoop
//		(1000000000000 / len(pattern)) - 2 => how many more loops it does
// 		Add (that * amountPerLoop) to the two loops storeMax you already did
//

/*
fmt.Println("---")
	totalLoops := (1000000000000 - 1749) / 1750
	remainingRockCount := (1000000000000 - 1749) - (totalLoops * 1750)
	fmt.Printf("How many times it loops: %d\n", totalLoops)
	fmt.Printf("Remaining: %d\n", remainingRockCount)

	// To find the height after n loops: 2798 + (2787*(n-1))
	// 2798 represents the height after the first loop
	bigMax := 2798 + (2787 * (totalLoops - 1))
	newGrid := [15000][7]int{}
	for i := thirdMax - 10; i <= thirdMax; i++ {
		for j := 0; j < 7; j++ {
			newGrid[i-(thirdMax-10)][j] = grid[i][j]
		}
	}
	max := 10

	for i := 9; ; i++ {
		if remainingRockCount == 0 {
			break
		}
		direction := pattern[i%len(pattern)]
		lowerLeft = Shift(direction, (rockCount)%5, lowerLeft, &grid)
		result := Lower((rockCount)%5, lowerLeft, &grid)

		if result > 0 {
			// fmt.Println("Settled!")
			max = int(math.Max(float64(max), float64(result)))
			remainingRockCount--
			lowerLeft = Coordinate{2, max + 3}
		} else {
			lowerLeft.y--
		}
	}
	bigMax += (max - 10)
	fmt.Printf("Answer: %d\n", bigMax)
*/
