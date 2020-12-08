package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	op  string
	arg int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	regex := regexp.MustCompile(`([a-z]+) ([-+][0-9]+)`)

	var bootCode []instruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := regex.FindStringSubmatch(scanner.Text())
		if len(matches) != 3 {
			log.Fatal("Malformed instruction")
		}

		arg, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal(err)
		}

		bootCode = append(bootCode, instruction{matches[1], arg})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(bootCode))
	fmt.Printf("Part two: %d\n", partTwo(bootCode))
}

func partOne(bootCode []instruction) int {
	var acc int

	run(bootCode, map[int]struct{}{}, 0, &acc)

	return acc
}

func partTwo(bootCode []instruction) int {
	for i, ins := range bootCode {
		if ins.op == "acc" {
			continue
		}

		var acc int

		if run(switchInstruction(bootCode, i), map[int]struct{}{}, 0, &acc) {
			return acc
		}
	}

	return 0
}

func run(bootCode []instruction, visited map[int]struct{}, index int, acc *int) bool {
	if index == len(bootCode) {
		return true
	}

	if _, exists := visited[index]; exists {
		return false
	}

	visited[index] = struct{}{}

	ins := bootCode[index]

	switch ins.op {
	case "nop":
		return run(bootCode, visited, index+1, acc)

	case "jmp":
		return run(bootCode, visited, index+ins.arg, acc)

	case "acc":
		*acc += ins.arg
		return run(bootCode, visited, index+1, acc)

	default:
		log.Fatal("Unsupported operation")
		return false
	}
}

func switchInstruction(bootCode []instruction, index int) []instruction {
	result := make([]instruction, len(bootCode))
	copy(result, bootCode)

	if result[index].op == "op" {
		result[index].op = "jmp"

	} else if result[index].op == "jmp" {
		result[index].op = "nop"
	}

	return result
}
