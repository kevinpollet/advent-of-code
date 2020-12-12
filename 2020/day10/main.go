package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	adapters := map[int]struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jolt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		adapters[jolt] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(adapters))
	fmt.Printf("Part two: %d\n", partTwo(adapters))
}

func partOne(adapters map[int]struct{}) int {
	oneCount, threeCount := partOneRec(adapters, 0)

	return oneCount * threeCount
}

func partOneRec(adapters map[int]struct{}, joultage int) (int, int) {
	for i := 1; i <= 3; i++ {

		if _, exists := adapters[joultage+i]; exists {
			oneCount, threeCount := partOneRec(adapters, joultage+i)

			if i == 1 {
				oneCount++

			} else if i == 3 {
				threeCount++
			}

			return oneCount, threeCount
		}
	}

	return 0, 1
}

func partTwo(adapters map[int]struct{}) int {
	return partTwoRec(adapters, map[int]int{}, 0)
}

func partTwoRec(adapters map[int]struct{}, visited map[int]int, joultage int) int {
	var count int

	for i := 1; i <= 3; i++ {
		if c, ok := visited[joultage+i]; ok {
			count += c
			continue
		}

		if _, exists := adapters[joultage+i]; exists {
			count += partTwoRec(adapters, visited, joultage+i)
			visited[joultage+i] = count
		}
	}

	if count == 0 {
		return 1
	}
	return count
}
