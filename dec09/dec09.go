package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func moveHead(head []int, direction string) {

	switch direction {
	case "R":
		head[1]++
	case "D":
		head[0]++
	case "L":
		head[1]--
	case "U":
		head[0]--
	}

}

func moveTail(head []int, tail []int) {

	rHead := head[0]
	cHead := head[1]
	rTail := tail[0]
	cTail := tail[1]

	rDistance := math.Abs(float64(rHead - rTail))
	cDistance := math.Abs(float64(cHead - cTail))

	// if tail not far enough
	if rDistance < 2 && cDistance < 2 {
		return
	}

	moveDirectionR := float64(rHead-rTail) / 2
	moveDirectionC := float64(cHead-cTail) / 2

	if moveDirectionR < 0 {
		moveDirectionR = math.Floor(moveDirectionR)
	} else {
		moveDirectionR = math.Ceil(moveDirectionR)
	}

	if moveDirectionC < 0 {
		moveDirectionC = math.Floor(moveDirectionC)
	} else {
		moveDirectionC = math.Ceil(moveDirectionC)
	}

	tail[0] += int(moveDirectionR)
	tail[1] += int(moveDirectionC)
	return

}

func main() {
	// Format input
	dat, _ := os.ReadFile("./input09.txt")
	input := strings.Split(string(dat), "\r\n")

	var motions [][]string
	for _, value := range input {
		motions = append(motions, strings.Split(value, " "))
	}

	// Part One
	visitedTailOne := make(map[string]int)

	head := []int{0, 0}
	tailOne := []int{0, 0}

	// Part Two
	visitedTailNine := make(map[string]int)

	tailTwo := []int{0, 0}
	tailThree := []int{0, 0}
	tailFour := []int{0, 0}
	tailFive := []int{0, 0}
	tailSix := []int{0, 0}
	tailSeven := []int{0, 0}
	tailEight := []int{0, 0}
	tailNine := []int{0, 0}

	for _, motion := range motions {
		direction := motion[0]
		steps, _ := strconv.Atoi(motion[1])
		for i := 0; i < steps; i++ {

			// Part One
			moveHead(head, direction)
			moveTail(head, tailOne)

			// Part Two
			moveTail(tailOne, tailTwo)
			moveTail(tailTwo, tailThree)
			moveTail(tailThree, tailFour)
			moveTail(tailFour, tailFive)
			moveTail(tailFive, tailSix)
			moveTail(tailSix, tailSeven)
			moveTail(tailSeven, tailEight)
			moveTail(tailEight, tailNine)

			// Part One - Track Tail One
			keyCoordTailOne := strconv.Itoa(tailOne[0]) + " " + strconv.Itoa(tailOne[1])
			visitedTailOne[keyCoordTailOne]++

			// Part One - Track Tail Nine
			keyCoordTailNine := strconv.Itoa(tailNine[0]) + " " + strconv.Itoa(tailNine[1])
			visitedTailNine[keyCoordTailNine]++

		}
	}

	fmt.Println("Part One Answer:", len(visitedTailOne))
	fmt.Println("Part Two Answer:", len(visitedTailNine))

}
