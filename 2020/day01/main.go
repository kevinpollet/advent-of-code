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

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		numberStr := scan.Text()

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(numbers))
	fmt.Printf("Part two: %d\n", partTwo(numbers))
}

func partOne(numbers []int) int {
	var result int

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {

			if numbers[i]+numbers[j] == 2020 {
				result = numbers[i] * numbers[j]
				break
			}
		}
	}

	return result
}

func partTwo(numbers []int) int {
	var result int

	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers)-1; j++ {
			for k := j + 1; k < len(numbers); k++ {

				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					result = numbers[i] * numbers[j] * numbers[k]
					break
				}
			}
		}
	}

	return result
}
