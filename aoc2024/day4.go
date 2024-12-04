package aoc2024

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// --- Day 4: Ceres Search ---
// https://adventofcode.com/2024/day/4
func day4Part1(content string, size int) int {
	xmasCount := 0
	matrix := createXmasMatrix(content, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] == 'X' {
				xmasCount += countXmasFromX(matrix, i, j, size)
			}
		}
	}

	return xmasCount
}

func countXmasFromX(m [][]rune, x, y, size int) int {
	count := 0
	// horizontal right XMAS
	if y+3 < size && m[x][y+1] == 'M' && m[x][y+2] == 'A' && m[x][y+3] == 'S' {
		// fmt.Printf("([%d,%d],[%d,%d],[%d,%d],[%d,%d]) - horizontal right.\n",
		// 	x, y, x, y+1, x, y+2, x, y+3)
		count++
	}

	// horizontal left XMAS
	if y-3 >= 0 && m[x][y-1] == 'M' && m[x][y-2] == 'A' && m[x][y-3] == 'S' {
		// fmt.Printf("([%d,%d],[%d,%d],[%d,%d],[%d,%d]) - horizontal left.\n",
		// 	x, y, x, y-1, x, y-2, x, y-3)
		count++
	}

	// vertical down XMAS
	if x+3 < size && m[x+1][y] == 'M' && m[x+2][y] == 'A' && m[x+3][y] == 'S' {
		// fmt.Printf("([%d,%d],[%d,%d],[%d,%d],[%d,%d]) - vertical down.\n",
		// 	x, y, x+1, y, x+2, y, x+3, y)
		count++
	}

	// vertical up XMAS
	if x-3 >= 0 && m[x-1][y] == 'M' && m[x-2][y] == 'A' && m[x-3][y] == 'S' {
		// fmt.Printf("([%d,%d],[%d,%d],[%d,%d],[%d,%d]) - vertical up.\n",
		// 	x, y, x-1, y, x-2, y, x-3, y)
		count++
	}

	// diagonal down right XMAS
	if x+3 < size && y+3 < size && m[x+1][y+1] == 'M' && m[x+2][y+2] == 'A' && m[x+3][y+3] == 'S' {
		// fmt.Printf("([%d,%d],[%d,%d],[%d,%d],[%d,%d]) - diagonal down right.\n",
		// 	x, y, x+1, y+1, x+2, y+2, x+3, y+3)
		count++
	}

	// diagonal down left XMAS
	if x+3 < size && y-3 >= 0 && m[x+1][y-1] == 'M' && m[x+2][y-2] == 'A' && m[x+3][y-3] == 'S' {
		// fmt.Printf("([%d,%d],[%d,%d],[%d,%d],[%d,%d]) - diagonal down left.\n",
		// 	x, y, x+1, y-1, x+2, y-2, x+3, y-3)
		count++
	}

	// diagonal up right XMAS
	if x-3 >= 0 && y+3 < size && m[x-1][y+1] == 'M' && m[x-2][y+2] == 'A' && m[x-3][y+3] == 'S' {
		// fmt.Printf("([%d,%d],[%d,%d],[%d,%d],[%d,%d]) - diagonal up right.\n",
		// 	x, y, x-1, y+1, x-2, y+2, x-3, y+3)
		count++
	}

	// diagonal up left XMAS
	if x-3 >= 0 && y-3 >= 0 && m[x-1][y-1] == 'M' && m[x-2][y-2] == 'A' && m[x-3][y-3] == 'S' {
		// fmt.Printf("([%d,%d],[%d,%d],[%d,%d],[%d,%d]) - diagonal up left.\n",
		// 	x, y, x-1, y-1, x-2, y-2, x-3, y-3)
		count++
	}

	return count
}

func createXmasMatrix(content string, size int) [][]rune {
	matrix := make([][]rune, size)
	for i := range matrix {
		matrix[i] = make([]rune, size)
	}

	// Populate the matrix with characters from content
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if i >= size {
			break
		}
		for j, char := range line {
			if j >= size {
				break
			}
			matrix[i][j] = char
		}
	}
	return matrix
}

func countXmasShapeFromX(m [][]rune, x, y, size int) int {
	count := 0

	// checking shape limits
	if x-1 >= 0 && y-1 >= 0 && x+1 < size && y+1 < size {
		// fmt.Printf("Pos [%d,%d]\n", x, y)

		// M.S
		// .A.
		// M.S
		if m[x-1][y-1] == 'M' && m[x-1][y+1] == 'S' && m[x+1][y-1] == 'M' && m[x+1][y+1] == 'S' {
			count++
			// fmt.Printf("XMAS shape 1 - MSAMS - [%d,%d]\n", x, y)
		}

		// S.S
		// .A.
		// M.M
		if m[x-1][y-1] == 'S' && m[x-1][y+1] == 'S' && m[x+1][y-1] == 'M' && m[x+1][y+1] == 'M' {
			count++
			// fmt.Printf("XMAS shape 2 - SSAMM- [%d,%d]\n", x, y)
		}

		// S.M
		// .A.
		// S.M
		if m[x-1][y-1] == 'S' && m[x-1][y+1] == 'M' && m[x+1][y-1] == 'S' && m[x+1][y+1] == 'M' {
			count++
			// fmt.Printf("XMAS shape 3 - SMASM - [%d,%d]\n", x, y)
		}

		// M.M
		// .A.
		// S.S
		if m[x-1][y-1] == 'M' && m[x-1][y+1] == 'M' && m[x+1][y-1] == 'S' && m[x+1][y+1] == 'S' {
			count++
			// fmt.Printf("XMAS shape 4 - MMASS [%d,%d]\n", x, y)
		}
	}

	return count
}

func day4Part2(content string, size int) int {
	xmasCount := 0
	matrix := createXmasMatrix(content, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] == 'A' {
				xmasCount += countXmasShapeFromX(matrix, i, j, size)
			}
		}
	}

	return xmasCount
}

func Day4() {
	byteArray, err := os.ReadFile("aoc2024/day4.input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 140
	sum := day4Part1(string(byteArray), 140) // 18 (test) & 2583
	fmt.Printf("Part 1: %d\n", sum)          //

	sum = day4Part2(string(byteArray), 140)
	fmt.Printf("Part 2: %d\n", sum) // 9 test
}
