package day01

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Part1() {
	file, _ := os.Open("day01/inputs/input1.txt")
	defer file.Close()
	reader := bufio.NewReader(file)

	max := 0
	sum := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if len(string(line)) == 0 {
			if max < sum {
				max = sum
			}
			sum = 0
		} else {
			num, _ := strconv.Atoi(string(line))
			sum += num
		}
	}
	if max < sum {
		max = sum
	}
	fmt.Println(max)
}
