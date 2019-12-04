package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getPasswordRange() (int, int) {
	data, _ := ioutil.ReadFile("Day 4/data.in")

	var raw = strings.Split(string(data), "-")

	var min, _ = strconv.Atoi(raw[0])
	var max, _ = strconv.Atoi(raw[1])

	return min, max
}

func isPossibleSolution(n int) bool {
	var digits []int

	for n > 0 {
		digits = append(digits, n%10)
		n = n/10
	}

	var isValid bool = false
	var identicalCnt = 0

	for i := 0; i< len(digits)-1; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
		if digits[i] == digits[i+1] {
			identicalCnt++
		} else {
			if identicalCnt == 1 {
				isValid = true
			}
			identicalCnt = 0
		}
	}

	return isValid || identicalCnt == 1
}

func getPossibleSolutions(min int, max int) int {

	var possibleSolutions int = 0

	for i := min; i <= max; i++ {
		if isPossibleSolution(i) {
			possibleSolutions++
		}
	}

	return possibleSolutions
}

func main() {

	var min, max = getPasswordRange()
	var possibleSolutions = getPossibleSolutions(min, max)

	fmt.Printf("Possible solutions: %d\n", possibleSolutions)
}