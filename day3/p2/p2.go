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
	regex := regexp.MustCompile("(mul\\((\\d+,\\d+)\\))|(don't\\(\\))|(do\\(\\))")

	str := readInput()

	res := regex.FindAllStringSubmatch(str, -1)

	fmt.Printf("%v\n", res)

	firstMatch := aoc_2024.StringsToInts(strings.Split(res[0][2], ","))

	sum := firstMatch[0] * firstMatch[1]
	dontFlag := "don't()"
	doFlag := "do()"
	isGoFlagProcessing := true
	for _, re := range res[1:] {
		switch re[0] {
		case doFlag:
			isGoFlagProcessing = true
		case dontFlag:
			isGoFlagProcessing = false
		}
		if strings.HasPrefix(re[0], "mul") && isGoFlagProcessing {
			nums := strings.Split(re[2], ",")
			fmt.Printf("res: %v\n", nums)
			numbers := aoc_2024.StringsToInts(nums)
			sum += numbers[0] * numbers[1]
		}

	}

	println(sum)

}
