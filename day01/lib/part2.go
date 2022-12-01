package day01

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func Part2() {
	file, _ := os.Open("inputs/input1.txt")
	defer file.Close()
	reader := bufio.NewReader(file)

	var elves []int
	sum := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if len(string(line)) == 0 {
			elves = append(elves, sum)
			sum = 0
		} else {
			num, _ := strconv.Atoi(string(line))
			sum += num
		}
	}
	elves = append(elves, sum)
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
