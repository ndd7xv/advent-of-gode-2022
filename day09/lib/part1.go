package day09

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part1() {
	input, _ := ioutil.ReadFile("inputs/input1.txt")

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		
	}
	fmt.Println("Part 1, ready to go!")
}
