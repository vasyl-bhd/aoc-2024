package main

import (
	aoc2024 "aoc-2024"
	"image"
)

type AntennaLocations = map[string][]image.Point

func readFile() [][]string {
	return aoc2024.ReadFileAsStringMatrix("day8/input.txt", "")
}

func main() {
	input := readFile()
	antennas, bounds := makeGrid(&input)

	res := map[image.Point]bool{}

	for _, antenna := range antennas {
		for _, p1 := range antenna {
			for _, p2 := range antenna {
				if p1 == p2 {
					continue
				}

				if a := p2.Add(p2.Sub(p1)); bounds[a] {
					res[a] = true
				}
			}
		}
	}

	println(len(res))
}

func makeGrid(input *[][]string) (AntennaLocations, map[image.Point]bool) {
	antennas := make(AntennaLocations, len(*input))
	bounds := make(map[image.Point]bool, len(*input))
	for y, row := range *input {
		for x, char := range row {
			bounds[image.Point{X: x, Y: y}] = true
			if char != "." {
				antennas[char] = append(antennas[char], image.Point{X: x, Y: y})
			}
		}
	}

	return antennas, bounds
}
