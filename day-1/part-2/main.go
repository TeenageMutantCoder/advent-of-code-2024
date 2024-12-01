package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getIntsFromLine(line string) ([]int, error) {
	splitLine := strings.Split(line, "   ")
	if len(splitLine) != 2 {
		return nil, fmt.Errorf("invalid input: %s", line)
	}

	string1 := splitLine[0]
	num1, err := strconv.Atoi(string1)

	if err != nil {
		return nil, fmt.Errorf("failed to parse the following input as an int: %s", string1)
	}

	string2 := splitLine[1]
	num2, err := strconv.Atoi(string2)

	if err != nil {
		return nil, fmt.Errorf("failed to parse the following input as an int: %s", string2)
	}

	return []int{num1, num2}, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	list1 := make([]int, 0)
	list2FrequencyMap := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		nums, err := getIntsFromLine(line)

		if err != nil {
			log.Fatalf("Error parsing line: %s", err)
		}

		list1 = append(list1, nums[0])
		list2FrequencyMap[nums[1]] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	sum := 0
	for i := 0; i < len(list1); i++ {
		num := list1[i]
		sum += (num * list2FrequencyMap[num])
	}

	fmt.Println("The total similarity score of the two lists is: ", strconv.Itoa(sum))
}