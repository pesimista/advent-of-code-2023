package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Part_1() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0

	lines := strings.Split(string(file), "\r\n")
	for i, line := range lines {
		index := strings.Index(line, ":")
		if index <= 0 {
			continue
		}

		rounds := strings.Split(line[index+1:], ";")

		possible := checkRounds(rounds)

		if possible {
			total += i + 1
		}
	}

	fmt.Println(total)
}

// 4 red, 18 green, 15 blue; 17 green, 18 blue, 9 red; 8 red, 14 green, 6 blue; 14 green, 12 blue, 2 red
func checkRounds(rounds []string) bool {
	for _, round := range rounds {
		if !checkCubes(round) {
			return false
		}
	}

	return true
}

// 4 red, 18 green, 15 blue
func checkCubes(round string) bool {
	cubes := strings.Split(round, ",")

	for _, cube := range cubes {
		cubes := strings.Split(cube, " ")

		amount, _ := strconv.Atoi(cubes[1])
		colour := cubes[2]

		fmt.Printf("%v %v, available %v \n", amount, colour, maxCubes[colour])

		if amount > maxCubes[colour] {
			return false
		}
	}

	return true
}
