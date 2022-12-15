package day15

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

func Part1() {
	input, _ := os.ReadFile("day15/inputs/input1.txt")

	sensors := []Coordinate{}
	beacons := []Coordinate{}
	distances := []int{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		sensor, beacon := ParseLine(line)

		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
		distances = append(distances, ManhattanDistance(sensor, beacon))

		// fmt.Printf("Sensor: (%d, %d) | Beacon: (%d, %d)\n", sensor.x, sensor.y, beacon.x, beacon.y)
	}

	invalidX := make(map[int]bool)

	for i, sensor := range sensors {
		dY := math.Abs(float64(sensor.y) - 2000000)
		dX := distances[i] - int(dY)
		// fmt.Printf("%f\n", dY)
		if dX >= 0 {
			for j := -dX + sensor.x; j <= dX+sensor.x; j++ {
				invalidX[int(j)] = true
				for _, beacon := range beacons {
					if beacon.x == j && beacon.y == 2000000 {
						delete(invalidX, int(j))
					}
				}
			}
		}
	}

	// keys := []int{}
	// for k := range invalidX {
	// 	keys = append(keys, k)
	// }
	// sort.Ints(keys)
	// fmt.Printf("%v", keys)
	fmt.Println(len(invalidX))
}

func ManhattanDistance(a, b Coordinate) int {
	return int(math.Abs(float64(a.x)-float64(b.x)) + math.Abs(float64(a.y)-float64(b.y)))
}

func ParseLine(s string) (Coordinate, Coordinate) {
	split := strings.Split(s, " closest ")

	var x, y int
	for _, word := range strings.Split(split[0], " ") {
		if strings.HasPrefix(word, "x=") {
			x, _ = strconv.Atoi(word[2 : len(word)-1])
		} else if strings.HasPrefix(word, "y=") {
			y, _ = strconv.Atoi(word[2 : len(word)-1])
		}
	}
	sensor := Coordinate{x, y}

	for _, word := range strings.Split(split[1], " ") {
		if strings.HasPrefix(word, "x=") {
			x, _ = strconv.Atoi(word[2 : len(word)-1])
		} else if strings.HasPrefix(word, "y=") {
			y, _ = strconv.Atoi(word[2:])
		}
	}
	beacon := Coordinate{x, y}
	return sensor, beacon
}
