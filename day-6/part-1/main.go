package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	fmt.Println("Number of distinct positions that the guard will visit:", getNumberOfDistinctPositionsInPath(puzzleMap, guardPosition))
}

func getNumberOfDistinctPositionsInPath(puzzleMap [][]rune, guardPosition [2]int) int {
	visitedPositions := make(map[[2]int]bool)
	visitedPositions[guardPosition] = true

	currentDirection := DIRECTION_UP

	for {
		positionX := guardPosition[0]
		positionY := guardPosition[1]
		if positionX < 0 || positionX >= len(puzzleMap[0]) || positionY < 0 || positionY >= len(puzzleMap) {
			break
		}

		visitedPositions[guardPosition] = true

		nextX := positionX + currentDirection[0]
		nextY := positionY + currentDirection[1]
		nextPositionIsOutOfBounds := nextX < 0 || nextX >= len(puzzleMap[0]) || nextY < 0 || nextY >= len(puzzleMap)
		if !nextPositionIsOutOfBounds && puzzleMap[nextY][nextX] == '#' {
			currentDirection = getNextDirection(currentDirection)
			continue
		}

		guardPosition = [2]int{nextX, nextY}
	}

	return len(visitedPositions)
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
