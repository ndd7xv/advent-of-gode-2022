package day16

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part2() {
	input, _ := ioutil.ReadFile("inputs/input1.txt")

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
	}

	fmt.Println("Part 2, ready to go!")
}
