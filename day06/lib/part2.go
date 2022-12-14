package day06

import (
	"fmt"
	"os"
)

func Part2() {
	input, _ := os.ReadFile("day06/inputs/input1.txt")

	for i := 0; i < len(input)-14; i++ {
		scan := input[i:i+14];

		m := make(map[byte]bool)
		unique := true


		for _, ch := range scan {
			if _, ok := m[ch]; ok {
				unique = false
				break
			}
			m[ch] = true
		}

		if unique {
			fmt.Println(i+14)
			break
		}
	}
}
