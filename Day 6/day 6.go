package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const You string = "YOU"
const Santa string = "SAN"

func getOrbitMap() map[string]string {
	var directOrbits map[string]string = make(map[string]string)

	file, _ := os.Open("Day 6/data.in")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ")")
		directOrbits[data[1]] = data[0]
	}

	return directOrbits
}

func getTotalOrbits(orbitMap map[string]string) int {
	var totalOrbits = 0

	for orbiter := range orbitMap {
		// orbiter orbits _
		var currentOrbiter = orbiter

		for orbitMap[currentOrbiter] != "" {
			totalOrbits++
			currentOrbiter = orbitMap[currentOrbiter]
		}

	}

	return totalOrbits
}

func getPathToCOM(orbitMap map[string]string, object string) map[string]int {

	var path map[string]int = make(map[string]int)
	var distance = 0

	var currentOrbiter = object

	for orbitMap[currentOrbiter] != "" {
		path[currentOrbiter] = distance
		currentOrbiter = orbitMap[currentOrbiter]
		distance++
	}

	return path
}

func getDistanceToSanta(orbitMap map[string]string) int {
	var minDistance = 9999999 // I still dont know how to maxInst

	var youPath = getPathToCOM(orbitMap, You)
	var santaPath = getPathToCOM(orbitMap, Santa)

	for object, distance := range youPath {
		// Object is at distance
		if santaPath[object] != 0{
			var thisDistance = santaPath[object] + distance
			if thisDistance < minDistance {
				minDistance = thisDistance
			}
		}
	}

	return minDistance - 2
}

func main() {
	var orbitMap = getOrbitMap()

	var totalOrbits = getTotalOrbits(orbitMap)
	fmt.Printf("Total orbits: %d\n", totalOrbits)

	var distanceToSanta = getDistanceToSanta(orbitMap)
	fmt.Printf("Distance to Santa: %d", distanceToSanta)
}