package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

var XmasDir = []Position{
	{x: 1, y: 0},   // right
	{x: 0, y: 1},   // down
	{x: -1, y: 0},  // left
	{x: 0, y: -1},  // up
	{x: 1, y: 1},   // right-down
	{x: 1, y: -1},  // right-up
	{x: -1, y: 1},  // left-down
	{x: -1, y: -1}, // left-up
}

var MasDir = []Position{
	{x: 1, y: 1},   // right-down
	{x: 1, y: -1},  // right-up
	{x: -1, y: 1},  // left-down
	{x: -1, y: -1}, // left-up
}

func getInput() string {
	data, err := os.ReadFile("in.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(data))
}

func getGrid(input string) [][]rune {
	grid := strings.Split(input, "\n")

	var res [][]rune
	for _, row := range grid {
		res = append(res, []rune(row))
	}

	return res
}

func countXMAS(grid [][]rune, word string, count *int, dir Position, index Position, lookFor int, maxX int, maxY int) {
	if lookFor == len(word) {
		*count++
		return
	}

	if index.x < 0 || index.x >= maxX || index.y < 0 || index.y >= maxY {
		return
	}

	if grid[index.y][index.x] == rune(word[lookFor]) {
		countXMAS(grid, word, count, dir, Position{index.x + dir.x, index.y + dir.y}, lookFor+1, maxX, maxY)
	}
}

func storeMAS(grid [][]rune, word string, list *map[Position]int, dir Position, index Position, lookFor int, maxX int, maxY int) {
	if lookFor == len(word) {
		(*list)[Position{index.x - 2*dir.x, index.y - 2*dir.y}]++
		return
	}

	if index.x < 0 || index.x >= maxX || index.y < 0 || index.y >= maxY {
		return
	}

	if grid[index.y][index.x] == rune(word[lookFor]) {
		storeMAS(grid, word, list, dir, Position{index.x + dir.x, index.y + dir.y}, lookFor+1, maxX, maxY)
	}
}

func part1() {
	input := getInput()
	grid := getGrid(input)

	count := 0
	maxX, maxY := len(grid[0]), len(grid)
	word := "XMAS"

	for y, row := range grid {
		for x, col := range row {
			if col == rune(word[0]) {
				for _, dir := range XmasDir {
					countXMAS(grid, word, &count, dir, Position{x, y}, 0, maxX, maxY)
				}
			}
		}
	}

	fmt.Println(count)
}

func part2() {
	input := getInput()
	grid := getGrid(input)

	count := 0
	list := make(map[Position]int)
	maxX, maxY := len(grid[0]), len(grid)
	word := "MAS"

	for y, row := range grid {
		for x, col := range row {
			if col == rune(word[0]) {
				for _, dir := range MasDir {
					storeMAS(grid, word, &list, dir, Position{x, y}, 0, maxX, maxY)
				}
			}
		}
	}

	for _, v := range list {
		if v >= 2 {
			count++
		}
	}

	fmt.Println(count)
}

func main() {
	part2()
}
