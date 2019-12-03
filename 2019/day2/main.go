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
	initialMemory := make([]int, len(opCodes))
	for i, strOpCode := range opCodes {
		initialMemory[i], err = strconv.Atoi(strOpCode)
		check(err)
	}

	var memory []int
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			memory = append([]int{}, initialMemory...)
			if err := runPrg(memory, noun, verb); err != nil {
				break
			}

			if memory[0] == 19690720 {
				fmt.Printf("(noun,verb)=(%d, %d)\n", memory[1], memory[2])
				return
			}
		}
	}

	fmt.Println("No result found")
}

func runPrg(memory []int, noun, verb int) error {
	memory[1] = noun
	memory[2] = verb

	for cursor := 0; memory[cursor] != 99; cursor += 4 {
		opCode := memory[cursor]
		if opCode != 1 && opCode != 2 {
			return fmt.Errorf("Unknown opcode: %d", cursor)
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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
