package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input and store in an array
	dat, _ := os.ReadFile("./input02.txt")
	input := strings.Split(string(dat), "\r\n")

	// Part One
	scoreMap := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 6,
			"Z": 0,
		},
		"B": {
			"X": 0,
			"Y": 3,
			"Z": 6,
		},
		"C": {
			"X": 6,
			"Y": 0,
			"Z": 3,
		},
		"shape": {
			"X": 1,
			"Y": 2,
			"Z": 3,
		},
	}

	var totalScore int
	var round string
	var opponentPlay string
	var myPlay string

	for i := 0; i < len(input); i++ {
		round = input[i]
		opponentPlay = round[0:1]
		myPlay = round[2:]

		totalScore += scoreMap[opponentPlay][myPlay] + scoreMap["shape"][myPlay]
	}

	fmt.Println("Part One Answer:", totalScore)

	// Part Two
	myPlayMap := map[string]map[string]string{
		"A": {
			"X": "Z",
			"Y": "X",
			"Z": "Y",
		},
		"B": {
			"X": "X",
			"Y": "Y",
			"Z": "Z",
		},
		"C": {
			"X": "Y",
			"Y": "Z",
			"Z": "X",
		},
	}

	var roundEnd string
	totalScore = 0

	for i := 0; i < len(input); i++ {
		round = input[i]
		opponentPlay = round[0:1]
		roundEnd = round[2:]
		myPlay = myPlayMap[opponentPlay][roundEnd]

		totalScore += scoreMap[opponentPlay][myPlay] + scoreMap["shape"][myPlay]
	}

	fmt.Println("Part Two Answer:", totalScore)
}
