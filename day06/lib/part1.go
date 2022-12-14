package day06

import (
	"fmt"
	"os"
)

func Part1() {
	input, _ := os.ReadFile("day06/inputs/input1.txt")

	for i := 0; i < len(input)-4; i++ {
		scan := input[i:i+4];

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
			fmt.Println(i+4)
			break
		}
	}
}
