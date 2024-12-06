package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	VALID_WORD           = "MAS"
	DIRECTION_UP_RIGHT   = [2]int{1, -1}
	DIRECTION_UP_LEFT    = [2]int{-1, -1}
	DIRECTION_DOWN_RIGHT = [2]int{1, 1}
	DIRECTION_DOWN_LEFT  = [2]int{-1, 1}
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	crossword := make([][]rune, 0)
	startingPositionsToCheck := make([][2]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		crosswordLine := make([]rune, 0)

		for i, char := range line {
			crosswordLine = append(crosswordLine, char)

			if char == 'A' {
				startingPositionsToCheck = append(startingPositionsToCheck, [2]int{len(crossword), i})
			}
		}

		crossword = append(crossword, crosswordLine)
	}

	fmt.Println("Number of times that X-MAS appears:", getNumberOfCrosswordSolutions(crossword, startingPositionsToCheck))
}

func getNumberOfCrosswordSolutions(crossword [][]rune, startingPositions [][2]int) int {
	if len(startingPositions) == 0 {
		return 0
	}

	count := 0

	for _, startingPosition := range startingPositions {
		count += getNumberOfCrosswordSolutionsAtPoint(crossword, startingPosition)
	}

	return count
}

func getNumberOfCrosswordSolutionsAtPoint(crossword [][]rune, startingPosition [2]int) int {
	count := 0

	hasDiagonalSolution1 :=
		checkDirectionIsSolution(DIRECTION_UP_RIGHT, VALID_WORD, crossword, startingPosition) == 1 ||
			checkDirectionIsSolution(DIRECTION_DOWN_LEFT, VALID_WORD, crossword, startingPosition) == 1

	hasDiagonalSolution2 :=
		checkDirectionIsSolution(DIRECTION_UP_LEFT, VALID_WORD, crossword, startingPosition) == 1 ||
			checkDirectionIsSolution(DIRECTION_DOWN_RIGHT, VALID_WORD, crossword, startingPosition) == 1

	if hasDiagonalSolution1 && hasDiagonalSolution2 {
		count++
	}
	return count
}

func checkDirectionIsSolution(direction [2]int, expectedString string, crossword [][]rune, startingPosition [2]int) int {
	currentPosition := [2]int{}
	for multiplier := -1; multiplier < 2; multiplier += 2 {
		currentPosition[0] = startingPosition[0] + direction[0]*multiplier
		currentPosition[1] = startingPosition[1] + direction[1]*multiplier
		if currentPosition[0] < 0 || currentPosition[0] >= len(crossword[0]) || currentPosition[1] < 0 || currentPosition[1] >= len(crossword) {
			return 0
		}

		if crossword[currentPosition[0]][currentPosition[1]] != rune(expectedString[multiplier+1]) {
			return 0
		}
	}

	return 1
}
