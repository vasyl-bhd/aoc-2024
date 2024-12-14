package aoc_2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFileAsIntMatrix(filename string, sep string) [][]int {
	f, _ := os.Open(filename)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	res := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, sep)
		res = append(res, StringsToInts(row))

	}

	return res
}

func ReadFileAsStringMatrix(filename string, sep string) [][]string {
	f, _ := os.Open(filename)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	res := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, sep)
		res = append(res, row)

	}

	return res
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func StringsToInts(ss []string) []int {
	res := make([]int, len(ss))
	for i, s := range ss {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res[i] = num
	}

	return res
}

func StrToUInt64(ss []string) []uint64 {
	res := make([]uint64, len(ss))
	for i, s := range ss {
		res[i] = SafeStrToUint64(s)
	}

	return res
}

func SafeStrToUint64(ss string) uint64 {
	num, err := strconv.ParseInt(ss, 10, 64)
	if err != nil {
		panic(err)
	}

	return uint64(num)
}

func SafeStrToInt(ss string) int {
	num, err := strconv.Atoi(ss)
	if err != nil {
		panic(err)
	}

	return num
}
