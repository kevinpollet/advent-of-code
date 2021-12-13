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
	defer f.Close()

	s := bufio.NewScanner(f)

	var r []string
	for s.Scan() {
		r = append(r, s.Text())
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part one: ", partOne(r))
	fmt.Println("Part two: ", partTwo(r))
}

func partOne(r []string) int {
	var g, e int

	for i := 0; i < len(r[0]); i++ {
		var nz, no int
		for _, b := range r {
			switch b[i] {
			case '0':
				nz++
			case '1':
				no++
			}
		}

		mask := 1 << (len(r[0]) - i - 1)

		if no >= nz {
			g |= mask
		} else {
			e |= mask
		}
	}

	return g * e
}

func partTwo(r []string) int {
	o := findRating(r, func(nz []string, no []string) []string {
		if len(no) >= len(nz) {
			return no
		}
		return nz
	})

	c := findRating(r, func(nz []string, no []string) []string {
		if len(nz) <= len(no) {
			return nz
		}
		return no
	})

	return o * c
}

func findRating(r []string, f func(nz []string, no []string) []string) int {
	var i int
	for len(r) != 1 {
		var z, o []string
		for _, s := range r {
			switch s[i] {
			case '0':
				z = append(z, s)
			case '1':
				o = append(o, s)
			}
		}

		r = f(z, o)
		i++
	}

	return parseBinary(r[0])
}

func parseBinary(b string) int {
	ret, err := strconv.ParseInt(b, 2, 32)
	if err != nil {
		panic(err)
	}

	return int(ret)
}
