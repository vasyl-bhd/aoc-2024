package main

import (
	aoc2024 "aoc-2024"
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
)

type Coords struct {
	x int
	y int
}

func readFile() [][]int {
	return aoc2024.ReadFileAsIntMatrix("day10/input.txt", "")
}

func main() {
	input := readFile()
	directions := []Coords{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
	}

	trailHeads := findTrailHeads(input)

	score := 0
	for _, head := range trailHeads {

		queue := llq.New()
		queue.Enqueue(head)

		for queue.Size() > 0 {
			var coord, _ = queue.Dequeue()
			current := coord.(Coords)

			for _, direction := range directions {
				newX, newY := current.x+direction.x, current.y+direction.y
				if newY >= 0 && newX >= 0 && newX < len(input[0]) && newY < len(input) {
					nextCoords := Coords{newX, newY}

					if input[newY][newX] == input[current.y][current.x]+1 {
						queue.Enqueue(nextCoords)
						if input[newY][newX] == 9 {
							score++
						}
					}
				}
			}

		}

	}
	println(score)

}

func findTrailHeads(input [][]int) []Coords {
	res := make([]Coords, 0)
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 0 {
				res = append(res, Coords{j, i})
			}
		}
	}

	return res
}
