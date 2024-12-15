package main

import (
	"container/list"
	"fmt"
	"os"
	"slices"
)

func readFile() string {
	f, _ := os.ReadFile("day9/input.txt")
	return string(f)
}

func main() {
	blocks := readFile()
	buffer := makeIndividualBlocks(blocks)
	replaced := replaceBlocks(buffer)

	fmt.Println(replaced)
	for _, val := range replaced {
		if val == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(val)
		}
	}
	fmt.Println()
	// 6511178035564
	// 2562652351669
	var hashSum = 0

	for idx := range replaced {
		if buffer[idx] == -1 {
			continue
		}
		hashSum += replaced[idx] * idx

	}
	println(hashSum)
}

func replaceBlocks(nonSorted []int) []int {
	grouped := groupBlocks(nonSorted)
	printList(*grouped)

	i := grouped.Len() - 1
	for back := grouped.Back(); back != nil; back = back.Prev() {
		nonEmptyBlock := back.Value.([]int)
		fmt.Printf("Curr: %v\n", nonEmptyBlock)

		if !slices.Contains(nonEmptyBlock, -1) {
			j := 0
			for front := grouped.Front(); front != nil && j < i; front = front.Next() {
				emptyBlock := front.Value.([]int)
				if emptyBlock[0] == -1 && len(nonEmptyBlock) <= len(emptyBlock) {
					hasChanged := replaceItems(emptyBlock, nonEmptyBlock)
					fmt.Println(hasChanged)
					if hasChanged && slices.Contains(emptyBlock, -1) {
						emptyIdx := slices.Index(emptyBlock, -1)

						empty := emptyBlock[emptyIdx:]
						containsVal := emptyBlock[:emptyIdx]

						front.Value = containsVal
						test := grouped.InsertAfter(empty, front)
						fmt.Printf("Empy: %v\n", test)
					}

					break
				}
				j++
			}
		}
		i--
	}

	printList(*grouped)

	return flatten(*grouped)
}

func flatten(grouped list.List) []int {
	res := make([]int, 0)
	for e := grouped.Front(); e != nil; e = e.Next() {
		for _, block := range e.Value.([]int) {
			res = append(res, block)
		}
	}
	return res
}

func replaceItems(dst []int, source []int) bool {
	replacedItems := copy(dst, source)
	for i2 := range source {
		source[i2] = -1
	}
	return replacedItems > 0
}

func groupBlocks(nonSorted []int) *list.List {
	grouped := list.New()
	currentBlock := make([]int, 0)

	for idx := 1; idx < len(nonSorted); idx++ {
		currentBlock = append(currentBlock, nonSorted[idx-1])
		if nonSorted[idx-1] != nonSorted[idx] {
			grouped.PushBack(currentBlock)
			currentBlock = make([]int, 0)
		}
	}
	currentBlock = append(currentBlock, nonSorted[len(nonSorted)-1])
	grouped.PushBack(currentBlock)

	return grouped
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

func printList(lst list.List) {
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}
