package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Generate a map for supplies and index when procedure instructions begin
func generateSupplies(input []string) (map[int][]string, int) {
	var procStart int
	supplies := make(map[int][]string)

	for i := 0; i < len(input); i++ {
		if input[i][1:2] == "1" {
			procStart = i + 2
			break
		}
		for j := 0; j < len(input[i]); j++ {
			char := input[i][j : j+1]
			match, _ := regexp.MatchString("[A-Z]", char)
			if match {
				stack := (j + 3) / 4
				supplies[stack] = append(supplies[stack], char)
			}
		}
	}

	return supplies, procStart
}

func main() {
	// Format input
	dat, _ := os.ReadFile("./input05.txt")
	input := strings.Split(string(dat), "\r\n")

	nStacks := (len(input[0]) + 2) / 4
	supplies, procStart := generateSupplies(input)

	// Part One
	for i := procStart; i < len(input); i++ {

		line := strings.Split(input[i], " ")

		n, _ := strconv.Atoi(line[1])
		from, _ := strconv.Atoi(line[3])
		to, _ := strconv.Atoi(line[5])

		for j := 0; j < n; j++ {
			crate := supplies[from][0]
			supplies[from] = supplies[from][1:]
			supplies[to] = append([]string{crate}, supplies[to]...)
		}

	}

	fmt.Print("Part One Answer: ")
	for stack := 1; stack <= nStacks; stack++ {
		fmt.Print(supplies[stack][0])
	}
	fmt.Print("\r\n")

	// Part Two
	supplies, _ = generateSupplies(input) // regenerate supplies map

	for i := procStart; i < len(input); i++ {

		line := strings.Split(input[i], " ")

		n, _ := strconv.Atoi(line[1])
		from, _ := strconv.Atoi(line[3])
		to, _ := strconv.Atoi(line[5])

		var crates []string

		for j := 0; j < n; j++ {
			crate := supplies[from][j]
			crates = append(crates, crate)
		}

		supplies[from] = supplies[from][n:]
		supplies[to] = append(crates, supplies[to]...)
	}

	fmt.Print("Part Two Answer: ")
	for stack := 1; stack <= nStacks; stack++ {
		fmt.Print(supplies[stack][0])
	}

}
