package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type asteroid struct {
	x, y     int
	distance float64
}

type void struct{}

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	asteroids := []asteroid{}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, value := range line {
			if string(value) == "#" {
				asteroids = append(asteroids, asteroid{x: x, y: y})
			}
		}
		y++
	}

	check(scanner.Err())

	result, asteroid := partOne(asteroids)

	fmt.Printf("Max asteroids detected: %d, position: (%d,%d)\n", result, asteroid.x, asteroid.y)

	result, err = partTwo(asteroids, asteroid)
	check(err)

	fmt.Println("Result:", result)
}

func partOne(asteroids []asteroid) (int, asteroid) {
	max, position := 0, asteroid{}

	for _, from := range asteroids {
		angles := map[float64]void{}

		for _, to := range asteroids {

			if from != to {
				vx, vy := to.x-from.x, to.y-from.y
				angle := math.Atan2(float64(vy), float64(vx)) + math.Atan2(1, 0)
				if angle < 0 {
					angle += 2 * math.Pi
				}

				if _, exists := angles[angle]; !exists {
					angles[angle] = void{}
				}
			}
		}

		if len(angles) > max {
			max = len(angles)
			position = from
		}
	}

	return max, position
}

func partTwo(asteroids []asteroid, from asteroid) (int, error) {
	anglesMap, angles := map[float64][]asteroid{}, []float64{}

	for _, asteroid := range asteroids {

		if from != asteroid {
			vx, vy := asteroid.x-from.x, asteroid.y-from.y
			angle := math.Atan2(float64(vy), float64(vx)) + math.Atan2(1, 0)
			if angle < 0 {
				angle += 2 * math.Pi
			}

			asteroid.distance = math.Sqrt(float64(vx*vx + vy*vy))

			value, exists := anglesMap[angle]

			anglesMap[angle] = append(value, asteroid)
			if !exists {
				angles = append(angles, angle)
			}
		}
	}

	sort.Float64s(angles)

	for i, angle := range angles {
		asteroid := closest(anglesMap[angle])
		anglesMap[angle] = remove(anglesMap[angle], asteroid)

		if i == 199 {
			return asteroid.x*100 + asteroid.y, nil
		}
	}
	return -1, errors.New("No result")
}

func closest(asteroids []asteroid) asteroid {
	min := asteroids[0]
	for i := 1; i < len(asteroids); i++ {
		if asteroids[i].distance < min.distance {
			min = asteroids[i]
		}
	}
	return min
}

func remove(asteroids []asteroid, a asteroid) []asteroid {
	result := []asteroid{}
	for _, asteroid := range asteroids {
		if asteroid != a {
			result = append(result, asteroid)
		}
	}
	return result
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
