package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	antennaPositionMap := make(map[string][][2]int, 0)
	mapHeight, mapWidth := 0, 0

	scanner := bufio.NewScanner(file)
	i := -1
	for scanner.Scan() {
		line := scanner.Text()
		i++
		mapHeight = i + 1
		for j, char := range line {
			mapWidth = j + 1
			if char != '.' {
				antennaPositionMap[string(char)] = append(antennaPositionMap[string(char)], [2]int{i, j})
			}
		}
	}

	antinodePositions := make(map[[2]int]bool, 0)

	for _, antennaPositions := range antennaPositionMap {
		for _, antennaPosition1 := range antennaPositions {
			for _, antennaPosition2 := range antennaPositions {
				if antennaPosition1 == antennaPosition2 {
					continue
				}

				y1, x1 := antennaPosition1[0], antennaPosition1[1]
				y2, x2 := antennaPosition2[0], antennaPosition2[1]

				yDiff := y2 - y1
				xDiff := x2 - x1
				slope := float64(yDiff) / float64(xDiff)
				for x := 0; x < mapWidth; x++ {
					y := slope*float64(x-x1) + float64(y1) // Point-slope form
					if y == math.Floor(y) && y >= 0 && int(y) < mapHeight {
						antinodePositions[[2]int{int(y), x}] = true
					}
				}
			}
		}
	}

	fmt.Println("Number of unique antinodes:", len(antinodePositions))
}
