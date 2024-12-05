package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1() {
	type Rule struct {
		A int
		B int
	}
	data, err := os.ReadFile("in.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	splits := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	rulesStr := strings.Split(splits[0], "\n")
	updates := strings.Split(strings.TrimSpace(splits[1]), "\n")

	rulesMap := make(map[Rule]bool, len(rulesStr))

	for _, r := range rulesStr {
		numsStr := strings.Split(r, "|")

		a, _ := strconv.Atoi(numsStr[0])
		b, _ := strconv.Atoi(numsStr[1])

		rulesMap[Rule{A: a, B: b}] = true
	}

	sum := 0
	for _, u := range updates {
		numsStr := strings.Split(u, ",")
		var nums []int

		for _, n := range numsStr {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		ok := true
		for i, a := range nums {
			for j, b := range nums {
				if j < i {
					if !rulesMap[Rule{A: b, B: a}] {
						ok = false
						break
					}
				}
			}
		}

		if ok {
			sum += nums[len(nums)/2]
		}

	}
	fmt.Println(sum)
}

func part2() {
	type Rule struct {
		A int
		B int
	}
	data, err := os.ReadFile("in.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	splits := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	rulesStr := strings.Split(splits[0], "\n")
	updates := strings.Split(strings.TrimSpace(splits[1]), "\n")

	rulesMap := make(map[Rule]int, len(rulesStr))

	for _, r := range rulesStr {
		numsStr := strings.Split(r, "|")

		a, _ := strconv.Atoi(numsStr[0])
		b, _ := strconv.Atoi(numsStr[1])

		rulesMap[Rule{A: a, B: b}] = 1
		rulesMap[Rule{A: b, B: a}] = -1
	}

	sum := 0
	for _, u := range updates {
		numsStr := strings.Split(u, ",")
		var nums []int

		for _, n := range numsStr {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		ok := true
		for i, a := range nums {
			for j, b := range nums {
				if j < i {
					if rulesMap[Rule{A: b, B: a}] != 1 {
						ok = false
						break
					}
				}
			}
		}

		if !ok {
			sorted := make([]int, len(nums))
			sorted = slices.Clone(nums)
			slices.SortFunc(sorted, func(a, b int) int {
				if value, ok := rulesMap[Rule{A: a, B: b}]; ok {
					return value
				}
				return 0
			})

			sum += sorted[len(sorted)/2]
		}
	}
	fmt.Println(sum)
}

func main() {
	part1()
}
