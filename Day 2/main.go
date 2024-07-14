package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseGame(gameStr string) (int, []map[string]int) {
	gameIDPattern := regexp.MustCompile(`Game (\d+):`)
	gameIDMatch := gameIDPattern.FindStringSubmatch(gameStr)
	if len(gameIDMatch) == 0 {
		return 0, nil
	}
	gameID, _ := strconv.Atoi(gameIDMatch[1])

	subsets := strings.Split(gameStr[strings.Index(gameStr, ":")+1:], ";")
	var parsedSubsets []map[string]int
	for _, subset := range subsets {
		colorPattern := regexp.MustCompile(`(\d+) (\w+)`)
		colorMatches := colorPattern.FindAllStringSubmatch(subset, -1)
		parsedColors := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, match := range colorMatches {
			count, _ := strconv.Atoi(match[1])
			color := match[2]
			parsedColors[color] = count
		}
		parsedSubsets = append(parsedSubsets, parsedColors)
	}

	return gameID, parsedSubsets
}

func minCubesRequired(subsets []map[string]int) (int, int, int) {
	minRed, minGreen, minBlue := 0, 0, 0
	for _, subset := range subsets {
		if subset["red"] > minRed {
			minRed = subset["red"]
		}
		if subset["green"] > minGreen {
			minGreen = subset["green"]
		}
		if subset["blue"] > minBlue {
			minBlue = subset["blue"]
		}
	}
	return minRed, minGreen, minBlue
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the games data: ")

	var inputData []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		inputData = append(inputData, line)
	}

	totalPower := 0

	for _, gameStr := range inputData {
		gameID, subsets := parseGame(gameStr)
		if gameID != 0 {
			minRed, minGreen, minBlue := minCubesRequired(subsets)
			power := minRed * minGreen * minBlue
			totalPower += power
		}
	}

	fmt.Printf("Sum of the power of the minimum sets: %d", totalPower)
}
