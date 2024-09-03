package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part_1() {
	content, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(content), "\r\n")

	exp, err := regexp.Compile(`\d`)
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int
	for _, line := range lines {
		var first, last string

		matches := exp.FindAll([]byte(line), -1)

		if matches == nil {
			continue
		}

		first = string(matches[0])
		last = string(matches[len(matches)-1])

		valueString := fmt.Sprintf("%v%v", first, last)
		valueInt, err := strconv.Atoi(valueString)
		if err != nil {
			fmt.Println(err)
			return
		}

		total += valueInt
	}

	fmt.Printf("Part 1 Total: %d", total)
}
