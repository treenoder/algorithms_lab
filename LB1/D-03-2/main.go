package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, nums := input()
	//nums := []int{5, 4, 3, 2, 1}
	fmt.Println(Solution(nums))
}

func input() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var nums []int
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return n, nums
}

func Solution(a []int) string {
	sorted := mergeSort(a)
	return fmt.Sprintf("%v", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sorted)), " "), "[]"))
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	slice := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			slice = append(slice, left[i])
			i++
		} else {
			slice = append(slice, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		slice = append(slice, left[i])
	}
	for ; j < len(right); j++ {
		slice = append(slice, right[j])
	}
	return slice
}
