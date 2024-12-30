package main

import (
	"bufio"
	"fmt"
	"log"
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

				antinode1Position := [2]int{y1 - yDiff, x1 - xDiff}
				antinode2Position := [2]int{y2 + yDiff, x2 + xDiff}

				if antinode1Position[0] >= 0 && antinode1Position[0] < mapHeight && antinode1Position[1] >= 0 && antinode1Position[1] < mapWidth {
					antinodePositions[antinode1Position] = true
				}
				if antinode2Position[0] >= 0 && antinode2Position[0] < mapHeight && antinode2Position[1] >= 0 && antinode2Position[1] < mapWidth {
					antinodePositions[antinode2Position] = true
				}
			}
		}
	}

	fmt.Println("Number of unique antinodes:", len(antinodePositions))
}
