package main

import (
	aoc2024 "aoc-2024"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
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

func removeIdx(s []int, i int) []int {
	r := make([]int, 0)
	r = append(r, s[:i]...)
	return append(r, s[i+1:]...)
}

func isCorrect(level []int) bool {
	if isCorrectArr(level) {
		return true
	}
	for i := range level {
		if isCorrectArr(removeIdx(level, i)) {
			return true
		}
	}
	return false
}

func main() {
	res := readFile()
	start := time.Now()

	var count = 0
	for _, level := range res {
		if isCorrect(level) {
			count++
		}

	}

	fmt.Printf("%v\n", time.Since(start))
	println(count)
}
