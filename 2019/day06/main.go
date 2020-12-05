package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	reader := bufio.NewScanner(file)

	orbits := make(map[string]string)
	for reader.Scan() {
		objects := strings.Split(reader.Text(), ")")
		orbits[objects[1]] = objects[0]
	}

	check(reader.Err())

	// result := partOne(orbits)
	result, err := partTwo(orbits)
	check(err)

	fmt.Println("Result:", result)
}

func partOne(orbits map[string]string) int {
	indirectOrbits := 0

	for _, value := range orbits {

		object, exists := orbits[value]
		for exists {
			indirectOrbits++
			object, exists = orbits[object]
		}
	}

	return indirectOrbits + len(orbits)
}

func partTwo(orbits map[string]string) (int, error) {
	youToCOM, youSteps, sanSteps := map[string]int{}, 0, 0

	for orbited, exists := orbits["YOU"]; exists; {
		youToCOM[orbited] = youSteps
		orbited, exists = orbits[orbited]
		youSteps++
	}

	for orbited, exists := orbits["SAN"]; exists; {

		if youSteps, intersect := youToCOM[orbited]; intersect {
			return youSteps + sanSteps, nil
		}

		orbited, exists = orbits[orbited]
		sanSteps++
	}

	return 0, errors.New("No path from YOU to SAN")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
