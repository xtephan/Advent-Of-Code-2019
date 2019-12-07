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
		program.Execute()
		program.SendInput(v)
		program.SendInput(lastOutput)
		lastOutput = program.Output[0]
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

func getMaxOutput(amplifiers []int) int {
	return permuteAmplifiers(amplifiers, 0)
}

func main() {

	var maxOutput = getMaxOutput([]int{0,1,2,3,4})
	//var maxOutput = getMaxOutput([]int{5,6,7,8,9})

	fmt.Printf("Max Output %d\n", maxOutput)
	fmt.Printf("Done diddly done!")
}