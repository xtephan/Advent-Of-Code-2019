package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFuelRequirementForMass(mass int) int {
	return mass / 3 - 2
}

func getFuelRequirementForModuleWithMass(mass int) int {
	var total int = 0
	var remainingMass int = mass
	for remainingMass > 0 {
		var fuelNeeded = getFuelRequirementForMass(remainingMass)
		if fuelNeeded > 0 {
			total += fuelNeeded
		}
		remainingMass = fuelNeeded
	}
	return total
}

func getModulesMasses() []int {

	var masses []int

	// Open the file
	file, err := os.Open("data.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		masses = append(masses, mass)
	}

	return masses
}

func getTotalFuelRequirement(masses []int) int {
	var totalFuelRequirement int = 0
	for _, thisMass := range masses {
		totalFuelRequirement += getFuelRequirementForModuleWithMass(thisMass)
	}
	return totalFuelRequirement
}

func main() {

	var modulesMass = getModulesMasses()
	var totalFuelRequirement = getTotalFuelRequirement(modulesMass)

	fmt.Printf("Total fuel requrement: %d", totalFuelRequirement)
}