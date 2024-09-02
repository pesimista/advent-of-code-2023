package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input/day1.txt")

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

	fmt.Printf("Total: %d", total)
}
