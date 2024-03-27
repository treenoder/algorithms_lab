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
	sorted := qs(a, 0, len(a)-1)
	return fmt.Sprintf("%v", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sorted)), " "), "[]"))
}

func part(a []int, low, high int) ([]int, int) {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return a, i
}

func qs(a []int, low, high int) []int {
	if low < high {
		var p int
		a, p = part(a, low, high)
		a = qs(a, low, p-1)
		a = qs(a, p+1, high)
	}
	return a
}
