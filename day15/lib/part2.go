package day15

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Part2() {
	input, _ := os.ReadFile("day15/inputs/input1.txt")

	sensors := []Coordinate{}
	// beacons := []Coordinate{}
	distances := []int{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		sensor, beacon := ParseLine(line)

		sensors = append(sensors, sensor)
		// beacons = append(beacons, beacon)
		distances = append(distances, ManhattanDistance(sensor, beacon))

		// fmt.Printf("Sensor: (%d, %d) | Beacon: (%d, %d)\n", sensor.x, sensor.y, beacon.x, beacon.y)
	}

out:
	for x := 0; x <= 4000000; x++ {
		for y := 0; y <= 4000000; y++ {
			// Apparently running this code in my terminal w/o the if check would eventually stop
			// printing to stdout?
			// if x > 4000000 {
			// 	fmt.Printf("Testing (%d, %d)\n", x, y)
			// }
			invalid := false
			for i, sensor := range sensors {
				if ManhattanDistance(sensor, Coordinate{x, y}) <= distances[i] {
					dX := int(math.Abs(float64(sensor.x - x)))
					dY := y - sensor.y
					y += distances[i] - dX - dY
					invalid = true
					break
				}
			}
			if !invalid {
				fmt.Printf("(%d, %d): %d\n", x, y, x*4000000+y)
				break out
			}
		}
	}

}
