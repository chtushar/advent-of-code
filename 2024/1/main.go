package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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

func part1() int {
	lines := strings.Split(getInput(), "\n")
	n := len(lines)

	pairs := make([][]int, 2)

	for i := 0; i < 2; i++ {
		pairs[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		numStr := strings.Split(lines[i], "   ")
		if len(numStr) != 2 {
			continue
		}
		for j := 0; j < 2; j++ {
			num, err := strconv.Atoi(numStr[j])
			if err != nil {
				fmt.Println("Error converting to int:", err)
				os.Exit(1)
			}
			pairs[j][i] = num
		}
	}

	for i := 0; i < 2; i++ {
		sort.Ints(pairs[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		dist := pairs[0][i] - pairs[1][i]
		if dist < 0 {
			dist = -dist
		}
		sum += dist
	}

	return sum
}

func part2() int {
	lines := strings.Split(getInput(), "\n")
	n := len(lines)

	locationMap := make(map[int]int)

	for i := 0; i < n; i++ {
		numStr := strings.Split(lines[i], "   ")
		if len(numStr) != 2 {
			continue
		}
		num, err := strconv.Atoi(numStr[1])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			os.Exit(1)
		}

		if _, ok := locationMap[num]; ok {
			locationMap[num]++
		} else {
			locationMap[num] = 1
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		numStr := strings.Split(lines[i], "   ")
		if len(numStr) != 2 {
			continue
		}
		num, err := strconv.Atoi(numStr[0])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			os.Exit(1)
		}

		if _, ok := locationMap[num]; ok {
			sum += num * locationMap[num]
		}
	}

	return sum
}

func main() {
	sum := part2()

	fmt.Println(sum)
}
