package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type Object string
type Paths string

const (
	Up    Paths = "^"
	Down        = "v"
	Left        = "<"
	Right       = ">"
)

const (
	Wall     Object = "#"
	Robot           = "@"
	Box             = "O"
	FreePath        = "."
)

type Input struct {
	gameMap       map[image.Point]Object
	startLocation image.Point
	paths         []Paths
}

func readFile() Input {
	f, _ := os.ReadFile("day15/input.txt")

	input := Input{}

	gameInfo := strings.Split(string(f), "\n\r")

	gameMap := parseMap(gameInfo[0], &input)
	input.gameMap = gameMap

	parsePaths(gameInfo[1], &input)

	return input
}

func parsePaths(s string, i *Input) {
	output := make([]Paths, 0)
	paths := strings.Split(strings.TrimSpace(s), "\n")
	for _, path := range paths {
		arr := strings.Split(path, "")

		for _, v := range arr {
			output = append(output, Paths(v))
		}
	}

	i.paths = output
}

func parseMap(pathMap string, input *Input) map[image.Point]Object {
	gameMap := map[image.Point]Object{}
	for y, line := range strings.Split(pathMap, "\n") {
		for x, str := range strings.Split(strings.TrimSpace(line), "") {
			if str != " " {
				if str == Robot {
					input.startLocation = image.Point{X: x, Y: y}
				}
				gameMap[image.Point{X: x, Y: y}] = Object(str)
			}
		}
	}
	return gameMap
}
func main() {
	input := readFile()

	fmt.Printf("%+v\n", input)
}
