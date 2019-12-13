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
const OpAdjustRelativeBase int = 9
const OpCodeHalt int = 99

const ParameterModePosition int = 0
const ParameterModeImmediate int = 1
const ParameterModeRelative int = 2
const ParameterModeRelativeInput int = -2

type OpcodeProgram struct {
	opcodes map[int]int
	breakOnInput bool
	input int
	pointer int
	Output []int
	Halted bool
	relativeBase int
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
	OpAdjustRelativeBase: 2,
}

func New(filepath string) OpcodeProgram {
	op := OpcodeProgram{
		opcodes: make(map[int]int),
		breakOnInput: true,
		input: 0,
		Output: []int {},
		pointer: 0,
		Halted: false,
		relativeBase: 0,
	}
	op.readOpCodes(filepath)
	return op
}

func (op *OpcodeProgram) SetMemory(pointer int, value int) {
	op.opcodes[pointer] = value
}

func (op *OpcodeProgram) SendInput(input int) {
	// have the input
	op.breakOnInput = false
	op.input = input

	// continue the execution
	op.Execute()
}

func (op *OpcodeProgram) readInput() int {
	op.breakOnInput = true
	return op.input
}

func (op OpcodeProgram) GetLastOutput() int {
	return op.Output[len(op.Output) - 1]
}

func (op *OpcodeProgram) ClearOutput()  {
	op.Output = []int {}
}

func (op OpcodeProgram) DumpOutput() {
	fmt.Println("Output:")
	for _,v:= range op.Output {
		fmt.Printf("%d ", v)
	}
	fmt.Println("\n-----")
}


func (op *OpcodeProgram) readOpCodes(filepath string) {

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	for index, thisCode := range strings.Split(string(data), ",") {
		thisParsedCode, _ := strconv.Atoi(thisCode)
		op.opcodes[index] = thisParsedCode
	}
}

func(op OpcodeProgram) getParameterByMode(offset int, mode int) int {
	var parameterPointer = op.pointer + offset

	switch mode {
	case ParameterModePosition:
		return op.opcodes[op.opcodes[parameterPointer]]
	case ParameterModeImmediate:
		return op.opcodes[parameterPointer]
	case ParameterModeRelative:
		return op.opcodes[op.relativeBase + op.opcodes[parameterPointer]]
	case ParameterModeRelativeInput:
		return op.relativeBase + op.opcodes[parameterPointer]
	default:
		fmt.Printf("Unknown parameter mode %d\n", mode)
		return -1
	}
}

func(op OpcodeProgram) getParameter(offset int) int {
	var mode = op.opcodes[op.pointer] / int(math.Pow10(offset + 1)) % 10
	return op.getParameterByMode(offset, mode)
}

func(op OpcodeProgram) getWriteParameter(offset int) int {
	var mode = op.opcodes[op.pointer] / int(math.Pow10(offset + 1)) % 10
	if mode == ParameterModeRelative {
		return op.getParameterByMode(offset, ParameterModeRelativeInput)
	} else {
		return op.getParameterByMode(offset, ParameterModeImmediate)
	}
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
			var p3 = op.getWriteParameter(3)
			op.opcodes[p3] = p1 + p2
		case OpCodeMultiply:
			var p3 = op.getWriteParameter(3)
			op.opcodes[p3] = p1 * p2
		case OpCodeInput:
			var target = op.getWriteParameter(1)
			if op.breakOnInput {
				break OpExecution
			}
			ip := op.readInput()
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
			var p3 = op.getWriteParameter(3)
			var result = 0
			if p1 < p2 {
				result = 1
			}
			op.opcodes[p3] = result
		case OpEquals:
			var p3 = op.getWriteParameter(3)
			var result = 0
			if p1 == p2 {
				result = 1
			}
			op.opcodes[p3] = result
		case OpAdjustRelativeBase:
			op.relativeBase += p1
		case OpCodeHalt:
			op.Halted = true
			break OpExecution
		default:
			fmt.Printf("Unknown opcode: %d\n", opCode)
		}

		op.pointer += pointerIncrements[opCode]
	}
}
