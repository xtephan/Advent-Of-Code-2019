package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

type Slope struct {
	dx       int
	dy       int
}

func getAsteroidPositions(path string) []Point {

	var result []Point

	file, _ := os.Open(path)
	defer file.Close()

	var y = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")
		for index,value := range data {
			if value == "#" {
				result = append(result, Point{x: index, y: y})
			}
		}
		y += 1
	}

	return result
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func getAbsolute(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func getGreatestCommonDenominator(x, y int) int {
	for y != 0 {

		x, y = y, x%y
	}
	return x
}

func getSlope(p1 Point, p2 Point) Slope {
	var dx = p2.x - p1.x
	var dy = p2.y - p1.y

	var gcd = getAbsolute(getGreatestCommonDenominator(dx, dy))

	return Slope{
		dx: dx / gcd,
		dy: dy / gcd,
	}
}

func getMostVisibleCount(positions []Point) int {
	var mostVisible = -1

	for _,thisStation := range positions {

		var existingSlopes = make(map[Slope]bool)

		for _,thisAsteroid := range positions {
			if thisStation.y != thisAsteroid.y || thisStation.x != thisAsteroid.x {
				var slope = getSlope(thisStation, thisAsteroid)
				existingSlopes[slope] = true
			}
		}


		if mostVisible <= len(existingSlopes) {
			mostVisible = len(existingSlopes)
		}

	}

	return mostVisible
}

func main() {

	//var asteroidPositions = getAsteroidPositions("Day 10/data-test.in")
	var asteroidPositions = getAsteroidPositions("Day 10/data.in")

	var mostVisible = getMostVisibleCount(asteroidPositions)

	fmt.Printf("Most visible asteroids: %d\n", mostVisible)
}
