package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var (
	DIRECTION_UP    = [2]int{0, -1}
	DIRECTION_DOWN  = [2]int{0, 1}
	DIRECTION_LEFT  = [2]int{-1, 0}
	DIRECTION_RIGHT = [2]int{1, 0}
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	puzzleMap := make([][]rune, 0)
	guardPosition := [2]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mapLine := make([]rune, 0)

		for i, char := range line {
			mapLine = append(mapLine, char)

			if char == '^' {
				guardPosition = [2]int{i, len(puzzleMap)}
			}
		}

		puzzleMap = append(puzzleMap, mapLine)
	}

	possibleObstructionPositions := make([][2]int, 0)

	for yPosition, row := range puzzleMap {
		for xPosition, column := range row {
			if column == '#' || column == '^' {
				continue
			}

			if checkIsLooping(puzzleMap, guardPosition, [2]int{xPosition, yPosition}) {
				possibleObstructionPositions = append(possibleObstructionPositions, [2]int{xPosition, yPosition})
			}
		}
	}

	fmt.Println("Number of possible positions to choose for obstruction:", len(possibleObstructionPositions))
}

func checkIsLooping(puzzleMap [][]rune, guardPosition [2]int, obstaclePosition [2]int) bool {
	currentDirection := DIRECTION_UP
	visitedPositions := make(map[[2]int][][2]int)
	isLooping := false

	for {
		positionX := guardPosition[0]
		positionY := guardPosition[1]
		if positionX < 0 || positionX >= len(puzzleMap[0]) || positionY < 0 || positionY >= len(puzzleMap) {
			break
		}

		if visitedPositions[guardPosition] != nil && slices.Contains(visitedPositions[guardPosition], currentDirection) {
			isLooping = true
			break
		}

		visitedPositions[guardPosition] = getUpdatedVisitedPosition(visitedPositions, guardPosition, currentDirection)

		nextX := positionX + currentDirection[0]
		nextY := positionY + currentDirection[1]
		nextPositionIsOutOfBounds := nextX < 0 || nextX >= len(puzzleMap[0]) || nextY < 0 || nextY >= len(puzzleMap)
		if !nextPositionIsOutOfBounds && (puzzleMap[nextY][nextX] == '#' || (nextX == obstaclePosition[0] && nextY == obstaclePosition[1])) {
			currentDirection = getNextDirection(currentDirection)
			continue
		}

		guardPosition = [2]int{nextX, nextY}
	}

	return isLooping
}

func getNextDirection(currentDirection [2]int) [2]int {
	if currentDirection == DIRECTION_UP {
		return DIRECTION_RIGHT
	} else if currentDirection == DIRECTION_RIGHT {
		return DIRECTION_DOWN
	} else if currentDirection == DIRECTION_DOWN {
		return DIRECTION_LEFT
	} else {
		return DIRECTION_UP
	}
}

func getUpdatedVisitedPosition(visitedPositions map[[2]int][][2]int, guardPosition [2]int, currentDirection [2]int) [][2]int {
	visitedPositionDirections, exists := visitedPositions[guardPosition]
	if !exists {
		visitedPositionDirections = make([][2]int, 0)
	}

	if !slices.Contains(visitedPositionDirections, currentDirection) {
		visitedPositionDirections = append(visitedPositionDirections, currentDirection)
	}

	return visitedPositionDirections
}
