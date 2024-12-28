package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ": ")
		expectedResult, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatalf("Error converting string to integer: %s", err)
		}

		splitNumbers := strings.Split(splitLine[1], " ")
		splitNumbersLength := len(splitNumbers)

		for i := 0; float64(i) < math.Pow(3, float64(splitNumbersLength-1)); i++ {
			currentOperators := getOperatorSlice(i, splitNumbersLength-1)
			numbers := getIntSliceFromStringSlice(splitNumbers)
			if checkIfCalibrationIsCorrect(numbers, currentOperators, expectedResult) {
				sum += expectedResult
				break
			}
		}
	}

	fmt.Println("Total calibration result:", sum)
}

func getIntSliceFromStringSlice(stringSlice []string) []int {
	intSlice := make([]int, len(stringSlice))
	for i := range stringSlice {
		number, err := strconv.Atoi(stringSlice[i])
		if err != nil {
			log.Fatalf("Error converting string to integer: %s", err)
		}
		intSlice[i] = number
	}

	return intSlice
}

func checkIfCalibrationIsCorrect(numbers []int, operators []string, expectedResult int) bool {
	sum := numbers[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == "+" {
			sum += numbers[i+1]
		} else if operators[i] == "*" {
			sum *= numbers[i+1]
		} else {
			updatedSum, err := strconv.Atoi(strconv.Itoa(sum) + strconv.Itoa(numbers[i+1]))
			if err != nil {
				log.Fatalf("Error converting string to integer: %s", err)
			}
			sum = updatedSum
		}

		if sum > expectedResult {
			return false
		}
	}
	return sum == expectedResult
}

func getOperatorSlice(seed int, size int) []string {
	bitField := fmt.Sprintf("%0*v", size, strconv.FormatInt(int64(seed), 3))
	operatorSlice := make([]string, size)
	for i := range operatorSlice {
		if bitField[i] == '0' {
			operatorSlice[i] = "+"
		} else if bitField[i] == '1' {
			operatorSlice[i] = "*"
		} else {
			operatorSlice[i] = "||"
		}
	}

	return operatorSlice
}
