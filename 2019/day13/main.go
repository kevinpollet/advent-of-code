package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	check(err)

	opCodes := strings.Split(string(data[:len(data)-1]), ",")
	memory := make([]int, len(opCodes)+1000)
	for i, opCode := range opCodes {
		memory[i] = atoi(opCode)
	}

	//result, err := partOne(memory)
	result, err := partTwo(memory)
	check(err)

	fmt.Println("Result:", result)
}

func partOne(memory []int) (int, error) {
	count, cursor, relativeBase := 0, 0, 0

	game, _, err := runPrg([]int{}, &cursor, &relativeBase, memory)
	if err != nil {
		return -1, err
	}

	for i := 0; i < len(game); i += 3 {
		_, _, t := game[i], game[i+1], game[i+2]

		if t == 2 {
			count++
		}
	}

	return count, nil
}

func partTwo(memory []int) (int, error) {
	jPos := 0
	score := 0
	cursor, relativeBase := 0, 0
	px, bx := 0, 0

	memory[0] = 2

	for {

		game, halted, err := runPrg([]int{jPos}, &cursor, &relativeBase, memory)
		if err != nil {
			return -1, err
		}

		for i := 0; i < len(game); i += 3 {
			x, y, t := game[i], game[i+1], game[i+2]

			if x == -1 && y == 0 {
				score = t

			} else if t == 3 {
				px = x

			} else if t == 4 {
				bx = x
			}
		}

		if halted {
			break
		}

		if bx < px {
			jPos = -1
		} else if bx > px {
			jPos = 1
		} else {
			jPos = 0
		}
	}

	return score, nil
}

func runPrg(inputs []int, cursor *int, relativeBase *int, memory []int) ([]int, bool, error) {
	opArity := map[int]int{1: 3, 2: 3, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 3, 9: 1}
	outputs := []int{}
	inputIndex := 0

	for memory[*cursor] != 99 {
		instruction := memory[*cursor]
		opCode := instruction % 100

		if opCode < 1 || opCode > 9 {
			return nil, true, fmt.Errorf("Unknown opcode: %d", opCode)
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
				return outputs, false, nil
			}
			memory[locs[0]] = inputs[inputIndex]
			inputIndex++

		//output
		case 4:
			outputs = append(outputs, memory[locs[0]])

		//jump-if-true
		case 5:
			if memory[locs[0]] != 0 {
				*cursor = memory[locs[1]]
				continue
			}

		//jump-if-false
		case 6:
			if memory[locs[0]] == 0 {
				*cursor = memory[locs[1]]
				continue
			}

		//less-than
		case 7:
			memory[locs[2]] = btoi(memory[locs[0]] < memory[locs[1]])

		//equals
		case 8:
			memory[locs[2]] = btoi(memory[locs[0]] == memory[locs[1]])

		//relative base
		case 9:
			*relativeBase += memory[locs[0]]
		}

		*cursor += opArity[opCode] + 1
	}
	return outputs, true, nil
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
