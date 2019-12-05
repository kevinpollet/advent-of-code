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
	memory := make([]int, len(opCodes))
	for i, opCode := range opCodes {
		memory[i] = atoi(opCode)
	}

	// partOne(memory)
	partTwo(memory)
}

func partOne(memory []int) {
	err := runIntcodePrg(memory, 12, 2)
	check(err)

	fmt.Println("Result:", memory[0])
}

func partTwo(initialMemory []int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			memory := append([]int{}, initialMemory...) // reset memory
			runIntcodePrg(memory, noun, verb)
			if memory[0] == 19690720 {
				fmt.Println("Result noun:", memory[1], "verb:", memory[2])
				return
			}
		}
	}
}

func runIntcodePrg(memory []int, noun, verb int) error {
	memory[1] = noun
	memory[2] = verb

	for cursor := 0; memory[cursor] != 99; cursor += 4 {
		opCode := memory[cursor]
		
		if opCode != 1 && opCode != 2 {
			return fmt.Errorf("Unknown opcode: %d", opCode)
		}

		l, r, o := memory[cursor+1], memory[cursor+2], memory[cursor+3]

		switch opCode {
		case 1:
			memory[o] = memory[l] + memory[r]

		case 2:
			memory[o] = memory[l] * memory[r]
		}
	}
	return nil
}

func atoi(value string) int {
	parsedValue, err := strconv.Atoi(value)
	check(err)

	return parsedValue
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
