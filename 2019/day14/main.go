package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type elt struct {
	n string
	q int
}

func (e *elt) String() string {
	return fmt.Sprintf("(%s, %d)", e.n, e.q)
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	tree := map[string][]*elt{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=>")

		output := parseElt(parts[1])
		tree[output.n] = append(tree[output.n], output)

		for _, input := range strings.Split(parts[0], ",") {
			tree[output.n] = append(tree[output.n], parseElt(input))
		}
	}

	fmt.Println("Parse End")

	check(scanner.Err())

	count := 0
	extras := map[string]int{}
	ore := 1000000000000

	// Part One
	produce("FUEL", 1, tree, &count, extras)

	fmt.Println("Result:", count)

	// Part Two
	max := 0
	lb, ub := 1, 1000000000000

	for {
		count = 0
		extras = map[string]int{}
		delta := (ub - lb) / 2

		produce("FUEL", lb+delta, tree, &count, extras)

		if count > ore {
			ub = lb + delta

		} else if count <= ore {
			lb = lb + delta

			if count == ore || delta == 0 {
				max = lb + delta
				break
			}
		}
	}

	fmt.Println("Result:", max)
}

func produce(name string, amount int, tree map[string][]*elt, count *int, extras map[string]int) {
	values := tree[name]

	if value := extras[name]; value > amount {
		extras[name] -= amount
		amount = 0

	} else if value < amount {
		extras[name] = 0
		amount -= value

	} else {
		extras[name] = 0
		amount = 0
	}

	factor := int(math.Ceil(float64(amount) / float64(values[0].q)))

	if factor*values[0].q > amount {
		extras[name] += factor*values[0].q - amount
	}

	for _, child := range values[1:] {
		if child.n == "ORE" {
			*count += factor * child.q
			return
		}

		produce(child.n, factor*child.q, tree, count, extras)
	}
}

func parseElt(str string) *elt {
	parts := strings.Split(strings.TrimSpace(str), " ")
	return &elt{parts[1], atoi(parts[0])}
}

func atoi(value string) int {
	intValue, err := strconv.Atoi(value)
	check(err)

	return intValue
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
