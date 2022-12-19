package day16

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	flowRate int
	valves   []*Node
}

func NewNode(name string, flow int) *Node {
	return &Node{
		name,
		flow,
		[]*Node{},
	}
}

type Answer struct {
	start     string
	remaining int
	opened    string
}

var cache map[Answer]int

func Part1() {
	input, _ := os.ReadFile("day16/inputs/test.txt")

	cache = make(map[Answer]int)

	nodes := ParseInput(string(input))
	// on := make(map[string]bool)

	fmt.Printf("%d\n", maxPressure("AA", 30, []string{}, nodes))
}

// This does NOT give the right answer for the example, and consequently I think it was mostly
// luck that I got this part. I'll describe an alternative/fix after I'm done with part 2 and
// have look at other solutions after that
func maxPressure(start string, remaining int, opened []string, nodes map[string]*Node) int {
	openedAsString := strings.Join(opened, " ")
	if answer, exists := cache[Answer{start, remaining, openedAsString}]; exists {
		return answer
	}

	if remaining <= 1 {
		return 0
	}

	max := 0

	for _, valve := range nodes[start].valves {
		// If you decide to skip...
		max = int(math.Max(float64(max), float64(maxPressure(valve.name, remaining-1, opened, nodes))))
		// If you turn on this valve
		if !contains(opened, start) && (nodes[start].flowRate*(remaining-1)) > 0 {
			opened = append(opened, start)
			sort.Strings(opened)
			max = int(math.Max(float64(max), float64((nodes[start].flowRate*(remaining-1))+maxPressure(valve.name, remaining-2, opened, nodes))))
			for i := range opened {
				if opened[i] == start {
					opened[i] = opened[len(opened)-1]
					opened = opened[:len(opened)-1]
					sort.Strings(opened)
					break
				}
			}
		}
	}

	cache[Answer{start, remaining, openedAsString}] = max
	return max
}

func ParseInput(input string) map[string]*Node {
	nodes := make(map[string]*Node)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}

		words := strings.Split(line, " ")
		valveName := words[1]
		var flowRate int
		neighbors := []*Node{}

		for i, word := range words {
			if strings.HasPrefix(word, "rate=") {
				flowRate, _ = strconv.Atoi(word[5 : len(word)-1])
			} else if strings.HasPrefix(word, "valves") {
				for _, valve := range words[i+1:] {
					valve = strings.TrimSuffix(valve, ",")
					if _, ok := nodes[valve]; !ok {
						nodes[valve] = NewNode(valve, 0)
					}
					neighbors = append(neighbors, nodes[valve])
				}
				break
			}
		}

		if _, ok := nodes[valveName]; !ok {
			nodes[valveName] = NewNode(valveName, flowRate)
		} else {
			nodes[valveName].flowRate = flowRate
		}
		nodes[valveName].valves = neighbors
	}

	return nodes
}

func Of[E any](e E) *E {
	return &e
}

func contains(arr []string, str string) bool {
	for _, word := range arr {
		if str == word {
			return true
		}
	}
	return false
}
