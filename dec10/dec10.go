package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addSignalStrength(cycle int, X int) int {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return cycle * X
	}

	return 0
}

func draw(cycle int, spritePositions [3]int, rowNumber int) string {
	for _, position := range spritePositions {
		if (cycle-1)-(rowNumber*40) == position {
			return "#"
		}
	}
	return "."
}

func moveCRT(cycle int, rowNumber int) int {
	return (cycle - 1) - (rowNumber * 40)
}

func isEndOfRow(cycle int) bool {
	if cycle == 40 || cycle == 80 || cycle == 120 || cycle == 160 || cycle == 200 || cycle == 240 {
		return true
	}
	return false
}

func main() {
	// Format input
	dat, _ := os.ReadFile("./input10.txt")
	input := strings.Split(string(dat), "\r\n")

	var program [][]string
	for _, value := range input {
		program = append(program, strings.Split(value, " "))
	}

	// Part One
	var cycle int
	var V int
	var signalStrength int

	X := 1

	// Part Two
	var image [6][40]string
	var row [40]string
	var crtPosition int
	spritePositions := [3]int{0, 1, 2}
	rowNumber := 0

	for _, instruction := range program {
		command := instruction[0]

		if len(instruction) > 1 {
			V, _ = strconv.Atoi(instruction[1])
		} else {
			V = 0
		}

		if command == "noop" {
			cycle++
			signalStrength += addSignalStrength(cycle, X)

			// Part Two
			crtPosition = moveCRT(cycle, rowNumber)
			row[crtPosition] = draw(cycle, spritePositions, rowNumber)

			if isEndOfRow(cycle) {
				image[rowNumber] = row
				rowNumber++
			}

		}

		if command == "addx" {
			// 1st cyle of addx command
			cycle++
			// Part One
			signalStrength += addSignalStrength(cycle, X)

			// Part Two
			crtPosition = moveCRT(cycle, rowNumber)
			row[crtPosition] = draw(cycle, spritePositions, rowNumber)

			if isEndOfRow(cycle) {
				image[rowNumber] = row
				rowNumber++
			}

			// 2nd cyle of addx command
			cycle++
			// Part One
			signalStrength += addSignalStrength(cycle, X)

			// Part Two
			crtPosition = moveCRT(cycle, rowNumber)
			row[crtPosition] = draw(cycle, spritePositions, rowNumber)

			if isEndOfRow(cycle) {
				image[rowNumber] = row
				rowNumber++
			}

			// End of 2nd cycle
			X += V

			spritePositions[0] = X - 1
			spritePositions[1] = X
			spritePositions[2] = X + 1
		}

	}

	fmt.Println("Part One Answer:", signalStrength)

	fmt.Println("")

	fmt.Println("Part Two Answer:")
	for _, row := range image {
		fmt.Println(row)
	}

}
