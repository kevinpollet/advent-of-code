package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		mass := toInt(scanner.Text())

		//partOne(mass, &sum)
		partTwo(mass, &sum)
	}

	check(scanner.Err())

	fmt.Println("Fuel requirements:", sum)
}

func partOne(mass int, sum *int) {
	*sum += mass/3 - 2
}

func partTwo(mass int, sum *int) {
	*sum += partTwoRec(mass, 0)
}

func partTwoRec(value, sum int) int {
	fuel := value/3 - 2
	if fuel <= 0 {
		return sum
	}

	return partTwoRec(fuel, sum+fuel)
}

func toInt(value string) int {
	parsedValue, err := strconv.Atoi(value)
	check(err)

	return parsedValue
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
