package day13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Item is either a list of more items, or an int.
type Item struct {
	items []*Item
	value *int
}

func (item Item) valueType() string {
	if item.items != nil {
		return "list"
	} else if item.value != nil {
		return "int"
	} else {
		return "empty"
	}
}

func Part1() {
	input, _ := os.ReadFile("day13/inputs/input1.txt")
	sum := 0
	for index, line := range strings.Split(strings.Trim(string(input), "\n"), "\n\n") {
		pair := strings.Split(line, "\n")

		first := Parse(pair[0])
		second := Parse(pair[1])

		fmt.Printf("Comparing the %d pair\n", index)

		if ro, _ := InRightOrder(*first, *second); ro {
			sum += (index + 1)
			fmt.Println("CORRECT")
		}
		//fmt.Println(line)
	}

	fmt.Println(sum)
}

func Parse(line string) *Item {
	// Find outer layer of brackets, assumes things are properly formatted
	var items []string
	// fmt.Println("Parsing", line)
	start := strings.Index(line, "[")
	if start != -1 {
		end := strings.LastIndex(line, "]")
		unparsedItems := line[start+1 : end]
		level := 0
		start := 0
		for i, ch := range unparsedItems {
			if ch == '[' {
				level++
			} else if ch == ']' {
				level--
			} else if ch == ',' && level == 0 {
				items = append(items, unparsedItems[start:i])
				start = i + 1
			}
		}
		items = append(items, unparsedItems[start:])
	} else {
		if strings.Contains(line, ",") {
			items = strings.Split(line, ",")
		} else if len(line) != 0 {
			num, _ := strconv.Atoi(line)
			return &Item{
				nil,
				&num,
			}
		}
	}

	// for _, item := range items {
	// 	fmt.Println("\t", item)
	// }
	// Create items from the list of items
	parsedItems := []*Item{}
	for _, item := range items {
		parsedItems = append(parsedItems, Parse(item))
	}

	return &Item{
		parsedItems,
		nil,
	}
}

func InRightOrder(first Item, second Item) (bool, string) {
	if first.valueType() == "empty" && first.valueType() != "empty" {
		return true, "stop"
	} else if first.valueType() != "empty" && first.valueType() == "empty" {
		return false, "stop"
	} else if first.valueType() == "int" && second.valueType() == "int" {
		fmt.Printf("\tComparing %d and %d:\n", *first.value, *second.value)
		if *first.value < *second.value {
			return true, "stop"
		} else if *first.value == *second.value {
			return true, ""
		} else {
			return false, "stop"
		}
	} else if first.valueType() == "list" && second.valueType() == "list" {
		fmt.Printf("Comparing %v and %v:\n", first.items, second.items)
		for i := 0; i < len(first.items); i++ {
			if i >= len(second.items) {
				return false, "stop"
			}
			res, cont := InRightOrder(*first.items[i], *second.items[i])
			if cont == "stop" {
				return res, "stop"
			}
		}
		if len(second.items) > len(first.items) {
			return true, "stop"
		} else if len(second.items) == len(first.items) {
			return true, ""
		}
	} else if first.valueType() == "int" {
		toList := Item{[]*Item{{nil, first.value}}, nil}
		for i := 0; i < len(toList.items); i++ {
			if i >= len(second.items) {
				return false, "stop"
			}
			res, cont := InRightOrder(*toList.items[i], *second.items[i])
			if cont == "stop" {
				return res, "stop"
			}
		}
		if len(second.items) > len(toList.items) {
			return true, "stop"
		} else if len(second.items) == len(toList.items) {
			return true, ""
		}
	} else if second.valueType() == "int" {
		toList := Item{[]*Item{{nil, second.value}}, nil}
		for i := 0; i < len(toList.items); i++ {
			if i >= len(first.items) {
				return true, "stop"
			}
			res, cont := InRightOrder(*first.items[i], *toList.items[i])
			if cont == "stop" {
				return res, "stop"
			}
		}
		if len(toList.items) > len(first.items) {
			return false, "stop"
		} else if len(toList.items) == len(first.items) {
			return true, ""
		}
	}
	return false, "stop"
}
