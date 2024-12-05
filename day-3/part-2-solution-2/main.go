package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	fileContent := string(file)

	doCommandRegexPattern := `do\(\)`
	dontCommandRegexPattern := `don't\(\)`
	mulCommandRegexPattern := `mul\((\d{1,3}),(\d{1,3})\)`

	genericCommandRegexPattern := fmt.Sprintf("(%s)|(%s)|(%s)", doCommandRegexPattern, dontCommandRegexPattern, mulCommandRegexPattern)
	mulCommandRegex := regexp.MustCompile(mulCommandRegexPattern)
	genericCommandRegex := regexp.MustCompile(genericCommandRegexPattern)

	sum := 0
	shouldCalculate := true
	commands := genericCommandRegex.FindAllString(fileContent, -1)

	for _, command := range commands {
		if command == "do()" {
			shouldCalculate = true
		} else if command == "don't()" {
			shouldCalculate = false
		} else {
			if !shouldCalculate {
				continue
			}
			numbers := mulCommandRegex.FindStringSubmatch(command)
			num1, err := strconv.Atoi(numbers[1])
			if err != nil {
				log.Fatalf("Error converting string to int: %s", err)
			}
			num2, err := strconv.Atoi(numbers[2])
			if err != nil {
				log.Fatalf("Error converting string to int: %s", err)
			}
			sum += num1 * num2
		}
	}

	fmt.Println("Sum of multiplication results:", sum)
}
