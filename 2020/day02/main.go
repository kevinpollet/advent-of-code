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

type password struct {
	min      int
	max      int
	letter   string
	password string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var passwords []password

	regex := regexp.MustCompile(`^(\d+)-(\d+) ([[:lower:]]): ([[:lower:]]+)$`)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()

		matches := regex.FindStringSubmatch(line)
		if len(matches) == 0 {
			log.Fatal("malformed password line")
		}

		passwords = append(passwords, password{
			min:      atoi(matches[1]),
			max:      atoi(matches[2]),
			letter:   matches[3],
			password: matches[4],
		})

	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(passwords))
	fmt.Printf("Part two: %d\n", partTwo(passwords))
}

func partOne(passwords []password) int {
	var result int

	for _, p := range passwords {
		count := strings.Count(p.password, p.letter)

		if count >= p.min && count <= p.max {
			result++
		}
	}

	return result
}

func partTwo(passwords []password) int {
	var result int

	for _, p := range passwords {
		if p.min < 1 || p.min > len(p.password) {
			continue
		}

		if p.max < 1 || p.max > len(p.password) {
			continue
		}

		firstMatch := string(p.password[p.min-1]) == p.letter
		secondMatch := string(p.password[p.max-1]) == p.letter

		if firstMatch != secondMatch {
			result++
		}
	}

	return result
}

func atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
