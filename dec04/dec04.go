package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Part One - function to check if a section is contained in another
func isContainedIn(s1, s2 string) bool {

	parsedS1 := strings.Split(s1, "-")
	parsedS2 := strings.Split(s2, "-")

	min1, _ := strconv.Atoi(parsedS1[0])
	max1, _ := strconv.Atoi(parsedS1[1])

	min2, _ := strconv.Atoi(parsedS2[0])
	max2, _ := strconv.Atoi(parsedS2[1])

	if min1 >= min2 && max1 <= max2 {
		return true
	}

	return false
}

// Part Two - function to check if a section overlaps another
func overlaps(s1, s2 string) bool {

	parsedS1 := strings.Split(s1, "-")
	parsedS2 := strings.Split(s2, "-")

	min1, _ := strconv.Atoi(parsedS1[0])
	max1, _ := strconv.Atoi(parsedS1[1])

	min2, _ := strconv.Atoi(parsedS2[0])
	max2, _ := strconv.Atoi(parsedS2[1])

	if min1 >= min2 && min1 <= max2 {
		return true
	}

	if max1 >= min2 && max1 <= max2 {
		return true
	}

	return false
}

func main() {
	// Format input
	dat, _ := os.ReadFile("./input04.txt")
	inputData := strings.Split(string(dat), "\r\n")

	var input [][]string

	for i := 0; i < len(inputData); i++ {
		input = append(input, strings.Split(inputData[i], ","))
	}

	// Part One
	var count int

	for i := 0; i < len(input); i++ {
		section1 := input[i][0]
		section2 := input[i][1]

		if isContainedIn(section1, section2) || isContainedIn(section2, section1) {
			count++
		}
	}
	fmt.Println("Part One Answer:", count)

	// Part Two
	count = 0

	for i := 0; i < len(input); i++ {
		section1 := input[i][0]
		section2 := input[i][1]

		if overlaps(section1, section2) || overlaps(section2, section1) {
			count++
		}
	}
	fmt.Println("Part Two Answer:", count)

}
