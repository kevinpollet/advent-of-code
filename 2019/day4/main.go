package main

import (
	"fmt"
	"strconv"
)

func main() {
	passwordCount := 0
	passwordInterval := [2]int{108457, 562041}

	for i := passwordInterval[0]; i <= passwordInterval[1]; i++ {
		//if matchPartOne(i) {
		if matchPartTwo(i) {
			passwordCount++
		}
	}

	fmt.Println("Result:", passwordCount)
}

func matchPartOne(value int) bool {
	number := strconv.Itoa(value)
	if len(number) != 6 {
		return false
	}

	hasDouble := false

	for i := 0; i < len(number)-1; i++ {
		if number[i] > number[i+1] {
			return false
		}
		hasDouble = hasDouble || number[i] == number[i+1]
	}

	return hasDouble
}

func matchPartTwo(value int) bool {
	number := strconv.Itoa(value)
	if len(number) != 6 {
		return false
	}

	hasDouble := false
	currentDouble := "🎄"

	for i := 0; i < len(number); i++ {
		if i < len(number)-1 && number[i] > number[i+1] {
			return false
		}

		if !hasDouble {
			if currentDouble[0] == number[i] {
				currentDouble += string(number[i])

			} else {
				hasDouble = len(currentDouble) == 2
				currentDouble = string(number[i])
			}
		}
	}

	hasDouble = hasDouble || len(currentDouble) == 2 // in case the password ends with a double

	return hasDouble
}
