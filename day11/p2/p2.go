package main

import (
	aoc_2024 "aoc-2024"
	"fmt"
	"strconv"
	"time"
)

type Result struct {
	first  int
	second int
}

func readFile() map[int]int {
	input := aoc_2024.StringsToInts(aoc_2024.ReadFileAsStringMatrix("day11/input.txt", " ")[0])

	res := map[int]int{}
	for _, val := range input {
		res[val] = 1
	}
	return res
}

func processInput(val int) Result {
	if val == 0 {
		return Result{first: 1, second: -1}

	} else if digits := strconv.Itoa(val); len(digits)%2 == 0 {
		firstHalf, _ := strconv.Atoi(digits[:len(digits)/2])
		secondHalf, _ := strconv.Atoi(digits[len(digits)/2:])
		return Result{first: firstHalf, second: secondHalf}
	} else {
		return Result{first: val * 2024, second: -1}
	}
}

func makeNewGeneration(input map[int]int) map[int]int {
	//fmt.Printf("BEFORE: %v\n", input)
	res := map[int]int{}

	add := func(k int, count int) {
		if _, ok := res[k]; !ok {
			res[k] = 0
		}
		res[k] += count

	}

	for k, count := range input {
		result := processInput(k)
		add(result.first, count)

		if result.second != -1 {
			add(result.second, count)
		}
	}

	return res
}

func main() {
	input := readFile()

	start := time.Now()
	for range 75 {
		input = makeNewGeneration(input)
	}

	fmt.Println(time.Since(start))

	sum := 0
	for _, v := range input {
		sum += v
	}
	fmt.Println(sum)
}
