package day07

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := ioutil.ReadFile("inputs/test.txt")

	root := NewResource("/", nil, true)
	parentDir := []*SystemResource{root}
	currentDir := root

	// We skip over the initial cd because we assume that is standard for all inputs and everyone
	// starts at /. We assume also that when `$ cd` is used, it is used correctly for an existing
	// and known file.
	for i, line := range strings.Split(string(input), "\n") {
		if i == 0 {
			continue
		}
		if line == "" {
			break
		}

		if strings.HasPrefix(line, "$ cd") {
			newDir := strings.Split(line, " ")[2]
			if newDir == ".." {
				currentDir = parentDir[len(parentDir)-1]
				parentDir = parentDir[:len(parentDir)-1]
				continue
			}
			for _, child := range currentDir.Resources {
				if child.Name == newDir {
					parentDir = append(parentDir, currentDir)
					currentDir = child
					break
				}
			}
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "dir ") {
			newDir := strings.Split(line, " ")[1]
			currentDir.AddFile(NewResource(newDir, nil, true))
		} else {
			parsedLine := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(parsedLine[0])
			newFile := parsedLine[1]
			currentDir.AddFile(NewResource(newFile, Of(fileSize), false))
		}
	}
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
