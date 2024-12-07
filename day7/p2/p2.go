package main

import (
	aoc2024 "aoc-2024"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	answer  uint64
	numbers []uint64
}

func readFile() []Input {
	f, _ := os.Open("day7/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	inputs := make([]Input, 0)
	for scanner.Scan() {
		text := scanner.Text()
		rowStr := strings.Split(text, ": ")

		answer, _ := strconv.ParseInt(rowStr[0], 10, 64)
		numbers := aoc2024.StrToUInt64(strings.Fields(rowStr[1]))

		inputs = append(inputs, Input{answer: uint64(answer), numbers: numbers})
	}

	return inputs
}

func main() {
	inputs := readFile()

	var sum uint64 = 0
	for _, input := range inputs {
		if isValid(input.answer, input.numbers) {
			sum += input.answer
		}
	}

	println(sum)
}

func isValid(answer uint64, numbers []uint64) bool {
	if len(numbers) == 1 {
		return numbers[0] == answer
	}

	head := numbers[:len(numbers)-1]
	tail := numbers[len(numbers)-1]

	if answer%tail == 0 && isValid(answer/tail, head) {
		return true
	}
	if answer-tail >= 0 && isValid(answer-tail, head) {
		return true
	}

	strTail := strconv.FormatUint(tail, 10)
	strAnswer := strconv.FormatUint(answer, 10)

	if len(strAnswer) > len(strTail) && strings.HasSuffix(strAnswer, strTail) {
		strAnswer = strAnswer[:len(strAnswer)-len(strTail)]

		newAns, _ := strconv.ParseInt(strAnswer, 10, 64)
		return isValid(uint64(newAns), numbers[:len(numbers)-1])
	}

	return false
}
