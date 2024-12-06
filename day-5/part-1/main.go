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

			if !isValidUpdateOrder(parsedUpdateNumbers, rules) {
				continue
			}

			arrayLength := len(updateNumbers)
			sum += parsedUpdateNumbers[(arrayLength-1)/2]
		}
	}

	fmt.Println("Sum of middle page numbers of correctly-ordered updates:", sum)
}

func isValidUpdateOrder(updateNumbers []int, rules map[int]Rule) bool {
	for i, updateNumber := range updateNumbers {
		rule, exists := rules[updateNumber]
		if !exists {
			continue
		}

		for j := i + 1; j < len(updateNumbers); j++ {
			if slices.Contains(rule.earlierNumbers, updateNumbers[j]) {
				return false
			}
		}
	}

	return true
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
