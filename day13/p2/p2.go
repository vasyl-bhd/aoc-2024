package main

import (
	aoc_2024 "aoc-2024"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Prize struct {
	x int
	y int
}

type Button struct {
	x int
	y int
}

type MachineInput struct {
	a     Button
	b     Button
	prize Prize
}

const A_COST = 3
const B_COST = 1
const PRIZE_MULTIPLYER = 10000000000000

var buttonPattern = regexp.MustCompile(`X.(\d+), Y.(\d+)`)

func readFile() []MachineInput {
	f, _ := os.ReadFile("day13/input.txt")
	input := string(f)

	res := make([]MachineInput, 0)
	strMachines := strings.Split(input, "\n\n")

	for _, strMachine := range strMachines {
		inputs := strings.Split(strMachine, "\n")
		buttonA := buttonPattern.FindStringSubmatch(inputs[0])
		buttonB := buttonPattern.FindStringSubmatch(inputs[1])
		prize := buttonPattern.FindStringSubmatch(inputs[2])

		res = append(res, MachineInput{
			a:     Button{aoc_2024.SafeStrToInt(buttonA[1]), aoc_2024.SafeStrToInt(buttonA[2])},
			b:     Button{aoc_2024.SafeStrToInt(buttonB[1]), aoc_2024.SafeStrToInt(buttonB[2])},
			prize: Prize{aoc_2024.SafeStrToInt(prize[1]) + PRIZE_MULTIPLYER, aoc_2024.SafeStrToInt(prize[2]) + PRIZE_MULTIPLYER},
		},
		)
	}

	return res
}

func main() {
	sum := 0
	for _, play := range readFile() {
		buttonA := play.a
		buttonB := play.b
		prize := play.prize
		div := play.a.x*play.b.y - play.a.y*play.b.x
		numA := (play.prize.x*play.b.y - play.prize.y*play.b.x) / div
		numB := (play.a.x*play.prize.y - play.a.y*play.prize.x) / div

		fmt.Println(prize)
		if numA*buttonA.x+numB*buttonB.x == prize.x && numA*buttonA.y+numB*buttonB.y == prize.y {
			sum += numA*A_COST + numB*B_COST
		}
	}

	fmt.Println(sum)
}
