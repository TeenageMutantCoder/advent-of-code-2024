package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	earlierNumbers []int
	laterNumbers   []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	sum := 0

	hasParsedRules := false
	rules := make(map[int]Rule)
	updatesSorterFunc := func(a, b int) int {
		if _, exists := rules[a]; !exists {
			return 0
		}
		if slices.Contains(rules[a].laterNumbers, b) {
			return -1
		}
		if slices.Contains(rules[a].earlierNumbers, b) {
			return 1
		}
		return 0
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			hasParsedRules = true
			continue
		}

		if !hasParsedRules {
			ruleNumbers := strings.Split(line, "|")
			parsedRuleNumbers := getIntSliceFromStringSlice(ruleNumbers)
			earlierNumber := parsedRuleNumbers[0]
			laterNumber := parsedRuleNumbers[1]

			if _, exists := rules[earlierNumber]; !exists {
				rules[earlierNumber] = Rule{}
			}
			if _, exists := rules[laterNumber]; !exists {
				rules[laterNumber] = Rule{}
			}

			earlierNumberRule := rules[earlierNumber]
			earlierNumberRule.laterNumbers = append(earlierNumberRule.laterNumbers, laterNumber)
			rules[earlierNumber] = earlierNumberRule

			laterNumberRule := rules[laterNumber]
			laterNumberRule.earlierNumbers = append(laterNumberRule.earlierNumbers, earlierNumber)
			rules[laterNumber] = laterNumberRule
		} else {
			updateNumbers := strings.Split(line, ",")
			parsedUpdateNumbers := getIntSliceFromStringSlice(updateNumbers)

			if slices.IsSortedFunc(parsedUpdateNumbers, updatesSorterFunc) {
				continue
			}

			slices.SortFunc(parsedUpdateNumbers, updatesSorterFunc)
			arrayLength := len(updateNumbers)
			sum += parsedUpdateNumbers[(arrayLength-1)/2]
		}
	}

	fmt.Println("Sum of middle page numbers of newly-correctly-ordered updates:", sum)
}

func getIntSliceFromStringSlice(stringSlice []string) []int {
	intSlice := make([]int, 0)

	for _, stringNumber := range stringSlice {
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}
		intSlice = append(intSlice, number)
	}

	return intSlice
}
