package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums [3]int
	fmt.Scan(&nums[0], &nums[1], &nums[2])
	sort.Ints(nums[:])
	fmt.Println(nums[2] * nums[2])
	fmt.Println(nums[0] * nums[0])
}
