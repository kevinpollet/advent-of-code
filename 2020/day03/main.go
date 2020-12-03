package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type slope struct {
	right int
	down  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var grid []string

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		grid = append(grid, line)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(grid))
	fmt.Printf("Part two: %d\n", partTwo(grid))
}

func partOne(grid []string) int {
	return traverse(
		slope{right: 3, down: 1},
		grid,
	)
}

func partTwo(grid []string) int {
	result := 1
	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	for _, slope := range slopes {
		result *= traverse(slope, grid)
	}

	return result
}

func traverse(s slope, grid []string) int {
	var i, result int

	for j := s.down; j < len(grid); j += s.down {
		i = (i + s.right) % len(grid[j])

		if grid[j][i] == '#' {
			result++
		}
	}

	return result
}
