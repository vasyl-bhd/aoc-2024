package main

import (
	aoc_2024 "aoc-2024"
	"fmt"
	"strconv"
	"strings"
)

func readFile() [][]string {
	return aoc_2024.ReadFileAsStringMatrix("day11/input.txt", " ")
}

func makeNewGeneration(input []string) []string {
	res := make([]string, 0)

	for _, val := range input {
		if val == "0" {
			res = append(res, "1")
		} else if len(val)%2 == 0 {
			firstHalf := val[:len(val)/2]
			secondHalf := val[len(val)/2:]
			for strings.HasPrefix(secondHalf, "0") && len(secondHalf) > 1 {
				secondHalf = secondHalf[1:]
			}

			res = append(res, firstHalf)
			res = append(res, secondHalf)
		} else {
			intVal := aoc_2024.SafeStrToInt(val)
			intVal = intVal * 2024
			res = append(res, strconv.FormatUint(intVal, 10))
		}
	}
	return res
}

func main() {
	input := readFile()[0]

	for range 25 {
		input = makeNewGeneration(input)
	}
	fmt.Printf("%v\n", len(input))

}
