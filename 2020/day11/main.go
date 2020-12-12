package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type adjCounterFunc func([][]rune, int, int) int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var seatLayout [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var rowLayout []rune
		for _, p := range scanner.Text() {
			rowLayout = append(rowLayout, p)
		}

		seatLayout = append(seatLayout, rowLayout)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(seatLayout))
	fmt.Printf("Part two: %d\n", partTwo(seatLayout))
}

func partOne(seatLayout [][]rune) int {
	return simulate(seatLayout, partOneAdjCounter, 4)
}

func partOneAdjCounter(seatLayout [][]rune, i, j int) int {
	var count int

	for k := i - 1; k <= i+1; k++ {
		if k < 0 || k >= len(seatLayout) {
			continue
		}

		for l := j - 1; l <= j+1; l++ {
			if k == i && j == l {
				continue
			}

			if l < 0 || l >= len(seatLayout[k]) {
				continue
			}

			if seatLayout[k][l] == '#' {
				count++
			}
		}
	}

	return count
}

func partTwo(seatLayout [][]rune) int {
	return simulate(seatLayout, partTwoAdjCounter, 5)
}

func partTwoAdjCounter(seatLayout [][]rune, i, j int) int {
	var count int

	for _, s := range []int{-1, 1} {
		for k := i + s; k >= 0 && k < len(seatLayout); k = k + s {
			if seatLayout[k][j] == 'L' {
				break
			}

			if seatLayout[k][j] == '#' {
				count++
				break
			}
		}

		for l := j + s; l >= 0 && l < len(seatLayout[i]); l = l + s {
			if seatLayout[i][l] == 'L' {
				break
			}

			if seatLayout[i][l] == '#' {
				count++
				break
			}
		}
	}

	for _, ks := range []int{-1, 1} {
		for _, js := range []int{-1, 1} {
			k, l := i, j

			for {
				k += ks
				l += js

				if k < 0 || k >= len(seatLayout) ||
					l < 0 || l >= len(seatLayout[k]) {
					break
				}

				if seatLayout[k][l] == 'L' {
					break
				}

				if seatLayout[k][l] == '#' {
					count++
					break
				}
			}
		}
	}

	return count
}

func simulate(seatLayout [][]rune, adjCounter adjCounterFunc, maxAdjCount int) int {
	for {
		newSeatLayout := make([][]rune, len(seatLayout))
		for i := 0; i < len(seatLayout); i++ {

			newRowLayout := make([]rune, len(seatLayout[i]))
			for j := 0; j < len(seatLayout[i]); j++ {
				switch seatLayout[i][j] {
				case 'L':
					if adjCounter(seatLayout, i, j) == 0 {
						newRowLayout[j] = '#'
						continue
					}

				case '#':
					if adjCounter(seatLayout, i, j) >= maxAdjCount {
						newRowLayout[j] = 'L'
						continue
					}
				}

				newRowLayout[j] = seatLayout[i][j]
			}

			newSeatLayout[i] = newRowLayout
		}

		if equal, count := isSameLayout(seatLayout, newSeatLayout); equal {
			return count
		}

		seatLayout = newSeatLayout
	}
}

func isSameLayout(one, two [][]rune) (bool, int) {
	var count int

	for i := 0; i < len(one); i++ {
		for j := 0; j < len(one[i]); j++ {
			if one[i][j] != two[i][j] {
				return false, 0
			}
			if one[i][j] == '#' {
				count++
			}
		}
	}

	return true, count
}
