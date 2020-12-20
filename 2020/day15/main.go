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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, number)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(numbers))
	fmt.Printf("Part two: %d\n", partTwo(numbers))
}

func partOne(numbers []int) int {
	return playGame(numbers, 2020)
}

func partTwo(numbers []int) int {
	return playGame(numbers, 30000000)
}

func playGame(numbers []int, maxTurn int) int {
	lastSpokenNumber := numbers[len(numbers)-1]

	spokenNumbers := map[int]int{}
	for i := 0; i < len(numbers)-1; i++ {
		spokenNumbers[numbers[i]] = i + 1
	}

	for i := len(numbers) + 1; i <= maxTurn; i++ {
		turn, exists := spokenNumbers[lastSpokenNumber]
		if !exists {
			spokenNumbers[lastSpokenNumber] = i - 1
			lastSpokenNumber = 0
			continue
		}

		spokenNumbers[lastSpokenNumber] = i - 1
		lastSpokenNumber = i - 1 - turn
	}

	return lastSpokenNumber
}
