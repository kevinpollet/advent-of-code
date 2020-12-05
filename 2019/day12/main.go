package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type moon struct {
	coordinates [3]int
	velocities  [3]int
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	moons := []*moon{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		moon := &moon{}
		fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &moon.coordinates[0], &moon.coordinates[1], &moon.coordinates[2])

		moons = append(moons, moon)
	}

	check(scanner.Err())

	//result := partOne(clone(moons), 1000)
	result := partTwo(clone(moons))

	fmt.Println(result)
}

func partOne(moons []*moon, steps int) int {
	total := 0

	for i := 0; i < steps; i++ {
		total = 0

		for j, moon := range moons {

			for _, moon2 := range moons[j+1:] {

				for c := 0; c < 3; c++ {
					if moon.coordinates[c] > moon2.coordinates[c] {
						moon.velocities[c]--
						moon2.velocities[c]++
					} else if moon2.coordinates[c] > moon.coordinates[c] {
						moon.velocities[c]++
						moon2.velocities[c]--
					}
				}
			}

			moon.coordinates[0] = moon.coordinates[0] + moon.velocities[0]
			moon.coordinates[1] = moon.coordinates[1] + moon.velocities[1]
			moon.coordinates[2] = moon.coordinates[2] + moon.velocities[2]

			kinetic := abs(moon.velocities[0]) + abs(moon.velocities[1]) + abs(moon.velocities[2])
			potential := abs(moon.coordinates[0]) + abs(moon.coordinates[1]) + abs(moon.coordinates[2])

			total += kinetic * potential
		}
	}

	return total
}

func partTwo(moons []*moon) int {
	count := 1
	revX, revY, revZ := 0, 0, 0
	initState := clone(moons)

	for revX == 0 || revY == 0 || revZ == 0 {

		for j, moon := range moons {

			for _, moon2 := range moons[j+1:] {

				for c := 0; c < 3; c++ {
					if moon.coordinates[c] > moon2.coordinates[c] {
						moon.velocities[c]--
						moon2.velocities[c]++
					} else if moon2.coordinates[c] > moon.coordinates[c] {
						moon.velocities[c]++
						moon2.velocities[c]--
					}
				}
			}

			moon.coordinates[0] = moon.coordinates[0] + moon.velocities[0]
			moon.coordinates[1] = moon.coordinates[1] + moon.velocities[1]
			moon.coordinates[2] = moon.coordinates[2] + moon.velocities[2]
		}

		isRevX, isRevY, isRevZ := revX == 0, revY == 0, revZ == 0
		for m := 0; m < len(moons); m++ {
			isRevX = isRevX && (moons[m].coordinates[0] == initState[m].coordinates[0] && moons[m].velocities[0] == initState[m].velocities[0])
			isRevY = isRevY && (moons[m].coordinates[1] == initState[m].coordinates[1] && moons[m].velocities[1] == initState[m].velocities[1])
			isRevZ = isRevZ && (moons[m].coordinates[2] == initState[m].coordinates[2] && moons[m].velocities[2] == initState[m].velocities[2])
		}

		if isRevX {
			revX = count
		}
		if isRevY {
			revY = count
		}
		if isRevZ {
			revZ = count
		}

		count++
	}

	return lcm(lcm(revX, revY), revZ)
}

func clone(moons []*moon) []*moon {
	result := []*moon{}
	for _, moon := range moons {
		moonCopy := *moon
		result = append(result, &moonCopy)
	}
	return result
}

func lcm(a, b int) int {
	return abs(a*b) / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
