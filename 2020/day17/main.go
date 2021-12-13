package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var gridLine []rune

		for _, c := range scanner.Text() {
			gridLine = append(gridLine, c)
		}

		grid = append(grid, gridLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(grid))
	// fmt.Printf("Part two: %d\n", partTwo(tickets, validators))
}

func partOne(grid [][]rune) int {
	fmt.Println(grid)
	return 0
}
