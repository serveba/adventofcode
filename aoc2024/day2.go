package aoc2024

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/2
// Red-Nosed Reports
func processReportsPart1(fileContent string) int {
	safeReports := 0
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		report := strings.Split(line, " ")
		safety := process(report)
		safeReports += safety
		// fmt.Printf("Levels: %v, Safety: %d\n", levels, safety)
	}
	return safeReports
}

func processReportsPart2(fileContent string) int {
	sReports := 0
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		report := strings.Split(line, " ")

		safety := process(report)

		if safety == 1 {
			sReports += 1
			continue
		}

		//fmt.Printf("LINE %d, REPORT: %v, Safety: %d\n", i+1, report, safety)
		for j := 0; j < len(report); j++ {
			newReport := make([]string, 0, len(report)-1)

			newReport = append(newReport, report[:j]...)
			newReport = append(newReport, report[j+1:]...)

			safety := process(newReport)
			//fmt.Printf("LINE %d, Level %d REMOVED: %v, Safety: %d\n", i+1, j, newReport, safety)
			if safety == 1 {
				sReports += safety
				break
			}
		}
		//fmt.Printf("LINE %d, Levels: %v, Safety: %d\n", i+1, report, safety)
	}
	return sReports
}

func isDiffCorrect(diff int) bool {
	return diff >= 1 && diff <= 3
}

func process(report []string) int {
	const sLevel = 1
	const uLevel = 0

	if len(report) <= 2 {
		return sLevel
	}

	prevLevel, _ := strconv.Atoi(report[0])
	curLevel, _ := strconv.Atoi(report[1])

	diff := int(math.Abs(float64(curLevel - prevLevel)))
	if !isDiffCorrect(diff) {
		return uLevel
	}

	isIncreasing := false
	if prevLevel < curLevel {
		isIncreasing = true
	}

	prevLevel = curLevel

	for i := 2; i < len(report); i++ {
		curLevel, _ = strconv.Atoi(report[i])
		diff = curLevel - prevLevel
		if isIncreasing && diff < 0 || !isIncreasing && diff > 0 {
			return uLevel
		}
		unsignedDiff := int(math.Abs(float64(diff)))
		if !isDiffCorrect(unsignedDiff) {
			return uLevel
		}
		prevLevel = curLevel
	}

	return sLevel
}

func Day2() {
	byteArray, err := os.ReadFile("aoc2024/day2.input.txt")
	if err != nil {
		log.Fatal(err)
	}

	safeReports := processReportsPart1(string(byteArray))
	fmt.Printf("Part 1: %d\n", safeReports) // 624

	safeReports = processReportsPart2(string(byteArray))
	fmt.Printf("Part 2: %d\n", safeReports) // 658

}
