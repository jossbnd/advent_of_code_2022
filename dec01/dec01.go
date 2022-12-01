package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input and store in an array
	dat, _ := os.ReadFile("./input01.txt")
	input := strings.Split(string(dat), "\r\n")

	// calculate and store sums
	var s int
	var sums []int

	for i := 0; i < len(input); i++ {
		n, _ := strconv.Atoi(input[i])
		s += n

		if n == 0 {
			sums = append(sums, s)
			s = 0
		}
	}

	// sort to get the maxs
	sort.Ints(sums)

	// Part One
	fmt.Println("Part One Answer:", sums[len(sums)-1])

	// Part Two
	fmt.Println("Part Two Answer:", sums[len(sums)-1]+sums[len(sums)-2]+sums[len(sums)-3])
}
