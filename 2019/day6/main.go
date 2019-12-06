package main

import (
	"bufio"
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

	//partOne(orbits)
	partTwo(orbits)
}

func partOne(orbits map[string]string) {
	indirectOrbits := 0

	for _, value := range orbits {

		object, exists := orbits[value]
		for exists {
			indirectOrbits++
			object, exists = orbits[object]
		}
	}

	fmt.Println("Total orbits:", indirectOrbits+len(orbits))
}

// not optimized
func partTwo(orbits map[string]string) {
	youToCOM, sanToCOM := path("YOU", orbits), path("SAN", orbits)

	for yIndex, yObject := range youToCOM {

		for sIndex := indexOf(sanToCOM, yObject); sIndex != -1; {
			fmt.Println("Orbital transfers:", sIndex+yIndex-2)
			return
		}
	}
}

func path(start string, orbitsMap map[string]string) []string {
	path := []string{start}

	for orbited, exists := orbitsMap[start]; exists; {
		path = append(path, orbited)
		orbited, exists = orbitsMap[orbited]
	}
	return path
}

func indexOf(values []string, value string) int {
	for index, v := range values {
		if v == value {
			return index
		}
	}
	return -1
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
