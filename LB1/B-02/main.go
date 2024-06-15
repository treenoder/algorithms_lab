package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, nums := input()
	fmt.Println(solution(nums, n))
}

func solution(a []int, n int) int {
	elements := make(map[int]struct{})
	for i := 0; i < n; i++ {
		elements[a[i]] = struct{}{}
	}
	return len(elements)
}

func input() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return n, nums
}
