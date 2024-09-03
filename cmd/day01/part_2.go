package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var dictionary = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Part_2() {
	content, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(content), "\r\n")

	re, err := regexp.Compile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)
	if err != err {
		fmt.Println(err)
		return
	}

	total := 0
	for _, line := range lines {

		matches := []string{}

		for i := 0; i < len(line); i++ {
			val := re.FindString(line[i:])

			if len(val) != 0 {
				matches = append(matches, wordIntoInt(val))
			}
		}

		if len(matches) == 0 {
			continue
		}

		first := matches[0]
		last := matches[len(matches)-1]

		valueString := fmt.Sprintf("%v%v", first, last)
		valueInt, err := strconv.Atoi(valueString)
		if err != err {
			fmt.Println(err)
			return
		}

		total += valueInt
	}

	fmt.Printf("Part 2 Total %d\n", total)
}

func wordIntoInt(key string) string {
	if match, err := regexp.MatchString(`\D`, key); match || err != nil {
		return dictionary[key]
	}

	return key
}
