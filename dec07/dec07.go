package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func addSizeToDirectoriesPaths(fileSize int, path []string, directories map[string]int) {
	for i := range path {
		directoryPath := strings.Join(path[0:i+1], "/")
		directories[directoryPath] += fileSize
	}
}

func sortDirectoriesPaths(directories map[string]int) []string {
	var sortedDirectoriesPaths []string

	for key := range directories {
		sortedDirectoriesPaths = append(sortedDirectoriesPaths, key)
	}

	sort.SliceStable(sortedDirectoriesPaths, func(i, j int) bool {
		return directories[sortedDirectoriesPaths[i]] < directories[sortedDirectoriesPaths[j]]
	})

	return sortedDirectoriesPaths
}

func main() {
	// Format input
	dat, _ := os.ReadFile("./input07.txt")
	input := strings.Split(string(dat), "\r\n")

	var path []string
	directories := make(map[string]int)

	// Part One
	for _, value := range input {
		log := strings.Split(value, " ")

		if log[0] == "$" {
			if log[1] == "cd" && log[2] != ".." {
				directoryName := log[2]
				path = append(path, directoryName)
			} else if log[1] == "cd" && log[2] == ".." {
				path = path[0 : len(path)-1]
			}
		} else if log[0] != "$" && log[0] != "dir" {
			fileSize, _ := strconv.Atoi(log[0])
			addSizeToDirectoriesPaths(fileSize, path, directories)
		}
	}

	var partOneResult int
	for _, value := range directories {
		if value <= 100000 {
			partOneResult += value
		}
	}

	fmt.Println("Part One Answer:", partOneResult)

	// Part Two
	totalSize := directories["/"]
	unusedSpace := 70000000 - totalSize
	spaceToDelete := 30000000 - unusedSpace

	sortedDirectoriesPaths := sortDirectoriesPaths(directories)

	for _, directoryPath := range sortedDirectoriesPaths {
		directorySize := directories[directoryPath]
		if directorySize >= spaceToDelete {
			fmt.Println("Part Two Answer:", directorySize)
			break
		}
	}
}
