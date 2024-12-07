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
	x int
	y int
}

type Player struct {
	location  Location
	direction Direction
}

func (l *Player) getNextDirection() {
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

func (l *Player) getNextLocation() (int, int) {
	switch l.direction {
	case Up:
		return l.location.x, l.location.y - 1
	case Down:
		return l.location.x, l.location.y + 1
	case Left:
		return l.location.x - 1, l.location.y
	case Right:
		return l.location.x + 1, l.location.y
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
	visitedObstacles := make(map[Location]bool)
	loopCount := 0

	for !isOutOfBounds(input, guardLocation.location.x, guardLocation.location.y) {
		fmt.Println(guardLocation)
		visitedLocations[Location{guardLocation.location.x, guardLocation.location.y}] = true

		x, y := guardLocation.getNextLocation()

		for input[y][x] == OBSTACLE {
			visitedObstacles[Location{x, y}] = true
			guardLocation.getNextDirection()
			x, y = guardLocation.getNextLocation()
		}
		move(&guardLocation.location, x, y)
		if obstacleCandidate(&guardLocation.location, visitedLocations, visitedObstacles) {
			loopCount++
		}

	}

	fmt.Println("----------------")

	println(len(visitedLocations) + 1)
	println(loopCount)
}

func obstacleCandidate(l *Location, visited map[Location]bool, obstacles map[Location]bool) bool {
	obstaclesIntersection := make(map[Location]bool)

	for k := range visited {
		if l.x+1 == k.x || l.y+1 == k.y || l.x-1 == k.x || l.y-1 == k.y {
			obstaclesIntersection[k] = true
		}
	}

	for k := range obstaclesIntersection {
		for vis, _ := range obstacles {
			if k.x+1 == vis.x || k.y+1 == vis.y || k.x-1 == vis.x || k.y-1 == vis.y {
				return true
			}
		}
	}

	return false
}

func move(location *Location, x int, y int) {
	location.x = x
	location.y = y
}

func isOutOfBounds(input [][]string, x int, y int) bool {
	return x-1 < 0 || x+1 > len(input[0])-1 || y-1 < 0 || y+1 > len(input)-1
}

func findStartingPoint(input [][]string) Player {
	for y, item := range input {
		for x := range item {
			if input[y][x] == "^" {
				return Player{Location{x, y}, Up}
			}
		}
	}
	panic("Could not find starting point")
}
