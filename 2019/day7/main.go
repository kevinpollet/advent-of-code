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

	//maxSignal := partOne(memory)
	maxSignal := partTwo(memory)

	fmt.Println("Highest signal:", maxSignal)
}

func partOne(memory []int) int {
	maxSignal := 0
	phasesForA := []int{0, 1, 2, 3, 4}

	for _, phaseA := range phasesForA {
		phasesForB := remove(phasesForA, phaseA)
		outputA := runPrg([]int{phaseA, 0}, clone(memory))

		for _, phaseB := range phasesForB {
			phasesForC := remove(phasesForB, phaseB)
			outputB := runPrg([]int{phaseB, outputA}, clone(memory))

			for _, phaseC := range phasesForC {
				phasesForD := remove(phasesForC, phaseC)
				outputC := runPrg([]int{phaseC, outputB}, clone(memory))

				for _, phaseD := range phasesForD {
					phaseE := remove(phasesForD, phaseD)[0]
					outputD := runPrg([]int{phaseD, outputC}, clone(memory))
					signal := runPrg([]int{phaseE, outputD}, clone(memory))

					if signal > maxSignal {
						maxSignal = signal
					}
				}
			}
		}
	}

	return maxSignal
}

func partTwo(memory []int) int {
	maxSignal := 0
	phasesForA := []int{5, 6, 7, 8, 9}

	for _, phaseA := range phasesForA {
		phasesForB := remove(phasesForA, phaseA)

		for _, phaseB := range phasesForB {
			phasesForC := remove(phasesForB, phaseB)

			for _, phaseC := range phasesForC {
				phasesForD := remove(phasesForC, phaseC)

				for _, phaseD := range phasesForD {
					phaseE := remove(phasesForD, phaseD)[0]

					ampA := newAmplifier(phaseA, clone(memory))
					ampB := newAmplifier(phaseB, clone(memory))
					ampC := newAmplifier(phaseC, clone(memory))
					ampD := newAmplifier(phaseD, clone(memory))
					ampE := newAmplifier(phaseE, clone(memory))

					signal := 0
					for {
						outputA := ampA(signal)
						outputB := ampB(outputA)
						outputC := ampC(outputB)
						outputD := ampD(outputC)
						outputE := ampE(outputD)

						if outputE != signal {
							signal = outputE

						} else {
							if signal > maxSignal {
								maxSignal = signal
							}
							break
						}
					}
				}
			}
		}
	}

	return maxSignal
}

func newAmplifier(phase int, memory []int) func(int) int {
	cursor := 0
	lastOutput := 0

	return func(signal int) int {
		inputs := []int{phase, signal}
		if cursor != 0 {
			inputs = []int{signal}
		}

		newOutput, newCursor := runPrgWithCursor(inputs, cursor, memory)

		cursor = newCursor
		if newOutput != -1 {
			lastOutput = newOutput
		}

		return lastOutput
	}
}

func runPrg(inputs []int, memory []int) int {
	output, _ := runPrgWithCursor(inputs, 0, memory)
	return output
}

func runPrgWithCursor(inputs []int, cursor int, memory []int) (int, int) {
	inputIndex := 0
	opArity := map[int]int{1: 3, 2: 3, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 3}

	for memory[cursor] != 99 {
		instruction := memory[cursor]

		opCode := instruction % 100
		if opCode < 1 || opCode > 8 {
			log.Fatalf("Unknown opcode: %d", opCode)
		}

		opCodeArity := opArity[opCode]

		// Read locations
		locs := make([]int, 0)
		str := fmt.Sprintf("%05d", instruction)

		for i := 1; i <= opCodeArity; i++ {
			mode := int(str[len(str)-2-i] - 48)

			switch mode {
			case 1:
				locs = append(locs, cursor+i)

			default:
				locs = append(locs, memory[cursor+i])
			}
		}

		// Execute
		switch opCode {
		// add
		case 1:
			memory[locs[2]] = memory[locs[0]] + memory[locs[1]]

		// multiply
		case 2:
			memory[locs[2]] = memory[locs[0]] * memory[locs[1]]

		// read
		case 3:
			memory[locs[0]] = inputs[inputIndex]
			inputIndex++

		//output
		case 4:
			return memory[locs[0]], cursor + opArity[opCode] + 1

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

		cursor += opArity[opCode] + 1
	}
	return -1, cursor
}

func remove(phases []int, phaseToRemove int) []int {
	var result []int

	for _, phase := range phases {
		if phaseToRemove != phase {
			result = append(result, phase)
		}
	}
	return result
}

func atoi(value string) int {
	intValue, err := strconv.Atoi(value)
	check(err)

	return intValue
}

func btoi(value bool) int {
	if value {
		return 1
	}
	return 0
}

func clone(slice []int) []int {
	return append([]int{}, slice...)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
