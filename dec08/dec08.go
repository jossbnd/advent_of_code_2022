package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isVisible(row int, col int, grid [][]int) bool {
	height := grid[row][col]

	// Check if tree on edges
	if row == 0 || row == len(grid)-1 || col == 0 || col == len(grid[0])-1 {
		return true
	}

	visibleUp := true
	for r := 0; r < row; r++ {
		if grid[r][col] >= height {
			visibleUp = false
		}
	}

	visibleDown := true
	for r := row + 1; r < len(grid); r++ {
		if grid[r][col] >= height {
			visibleDown = false
		}
	}

	visibleLeft := true
	for c := 0; c < col; c++ {
		if grid[row][c] >= height {
			visibleLeft = false
		}
	}

	visibleRight := true
	for c := col + 1; c < len(grid[0]); c++ {
		if grid[row][c] >= height {
			visibleRight = false
		}
	}

	if visibleUp || visibleDown || visibleLeft || visibleRight {
		return true
	} else {
		return false
	}
}

func calculateTreeScore(row int, col int, grid [][]int) int {
	height := grid[row][col]

	// Check if tree on edges
	if row == 0 || row == len(grid)-1 || col == 0 || col == len(grid[0])-1 {
		return 0
	}

	upTrees := 0
	for r := row - 1; r >= 0; r-- {
		if grid[r][col] >= height || r == 0 {
			upTrees = row - r
			break
		}
	}

	downTrees := 0
	for r := row + 1; r < len(grid); r++ {
		if grid[r][col] >= height || r == len(grid)-1 {
			downTrees = r - row
			break
		}
	}

	leftTrees := 0
	for c := col - 1; c >= 0; c-- {
		if grid[row][c] >= height || c == 0 {
			leftTrees = col - c
			break
		}
	}

	rightTrees := 0
	for c := col + 1; c < len(grid[0]); c++ {
		if grid[row][c] >= height || c == len(grid[0])-1 {
			rightTrees = c - col
			break
		}
	}

	return upTrees * downTrees * leftTrees * rightTrees
}

func main() {
	// Format input
	dat, _ := os.ReadFile("./input08.txt")
	input := strings.Split(string(dat), "\r\n")

	var grid [][]int
	for _, line := range input {
		rowString := strings.Split(line, "")
		var row []int
		for _, s := range rowString {
			n, _ := strconv.Atoi(s)
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	// Part One & Two
	var countVisibles int
	var maxTreeScore int

	for row := range grid {
		for col := range grid[0] {

			// Part One
			if isVisible(row, col, grid) {
				countVisibles++
			}

			// Part Two
			treeScore := calculateTreeScore(row, col, grid)
			if treeScore > maxTreeScore {
				maxTreeScore = treeScore
			}
		}
	}

	fmt.Println("Part One Answer:", countVisibles)
	fmt.Println("Part Two Answer:", maxTreeScore)

}
