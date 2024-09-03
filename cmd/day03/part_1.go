package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part_1() {
	file, err := os.ReadFile("./cmd/day03/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\r\n")

	re := regexp.MustCompile(`^\d+`)
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for lineIndex, currentLine := range lines {
		for i := 0; i < len(currentLine); i++ {
			match := re.FindString(currentLine[i:])

			if len(match) == 0 {
				continue
			}

			skip := len(match) - 1

			isValid := validatePossition(lines, skip, lineIndex, i)
			fmt.Println()
			i += skip

			if !isValid {
				continue
			}

			// fmt.Println(match)
			value, err := strconv.Atoi(match)
			if err != nil {
				fmt.Println(err)
				return
			}

			total += value
		}
	}

	fmt.Printf("total: %d\n", total)
}

func validatePossition(lines []string, steps, rowIndex, colIndex int) bool {
	maxRow := len(lines)
	re := regexp.MustCompile(`[^\d\.\s]`)

	for y := -1; y <= 1; y++ {
		row := rowIndex + y

		if row < 0 || row >= maxRow {
			continue
		}

		for x := -1; x <= steps+1; x++ {
			col := colIndex + x
			maxCol := len(lines[row])

			if col < 0 || col >= maxCol {
				continue
			}

			fmt.Printf("row %v col %v | ", row, col)
			cell := lines[row][col]

			if re.MatchString(string(cell)) {
				return true
			}
		}
	}

	return false
}
