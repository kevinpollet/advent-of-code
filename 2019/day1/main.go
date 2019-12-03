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

		//partyOne(mass, &sum)
		partyTwo(mass, &sum)
	}

	check(scanner.Err())

	fmt.Printf("Fuel requirements: %d\n", sum)
}

func partyOne(mass int, sum *int) {
	*sum += mass/3 - 2
}

func partyTwo(mass int, sum *int) {
	*sum += partyTwoRec(mass, 0)
}

func partyTwoRec(value, sum int) int {
	fuel := value/3 - 2
	if fuel <= 0 {
		return sum
	}

	return partyTwoRec(fuel, sum+fuel)
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
