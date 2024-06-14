package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func evaluate(ex string) int {
	operators := []rune{}
	parts := []int{}
	start := 0

	for i, ch := range ex {
		if ch == '+' || ch == '-' {
			operators = append(operators, ch)
			part, _ := strconv.Atoi(ex[start:i])
			parts = append(parts, part)
			start = i + 1
		}
	}
	part, _ := strconv.Atoi(ex[start:])
	parts = append(parts, part)

	result := parts[0]
	for i, part := range parts[1:] {
		if operators[i] == '+' {
			result += part
		} else {
			result -= part
		}
	}
	return result
}

func solution(ex string) int {
	maxVal := math.MinInt
	for i := range ex {
		newEx := ex[:i] + ex[i+1:]
		if newEx == "" {
			continue
		}
		evl := evaluate(newEx)
		if evl > maxVal {
			maxVal = evl
		}
	}
	return maxVal
}

func main() {
	var ex string
	fmt.Scanln(&ex)
	result := solution(strings.ReplaceAll(ex, " ", ""))
	fmt.Println(result)
}
