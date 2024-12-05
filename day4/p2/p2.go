package main

import (
	aoc_2024 "aoc-2024"
	"bufio"
	"os"
	"strings"
)

func readFile() [][]string {
	file, err := os.Open("day4/input.txt")
	aoc_2024.Check(err)

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	for scanner.Scan() {

		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	return lines
}

func main() {
	str := readFile()
	count := 0
	for i := 1; i < len(str)-1; i++ {
		for j := 1; j < len(str[i])-1; j++ {
			if str[i][j] == "A" {

				word := str[i+1][j+1] + str[i][j] + str[i-1][j-1]
				word2 := str[i+1][j-1] + str[i][j] + str[i-1][j+1]
				if (word == "MAS" || word == "SAM") && (word2 == "MAS" || word2 == "SAM") {
					println(word)
					println(word2)
					count++
				}
			}
		}
	}

	println(count)
}
