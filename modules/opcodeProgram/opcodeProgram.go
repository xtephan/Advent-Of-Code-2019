package opcodeProgram

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
const OpJumpIfTrue int = 5
const OpJumpIfFalse int = 6
const OpLessThan int = 7
const OpEquals int = 8
const OpCodeHalt int = 99

type OpcodeProgram struct {
	opcodes []int
	breakOnInput bool
	input int
	pointer int
	Output []int
}

var pointerIncrements = map[int]int{
	OpCodeAdd: 4,
	OpCodeMultiply: 4,
	OpCodeInput: 2,
	OpCodeOutput: 2,
	OpJumpIfTrue: 0,
	OpJumpIfFalse: 0,
	OpLessThan: 4,
	OpEquals: 4,
}

func New(filepath string) OpcodeProgram {
	op := OpcodeProgram{
		opcodes: getProgramOpcodes(filepath),
		breakOnInput: true,
		input: 0,
		Output: make([]int, 0),
		pointer: 0,
	}
	return op
}

func (op *OpcodeProgram) SendInput(input int) {
	// have the input
	op.breakOnInput = false
	op.input = input

	// continue the execution
	op.Execute()
}

func (op *OpcodeProgram) getInput() int {
	op.breakOnInput = true
	return op.input
}

func getProgramOpcodes(filepath string) []int {
	var opcodes []int

	data, err := ioutil.ReadFile(filepath)
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

func(op OpcodeProgram) getParameterByMode(offset int, mode int) int {
	var pointer int

	var currentPosition = op.pointer + offset

	if currentPosition >= len(op.opcodes) {
		return -1
	}

	if mode == 1 {
		pointer = currentPosition
	} else {
		pointer = op.opcodes[currentPosition]
	}

	if pointer < len(op.opcodes) {
		return op.opcodes[pointer]
	} else {
		return -1
	}
}

func(op OpcodeProgram) getParameter(offset int) int {
	var mode = op.opcodes[op.pointer] / int(math.Pow10(offset + 1)) % 10
	return op.getParameterByMode(offset, mode)
}

func(op *OpcodeProgram) Execute() {

	var opCode = 0

OpExecution:
	for {

		// Split it
		opCode = op.opcodes[op.pointer] % 100
		var p1 = op.getParameter(1)
		var p2 = op.getParameter(2)

		switch opCode {
		case OpCodeAdd:
			var p3 = op.getParameterByMode(3, 1)
			op.opcodes[p3] = p1 + p2
		case OpCodeMultiply:
			var p3 = op.getParameterByMode(3, 1)
			op.opcodes[p3] = p1 * p2
		case OpCodeInput:
			var target = op.getParameterByMode(1, 1)
			if op.breakOnInput {
				break OpExecution
			}
			ip := op.getInput()
			//fmt.Printf("Read input: %d\n", ip)
			op.opcodes[target] = ip
		case OpCodeOutput:
			op.Output = append(op.Output, p1)
			//fmt.Printf("Program Output: %d\n", p1)
		case OpJumpIfTrue:
			if p1 != 0 {
				op.pointer = p2
			} else {
				op.pointer += 3
			}
		case OpJumpIfFalse:
			if p1 == 0 {
				op.pointer = p2
			} else {
				op.pointer += 3
			}
		case OpLessThan:
			var p3 = op.getParameterByMode(3, 1)
			var result = 0
			if p1 < p2 {
				result = 1
			}
			op.opcodes[p3] = result
		case OpEquals:
			var p3 = op.getParameterByMode(3, 1)
			var result = 0
			if p1 == p2 {
				result = 1
			}
			op.opcodes[p3] = result
		case OpCodeHalt:
			break OpExecution
		default:
			fmt.Printf("Unknown opcode: %d\n", opCode)
		}

		op.pointer += pointerIncrements[opCode]
	}
}
