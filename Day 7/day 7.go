package main

import "fmt"
import "../modules/opcodeProgram"

func printAmplifiers(amplifiers []int) {
	fmt.Printf("Amplifiers: ")
	for _,v := range amplifiers {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

func getAmplifiersOutput(amplifiers []int) int {
	var lastOutput = 0

	for _,v := range amplifiers {
		program := opcodeProgram.New("Day 7/data.in")
		var output = program.Execute( []int{v, lastOutput} )
		lastOutput = output[0]
	}

	return lastOutput
}

func permuteAmplifiers(amplifiers []int, index int) int {
	if index >= len(amplifiers) - 1 {
		printAmplifiers(amplifiers)
		var output = getAmplifiersOutput(amplifiers)
		fmt.Printf("Output: %d \n", output)
		return output
	}

	var maxOutput = 0

	for i:=index; i<len(amplifiers); i++ {
		var tmp int

		tmp = amplifiers[index]
		amplifiers[index] = amplifiers[i]
		amplifiers[i] = tmp

		var output = permuteAmplifiers(amplifiers, index+1)

		if output > maxOutput {
			maxOutput = output
		}

		tmp = amplifiers[index]
		amplifiers[index] = amplifiers[i]
		amplifiers[i] = tmp
	}

	return maxOutput
}

func getMaxOutput() int {
	var amplifiers = []int{5,6,7,8,9}
	var max = permuteAmplifiers(amplifiers, 0)
	return max
}

func main() {

	var maxOutput = getMaxOutput()

	fmt.Printf("Max Output %d\n", maxOutput)
	fmt.Printf("Done diddly done!")
}