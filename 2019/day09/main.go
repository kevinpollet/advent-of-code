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
	memory := make([]int, len(opCodes)+1000)
	for i, opCode := range opCodes {
		memory[i] = atoi(opCode)
	}

	err = runPrg(memory)
	check(err)
}

func runPrg(memory []int) error {
	reader := bufio.NewScanner(os.Stdin)
	relativeBase := 0
	opArity := map[int]int{1: 3, 2: 3, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 3, 9: 1}

	for cursor := 0; memory[cursor] != 99; {
		instruction := memory[cursor]
		opCode := instruction % 100

		if opCode < 1 || opCode > 9 {
			return fmt.Errorf("Unknown opcode: %d", opCode)
		}

		locs := readLocs(opArity[opCode], instruction, memory, cursor, relativeBase)

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

		// output
		case 4:
			fmt.Println("Output:", memory[locs[0]])

		// jump-if-true
		case 5:
			if memory[locs[0]] != 0 {
				cursor = memory[locs[1]]
				continue
			}

		// jump-if-false
		case 6:
			if memory[locs[0]] == 0 {
				cursor = memory[locs[1]]
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
			relativeBase += memory[locs[0]]
		}

		cursor += opArity[opCode] + 1
	}
	return nil
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
