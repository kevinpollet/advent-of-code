package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type instruction struct {
	dir   rune
	value int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var instructions []instruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		instructions = append(instructions, instruction{
			dir:   rune(line[0]),
			value: value,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(instructions))
	fmt.Printf("Part two: %d\n", partTwo(instructions))
}

func partOne(instructions []instruction) int {
	sDir := []int{'E', 'S', 'W', 'N'}
	sPos := map[rune]int{'N': 0, 'E': 0, 'S': 0, 'W': 0}

	for _, i := range instructions {
		switch i.dir {
		case 'F':
			sPos[rune(sDir[0])] += i.value

		case 'R':
			sDir = rotate(sDir, -i.value)

		case 'L':
			sDir = rotate(sDir, i.value)

		default:
			sPos[i.dir] += i.value
		}
	}

	return abs(sPos['E']-sPos['W']) + abs(sPos['N']-sPos['S'])
}

func partTwo(instructions []instruction) int {
	sPos := []int{0, 0, 0, 0}  // N, E, S, W
	wPos := []int{1, 10, 0, 0} // N, E, S, W
	wPosIdx := map[rune]int{'N': 0, 'E': 1, 'S': 2, 'W': 3}

	for _, i := range instructions {
		switch i.dir {
		case 'F':
			for k := 0; k < len(sPos); k++ {
				sPos[k] += i.value * wPos[k]
			}

		case 'R':
			wPos = rotate(wPos, i.value)

		case 'L':
			wPos = rotate(wPos, -i.value)

		default:
			wPos[wPosIdx[i.dir]] += i.value
		}
	}

	return abs(sPos[1]-sPos[3]) + abs(sPos[0]-sPos[2])
}

func rotate(src []int, degree int) []int {
	result := make([]int, len(src))
	for i, v := range src {
		idx := (i + degree/90) % len(src)
		if idx < 0 {
			idx = len(src) + idx
		}

		result[idx] = v
	}
	return result
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
