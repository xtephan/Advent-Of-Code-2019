package main

import (
	"../modules/opcodeProgram"
	"fmt"
)

type Point struct {
	x int
	y int
}

var scorePoint = Point{
	x: -1,
	y: 0,
}

const MaxSize = 100

const TileEmpty = 0
const TileWall = 1
const TileBlock = 2
const TilePaddle = 3
const TileBall = 4

const JoystickNeutral = 0
const JoystickLeft = -1
const JoystickRight = 1

func getGameBoard() [MaxSize][MaxSize]int {
	gameBoard := [MaxSize][MaxSize]int{}

	var program = opcodeProgram.New("Day 13/data.in")
	program.Execute()

	for i:=0; i<len(program.Output); i+=3 {
		gameBoard[program.Output[i]][program.Output[i+1]] = program.Output[i+2]
	}

	return gameBoard
}

func countTiles(gameBoard [MaxSize][MaxSize]int) int  {
	tiles := 0

	for i:= range gameBoard {
		for j:= range gameBoard[i] {
			if gameBoard[i][j] == TileBlock {
				tiles++
			}
		}
	}

	return tiles
}

func countTiles2(blocks map[Point]bool) int  {
	tiles := 0

	for _,isBlock := range blocks {
		if isBlock {
			tiles++
		}
	}

	return tiles
}

func playGame() int {
	score := 0
	blocks := make(map[Point]bool)
	tiles := 100

	paddlePosition := Point{
		x: -1,
		y: -1,
	}
	ballPosition := Point{
		x: -1,
		y: -1,
	}

	var program = opcodeProgram.New("Day 13/data.in")
	program.SetMemory(0, 2)
	program.Execute()

	for tiles != 0 {

		// Update the game board
		for i := 0; i < len(program.Output); i += 3 {
			thisPoint := Point{
				x: program.Output[i],
				y: program.Output[i + 1],
			}
			payload := program.Output[i+2]

			if thisPoint == scorePoint {
				score = payload
			} else {
				switch payload {
				case TileBlock:
					blocks[thisPoint] = true
				case TileEmpty:
					if blocks[thisPoint] {
						blocks[thisPoint] = false
					}
				case TileBall:
					ballPosition = thisPoint
				case TilePaddle:
					paddlePosition = thisPoint
				}
			}
		}

		program.ClearOutput()
		tiles = countTiles2(blocks)

		// Move the paddle
		if paddlePosition.x == ballPosition.x {
			program.SendInput(JoystickNeutral)
		} else if paddlePosition.x < ballPosition.x {
			program.SendInput(JoystickRight)
		} else {
			program.SendInput(JoystickLeft)
		}
	}

	return score
}

func main() {

	//gameBoard := getGameBoard()
	//tiles := countTiles(gameBoard)
	//fmt.Printf("Tiles: %d\n", tiles)

	score := playGame()
	fmt.Printf("Score: %d\n", score)

}