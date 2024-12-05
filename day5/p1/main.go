package main

import (
	"bufio"
	"math"
	"os"
	slices "slices"
	"strconv"
	"strings"
)

type OrderingRules = map[int][]int

type Input struct {
	pages []Page
	paths [][]string
}

type Page struct {
	left  string
	right string
}

func readFile() Input {
	f, _ := os.Open("day5/input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanningPages := true
	paths := make([][]string, 0)
	pages := make([]Page, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			scanningPages = false
			continue
		}

		if scanningPages {
			str := strings.Split(text, "|")
			pages = append(pages, Page{str[0], str[1]})
		} else {
			paths = append(paths, strings.Split(text, ","))
		}
	}

	return Input{paths: paths, pages: pages}
}

func buildGraph(pages []Page) map[string][]string {
	graph := make(map[string][]string)

	for _, page := range pages {
		graph[page.left] = append(graph[page.left], page.right)
	}

	return graph
}

func main() {
	input := readFile()

	pageGraph := buildGraph(input.pages)

	sum := 0
	for _, path := range input.paths {
		shouldInclude := true
		for i := 0; i < len(path)-1; i++ {
			currentTraversal := pageGraph[path[i]]
			if !slices.Contains(currentTraversal, path[i+1]) {
				shouldInclude = false
				break
			}
		}

		if shouldInclude {
			midEl := path[(int(math.Ceil(float64(len(path) / 2))))]
			m, _ := strconv.Atoi(midEl)
			sum += m
		}
	}

	print(sum)

}
