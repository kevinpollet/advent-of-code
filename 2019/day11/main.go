package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type pos struct{ x, y int }

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	check(err)

	opCodes := strings.Split(string(data[:len(data)-1]), ",")
	memory := make([]int, len(opCodes)+1000)
	for i, opCode := range opCodes {
		memory[i] = atoi(opCode)
	}

	paintCount, _, err := partOne(0, memory)
	check(err)

	fmt.Println("Result:", paintCount)

	plate, err := partTwo(memory)
	check(err)

	fmt.Println(plate)
}

func partOne(startColor int, memory []int) (int, map[pos]int, error) {
	cursor, relativeBase, count, directionIndex := 0, 0, 0, 0
	position := pos{}
	positions := map[pos]int{}
	directions := [...]pos{pos{0, -1}, pos{1, 0}, pos{0, 1}, pos{-1, 0}}

	for {
		color, exists := positions[position]
		if cursor == 0 {
			color = startColor
		}
		if !exists {
			count++
		}

		result, err := runPrg([]int{color}, &cursor, &relativeBase, memory)
		if err != nil {
			return -1, positions, err
		}

		if len(result) == 0 {
			break
		}

		positions[position] = result[0]

		switch result[1] {
		case 0: // left 90
			if directionIndex == 0 {
				directionIndex = len(directions) - 1
			} else {
				directionIndex--
			}

		case 1: // right 90
			directionIndex = (directionIndex + 1) % len(directions)
		}

		position = pos{
			x: position.x + directions[directionIndex].x,
			y: position.y + directions[directionIndex].y,
		}
	}
	return count, positions, nil
}

func partTwo(memory []int) (string, error) {
	plate := ""

	_, positions, err := partOne(1, memory)
	if err != nil {
		return plate, err
	}

	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if color := positions[pos{x, y}]; color == 1 {
				plate += "*"
			} else {
				plate += " "
			}
		}
		plate += "\n"
	}

	return plate, nil
}

func runPrg(inputs []int, cursor *int, relativeBase *int, memory []int) ([]int, error) {
	opArity := map[int]int{1: 3, 2: 3, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 3, 9: 1}
	outputs := []int{}
	inputIndex := 0

	for memory[*cursor] != 99 {
		instruction := memory[*cursor]
		opCode := instruction % 100

		if opCode < 1 || opCode > 9 {
			return nil, fmt.Errorf("Unknown opcode: %d", opCode)
		}

		locs := readLocs(opArity[opCode], instruction, memory, *cursor, *relativeBase)

		switch opCode {
		// add
		case 1:
			memory[locs[2]] = memory[locs[0]] + memory[locs[1]]

		// multiply
		case 2:
			memory[locs[2]] = memory[locs[0]] * memory[locs[1]]

		// read
		case 3:
			if inputIndex >= len(inputs) {
				return outputs, nil
			}
			memory[locs[0]] = inputs[inputIndex]
			inputIndex++

		// output
		case 4:
			outputs = append(outputs, memory[locs[0]])

		// jump-if-true
		case 5:
			if memory[locs[0]] != 0 {
				*cursor = memory[locs[1]]
				continue
			}

		// jump-if-false
		case 6:
			if memory[locs[0]] == 0 {
				*cursor = memory[locs[1]]
				continue
			}

		// less-than
		case 7:
			memory[locs[2]] = btoi(memory[locs[0]] < memory[locs[1]])

		// equals
		case 8:
			memory[locs[2]] = btoi(memory[locs[0]] == memory[locs[1]])

		// relative base
		case 9:
			*relativeBase += memory[locs[0]]
		}

		*cursor += opArity[opCode] + 1
	}
	return outputs, nil
}

func readLocs(arity int, instruction int, memory []int, cursor int, relativeBase int) []int {
	var locs []int

	str := fmt.Sprintf("%05d", instruction)

	for i := 1; i <= arity; i++ {
		mode := int(str[len(str)-2-i] - 48)

		switch mode {
		case 0:
			locs = append(locs, memory[cursor+i])

		case 1:
			locs = append(locs, cursor+i)

		case 2:
			locs = append(locs, memory[cursor+i]+relativeBase)
		}
	}

	return locs
}

func btoi(value bool) int {
	if value {
		return 1
	}
	return 0
}

func atoi(value string) int {
	intValue, err := strconv.Atoi(value)
	check(err)

	return intValue
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
