package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction string

const OBSTACLE = "#"

const (
	Up    Direction = "^"
	Down            = "v"
	Left            = "<"
	Right           = ">"
)

type Location struct {
	x         int
	y         int
	direction Direction
}

func (l *Location) getNextDirection() {
	switch l.direction {
	case Up:
		l.direction = Right
	case Right:
		l.direction = Down
	case Down:
		l.direction = Left
	case Left:
		l.direction = Up
	default:
		panic("invalid direction " + l.direction)

	}

}

func (l *Location) getNextLocation() (int, int) {
	switch l.direction {
	case Up:
		return l.x, l.y - 1
	case Down:
		return l.x, l.y + 1
	case Left:
		return l.x - 1, l.y
	case Right:
		return l.x + 1, l.y
	default:
		panic("nani?")
	}

}

func readFile() [][]string {
	file, _ := os.Open("day6/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([][]string, 0)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		input = append(input, row)
	}

	return input
}

func main() {
	input := readFile()

	guardLocation := findStartingPoint(input)
	visitedLocations := make(map[Location]bool)

	for !isOutOfBounds(input, guardLocation) {
		fmt.Println(guardLocation)
		visitedLocations[Location{guardLocation.x, guardLocation.y, ""}] = true

		x, y := guardLocation.getNextLocation()

		for input[y][x] == OBSTACLE {
			guardLocation.getNextDirection()
			x, y = guardLocation.getNextLocation()
		}
		move(&guardLocation, x, y)
	}

	println(len(visitedLocations) + 1)
}

func move(location *Location, x int, y int) {
	location.x = x
	location.y = y
}

func isOutOfBounds(input [][]string, location Location) bool {
	return location.x-1 < 0 || location.x+1 > len(input[0])-1 || location.y-1 < 0 || location.y+1 > len(input)-1
}

func findStartingPoint(input [][]string) Location {
	for y, item := range input {
		for x := range item {
			if input[y][x] == "^" {
				return Location{x, y, Up}
			}
		}
	}
	panic("Could not find starting point")
}
