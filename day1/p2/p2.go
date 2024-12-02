package main

import (
	aoc_2024 "aoc-2024"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputPair struct {
	first  []int
	second []int
}

func readFile() InputPair {
	f, err := os.Open("day1/input.txt")
	defer f.Close()

	aoc_2024.Check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	first := make([]int, 0)
	second := make([]int, 0)
	for scanner.Scan() {
		str := strings.Fields(scanner.Text())

		num, err := strconv.Atoi(str[0])
		num1, err := strconv.Atoi(str[1])
		aoc_2024.Check(err)

		first = append(first, num)
		second = append(second, num1)
	}

	return InputPair{first: first, second: second}
}

func main() {
	input := readFile()

	frequencyMap := make(map[int]int)
	for _, num := range input.second {
		frequencyMap[num] = frequencyMap[num] + 1
	}

	// Step 2: Create the result map for 'a' based on the frequency map
	occurrences := make(map[int]int)
	for _, num := range input.first {
		occurrences[num] = frequencyMap[num]
	}

	fmt.Println(frequencyMap)

	res := 0

	for k, v := range occurrences {
		res += v * k
	}

	println(res)

}
