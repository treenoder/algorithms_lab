package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, nums := input()
	//nums := []int{5, 4, 3, 2, 1, 5, 8, 9}
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
	if !isSorted(a) {
		shuffle(a)
		quickSort(a, 0, len(a)-1)
	}
	return fmt.Sprintf("%v", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), " "), "[]"))
}

func quickSort(a []int, low, high int) {
	if low < high {
		p := partition(a, low, high)
		quickSort(a, low, p-1)
		quickSort(a, p+1, high)
	}
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1

	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

func shuffle(a []int) {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
