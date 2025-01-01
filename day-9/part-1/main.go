package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	fileSystem := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for index, char := range line {
			charString := string(char)
			number, err := strconv.Atoi(charString)
			if err != nil {
				log.Fatalf("Error converting char to int: %s", err)
			}

			isFile := index%2 == 0
			for range number {
				var block string
				if isFile {
					fileIndex := index / 2
					block = strconv.Itoa(fileIndex)
				} else {
					block = "."
				}
				fileSystem = append(fileSystem, block)
			}
		}
	}

	leftIndex := 0
	rightIndex := len(fileSystem) - 1
	for {
		if leftIndex >= rightIndex {
			break
		}

		if fileSystem[leftIndex] == "." && fileSystem[rightIndex] != "." {
			fileSystem[leftIndex] = fileSystem[rightIndex]
			fileSystem[rightIndex] = "."
			leftIndex++
			rightIndex--
		} else {
			if fileSystem[leftIndex] != "." {
				leftIndex++
			}
			if fileSystem[rightIndex] == "." {
				rightIndex--
			}
		}
	}

	sum := 0

	for index, char := range fileSystem {
		if char == "." {
			break
		}

		number, err := strconv.Atoi(char)
		if err != nil {
			log.Fatalf("Error converting char to int: %s", err)
		}

		sum += (number * index)
	}

	fmt.Println("Resulting filesystem checksum:", sum)
}
