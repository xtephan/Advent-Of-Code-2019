package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const OpCodePositionIncrement int = 4
const OpCodeAdd int = 1
const OpCodeMultiply int = 2
const OpCodeHalt int = 99

func getProgramOpcodes() []int {
	var opcodes []int

	data, err := ioutil.ReadFile("Day 2/data.in")
	if err != nil {
		fmt.Println("File reading error", err)
		return opcodes
	}

	for _, thisCode := range strings.Split(string(data), ",") {
		thisParsedCode, _ := strconv.Atoi(thisCode)
		opcodes = append(opcodes, thisParsedCode)
	}

	return opcodes
}

func executeOpcodes(_opcodes []int) []int {

	var opcodes = append(_opcodes[:0:0], _opcodes...)
	var currentPosition = 0

	for opcodes[currentPosition] != OpCodeHalt {

		var lho = opcodes[currentPosition + 1]
		var rho = opcodes[currentPosition + 2]
		var target = opcodes[currentPosition + 3]

		switch opcodes[currentPosition] {
		case OpCodeAdd:
			opcodes[target] = opcodes[lho] + opcodes[rho]
		case OpCodeMultiply:
			opcodes[target] = opcodes[lho] * opcodes[rho]
		}

		currentPosition += OpCodePositionIncrement
	}

	return opcodes
}

func findMagicSolution(opcodes []int) int {

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			opcodes[1] = noun
			opcodes[2] = verb

			var opcodes2 = executeOpcodes(opcodes)

			if opcodes2[0] == 19690720 {
				return 100 * noun + verb;
			}
		}
	}

	return -1
}

func main() {

	var opcodes = getProgramOpcodes()
	var magicNumber = findMagicSolution(opcodes)

	if magicNumber > 0 {
		fmt.Printf("Found solution: %d", magicNumber)
	} else {
		fmt.Printf("Nothing found :(")
	}
}