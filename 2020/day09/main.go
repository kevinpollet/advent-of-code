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

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	invalidNumber := partOne(numbers, 25)

	fmt.Printf("Part one: %d\n", invalidNumber)
	fmt.Printf("Part two: %d\n", partTwo(numbers, invalidNumber))
}

func partOne(numbers []int, pLen int) int {
	for i := pLen; i < len(numbers); i++ {
		if !sumExists(numbers[i-pLen:i], numbers[i]) {
			return numbers[i]
		}
	}

	return 0
}

func sumExists(preamble []int, sum int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == sum {
				return true
			}
		}
	}
	return false
}

func partTwo(numbers []int, invalidNum int) int {
	var result, maxLen int

	for i := 0; i < len(numbers); i++ {
		curLen := 1
		sum := numbers[i]
		min := numbers[i]
		max := numbers[i]

		for j := i + 1; j < len(numbers) && sum < invalidNum; j++ {
			curLen += 1
			sum += numbers[j]

			switch {
			case numbers[j] < min:
				min = numbers[j]

			case numbers[j] > max:
				max = numbers[j]
			}
		}

		if sum == invalidNum && curLen > maxLen {
			maxLen = curLen
			result = min + max
		}
	}

	return result
}
