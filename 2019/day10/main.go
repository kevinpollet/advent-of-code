package main

import (
	"bufio"
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

	result = partTwo(asteroids, asteroid)

	fmt.Println("Result:", result)
}

func partOne(asteroids []asteroid) (int, asteroid) {
	max := 0
	asteroid := asteroid{}

	for _, from := range asteroids {
		angles := map[float64]int{}

		for _, to := range asteroids {

			if from != to {
				vx, vy := to.x-from.x, to.y-from.y
				angle := math.Atan2(float64(vy), float64(vx)) + math.Atan2(1, 0)
				if angle < 0 {
					angle += 2 * math.Pi
				}
				if _, exists := angles[angle]; !exists {
					angles[angle] = 0
				}
			}
		}

		if len(angles) > max {
			max = len(angles)
			asteroid = from
		}
	}

	return max, asteroid
}

func partTwo(asteroids []asteroid, from asteroid) int {
	anglesMap := map[float64][]asteroid{}
	angles := []float64{}

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
		min := min(anglesMap[angle])
		anglesMap[angle] = remove(anglesMap[angle], min)

		if i == 199 {
			return min.x*100 + min.y
		}
	}
	return -1
}

func min(asteroids []asteroid) asteroid {
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
