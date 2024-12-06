package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	VALID_WORD           = "XMAS"
	DIRECTION_UP         = [2]int{0, -1}
	DIRECTION_RIGHT      = [2]int{1, 0}
	DIRECTION_DOWN       = [2]int{0, 1}
	DIRECTION_LEFT       = [2]int{-1, 0}
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

			if char == 'X' {
				startingPositionsToCheck = append(startingPositionsToCheck, [2]int{len(crossword), i})
			}
		}

		crossword = append(crossword, crosswordLine)
	}

	fmt.Println("Number of times that XMAS appears:", getNumberOfCrosswordSolutions(crossword, startingPositionsToCheck))
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

	count += checkDirectionIsSolution(DIRECTION_UP, VALID_WORD, crossword, startingPosition)
	count += checkDirectionIsSolution(DIRECTION_DOWN, VALID_WORD, crossword, startingPosition)
	count += checkDirectionIsSolution(DIRECTION_RIGHT, VALID_WORD, crossword, startingPosition)
	count += checkDirectionIsSolution(DIRECTION_LEFT, VALID_WORD, crossword, startingPosition)
	count += checkDirectionIsSolution(DIRECTION_UP_RIGHT, VALID_WORD, crossword, startingPosition)
	count += checkDirectionIsSolution(DIRECTION_UP_LEFT, VALID_WORD, crossword, startingPosition)
	count += checkDirectionIsSolution(DIRECTION_DOWN_RIGHT, VALID_WORD, crossword, startingPosition)
	count += checkDirectionIsSolution(DIRECTION_DOWN_LEFT, VALID_WORD, crossword, startingPosition)
	return count
}

func checkDirectionIsSolution(direction [2]int, expectedString string, crossword [][]rune, startingPosition [2]int) int {
	expectedStringLength := len(expectedString)
	currentPosition := startingPosition
	for offset := range expectedStringLength {
		if currentPosition[0] < 0 || currentPosition[0] >= len(crossword[0]) || currentPosition[1] < 0 || currentPosition[1] >= len(crossword) {
			return 0
		}

		if crossword[currentPosition[0]][currentPosition[1]] != rune(expectedString[offset]) {
			return 0
		}

		currentPosition[0] += direction[0]
		currentPosition[1] += direction[1]
	}

	return 1
}
