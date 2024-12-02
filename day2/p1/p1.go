package main

import (
	aoc2024 "aoc-2024"
	"bufio"
	"math"
	"os"
	"strings"
)

func readFile() [][]int {
	res := make([][]int, 0)

	file, err := os.Open("day2/input.txt")
	defer file.Close()

	aoc2024.Check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		res = append(res, aoc2024.StringsToInts(line))
	}

	return res
}

func isCorrectArr(nums []int) bool {
	cmp := func(a, b int) bool { return a < b }
	for i := range nums {
		if nums[i]-nums[i+1] == 0 {
			continue
		}
		if nums[i]-nums[i+1] > 0 {
			cmp = func(a, b int) bool { return a > b }
		}
		break
	}

	for i := 1; i < len(nums); i++ {
		diff := math.Abs(float64(nums[i] - nums[i-1]))

		if diff < 1 || diff > 3 || !cmp(nums[i-1], nums[i]) {
			return false
		}
	}

	return true
}

func main() {
	res := readFile()

	var count = 0
	for _, level := range res {
		if isCorrectArr(level) {
			count++
		}
	}

	println(count)
}
