package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lexer struct {
	input  string
	active bool
	sum    int64
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func (l *Lexer) ScanMul() {
	for i := 0; i < len(l.input); i++ {
		if i+4 < len(l.input) && l.input[i:i+4] == "mul(" {
			i += 4
			numStr := ""

			for l.input[i] != ')' {
				if (l.input[i] >= '0' && l.input[i] <= '9') || (l.input[i] == ',' && len(numStr) > 0) {
					numStr += string(l.input[i])
					i++
					continue
				} else {
					numStr = ""
					break
				}
			}

			if numStr != "" {
				nums := strings.Split(numStr, ",")

				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])

				l.sum += int64(num1) * int64(num2)
			}

		} else {
			continue
		}
	}
}

func (l *Lexer) ScanMulPart2() {
	enabled := true
	for i := 0; i < len(l.input); i++ {
		if i+7 < len(l.input) && l.input[i:i+7] == "don't()" {
			enabled = false
			continue
		}

		if i+4 < len(l.input) && l.input[i:i+4] == "do()" {
			enabled = true
			continue
		}

		if !enabled {
			continue
		}

		if i+4 < len(l.input) && l.input[i:i+4] == "mul(" {
			i += 4
			numStr := ""

			for l.input[i] != ')' {
				if (l.input[i] >= '0' && l.input[i] <= '9') || (l.input[i] == ',' && len(numStr) > 0) {
					numStr += string(l.input[i])
					i++
					continue
				} else {
					numStr = ""
					break
				}
			}

			if numStr != "" {
				nums := strings.Split(numStr, ",")

				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])

				l.sum += int64(num1) * int64(num2)
			}

		} else {
			continue
		}
	}
}

func (l *Lexer) Sum() int64 {
	return l.sum
}

func getInput() string {
	data, err := os.ReadFile("in.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(data))
}

func part1() int64 {
	input := getInput()
	lexer := New(input)

	lexer.ScanMulPart2()

	return lexer.Sum()
}

func main() {
	fmt.Println(part1())
}
