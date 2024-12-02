package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type InputPair struct {
	first  []int
	second []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() InputPair {
	f, err := os.Open("day1/p1input.txt")
	defer f.Close()

	check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	first := make([]int, 0)
	second := make([]int, 0)
	for scanner.Scan() {
		str := strings.Fields(scanner.Text())

		num, err := strconv.Atoi(str[0])
		num1, err := strconv.Atoi(str[1])
		check(err)

		first = append(first, num)
		second = append(second, num1)
	}

	return InputPair{first: first, second: second}
}

func main() {
	input := readFile()

	slices.Sort(input.first)
	slices.Sort(input.second)

	res := 0
	for i := range input.second {
		fmt.Printf("%d %d\n", input.second[i], input.first[i])
		if input.second[i] < input.first[i] {
			res += input.first[i] - input.second[i]
		} else {
			res += input.second[i] - input.first[i]
		}
	}

	println(res)
}
