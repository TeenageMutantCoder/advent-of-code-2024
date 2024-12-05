package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type TokenType int

const MAX_NUMBER_SIZE = 3
const MAX_NUMBER_COUNT = 2

const (
	COMMAND TokenType = iota
	OPEN_PARENTHESIS
	NUMBER
	SEPARATOR
	CLOSING_PARENTHESIS
	INVALID
	NULL
)

type Token struct {
	tokenType TokenType
	value     string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	tokens := make([]Token, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens = slices.Concat(tokens, getTokensFromLine(line))
	}

	fmt.Println("Sum of multiplication results:", getSumFromTokens(tokens))
}

func getSumFromTokens(tokens []Token) int {
	sum := 0
	expectedTokenType := COMMAND
	currentCommand := ""
	shouldCalculate := true
	numbers := make([]int, 0)

	for _, token := range tokens {
		if expectedTokenType == COMMAND && token.tokenType == COMMAND {
			currentCommand = token.value
			expectedTokenType = OPEN_PARENTHESIS
			continue
		}
		if expectedTokenType == OPEN_PARENTHESIS && token.tokenType == OPEN_PARENTHESIS {
			switch currentCommand {
			case "do", "don't":
				expectedTokenType = CLOSING_PARENTHESIS
			case "mul":
				expectedTokenType = NUMBER
			default:
				log.Fatalf("Invalid command: %s", currentCommand)
			}

			continue
		}
		if expectedTokenType == NUMBER && token.tokenType == NUMBER && len(numbers) < MAX_NUMBER_COUNT {
			number, err := strconv.Atoi(token.value)
			if err != nil {
				log.Fatalf("Error converting string to int: %s", err)
			}

			numbers = append(numbers, number)
			if len(numbers) == MAX_NUMBER_COUNT {
				expectedTokenType = CLOSING_PARENTHESIS
			} else {
				expectedTokenType = SEPARATOR
			}
			continue
		}
		if expectedTokenType == SEPARATOR && token.tokenType == SEPARATOR {
			expectedTokenType = NUMBER
			continue
		}
		if expectedTokenType == CLOSING_PARENTHESIS && token.tokenType == CLOSING_PARENTHESIS {
			switch currentCommand {
			case "mul":
				if shouldCalculate {
					product := 1
					for _, number := range numbers {
						product *= number
					}
					sum += product
				}
			case "do":
				shouldCalculate = true
			case "don't":
				shouldCalculate = false
			default:
				log.Fatalf("Invalid command: %s", currentCommand)
			}
		}

		expectedTokenType = COMMAND
		numbers = make([]int, 0)
	}

	return sum
}

func getTokensFromLine(line string) []Token {
	tokens := make([]Token, 0)
	currentToken := Token{tokenType: NULL, value: ""}

	for _, char := range line {
		tokenIsNull := currentToken.tokenType == NULL
		if unicode.IsDigit(char) {
			if currentToken.tokenType != NUMBER && !tokenIsNull {
				tokens = append(tokens, validateToken(currentToken))
				currentToken = Token{tokenType: NULL, value: ""}
			}
			currentToken.tokenType = NUMBER
			currentToken.value += string(char)
			continue
		}

		// START OF mul command
		if char == 'm' {
			if !tokenIsNull {
				tokens = append(tokens, validateToken(currentToken))
			}
			currentToken = Token{tokenType: COMMAND, value: string(char)}
			continue
		}

		if char == 'u' && currentToken.tokenType == COMMAND && currentToken.value == "m" {
			currentToken.value += string(char)
			continue
		}

		if char == 'l' && currentToken.tokenType == COMMAND && currentToken.value == "mu" {
			currentToken.value += string(char)
			tokens = append(tokens, validateToken(currentToken))
			currentToken = Token{tokenType: NULL, value: ""}
			continue
		}
		// END of mul command

		// START OF do and don't commands
		if char == 'd' {
			if !tokenIsNull {
				tokens = append(tokens, validateToken(currentToken))
			}
			currentToken = Token{tokenType: COMMAND, value: string(char)}
			continue
		}

		if char == 'o' && currentToken.tokenType == COMMAND && currentToken.value == "d" {
			currentToken.value += string(char)
			continue
		}

		if char == 'n' && currentToken.tokenType == COMMAND && currentToken.value == "do" {
			currentToken.value += string(char)
			continue
		}

		if char == '\'' && currentToken.tokenType == COMMAND && currentToken.value == "don" {
			currentToken.value += string(char)
			continue
		}

		if char == 't' && currentToken.tokenType == COMMAND && currentToken.value == "don'" {
			currentToken.value += string(char)
			tokens = append(tokens, validateToken(currentToken))
			currentToken = Token{tokenType: NULL, value: ""}
			continue
		}
		// END OF do and don't commands

		if char == '(' {
			if !tokenIsNull {
				tokens = append(tokens, validateToken(currentToken))
			}
			currentToken = Token{tokenType: OPEN_PARENTHESIS, value: string(char)}
			tokens = append(tokens, validateToken(currentToken))
			currentToken = Token{tokenType: NULL, value: ""}
			continue
		}
		if char == ')' {
			if !tokenIsNull {
				tokens = append(tokens, validateToken(currentToken))
			}
			currentToken = Token{tokenType: CLOSING_PARENTHESIS, value: string(char)}
			tokens = append(tokens, validateToken(currentToken))
			currentToken = Token{tokenType: NULL, value: ""}
			continue
		}
		if char == ',' {
			if !tokenIsNull {
				tokens = append(tokens, validateToken(currentToken))
			}
			currentToken = Token{tokenType: SEPARATOR, value: string(char)}
			tokens = append(tokens, validateToken(currentToken))
			currentToken = Token{tokenType: NULL, value: ""}
			continue
		}

		tokens = append(tokens, validateToken(currentToken))
		currentToken = Token{tokenType: INVALID, value: string(char)}
	}

	return tokens
}

func validateToken(token Token) Token {
	if !checkTokenIsValid(token) {
		token.tokenType = INVALID
	}

	return token
}

func checkTokenIsValid(token Token) bool {
	if token.tokenType == COMMAND {
		return token.value == "mul" || token.value == "do" || token.value == "don't"
	}
	if token.tokenType == OPEN_PARENTHESIS {
		return token.value == "("
	}
	if token.tokenType == CLOSING_PARENTHESIS {
		return token.value == ")"
	}
	if token.tokenType == SEPARATOR {
		return token.value == ","
	}
	if token.tokenType == NUMBER {
		_, err := strconv.Atoi(token.value)
		numberLength := len(token.value)
		return err == nil && numberLength <= MAX_NUMBER_SIZE && numberLength > 0
	}
	return false
}
