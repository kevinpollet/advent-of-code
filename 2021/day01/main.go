package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	var d []int
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		d = append(d, i)
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part one: ", partOne(d))
	fmt.Println("Part one: ", partTwo(d))

}

func partOne(d []int) int {
	var ret int
	for i := 1; i < len(d); i++ {
		if d[i] > d[i-1] {
			ret++
		}
	}

	return ret
}

func partTwo(d []int) int {
	var ret int
	for i := 0; i < len(d)-3; i++ {
		c := d[i] + d[i+1] + d[i+2]
		n := d[i+1] + d[i+2] + d[i+3]

		if n > c {
			ret++
		}
	}

	return ret
}
