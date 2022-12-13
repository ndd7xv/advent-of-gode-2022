package day13

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day13/inputs/input1.txt")

	messages := []Item{}
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		if line == "" {
			continue
		}
		messages = append(messages, *Parse(line))
	}

	messages = append(messages, *Parse("[[2]]"))
	messages = append(messages, *Parse("[[6]]"))
	sort.Slice(messages, func(first, second int) bool {
		res, _ := InRightOrder(messages[first], messages[second])
		return res
	})

	decoderProduct := 1
	for i, message := range messages {
		if message.items != nil && len(message.items) == 1 {
			if message.items[0].items != nil && len(message.items[0].items) == 1 && message.items[0].items[0].value != nil && (*message.items[0].items[0].value == 2 || *message.items[0].items[0].value == 6) {
				decoderProduct *= (1 + i)
			}
		}
	}
	fmt.Println(decoderProduct)
}
