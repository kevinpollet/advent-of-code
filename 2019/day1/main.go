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
		mass, _ := strconv.Atoi(scanner.Text())
		sum += fuelRec(mass, 0)
	}

	check(scanner.Err())

	fmt.Printf("Fuel requirements: %d\n", sum)
}

func fuelRec(value, sum int) int {
	fuel := value/3 - 2
	if fuel <= 0 {
		return sum
	}

	return fuelRec(fuel, sum+fuel)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
