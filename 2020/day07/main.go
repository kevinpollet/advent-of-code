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

const shinyGoldBagType = "shiny gold"

type rule struct {
	n       int
	bagType string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rulesByType := map[string][]rule{}
	regex := regexp.MustCompile(`([0-9]) ([a-z ]+) bags?`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " bags contain ")
		if len(parts) != 2 {
			log.Fatal("Malformed bag rules")
		}

		var rules []rule

		matches := regex.FindAllStringSubmatch(parts[1], -1)
		for _, match := range matches {
			n, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}

			rules = append(rules, rule{n: n, bagType: match[2]})
		}

		rulesByType[parts[0]] = rules
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(rulesByType))
	fmt.Printf("Part two: %d\n", partTwo(rulesByType))
}

func partOne(rulesByType map[string][]rule) int {
	var result int

	for bagType := range rulesByType {
		if bagType == shinyGoldBagType {
			continue
		}

		partOneRec(rulesByType, bagType, &result)
	}

	return result
}

func partOneRec(rulesByType map[string][]rule, bagType string, count *int) bool {
	if bagType == shinyGoldBagType {
		*count++
		return true
	}

	for _, rule := range rulesByType[bagType] {
		if partOneRec(rulesByType, rule.bagType, count) {
			return true
		}
	}

	return false
}

func partTwo(rulesByType map[string][]rule) int {
	return partTwoRec(rulesByType, shinyGoldBagType) - 1
}

func partTwoRec(rulesByType map[string][]rule, bagType string) int {
	rules := rulesByType[bagType]
	if len(rules) == 0 {
		return 1
	}

	var count int

	for _, rule := range rules {
		count += rule.n * partTwoRec(rulesByType, rule.bagType)
	}

	return 1 + count
}
