package main

import (
	aoc_2024 "aoc-2024"
	"bufio"
	"os"
	"regexp"
)

func readFile() [][]rune {
	file, err := os.Open("day4/input.txt")
	aoc_2024.Check(err)

	scanner := bufio.NewScanner(file)

	lines := make([][]rune, 0)
	for scanner.Scan() {

		lines = append(lines, []rune(scanner.Text()))
	}

	return lines
}

func rotate90(in [][]rune) [][]rune {
	rows := len(in)
	cols := len(in[0])
	rotated := make([][]rune, cols)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][i] = in[i][j]
		}
	}

	return rotated
}

func rotate45(in [][]rune) [][]rune {
	rows := len(in)
	cols := len(in[0])

	newRows := rows + cols - 1

	rotated := make([][]rune, newRows)

	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[i+j][i] = in[i][j]
		}
	}
	return rotated
}

func rotate135(in [][]rune) [][]rune {
	rows := len(in)
	cols := len(in[0])
	rotated := make([][]rune, cols)
	for i := 0; i < rows; i++ {
		rotated[i] = make([]rune, rows)
		for j := 0; j < cols; j++ {
			rotated[i][j] = in[rows-1-j][i]
		}
	}
	return rotate45(rotated)
}

func calcXMAX(in [][]rune) int {
	result := 0
	for _, line := range in {
		result += calculateLine(line)
	}
	return result
}

func calculateLine(input []rune) int {
	result := 0
	r, _ := regexp.Compile(`XMAS|SAMX`)
	for i := 0; i < len(input)-3; i++ {
		newStr := string(input[i : i+4])
		if r.FindString(newStr) != "" {
			result++
		}
	}
	return result
}

func main() {
	str := readFile()

	count := 0

	count += calcXMAX(str)
	count += calcXMAX(rotate45(str))
	count += calcXMAX(rotate90(str))
	count += calcXMAX(rotate135(str))

	println(count)
}
