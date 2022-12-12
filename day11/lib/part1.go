package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	input, _ := os.ReadFile("day11/inputs/input1.txt")

	monkeys := ParseInput(input)

	// Hate to see it; you cannot range over an array and expect a reference to it - gotta index
	// https://stackoverflow.com/questions/41706212/impossible-to-modify-struct-inside-array
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := &monkeys[j]
			items := len(monkey.items)
			for k := 0; k < items; k++ {
				item := monkey.items[k]

				item = monkey.operation(item) / 3
				monkeys[monkey.test(item)].items = append(monkeys[monkey.test(item)].items, item)
				monkey.inspected++
			}
			monkey.items = []int{}
		}
	}

	first, second := FindTopTwoInspections(monkeys)

	fmt.Printf("Answer: %d\n", first*second)
}

// FindTopTwoInspections finds the first and second greatest numbers and returns them.
func FindTopTwoInspections(monkeys []Monkey) (int, int) {
	first := 0
	second := 0

	for _, monkey := range monkeys {
		if monkey.inspected > first {
			second = first
			first = monkey.inspected
		} else if monkey.inspected > second {
			second = monkey.inspected
		}
	}
	return first, second
}

// ParseInput goes through the input file and creates the monkeys with the right items, operations,
// and tests.
func ParseInput(input []byte) []Monkey {
	monkeys := []Monkey{}
	for id, monkey := range strings.Split(strings.Trim(string(input), "\n"), "\n\n") {
		monkeys = append(monkeys, New())
		monkeyAttributes := strings.Split(monkey, "\n")

		for i := 0; i < len(monkeyAttributes); i++ {
			parsed := strings.Split(strings.Trim(monkeyAttributes[i], "\n\t "), " ")
			switch parsed[0] {
			case "Starting":
				nums := strings.Split(strings.Join(parsed[2:], ""), ",")
				for _, num := range nums {
					converted, _ := strconv.Atoi(num)
					monkeys[id].items = append(monkeys[id].items, converted)
				}
			case "Operation:":
				// Assumes every "operation" is in the format new = old <OPERATOR> <NUMBER> except
				// the case where <NUMBER> is old
				operator := parsed[4]
				number := parsed[5]

				// In the case where <NUMBER> is old, just hard code it lol
				if number == "old" {
					monkeys[id].operation = func(i int) int {
						return i * i
					}
					continue
				}

				parsedNumber, _ := strconv.Atoi(parsed[5])
				switch operator {
				case "*":
					monkeys[id].operation = func(i int) int {
						return i * parsedNumber
					}
				case "+":
					monkeys[id].operation = func(i int) int {
						return i + parsedNumber
					}
				}

			case "Test:":
				// Assumes that test is always a divisible statement and the result is throwing to
				// a monkey
				num, _ := strconv.Atoi(parsed[3])
				ifTrue, _ := strconv.Atoi(strings.Split(strings.Trim(monkeyAttributes[i+1], " "), " ")[5])
				ifFalse, _ := strconv.Atoi(strings.Split(strings.Trim(monkeyAttributes[i+2], " "), " ")[5])
				monkeys[id].test = func(i int) int {
					if i%num == 0 {
						return ifTrue
					}
					return ifFalse
				}
			}
		}
	}
	return monkeys
}
