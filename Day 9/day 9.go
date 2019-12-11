package main

import (
	"../modules/opcodeProgram"
	"fmt"
)

func main() {


	var program2 = opcodeProgram.New("Day 9/data.in")
	program2.Execute()
	program2.SendInput(1)
	program2.DumpOutput()

	//var program = opcodeProgram.New("Day 9/data-test.in")
	//program.Execute()
	//program.DumpOutput()

	fmt.Println("\n-----")
}