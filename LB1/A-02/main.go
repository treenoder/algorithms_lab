package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, nums := input()
	sort2(nums, n)
}

func sort1(a []int64, n int) {
	toPrint := make([]string, 0)
	isSorted := true
	for i := 1; i < n; i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			isSorted = false
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
		numsStr := strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), " "), "[]")
		toPrint = append(toPrint, numsStr)
	}
	if isSorted {
		fmt.Println("The array is already sorted")
		return
	}
	for _, s := range toPrint {
		fmt.Println(s)
	}
}

func sort2(a []int64, n int) {
	isSorted := true
	for i := 1; i < n; i++ {
		if a[i-1] > a[i] {
			isSorted = false
			for j := i; j > 0; j-- {
				if a[j-1] > a[j] {
					a[j], a[j-1] = a[j-1], a[j]
				} else {
					break
				}
			}
			for _, now := range a {
				fmt.Print(now, " ")
			}
			fmt.Println()
		}
	}
	if isSorted {
		fmt.Println("The array is already sorted")
	}
}

func input() (int, []int64) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		nums = append(nums, num)
	}
	return n, nums
}
