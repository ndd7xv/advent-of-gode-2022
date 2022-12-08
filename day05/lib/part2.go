package day05

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("inputs/input1.txt")

	splitInput := strings.Split(string(input), "\n\n")
	cratePositions := splitInput[0]
	directions := splitInput[1]

	columns := setupCrates(cratePositions)
	updatedPositions := moveCratesPart2(columns, directions)

	crateTops := ""

	for _, column := range updatedPositions {
		crateTops += string(column.Pop().(int32))
	}

	fmt.Println(crateTops)
}

func moveCratesPart2(columns []*stack.Stack, directions string) []*stack.Stack {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	for _, direction := range strings.Split(directions, "\n") {
		if direction == "" {
			break
		}

		matches := re.FindAllString(direction, -1)
		nums := make([]int, len(matches))
		for i, num := range matches {
			parsed, _ := strconv.Atoi(num)
			nums[i] = parsed
		}
		fmt.Println(nums)
		fmt.Printf("move %d from %d to %d\n", nums[0], nums[1], nums[2])

		source := columns[nums[1]-1]
		dest := columns[nums[2]-1]

		tempStack := stack.New()
		for i := 0; i < nums[0]; i++ {
			tempStack.Push(source.Pop())
		}

		for tempStack.Peek() != nil {
			dest.Push(tempStack.Pop())
		}
	}

	return columns
}
