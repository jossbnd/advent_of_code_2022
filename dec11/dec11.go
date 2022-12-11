package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func throw(monkey string, worryLevel int, receiver string, monkeys map[string]map[string][]string) {
	item := strconv.Itoa(worryLevel)

	monkeys[monkey]["items"] = monkeys[monkey]["items"][1:]
	monkeys[receiver]["items"] = append(monkeys[receiver]["items"], item)

}

func main() {
	// Format input
	dat, _ := os.ReadFile("./input11.txt")
	input := strings.Split(string(dat), "\r\n")

	monkeys := make(map[int][]uint64)
	monkey := 0

	var operationSign string
	var operationNumber uint64
	var old bool
	var test uint64
	var receiverTrue int
	var receiverFalse int

	var count [8]int

	lcm := uint64(1)

	for i := 0; i < 20; i++ {
		monkey = 0

		for index, line := range input {

			if i == 0 {
				if index == 1+monkey*7 {
					re := regexp.MustCompile(`[0-9]+`)
					items := re.FindAllString(line, -1)
					for _, val := range items {
						item, _ := strconv.ParseUint(val, 10, 64)
						monkeys[monkey] = append(monkeys[monkey], item)
					}

					continue
				}
			}

			if index == 2+monkey*7 {
				l := strings.Split(string(line), " ")
				operationSign = l[6]

				if l[7] == "old" {
					old = true
				} else {
					old = false
				}

				if !old {
					operationNumber, _ = strconv.ParseUint(l[7], 10, 64)
				}
				continue
			}

			if index == 3+monkey*7 {
				l := strings.Split(string(line), " ")
				test, _ = strconv.ParseUint(l[5], 10, 64)
				continue
			}

			if index == 4+monkey*7 {
				l := strings.Split(string(line), " ")
				receiverTrue, _ = strconv.Atoi(l[9])

				continue
			}

			if index == 5+monkey*7 {
				l := strings.Split(string(line), " ")
				receiverFalse, _ = strconv.Atoi(l[9])
				continue
			}

			if len(line) == 0 {
				for _, worryLevel := range monkeys[monkey] {
					count[monkey]++

					if operationSign == "+" {
						worryLevel += operationNumber
					} else if operationSign == "*" && !old {
						worryLevel *= operationNumber
					} else if operationSign == "*" && old {
						worryLevel *= worryLevel
					}

					worryLevel = worryLevel / 3
					// worryLevel = worryLevel % lcm

					var receiver int
					if worryLevel%test == 0 {
						receiver = receiverTrue
					} else {
						receiver = receiverFalse
					}

					monkeys[monkey] = monkeys[monkey][1:]
					monkeys[receiver] = append(monkeys[receiver], worryLevel)

				}

				if i == 0 {
					lcm *= test
				}
				monkey++
				continue
			}

		}

	}

	countSlice := count[:]
	sort.Ints(countSlice)
	fmt.Println("Part One Answer:", countSlice[len(countSlice)-1]*countSlice[len(countSlice)-2])
}
