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
		memory[i] = toInt(opCode)
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
				fmt.Printf("Result noun:", memory[1], ", verb:", memory[2])
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

		lIdx, rIdx, oIdx := memory[cursor+1], memory[cursor+2], memory[cursor+3]
		if opCode == 1 {
			memory[oIdx] = memory[lIdx] + memory[rIdx]
		} else {
			memory[oIdx] = memory[lIdx] * memory[rIdx]
		}
	}
	return nil
}

func toInt(value string) int {
	parsedValue, err := strconv.Atoi(value)
	check(err)

	return parsedValue
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
