package day07

import (
	"fmt"
	"os"
	"math"
)

func Part2() {
	input, _ := os.ReadFile("inputs/test.txt")

	root := CreateFileSystem(input)
	freeTarget := 30000000 - (70000000 - root.GetSize())
	fmt.Println(findSmallestWorkingDirectorySize(root, freeTarget))
}

func findSmallestWorkingDirectorySize(root *SystemResource, freeTarget int) int {
	smallestSize := root.GetSize()
	for _, resource := range root.Resources {
		if resource.Resources != nil && resource.GetSize() >= freeTarget {
			smallestSize = int(math.Min(float64(smallestSize), float64(findSmallestWorkingDirectorySize(resource, freeTarget))))
		}
	}
	return smallestSize
}
