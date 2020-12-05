package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var boardingPasses []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boardingPasses = append(boardingPasses, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(boardingPasses))
	fmt.Printf("Part two: %d\n", partTwo(boardingPasses))
}

func partOne(boardingPasses []string) int {
	var max int

	for _, boardingPass := range boardingPasses {
		if id := calcSeatID(boardingPass); id > max {
			max = id
		}
	}

	return max
}

func partTwo(boardingPasses []string) int {
	var seats []int
	for _, boardingPass := range boardingPasses {
		seats = append(seats, calcSeatID(boardingPass))
	}

	sort.Ints(seats)

	for i, id := range seats {
		if i+1 == len(seats) {
			continue
		}

		if seats[i+1] == id+2 {
			return id + 1
		}
	}

	return 0
}

func calcSeatID(boardingPass string) int {
	rowL, rowU := 0, 127
	colL, colU := 0, 7

	for k := 0; k < len(boardingPass); k++ {
		switch boardingPass[k] {
		// lower half
		case 'F':
			rowU = rowL + (rowU-rowL)/2
		case 'L':
			colU = colL + (colU-colL)/2

		// upper half
		case 'B':
			rowL = rowU - (rowU-rowL)/2
		case 'R':
			colL = colU - (colU-colL)/2
		}
	}

	return rowL*8 + colL
}
