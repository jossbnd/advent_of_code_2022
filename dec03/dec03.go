package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input and store in an array
	dat, _ := os.ReadFile("./input03.txt")
	input := strings.Split(string(dat), "\r\n")

	// Part One
	var item string
	var rucksack string
	var errors []string

	// find errors
	for i := 0; i < len(input); i++ {
		var inventory = make(map[string]int)
		rucksack = input[i]
		for j := 0; j < len(rucksack); j++ {
			item = rucksack[j : j+1]
			if j < len(rucksack)/2 {
				inventory[item] = 1
			} else if j >= len(rucksack)/2 && inventory[item] == 1 {
				inventory[item]++
				errors = append(errors, item)
			}
		}
	}

	// sum priorities
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sum int

	for i := 0; i < len(errors); i++ {
		sum += strings.Index(alphabet, errors[i]) + 1
	}

	fmt.Println("Part One Answer:", sum)

	// Part two
	var badges []string

	// find badges
	for i := 0; i < len(input); i += 3 {
		var inventory = make(map[string]int)
		for k := 0; k < 3; k++ {
			rucksack = input[i+k]
			for j := 0; j < len(rucksack); j++ {
				item = rucksack[j : j+1]
				if inventory[item] == k {
					inventory[item]++
				}
				if inventory[item] == 3 {
					inventory[item]++
					badges = append(badges, item)
				}
			}
		}
	}

	// sum priorities
	sum = 0
	for i := 0; i < len(badges); i++ {
		sum += strings.Index(alphabet, badges[i]) + 1
	}

	fmt.Println("Part Two Answer:", sum)
}
