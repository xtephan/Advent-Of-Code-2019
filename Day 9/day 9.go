package main

import (
	"../modules/opcodeProgram"
)

func main() {


	// Part 1
	//var program = opcodeProgram.New("Day 9/data.in")
	//program.Execute()
	//program.SendInput(1)
	//program.DumpOutput()

	// Part 2
	var program2 = opcodeProgram.New("Day 9/data.in")
	program2.Execute()
	program2.SendInput(2)
	program2.DumpOutput()

	// Test
	//var programTest = opcodeProgram.New("Day 9/data-test.in")
	//programTest.Execute()
	//programTest.DumpOutput()
}