package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var answersByGroup [][]string
	var groupAnswers []string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			answersByGroup = append(answersByGroup, groupAnswers)
			groupAnswers = make([]string, 0)
			continue
		}

		groupAnswers = append(groupAnswers, line)
	}

	err = scanner.Err()
	if err == nil && len(groupAnswers) != 0 {
		answersByGroup = append(answersByGroup, groupAnswers)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(answersByGroup))
	fmt.Printf("Part two: %d\n", partTwo(answersByGroup))
}

func partOne(answersByGroup [][]string) int {
	var result int

	for _, groupAnswers := range answersByGroup {
		uniqAnswers := make(map[rune]struct{})

		for _, answers := range groupAnswers {
			for _, answer := range answers {
				uniqAnswers[answer] = struct{}{}
			}
		}

		result += len(uniqAnswers)
	}

	return result
}

func partTwo(answersByGroup [][]string) int {
	var result int

	for _, groupAnswers := range answersByGroup {
		answersCount := make(map[rune]int)

		for _, answers := range groupAnswers {
			for _, answer := range answers {
				answersCount[answer] += 1
			}
		}

		for _, count := range answersCount {
			if count == len(groupAnswers) {
				result++
			}
		}
	}

	return result
}
