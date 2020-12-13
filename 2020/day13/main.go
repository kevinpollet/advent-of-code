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

type bus struct {
	id     int
	offset int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var (
		depTimestamp int
		buses        []bus
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		switch len(parts) {
		case 1:
			depTimestamp, err = strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
		default:
			for i, part := range parts {
				if part == "x" {
					continue
				}

				id, err := strconv.Atoi(part)
				if err != nil {
					log.Fatal(id)
				}

				buses = append(buses, bus{
					id:     id,
					offset: i,
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one: %d\n", partOne(buses, depTimestamp))
}

func partOne(buses []bus, depTimestamp int) int {
	var firstBusID int
	firstBusTS := math.MaxInt64

	for _, bus := range buses {
		nextTS := bus.id * int(math.Ceil(float64(depTimestamp)/float64(bus.id)))
		if nextTS < firstBusTS {
			firstBusID = bus.id
			firstBusTS = nextTS
		}
	}

	return firstBusID * (firstBusTS - depTimestamp)
}
