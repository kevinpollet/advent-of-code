package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type maskIns struct {
	value string
}

type memIns struct {
	index int
	value int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var program []interface{}

	regex := regexp.MustCompile(`^mem\[(\d+)] = (\d+)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "mask = ")
		if len(parts) == 2 {
			program = append(program, maskIns{parts[1]})
			continue
		}

		matches := regex.FindStringSubmatch(scanner.Text())
		if len(matches) == 3 {
			index, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
			}

			value, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatal(err)
			}

			program = append(program, memIns{index, value})
			continue
		}

		log.Fatal("Malformed instruction")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(program))
	fmt.Printf("Part two: %d\n", partTwo(program))
}

func partOne(program []interface{}) int {
	var mask string

	memory := map[int]int{}

	for _, ins := range program {
		switch i := ins.(type) {
		case maskIns:
			mask = i.value

		case memIns:
			memValue := []rune(fmt.Sprintf("%036b", i.value))
			for i := range memValue {
				if mask[i] == 'X' {
					continue
				}
				memValue[i] = rune(mask[i])
			}

			memory[i.index] = parse(memValue)
		}
	}

	return sum(memory)
}

func partTwo(program []interface{}) int {
	var mask string

	memory := map[int]int{}

	for _, ins := range program {
		switch i := ins.(type) {
		case maskIns:
			mask = i.value

		case memIns:
			address := []rune(fmt.Sprintf("%036b", i.index))
			for i := range address {
				if mask[i] == '0' {
					continue
				}
				address[i] = rune(mask[i])
			}

			write(memory, address, i.value, 0)
		}
	}

	return sum(memory)
}

func write(memory map[int]int, address []rune, value, idx int) {
	if len(address) == idx {
		memory[parse(address)] = value
		return
	}

	if address[idx] != 'X' {
		write(memory, address, value, idx+1)
		return
	}

	write(memory, replace(address, '0', idx), value, idx+1)
	write(memory, replace(address, '1', idx), value, idx+1)
}

func parse(value []rune) int {
	if len(value) == 0 {
		return 0
	}

	res, err := strconv.ParseInt(string(value), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(res)
}

func sum(memory map[int]int) int {
	var sum int
	for _, value := range memory {
		sum += value
	}

	return sum
}

func replace(array []rune, value rune, index int) []rune {
	newArray := append([]rune{}, array...)
	newArray[index] = value

	return newArray
}
