package main

import (
	"fmt"
)

func main() {
	var n int64
	var k uint8
	fmt.Scan(&n, &k)
	for i := uint8(0); i < k; i++ {
		fmt.Println(n)
	}
}
