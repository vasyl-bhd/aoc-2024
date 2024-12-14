package main

import (
	aoc_2024 "aoc-2024"
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

const MaxX = 101
const MaxY = 103

const SafezoneX = MaxX / 2
const SafezoneY = MaxY / 2

var Q1 = image.Rect(0, 0, SafezoneX, SafezoneY)
var Q2 = image.Rect(SafezoneX+1, 0, MaxX, SafezoneY)
var Q3 = image.Rect(0, SafezoneY+1, SafezoneX, MaxY)
var Q4 = image.Rect(SafezoneX+1, SafezoneY+1, MaxX, MaxY)

type RobotInfo struct {
	position image.Point
	velocity image.Point
}

func readFile() []RobotInfo {
	f, _ := os.Open("day14/input.txt")
	defer f.Close()

	regexp := regexp.MustCompile("p=(-?\\d+),(-?\\d+) v=(-?\\d+),(-?\\d+)")
	scanner := bufio.NewScanner(f)

	res := make([]RobotInfo, 0)
	for scanner.Scan() {
		txt := scanner.Text()
		match := regexp.FindStringSubmatch(txt)

		res = append(res, RobotInfo{
			position: image.Point{aoc_2024.SafeStrToInt(match[1]), aoc_2024.SafeStrToInt(match[2])},
			velocity: image.Point{X: aoc_2024.SafeStrToInt(match[3]), Y: aoc_2024.SafeStrToInt(match[4])},
		})

	}

	return res
}

func main() {
	input := readFile()
	fmt.Println("BEGIN ")
	for j := 0; j < 10000; j++ {
		for i := range input {
			robot := &input[i]

			newPos := robot.position.Add(robot.velocity)
			normalizePos(&newPos)
			robot.position = newPos
		}
		if j > 7500 {
			drawImage(input, j+1)
		}
	}

	quadrantCount := map[image.Rectangle]int{}
	for _, robot := range input {
		if robot.position.In(Q1) {
			quadrantCount[Q1]++
		}
		if robot.position.In(Q2) {
			quadrantCount[Q2]++
		}
		if robot.position.In(Q3) {
			quadrantCount[Q3]++
		}
		if robot.position.In(Q4) {
			quadrantCount[Q4]++
		}
	}

	fmt.Println(quadrantCount)

	sum := 1
	for _, v := range quadrantCount {
		fmt.Println(v)
		sum *= v
	}

	fmt.Println(sum)
}

func normalizePos(newPos *image.Point) {
	newPos.X = (newPos.X%MaxX + MaxX) % MaxX
	newPos.Y = (newPos.Y%MaxY + MaxY) % MaxY
}

func printRobotLocation(locations []RobotInfo) {
	board := make([][]string, MaxY)
	for i := range board {
		board[i] = make([]string, MaxX)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	for _, val := range locations {
		j := val.position.X
		//if j == MaxX {
		//	j = MaxX - 1
		//}
		i := val.position.Y
		//if i == MaxY {
		//	i = MaxY - 1
		//}
		if board[i][j] != "." {
			newVal := aoc_2024.SafeStrToInt(board[i][j]) + 1
			board[i][j] = strconv.Itoa(newVal)
		} else {
			board[i][j] = "1"
		}
	}

	for _, row := range board {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
	fmt.Println("--------------")
}

func drawImage(locs []RobotInfo, idx int) {
	img := image.NewRGBA(image.Rect(0, 0, MaxX, MaxY))
	blackCol := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: white}, image.Point{}, draw.Src)

	for _, robot := range locs {
		img.Set(robot.position.X, robot.position.Y, blackCol)
	}

	f, err := os.Create(fmt.Sprintf("day14/output/%d.png", idx))
	aoc_2024.Check(err)
	defer f.Close()

	png.Encode(f, img)
}
