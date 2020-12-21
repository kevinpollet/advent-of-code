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

type validatorFunc func(int) bool

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var (
		tickets    [][]int
		validators []validatorFunc
	)

	regex := regexp.MustCompile(`^[[:lower:] ]+: (\d+)-(\d+) or (\d+)-(\d+)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line == "your ticket:" || line == "nearby tickets:" {
			continue
		}

		matches := regex.FindStringSubmatch(line)
		if len(matches) == 5 {
			validators = append(validators, func(value int) bool {
				return (value >= atoi(matches[1]) && value <= atoi(matches[2])) ||
					(value >= atoi(matches[3]) && value <= atoi(matches[4]))
			})
			continue
		}

		var ticket []int
		for _, n := range strings.Split(line, ",") {
			ticket = append(ticket, atoi(n))
		}

		tickets = append(tickets, ticket)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(tickets, validators))
	fmt.Printf("Part two: %d\n", partTwo(tickets, validators))
}

func partOne(tickets [][]int, validators []validatorFunc) int {
	var sum int

	for i := 1; i < len(tickets); i++ {
		for k := 0; k < len(tickets[i]); k++ {
			if !isValueValid(tickets[i][k], validators...) {
				sum += tickets[i][k]
			}
		}
	}

	return sum
}

func partTwo(tickets [][]int, validators []validatorFunc) int {
	var validTickets [][]int

	for i := 0; i < len(tickets); i++ {
		valid := true

		for k := 0; k < len(tickets[i]); k++ {
			if !isValueValid(tickets[i][k], validators...) {
				valid = false
				break
			}
		}

		if valid {
			validTickets = append(validTickets, tickets[i])
		}
	}

	validPositions := make([][]int, len(validators))

	for k, validator := range validators {
		for i := 0; i < len(tickets[0]); i++ {
			valid := true

			for _, ticket := range validTickets {
				if !isValueValid(ticket[i], validator) {
					valid = false
					break
				}
			}

			if valid {
				validPositions[k] = append(validPositions[k], i)
			}
		}
	}

	result := 1
	length := 1
	positions := map[int]int{}

	for len(positions) != len(validators) {

		for i, validPosition := range validPositions {
			if length != len(validPosition) {
				continue
			}

			for _, p := range validPosition {
				if _, ok := positions[p]; !ok {
					if i < 6 {
						result *= tickets[0][p]
					}

					positions[p] = i
					break
				}
			}
		}

		length++
	}

	return result
}

func isValueValid(value int, validators ...validatorFunc) bool {
	for _, validator := range validators {
		if validator(value) {
			return true
		}
	}

	return false
}

func atoi(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
