package aoc2015

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// day 2
// https://adventofcode.com/2015/day/2
// --- Day 2: I Was Told There Would Be No Math ---

func processSquareFeet(present string) (int, int) {
	dimension := strings.Split(present, "x")

	l, err := strconv.Atoi(dimension[0])
	if err != nil {
		log.Fatal(err)
	}
	w, err := strconv.Atoi(dimension[1])
	if err != nil {
		log.Fatal(err)
	}
	h, err := strconv.Atoi(dimension[2])
	if err != nil {
		log.Fatal(err)
	}

	min := l * w

	if w*h < min {
		min = w * h
	}

	if h*l < min {
		min = h * l
	}

	surfaceArea := 2*l*w + 2*w*h + 2*h*l + min
	// LxWxH
	ribbon := l + l + w + w + l*w*h

	return surfaceArea, ribbon
}

func Day2() {
	// line by line
	readFile, err := os.Open("aoc2015/day2.input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalSquareFeet := 0
	totalRibbon := 0

	for fileScanner.Scan() {
		input := strings.TrimSpace(fileScanner.Text())

		squareFeet, ribbon := processSquareFeet(input)
		totalSquareFeet += squareFeet
		totalRibbon += ribbon
	}

	readFile.Close()
	fmt.Printf("Total square feet and ribbon: %d, %d\n",
		totalSquareFeet, totalRibbon)
}
