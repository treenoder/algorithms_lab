package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	k := input()
	fmt.Println(Solution(k))
	//for i := 1; i <= 20; i++ {
	//	println(Solution(i))
	//}
}

func Solution(index int) int64 {
	c := make([]int64, 0)
	j := int64(1)
	for i := int64(1); i <= int64(index) && len(c) < index; i++ {
		cube := i * i * i
		for ; j*j <= cube && len(c) < index; j++ {
			if j*j == cube {
				j++
				break
			}
			c = append(c, j*j)
		}
		if len(c) < index {
			c = append(c, cube)
		} else {
			break
		}
	}

	return c[index-1]
}

func input() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
