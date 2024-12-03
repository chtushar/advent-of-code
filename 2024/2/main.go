package main

import (
	"fmt"
	"os"
	"slices"
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

func isValid(nums []int) bool {
	asc := nums[1] > nums[0]
	for i := 1; i < len(nums); i++ {
		if asc {
			diff := nums[i] - nums[i-1]
			if nums[i] < nums[i-1] || diff > 3 || diff < 1 {
				return false
			}
		} else {
			diff := nums[i-1] - nums[i]
			if nums[i] > nums[i-1] || diff > 3 || diff < 1 {
				return false
			}
		}
	}

	return true
}

func isTolerable(nums []int) bool {
	if isValid(nums) {
		return true
	}
	for i := range nums {
		if isValid(slices.Concat(nums[:i], nums[i+1:])) {
			return true
		}
	}
	return false
}

func part1() int {
	lines := strings.Split(getInput(), "\n")
	n := len(lines)
	count := 0
	for i := 0; i < n; i++ {
		numStr := strings.Split(lines[i], " ")
		nums := make([]int, len(numStr))
		for j := 0; j < len(numStr); j++ {
			num, err := strconv.Atoi(numStr[j])
			if err != nil {
				fmt.Println("Error converting to int:", err)
				os.Exit(1)
			}
			nums[j] = num
		}

		if isValid(nums) {
			count++
		}
	}

	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2() int {
	lines := strings.Split(getInput(), "\n")
	n := len(lines)

	count := 0
	for i := 0; i < n; i++ {
		numStr := strings.Split(lines[i], " ")
		nums := make([]int, len(numStr))

		for j := 0; j < len(numStr); j++ {
			num, err := strconv.Atoi(numStr[j])
			if err != nil {
				fmt.Println("Error converting to int:", err)
				os.Exit(1)
			}
			nums[j] = num
		}

		if isTolerable(nums) {
			count++
		}
	}

	return count
}

func main() {
	ans := part2()
	fmt.Println(ans)
}
