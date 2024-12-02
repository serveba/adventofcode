package aoc2024

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// day 1
// https://adventofcode.com/2024/day/2
// Red-Nosed Reports
func processReports(fileContent string) int {
	safeReports := 0
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		levels := strings.Split(line, " ")
		safety := processReport(levels)
		safeReports += safety
		// fmt.Printf("Levels: %v, Safety: %d\n", levels, safety)
	}
	return safeReports
}

func isDiffCorrect(diff int) bool {
	return diff >= 1 && diff <= 3
}

func processReport(levels []string) int {
	const safeLevel = 1
	const unsafeLevel = 0

	if len(levels) <= 2 {
		return safeLevel
	}

	previousLevel, _ := strconv.Atoi(levels[0])
	currentLevel, _ := strconv.Atoi(levels[1])

	diff := int(math.Abs(float64(currentLevel - previousLevel)))
	if !isDiffCorrect(diff) {
		return unsafeLevel
	}

	isIncreasing := false
	if previousLevel < currentLevel {
		isIncreasing = true
	}

	previousLevel = currentLevel

	for i := 2; i < len(levels); i++ {
		currentLevel, _ = strconv.Atoi(levels[i])
		diff = currentLevel - previousLevel
		if isIncreasing && diff < 0 || !isIncreasing && diff > 0 {
			return unsafeLevel
		}
		unsignedDiff := int(math.Abs(float64(diff)))
		if !isDiffCorrect(unsignedDiff) {
			return unsafeLevel
		}
		previousLevel = currentLevel
	}

	return safeLevel
}

func Day2() {
	byteArray, err := os.ReadFile("aoc2024/day2.input.txt")
	if err != nil {
		log.Fatal(err)
	}

	safeReports := processReports(string(byteArray))
	fmt.Printf("Part 1: %d\n", safeReports)
}
