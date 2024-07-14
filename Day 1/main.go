package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalSum := 0

	digitMap := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	expression := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|\d`)
	reverseExpression := regexp.MustCompile(`eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d`)

	fmt.Println("Enter lines (end input with an empty line):")

	for {
		scanner.Scan()
		line := scanner.Text()

		if line == "" {
			break
		}

		matches := expression.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}

		firstMatch := matches[0]
		firstDigit := convertToDigit(firstMatch, digitMap)

		reverseLine := reverseString(line)
		reverseMatches := reverseExpression.FindAllString(reverseLine, -1)
		if len(reverseMatches) == 0 {
			continue
		}

		lastMatch := reverseMatches[0]
		lastDigit := convertToDigit(reverseString(lastMatch), digitMap)

		calibrationValue := firstDigit*10 + lastDigit
		totalSum += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read standard input:", err)
	}

	fmt.Printf("Total Sum: %d", totalSum)
}

func convertToDigit(s string, digitMap map[string]int) int {
	if val, ok := digitMap[s]; ok {
		return val
	}
	return int(s[0] - '0')
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
