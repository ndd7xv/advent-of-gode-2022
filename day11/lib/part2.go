package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day11/inputs/input1.txt")

	monkeys := []Monkey{}

	commonDenominator := 1

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
				fmt.Printf("Monkey %d has items: %v\n", id, monkeys[id].items)
			case "Operation:":
				// Assumes every "operation" is in the format new = old <OPERATOR> <NUMBER> except
				// the case where <NUMBER> is old
				operator := parsed[4]
				number := parsed[5]

				// In the case where <NUMBER> is old, just hard code it lol
				if number == "old" {
					fmt.Printf("Operation is squared\n")
					monkeys[id].operation = func(i int) int {
						return i * i
					}
					continue
				}

				parsedNumber, _ := strconv.Atoi(parsed[5])
				switch operator {
				case "*":
					fmt.Printf("Operation is * %d\n", parsedNumber)
					monkeys[id].operation = func(i int) int {
						return i * parsedNumber
					}
				case "+":
					fmt.Printf("Operation is + %d\n", parsedNumber)
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
				commonDenominator *= num
				fmt.Printf("If divisible by %d throw to monkey %d else %d\n", num, ifTrue, ifFalse)
				monkeys[id].test = func(i int) int {
					if i%num == 0 {
						return ifTrue
					}
					return ifFalse
				}
			}
		}
	}

	for id, monkey := range monkeys {
		fmt.Printf("Monkey %d has:  %v\n", id, monkey.items)
	}

	// Hate to see it
	// https://stackoverflow.com/questions/41706212/impossible-to-modify-struct-inside-array
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			//fmt.Printf("Monkey %d: items are %v\n", id, monkey.items)
			items := len(monkeys[j].items)
			for k := 0; k < items; k++ {
				//fmt.Printf("Inspects %d -> %d\n", monkey.items[j], monkey.operation(monkey.items[j])/3)
				//fmt.Printf("Thrown to Monkey %d\n", monkey.test(monkey.items[j]))
				monkeys[j].items[k] = monkeys[j].operation(monkeys[j].items[k]) % commonDenominator
				monkeys[monkeys[j].test(monkeys[j].items[k])].items = append(monkeys[monkeys[j].test(monkeys[j].items[k])].items, monkeys[j].items[k])
				monkeys[j].inspected++
			}
			monkeys[j].items = []int{}
		}

		for id, monkey := range monkeys {
			fmt.Printf("After round %d, Monkey %d with %d inspections has: %v\n", i+1, id, monkey.inspected, monkey.items)
		}
	}

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

	fmt.Printf("Answer: %d", first*second)
}
