package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const ColorBlack int = 0
const ColorWhite int = 1
const ColorTransparent int = 2

func getFingerPrint(image string, width int, height int) int {
	var layerSize = width * height
	var layersCount = len(image) / layerSize
	var runes = []rune(image)

	var min0 = 9999999
	var currentHash = 0

	for i:=0; i<layersCount; i++ {

		var digitsMap = make(map[string]int)

		for j:=0; j < layerSize; j++ {
			var currentIndex = layerSize * i + j
			var key = string(runes[currentIndex])
			digitsMap[ key ] = digitsMap[ key ] + 1
		}

		if digitsMap["0"] < min0 {
			min0 = digitsMap["0"]
			currentHash = digitsMap["1"] * digitsMap["2"]
		}

	}

	return currentHash
}

func decodeImage(image string, width int, height int) []int {

	var layerSize = width * height
	var layersCount = len(image) / layerSize

	// create and allocate result
	var decodedImage = make([]int, layerSize)
	for i:=0; i<layerSize; i++ {
		decodedImage[i] = ColorTransparent
	}

	var runes = []rune(image)

	for i:=0; i<layersCount; i++ {
		for j := 0; j < layerSize; j++ {
			var currentIndex = layerSize*i + j
			thisPixel, _ := strconv.Atoi(string(runes[currentIndex]))
			if decodedImage[j] == ColorTransparent && thisPixel != ColorTransparent {
				decodedImage[j] = thisPixel
			}
		}
	}

	return decodedImage
}

func getImage(path string) string {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func printDecodeImage(image []int, width int, height int)  {
	for i:=0; i< height; i++ {
		for j:=0; j<width; j++ {
			var index = i*width + j
			if image[index] == ColorWhite {
				fmt.Printf("■")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	var image = getImage("Day 8/data.in")
	var width = 25
	var height = 6

	var fingerprint = getFingerPrint(image, width, height)
	fmt.Printf("Fingerprint: %d\n", fingerprint)

	var decodedImage = decodeImage(image, width, height)
	printDecodeImage(decodedImage, width, height)
}
