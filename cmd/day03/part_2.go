package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part_2() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0

	lines := strings.Split(string(file), "\r\n")
	for rowIndex, line := range lines {
		for colIndex, char := range line {
			if char != '*' {
				continue
			}

			numbers := []int{}

			// check around
			for row := rowIndex - 1; row < rowIndex+2; row++ {
				for col := colIndex - 1; col < colIndex+2; col++ {
					cell := lines[row][col]

					if !isNumber(cell) {

						continue
					}

					figure, lastIndex := getNumber(lines[row], col)

					numbers = append(numbers, figure)
					col = lastIndex
				}
			}

			if len(numbers) != 2 {
				continue
			}

			total += numbers[0] * numbers[1]
		}
	}

	fmt.Println(total)
}

func getNumber(line string, index int) (int, int) {
	figure := string(line[index])

	i := index
	for {
		i--

		if i < 0 {
			break
		}

		char := line[i]
		if !isNumber(char) {
			break
		}

		figure = string(char) + figure
	}

	i = index
	for {
		i++

		if i >= len(line) {
			break
		}

		char := line[i]
		if !isNumber(char) {
			break
		}

		figure = figure + string(char)
	}

	number, _ := strconv.Atoi(figure)

	return number, i - 1
}

func isNumber(input byte) bool {
	return input >= '0' && input <= '9'
}
