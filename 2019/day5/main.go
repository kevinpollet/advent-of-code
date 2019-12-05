package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	reader := bufio.NewReader(file)
	bytes, _, err := reader.ReadLine()
	check(err)

	opCodes := strings.Split(string(bytes), ",")

	memory := make([]int, len(opCodes))
	for i, opCode := range opCodes {
		memory[i] = atoi(opCode)
	}

	runIntcodePrg(memory)
}

func runIntcodePrg(memory []int) error {
	for cursor := 0; memory[cursor] != 99; {
		instruction := memory[cursor]

		opCode := instruction % 100
		if opCode < 1 || opCode > 8 {
			return fmt.Errorf("Unknown opcode: %d", opCode)
		}

		switch opCode {
		case 1: // add
			p1, p2 := readParam(1, instruction, memory, cursor), readParam(2, instruction, memory, cursor)

			writeParamValue(3, p1+p2, instruction, memory, cursor)

			cursor += 4

		case 2: // multiply
			p1, p2 := readParam(1, instruction, memory, cursor), readParam(2, instruction, memory, cursor)

			writeParamValue(3, p1*p2, instruction, memory, cursor)

			cursor += 4

		case 3: // read
			reader := bufio.NewScanner(os.Stdin)
			fmt.Print("Input: ")

			reader.Scan()
			value := atoi(reader.Text())

			writeParamValue(1, value, instruction, memory, cursor)

			cursor += 2

		case 4: //output
			p1 := readParam(1, instruction, memory, cursor)
			fmt.Println("Output:", p1)

			cursor += 2

		case 5: //jump-if-true
			p1 := readParam(1, instruction, memory, cursor)
			if p1 != 0 {
				cursor = readParam(2, instruction, memory, cursor)
			} else {
				cursor += 3
			}

		case 6: //jump-if-false
			p1 := readParam(1, instruction, memory, cursor)
			if p1 == 0 {
				cursor = readParam(2, instruction, memory, cursor)
			} else {
				cursor += 3
			}

		case 7: //less-than
			p1, p2 := readParam(1, instruction, memory, cursor), readParam(2, instruction, memory, cursor)

			writeParamValue(1, btoi(p1 < p2), instruction, memory, cursor)

			cursor += 4

		case 8: //equals
			p1, p2 := readParam(1, instruction, memory, cursor), readParam(2, instruction, memory, cursor)

			writeParamValue(1, btoi(p1 == p2), instruction, memory, cursor)

			cursor += 4
		}

	}
	return nil
}

func readParam(index int, instruction int, memory []int, cursor int) int {
	mode := mode(index, instruction)
	if mode == 1 {
		return memory[cursor+index]
	}
	return memory[memory[cursor+index]]
}

func writeParamValue(paramIndex int, value int, instruction int, memory []int, cursor int) {
	mode := mode(paramIndex, instruction)
	if mode == 1 {
		memory[cursor+paramIndex] = value
	}
	memory[memory[cursor+paramIndex]] = value
}

func mode(paramIndex int, instruction int) int {
	str := strconv.Itoa(instruction)

	if len(str) >= paramIndex+2 {
		return int(str[len(str)-2-paramIndex] - 48)
	}
	return 0
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
