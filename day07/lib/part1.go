package day07

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Part1() {
	input, _ := ioutil.ReadFile("inputs/input1.txt")

	root := CreateFileSystem(input)

	fmt.Println(sumSizesLessThan100000(root))
}

// CreateFileSystem reads input and interprets the commands to formulate a tree.
// We skip over the initial cd because we assume that is standard for all inputs and everyone
// starts at /. We assume also that when `$ cd` is used, it is used correctly for an existing
// and known file.
func CreateFileSystem(input []byte) *SystemResource {
	root := NewResource("/", nil, true)
	parentDir := []*SystemResource{root}
	currentDir := root

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
	return root
}

// Recursive call on directories that returns their cumulative sizes if less than 100000
func sumSizesLessThan100000(node *SystemResource) int {
	total := 0
	for _, resource := range node.Resources {
		if resource.Resources != nil {
			total += sumSizesLessThan100000(resource)
		}
	}
	if node.GetSize() <= 100000 {
		total += node.GetSize()
	}
	return total
}
