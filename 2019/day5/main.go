package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	check(err)

	opCodes := strings.Split(string(data[:len(data)-1]), ",")
	memory := make([]int, len(opCodes))
	for i, opCode := range opCodes {
		memory[i] = atoi(opCode)
	}

	runIntcodePrg(memory)
}

func runIntcodePrg(memory []int) error {
	cursor := 0
	reader := bufio.NewScanner(os.Stdin)
	opArity := map[int]int{1: 3, 2: 3, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 3}

	for memory[cursor] != 99 {
		instruction := memory[cursor]
		opCode := instruction % 100

		if opCode < 1 || opCode > 8 {
			return fmt.Errorf("Unknown opcode: %d", opCode)
		}

		locs := readLocs(opArity[opCode], instruction, memory, cursor)

		switch opCode {
		// add
		case 1:
			memory[locs[2]] = memory[locs[0]] + memory[locs[1]]

		// multiply
		case 2:
			memory[locs[2]] = memory[locs[0]] * memory[locs[1]]

		// read
		case 3:
			fmt.Print("Input: ")
			reader.Scan()
			memory[locs[0]] = atoi(reader.Text())

		//output
		case 4:
			fmt.Println("Output:", memory[locs[0]])

		//jump-if-true
		case 5:
			if memory[locs[0]] != 0 {
				cursor = memory[locs[1]]
				continue
			}

		//jump-if-false
		case 6:
			if memory[locs[0]] == 0 {
				cursor = memory[locs[1]]
				continue
			}

		//less-than
		case 7:
			memory[locs[0]] = btoi(memory[locs[0]] < memory[locs[1]])

		//equals
		case 8:
			memory[locs[0]] = btoi(memory[locs[0]] == memory[locs[1]])
		}

		cursor += opArity[opCode] + 1
	}
	return nil
}

func readLocs(arity int, instruction int, memory []int, cursor int) []int {
	var locs []int
	str := strconv.Itoa(instruction)

	for i := 1; i <= arity; i++ {
		mode := 0
		if len(str) >= i+2 {
			mode = int(str[len(str)-2-i] - 48)
		}

		if mode == 1 {
			locs = append(locs, cursor+i)
		} else {
			locs = append(locs, memory[cursor+i])
		}
	}

	return locs
}

func atoi(value string) int {
	parsedValue, err := strconv.Atoi(value)
	check(err)

	return parsedValue
}

func btoi(value bool) int {
	if value {
		return 1
	}
	return 0
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
