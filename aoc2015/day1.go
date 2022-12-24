package aoc2015

import (
	"fmt"
	"io/ioutil"
	"log"
)

// day 1
// https://adventofcode.com/2015/day/1
// not quite lisp

func computeResultFloor(movements string) (int, int) {
	floor := 0
	basement := -1
	basement_reached := false
	for i, ch := range movements {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}

		if !basement_reached && floor == -1 {
			basement = i + 1
			basement_reached = true
		}
	}
	return floor, basement
}

func Day1() {
	byteArray, err := ioutil.ReadFile("aoc2015/day1.input.txt")
	if err != nil {
		log.Fatal(err)
	}

	finalFloor, basementReachedAt := computeResultFloor(string(byteArray))
	fmt.Printf("Final floor: %d\nBasement (-1) first reached at position %d\n",
		finalFloor, basementReachedAt)
}
