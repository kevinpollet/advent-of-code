package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type step struct {
	x, y int
}

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

		for _, token := range strings.Split(scanner.Text(), ",") {
			direction := token[:1]
			nbSteps, err := strconv.Atoi(token[1:])
			check(err)

			for i := 0; i < nbSteps; i++ {
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

				if firstWireStepCount, exists := steps[step]; !exists && isFirstWire {
					steps[step] = stepCount

				} else if exists && !isFirstWire {
					// distance := abs(-step.x) + abs(-step.y) // manhattan distance part One
					distance := firstWireStepCount + stepCount // step count distance part Two

					if minDistance > distance {
						minDistance = distance
					}
				}
			}
		}

		isFirstWire = false
	}

	check(scanner.Err())

	fmt.Println("Distance:", minDistance)
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
