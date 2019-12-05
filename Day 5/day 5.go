package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const OpCodeAdd int = 1
const OpCodeMultiply int = 2
const OpCodeInput int = 3
const OpCodeOutput int = 4
const OpCodeHalt int = 99

func getProgramOpcodes() []int {
	var opcodes []int

	//data, err := ioutil.ReadFile("Day 5/data-inout.in")
	//data, err := ioutil.ReadFile("Day 5/data-parammode.in")
	data, err := ioutil.ReadFile("Day 5/data.in")
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

func getUserInput() int {
	return 1
}

func getParameterByMode(opcodes []int, position int, offset int, mode int) int {
	var pointer int

	if mode == 1 {
		pointer = position + offset
	} else {
		pointer = opcodes[position + offset]
	}

	if pointer < len(opcodes) {
		return opcodes[pointer]
	} else {
		return -1
	}
}

func getParameter(opcodes []int, position int, offset int) int {
	var mode = opcodes[position] / int(math.Pow10(offset + 1)) % 10
	return getParameterByMode(opcodes, position, offset, mode)
}

func executeOpcodes(_opcodes []int) []int {

	var opcodes = append(_opcodes[:0:0], _opcodes...)
	var currentPosition = 0
	var opCode = 0

	var pointerIncrements = map[int]int{
		OpCodeAdd: 4,
		OpCodeMultiply: 4,
		OpCodeInput: 2,
		OpCodeOutput: 2,
	}

	for opCode != OpCodeHalt {

		// Split it
		opCode = opcodes[currentPosition] % 100;
		var p1 = getParameter(opcodes, currentPosition, 1)
		var p2 = getParameter(opcodes, currentPosition, 2)

		switch opCode {
		case OpCodeAdd:
			var p3 = getParameterByMode(opcodes, currentPosition, 3, 1)
			opcodes[p3] = p1 + p2
		case OpCodeMultiply:
			var p3 = getParameterByMode(opcodes, currentPosition, 3, 1)
			opcodes[p3] =	p1 * p2
		case OpCodeInput:
			var target = opcodes[currentPosition + 1]
			opcodes[target] = getUserInput()
		case OpCodeOutput:
			fmt.Printf("Output: %d\n", p1)
		}

		currentPosition += pointerIncrements[opCode]
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
	executeOpcodes(opcodes)

	fmt.Printf("Done diddly done!")
}