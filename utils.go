package aoc_2024

import (
	"strconv"
)

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
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		res[i] = uint64(num)
	}

	return res
}
