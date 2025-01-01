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

	emptyStartIndex := 0
	emptySize := 0
	fileBlockStartIndex := len(fileSystem) - 1
	fileBlockSize := 0
	currentFileBlockId := ""
	for {
		if fileBlockStartIndex <= emptyStartIndex {
			if emptyStartIndex > 0 && fileBlockStartIndex > 0 {
				emptyStartIndex = 0
				fileBlockStartIndex--
			} else {
				break
			}
		}

		// Find the file block
		if fileSystem[fileBlockStartIndex] == currentFileBlockId {
			// Duplicate case where file block is the same as the previous one. Nothing needs to be done.
		} else if fileSystem[fileBlockStartIndex] != "." {
			currentFileBlockId = fileSystem[fileBlockStartIndex]
			fileBlockSize = 1
			for i := fileBlockStartIndex - 1; i > 0 && fileSystem[i] == currentFileBlockId; i-- {
				fileBlockSize++
				fileBlockStartIndex--
			}
		} else {
			for fileBlockStartIndex > 0 && fileSystem[fileBlockStartIndex] == "." {
				fileBlockStartIndex--
			}
			continue
		}

		// Find the empty block
		if fileSystem[emptyStartIndex] == "." {
			emptySize = 1
			for i := emptyStartIndex + 1; i < len(fileSystem) && fileSystem[i] == "."; i++ {
				emptySize++
			}
		} else {
			for emptyStartIndex < len(fileSystem) && fileSystem[emptyStartIndex] != "." {
				emptyStartIndex++
			}
			continue
		}

		// Put file in empty block if possible
		if emptySize >= fileBlockSize {
			for i := 0; i < fileBlockSize; i++ {
				fileSystem[emptyStartIndex+i] = fileSystem[fileBlockStartIndex+i]
				fileSystem[fileBlockStartIndex+i] = "."
			}
			emptyStartIndex = 0
			emptySize = 0
			fileBlockStartIndex--
			fileBlockSize = 0
		} else {
			emptyStartIndex += emptySize
			emptySize = 0
		}
	}

	sum := 0

	for index, char := range fileSystem {
		if char == "." {
			continue
		}

		number, err := strconv.Atoi(char)
		if err != nil {
			log.Fatalf("Error converting char to int: %s", err)
		}

		sum += (number * index)
	}

	fmt.Println("Resulting filesystem checksum:", sum)
}
