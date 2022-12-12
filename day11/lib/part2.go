package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day11/inputs/input1.txt")

	monkeys := ParseInput(input)

	commonDenominator := 1
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		words := strings.Split(strings.Trim(line, " "), " ")
		if words[0] == "Test:" {
			divisor, _ := strconv.Atoi(words[3])
			commonDenominator *= divisor
		}
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := &monkeys[j]
			items := len(monkey.items)
			for k := 0; k < items; k++ {
				item := monkey.items[k]

				item = monkey.operation(item) % commonDenominator
				monkeys[monkey.test(item)].items = append(monkeys[monkey.test(item)].items, item)
				monkey.inspected++
			}
			monkey.items = []int{}
		}
	}

	first, second := FindTopTwoInspections(monkeys)

	fmt.Printf("Answer: %d\n", first*second)
}
