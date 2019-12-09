package main

import (
	"bufio"
	"errors"
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

	memory := []int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		opCodes := strings.Split(scanner.Text(), ",")

		for _, opCode := range opCodes {
			digit, err := strconv.Atoi(opCode)
			check(err)

			memory = append(memory, digit)
		}
	}

	check(scanner.Err())

	//result, err := partOne(memory, 12, 2)
	result, err := partTwo(memory, 19690720)
	check(err)

	fmt.Println("Result:", result)

}

func partOne(memory []int, noun int, verb int) (int, error) {
	if err := runPrg(memory, noun, verb); err != nil {
		return -1, err
	}

	return memory[0], nil
}

func partTwo(memory []int, expectedOutput int) (int, error) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {

			if result, err := partOne(clone(memory), noun, verb); err != nil {
				return -1, err

			} else if result == expectedOutput {
				return 100*noun + verb, nil
			}
		}
	}
	return -1, errors.New("No solution")
}

func runPrg(memory []int, noun, verb int) error {
	memory[1], memory[2] = noun, verb

	for cursor := 0; memory[cursor] != 99; cursor += 4 {
		opCode := memory[cursor]
		if opCode < 1 || opCode > 2 {
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

func clone(memory []int) []int {
	return append([]int{}, memory...)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
