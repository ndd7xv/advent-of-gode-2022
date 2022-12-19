package day17

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// Used for checking if we have entered a loop
// Used as a part of map to see if the previous index has the same rock
type Event struct {
	rock      int
	x         int
	y         int
	iteration int
	max       int
}

func Part2() {
	input, _ := os.ReadFile("day17/inputs/input1.txt")

	rockCount := 0

	grid := [60000][7]int{}
	for i := 0; i < 7; i++ {
		grid[0][i] = 1
	}

	max := 0
	lowerLeft := Coordinate{2, max + 3}
	pattern := strings.Trim(string(input), "\n")
	eventMapper := make(map[int]Event)

	rocksPerLoop := -1
	heightPerLoop := -1

	for index := 0; ; index++ {
		// if index%len(pattern) == 0 && index != 0 && lowerLeft.y == max+3 {
		// 	fmt.Printf("Fresh start: %d\n", index/len(pattern))
		// 	break
		// }
		direction := pattern[index%len(pattern)]

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

		if event, ok := eventMapper[index%len(pattern)]; ok {
			if event.x == lowerLeft.x && event.max-event.y == max-lowerLeft.y && event.iteration > 0 {
				fmt.Printf("Index %d, Iteration %d to %d!\n", index, event.iteration, index/len(pattern))
				fmt.Printf("Before: %d | After: %d\n", event.rock, rockCount)
				fmt.Printf("Loop begins at index %d, relY: %d\n\n", index%len(pattern), max-lowerLeft.y)
				rocksPerLoop = rockCount - event.rock
				heightPerLoop = max - event.max
				break
			} else {
				eventMapper[index%len(pattern)] = Event{
					rockCount,
					lowerLeft.x,
					lowerLeft.y,
					index / len(pattern),
					max,
				}
			}
		} else {
			eventMapper[index%len(pattern)] = Event{
				rockCount,
				lowerLeft.x,
				lowerLeft.y,
				index / len(pattern),
				max,
			}
		}
	}

	fmt.Printf("Rocks/Loop: %d | Height/Loop: %d\n", rocksPerLoop, heightPerLoop)

	remainingRocks := 1000000000000 - rockCount
	loops := remainingRocks / rocksPerLoop

	rocksSkipped := (loops * rocksPerLoop)
	heightAdded := (loops * heightPerLoop)

	remainingRocks = remainingRocks - rocksSkipped

	fmt.Printf("Rocks skipped: %d | heightAdded: %d | remainingRocks: %d\n", rocksSkipped, heightAdded, remainingRocks)

	rockCount += rocksSkipped

	for index := 1; ; index++ {
		if rockCount == 1000000000000 {
			break
		}
		direction := pattern[index%len(pattern)]

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

	fmt.Println(heightAdded + max)
}
