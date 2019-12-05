package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	isFirstWire := true
	steps := make(map[step]int)
	scanner := bufio.NewScanner(file)
	minDistance := int(^uint(0) >> 1)

	for scanner.Scan() {
		step := step{}
		stepCount := 0
		line := scanner.Text()

		for _, token := range strings.Split(line, ",") {
			direction, count := token[:1], atoi(token[1:])

			for i := 0; i < count; i++ {
				switch direction {
				case "R":
					step.x++

				case "L":
					step.x--

				case "U":
					step.y++

				case "D":
					step.y--
				}

				stepCount++

				if isFirstWire {
					if _, exists := steps[step]; !exists {
						steps[step] = stepCount
					}

				} else {
					//partOne(steps, step, &minDistance)
					partTwo(steps, step, stepCount, &minDistance)
				}

			}

		}

		isFirstWire = false
	}

	check(scanner.Err())

	fmt.Println("Distance:", minDistance)
}

type step struct{ x, y int }

func partOne(steps map[step]int, secondWireStep step, minDistance *int) {
	if _, exists := steps[secondWireStep]; exists {

		distance := abs(-secondWireStep.x) + abs(-secondWireStep.y)
		if *minDistance > distance {
			*minDistance = distance
		}
	}
}

func partTwo(steps map[step]int, secondWireStep step, secondWireStepCount int, minDistance *int) {
	if stepCount, exists := steps[secondWireStep]; exists {

		distance := stepCount + secondWireStepCount
		if *minDistance > distance {
			*minDistance = distance
		}
	}
}

func atoi(value string) int {
	parsedValue, err := strconv.Atoi(value)
	check(err)

	return parsedValue
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
