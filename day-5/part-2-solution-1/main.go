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

			wasInitiallyCorrect := true
			for {
				incorrectPosition := getIncorrectPosition(parsedUpdateNumbers, rules)
				if incorrectPosition == -1 {
					break
				}
				wasInitiallyCorrect = false
				parsedUpdateNumbers = fixIncorrectPosition(parsedUpdateNumbers, rules, incorrectPosition)
			}

			if !wasInitiallyCorrect {
				arrayLength := len(updateNumbers)
				sum += parsedUpdateNumbers[(arrayLength-1)/2]
			}
		}
	}

	fmt.Println("Sum of middle page numbers of newly-correctly-ordered updates:", sum)
}

func fixIncorrectPosition(updateNumbers []int, rules map[int]Rule, incorrectPosition int) []int {
	correctedUpdateNumbers := make([]int, len(updateNumbers))
	incorrectNumber := updateNumbers[incorrectPosition]
	newPosition := -1

	rule := rules[incorrectNumber]

	for i := incorrectPosition; i > -1; i-- {
		currentNumberToCheck := updateNumbers[i]
		_, currentNumberRuleExists := rules[currentNumberToCheck]

		if !slices.Contains(rule.earlierNumbers, currentNumberToCheck) &&
			(!currentNumberRuleExists || slices.Contains(rule.laterNumbers, currentNumberToCheck)) {
			newPosition = i
			break
		}
	}

	if newPosition == -1 {
		log.Fatalf("Could not find a correct position for number %d in %v. Earlier rules: %v. Later rules: %v.",
			incorrectNumber, updateNumbers, rule.earlierNumbers, rule.laterNumbers)
	}

	for i := 0; i < len(updateNumbers); i++ {
		if i == newPosition {
			correctedUpdateNumbers[i] = incorrectNumber
			continue
		}

		if i < newPosition {
			correctedUpdateNumbers[i] = updateNumbers[i]
			continue
		}

		if i <= incorrectPosition {
			correctedUpdateNumbers[i] = updateNumbers[i-1]
			continue
		}

		correctedUpdateNumbers[i] = updateNumbers[i]
	}

	return correctedUpdateNumbers
}

func getIncorrectPosition(updateNumbers []int, rules map[int]Rule) int {
	updateNumbersLength := len(updateNumbers)
	for i := 0; i < updateNumbersLength; i++ {
		updateNumber := updateNumbers[i]
		rule, exists := rules[updateNumber]
		if !exists {
			continue
		}

		for j := i + 1; j < updateNumbersLength; j++ {
			if slices.Contains(rule.earlierNumbers, updateNumbers[j]) {
				return j
			}
		}
	}

	return -1
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
