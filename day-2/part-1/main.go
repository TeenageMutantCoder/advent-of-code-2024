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

const (
	MIN_DIFFERENCE = 1
	MAX_DIFFERENCE = 3
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	reports := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		report := getSliceFromLine(line)

		if checkIsValidAscending(report) || checkIsValidDescending(report) {
			reports = append(reports, report)
		}
	}

	fmt.Println("Number of valid reports:", len(reports))
}

func getSliceFromLine(line string) []int {
	splitLine := strings.Split(line, " ")
	parsedSplitLine := make([]int, 0)
	for _, num := range splitLine {
		parsedNum, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("Error parsing number: %s", err)
		}

		parsedSplitLine = append(parsedSplitLine, parsedNum)
	}

	return parsedSplitLine
}

func checkValueWithinRange(value int) bool {
	return value >= MIN_DIFFERENCE && value <= MAX_DIFFERENCE
}

func checkIsValidAscending(report []int) bool {
	prevNum := 0
	isValid := true
	for _, num := range report {
		if prevNum == 0 || (num > prevNum && checkValueWithinRange(int(math.Abs(float64(num-prevNum))))) {
			prevNum = num
			continue
		}

		isValid = false
		break
	}

	return isValid
}

func checkIsValidDescending(report []int) bool {
	prevNum := 0
	isValid := true
	for _, num := range report {
		if prevNum == 0 || (num < prevNum && checkValueWithinRange(int(math.Abs(float64(num-prevNum))))) {
			prevNum = num
			continue
		}

		isValid = false
		break
	}

	return isValid
}