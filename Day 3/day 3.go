package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y  int
}

const DirectionUp string = "U"
const DirectionDown string = "D"
const DirectionLeft string = "L"
const DirectionRight string = "R"

func getWirePath(directions []string) []point {
	var path []point

	var currentX = 0
	var currentY = 0

	for _, thisDirection := range directions {
		var runes = []rune(thisDirection)
		var wireDirection = string(runes[0:1])
		var wireDirectionLength, _ =  strconv.Atoi(string(runes[1:]))
		for i := 0; i < wireDirectionLength; i++ {
			switch wireDirection {
			case DirectionDown:
				currentY--
			case DirectionUp:
				currentY++
			case DirectionLeft:
				currentX--
			case DirectionRight:
				currentX++
			}
			path = append(path, point{x: currentX, y: currentY})
		}
	}

	return path
}

func getWirePaths() [][]point {
	var result [][]point

	file, _ := os.Open("Day 3/data.in")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, getWirePath(strings.Split(scanner.Text(),",")))
	}

	return result
}

func isSamePoint(a point, b point) bool  {
	return a.x == b.x && a.y == b.y
}

func getManhattanDistance(a point, b point) int  {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func findClosestIntersection(wire1 []point, wire2 []point) int  {

	var minDistance = 999999 // hardcoded stupid big int
	var originPoint = point{x: 0, y:0}

	for i := 0; i < len(wire1); i++ {
		for j := 0; j < len(wire2); j++ {
			if isSamePoint(wire1[i], wire2[j]) {
				var distance = getManhattanDistance(wire1[i], originPoint)
				if minDistance > distance {
					minDistance = distance
				}
			}
		}
	}

	return minDistance
}


func findClosestIntersection2(wire1 []point, wire2 []point) int  {

	var minDistance = 999999 // hardcoded stupid big int

	for i := 0; i < len(wire1); i++ {
		for j := 0; j < len(wire2); j++ {
			if isSamePoint(wire1[i], wire2[j]) {
				var distance = i + j + 2 // 2 origin points
				if minDistance > distance {
					minDistance = distance
				}
			}
		}
	}

	return minDistance
}

func main() {

	var wires = getWirePaths()

	var closestManhattan = findClosestIntersection(wires[0], wires[1])
	var closestSteps = findClosestIntersection2(wires[0], wires[1])

	fmt.Printf("Closest Manhattan Distance: %d\n", closestManhattan)
	fmt.Printf("Closest Distance By Steps: %d", closestSteps)
}