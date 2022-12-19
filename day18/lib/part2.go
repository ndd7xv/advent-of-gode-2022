package day18

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day18/inputs/input1.txt")

	cubes := make(map[Cube]bool)

	min := math.MaxInt
	max := math.MinInt

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}

		dimensions := strings.Split(line, ",")
		x, _ := strconv.Atoi(dimensions[0])
		y, _ := strconv.Atoi(dimensions[1])
		z, _ := strconv.Atoi(dimensions[2])

		min = int(math.Min(float64(min), math.Min(math.Min(float64(x), float64(y)), float64(z))))
		max = int(math.Max(float64(max), math.Max(math.Max(float64(x), float64(y)), float64(z))))

		cubes[NewCube(x, y, z)] = true
	}
	cubesAsSlice := []Cube{}
	for k := range cubes {
		cubesAsSlice = append(cubesAsSlice, k)
	}
	totalSurfaceArea := ExternalSurfaceArea(cubesAsSlice)

	// Represents all of the isolated areas of space that are not occupied as magma cubes; adjacent
	// are clumped together
	unselectedBlocks := []map[Cube]bool{}
	visited := make(map[Cube]bool)

	// Basically encases all input points in a cube and uses those inputs as boundaries
	// On the very first of the three-leveled for-loop, we find only the external cubes; additional
	// cubes added to unselectedBlocks are internal cubes
	for x := min - 1; x <= max+1; x++ {
		for y := min - 1; y <= max+1; y++ {
			for z := min - 1; z <= max+1; z++ {
				cube := NewCube(x, y, z)
				if _, processed := visited[cube]; !processed {
					if _, input := cubes[cube]; !input {
						newUnselectedBlock := map[Cube]bool{}
						queue := []Cube{cube}
						for len(queue) != 0 {
							current := queue[0]
							_, seen := visited[current]
							_, isInput := cubes[current]
							if current.x >= min-1 &&
								current.x <= max+1 &&
								current.y >= min-1 &&
								current.y <= max+1 &&
								current.z >= min-1 &&
								current.z <= max+1 && !seen && !isInput {
								visited[current] = true
								newUnselectedBlock[current] = true
								queue = append(queue, NewCube(current.x, current.y, current.z+1))
								queue = append(queue, NewCube(current.x, current.y, current.z-1))
								queue = append(queue, NewCube(current.x, current.y+1, current.z))
								queue = append(queue, NewCube(current.x, current.y-1, current.z))
								queue = append(queue, NewCube(current.x+1, current.y, current.z))
								queue = append(queue, NewCube(current.x-1, current.y, current.z))
							}
							queue = queue[1:]
						}
						unselectedBlocks = append(unselectedBlocks, newUnselectedBlock)
					}
				}
			}
		}
	}

	unselectedBlocks = unselectedBlocks[1:]
	fmt.Println(totalSurfaceArea)

	for i := range unselectedBlocks {
		// fmt.Printf("%v\n", cubes)
		blockList := []Cube{}
		for c := range unselectedBlocks[i] {
			blockList = append(blockList, c)
		}
		totalSurfaceArea -= ExternalSurfaceArea(blockList)
	}

	fmt.Println(totalSurfaceArea)
}

func ExternalSurfaceArea(cubes []Cube) int {
	// fmt.Printf("%v\n", cubes)

	totalSurfaceArea := len(cubes) * 6
	for i := 0; i < len(cubes)-1; i++ {
		for j := i + 1; j < len(cubes); j++ {
			if SharesSide(cubes[i], cubes[j]) {
				totalSurfaceArea -= 2
			}
		}
	}

	return totalSurfaceArea
}
