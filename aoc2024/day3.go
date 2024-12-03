package aoc2024

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// --- Day 3: Mull It Over ---
// https://adventofcode.com/2024/day/3
func day3Part1(content string) int {
	var mulRe = regexp.MustCompile(`mul\(\d+,\d+\)`)
	sum := 0

	ops := mulRe.FindAllString(content, -1)

	var re = regexp.MustCompile(`mul\((?P<X>\d+),(?P<Y>\d+)\)`)
	for _, op := range ops {
		tokens := re.FindStringSubmatch(op)
		x, _ := strconv.Atoi(tokens[1])
		y, _ := strconv.Atoi(tokens[2])
		sum += x * y
	}
	return sum
}

func createInvalidRanges(dontReIndexes, doReIndexes [][]int, lastPos int) [][]int {
	var invalidRanges [][]int
	for i := 0; i < len(dontReIndexes); i++ {
		for j := 0; j < len(doReIndexes); j++ {
			if doReIndexes[j][0] > dontReIndexes[i][0] {
				invalidRanges = append(invalidRanges, []int{dontReIndexes[i][0], doReIndexes[j][0]})
				break
			}
			if j == len(doReIndexes)-1 {
				invalidRanges = append(invalidRanges, []int{dontReIndexes[i][0], lastPos})
			}
		}
	}
	return invalidRanges
}

func isInvalidRange(invalidRanges [][]int, index int) bool {
	for i := 0; i < len(invalidRanges); i++ {
		if index > invalidRanges[i][0] && index < invalidRanges[i][1] {
			return true
		}
	}
	return false
}

func day3Part2(content string) int {
	var doRe = regexp.MustCompile(`do\(\)`)
	var dontRe = regexp.MustCompile(`don\'t\(\)`)
	var mulRe = regexp.MustCompile(`mul\(\d+,\d+\)`)
	sum := 0

	doReIndexes := doRe.FindAllStringIndex(content, -1)
	dontReIndexes := dontRe.FindAllStringIndex(content, -1)
	mulReIndexes := mulRe.FindAllStringIndex(content, -1)

	ops := mulRe.FindAllString(content, -1)

	invalidRanges := createInvalidRanges(dontReIndexes, doReIndexes, len(content))

	var re = regexp.MustCompile(`mul\((?P<X>\d+),(?P<Y>\d+)\)`)

	for i, op := range ops {
		if isInvalidRange(invalidRanges, mulReIndexes[i][0]) {
			continue
		}
		tokens := re.FindStringSubmatch(op)
		x, _ := strconv.Atoi(tokens[1])
		y, _ := strconv.Atoi(tokens[2])
		sum += x * y
	}
	return sum
}

func Day3() {
	byteArray, err := os.ReadFile("aoc2024/day3.input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := day3Part1(string(byteArray))
	fmt.Printf("Part 1: %d\n", sum) // 160672468

	sum = day3Part2(string(byteArray))
	fmt.Printf("Part 2: %d\n", sum) // 84893551
}
