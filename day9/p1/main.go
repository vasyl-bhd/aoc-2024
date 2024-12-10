package main

import (
	"os"
)

func readFile() string {
	f, _ := os.ReadFile("day9/input.txt")

	return string(f)
}

func main() {
	blocks := readFile()
	buffer := makeIndividualBlocks(blocks)
	replaceBlocks(buffer)

	var hashSum int64 = 0
	for idx := range buffer {
		if buffer[idx] != -1 {
			//fmt.Printf("%d * %d = %d| ", num, idx, int64(num)*int64(idx))
			hashSum += int64(buffer[idx]) * int64(idx)
		}

	}

	println(hashSum)
	//6471961544878
	//6471961544878

	// 95217
	// 95217

	// [0 0 9 9 8 1 1 1 8 8 8 2 7 7 7 3 3 3 6 4 4 6 5 5 5 5 6 6]
	// [0 0 9 9 8 1 1 1 8 8 8 2 7 7 7 3 3 3 6 4 4 6 5 5 5 5 6 6]
}

func findLastIdx(strs []int, idx int) int {
	lastNumIdx := -1
	for i := len(strs) - 1; i > idx; i-- {
		if strs[i] != -1 {
			lastNumIdx = i
			break
		}
	}

	return lastNumIdx
}

func replaceBlocks(nonSorted []int) {

	for idx, block := range nonSorted {
		if block == -1 {
			lastNumIdx := findLastIdx(nonSorted, idx)
			if lastNumIdx == -1 {
				break
			}
			nonSorted[idx], nonSorted[lastNumIdx] = nonSorted[lastNumIdx], nonSorted[idx]
		}
	}

}

func makeIndividualBlocks(blocks string) []int {
	var indBlocks []int
	for i := range blocks {
		for range blocks[i] - 48 {
			if i%2 == 0 {
				indBlocks = append(indBlocks, i/2)
			} else {
				indBlocks = append(indBlocks, -1)
			}
		}
	}

	return indBlocks
}
