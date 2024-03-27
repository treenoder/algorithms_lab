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
	fmt.Println(Solution(nums))
}

func input() (int, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var nums []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums = append(nums, scanner.Text())
	}
	return n, nums
}

func Solution(a []string) string {
	for i := 1; i < len(a); i++ {
		if digitSum(a[i-1]) >= digitSum(a[i]) {
			for j := i; j > 0; j-- {
				if digitSum(a[j-1]) > digitSum(a[j]) || (digitSum(a[j-1]) == digitSum(a[j]) && toInt(a[j-1]) > toInt(a[j])) {
					a[j], a[j-1] = a[j-1], a[j]
				} else {
					break
				}
			}
		}
	}
	return strings.Join(a, " ")
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func digitSum(s string) int {
	sum := 0
	sign := ""
	for _, ch := range s {
		if ch == '-' {
			sign = "-"
			continue
		}
		digit, err := strconv.Atoi(sign + string(ch))
		if err != nil {
			panic(err)
		}
		sum += digit
		sign = ""
	}
	return sum
}
