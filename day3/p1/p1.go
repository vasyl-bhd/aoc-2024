package main

import (
	aoc_2024 "aoc-2024"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readInput() string {
	f, err := os.ReadFile("day3/input.txt")
	aoc_2024.Check(err)

	return string(f)
}

func main() {
	regex := regexp.MustCompile("mul\\((\\d+,\\d+)\\)")

	str := readInput()

	res := regex.FindAllStringSubmatch(str, -1)

	sum := 0
	for _, re := range res {
		nums := strings.Split(re[1], ",")
		fmt.Printf("res: %v\n", nums)
		numbers := aoc_2024.StringsToInts(nums)
		sum += numbers[0] * numbers[1]
	}

	println(sum)

}
