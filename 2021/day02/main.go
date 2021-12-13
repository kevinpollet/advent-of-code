package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	name  string
	value int
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var cs []command
	for s.Scan() {
		p := strings.Split(s.Text(), " ")
		if len(p) != 2 {
			panic("malformed input")
		}

		cs = append(cs, command{
			name:  p[0],
			value: atoi(p[1]),
		})
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part one: ", partOne(cs))
	fmt.Println("Part two: ", partTwo(cs))
}

func partOne(cs []command) int {
	var hp, dp int

	for _, c := range cs {
		switch c.name {
		case "forward":
			hp += c.value

		case "up":
			dp -= c.value

		case "down":
			dp += c.value
		}
	}

	return hp * dp
}

func partTwo(cs []command) int {
	var hp, dp, ap int

	for _, c := range cs {
		switch c.name {
		case "forward":
			hp += c.value
			dp += ap * c.value

		case "up":
			ap -= c.value

		case "down":
			ap += c.value
		}
	}

	return hp * dp
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
