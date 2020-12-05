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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var passports []map[string]string

	passport := map[string]string{}
	fieldRegex := regexp.MustCompile(`([[:lower:]]+):([^ ]+)`)

	scan := bufio.NewScanner(file)

	for {
		hasNext := scan.Scan()
		if !hasNext {
			passports = append(passports, passport)
			break
		}

		line := scan.Text()
		if len(line) == 0 {
			passports = append(passports, passport)
			passport = map[string]string{}
			continue
		}

		matches := fieldRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			passport[match[1]] = match[2]
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(passports))
	fmt.Printf("Part two: %d\n", partTwo(passports))
}

func partOne(passports []map[string]string) int {
	var result int

	for _, passport := range passports {
		if hasRequiredFields(passport) {
			result++
		}
	}

	return result
}

func partTwo(passports []map[string]string) int {
	var result int

	hclRegex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eclRegex := regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	pidRegex := regexp.MustCompile(`^[0-9]{9}$`)

	for _, passport := range passports {
		if !hasRequiredFields(passport) {
			continue
		}

		// byr
		if !matchRange(passport["byr"], 1920, 2002) {
			continue
		}

		// iyr
		if !matchRange(passport["iyr"], 2010, 2020) {
			continue
		}

		// eyr
		if !matchRange(passport["eyr"], 2020, 2030) {
			continue
		}

		// hgt
		hgt := passport["hgt"]
		if !strings.HasSuffix(hgt, "cm") && !strings.HasSuffix(hgt, "in") {
			continue
		}

		switch hgt[len(hgt)-2:] {
		case "cm":
			if !matchRange(hgt[:len(hgt)-2], 150, 193) {
				continue
			}

		case "in":
			if !matchRange(hgt[:len(hgt)-2], 59, 76) {
				continue
			}
		}

		// hcl
		if !hclRegex.MatchString(passport["hcl"]) {
			continue
		}

		// ecl
		if !eclRegex.MatchString(passport["ecl"]) {
			continue
		}

		// pid
		if !pidRegex.MatchString(passport["pid"]) {
			continue
		}

		result++
	}

	return result
}

func hasRequiredFields(passport map[string]string) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range fields {
		if _, exists := passport[field]; !exists {
			return false
		}
	}

	return true
}

func matchRange(value string, min, max int) bool {
	i, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return i >= min && i <= max
}
