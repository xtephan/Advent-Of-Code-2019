package main

import (
	"../modules/opcodeProgram"
	"fmt"
)

type Point struct {
	x int
	y int
}

func paintPanel(panel map[Point]int) map[Point]int {

	var orientation int = 0
	var currentPosition Point = Point{
		x: 0,
		y: 0,
	}

	var program = opcodeProgram.New("Day 11/data.in")
	program.Execute()

	for !program.Halted {
		program.SendInput( panel[currentPosition] )

		panel[currentPosition] = program.Output[ len(program.Output) - 2 ]

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

	return panel
}

func drawPanel(panel map[Point]int)  {
	fmt.Printf("Your panel here, Panel Count: %d \n", len(panel))

	var start = Point{
		x: 9999,
		y: 9999,
	}

	var end = Point{
		x: -9999,
		y: -9999,
	}

	for point,_ := range panel {
		if point.x < start.x {
			start.x = point.x
		}
		if point.y < start.y {
			start.y = point.y
		}

		if point.x > end.x {
			end.x = point.x
		}
		if point.y > end.y {
			end.y = point.y
		}
	}

	for i := start.y; i<=end.y; i++ {
		for j := start.x; j<=end.x; j++ {
			var thisPoint = Point{
				x: j,
				y: i,
			}
			if panel[thisPoint] == 1 {
				fmt.Print("■")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}


}

func main() {

	//var panel map[Point]int = make(map[Point]int)
	//panel = paintPanel(panel)
	//fmt.Printf("Panel Count: %d", len(panel))

	var panel map[Point]int = make(map[Point]int)
	panel[Point{x: 0, y: 0}] = 1
	panel = paintPanel(panel)
	drawPanel(panel)

}
