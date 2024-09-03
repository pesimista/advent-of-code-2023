package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bag map[string]int

func main() {
	file, err := os.ReadFile("./cmd/day02/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0

	lines := strings.Split(string(file), "\r\n")
	for _, line := range lines {
		index := strings.Index(line, ":")
		if index <= 0 {
			continue
		}

		rounds := strings.Split(line[index+1:], ";")

		bagPower := getBagPower(rounds)
		total += bagPower
	}

	fmt.Println(total)
}

// 4 red, 18 green, 15 blue; 17 green, 18 blue, 9 red; 8 red, 14 green, 6 blue; 14 green, 12 blue, 2 red
func getBagPower(rounds []string) int {
	bag := &Bag{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	for _, round := range rounds {
		updateCubesForRound(round, bag)
	}

	var power int
	for _, value := range *bag {
		if power == 0 {
			power = value

			continue
		}
		power *= value
	}

	fmt.Println(bag)
	return power
}

// 4 red, 18 green, 15 blue
func updateCubesForRound(round string, bag *Bag) {
	cubes := strings.Split(round, ",")

	for _, cube := range cubes {
		cubes := strings.Split(cube, " ")

		amount, _ := strconv.Atoi(cubes[1])
		colour := cubes[2]

		if (*bag)[colour] < amount {
			(*bag)[colour] = amount
		}
	}

}
