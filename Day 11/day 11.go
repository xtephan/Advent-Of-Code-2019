package main

import (
	"../modules/opcodeProgram"
	"fmt"
)

type Point struct {
	x int
	y int
}

func getPanelCount() int {

	var panelColor map[Point]int = make(map[Point]int)

	var orientation int = 0
	var currentPosition Point = Point{
		x: 0,
		y: 0,
	}

	var program = opcodeProgram.New("Day 11/data.in")
	program.Execute()

	for !program.Halted {
		program.SendInput( panelColor[currentPosition] )

		panelColor[currentPosition] = program.Output[ len(program.Output) - 2 ]

		// turn right or left
		if program.Output[ len(program.Output) - 1 ] == 1 {
			orientation = (orientation + 90) % 360
		} else {
			orientation = (orientation + 270) % 360
		}

		switch orientation {
		case 0:
			currentPosition.y = currentPosition.y - 1
		case 90:
			currentPosition.x = currentPosition.x + 1
		case 180:
			currentPosition.y = currentPosition.y + 1
		case 270:
			currentPosition.x = currentPosition.x - 1
		}
	}

	return len(panelColor)
}

func main() {
	var panelCount = getPanelCount()
	fmt.Printf("Panel Count: %d", panelCount)
}
