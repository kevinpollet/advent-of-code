package main

import (
	"bufio"
	"fmt"
	"io"
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
		memory[i], err = strconv.Atoi(opCode)
		check(err)
	}

	fmt.Print("Input: ")
	err = runPrg(memory, bufio.NewReader(os.Stdin), os.Stdout)
	check(err)
}

func runPrg(memory []int, input io.Reader, output io.Writer) error {
	reader := bufio.NewScanner(input)
	arity := map[int]int{1: 3, 2: 3, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 3}

	for cursor := 0; memory[cursor] != 99; {
		instruction := memory[cursor]
		opCode := instruction % 100

		if opCode < 1 || opCode > 8 {
			return fmt.Errorf("Unknown opcode: %d", opCode)
		}

		locs := readLocs(arity[opCode], instruction, memory, cursor)

		switch opCode {
		// add
		case 1:
			memory[locs[2]] = memory[locs[0]] + memory[locs[1]]

		// multiply
		case 2:
			memory[locs[2]] = memory[locs[0]] * memory[locs[1]]

		// read
		case 3:
			for reader.Scan() {
				input, err := strconv.Atoi(reader.Text())
				if err != nil {
					return err
				}

				memory[locs[0]] = input
				break
			}

			if reader.Err() != nil {
				return reader.Err()
			}

		//output
		case 4:
			output.Write([]byte(strconv.Itoa(memory[locs[0]])))
			output.Write([]byte("\n"))

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
			memory[locs[2]] = btoi(memory[locs[0]] < memory[locs[1]])

		//equals
		case 8:
			memory[locs[2]] = btoi(memory[locs[0]] == memory[locs[1]])
		}

		cursor += arity[opCode] + 1
	}
	return nil
}

func readLocs(arity int, instruction int, memory []int, cursor int) []int {
	var locs []int

	str := fmt.Sprintf("%05d", instruction)

	for i := 1; i <= arity; i++ {
		mode := int(str[len(str)-2-i] - 48)
		if mode == 1 {
			locs = append(locs, cursor+i)
		} else {
			locs = append(locs, memory[cursor+i])
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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
