package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	check(err)

	layers := data[:len(data)-1]

	// partOne(layers, 25, 6)
	partTwo(layers, 25, 6)
}

func partOne(layers []byte, layerWidth int, layerHeight int) {
	nbLayers := len(layers) / (layerWidth * layerHeight)

	nbZero, nbOne, nbTwo := int(^uint(0)>>1), 0, 0
	for i := 0; i < nbLayers; i++ {
		start := i * layerWidth * layerHeight
		end := start + layerWidth*layerHeight
		layer := layers[start:end]

		if zeroCount, oneCount, twoCount := countZeroOneAndTwoDigits(layer); nbZero > zeroCount {
			nbZero = zeroCount
			nbOne = oneCount
			nbTwo = twoCount
		}
	}

	fmt.Println("Result:", nbOne*nbTwo)
}

func partTwo(layers []byte, layerWidth int, layerHeight int) {
	nbLayers := len(layers) / (layerWidth * layerHeight)
	image := append([]byte{}, layers[0:layerWidth*layerHeight]...)

	for i := 1; i < nbLayers; i++ {
		start := i * layerWidth * layerHeight
		end := start + layerWidth*layerHeight
		layer := layers[start:end]

		for j := 0; j < len(layer); j++ {
			if image[j] == 50 {
				image[j] = layer[j]
			}
		}
	}

	for i := 0; i < len(image); i++ {
		switch image[i] {
		case 49:
			fmt.Print("*")

		default:
			fmt.Print(" ")
		}

		if (i+1)%layerWidth == 0 {
			fmt.Println()
		}
	}
}

func countZeroOneAndTwoDigits(layer []byte) (int, int, int) {
	nbZero, nbOne, nbTwo := 0, 0, 0

	for _, digit := range layer {
		switch digit {
		case 48:
			nbZero++
		case 49:
			nbOne++
		case 50:
			nbTwo++
		}
	}
	return nbZero, nbOne, nbTwo
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
