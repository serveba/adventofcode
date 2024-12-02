package aoc2024

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// day 1
// https://adventofcode.com/2024/day/1
// Historian Hysteria
func distanceSum(fileContent string) int {
	sum := 0
	leftList, rightList := extractSortedLists(fileContent)
	elements := len(leftList) // lists have the same length
	for i := 0; i < elements; i++ {
		sum += int(math.Abs(float64(leftList[i] - rightList[i])))
	}
	return sum
}

func extractLists(fileContent string) ([]int, []int) {
	// declare two dynamic arrays
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	// split the file content by new line
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		elements := strings.Split(line, "   ")
		left, _ := strconv.Atoi(elements[0])
		right, _ := strconv.Atoi(elements[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return leftList, rightList
}

func extractSortedLists(fileContent string) ([]int, []int) {
	leftList, rightList := extractLists(fileContent)

	// sort the lists in ascending order
	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList
}

func silimarityScores(fileContent string) int {
	similarityScore := 0
	leftList, rightList := extractSortedLists(fileContent)

	for i := 0; i < len(leftList); i++ {
		coef := 0
		for j := 0; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				coef += 1
			}
		}
		if coef > 0 {
			// fmt.Printf("%d appears %d times\n", leftList[i], coef)
			similarityScore += leftList[i] * coef
		}
	}

	return similarityScore
}

func Day1() {
	byteArray, err := os.ReadFile("aoc2024/day1.input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := distanceSum(string(byteArray))
	fmt.Printf("Part 1: %d\n", sum)

	similarityScore := silimarityScores(string(byteArray))
	fmt.Printf("Part 2: %d\n", similarityScore)
}
