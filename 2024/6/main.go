package main

import (
	"fmt"
	"os"
	"strings"
)

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

type Block rune

const (
	Player Block = '^'
	Empty  Block = '.'
	Wall   Block = '#'
)

func getPlayerPositionInRow(row []rune) int {
	for c, block := range row {
		if block == rune(Player) {
			return c
		}
	}

	return -1
}

type Position struct {
	C int
	R int
}

func part1() {

	input := getInput()

	visited := make(map[Position]bool)
	player := Position{C: 0, R: 0}
	direction := Position{C: 0, R: -1}

	// find the player
	grid := getGrid(input)

	maxC := len(grid[0])
	maxR := len(grid)

	for r, row := range grid {
		if c := getPlayerPositionInRow(row); c != -1 {
			player.C = c
			player.R = r

			visited[player] = true
			break
		}
	}

	// move the player
	for {
		if player.C+direction.C < 0 || player.C+direction.C >= maxC || player.R+direction.R < 0 || player.R+direction.R >= maxR {
			break
		}

		if grid[player.R+direction.R][player.C+direction.C] == rune(Wall) { // turn right
			if direction.C == 0 && direction.R == -1 {
				direction = Position{C: 1, R: 0}
			} else if direction.C == 1 && direction.R == 0 {
				direction = Position{C: 0, R: 1}
			} else if direction.C == 0 && direction.R == 1 {
				direction = Position{C: -1, R: 0}
			} else if direction.C == -1 && direction.R == 0 {
				direction = Position{C: 0, R: -1}
			}
			continue
		}

		player.C += direction.C
		player.R += direction.R

		visited[player] = true
	}

	fmt.Println(len(visited))
}

func getPlayerPosition(grid [][]rune) Position {
	for r, row := range grid {
		if c := getPlayerPositionInRow(row); c != -1 {
			return Position{C: c, R: r}
		}
	}
	return Position{}
}

func part2() {
	type Visit struct {
		Position
		Dir Position
	}

	input := getInput()

	// find the player
	grid := getGrid(input)

	maxC := len(grid[0])
	maxR := len(grid)
	count := 0

	for r, row := range grid {
		for c, block := range row {
			if block == rune(Empty) {
				grid[r][c] = rune(Wall)

				visited := make(map[Visit]bool)
				player := Position{C: c, R: r}
				direction := Position{C: 0, R: -1}

				player = getPlayerPosition(grid)

				visited[Visit{Position: player, Dir: direction}] = true

				for {
					if player.C+direction.C < 0 || player.C+direction.C >= maxC || player.R+direction.R < 0 || player.R+direction.R >= maxR {
						break
					}

					if visited[Visit{Position: Position{C: player.C + direction.C, R: player.R + direction.R}, Dir: direction}] {
						count++
						break
					}

					if grid[player.R+direction.R][player.C+direction.C] == rune(Wall) { // turn right
						if direction.C == 0 && direction.R == -1 {
							direction = Position{C: 1, R: 0}
						} else if direction.C == 1 && direction.R == 0 {
							direction = Position{C: 0, R: 1}
						} else if direction.C == 0 && direction.R == 1 {
							direction = Position{C: -1, R: 0}
						} else if direction.C == -1 && direction.R == 0 {
							direction = Position{C: 0, R: -1}
						}
						continue
					}

					player.C += direction.C
					player.R += direction.R

					visited[Visit{Position: player, Dir: direction}] = true
				}

				grid[r][c] = rune(Empty)
			}

		}

	}

	fmt.Println(count)
}

func main() {
	part2()
}
